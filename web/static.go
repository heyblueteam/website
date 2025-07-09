package web

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)


// CachePolicy defines caching behavior for different file types
type CachePolicy struct {
	MaxAge    int  // Cache duration in seconds
	Immutable bool // Whether the resource is immutable
	Public    bool // Whether the cache is public
}

// CacheFileServer wraps http.FileServer with cache headers
type CacheFileServer struct {
	handler  http.Handler
	policies map[string]CachePolicy
}

// NewCacheFileServer creates a new cache-enabled file server
func NewCacheFileServer(root string) *CacheFileServer {
	// Define cache policies for different file types
	policies := map[string]CachePolicy{
		// Images - long cache (1 year) since they rarely change
		".png":  {MaxAge: 31536000, Public: true, Immutable: true},
		".jpg":  {MaxAge: 31536000, Public: true, Immutable: true},
		".jpeg": {MaxAge: 31536000, Public: true, Immutable: true},
		".webp": {MaxAge: 31536000, Public: true, Immutable: true},
		".svg":  {MaxAge: 31536000, Public: true, Immutable: true},
		".ico":  {MaxAge: 31536000, Public: true, Immutable: true},
		".gif":  {MaxAge: 31536000, Public: true, Immutable: true},

		// Fonts - long cache (1 year)
		".woff":  {MaxAge: 31536000, Public: true, Immutable: true},
		".woff2": {MaxAge: 31536000, Public: true, Immutable: true},
		".ttf":   {MaxAge: 31536000, Public: true, Immutable: true},
		".eot":   {MaxAge: 31536000, Public: true, Immutable: true},

		// JavaScript - medium cache (1 month) since you're actively developing
		".js": {MaxAge: 2592000, Public: true},

		// CSS - shorter cache (1 day) since you're actively updating the site
		".css": {MaxAge: 86400, Public: true},

		// Static assets - medium cache (1 week)
		".pdf": {MaxAge: 604800, Public: true},
		".zip": {MaxAge: 604800, Public: true},

		// Videos - long cache (1 month)
		".mp4":  {MaxAge: 2592000, Public: true},
		".webm": {MaxAge: 2592000, Public: true},
		".mov":  {MaxAge: 2592000, Public: true},

		// HTML - no cache for dynamic content
		".html": {MaxAge: 0, Public: false},
	}

	// Override cache policies in development
	if os.Getenv("ENV") == "development" {
		// Shorter cache durations for development
		policies[".css"] = CachePolicy{MaxAge: 300, Public: true} // 5 minutes
		policies[".js"] = CachePolicy{MaxAge: 300, Public: true}  // 5 minutes
	}

	return &CacheFileServer{
		handler:  http.FileServer(http.Dir(root)),
		policies: policies,
	}
}

// ServeHTTP implements http.Handler interface
func (cfs *CacheFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get file extension
	ext := strings.ToLower(filepath.Ext(r.URL.Path))

	// Apply cache policy if one exists for this file type
	if policy, exists := cfs.policies[ext]; exists {
		cfs.applyCacheHeaders(w, policy)
	} else {
		// Default: no cache for unknown file types
		cfs.applyNoCacheHeaders(w)
	}

	// Delegate to the underlying file server
	cfs.handler.ServeHTTP(w, r)
}

// applyCacheHeaders applies cache headers based on the policy
func (cfs *CacheFileServer) applyCacheHeaders(w http.ResponseWriter, policy CachePolicy) {
	if policy.MaxAge > 0 {
		// Build cache control header
		parts := []string{}

		if policy.Public {
			parts = append(parts, "public")
		} else {
			parts = append(parts, "private")
		}

		parts = append(parts, "max-age="+fmt.Sprintf("%d", policy.MaxAge))

		if policy.Immutable {
			parts = append(parts, "immutable")
		}

		w.Header().Set("Cache-Control", strings.Join(parts, ", "))

		// Add ETag for conditional requests
		w.Header().Set("ETag", `"blue-website-`+time.Now().Format("20060102")+`"`)
	} else {
		// No cache
		cfs.applyNoCacheHeaders(w)
	}
}

// applyNoCacheHeaders applies headers to prevent caching
func (cfs *CacheFileServer) applyNoCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

// GetCachePolicy returns the cache policy for a given file extension
func (cfs *CacheFileServer) GetCachePolicy(ext string) (CachePolicy, bool) {
	policy, exists := cfs.policies[strings.ToLower(ext)]
	return policy, exists
}

// SetCachePolicy allows runtime modification of cache policies
func (cfs *CacheFileServer) SetCachePolicy(ext string, policy CachePolicy) {
	cfs.policies[strings.ToLower(ext)] = policy
}

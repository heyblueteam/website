/**
 * Dynamic Highlight.js Loader
 * Loads Highlight.js on demand when navigating to pages with code blocks
 */
window.HighlightLoader = {
    loading: false,
    loaded: false,
    
    /**
     * Load Highlight.js dynamically
     * @returns {Promise} Promise that resolves when Highlight.js is loaded
     */
    async load() {
        // If already loaded, return immediately
        if (this.loaded || typeof hljs !== 'undefined') {
            this.loaded = true;
            return Promise.resolve();
        }
        
        // If already loading, wait for it
        if (this.loading) {
            return new Promise((resolve) => {
                const checkInterval = setInterval(() => {
                    if (this.loaded || typeof hljs !== 'undefined') {
                        clearInterval(checkInterval);
                        this.loaded = true;
                        resolve();
                    }
                }, 50);
            });
        }
        
        this.loading = true;
        
        return new Promise((resolve, reject) => {
            // Load CSS first
            const link = document.createElement('link');
            link.rel = 'stylesheet';
            link.href = '/css/highlight-github-dark.min.css';
            document.head.appendChild(link);
            
            // Then load JavaScript
            const script = document.createElement('script');
            script.src = '/js/vendor/highlight.min.js';
            script.async = true;
            
            script.onload = () => {
                this.loaded = true;
                this.loading = false;
                console.log('✅ Highlight.js loaded dynamically');
                resolve();
            };
            
            script.onerror = () => {
                this.loading = false;
                console.error('❌ Failed to load Highlight.js');
                reject(new Error('Failed to load Highlight.js'));
            };
            
            document.head.appendChild(script);
        });
    },
    
    /**
     * Initialize syntax highlighting on the page
     */
    async initialize() {
        // Load if not already loaded
        await this.load();
        
        // Initialize highlighting
        if (typeof hljs !== 'undefined') {
            hljs.highlightAll();
        }
    }
};
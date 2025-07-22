/**
 * Auth Cookie Utility
 * Reads and parses the blueAuthState cookie set by the Blue app
 * This allows the marketing website to show personalized content
 */
window.AuthCookie = {
    /**
     * Read and parse the auth state cookie
     * @returns {Object|null} Auth state object with isLoggedIn and firstName, or null if not found
     */
    read() {
        const cookieValue = this.getCookieValue('blueAuthState');
        if (!cookieValue) {
            return null;
        }

        try {
            // The cookie value is URL encoded, so decode it first
            const decodedValue = decodeURIComponent(cookieValue);
            const authState = JSON.parse(decodedValue);
            
            // Validate the structure
            if (typeof authState === 'object' && 
                authState !== null && 
                typeof authState.isLoggedIn === 'boolean') {
                return {
                    isLoggedIn: authState.isLoggedIn,
                    firstName: this.sanitizeFirstName(authState.firstName)
                };
            }
            
            return null;
        } catch (error) {
            console.warn('Failed to parse auth cookie:', error);
            return null;
        }
    },

    /**
     * Sanitize and format firstName for display
     * @param {string} firstName - Raw first name from cookie
     * @returns {string} Sanitized first name or 'there' as fallback
     */
    sanitizeFirstName(firstName) {
        // Handle empty, null, or whitespace-only names
        if (!firstName || firstName.trim() === '') {
            return 'there';
        }
        
        // Remove any HTML tags to prevent XSS
        let sanitized = firstName.replace(/<[^>]*>/g, '');
        
        // Trim whitespace
        sanitized = sanitized.trim();
        
        // If name is too long (more than 20 chars), use "there" instead
        if (sanitized.length > 20) {
            return 'there';
        }
        
        return sanitized;
    },

    /**
     * Get a cookie value by name
     * @param {string} name - Cookie name
     * @returns {string|null} Cookie value or null if not found
     */
    getCookieValue(name) {
        const cookies = document.cookie.split('; ');
        
        for (const cookie of cookies) {
            const [cookieName, cookieValue] = cookie.split('=');
            if (cookieName === name) {
                return cookieValue;
            }
        }
        
        return null;
    },

    /**
     * Check if user is logged in
     * @returns {boolean} True if logged in, false otherwise
     */
    isLoggedIn() {
        const authState = this.read();
        return authState?.isLoggedIn === true;
    },

    /**
     * Get user's first name
     * @returns {string} User's first name or 'there' as fallback
     */
    getFirstName() {
        const authState = this.read();
        return authState?.firstName || 'there';
    }
};
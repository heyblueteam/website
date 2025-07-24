/**
 * Copy Code Button Utilities
 * Handles adding copy buttons to code blocks and managing clipboard operations
 */
window.CopyCodeUtils = {
    /**
     * Initialize copy buttons on all code blocks
     */
    init() {
        this.addCopyButtons();
    },

    /**
     * Add copy buttons to all code blocks that don't already have them
     */
    addCopyButtons() {
        const codeBlocks = document.querySelectorAll('.prose pre:not([data-copy-added])');
        codeBlocks.forEach(pre => {
            pre.setAttribute('data-copy-added', 'true');
            
            // Create a wrapper div to contain the pre and button
            const wrapper = document.createElement('div');
            wrapper.style.position = 'relative';
            
            // Insert wrapper before pre
            pre.parentNode.insertBefore(wrapper, pre);
            
            // Move pre into wrapper
            wrapper.appendChild(pre);
            
            const copyBtn = document.createElement('button');
            copyBtn.className = 'code-copy-btn';
            copyBtn.innerHTML = '<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path></svg>';
            copyBtn.title = 'Copy code';
            
            copyBtn.addEventListener('click', async () => {
                const code = pre.querySelector('code');
                const text = code ? code.textContent : pre.textContent;
                
                try {
                    await navigator.clipboard.writeText(text);
                    copyBtn.innerHTML = '<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>';
                    copyBtn.classList.add('copied');
                    copyBtn.title = 'Copied!';
                    
                    setTimeout(() => {
                        copyBtn.innerHTML = '<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path></svg>';
                        copyBtn.classList.remove('copied');
                        copyBtn.title = 'Copy code';
                    }, 2000);
                } catch (err) {
                    console.error('Failed to copy:', err);
                    // Fallback to old clipboard API if needed
                    const textArea = document.createElement('textarea');
                    textArea.value = text;
                    textArea.style.position = 'fixed';
                    textArea.style.left = '-999999px';
                    document.body.appendChild(textArea);
                    textArea.select();
                    try {
                        document.execCommand('copy');
                        copyBtn.innerHTML = '<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>';
                        copyBtn.classList.add('copied');
                        copyBtn.title = 'Copied!';
                        
                        setTimeout(() => {
                            copyBtn.innerHTML = '<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path></svg>';
                            copyBtn.classList.remove('copied');
                            copyBtn.title = 'Copy code';
                        }, 2000);
                    } catch (fallbackErr) {
                        console.error('Fallback copy also failed:', fallbackErr);
                    } finally {
                        document.body.removeChild(textArea);
                    }
                }
            });
            
            // Append button to wrapper, not pre
            wrapper.appendChild(copyBtn);
        });
    }
};
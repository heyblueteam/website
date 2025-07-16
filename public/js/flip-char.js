/**
 * Flip Character Animation
 * Creates a glitchy character flip effect on hover
 * Used on the security page for interactive text animations
 */
window.flipChar = function(event) {
    const span = event.target;
    if (span.dataset.flipping === 'true' || !span.classList.contains('flip-char')) return;
    
    span.dataset.flipping = 'true';
    const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*';
    const original = span.textContent;
    let iterations = 0;
    
    // Add glitch effect class
    span.classList.add('glitching');
    
    const interval = setInterval(() => {
        span.textContent = chars[Math.floor(Math.random() * chars.length)];
        
        if (iterations > 10) {
            span.textContent = original;
            span.dataset.flipping = 'false';
            span.classList.remove('glitching');
            clearInterval(interval);
        }
        iterations++;
    }, 40);
}
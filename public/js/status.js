// Global status functionality for both full status page and API page
window.StatusUtils = {
    // Factory function for full status page (multiple services)
    createStatusPageData() {
        return {
            services: [],
            loading: true,
            error: null,
            
            async loadData() {
                try {
                    this.loading = true;
                    this.error = null;
                    
                    // Fetch both historical data and current status
                    const [historyResponse, currentResponse] = await Promise.all([
                        fetch('/api/status/history'),
                        fetch('/api/status/current')
                    ]);
                    
                    if (!historyResponse.ok) {
                        throw new Error(`HTTP ${historyResponse.status}: ${historyResponse.statusText}`);
                    }
                    
                    const historyData = await historyResponse.json();
                    const currentData = currentResponse.ok ? await currentResponse.json() : { services: [] };
                    
                    // Merge current status with historical data
                    this.services = (historyData.services || []).map(service => {
                        const currentService = currentData.services.find(c => c.name === service.name);
                        return {
                            ...service,
                            currentStatus: currentService ? currentService.status : 'unknown'
                        };
                    });
                    
                } catch (err) {
                    this.error = err.message;
                    console.error('Failed to load status data:', err);
                } finally {
                    this.loading = false;
                }
            },
            
            formatTooltip: StatusUtils.formatTooltip,
            getStatusColor: StatusUtils.getStatusColor,
            getStatusDotColor: StatusUtils.getStatusDotColor,
            getPulseColor: StatusUtils.getPulseColor
        }
    },

    // Factory function for API-only status
    createApiStatusData() {
        return {
            apiService: null,
            loading: true,
            error: null,
            
            async loadData() {
                try {
                    this.loading = true;
                    this.error = null;
                    
                    // Fetch both historical data and current status
                    const [historyResponse, currentResponse] = await Promise.all([
                        fetch('/api/status/history'),
                        fetch('/api/status/current')
                    ]);
                    
                    if (!historyResponse.ok) {
                        throw new Error(`HTTP ${historyResponse.status}: ${historyResponse.statusText}`);
                    }
                    
                    const historyData = await historyResponse.json();
                    const currentData = currentResponse.ok ? await currentResponse.json() : { services: [] };
                    
                    // Find API service specifically
                    const apiServiceHistory = (historyData.services || []).find(service => 
                        service.name.toLowerCase().includes('api') || service.name.toLowerCase().includes('graphql')
                    );
                    
                    if (apiServiceHistory) {
                        const currentService = currentData.services.find(c => 
                            c.name.toLowerCase().includes('api') || c.name.toLowerCase().includes('graphql')
                        );
                        
                        this.apiService = {
                            ...apiServiceHistory,
                            currentStatus: currentService ? currentService.status : 'operational'
                        };
                    }
                    
                } catch (err) {
                    this.error = err.message;
                    console.error('Failed to load API status data:', err);
                } finally {
                    this.loading = false;
                }
            },
            
            formatTooltip: StatusUtils.formatTooltip,
            getStatusColor: StatusUtils.getStatusColor,
            getStatusDotColor: StatusUtils.getStatusDotColor,
            getPulseColor: StatusUtils.getPulseColor
        }
    },

    // Shared utility functions
    formatTooltip(day) {
        const date = new Date(day.date);
        const options = { 
            weekday: 'short', 
            year: 'numeric', 
            month: 'long', 
            day: 'numeric' 
        };
        const formattedDate = date.toLocaleDateString('en-US', options);
        
        let statusText = 'Operational';
        if (day.status === 'degraded') {
            statusText = 'Degraded Performance';
        } else if (day.status === 'outage') {
            statusText = 'Major Outage';
        }
        
        return `${formattedDate}<br>${statusText} (${day.uptime.toFixed(1)}%)`;
    },
    
    getStatusColor(status) {
        switch (status) {
            case 'operational':
                return 'bg-green-500';
            case 'degraded':
                return 'bg-yellow-500';
            case 'outage':
                return 'bg-red-500';
            default:
                return 'bg-gray-300';
        }
    },
    
    getStatusDotColor(status) {
        switch (status) {
            case 'up':
            case 'operational':
                return 'bg-green-500';
            case 'down':
                return 'bg-red-500';
            case 'unknown':
            default:
                return 'bg-gray-400';
        }
    },
    
    getPulseColor(status) {
        switch (status) {
            case 'up':
            case 'operational':
                return 'bg-green-500';
            case 'down':
                return 'bg-red-500';
            case 'unknown':
            default:
                return 'bg-gray-400';
        }
    }
};
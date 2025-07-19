/**
 * Pricing Calculator Component
 * Interactive tool for calculating savings when switching to Blue
 */

// Define the function immediately to ensure it's available
window.pricingCalculator = function() {
  return {
    competitors: [
      { id: 1, name: 'Trello', logo: '/competitors/trello-icon.svg', cost: 12.50 },
      { id: 2, name: 'Basecamp', logo: '/competitors/basecamp-icon.svg', cost: 15 },
      { id: 3, name: 'Asana', logo: '/competitors/asana-icon.svg', cost: 30.49 },
      { id: 4, name: 'Monday', logo: '/competitors/monday-icon.svg', cost: 24 },
      { id: 5, name: 'Clickup', logo: '/competitors/clickup-icon.svg', cost: 19 },
      { id: 6, name: 'Notion', logo: '/competitors/notion-icon.svg', cost: 18 },
      { id: 7, name: 'Airtable', logo: '/competitors/airtable-icon.svg', cost: 54 },
      { id: 8, name: 'Slack', logo: '/competitors/slack-icon.svg', cost: 15 },
      { id: 9, name: 'Salesforce', logo: '/competitors/salesforce-icon.svg', cost: 25 },
      { id: 10, name: 'Microsoft', logo: '/competitors/msprojects-icon.svg', cost: 14 }
    ],
    selectedApps: [],
    companySize: 1,
    blueCostPerUser: 7,
    tooltipVisible: null,
    
    toggleApp(id) {
      const index = this.selectedApps.indexOf(id);
      if (index === -1) {
        this.selectedApps.push(id);
      } else {
        this.selectedApps.splice(index, 1);
      }
    },
    
    get annualSavings() {
      const competitorsCost = this.selectedApps.reduce((total, id) => {
        const competitor = this.competitors.find(c => c.id === id);
        return total + (competitor ? competitor.cost : 0);
      }, 0);
      
      const totalCompetitorsCost = competitorsCost * this.companySize * 12;
      const totalBlueCost = this.blueCostPerUser * this.companySize * 12;
      
      return totalCompetitorsCost - totalBlueCost;
    },
    
    get formattedSavings() {
      return Math.max(0, this.annualSavings).toLocaleString();
    }
  }
};
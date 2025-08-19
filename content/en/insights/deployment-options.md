---
title: "Deployment Options"
category: "Partner Resources"
description: "Complete guide to Blue's flexible deployment models - Cloud, Hybrid, On-Premise, and White-Label. Help customers choose the right architecture for their needs."
date: 2025-08-12
---


Unlike rigid enterprise platforms, Blue offers four deployment models with no feature penalties. Every customer gets the latest updates and full functionality, _regardless_ of how they deploy. This guide helps partners match customers with their ideal deployment architecture.

<video controls>
  <source src="https://media.blue.cc/options-for-deployment.mp4" type="video/mp4">
</video>

## Deployment Models Overview

| **Model** | **Infrastructure** | **Data Location** | **Setup Time** | **Ideal For** |
|-----------|-------------------|-------------------|----------------|---------------|
| **Cloud** | Blue-managed multi-tenant | AWS Europe | Immediate | Most enterprises |
| **Hybrid** | Dedicated Render instances + Customer AWS database | Customer-controlled database | 1-2 days | Data sovereignty needs |
| **On-Premise** | Customer infrastructure | Fully on-premise | 1-2 weeks | Government & high security |
| **White-Label** | Blue infrastructure, partner branding | Blue-managed | 1 week | MSPs & resellers |

## Cloud

### Architecture
- **Infrastructure:** Multi-tenant SaaS on AWS (database) and Render (application)
- **Region:** Primary hosting in Europe (GDPR compliant)
- **Availability:** 99.9% uptime SLA
- **Setup:** Immediate - same platform as self-service customers

### Benefits
- **Zero infrastructure management** - Blue handles everything
- **Instant provisioning** - Start immediately
- **Automatic updates** - Always on latest version
- **Full backup & DR** included
- **No size limitations** - From 10 to 10,000 users

### Pricing Structure
- Simple license fee per month
- No infrastructure costs
- No setup fees
- Predictable monthly billing

### Best For
- Companies wanting minimal IT involvement
- Organizations prioritizing speed to deployment
- Businesses without specific data residency requirements
- Growing companies needing to scale quickly

## Hybrid 

### Architecture
- **Application Layer:** Dedicated Render instances (Blue-managed)
- **Database:** Customer's AWS account in same region
- **Connection:** Direct secure connection between layers
- **Updates:** Blue-managed for application, customer-controlled for database

### Unique Advantages
- **Direct database access** - Query with MySQL, Power BI, or any BI tool
- **Lightning performance** - Dedicated resources just for them
- **Data sovereignty** - Complete control over data location
- **Best of both worlds** - Managed application, controlled data

### Technical Requirements
- AWS account in same region as Blue services
- Database administration capabilities
- Backup strategy for database (Blue can help setup)

### Pricing Structure
- Monthly license fee
- Base infrastructure cost
- 20% markup on Render hosting costs
- Usage-based scaling for resources

### Best For
- Organizations needing direct database access for BI/reporting
- Companies with data residency requirements
- Enterprises wanting performance guarantees
- Businesses with existing AWS infrastructure

## On-Premise

### Architecture
- **Deployment:** Docker containers provided by Blue
- **Infrastructure:** Customer's servers/cloud
- **Updates:** Customer-controlled (must stay current for support)
- **Mobile Apps:** White-labeled and pointed to on-premise instance

### Key Capabilities
- **Full functionality** including all enterprise features
- **Cloud integrations work** - Zapier, Make, etc. via custom API endpoints
- **Mobile apps included** - Fully branded iOS/Android apps
- **Complete data isolation** - Nothing leaves your infrastructure

### Technical Requirements
- Docker/Kubernetes environment
- SSL certificates and domain management
- Database administration expertise
- Network configuration for API access

### Support Model
- Local IT partner typically required
- Blue supports latest version only
- Update required for troubleshooting older versions
- Customer responsible for infrastructure issues

### Pricing Structure
- License fee
- One-time setup fee (varies by complexity)
- Optional professional services for deployment
- Annual support contract

### Best For
- Government agencies with strict data requirements
- Financial institutions with regulatory constraints
- Organizations in countries with data localization laws
- Companies with existing container infrastructure

## White-Label Cloud: Your Brand, Our Platform

### Architecture
- **Infrastructure:** Blue-managed cloud
- **Branding:** Complete white-label
- **Multi-tenancy:** Support for multiple client instances
- **Cascading white-label:** Your customers can white-label too

### MSP/Reseller Benefits
- **Set your own pricing** to end customers
- **Multiple branded instances** for different market segments
- **Your brand everywhere** - apps, emails, URLs, API
- **No infrastructure headaches** - Blue manages everything

### Commercial Model
- Volume-based licensing
- You set end-customer pricing
- Monthly billing from Blue
- Support can be Blue or partner-provided

### Best For
- Managed Service Providers (MSPs)
- Industry-specific solution providers
- Consulting firms with proprietary methodologies
- Regional resellers wanting local branding

## Migration Paths Between Models

| **Migration Path** | **Timeline** | **Downtime** | **Process** | **Cost** |
|-------------------|-------------|--------------|-------------|----------|
| **Cloud → Hybrid** | 2-3 days | 2-4 hours | Database migration to customer AWS | One-time migration fee |
| **Cloud → On-Premise** | 1-2 weeks | 4-6 hours | Full export, container deployment, data import | Setup fee + migration services |
| **Hybrid → On-Premise** | 1 week | 2-4 hours | Application container deployment, database migration | Setup fee + professional services |
| **Any Model → Cloud** | 1-2 days | 2-4 hours | Data export and import to Blue cloud | Minimal migration fee |

## Decision Framework

When speaking to a prospect, you can Figure out which deployment method makes the most sense for them based on these key factors:

1. **Data Control Requirements?**
   - No special requirements → Cloud
   - Need database access for BI → Hybrid
   - Strict data sovereignty → On-Premise

2. **IT Capabilities?**
   - Minimal IT resources → Cloud
   - AWS expertise → Hybrid
   - Full DevOps team → On-Premise

3. **Compliance Needs?**
   - Standard business → Cloud
   - GDPR with data residency → Hybrid
   - Government/military → On-Premise

4. **Business Model?**
   - Direct enterprise sale → Cloud/Hybrid/On-Premise
   - Reselling to multiple clients → White-Label

5. **Performance Requirements?**
   - Standard performance → Cloud
   - Guaranteed resources → Hybrid
   - Complete isolation → On-Premise


## The No-Penalty Promise

Regardless of deployment model:

- **Same features** - Every model gets full functionality
- **Same updates** - Latest version available to all
- **Same integrations** - API, Zapier, Make all work
- **Same mobile apps** - iOS/Android included
- **Same support quality** - Enterprise support for all

The only differences are where data lives and who manages infrastructure.

## Responsibilities by Model

| **Responsibility** | **Cloud** | **Hybrid** | **On-Premise** | **White-Label** |
|-------------------|-----------|------------|----------------|-----------------|
| Application Servers | Blue | Blue | Customer | Blue |
| Database | Blue | Customer | Customer | Blue |
| Backups | Blue | Shared | Customer | Blue |
| Updates | Blue | Blue/Customer | Customer | Blue |
| Security Patches | Blue | Shared | Customer | Blue |
| Scaling | Blue | Blue/Customer | Customer | Blue |
| Monitoring | Blue | Shared | Customer | Blue |
| SSL/Certificates | Blue | Blue | Customer | Blue |


---
title: FAQ om Blue Säkerhet
description: Detta är en lista över de mest frekvent ställda frågorna om säkerhetsprotokoll och praxis på Blue.
category: "FAQ"
date: 2024-07-19
---



Vår mission är att organisera världens arbete genom att bygga den bästa projektledningsplattformen på planeten.

Centralt för att uppnå denna mission är att säkerställa att vår plattform är säker, pålitlig och trovärdig. Vi förstår att för att vara din enda sanna källa, måste Blue skydda dina känsliga affärsdata mot externa hot, dataloss och driftstopp.

Detta innebär att vi tar säkerhet på allvar på Blue.

När vi tänker på säkerhet, överväger vi en helhetssyn som fokuserar på tre nyckelområden:

1.  **Infrastruktur & Nätverkssäkerhet**: Säkerställer att våra fysiska och virtuella system är skyddade mot externa hot och obehörig åtkomst.
2.  **Programvarusäkerhet**: Fokuserar på säkerheten i själva koden, inklusive säkra kodningspraxis, regelbundna kodgranskningar och sårbarhetshantering.
3.  **Plattformsäkerhet**: Inkluderar funktioner inom Blue, såsom [sofistikerade åtkomstkontroller](/platform/features/user-permissions), som säkerställer att projekt är privata som standard, och andra åtgärder för att skydda användardata och integritet.


## Hur skalbar är Blue?

Detta är en viktig fråga, eftersom du vill ha ett system som kan *växa* med dig. Du vill inte behöva byta din projekt- och processhanteringsplattform om sex eller tolv månader.

Vi väljer plattformsleverantörer med omsorg, för att säkerställa att de kan hantera de krävande arbetsbelastningarna från våra kunder. Vi använder molntjänster från några av världens främsta molnleverantörer som driver företag som [Spotify](https://spotify.com) och [Netflix](https://netflix.com), som har flera ordrar av magnitud mer trafik än vi har.

De huvudsakliga molnleverantörerna vi använder är:

- **[Cloudflare](https://cloudflare.com)**: Vi hanterar vår DNS (Domain Name Service) via Cloudflare samt vår marknadsföringswebbplats som körs på [Cloudflare Pages](https://pages.cloudflare.com/).
- **[Amazon Web Services](https://aws.amazon.com/)**: Vi använder AWS för vår databas, som är [Aurora](https://aws.amazon.com/rds/aurora/), för fillagring via [Simple Storage Service (S3)](https://aws.amazon.com/s3/), och även för att skicka e-post via [Simple Email Service (SES)](https://aws.amazon.com/ses/)
- **[Render](https://render.com)**: Vi använder Render för våra front-end servrar, applikations/API servrar, våra bakgrundstjänster, köhanteringssystem och Redis-databas. Intressant nog är Render faktiskt byggt *ovanpå* AWS!


## Hur säkra är filer i Blue?

Låt oss börja med datalagring. Våra filer är värd på [AWS S3](https://aws.amazon.com/s3/), som är världens mest populära molnobjektslagring med branschledande skalbarhet, datatillgänglighet, säkerhet och prestanda.

Vi har 99,99% filtillgänglighet och 99,999999999% hög hållbarhet.

Låt oss bryta ner vad detta betyder.

Tillgänglighet hänvisar till den tid som datan är operativ och tillgänglig. Den 99,99% filtillgängligheten innebär att vi kan förvänta oss att filer är otillgängliga i högst cirka 8,76 timmar per år.

Hållbarhet hänvisar till sannolikheten att data förblir intakt och okorrumperad över tid. Denna nivå av hållbarhet innebär att vi kan förvänta oss att förlora högst en fil av 10 miljarder uppladdade filer, tack vare omfattande redundans och datareplication över flera datacenter.

Vi använder [S3 Intelligent-Tiering](https://aws.amazon.com/s3/storage-classes/intelligent-tiering/) för att automatiskt flytta filer till olika lagringsklasser baserat på åtkomstfrekvens. Baserat på aktivitetsmönster från hundratusentals projekt, märker vi att de flesta filer nås i ett mönster som liknar en exponentiell backoff-kurva. Detta innebär att de flesta filer nås mycket frekvent under de första dagarna, och sedan snabbt nås mindre och mindre frekvent. Detta gör att vi kan flytta äldre filer till långsammare, men betydligt billigare, lagring utan att påverka användarupplevelsen på ett meningsfullt sätt.

Kostnadsbesparingarna för detta är betydande. S3 Standard-Infrequent Access (S3 Standard-IA) är cirka 1,84 gånger billigare än S3 Standard. Detta innebär att för varje dollar vi skulle ha spenderat på S3 Standard, spenderar vi endast cirka 54 cent på S3 Standard-IA för samma mängd lagrad data.

| Funktion                  | S3 Standard             | S3 Standard-IA       |
|--------------------------|-------------------------|-----------------------|
| Lagringskostnad          | $0.023 - $0.021 per GB  | $0.0125 per GB        |
| Begärningskostnad (PUT, etc.) | $0.005 per 1,000 requests | $0.01 per 1,000 requests |
| Begärningskostnad (GET)  | $0.0004 per 1,000 requests | $0.001 per 1,000 requests |
| Dataåterställningskostnad | $0.00 per GB            | $0.01 per GB          |


De filer du laddar upp genom Blue är krypterade både under överföring och i vila. Data som överförs till och från Amazon S3 är säkrad med [Transport Layer Security (TLS)](https://www.internetsociety.org/deploy360/tls/basics), vilket skyddar mot [avlyssning](https://en.wikipedia.org/wiki/Network_eavesdropping) och [man-in-the-middle-attacker](https://en.wikipedia.org/wiki/Man-in-the-middle_attack). För kryptering i vila använder Amazon S3 Server-Side Encryption (SSE-S3), som automatiskt krypterar alla nya uppladdningar med AES-256-kryptering, där Amazon hanterar krypteringsnycklarna. Detta säkerställer att dina data förblir säkra genom hela sin livscykel.

## Vad gäller icke-fildata?

Vår databas drivs av [AWS Aurora](https://aws.amazon.com/rds/aurora/), en modern relationsdatabasservice som säkerställer hög prestanda, tillgänglighet och säkerhet för dina data.

Data i Aurora är krypterad både under överföring och i vila. Vi använder SSL (AES-256) för att säkra anslutningar mellan din databasinstans och din applikation, vilket skyddar data under överföring. För kryptering i vila använder Aurora nycklar som hanteras genom AWS Key Management Service (KMS), vilket säkerställer att all lagrad data, inklusive automatiserade säkerhetskopior, snapshots och repliker, är krypterad och skyddad.

Aurora har ett distribuerat, fel-tolerant och självåterställande lagringssystem. Detta system är frikopplat från datorkapacitet och kan autoskala upp till 128 TiB per databasinstans. Data replikeras över tre [Tillgänglighetszoner](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) (AZ), vilket ger motståndskraft mot dataloss och säkerställer hög tillgänglighet. Vid en databas krasch minskar Aurora återställningstiderna till mindre än 60 sekunder, vilket säkerställer minimal störning.

Blue säkerhetskopierar kontinuerligt vår databas till Amazon S3, vilket möjliggör punkt-i-tid återställning. Detta innebär att vi kan återställa den blå huvuddatabasen till vilken specifik tid som helst inom de senaste fem minuterna, vilket säkerställer att dina data alltid är återställbara. Vi tar också regelbundna snapshots av databasen för längre säkerhetskopieringsperioder.

Som en helt hanterad tjänst automatiserar Aurora tidskrävande administrationsuppgifter som hårdvaruprovionering, databasinställning, patchning och säkerhetskopior. Detta minskar den operativa belastningen och säkerställer att vår databas alltid är uppdaterad med de senaste säkerhetsuppdateringarna och prestandaförbättringarna.

Om vi är mer effektiva kan vi överföra våra kostnadsbesparingar till våra kunder med vår [branschledande prissättning](/pricing).

Aurora följer olika branschstandarder som HIPAA, GDPR och SOC 2, vilket säkerställer att dina databehandlingspraxis uppfyller strikta regulatoriska krav. Regelbundna säkerhetsrevisioner och integration med [Amazon GuardDuty](https://aws.amazon.com/guardduty/) hjälper till att upptäcka och mildra potentiella säkerhetshot.

## Hur säkerställer Blue inloggningssäkerhet?

Blue använder [magiska länkar via e-post](https://documentation.blue.cc/user-management/magic-links) för att ge säker och bekväm åtkomst till ditt konto, vilket eliminerar behovet av traditionella lösenord.

Denna metod förbättrar säkerheten avsevärt genom att mildra vanliga hot som är förknippade med lösenordsbaserade inloggningar. Genom att eliminera lösenord skyddar magiska länkar mot phishing-attacker och stöld av lösenord, *eftersom det inte finns något lösenord att stjäla eller utnyttja.*

Varje magisk länk är giltig för endast en inloggningssession, vilket minskar risken för obehörig åtkomst. Dessutom går dessa länkar ut efter 15 minuter, vilket säkerställer att oanvända länkar inte kan utnyttjas, vilket ytterligare förbättrar säkerheten.

Bekvämligheten som erbjuds av magiska länkar är också anmärkningsvärd. Magiska länkar ger en problemfri inloggningsupplevelse, vilket gör att du kan få åtkomst till ditt konto *utan* att behöva komma ihåg komplexa lösenord.

Detta förenklar inte bara inloggningsprocessen utan förhindrar också säkerhetsbrott som inträffar när lösenord används om på flera tjänster. Många användare tenderar att använda samma lösenord på olika plattformar, vilket innebär att ett säkerhetsbrott på en tjänst kan kompromettera deras konton på andra tjänster, inklusive Blue. Genom att använda magiska länkar är Blues säkerhet inte beroende av säkerhetspraxis från andra tjänster, vilket ger ett mer robust och oberoende skydd för våra användare.

När du begär att logga in på ditt Blue-konto skickas en unik inloggnings-URL till din e-post. Genom att klicka på denna länk loggas du omedelbart in på ditt konto. Länken är utformad för att gå ut efter en enda användning eller efter 15 minuter, vilket som kommer först, vilket ger ett extra lager av säkerhet. Genom att använda magiska länkar säkerställer Blue att din inloggningsprocess är både säker och användarvänlig, vilket ger sinnesro och bekvämlighet.

## Hur kan jag kontrollera tillförlitligheten och drifttiden för Blue?

På Blue är vi engagerade i att upprätthålla en hög nivå av tillförlitlighet och transparens för våra användare. För att ge insyn i vår plattforms prestanda erbjuder vi en [dedikerad systemstatussida](https://status.blue.cc) som också är länkad från vår sidfot på varje sida av vår webbplats.

![](/insights/status-page.png)

Denna sida visar vår historiska drifttidsdata, vilket gör att du kan se hur konsekvent våra tjänster har varit tillgängliga över tid. Dessutom inkluderar statusidan detaljerade incidentrapporter, vilket ger transparens om eventuella tidigare problem, deras påverkan och de åtgärder vi har vidtagit för att lösa dem och förhindra framtida förekomster.
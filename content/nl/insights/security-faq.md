---
title: FAQ over Blue Beveiliging
description: Dit is een lijst van de meest gestelde vragen over de beveiligingsprotocollen en -praktijken bij Blue.
category: "FAQ"
date: 2024-07-19
---



Onze missie is om het werk in de wereld te organiseren door het beste projectmanagementplatform op de planeet te bouwen.

Centraal in het bereiken van deze missie staat het waarborgen dat ons platform veilig, betrouwbaar en vertrouwd is. We begrijpen dat Blue uw enige bron van waarheid moet zijn en uw gevoelige bedrijfsgegevens moet beschermen tegen externe bedreigingen, gegevensverlies en downtime.

Dit betekent dat we bij Blue beveiliging serieus nemen.

Wanneer we aan beveiliging denken, overwegen we een holistische benadering die zich richt op drie belangrijke gebieden:

1.  **Infrastructuur & Netwerkbeveiliging**: Zorgt ervoor dat onze fysieke en virtuele systemen beschermd zijn tegen externe bedreigingen en ongeautoriseerde toegang.
2.  **Softwarebeveiliging**: Richt zich op de beveiliging van de code zelf, inclusief veilige codering, regelmatige codebeoordelingen en kwetsbaarheidsbeheer.
3.  **Platformbeveiliging**: Bevat de functies binnen Blue, zoals [geavanceerde toegangscontroles](/platform/features/user-permissions), die ervoor zorgen dat projecten standaard privé zijn, en andere maatregelen om gebruikersgegevens en privacy te beschermen.


## Hoe schaalbaar is Blue?

Dit is een belangrijke vraag, aangezien u een systeem wilt dat met u kan *meegroeien*. U wilt niet dat u uw project- en procesmanagementplatform binnen zes of twaalf maanden moet vervangen.

We kiezen platformleveranciers zorgvuldig uit om ervoor te zorgen dat ze de veeleisende werklasten van onze klanten aankunnen. We gebruiken cloudservices van enkele van de beste cloudproviders ter wereld die bedrijven zoals [Spotify](https://spotify.com) en [Netflix](https://netflix.com) aandrijven, die verschillende ordes van grootte meer verkeer hebben dan wij.

De belangrijkste cloudproviders die we gebruiken zijn:

- **[Cloudflare](https://cloudflare.com)**: We beheren onze DNS (Domain Name Service) via Cloudflare, evenals onze marketingwebsite die draait op [Cloudflare Pages](https://pages.cloudflare.com/). 
- **[Amazon Web Services](https://aws.amazon.com/)**: We gebruiken AWS voor onze database, die [Aurora](https://aws.amazon.com/rds/aurora/) is, voor bestandsopslag via [Simple Storage Service (S3)](https://aws.amazon.com/s3/), en ook voor het verzenden van e-mails via [Simple Email Service (SES)](https://aws.amazon.com/ses/)
- **[Render](https://render.com)**: We gebruiken Render voor onze front-end servers, applicatie/API servers, onze achtergrondservices, queuesysteem en Redis-database. Interessant genoeg is Render eigenlijk *bovenop* AWS gebouwd! 


## Hoe veilig zijn bestanden in Blue? 

Laten we beginnen met gegevensopslag. Onze bestanden worden gehost op [AWS S3](https://aws.amazon.com/s3/), de meest populaire cloudobjectopslag ter wereld met toonaangevende schaalbaarheid, gegevensbeschikbaarheid, beveiliging en prestaties.

We hebben 99,99% bestandsbeschikbaarheid en 99,999999999% hoge duurzaamheid. 

Laten we opsplitsen wat dit betekent.

Beschikbaarheid verwijst naar de hoeveelheid tijd dat de gegevens operationeel en toegankelijk zijn. De 99,99% bestandsbeschikbaarheid betekent dat we kunnen verwachten dat bestanden niet langer dan ongeveer 8,76 uur per jaar onbeschikbaar zijn.

Duurzaamheid verwijst naar de waarschijnlijkheid dat gegevens intact en onbeschadigd blijven in de loop van de tijd. Dit niveau van duurzaamheid betekent dat we kunnen verwachten niet meer dan één bestand te verliezen van de 10 miljard geüploade bestanden, dankzij uitgebreide redundantie en gegevensreplicatie over meerdere datacenters.

We gebruiken [S3 Intelligent-Tiering](https://aws.amazon.com/s3/storage-classes/intelligent-tiering/) om bestanden automatisch naar verschillende opslagklassen te verplaatsen op basis van de frequentie van toegang. Op basis van de activiteitspatronen van honderden duizenden projecten merken we dat de meeste bestanden in een patroon worden geopend dat lijkt op een exponentiële backoff-curve. Dit betekent dat de meeste bestanden in de eerste paar dagen zeer frequent worden geopend en daarna snel steeds minder vaak. Dit stelt ons in staat om oudere bestanden naar langzamere, maar aanzienlijk goedkopere opslag te verplaatsen zonder de gebruikerservaring op een significante manier te beïnvloeden.

De kostenbesparingen hiervoor zijn aanzienlijk. S3 Standard-Infrequent Access (S3 Standard-IA) is ongeveer 1,84 keer goedkoper dan S3 Standard. Dit betekent dat we voor elke dollar die we aan S3 Standard zouden hebben uitgegeven, we slechts ongeveer 54 cent uitgeven aan S3 Standard-IA voor dezelfde hoeveelheid opgeslagen gegevens.

| Kenmerk                  | S3 Standard             | S3 Standard-IA       |
|--------------------------|-------------------------|-----------------------|
| Opslagkosten             | $0.023 - $0.021 per GB  | $0.0125 per GB        |
| Verzoekskosten (PUT, etc.) | $0.005 per 1.000 verzoeken | $0.01 per 1.000 verzoeken |
| Verzoekskosten (GET)       | $0.0004 per 1.000 verzoeken | $0.001 per 1.000 verzoeken |
| Gegevensherstelkosten      | $0.00 per GB            | $0.01 per GB          |


De bestanden die u via Blue uploadt, zijn zowel tijdens de overdracht als in rust versleuteld. Gegevens die naar en van Amazon S3 worden overgedragen, zijn beveiligd met [Transport Layer Security (TLS)](https://www.internetsociety.org/deploy360/tls/basics), wat bescherming biedt tegen [afluisteren](https://en.wikipedia.org/wiki/Network_eavesdropping) en [man-in-the-middle-aanvallen](https://en.wikipedia.org/wiki/Man-in-the-middle_attack). Voor versleuteling in rust gebruikt Amazon S3 Server-Side Encryption (SSE-S3), dat automatisch alle nieuwe uploads versleutelt met AES-256-versleuteling, waarbij Amazon de versleutelingssleutels beheert. Dit zorgt ervoor dat uw gegevens gedurende de hele levenscyclus veilig blijven.

## Wat betreft niet-bestandsgegevens? 

Onze database wordt aangedreven door [AWS Aurora](https://aws.amazon.com/rds/aurora/), een moderne relationele database-service die hoge prestaties, beschikbaarheid en beveiliging voor uw gegevens waarborgt.

Gegevens in Aurora zijn zowel tijdens de overdracht als in rust versleuteld. We gebruiken SSL (AES-256) om verbindingen tussen uw database-instantie en uw applicatie te beveiligen, waardoor gegevens tijdens de overdracht worden beschermd. Voor versleuteling in rust gebruikt Aurora sleutels die worden beheerd via de AWS Key Management Service (KMS), waardoor alle opgeslagen gegevens, inclusief geautomatiseerde back-ups, snapshots en replica's, zijn versleuteld en beschermd.

Aurora beschikt over een gedistribueerd, fouttolerant en zelfherstellend opslagsysteem. Dit systeem is ontkoppeld van rekenresources en kan automatisch opschalen tot 128 TiB per database-instantie. Gegevens worden gerepliceerd over drie [Beschikbaarheid Zones](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) (AZ's), wat veerkracht biedt tegen gegevensverlies en zorgt voor hoge beschikbaarheid. In het geval van een databasecrash vermindert Aurora de hersteltijden tot minder dan 60 seconden, wat zorgt voor minimale verstoring.

Blue maakt continu back-ups van onze database naar Amazon S3, waardoor punt-in-de-tijd herstel mogelijk is. Dit betekent dat we de blauwe masterdatabase naar elk specifiek tijdstip binnen de laatste vijf minuten kunnen herstellen, zodat uw gegevens altijd herstelbaar zijn. We maken ook regelmatig snapshots van de database voor langere back-upretentietijden. 

Als een volledig beheerde service automatiseert Aurora tijdrovende beheertaken zoals hardwarevoorziening, database-instelling, patching en back-ups. Dit vermindert de operationele overhead en zorgt ervoor dat onze database altijd up-to-date is met de nieuwste beveiligingspatches en prestatieverbeteringen. 

Als we efficiënter zijn, kunnen we onze kostenbesparingen doorgeven aan onze klanten met onze [toonaangevende prijzen](/pricing). 

Aurora voldoet aan verschillende industriestandaarden zoals HIPAA, GDPR en SOC 2, waardoor uw gegevensbeheerpraktijken voldoen aan strenge wettelijke vereisten. Regelmatige beveiligingsaudits en integratie met [Amazon GuardDuty](https://aws.amazon.com/guardduty/) helpen bij het detecteren en mitigeren van potentiële beveiligingsbedreigingen.

## Hoe zorgt Blue voor inlogbeveiliging?

Blue gebruikt [magische links via e-mail](https://documentation.blue.cc/user-management/magic-links) om veilige en handige toegang tot uw account te bieden, waardoor de noodzaak voor traditionele wachtwoorden vervalt.

Deze aanpak verbetert de beveiliging aanzienlijk door veelvoorkomende bedreigingen die gepaard gaan met wachtwoordgebaseerde inlogmethoden te verminderen. Door wachtwoorden te elimineren, beschermen magische links tegen phishing-aanvallen en wachtwoorddiefstal, *aangezien er geen wachtwoord is om te stelen of te exploiteren.* 

Elke magische link is slechts geldig voor één inlogsessie, waardoor het risico op ongeautoriseerde toegang wordt verminderd. Bovendien vervallen deze links na 15 minuten, zodat ongebruikte links niet kunnen worden geëxploiteerd, wat de beveiliging verder verbetert.

Het gemak dat magische links bieden is ook opmerkelijk. Magische links bieden een probleemloze inlogervaring, waardoor u toegang heeft tot uw account *zonder* de noodzaak om complexe wachtwoorden te onthouden.

Dit vereenvoudigt niet alleen het inlogproces, maar voorkomt ook beveiligingsinbreuken die optreden wanneer wachtwoorden op meerdere diensten worden hergebruikt. Veel gebruikers hebben de neiging om hetzelfde wachtwoord op verschillende platforms te gebruiken, wat betekent dat een beveiligingsinbreuk op één dienst hun accounts op andere diensten, inclusief Blue, kan compromitteren. Door magische links te gebruiken, is de beveiliging van Blue niet afhankelijk van de beveiligingspraktijken van andere diensten, wat een robuustere en onafhankelijke beschermingslaag voor onze gebruikers biedt.

Wanneer u vraagt om in te loggen op uw Blue-account, wordt er een unieke inlog-URL naar uw e-mail gestuurd. Door op deze link te klikken, wordt u onmiddellijk ingelogd op uw account. De link is ontworpen om te vervallen na één gebruik of na 15 minuten, afhankelijk van wat het eerst komt, wat een extra beveiligingslaag toevoegt. Door magische links te gebruiken, zorgt Blue ervoor dat uw inlogproces zowel veilig als gebruiksvriendelijk is, wat gemoedsrust en gemak biedt.

## Hoe kan ik de betrouwbaarheid en uptime van Blue controleren?

Bij Blue zijn we toegewijd aan het handhaven van een hoog niveau van betrouwbaarheid en transparantie voor onze gebruikers. Om inzicht te bieden in de prestaties van ons platform, bieden we een [gewijd systeemstatuspagina](https://status.blue.cc) die ook is gelinkt vanuit onze footer op elke pagina van onze website. 

![](/insights/status-page.png)

Deze pagina toont onze historische uptime-gegevens, zodat u kunt zien hoe consistent onze diensten in de loop van de tijd beschikbaar zijn geweest. Bovendien bevat de statuspagina gedetailleerde incidentrapporten, die transparantie bieden over eventuele eerdere problemen, hun impact en de stappen die we hebben genomen om ze op te lossen en toekomstige voorvallen te voorkomen.
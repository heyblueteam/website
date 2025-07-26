---
title:  Skapa Blues Anpassade Behörighetsmotor
description: Gå bakom kulisserna med Blue ingenjörsteam när de förklarar hur man bygger en AI-driven automatisk kategorisering och taggning funktion.
category: "Engineering"
date: 2024-07-25
---



Effektiv projekt- och processhantering är avgörande för organisationer av alla storlekar.

På Blue har [vi gjort det till vår mission](/about) att organisera världens arbete genom att bygga den bästa projektledningsplattformen på planeten—enkel, kraftfull, flexibel och prisvärd för alla.

Detta innebär att vår plattform måste anpassa sig till de unika behoven hos varje team. Idag är vi glada att dra tillbaka ridån på en av våra mest kraftfulla funktioner: Anpassade Behörigheter.

Projektledningsverktyg är ryggraden i moderna arbetsflöden, som rymmer känslig data, avgörande kommunikation och strategiska planer. Därför är förmågan att noggrant kontrollera åtkomst till denna information inte bara en lyx—det är en nödvändighet.

<video autoplay loop muted playsinline>
  <source src="/videos/user-roles.mp4" type="video/mp4">
</video>


Anpassade behörigheter spelar en kritisk roll i B2B SaaS-plattformar, särskilt i projektledningsverktyg, där balansen mellan samarbete och säkerhet kan avgöra ett projekts framgång.

Men här tar Blue en annan väg: **vi tror att företagsklassade funktioner inte bör reserveras för företagsstora budgetar.**

I en tid där AI möjliggör för små team att verka i oöverträffade skala, varför ska robust säkerhet och anpassning vara utom räckhåll?

I denna bakom kulisserna-titt kommer vi att utforska hur vi utvecklade vår funktion för Anpassade Behörigheter, utmanande status quo för SaaS-prissättning och förde kraftfulla, flexibla säkerhetsalternativ till företag av alla storlekar.

Oavsett om du är en startup med stora drömmar eller en etablerad aktör som vill optimera dina processer, kan anpassade behörigheter möjliggöra nya användningsfall som du aldrig ens visste var möjliga.

## Förstå Anpassade Användarbehörigheter

Innan vi dyker ner i vår resa att utveckla anpassade behörigheter för Blue, låt oss ta ett ögonblick för att förstå vad anpassade användarbehörigheter är och varför de är så avgörande i projektledningsprogram.

Anpassade användarbehörigheter avser möjligheten att skräddarsy åtkomsträttigheter för individuella användare eller grupper inom ett program. Istället för att förlita sig på fördefinierade roller med fasta uppsättningar av behörigheter, tillåter anpassade behörigheter administratörer att skapa mycket specifika åtkomstprofiler som perfekt överensstämmer med deras organisationsstruktur och arbetsflödesbehov.

I sammanhanget av projektledningsprogram som Blue inkluderar anpassade behörigheter:

* **Granulär åtkomstkontroll**: Bestämma vem som kan se, redigera eller ta bort specifika typer av projektdata.
* **Funktionsbaserade begränsningar**: Aktivera eller inaktivera vissa funktioner för specifika användare eller team.
* **Datakänslighetsnivåer**: Ställa in olika nivåer av åtkomst till känslig information inom projekt.
* **Arbetsflödes-specifika behörigheter**: Justera användarkapabiliteter med specifika steg eller aspekter av ditt projektarbetsflöde.

Vikten av anpassade behörigheter i projektledning kan inte överskattas:

* **Förbättrad säkerhet**: Genom att ge användare endast den åtkomst de behöver, minskar du risken för dataintrång eller obehöriga ändringar.
* **Förbättrad efterlevnad**: Anpassade behörigheter hjälper organisationer att uppfylla branschspecifika regleringskrav genom att kontrollera dataåtkomst.
* **Strömlinjeformad samarbete**: Team kan arbeta mer effektivt när varje medlem har rätt nivå av åtkomst för att utföra sin roll utan onödiga begränsningar eller överväldigande privilegier.
* **Flexibilitet för komplexa organisationer**: När företag växer och utvecklas, tillåter anpassade behörigheter programvaran att anpassa sig till förändrade organisationsstrukturer och processer.

## Komma till JA

[Vi har skrivit tidigare](/insights/value-proposition-blue) att varje funktion i Blue måste vara ett **hårt** JA innan vi beslutar att bygga den. Vi har inte lyxen av hundratals ingenjörer och att slösa tid och pengar på att bygga saker som kunderna inte behöver.

Och så var vägen till att implementera anpassade behörigheter i Blue inte en rak linje. Som många kraftfulla funktioner började det med ett klart behov från våra användare och utvecklades genom noggrant övervägande och planering.

I flera år hade våra kunder begärt mer granulär kontroll över användarbehörigheter. När organisationer av alla storlekar började hantera alltmer komplexa och känsliga projekt blev begränsningarna i vår standard rollbaserade åtkomstkontroll uppenbara.

Små startups som arbetar med externa kunder, medelstora företag med intrikata godkännandeprocesser och stora företag med strikta efterlevnadskrav uttryckte alla samma behov:

Mer flexibilitet i hanteringen av användartillgång.

Trots den tydliga efterfrågan tvekade vi initialt att dyka in i utvecklingen av anpassade behörigheter.

Varför?

Vi förstod den involverade komplexiteten!

Anpassade behörigheter berör varje del av ett projektledningssystem, från användargränssnittet ner till databasstrukturen. Vi visste att implementeringen av denna funktion skulle kräva betydande förändringar i vår kärnarkitektur och noggrant övervägande av prestandaimplikationer.

När vi granskade landskapet märkte vi att mycket få av våra konkurrenter hade försökt implementera en kraftfull anpassad behörighetsmotor som den våra kunder efterfrågade. De som gjorde det reserverade ofta den för sina högsta företagsplaner.

Det blev tydligt varför: utvecklingsinsatsen är betydande, och insatserna är höga.

Att implementera anpassade behörigheter felaktigt skulle kunna introducera kritiska buggar eller säkerhetsbrister, vilket potentiellt skulle kunna kompromettera hela systemet. Denna insikt underströk storleken på den utmaning vi övervägde.

### Utmana Status Quo

Men när vi fortsatte att växa och utvecklas nådde vi en avgörande insikt:

**Den traditionella SaaS-modellen att reservera kraftfulla funktioner för företagskunder är inte längre rimlig i dagens affärslandskap.**

År 2024, med kraften av AI och avancerade verktyg, kan små team verka i en skala och komplexitet som rivaliserar mycket större organisationer. En startup kan hantera känslig kunddata över flera länder. En liten marknadsföringsbyrå kan jonglera dussintals kundprojekt med varierande konfidentialitetskrav. Dessa företag behöver samma nivå av säkerhet och anpassning som *vilket* stort företag som helst.

Vi frågade oss själva: Varför ska storleken på ett företags arbetskraft eller budget avgöra deras förmåga att hålla sin data säker och sina processer effektiva?

### Företagsklassad för Alla

Denna insikt ledde oss till en kärnfilosofi som nu driver mycket av vår utveckling på Blue: Företagsklassade funktioner bör vara tillgängliga för företag av alla storlekar.

Vi tror att:

- **Säkerhet bör inte vara en lyx.** Varje företag, oavsett storlek, förtjänar verktygen för att skydda sin data och sina processer.
- **Flexibilitet driver innovation.** Genom att ge alla våra användare kraftfulla verktyg, möjliggör vi för dem att skapa arbetsflöden och system som driver deras branscher framåt.
- **Tillväxt bör inte kräva plattformsändringar.** När våra kunder växer, bör deras verktyg växa sömlöst med dem.

Med detta tankesätt beslutade vi att ta itu med utmaningen av anpassade behörigheter direkt, engagerade att göra det tillgängligt för alla våra användare, inte bara de på högre nivåer.

Detta beslut satte oss på en väg av noggrant design, iterativ utveckling och kontinuerlig användarfeedback som slutligen ledde till den anpassade behörighetsfunktion vi är stolta över att erbjuda idag.

I nästa avsnitt kommer vi att dyka ner i hur vi närmade oss design- och utvecklingsprocessen för att förverkliga denna komplexa funktion.

### Design och Utveckling

När vi beslutade att ta itu med anpassade behörigheter insåg vi snabbt att vi stod inför en kolossal uppgift.

Vid första anblicken kan "anpassade behörigheter" låta enkelt, men det är en bedrägligt komplex funktion som berör varje aspekt av vårt system.

Utmaningen var skrämmande: vi behövde implementera kaskadbehörigheter, tillåta redigeringar i farten, göra betydande ändringar i databasens schema och säkerställa sömlös funktionalitet över hela vårt ekosystem – webb, Mac, Windows, iOS och Android-appar, liksom vår API och webhooks.

Komplexiteten var tillräcklig för att få även de mest erfarna utvecklarna att tveka.

Vår strategi centrerade kring två nyckelprinciper:

1. Bryta ner funktionen i hanterbara versioner
2. Omfamna inkrementell leverans.

Konfronterade med komplexiteten av fullskala anpassade behörigheter ställde vi oss en avgörande fråga:

> Vad skulle vara den enklaste möjliga första versionen av denna funktion?

Denna strategi stämmer överens med den agila principen att leverera en Minimum Viable Product (MVP) och iterera baserat på feedback.

Vårt svar var uppfriskande enkelt:

1. Introducera en växel för att dölja projektaktivitetstabellen
2. Lägga till en annan växel för att dölja formulärtabellen

**Det var allt.**

Inga klockor och visselpipor, inga komplexa behörighetsmatriser—bara två enkla av/på-brytare.

Även om det kan verka otillräckligt vid första anblicken, erbjöd denna strategi flera betydande fördelar:

* **Snabb implementering**: Dessa enkla brytare kunde utvecklas och testas snabbt, vilket gjorde att vi kunde få en grundläggande version av anpassade behörigheter i användarnas händer snabbt.
* **Tydligt användarvärde**: Även med bara dessa två alternativ gav vi konkret värde. Vissa team kanske vill dölja aktivitetsflödet från kunder, medan andra kanske behöver begränsa åtkomsten till formulär för vissa användargrupper.
* **Grund för tillväxt**: Denna enkla start lade grunden för mer komplexa behörigheter. Det gjorde att vi kunde sätta upp den grundläggande infrastrukturen för anpassade behörigheter utan att fastna i komplexitet från början.
* **Användarfeedback**: Genom att släppa denna enkla version kunde vi samla verklig feedback om hur användare interagerade med anpassade behörigheter, vilket informerade vår framtida utveckling.
* **Teknisk inlärning**: Denna initiala implementering gav vårt utvecklingsteam praktisk erfarenhet av att modifiera behörigheter över vår plattform, vilket förberedde oss för mer komplexa iterationer.

Och du vet, det är faktiskt ganska ödmjukande att ha en stor vision för något, och sedan att leverera något som är en så liten procentandel av den visionen.

Efter att ha levererat dessa första två brytare beslutade vi att ta itu med något mer sofistikerat. Vi landade på två nya anpassade användarrollbehörigheter.

Den första var möjligheten att begränsa användare till att endast se poster som har tilldelats dem specifikt. Detta är mycket användbart om du har en kund i ett projekt och du bara vill att de ska se poster som är specifikt tilldelade dem istället för allt som du arbetar med för dem.

Den andra var ett alternativ för projektadministratörer att blockera användargrupper från att kunna bjuda in andra användare. Detta är bra om du har ett känsligt projekt som du vill säkerställa förblir på en "behöver se"-basis.

När vi hade levererat detta, fick vi mer självförtroende och för vår tredje version tog vi oss an kolumnnivåbehörigheter, vilket innebär att kunna bestämma vilka anpassade fält en specifik användargrupp kan se eller redigera.

Detta är extremt kraftfullt. Tänk dig att du har ett CRM-projekt, och du har data där som inte bara är relaterad till beloppen som kunden kommer att betala, utan också dina kostnader och vinstmarginaler. Du kanske inte vill att dina kostnadsfält och projektmarginalformelfält ska vara synliga för junior personal, och anpassad behörighet gör att du kan låsa dessa fält så att de inte visas.

Nästa steg var att skapa listbaserade behörigheter, där projektadministratörer kan bestämma om en användargrupp kan se, redigera och ta bort en specifik lista. Om de döljer en lista, blir alla poster inuti den listan också dolda, vilket är bra eftersom det innebär att du kan dölja vissa delar av din process från dina teammedlemmar eller kunder.

Detta är det slutgiltiga resultatet:

<video autoplay loop muted playsinline>
  <source src="/videos/custom-user-roles.mp4" type="video/mp4">
</video>

## Tekniska Överväganden

I hjärtat av Blues tekniska arkitektur ligger GraphQL, ett avgörande val som har påverkat vår förmåga att implementera komplexa funktioner som anpassade behörigheter. Men innan vi dyker ner i detaljerna, låt oss ta ett steg tillbaka och förstå vad GraphQL är och hur det skiljer sig från den mer traditionella REST API-ansatsen.
GraphQL vs REST API: En Tillgänglig Förklaring

Tänk dig att du är på en restaurang. Med ett REST API är det som att beställa från en fast meny. Du ber om en specifik rätt (ändpunkt), och du får allt som följer med den, oavsett om du vill ha allt eller inte. Om du vill anpassa din måltid kan du behöva göra flera beställningar (API-anrop) eller be om en specialberedd rätt (anpassad ändpunkt).

GraphQL, å andra sidan, är som att ha en konversation med en kock som kan förbereda vad som helst. Du berättar för kocken exakt vilka ingredienser du vill ha (datafält), och i vilka kvantiteter. Kocken förbereder sedan en rätt som är precis vad du bad om - inte mer, inte mindre. Detta är i grunden vad GraphQL gör - det tillåter klienten att be om exakt den data den behöver, och servern tillhandahåller just det.

### En Viktig Lunch

Ungefär sex veckor in i Blues initiala utveckling gick vår huvudingenjör och VD ut för att äta lunch.

Ämnet för diskussion? 

Huruvida vi skulle byta från REST API:er till GraphQL. Detta var inte ett beslut att ta lätt på - att anta GraphQL skulle innebära att vi skulle kassera sex veckors initialt arbete.

På vägen tillbaka till kontoret ställde VD:n en avgörande fråga till huvudingenjören: "Skulle vi ångra att vi inte gjorde detta om fem år?"

Svaret blev tydligt: GraphQL var vägen framåt.

Vi insåg potentialen i denna teknik tidigt, och såg hur den kunde stödja vår vision för en flexibel, kraftfull projektledningsplattform.

Vår förutseende i att anta GraphQL betalade sig när det kom till att implementera anpassade behörigheter. Med ett REST API skulle vi ha behövt en annan ändpunkt för varje möjlig konfiguration av anpassade behörigheter - en metod som snabbt skulle bli oöverskådlig och svår att underhålla.

GraphQL, å andra sidan, tillåter oss att hantera anpassade behörigheter dynamiskt. Så här fungerar det:

- **On-the-fly Behörighetskontroller**: När en klient gör en begäran kan vår GraphQL-server kontrollera användarens behörigheter direkt från vår databas.
- **Precise Data Retrieval**: Baserat på dessa behörigheter returnerar GraphQL endast den begärda data som passar inom användarens åtkomsträttigheter.
- **Flexibla Frågor**: När behörigheter ändras behöver vi inte skapa nya ändpunkter eller ändra befintliga. Den samma GraphQL-frågan kan anpassas till olika behörighetsinställningar.
- **Effektiv Datahämtning**: GraphQL tillåter klienter att begära exakt vad de behöver. Detta innebär att vi inte överhämtar data, vilket potentiellt skulle kunna exponera information som användaren inte borde få åtkomst till.

Denna flexibilitet är avgörande för en funktion så komplex som anpassade behörigheter. Det gör att vi kan erbjuda granulär kontroll *utan* att kompromissa med prestanda eller underhållbarhet.

## Utmaningar

Implementeringen av anpassade behörigheter i Blue medförde sina egna utmaningar, var och en som pressade oss att innovera och förfina vår strategi. Prestandaoptimering blev snabbt en kritisk fråga. När vi lade till fler granulära behörighetskontroller riskerade vi att sakta ner vårt system, särskilt för stora projekt med många användare och komplexa behörighetsinställningar. För att hantera detta implementerade vi en flerlagers cache-strategi, optimerade våra databasfrågor och utnyttjade GraphQL:s förmåga att begära endast nödvändig data. Denna metod gjorde att vi kunde bibehålla snabba svarstider även när projekten växte och behörighetskomplexiteten ökade.

Användargränssnittet för anpassade behörigheter presenterade en annan betydande hinder. Vi behövde göra gränssnittet intuitivt och hanterbart för administratörer, även när vi lade till fler alternativ och ökade systemets komplexitet.

Vår lösning involverade flera omgångar av användartester och iterativ design.

Vi introducerade en visuell behörighetsmatris som tillät administratörer att snabbt se och ändra behörigheter över olika roller och projektområden.

Att säkerställa plattformsövergripande konsistens presenterade en egen uppsättning utmaningar. Vi behövde implementera anpassade behörigheter enhetligt över våra webb-, skrivbords- och mobilapplikationer, var och en med sitt unika gränssnitt och användarupplevelseöverväganden. Detta var särskilt knepigt för våra mobilappar, som måste dynamiskt dölja och visa funktioner baserat på användarens behörigheter. Vi hanterade detta genom att centralisera vår behörighetslogik i API-lagret, vilket säkerställde att alla plattformar fick konsekvent behörighetsdata.

Sedan utvecklade vi ett flexibelt UI-ramverk som kunde anpassa sig till dessa behörighetsändringar i realtid, vilket gav en sömlös upplevelse oavsett vilken plattform som användes.

Användarutbildning och adoption presenterade det sista hindret i vår resa med anpassade behörigheter. Att introducera en så kraftfull funktion innebar att vi behövde hjälpa våra användare att förstå och effektivt utnyttja anpassade behörigheter.

Vi lanserade initialt anpassade behörigheter till en delmängd av vår användarbas, noggrant övervakande deras upplevelser och samlade insikter. Denna metod gjorde att vi kunde förfina funktionen och våra utbildningsmaterial baserat på verklig användning innan vi lanserade till hela vår användarbas.

Den fasade lanseringen visade sig vara ovärderlig, vilket hjälpte oss att identifiera och åtgärda mindre problem och användarförvirringspunkter som vi inte hade förutsett, vilket slutligen ledde till en mer polerad och användarvänlig funktion för alla våra användare.

Denna metod att lansera till en delmängd av användare, samt vår typiska 2-3 veckors "Beta"-period på vår offentliga Beta hjälper oss att sova gott om natten. :)

## Titta Framåt

Som med alla funktioner är inget någonsin *"klart"*.

Vår långsiktiga vision för funktionen för anpassade behörigheter sträcker sig över taggar, anpassade fältfilter, anpassningsbar projektnavigation och kommentarskontroller.

Låt oss dyka ner i varje aspekt.

### Taggbehörigheter

Vi tycker att det skulle vara fantastiskt att kunna skapa behörigheter baserat på huruvida en post har en eller flera taggar. Det mest uppenbara användningsfallet skulle vara att du skapar en anpassad användarroll som kallas "Kunder" och endast tillåter användare i den rollen att se poster som har taggen "Kunder".

Detta ger dig en översikt över huruvida en post kan eller inte kan ses av dina kunder.

Detta kan bli ännu kraftfullare med OCH/ELLER-kombinatorer, där du kan specificera mer komplexa regler. Till exempel kan du ställa in en regel som tillåter åtkomst till poster som är taggade både "Kunder" OCH "Offentlig", eller poster som är taggade antingen "Intern" ELLER "Konfidentiell". Denna nivå av flexibilitet skulle möjliggöra otroligt nyanserade behörighetsinställningar, som tillgodoser även de mest komplexa organisationsstrukturer och arbetsflöden.

De potentiella tillämpningarna är stora. Projektledare skulle enkelt kunna separera känslig information, säljteam skulle kunna ha automatisk åtkomst till relevant kunddata, och externa samarbetspartners skulle kunna integreras sömlöst i specifika delar av ett projekt utan att riskera exponering för känslig intern information.

### Anpassade Fältfilter

Vår vision för Anpassade Fältfilter representerar ett betydande framsteg inom granulär åtkomstkontroll. Denna funktion kommer att ge projektadministratörer möjlighet att definiera vilka poster specifika användargrupper kan se baserat på värdena av anpassade fält. Det handlar om att skapa dynamiska, datadrivna gränser för informationsåtkomst.

Tänk dig att kunna ställa in behörigheter som detta:

- Visa endast poster där "Projektstatus" dropdown är inställd på "Offentlig"
- Begränsa synlighet till objekt där "Avdelning" multi-select fält inkluderar "Marknadsföring"
- Tillåta åtkomst till uppgifter där "Prioritet" checkbox är ikryssad
- Visa projekt där "Budget" nummerfält är över en viss gräns

### Anpassningsbar Projektnavigation

Detta är helt enkelt en förlängning av de brytare som vi redan har. Istället för att bara ha brytare för "aktivitet" och "formulär", vill vi utöka det till varje del av projektets navigation. På så sätt kan projektadministratörer skapa fokuserade gränssnitt och ta bort verktyg som de inte behöver.

### Kommentarskontroller

I framtiden vill vi vara kreativa i hur vi låter våra kunder bestämma vem som kan och inte kan se kommentarer. Vi kan tillåta flera flikade kommentarsområden under en post, och varje kan vara synlig eller inte synlig för olika användargrupper.

Dessutom kan vi också tillåta en funktion där endast kommentarer där en användare *specifikt* nämns är synliga, och inget annat. Detta skulle möjliggöra för team som har kunder i projekt att säkerställa att endast kommentarer som de vill att kunder ska se är synliga.

## Slutsats

Så där har vi det, så här närmade vi oss att bygga en av de mest intressanta och kraftfulla funktionerna! [Som du kan se på vårt verktyg för jämförelse av projektledning](/compare) har mycket få projektledningssystem en så kraftfull behörighetsmatrisuppsättning, och de som har det reserverar det för sina dyraste företagsplaner, vilket gör det otillgängligt för ett typiskt litet eller medelstort företag.

Med Blue har du *alla* funktioner tillgängliga med vår plan — vi tror inte att företagsklassade funktioner bör reserveras för företagskunder!
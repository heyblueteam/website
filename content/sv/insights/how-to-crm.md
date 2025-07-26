---
title: Hur man ställer in Blue som ett CRM
description: Lär dig hur du ställer in Blue för att spåra dina kunder och affärer på ett enkelt sätt.
category: "Best Practices"
date: 2024-08-11
---



## Introduktion

En av de viktigaste fördelarna med att använda Blue är att det inte används för ett *specifikt* användningsområde, utan används *över* användningsområden. Detta innebär att du inte behöver betala för flera verktyg, och du har också en plats där du enkelt kan växla mellan dina olika projekt och processer som rekrytering, försäljning, marknadsföring och mer.

Genom att hjälpa tusentals kunder att komma igång med Blue genom åren har vi märkt att den svåra delen inte är att *ställa in* Blue själv, utan att tänka igenom processerna och få ut det mesta av vår plattform.

De centrala delarna handlar om att tänka på arbetsflödet steg för steg för varje affärsprocess som du vill spåra, samt specifikationerna för den data som du vill fånga och hur detta översätts till de anpassade fälten som du ställer in.

Idag kommer vi att gå igenom hur du skapar [ett lättanvänt, men kraftfullt försäljnings-CRM-system](/solutions/use-case/sales-crm) med en kunddatabas som är kopplad till en pipeline av möjligheter. All denna data kommer att flöda in i en instrumentpanel där du kan se realtidsdata om dina totala försäljningar, prognostiserade försäljningar och mer.

## Kunddatabas

Det första steget är att ställa in ett nytt projekt för att lagra din kunddata. Denna data kommer sedan att korsrefereras i ett annat projekt där du spårar specifika försäljningsmöjligheter.

Anledningen till att vi separerar din kundinformation från möjligheterna är att de inte kartläggs ett till ett.

En kund kan ha flera möjligheter eller projekt.

Till exempel, om du är en marknadsförings- och designbyrå, kan du först samarbeta med en kund för deras varumärke, och sedan göra ett separat projekt för deras webbplats, och sedan ett annat för deras hantering av sociala medier.

Alla dessa skulle vara separata försäljningsmöjligheter som kräver sin egen spårning och förslag, men de är alla kopplade till den ena kunden.

Fördelen med att separera din kunddatabas till ett eget projekt är att om du uppdaterar några detaljer i din kunddatabas, kommer alla dina möjligheter automatiskt att ha den nya datan, vilket innebär att du nu har en sanningskälla i ditt företag! Du behöver inte gå tillbaka och redigera allt manuellt!

Så, det första du behöver bestämma är om du kommer att vara företagscentrerad eller personcentrerad.

Detta beslut beror verkligen på vad du säljer och till vem du säljer. Om du säljer främst till företag, så vill du förmodligen att postnamnet ska vara företagsnamnet. Men om du säljer mest till individer (dvs. du är en personlig hälsocoach eller en personlig varumärkesrådgivare), så skulle du troligen ta en personcentrerad strategi.

Så postnamnsfältet kommer antingen att vara företagsnamnet eller personnamnet, beroende på ditt val. Anledningen till detta är att det innebär att du enkelt kan identifiera en kund vid en blick, bara genom att titta på din tavla eller databas.

Nästa steg är att överväga vilken information du vill fånga som en del av din kunddatabas. Dessa kommer att bli dina anpassade fält.

De vanliga misstänkta här är:

- E-post
- Telefonnummer
- Webbplats
- Adress
- Källa (dvs. var kom denna kund först ifrån?)
- Kategori

I Blue kan du också ta bort eventuella standardfält som du inte behöver. För denna kunddatabas rekommenderar vi vanligtvis att du tar bort förfallodatum, tilldelad, beroenden och checklistor. Du kanske vill behålla vårt standardbeskrivningsfält tillgängligt ifall du har allmänna anteckningar om den kunden som inte är specifika för någon försäljningsmöjlighet.

Vi rekommenderar att du behåller fältet "Referens av", eftersom detta kommer att vara användbart senare. När vi ställer in vår möjlighetsdatabas kommer vi att kunna se varje försäljningspost som är kopplad till denna specifika kund här.

När det gäller listor ser vi vanligtvis att våra kunder bara håller det enkelt och har en lista som heter "Kunder" och lämnar det så. Det är bättre att använda taggar eller anpassade fält för kategorisering.

Det som är bra här är att när du har detta inställt kan du enkelt importera din data från andra system eller Excel-ark till Blue via vår CSV-importfunktion, och du kan också skapa ett formulär för nya potentiella kunder att skicka in sina uppgifter så att du kan **automatiskt** fånga dem i din databas.

## Möjlighetsdatabas

Nu när vi har vår kunddatabas behöver vi skapa ett annat projekt för att fånga våra faktiska försäljningsmöjligheter. Du kan kalla detta projekt "Försäljnings-CRM" eller "Möjligheter".

### Listor som Processsteg

För att ställa in din försäljningsprocess behöver du tänka på vilka vanliga steg en möjlighet går igenom från det ögonblick du får en förfrågan från en kund hela vägen till att få ett undertecknat kontrakt.

Varje lista i ditt projekt kommer att vara ett steg i din process.

Oavsett din specifika process kommer det att finnas några gemensamma listor som ALLA försäljnings-CRM bör ha:

- Oqualificerad — Alla inkommande förfrågningar, där du ännu inte har kvalificerat en kund.
- Stängd Vunnen — Alla möjligheter som du har vunnit och omvandlat till försäljningar!
- Stängd Förlorad — Alla möjligheter där du har gett en offert till en kund, och de inte accepterade.
- N/A — Här placerar du alla möjligheter som du inte vann, men som inte heller var "förlorade". Det kan vara de som du avböjde, de där kunden, av vilken anledning som helst, spökade dig, och så vidare.

När det gäller att tänka igenom din försäljnings-CRM-affärsprocess bör du överväga nivån av detaljrikedom som du vill ha. Vi rekommenderar inte att ha 20 eller 30 kolumner, detta blir vanligtvis förvirrande och hindrar dig från att kunna se den större bilden.

Det är dock också viktigt att inte göra varje process för bred, annars kommer affärer att "fastna" i ett specifikt steg i veckor eller månader, även när de faktiskt går framåt. Här är en typisk rekommenderad strategi:

- **Oqualificerad**: Alla inkommande förfrågningar, där du ännu inte har kvalificerat en kund.
- **Kvalificering**: Här tar du möjligheten och börjar processen att förstå om detta är en bra match för ditt företag.
- **Skriva förslag**: Här börjar du omvandla möjligheten till ett erbjudande för ditt företag. Detta är ett dokument som du skulle skicka till kunden.
- **Förslag skickat**: Här har du skickat förslaget till kunden och väntar på svar.
- **Förhandlingar**: Här är du i processen att slutföra affären.
- **Kontrakt ute för underskrift**: Här väntar du bara på att kunden ska skriva under kontraktet.
- **Stängd Vunnen**: Här har du vunnit affären och arbetar nu med projektet.
- **Stängd Förlorad**: Här har du gett kunden en offert, men de har inte accepterat villkoren.
- **N/A**: Här placerar du alla möjligheter som du inte vann, men som inte heller var "förlorade". Det kan vara de som du avböjde, de där kunden, av vilken anledning som helst, spökade dig, och så vidare.

### Taggar som Tjänstekategorier
Låt oss nu prata om taggar.

Vi rekommenderar att du använder taggar för de olika typer av tjänster som du erbjuder. Så, tillbaka till vårt exempel med marknadsförings- och designbyrån, kan du ha taggar för "varumärke", "webbplats", "SEO", "Facebook-hantering" och så vidare.

Fördelarna här är att du enkelt kan filtrera ner efter tjänst med ett klick, vilket kan ge dig en kort översikt över vilka tjänster som är mer populära, och detta kan också informera framtida rekrytering, eftersom olika tjänster vanligtvis kräver olika teammedlemmar.

### Anpassade fält för Försäljnings-CRM

Nästa steg är att överväga vilka anpassade fält vi vill ha.

Typiska fält som vi ser användas är:

- **Belopp**: Detta är ett valutafält för beloppet av projektet
- **Kostnad**: Din förväntade kostnad för att uppfylla försäljningen, också ett valutafält
- **Vinst**: Ett formelfält för att beräkna vinsten baserat på belopp- och kostnadsfälten.
- **Förslag URL**: Detta kan inkludera en länk till ett online Google-dokument eller Word-dokument av ditt förslag, så att du enkelt kan klicka och granska det.
- **Mottagna filer**: Detta kan vara ett filanpassat fält där du kan släppa in filer som mottagits från kunden, såsom forskningsmaterial, NDA:er och så vidare.
- **Kontrakt**: Ett annat filanpassat fält där du kan lägga till signerade kontrakt för säker förvaring.
- **Konfidensnivå**: Ett stjärnanpassat fält med 5 stjärnor, som indikerar hur säkra du är på att vinna denna specifika möjlighet. Detta kan användas senare i instrumentpanelen för prognostisering!
- **Förväntat stängningsdatum**: Ett datumfält för att uppskatta när affären sannolikt kommer att stängas.
- **Kund**: Ett referensfält som länkar till den primära kontaktpersonen i kunddatabasen.
- **Kundnamn**: Ett uppslagsfält som hämtar kundnamnet från den specifika länkade posten i kunddatabasen.
- **Kundens e-post**: Ett uppslagsfält som hämtar kundens e-post från den specifika länkade posten i kunddatabasen.
- **Affärskälla**: Ett rullgardinsfält för att spåra var möjligheten kom ifrån (t.ex. referens, webbplats, kall samtal, mässa).
- **Orsak till förlust**: Ett rullgardinsfält (för stängda förlorade affärer) för att kategorisera varför möjligheten förlorades.
- **Kundstorlek**: Ett rullgardinsfält för att kategorisera kunder efter storlek (t.ex. liten, medel, stor företag).

Återigen, det är verkligen **upp till dig** att bestämma exakt vilka fält du vill ha. Ett varningens ord: det är lätt att lägga till massor av fält i ditt försäljnings-CRM av data du vill fånga. Men du måste vara realistisk när det gäller disciplin och tidsåtagande. Det finns ingen poäng i att ha 30 fält i ditt försäljnings-CRM om 90% av posterna inte kommer att ha någon data i dem.

Det fantastiska med anpassade fält är att de integreras väl i [Anpassade behörigheter](/platform/features/user-permissions). Detta innebär att du kan bestämma exakt vilka fält teammedlemmar i ditt team kan se eller redigera. Till exempel kan du vilja dölja kostnads- och vinstinformation från junior personal.

### Automatiseringar

[Automatiseringar för försäljnings-CRM](/platform/features/automations) är en kraftfull funktion i Blue som kan effektivisera din försäljningsprocess, säkerställa konsekvens och spara tid på repetitiva uppgifter. Genom att ställa in intelligenta automatiseringar kan du förbättra effektiviteten i ditt försäljnings-CRM och låta ditt team fokusera på det som är viktigast - att stänga affärer. Här är några viktiga automatiseringar att överväga för ditt försäljnings-CRM:

- **Ny lead-tilldelning**: Automatisk tilldelning av nya leads till försäljningsrepresentanter baserat på fördefinierade kriterier som plats, affärsstorlek eller bransch. Detta säkerställer snabb uppföljning och balanserad arbetsbelastning.
- **Uppföljningspåminnelser**: Ställ in automatiska påminnelser för försäljningsrepresentanter att följa upp med prospekter efter en viss period av inaktivitet. Detta hjälper till att förhindra att leads faller mellan stolarna.
- **Stegprogressionsnotifikationer**: Meddela relevanta teammedlemmar när en affär går till ett nytt steg i pipelinen. Detta håller alla informerade om framsteg och möjliggör snabba insatser om det behövs.
- **Affärsåldersvarningar**: Skapa varningar för affärer som har varit i ett visst steg längre än förväntat. Detta hjälper till att identifiera stillastående affärer som kan behöva extra uppmärksamhet.

## Koppla kunder och affärer

En av de mest kraftfulla funktionerna i Blue för att skapa ett effektivt CRM-system är möjligheten att koppla din kunddatabas med dina försäljningsmöjligheter. Denna koppling gör att du kan upprätthålla en enda sanningskälla för kundinformation samtidigt som du spårar flera affärer kopplade till varje kund. Låt oss utforska hur man ställer in detta med hjälp av referens- och uppslagsanpassade fält.

### Ställa in referensfältet

1. I ditt projekt för Möjligheter (eller Försäljnings-CRM), skapa ett nytt anpassat fält.
2. Välj fälttypen "Referens".
3. Välj ditt projekt för Kunddatabas som källa för referensen.
4. Konfigurera fältet för att tillåta enstaka val (eftersom varje möjlighet vanligtvis är kopplad till en kund).
5. Namnge detta fält något som "Kund" eller "Kopplad Företag".

Nu, när du skapar eller redigerar en möjlighet, kommer du att kunna välja den kopplade kunden från en rullgardinsmeny som fylls med poster från din kunddatabas.

### Förbättra med uppslagsfält

När du har etablerat referenskopplingen kan du använda uppslagsfält för att ta in relevant kundinformation direkt i din möjlighetsvy. Så här gör du:

1. I ditt projekt för Möjligheter, skapa ett nytt anpassat fält.
2. Välj fälttypen "Uppslag".
3. Välj referensfältet du just skapade ("Kund" eller "Kopplad Företag") som källa.
4. Välj vilken kundinformation du vill visa. Du kan överväga fält som: E-post, Telefonnummer, Kundkategori, Kontohanterare.

Upprepa denna process för varje bit av kundinformation som du vill visa i din möjlighetsvy.

Fördelarna med detta är:

- **Enkel sanningskälla**: Uppdatera kundinformation en gång i kunddatabasen, och det återspeglas automatiskt i alla kopplade möjligheter.
- **Effektivitet**: Snabbt få tillgång till relevant kundinformation medan du arbetar med möjligheter utan att växla mellan projekt.
- **Dataintegritet**: Minska fel från manuell datainmatning genom att automatiskt hämta kundinformation.
- **Helhetssyn**: Se enkelt alla möjligheter kopplade till en kund genom att använda fältet "Refererad av" i din kunddatabas.

### Avancerat tips: Uppslag av ett uppslag

Blue erbjuder en avancerad funktion som kallas "Uppslag av ett uppslag" som kan vara otroligt användbar för komplexa CRM-inställningar. Denna funktion gör att du kan skapa kopplingar över flera projekt, vilket gör att du kan få tillgång till information från både din kunddatabas och ditt projekt för möjligheter i ett tredje projekt.

Till exempel, låt oss säga att du har en "Projekt"-arbetsyta där du hanterar det faktiska arbetet för dina kunder. Du vill att denna arbetsyta ska ha tillgång till både kunduppgifter och möjlighetsinformation. Så här kan du ställa in detta:

Först, skapa ett referensfält i din arbetsyta för Projekt som länkar till projektet för Möjligheter. Detta etablerar den initiala kopplingen. Nästa steg är att skapa uppslagsfält baserat på denna referens för att hämta specifika detaljer från möjligheterna, såsom affärsvärde eller förväntat stängningsdatum.

Den verkliga kraften kommer i nästa steg: du kan skapa ytterligare uppslagsfält som når genom möjligheternas referens till kunddatabasen. Detta gör att du kan hämta kundinformation som kontaktuppgifter eller kontostatus direkt in i din arbetsyta för Projekt.

Denna kedja av kopplingar ger dig en omfattande vy i din arbetsyta för Projekt, som kombinerar data från både dina möjligheter och kunddatabas. Det är ett kraftfullt sätt att säkerställa att dina projektteam har all relevant information till hands utan att behöva växla mellan olika projekt.

### Bästa praxis för länkade CRM-system

Underhåll din kunddatabas som den enda sanningskällan för all kundinformation. När du behöver uppdatera kunduppgifter, gör alltid det i kunddatabasen först. Detta säkerställer att informationen förblir konsekvent över alla länkade projekt.

När du skapar referens- och uppslagsfält, använd tydliga och meningsfulla namn. Detta hjälper till att upprätthålla tydlighet, särskilt när ditt system växer mer komplext.

Granska regelbundet din inställning för att säkerställa att du hämtar den mest relevanta informationen. När dina affärsbehov utvecklas kan det hända att du behöver lägga till nya uppslagsfält eller ta bort sådana som inte längre är användbara. Periodiska granskningar hjälper till att hålla ditt system strömlinjeformat och effektivt.

Överväg att utnyttja Blues automatiseringsfunktioner för att hålla dina data synkroniserade och uppdaterade över projekt. Till exempel kan du ställa in en automatisering för att meddela relevanta teammedlemmar när viktig kundinformation uppdateras i kunddatabasen.

Genom att effektivt implementera dessa strategier och utnyttja referens- och uppslagsfält kan du skapa ett kraftfullt, sammanlänkat CRM-system i Blue. Detta system kommer att ge dig en omfattande 360-graders vy av dina kundrelationer och försäljningspipeline, vilket möjliggör mer informerat beslutsfattande och smidigare verksamhet i hela din organisation.

## Instrumentpaneler

Instrumentpaneler är en avgörande komponent i varje effektivt CRM-system, som ger ögonblickliga insikter i din försäljningsprestanda och kundrelationer. Blues instrumentpanelfunktion är särskilt kraftfull eftersom den gör att du kan kombinera realtidsdata från flera projekt samtidigt, vilket ger dig en omfattande vy av dina försäljningsoperationer.

När du ställer in din CRM-instrumentpanel i Blue, överväg att inkludera flera viktiga mätvärden. Pipeline genererad per månad visar det totala värdet av nya möjligheter som lagts till din pipeline, vilket hjälper dig att spåra ditt teams förmåga att generera nya affärer. Försäljning per månad visar dina faktiska stängda affärer, vilket gör att du kan övervaka ditt teams prestation i att omvandla möjligheter till försäljningar.

Att införa konceptet pipeline-rabatter kan leda till mer exakta prognoser. Till exempel kan du räkna 90% av värdet av affärer i "Kontrakt ute för underskrift"-steget, men bara 50% av affärer i "Förslag skickat"-steget. Denna viktade strategi ger en mer realistisk försäljningsprognos.

Att spåra nya möjligheter per månad hjälper dig att övervaka antalet nya potentiella affärer som kommer in i din pipeline, vilket är en bra indikator på ditt försäljningsteams prospekteringsinsatser. Att bryta ner försäljning efter typ kan hjälpa dig att identifiera dina mest framgångsrika erbjudanden. Om du ställer in ett fakturatrackingprojekt kopplat till dina möjligheter kan du också spåra faktisk intäkt på din instrumentpanel, vilket ger en komplett bild från möjlighet till kontanter.

Blue erbjuder flera kraftfulla funktioner för att hjälpa dig skapa en informativ och interaktiv CRM-instrumentpanel. Plattformen tillhandahåller tre huvudtyper av diagram: statistikkort, cirkeldiagram och stapeldiagram. Statistikkort är idealiska för att visa nyckelmätvärden som totalt pipelinevärde eller antal aktiva möjligheter. Cirkeldiagram är perfekta för att visa sammansättningen av din försäljning efter typ eller fördelningen av affärer över olika steg. Stapeldiagram är utmärkta för att jämföra mätvärden över tid, såsom månatlig försäljning eller nya möjligheter.

Blues sofistikerade filtreringsmöjligheter gör att du kan segmentera dina data efter projekt, lista, tagg och tidsram. Detta är särskilt användbart för att gräva ner i specifika aspekter av din försäljningsdata eller jämföra prestationer över olika team eller produkter. Plattformen smart konsoliderar listor och taggar med samma namn över projekt, vilket möjliggör sömlös analys över projekt. Detta är ovärderligt för en CRM-inställning där du kan ha separata projekt för kunder, möjligheter och fakturor.

Anpassning är en nyckelstyrka i Blues instrumentpanelfunktion. Drag-och-släpp-funktionen och visningsflexibiliteten gör att du kan skapa en instrumentpanel som perfekt passar dina behov. Du kan enkelt omorganisera diagram och välja den mest lämpliga visualiseringen för varje mätvärde.
Även om instrumentpaneler för närvarande endast är för internt bruk kan du enkelt dela dem med teammedlemmar och ge antingen visnings- eller redigeringsbehörigheter. Detta säkerställer att alla i ditt försäljningsteam har tillgång till de insikter de behöver.

Genom att utnyttja dessa funktioner och inkludera de nyckelmätvärden vi har diskuterat kan du skapa en omfattande CRM-instrumentpanel i Blue som ger realtidsinsikter i din försäljningsprestanda, pipelinehälsa och övergripande affärstillväxt. Denna instrumentpanel kommer att bli ett ovärderligt verktyg för att fatta datadrivna beslut och hålla hela ditt team i linje med dina försäljningsmål och framsteg.

## Slutsats

Att ställa in ett omfattande försäljnings-CRM i Blue är ett kraftfullt sätt att effektivisera din försäljningsprocess och få värdefulla insikter i dina kundrelationer och affärsprestanda. Genom att följa stegen i denna guide har du skapat ett robust system som integrerar kundinformation, försäljningsmöjligheter och prestationsmätvärden i en enda, sammanhängande plattform.

Vi började med att skapa en kunddatabas, vilket etablerade en enda sanningskälla för all din kundinformation. Denna grund gör att du kan upprätthålla korrekta och aktuella register för alla dina kunder och prospekt. Vi byggde sedan vidare på detta med en möjlighetsdatabas, vilket möjliggör att du effektivt kan spåra och hantera din försäljningspipeline.

En av de viktigaste styrkorna med att använda Blue för ditt CRM är möjligheten att länka dessa databaser med hjälp av referens- och uppslagsfält. Denna integration skapar ett dynamiskt system där uppdateringar av kundinformation omedelbart återspeglas i alla relaterade möjligheter, vilket säkerställer datakonsistens och sparar tid på manuella uppdateringar.
Vi utforskade hur man utnyttjar Blues kraftfulla automatiseringsfunktioner för att effektivisera ditt arbetsflöde, från att tilldela nya leads till att skicka uppföljningspåminnelser. Dessa automatiseringar hjälper till att säkerställa att inga möjligheter faller mellan stolarna och att ditt team kan fokusera på aktiviteter med högre värde snarare än administrativa uppgifter.

Slutligen fördjupade vi oss i att skapa instrumentpaneler som ger ögonblickliga insikter i din försäljningsprestanda. Genom att kombinera data från dina kund- och möjlighetsdatabaser erbjuder dessa instrumentpaneler en omfattande vy av din försäljningspipeline, stängda affärer och övergripande affärshälsa.

Kom ihåg, nyckeln till att få ut det mesta av ditt CRM är konsekvent användning och regelbunden förfining. Uppmuntra ditt team att fullt ut adoptera systemet, granska regelbundet dina processer och automatiseringar, och fortsätta utforska nya sätt att utnyttja Blues funktioner för att stödja dina försäljningsinsatser.

Med denna försäljnings-CRM-inställning i Blue är du väl rustad för att vårda kundrelationer, stänga fler affärer och driva ditt företag framåt.
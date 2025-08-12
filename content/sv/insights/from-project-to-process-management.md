---
title: Från projekt till processer  
description: Hur vi omvandlar Blue från ett projektledningsverktyg till en omfattande plattform för processhantering, och varför denna förändring är viktig för framtiden för teamwork.
category: "CEO Blog"
date: 2025-08-12
---
Hej, det här är Manny, grundaren och VD:n för Blue.

Jag vill dela med mig av något som har funnits i mina tankar på sistone.

Det är en grundläggande förändring i hur vi tänker på Blue och, mer generellt, hur team hanterar sitt arbete. Det är en insikt som kommit från otaliga samtal med våra kunder och att observera hur de faktiskt använder Blue i verkligheten.

Här är saken: Jag byggde Blue som ett projektledningsverktyg. Med en bakgrund inom byråvärlden var det vad som kändes logiskt för mig. Vi startade ett kundprojekt, lade till alla uppgifter, bockade av dem och arkiverade så småningom det avslutade projektet.

Rent, enkelt, klart.

När Blue växte och jag pratade med fler kunder, märkte jag något som ärligt talat överraskade mig:

**Över 90 % av våra kunder använde inte Blue för projekt alls**.

De använde det för att hantera pågående processer—den typen av arbete som aldrig riktigt "slutar".

Jag har sett några vilda användningsfall genom åren som jag aldrig skulle ha kunnat föreställa mig.

Och nyligen klickade det: varför försöker vi hacka ett "projektledningssystem" för processhantering?

Ja, det fungerar, men varför inte bara lösa det större problemet direkt?

## Projekt vs. Processer: Den Grundläggande Skillnaden

Låt mig bryta ner detta enkelt:

**Projekt** har ett bestämt startdatum och (förhoppningsvis) ett eventualt slutdatum. De är tillfälliga insatser med specifika leveranser. Tänk på att lansera en ny webbplats, organisera ett evenemang eller utveckla en ny produktfunktion.

**Processer** är pågående, ständiga delar av ditt företag. De är de upprepbara arbetsflöden som håller ditt företag igång: försäljningsprocesser, rekryteringspipeline, kundintroduktion, hantering av supportärenden, arbetsflöden för innehållsskapande. Dessa "slutar" aldrig riktigt—de fortsätter bara att cykla genom.

Ironin?

Även inom Blue fann vi oss använda vårt eget verktyg mer för processer än för projekt.

Vi har processer för mjukvaruutveckling, rekrytering, kundförfrågningar om anpassade domäner och dussintals fler.

Det var rakt framför oss hela tiden.

## Vad Detta Betyder för Blue

Denna insikt har fundamentalt förändrat hur vi bygger Blue.

### 1. **Från "Att-göra" till "Poster"**
Vi har döpt om våra kärnbyggstenar från "att-göra" till "poster"—och gjort detta anpassningsbart. Ditt säljteam kan kalla dem "möjligheter", support kan använda "ärenden", HR kanske föredrar "kandidater". Det är en liten förändring med stora konsekvenser för hur team tänker på sitt arbete. Vi gör samma sak med "projekt" och förvandlar dem till "arbetsytor".

### 2. **Tidsintelligens**
Vi har lagt till anpassade fält som spårar tid på sätt som är viktiga för processer:
- **Tid i lista**: Se exakt hur länge en post har tillbringat i varje steg av din process
- **Varaktighet mellan händelser**: Mät cykeltid mellan två steg (som från "Pågående" till "Färdig")

Detta är inte bara trevliga mätvärden—de är avgörande för att identifiera flaskhalsar och förbättra processer över tid.

### 3. **Smart automatisering**

Vi arbetar för närvarande med schemalagda automatiseringar som gör att våra automatiseringar kan aktiveras upprepade gånger enligt ett regelbundet schema och endast agera på poster som uppfyller specifika kriterier.

De andra två stora utmaningarna är:

1. Villkorlig automatisering (där du kan ha flera kriterier som måste vara sanna innan något händer)
2. Försöka lösa problemet med oändliga loopar (där en automatisering för närvarande inte kan utlösa en annan automatisering eftersom det skulle kunna orsaka en oändlig loop). I verkligheten skulle vi gärna se att detta är möjligt och vi upptäcker bara oändliga loopar och ser till att du inte kan skapa dem.

### 4. **Multi-Homing**

Detta är stort: en post kommer att kunna finnas i flera projekt samtidigt som den automatiskt synkroniserar data från anpassade fält. Tänk dig en kundförfrågan som behöver spåras av både support och teknik—inga fler duplicerade arbeten eller förlorad kontext.

## Hitta vår marknadsposition

Denna förändring placerar oss i en intressant position. Vi konkurrerar inte längre bara med Asana, Monday eller Trello om projektledning. Men vi försöker inte heller vara ett massivt ERP-system som kostar miljoner att implementera och tar år att rulla ut.

Jag tror att vi skapar något nytt—en sweet spot mellan lätta projektverktyg och tunga företagsystem. Det finns en enorm lucka på marknaden för team som behöver mer än grundläggande uppgiftshantering men som inte kan rättfärdiga (eller inte vill ha) komplexiteten hos traditionell affärsprocesshanteringsprogramvara.

## Utmaningen med enkelhet

Nu ska jag vara ärlig—detta är något vi brottas med varje dag. Hur lägger vi till sofistikerade processhanteringsfunktioner samtidigt som vi håller oss trogna vår "Håll det enkelt"-filosofi?

Det finns inget enkelt svar. Det handlar om att göra kontinuerliga avvägningar. En strategi vi använder är att acceptera lite mer komplexitet för administratörer som konfigurerar processerna, samtidigt som vi håller saker dödligt enkla för slutanvändarna som arbetar inom dessa processer dagligen. De som sätter upp arbetsflödet kanske behöver tänka igenom automatiseringar och fältkonfigurationer, men personen som hanterar en kundförfrågan bör bara se ett klart, enkelt gränssnitt som berättar vad de ska göra härnäst.

## Framtiden

Vi tillkännager denna förändring offentligt med vår nya webbplats och uppdaterad dokumentation. Många av dessa processfokuserade funktioner är redan aktiva, med många fler på gång. Responsen från kunder som vi har delat denna vision med har varit otroligt positiv—de förstår det omedelbart eftersom de redan använder Blue på detta sätt.

Framtiden för arbete handlar inte om att hantera isolerade projekt—det handlar om att orkestrera de sammanlänkade processerna som får företag att fungera. Det handlar om att ge team flexibiliteten att definiera sina egna arbetsflöden samtidigt som man upprätthåller den struktur som behövs för konsekvens och förbättring.

Vi bygger Blue för att vara den enklaste, mest intuitiva plattformen för att hantera alla typer av affärsprocesser. Tänk schemalagda automatiseringar, villkorlig logik, avancerad flaskhalsspårning, och kanske till och med helt nya vyer bortom våra nuvarande sex som låter dig se processer på en strategisk nivå.

Och när jag tänker på den verkliga konkurrensen till Blue handlar det inte riktigt om programvarusystem utan det handlar verkligen om att möjliggöra organisationer att utveckla det tankesättet kring processhantering, och att skifta bort från kaotiska e-posttrådar, klisterlappar och löst kopplade kalkylblad.

Denna evolution från projekt till processhantering handlar inte bara om Blue—det handlar om att erkänna hur moderna team faktiskt arbetar och bygga verktyg som matchar den verkligheten. Vi är entusiastiska över denna riktning och kan inte vänta på att se hur team använder dessa funktioner för att strömlinjeforma sina viktigaste processer.

Om du redan använder Blue för processhantering, skulle vi gärna höra din historia. Och om du fortfarande tänker i projekt när du verkligen behöver processhantering, kanske det är dags att göra skiftet.

Håll dig produktiv,  

Manny

**PS**: Vi lanserar regelbundet nya processfokuserade funktioner. Håll ett öga på vår [ändringslogg](/platform/changelog) för de senaste uppdateringarna, och tveka inte att höra av dig om du har idéer om hur vi kan stödja ditt teams processer bättre.
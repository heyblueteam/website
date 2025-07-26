---
title:  Hur vi använder Blue för att bygga Blue. 
description: Lär dig hur vi använder vår egen projektledningsplattform för att bygga vår projektledningsplattform!
category: "CEO Blog"
date: 2024-08-07
---



Du är på väg att få en insider-tur om hur Blue bygger Blue.

På Blue äter vi vår egen hundmat.

Detta betyder att vi använder Blue för att *bygga* Blue.

Denna konstiga term, som ofta kallas "dogfooding", tillskrivs ofta Paul Maritz, en chef på Microsoft under 1980-talet. Han ska ha skickat ett e-postmeddelande med ämnesraden *"Äta vår egen hundmat"* för att uppmuntra Microsoft-anställda att använda företagets produkter.

Idén att använda sina egna verktyg för att bygga sina verktyg leder till en positiv feedbackcykel.

Idén att använda sina egna verktyg för att bygga sina verktyg leder till en positiv feedbackcykel, vilket skapar många fördelar:

- **Det hjälper oss att snabbt identifiera verkliga användbarhetsproblem.** Eftersom vi använder Blue dagligen stöter vi på samma utmaningar som våra användare kan möta, vilket gör att vi kan ta itu med dem proaktivt.
- **Det påskyndar upptäckten av buggar.** Intern användning avslöjar ofta buggar innan de når våra kunder, vilket förbättrar den övergripande produktkvaliteten.
- **Det ökar vår empati för slutanvändare.** Vårt team får förstahandsupplevelse av Blues styrkor och svagheter, vilket hjälper oss att fatta mer användarcentrerade beslut.
- **Det driver en kultur av kvalitet inom vår organisation.** När alla använder produkten finns det en gemensam insats i dess excellens.
- **Det främjar innovation.** Regelbunden användning väcker ofta idéer för nya funktioner eller förbättringar, vilket håller Blue i framkant.


[Vi har tidigare pratat om varför vi inte har något dedikerat testteam](/insights/open-beta) och detta är ännu en anledning.

Om det finns buggar i vårt system hittar vi dem nästan alltid i vår ständiga dagliga användning av plattformen. Och detta skapar också en tvingande funktion för att åtgärda dem, eftersom vi uppenbarligen kommer att tycka att de är mycket irriterande då vi förmodligen är en av de största användarna av Blue!

Denna metod visar vårt engagemang för produkten. Genom att förlita oss på Blue själva visar vi våra kunder att vi verkligen tror på det vi bygger. Det är inte bara en produkt vi säljer – det är ett verktyg vi är beroende av varje dag.

## Huvudprocess

Vi har ett projekt i Blue, passande nog kallat "Produkt".

**Allt** relaterat till vår produktutveckling spåras här. Kundfeedback, buggar, funktionsidéer, pågående arbete och så vidare. Idén med att ha ett projekt där vi spårar allt är att det [främjar bättre teamwork.](/insights/great-teamwork)

Varje post är en funktion eller del av en funktion. Så här går vi från "skulle det inte vara coolt om..." till "kolla in denna fantastiska nya funktion!"

Projektet har följande listor:

- **Idéer/Feedback**: Detta är en lista över teamidéer eller kundfeedback baserat på samtal eller e-postutbyten. Känn dig fri att lägga till idéer här! I denna lista har vi ännu inte beslutat att vi kommer att bygga några av dessa funktioner, men vi granskar detta regelbundet för idéer som vi vill utforska vidare.
- **Backlog (Långsiktig)**: Detta är där funktioner från Idéer/Feedback-listan går om vi beslutar att de skulle vara ett bra tillskott till Blue.
- **{Aktuellt Kvartal}**: Detta är vanligtvis strukturerat som "Qx ÅÅÅÅ" och visar våra kvartalsprioriteringar.
- **Buggar**: Detta är en lista över kända buggar som rapporterats av teamet eller kunder. Buggar som läggs till här får automatiskt taggen "Bugg".
- **Specifikationer**: Dessa funktioner specificeras för närvarande. Inte varje funktion kräver en specifikation eller design; det beror på den förväntade storleken på funktionen och den säkerhetsnivå vi har när det gäller kantfall och komplexitet.
- **Design Backlog**: Detta är backloggen för designers; varje gång de har avslutat något som är under arbete kan de välja ett objekt från denna lista.
- **Under Utveckling Design**: Detta är de aktuella funktionerna som designerna arbetar med.
- **Design Granskning**: Detta är där funktionerna vars designer för närvarande granskas.
- **Backlog (Kortsiktig)**: Detta är en lista över funktioner som vi sannolikt kommer att börja arbeta med under de kommande veckorna. Här sker tilldelningar. VD:n och teknikchefen beslutar vilka funktioner som tilldelas vilken ingenjör baserat på tidigare erfarenhet och arbetsbelastning. [Teammedlemmar kan sedan dra in dessa i Under Utveckling](/insights/push-vs-pull-kanban) när de har slutfört sitt nuvarande arbete.
- **Under Utveckling**: Dessa är funktioner som för närvarande utvecklas.
- **Kodgranskning**: När en funktion har avslutat utvecklingen genomgår den en kodgranskning. Då kommer den antingen att flyttas tillbaka till "Under Utveckling" om justeringar behövs eller distribueras till utvecklingsmiljön.
- **Utveckling**: Dessa är alla funktioner som för närvarande finns i utvecklingsmiljön. Andra teammedlemmar och vissa kunder kan granska dessa.
- **Beta**: Dessa är alla funktioner som för närvarande finns i [Beta-miljön](https://beta.app.blue.cc). Många kunder använder detta som sin dagliga Blue-plattform och ger också feedback.
- **Produktion**: När en funktion når produktion anses den vara klar.

Ibland, när vi utvecklar en funktion, inser vi att vissa underfunktioner är svårare att implementera än vi först förväntade oss, och vi kan välja att inte göra dem i den första versionen som vi distribuerar till kunder. I detta fall kan vi skapa en ny post med ett namn som följer formatet "{FeatureName} V2" och inkludera alla underfunktioner som checklistapunkter.

## Taggar

- **Mobil**: Detta betyder att funktionen är specifik för antingen våra iOS-, Android- eller iPad-appar.
- **{EnterpriseCustomerName}**: En funktion byggs specifikt för en företagskund. Spårning är viktig eftersom det vanligtvis finns ytterligare kommersiella avtal för varje funktion.
- **Bugg**: Detta betyder att detta är en bugg som kräver åtgärd.
- **Snabbspår**: Detta betyder att detta är en snabbspårändring som inte behöver gå igenom hela releasecykeln som beskrivs ovan.
- **Huvud**: Detta är en större funktionsutveckling. Det är vanligtvis reserverat för större infrastrukturarbete, stora beroendeuppgraderingar och betydande nya moduler inom Blue.
- **AI**: Denna funktion innehåller en komponent för artificiell intelligens.
- **Säkerhet**: Detta betyder att en säkerhetsimplikation måste granskas eller en patch krävs.


Snabbspårtaggen är särskilt intressant. Denna är reserverad för mindre, mindre komplexa uppdateringar som inte kräver vår fulla releasecykel och som vi vill skicka till kunder inom 24-48 timmar.

Snabbspårändringar är vanligtvis mindre justeringar som kan förbättra användarupplevelsen avsevärt utan att ändra kärnfunktionaliteten. Tänk på att fixa stavfel i användargränssnittet, justera knappavstånd eller lägga till nya ikoner för bättre visuell vägledning. Dessa är de typer av förändringar som, även om de är små, kan göra stor skillnad i hur användare uppfattar och interagerar med vår produkt. De är också irriterande om de tar lång tid att skicka!

Vår snabbspårsprocess är enkel.

Den börjar med att skapa en ny gren från huvudgrenen, implementera ändringarna och sedan skapa sammanslagningsförfrågningar för varje målgren - Utveckling, Beta och Produktion. Vi genererar en förhandsgranskningslänk för granskning, vilket säkerställer att även dessa små förändringar uppfyller våra kvalitetsstandarder. När de är godkända slås ändringarna samman samtidigt i alla grenar, vilket håller våra miljöer synkroniserade.

## Anpassade Fält

Vi har inte många anpassade fält i vårt produktprojekt.

- **Specifikationer**: Detta länkar till ett Blue-dokument som har specifikationen för den specifika funktionen. Detta görs inte alltid, eftersom det beror på funktionens komplexitet.
- **MR**: Detta är länken till Merge Request i [Gitlab](https://gitlab.com) där vi hostar vår kod.
- **Förhandsgranskningslänk**: För funktioner som främst ändrar front-end kan vi skapa en unik URL som har dessa ändringar för varje commit, så att vi enkelt kan granska ändringarna.
- **Ledare**: Detta fält berättar för oss vilken senior ingenjör som tar ledningen i kodgranskningen. Det säkerställer att varje funktion får den expertuppmärksamhet den förtjänar, och det finns alltid en tydlig kontaktperson för frågor eller bekymmer.

## Checklistor

Under våra veckovisa demonstrationer kommer vi att lägga den diskuterade feedbacken i en checklista som kallas "feedback" och det kommer också att finnas en annan checklista som innehåller den huvudsakliga [WBS (Work Breakdown Scope)](/insights/simple-work-breakdown-structure) för funktionen, så att vi enkelt kan se vad som är klart och vad som återstår att göra.

## Slutsats

Och det är det!

Vi tror att människor ibland blir överraskade över hur enkelt vår process är, men vi anser att enkla processer ofta är mycket överlägsna än överdrivet komplexa processer som man inte lätt kan förstå.

Denna enkelhet är avsiktlig. Den gör att vi kan förbli agila, snabbt svara på kundernas behov och hålla hela vårt team i linje.

Genom att använda Blue för att bygga Blue utvecklar vi inte bara en produkt – vi lever det.

Så nästa gång du använder Blue, kom ihåg: du använder inte bara en produkt vi har byggt. Du använder en produkt som vi personligen är beroende av varje dag.

Och det gör all skillnad.
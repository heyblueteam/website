---
title:  Varför Blue har en öppen beta
description: Lär dig varför vårt projektledningssystem har en pågående öppen beta.
category: "Engineering"
date: 2024-08-03
---



Många B2B SaaS-startups lanserar i beta, och av goda skäl. Det är en del av det traditionella Silicon Valley-mottot *“move fast and break things”*.

Att sätta en “beta”-etikett på en produkt sänker förväntningarna.

Är något trasigt? Jaha, det är bara en beta.

Är systemet långsamt? Jaha, det är bara en beta.

[Dokumentationen](https://blue.cc/docs) är icke-existerande? Jaha…du förstår poängen.

Och detta är *faktiskt* en bra sak. Reid Hoffman, grundaren av LinkedIn, sa berömt:

> Om du inte skäms över den första versionen av din produkt, har du lanserat för sent.

Och beta-etiketten är också bra för kunderna. Den hjälper dem att självvälja.

Kunderna som testar beta-produkter är de som befinner sig i de tidiga faserna av Teknologiadoptionslivscykeln, även känd som Produktadoptionskurvan.

Teknologiadoptionslivscykeln är vanligtvis indelad i fem huvudsegment:

1. Innovatörer
2. Tidiga användare
3. Tidig majoritet
4. Sen majoritet
5. Eftersläntrare

![](/insights/technology-adoption-lifecycle-graph.png)


Men så småningom måste produkten mogna, och kunderna förväntar sig en stabil, fungerande produkt. De vill inte ha tillgång till en “beta”-miljö där saker går sönder.

Eller vill de?

*Detta* är frågan vi ställde oss själva.

Vi tror att vi ställde oss denna fråga på grund av hur Blue ursprungligen byggdes. [Blue började som en avknoppning av en upptagen designbyrå](/insights/agency-success-playbook), och vi arbetade *inne* på kontoret hos ett företag som aktivt använde Blue för att driva alla sina projekt.

Detta innebär att vi under flera år kunde observera hur *verkliga* människor — som satt precis bredvid oss! — använde Blue i sina dagliga liv.

Och eftersom de använde Blue från de tidiga dagarna, använde detta team alltid Blue Beta!

Och därför var det naturligt för oss att låta alla våra andra kunder använda det också.

**Och detta är varför vi inte har ett dedikerat testteam.**

Det stämmer.

Ingen på Blue har det *enda* ansvaret för att säkerställa att vår plattform fungerar bra och stabilt.

Detta beror på flera skäl.

Det första är en lägre kostnadsbas.

Att inte ha ett heltids testteam minskar våra kostnader avsevärt, och vi kan överföra dessa besparingar till våra kunder med de lägsta priserna i branschen.

För att sätta detta i perspektiv, erbjuder vi företagsnivå-funktioner som vår konkurrens tar $30-$55/användare/månad för, för endast $7/månad.

Detta händer inte av en slump, det är *medvetet*.

Men det är inte en bra strategi att sälja en billigare produkt om den inte fungerar.

Så *den verkliga frågan är*, hur lyckas vi skapa en stabil plattform som tusentals kunder kan använda utan ett dedikerat testteam?

Självklart är vår strategi för att ha en öppen beta avgörande för detta, men innan vi dyker ner i detta, vill vi beröra utvecklaransvar.

Vi fattade det tidiga beslutet på Blue att vi aldrig skulle dela ansvar för front-end och back-end teknologier. Vi skulle endast anställa eller utbilda fullstack-utvecklare.

Anledningen till att vi fattade detta beslut var att säkerställa att en utvecklare skulle äga den funktion de arbetade med fullt ut. Så det skulle inte finnas någon av *“släng problemet över trädgårdsstaketet”* mentaliteten som man ibland får när det finns gemensamt ansvar för funktioner.

Och detta sträcker sig till testningen av funktionen, att förstå kundernas användningsfall och förfrågningar, samt att läsa och kommentera specifikationerna.

Med andra ord, varje utvecklare bygger en djup och intuitiv förståelse för den funktion de bygger.

Okej, låt oss nu prata om vår öppna beta.

När vi säger att den är “öppen” — menar vi det. Varje kund kan prova den helt enkelt genom att lägga till “beta” framför vår webbtillämpnings-URL.

Så “app.blue.cc” blir “beta.app.blue.cc”

När de gör detta kan de se sina vanliga data, eftersom både Beta- och Produktionsmiljöerna delar samma databas, men de kommer också att kunna se nya funktioner.

Kunder kan enkelt arbeta även om de har vissa teammedlemmar på Produktion och några nyfikna på Beta.

Vi har vanligtvis ett par hundra kunder som använder Beta vid varje given tidpunkt, och vi lägger upp funktionsförhandsvisningar på våra community-forum så att de kan kolla vad som är nytt och prova det.

Och detta är poängen: vi har *flera hundra* testare!

Alla dessa kunder kommer att prova funktioner i sina arbetsflöden och vara ganska verbala om något inte är helt rätt, eftersom de *redan* implementerar funktionen inom sitt företag!

Den vanligaste feedbacken är små men mycket användbara ändringar som adresserar kantfall som vi inte övervägde.

Vi lämnar nya funktioner på Beta mellan 2-4 veckor. När vi känner att de är stabila, släpper vi dem till produktion.

Vi har också möjlighet att kringgå Beta om det behövs, med hjälp av en snabbspårflagga. Detta görs vanligtvis för buggfixar som vi inte vill hålla för 2-4 veckor innan vi skickar till produktion.

Resultatet?

Att trycka till produktion känns…tja tråkigt! Som ingenting — det är helt enkelt ingen stor grej för oss.

Och det innebär att detta jämnar ut vårt releaseschema, vilket är vad som har möjliggjort att vi [kan skicka funktioner varje månad som klockverk under de senaste sex åren.](/changelog).

Men precis som med alla val finns det vissa avvägningar.

Kundsupport är något mer komplext, eftersom vi måste stödja kunder över två versioner av vår plattform. Ibland kan detta orsaka förvirring för kunder som har teammedlemmar som använder två olika versioner.

En annan smärtpunk är att denna strategi ibland kan sakta ner det övergripande releaseschemat till produktion. Detta är särskilt sant för större funktioner som kan bli "fast" i Beta om det finns en annan relaterad funktion som har problem och behöver mer arbete.

Men på det hela taget tycker vi att dessa avvägningar är värda fördelarna med en lägre kostnadsbas och större kundengagemang.

Vi är ett av få mjukvaruföretag som omfamnar denna strategi, men det är nu en grundläggande del av vår produktutvecklingsstrategi.
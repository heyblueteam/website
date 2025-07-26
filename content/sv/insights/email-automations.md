---
title: Hur man skapar anpassade e-postautomatiseringar
description: Anpassade e-postnotifikationer är en otroligt kraftfull funktion i Blue som kan hjälpa till att hålla arbetet framåt och säkerställa att kommunikationen är på autopilot.
category: "Product Updates"
---



E-postautomatiseringar i Blue är en [kraftfull projektledningsautomatisering](/platform/features/automations) för att effektivisera kommunikationen, säkerställa [bra teamwork](/insights/great-teamwork) och hålla projekten på rätt spår. Genom att utnyttja data som lagras inom dina poster kan du automatiskt skicka personliga e-postmeddelanden när vissa utlösare inträffar, såsom att en ny post skapas eller en uppgift blir försenad.

I den här artikeln kommer vi att utforska hur man ställer in och använder e-postautomatiseringar i Blue.

## Ställa in e-postautomatiseringar.

Att skapa en e-postautomatisering i Blue är en enkel process. Först, välj utlösaren som kommer att initiera det automatiserade e-postmeddelandet. Några vanliga utlösare inkluderar:

- En ny post skapas
- En tagg läggs till en post
- En post flyttas till en annan lista


Nästa steg är att konfigurera e-postdetaljerna, inklusive:

- Från namn och svara-till-adress
- Till adress (kan vara statisk eller dynamiskt hämtad från ett anpassat e-postfält)
- CC eller BCC-adresser (valfritt)

![](/insights/email-automations-image.webp)

En av de viktigaste fördelarna med e-postautomatiseringar i Blue är möjligheten att anpassa innehållet med hjälp av sammanfogningsetiketter. När du anpassar e-postens ämne och kropp kan du infoga sammanfogningsetiketter som refererar till specifik postdata, såsom postens namn eller värden från anpassade fält. Använd helt enkelt {klammerparentes} syntaxen för att infoga sammanfogningsetiketter.

Du kan också inkludera filbilagor genom att dra och släppa dem i e-postmeddelandet eller använda bilagans ikon. Filer från anpassade fält kan automatiskt bifogas om de är under 10 MB.

Innan du slutför din e-postautomatisering rekommenderas det att skicka ett test-e-postmeddelande till dig själv eller en kollega för att säkerställa att allt fungerar som det ska.

## Användningsfall och exempel

E-postautomatiseringar i Blue kan användas för en mängd olika syften. Här är några exempel:

1. Skicka ett bekräftelse-e-postmeddelande när en klient skickar en begäran via ett intagningsformulär. Ställ in utlösaren för att skicka ett e-postmeddelande när en ny post skapas genom formuläret, och se till att inkludera ett e-postfält i formuläret för att fånga klientens adress.
2. Meddela en tilldelad när en ny högprioriterad uppgift skapas. Ställ in utlösaren för att skicka ett e-postmeddelande när en "Prioritet"-tagg läggs till en post, och använd {Tilldelad} sammanfogningsetiketten för att automatiskt skicka e-postmeddelandet till den tilldelade användaren.
3. Skicka en enkät till en kund efter att en supportärende har markerats som löst. Ställ in utlösaren för att skicka ett e-postmeddelande när en post markeras som slutförd och flyttas till "Färdig"-listan. Inkludera kundens e-post i ett anpassat fält och ge detaljerad information om det lösta problemet i e-postens kropp.
4. Automatisera ett rekryteringsprogram genom att skicka bekräftelse-e-postmeddelanden till sökande. Ställ in utlösaren för att skicka ett e-postmeddelande när en ansökan skickas in via ett formulär och läggs till i "Mottagna"-listan. Fånga den sökandes e-post i formuläret och använd den för att skicka ett tackmeddelande.

## Fördelar med e-postautomatiseringar

E-postautomatiseringar i Blue erbjuder flera viktiga fördelar:

- Personlig kommunikation genom användning av sammanfogningsetiketter och data från anpassade fält
- Automatiska meddelanden som minskar manuellt arbete och säkerställer tidsenliga uppdateringar
- Strukturerade, datadrivna arbetsflöden som driver projekten framåt baserat på postdata

## Slutsats 

E-postautomatiseringar i Blue är ett värdefullt verktyg för att effektivisera kommunikationen och hålla projekten på rätt spår. Genom att utnyttja utlösare, sammanfogningsetiketter och data från anpassade fält kan du skapa personliga, automatiserade e-postmeddelanden som ökar ditt teams produktivitet och säkerställer att viktiga uppdateringar aldrig missas. Med ett brett utbud av användningsfall och enkel installation är e-postautomatiseringar en oumbärlig funktion för alla Blue-användare som vill optimera sitt arbetsflöde.
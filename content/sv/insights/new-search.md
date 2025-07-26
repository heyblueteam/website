---
title: Sökning i realtid
description: Blue presenterar en ny blixtsnabb sökmotor som returnerar resultat över alla dina projekt på millisekunder, vilket gör att du kan byta kontext på ett ögonblick.
category: "Product Updates"
date: 2024-03-01
---


Vi är glada att kunna meddela lanseringen av vår nya sökmotor, utformad för att revolutionera hur du hittar information inom Blue. Effektiv sökfunktionalitet är avgörande för sömlös projektledning, och vår senaste uppdatering säkerställer att du kan få tillgång till dina data snabbare än någonsin.

Vår nya sökmotor låter dig söka efter alla kommentarer, filer, poster, anpassade fält, beskrivningar och checklistor. Oavsett om du behöver hitta en specifik kommentar som gjorts på ett projekt, snabbt lokalisera en fil eller söka efter en viss post eller ett fält, ger vår sökmotor blixtsnabba resultat.

När verktyg når en responsivitet på 50-100 ms tenderar de att blekna bort och smälta in i bakgrunden, vilket ger en sömlös och nästan osynlig användarupplevelse. För sammanhang tar en mänsklig blinkning ungefär 60-120 ms, så 50 ms är faktiskt snabbare än en blinkning! Denna nivå av responsivitet gör att du kan interagera med Blue utan att ens inse att det finns där, vilket frigör dig att fokusera på det faktiska arbetet. Genom att utnyttja denna nivå av prestanda säkerställer vår nya sökmotor att du snabbt kan få tillgång till den information du behöver, utan att det någonsin kommer i vägen för ditt arbetsflöde.

För att uppnå vårt mål om blixtsnabb sökning har vi utnyttjat den senaste öppen källkodsteknologin. Vår sökmotor är byggd ovanpå MeiliSearch, en populär öppen källkodssökning som tjänst som använder naturlig språkbehandling och vektorsökning för att snabbt hitta relevanta resultat. Dessutom har vi implementerat lagring i minnet, vilket gör att vi kan lagra ofta åtkomna data i RAM, vilket minskar tiden det tar att returnera sökresultat. Denna kombination av MeiliSearch och lagring i minnet gör att vår sökmotor kan leverera resultat på millisekunder, vilket gör det möjligt för dig att snabbt hitta vad du behöver utan att behöva tänka på den underliggande teknologin.

Den nya sökfältet är bekvämt placerat på navigeringsfältet, vilket gör att du kan börja söka direkt. För en mer detaljerad sökupplevelse, tryck helt enkelt på Tab-tangenten medan du söker för att få tillgång till hela söksidan. Dessutom kan du snabbt aktivera sökfunktionen från var som helst med CMD/Ctrl+K-genvägen, vilket gör det ännu enklare att hitta vad du behöver.

<video autoplay loop muted playsinline>
  <source src="/videos/search-demo.mp4" type="video/mp4">
</video>

## Framtida Utvecklingar

Detta är bara början. Nu när vi har en nästa generations sökinfrastruktur kan vi göra några riktigt intressanta saker i framtiden.

Nästa steg kommer att vara semantisk sökning, vilket är en betydande förbättring av den typiska nyckelordsökningen. Låt oss förklara.

Denna funktion kommer att göra det möjligt för sökmotorn att förstå kontexten av dina frågor. Till exempel, att söka efter "hav" kommer att hämta relevanta dokument även om den exakta frasen inte används. Du kanske tänker "men jag skrev 'ocean' istället!" - och du har rätt. Sökmotorn kommer också att förstå likheten mellan "hav" och "ocean", och returnera relevanta dokument även om den exakta frasen inte används. Denna funktion är särskilt användbar när man söker efter dokument som innehåller tekniska termer, akronymer eller bara vanliga ord som har flera variationer eller stavfel.

En annan kommande funktion är möjligheten att söka efter bilder utifrån deras innehåll. För att uppnå detta kommer vi att bearbeta varje bild i ditt projekt och skapa en inbäddning för var och en. I hög nivå är en inbäddning en matematisk uppsättning koordinater som motsvarar betydelsen av en bild. Detta innebär att alla bilder kan sökas baserat på vad de innehåller, oavsett deras filnamn eller metadata. Tänk dig att söka efter "flödesschema" och hitta alla bilder relaterade till flödesscheman, *oavsett deras filnamn.*
---
title: Automatisering av projektledning — e-post till intressenter.
description: Ofta vill du ha kontroll över dina automatiseringar för projektledning
category: "Product Updates"
date: 2024-07-08
---


Vi har tidigare gått igenom hur man [skapar e-postautomatiseringar.](/insights/email-automations)

Men ofta finns det intressenter i projekt som bara behöver bli informerade när det är något *verkligen* viktigt.

Skulle det inte vara bra om det fanns en automatisering för projektledning där du, som projektledare, kunde ha kontroll över *exakt* när du ska informera en nyckelintressent med ett knapptryck?

Det visar sig att med Blue kan du göra just detta!

Idag ska vi lära oss hur man skapar en riktigt användbar automatisering för projektledning:

En kryssruta som automatiskt informerar en eller flera nyckelintressenter, och ger dem all viktig kontext om vad du informerar dem om. Som en bonus kommer vi också att lära oss hur man låser denna funktion så att endast vissa medlemmar i ditt projekt kan utlösa denna e-postnotifikation.

Detta kommer att se ut ungefär så här när du är klar:

![](/insights/checkbox-email-automation.png)

Och genom att bara kryssa i denna kryssruta kan du utlösa en automatisering för projektledning för att skicka ett anpassat e-postmeddelande till intressenter.

Låt oss gå steg för steg.

## 1. Skapa ditt anpassade fält för kryssruta

Detta är mycket enkelt, du kan kolla in vår [detaljerade dokumentation](https://documentation.blue.cc/custom-fields/introduction#creating-custom-fields) om hur man skapar anpassade fält.

Se till att du namnger detta fält något uppenbart som du kommer att komma ihåg, till exempel "informera ledningen" eller "informera intressenter".

## 2. Skapa din utlösare för automatisering av projektledning.

I postvyn i ditt projekt, klicka på den lilla roboten uppe till höger för att öppna inställningarna för automatisering:

<video autoplay loop muted playsinline>
  <source src="/videos/notify-stakeholders-automation-setup.mp4" type="video/mp4">
</video>

## 3. Skapa din åtgärd för automatisering av projektledning.

I det här fallet kommer vår åtgärd att vara att skicka en anpassad e-postnotifikation till en eller flera e-postadresser. Det är bra att notera här att dessa personer **inte** behöver vara i Blue för att ta emot dessa e-postmeddelanden, du kan skicka e-post till *vilken som helst* e-postadress.

Du kan lära dig mer i vår [detaljerade dokumentationsguide om hur man ställer in e-postautomatiseringar](https://documentation.blue.cc/automations/actions/email-automations)

Ditt slutresultat bör se ut ungefär så här:

![](/insights/email-automation-example.png)

## 4. Bonus: Begränsa åtkomst till kryssrutan.

Du kan använda [anpassade användarroller i Blue](/platform/features/user-permissions) för att begränsa åtkomsten till de anpassade fälten för kryssrutor, vilket säkerställer att endast auktoriserade teammedlemmar kan utlösa e-postnotifikationer.

Blue tillåter projektadministratörer att definiera roller och tilldela behörigheter till användargrupper. Detta system är avgörande för att upprätthålla kontroll över vem som kan interagera med specifika element i ditt projekt, inklusive anpassade fält som notifikationskryssrutan.

1. Navigera till avsnittet Användarhantering i Blue och välj "Anpassade användarroller."
2. Skapa en ny roll genom att ge den ett beskrivande namn och en valfri beskrivning.
3. Inom rollens behörigheter, lokalisera avsnittet för Åtkomst till anpassade fält.
4. Ange om rollen kan se eller redigera det anpassade fältet för kryssruta. Till exempel, begränsa redigeringsåtkomst till roller som "Projektadministratör" medan du tillåter en ny skapad anpassad roll att hantera detta fält.
5. Tilldela den nyss skapade rollen till lämpliga användare eller användargrupper. Detta säkerställer att endast de utsedda individerna har möjlighet att interagera med notifikationskryssrutan.

[Läs mer på vår officiella dokumentationssida.](https://documentation.blue.cc/user-management/roles/custom-user-roles)

Genom att implementera dessa anpassade roller förbättrar du säkerheten och integriteten i dina processer för projektledning. Endast auktoriserade teammedlemmar kan utlösa kritiska e-postnotifikationer, vilket säkerställer att intressenter får viktiga uppdateringar utan onödiga varningar.

## Slutsats

Genom att implementera automatiseringen för projektledning som beskrivs ovan får du exakt kontroll över när och hur du informerar nyckelintressenter. Denna metod säkerställer att viktiga uppdateringar kommuniceras effektivt, utan att överväldiga dina intressenter med onödig information. Genom att använda Blues anpassade fält och automatiseringsfunktioner kan du strömlinjeforma din projektledningsprocess, förbättra kommunikationen och upprätthålla en hög nivå av effektivitet.

Med bara en enkel kryssruta kan du utlösa anpassade e-postnotifikationer anpassade efter ditt projekts behov, vilket säkerställer att rätt personer informeras vid rätt tidpunkt. Dessutom ger möjligheten att begränsa denna funktionalitet till specifika teammedlemmar ett extra lager av kontroll och säkerhet.

Börja utnyttja denna kraftfulla funktion i Blue idag för att hålla dina intressenter informerade och dina projekt igång smidigt. För mer detaljerade steg och ytterligare anpassningsalternativ, se de angivna dokumentationslänkarna. Lycka till med automatiseringen!
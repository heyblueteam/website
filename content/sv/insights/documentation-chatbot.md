---
title: Varför vi byggde vår egen AI-dokumentationschattbot
description: Vi byggde vår egen dokumentations-AI-chattbot som är tränad på Blue-plattformens dokumentation.
category: "Product Updates"
date: 2024-07-09
---



På Blue letar vi alltid efter sätt att göra livet enklare för våra kunder. Vi har [djupgående dokumentation av varje funktion](https://documentation.blue.cc), [YouTube-videor](https://www.youtube.com/@HeyBlueTeam), [Tips och tricks](/insights/tips-tricks) och [olika supportkanaler](/support). 

Vi har hållit ett nära öga på utvecklingen av AI (Artificiell Intelligens) eftersom vi är mycket intresserade av [automatiseringar inom projektledning](/platform/features/automations). Vi har också släppt funktioner som [AI Auto Kategorisering](/insights/ai-auto-categorization) och [AI Sammanfattningar](/insights/ai-content-summarization) för att göra arbetet enklare för våra tusentals kunder. 

En sak som är klar är att AI är här för att stanna, och det kommer att ha en otrolig effekt på de flesta industrier, och projektledning är inget undantag. Så vi frågade oss hur vi ytterligare kunde utnyttja AI för att hjälpa hela livscykeln för en kund, från upptäckten, försäljning, onboarding och även med pågående frågor.

Svaret var ganska klart: **Vi behövde en AI-chattbot tränad på vår dokumentation.**

Låt oss inse det: *varje* organisation borde förmodligen ha en chattbot. De är fantastiska sätt för kunder att få omedelbara svar på typiska frågor, utan att behöva gräva igenom sidor av tät dokumentation eller din webbplats. Betydelsen av chattbotar på moderna marknadsföringswebbplatser kan inte överskattas. 

![](/insights/ai-chatbot-regular.png)

För mjukvaruföretag specifikt bör man inte betrakta marknadsföringswebbplatsen som en separat "sak" — den *är* en del av din produkt. Detta beror på att den passar in i den typiska kundlivscykeln:

- **Medvetenhet** (Upptäckte): Det är här potentiella kunder först snubblar över din fantastiska produkt. Din chattbot kan vara deras vänliga guide, som pekar dem mot viktiga funktioner och fördelar direkt från början.
- **Övervägande** (Utbildning): Nu är de nyfikna och vill lära sig mer. Din chattbot blir deras personliga handledare, som delar ut information skräddarsydd för deras specifika behov och frågor.
- **Köp/Omvandling**: Detta är sanningen - när en potentiell kund bestämmer sig för att ta steget och bli en kund. Din chattbot kan jämna ut eventuella sista-minuten-problem, svara på de "precis innan jag köper"-frågorna och kanske till och med ge ett bra erbjudande för att stänga affären.
- **Onboarding**: De har köpt in, vad nu? Din chattbot förvandlas till en hjälpsam sidekick, som går igenom inställningen med nya användare, visar dem repen och ser till att de inte känner sig vilse i din produkts underland.
- **Behållning**: Att hålla kunder nöjda är spelets namn. Din chattbot är tillgänglig dygnet runt, redo att lösa problem, erbjuda tips och tricks och se till att dina kunder känner kärleken.
- **Expansion**: Dags att ta det till nästa nivå! Din chattbot kan subtilt föreslå nya funktioner, uppgraderingar eller korsförsäljningar som stämmer överens med hur kunden redan använder din produkt. Det är som att ha en riktigt smart, icke-påträngande säljare alltid redo.
- **Advokatur**: Nöjda kunder blir dina största hejaklackare. Din chattbot kan uppmuntra nöjda användare att sprida ordet, lämna recensioner eller delta i referensprogram. Det är som att ha en hype-maskin inbyggd direkt i din produkt!

## Bygg vs Köp Beslut

När vi bestämde oss för att implementera en AI-chattbot var nästa stora fråga: bygga eller köpa? Som ett litet team som är laserfokuserat på vår kärnprodukt föredrar vi generellt "som-en-tjänst"-lösningar eller populära open-source-plattformar. Vi är trots allt inte i affären att återuppfinna hjulet för varje del av vår teknikstack.
Så, vi rullade upp ärmarna och dykt ner i marknaden, på jakt efter både betalda och open-source AI-chattbot-lösningar. 

Våra krav var tydliga, men icke-förhandlingsbara:

- **Omärkt Upplevelse**: Denna chattbot är inte bara en trevlig widget; den ska gå på vår marknadsföringswebbplats och så småningom i vår produkt. Vi är inte intresserade av att annonsera någon annans varumärke i vår egen digitala fastighet.
- **Bra UX**: För många potentiella kunder kan denna chattbot vara deras första kontaktpunkt med Blue. Den sätter tonen för deras uppfattning om vårt företag. Låt oss inse det: om vi inte kan få till en ordentlig chattbot på vår webbplats, hur kan vi då förvänta oss att kunder ska lita på oss med sina kritiska projekt och processer?
- **Rimlig Kostnad**: Med en stor användarbas och planer på att integrera chattboten i vår kärnprodukt behövde vi en lösning som inte skulle kosta skjortan när användningen ökar. Idealiskt ville vi ha ett **BYOK (Bring Your Own Key) alternativ**. Detta skulle tillåta oss att använda vår egen OpenAI- eller annan AI-tjänstnyckel, betala för direkta variabla kostnader istället för ett påslag till en tredje part som faktiskt inte kör modellerna.
- **Kompatibel med OpenAI Assistants API**: Om vi skulle gå med en open-source-lösning ville vi inte ha besväret med att hantera en pipeline för dokumentinmatning, indexering, vektordatabaser och allt det där. Vi ville använda [OpenAI Assistants API](https://platform.openai.com/docs/assistants/overview) som skulle abstrahera bort all komplexitet bakom en API. Ärligt talat — detta är verkligen välgjort. 
- **Skalbarhet**: Vi vill ha denna chattbot på flera ställen, med potentiellt tiotusentals användare per år. Vi förväntar oss betydande användning, och vi vill inte vara låsta till en lösning som inte kan växa med våra behov.

## Kommersiella AI Chattbotar

De vi granskade hade en tendens att ha en bättre UX än open-source-lösningar — som tyvärr ofta är fallet. Det finns förmodligen en separat diskussion att ha en dag om *varför* många open-source-lösningar ignorerar eller undervärderar vikten av UX. 

Vi kommer att ge en lista här, ifall du letar efter några solida kommersiella erbjudanden:

- **[Chatbase](https://chatbase.co):** Chatbase låter dig bygga en anpassad AI-chattbot tränad på din kunskapsbas och lägga till den på din webbplats eller interagera med den genom deras API. Den erbjuder funktioner som pålitliga svar, leadgenerering, avancerad analys och möjligheten att koppla till flera datakällor. För oss kändes detta som ett av de mest polerade kommersiella erbjudandena där ute. 
- **[DocsBot AI](https://docsbot.ai/):** DocsBot AI skapar anpassade ChatGPT-botar tränade på din dokumentation och innehåll för support, försäljning, forskning och mer. Den tillhandahåller inbäddningswidgetar för att enkelt lägga till chattboten på din webbplats, möjligheten att automatiskt svara på supportärenden och ett kraftfullt API för integration.
- **[CustomGPT.ai](https://customgpt.ai):** CustomGPT.ai skapar en personlig chattbotupplevelse genom att ta in dina affärsdata, inklusive webbplatsinnehåll, hjälpdesk, kunskapsbaser, dokument och mer. Det gör att leads kan ställa frågor och få omedelbara svar baserat på ditt innehåll, utan att behöva söka. Intressant nog påstår de också [att slå OpenAI i RAG (Retrieval Augmented Generation)!](https://customgpt.ai/customgpt-beats-open-ai-in-rag-benchmark/)
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Detta är ett intressant kommersiellt erbjudande, eftersom det *också* råkar vara open-source mjukvara. Det verkar vara lite i tidig utvecklingsfas, och prissättningen kändes inte realistisk (27 dollar/månad för obegränsade meddelanden kommer aldrig att fungera kommersiellt för dem).

Vi tittade också på [InterCom Fin](https://www.intercom.com/fin) som är en del av deras kundsupportmjukvara. Detta skulle ha inneburit att vi bytte bort från [HelpScout](https://wwww.helpscout.com) som vi har använt sedan vi startade Blue. Detta kunde ha varit möjligt, men InterCom Fin har en galen prissättning som helt enkelt uteslöt det från övervägande.

Och detta är faktiskt problemet med många av de kommersiella erbjudandena. InterCom Fin tar ut 0,99 dollar per kundsupportförfrågan som hanteras, och ChatBase tar ut 399 dollar/månad för 40 000 meddelanden. Det är nästan 5 000 dollar per år för en enkel chattwidget. 

Med tanke på att priserna för AI-inferens sjunker som en sten. OpenAI sänkte sina priser ganska dramatiskt:

- Den ursprungliga GPT-4 (8k kontext) var prissatt till 0,03 dollar per 1K prompttokens.
- GPT-4 Turbo (128k kontext) var prissatt till 0,01 dollar per 1K prompttokens, en 50% reduktion från den ursprungliga GPT-4.
- GPT-4o-modellen är prissatt till 0,005 dollar per 1K tokens, vilket är en ytterligare 50% reduktion från GPT-4 Turbo-prissättningen.

Det är en 83% reduktion i kostnader, och vi förväntar oss inte att det kommer att förbli stillastående. 

Med tanke på att vi letade efter en skalbar lösning som skulle användas av tiotusentals användare per år med en betydande mängd meddelanden, är det logiskt att gå direkt till källan och betala för API-kostnaderna direkt, istället för att använda en kommersiell version som lägger på kostnaderna.

## Open Source AI Chattbotar

Som nämnts var de open-source-alternativ vi granskade mestadels besvikande när det gäller kravet på "Bra UX". 

Vi tittade på:

- **[Deepchat](https://deepchat.dev/)**: Detta är en ramagnostisk chattkomponent för AI-tjänster, som kopplar till olika AI-API:er inklusive OpenAI. Den har också möjligheten för användare att ladda ner en AI-modell som körs direkt i webbläsaren. Vi lekte med detta och fick en version att fungera, men OpenAI Assistants API som implementerades kändes ganska buggig med flera problem. Men detta är ett mycket lovande projekt, och deras lekplats är verkligen välgjord. 
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: När vi ser på detta igen från ett open-source-perspektiv, skulle detta kräva att vi sätter upp ganska mycket infrastruktur, något som vi inte ville göra, eftersom vi ville förlita oss så mycket som möjligt på OpenAI:s Assistants API. 


## Bygga Vår Egen Chattbot

Och så, utan att kunna hitta något som matchade alla våra krav, bestämde vi oss för att bygga vår egen AI-chattbot som kunde interagera med OpenAI Assistants API. Detta visade sig i slutändan vara relativt smärtfritt! 

Vår webbplats använder [Nuxt3](https://nuxt.com), [Vue3](https://vuejs.org/) (som är samma ramverk som Blue-plattformen), och [TailwindUI](https://tailwindui.com/).

Det första steget var att skapa ett API (Application Programming Interface) i Nuxt3 som kan "prata" med OpenAI Assistants API. Detta var nödvändigt eftersom vi inte ville göra allt på front-end, eftersom detta skulle exponera vår OpenAI API-nyckel för världen, med potential för missbruk. 

Vårt backend-API fungerar som en säker mellanhand mellan användarens webbläsare och OpenAI. Här är vad det gör:

- **Konversationshantering:** Det skapar och hanterar "trådar" för varje konversation. Tänk på en tråd som en unik chatt-session som kommer ihåg allt du har sagt.
- **Meddelandehantering:** När du skickar ett meddelande lägger vårt API till det i rätt tråd och ber OpenAI:s assistent att skapa ett svar.
- **Smart Väntan:** Istället för att få dig att stirra på en laddningsskärm, kollar vårt API med OpenAI varje sekund för att se om ditt svar är klart. Det är som att ha en servitör som håller ett öga på din beställning utan att störa kocken varannan sekund.
- **Säkerhet Först:** Genom att hantera allt detta på servern håller vi dina data och våra API-nycklar säkra och trygga.

Sedan var det front-end och användarupplevelsen. Som diskuterats tidigare var detta *kritiskt* viktigt, eftersom vi inte får en andra chans att göra ett första intryck! 

När vi designade vår chattbot lade vi stor vikt vid användarupplevelsen, och säkerställde att varje interaktion är smidig, intuitiv och reflekterar Blues engagemang för kvalitet. Chattbotens gränssnitt börjar med en enkel, elegant blå cirkel, som använder [HeroIcons för våra ikoner](https://heroicons.com/) (som vi använder genom hela Blue-webbplatsen) för att fungera som vår chattbots öppningswidget. Detta designval säkerställer visuell konsistens och omedelbar varumärkesigenkänning.

![](/insights/ai-chatbot-circle.png)

Vi förstår att användare ibland kan behöva ytterligare support eller mer djupgående information. Därför har vi inkluderat praktiska länkar inom chattbotens gränssnitt. En e-postlänk för support är lätt tillgänglig, vilket gör att användare kan kontakta vårt team direkt om de behöver mer personlig hjälp. Dessutom har vi integrerat en dokumentationslänk, som ger enkel tillgång till mer omfattande resurser för dem som vill dyka djupare in i Blues erbjudanden.

Användarupplevelsen förbättras ytterligare av smakfulla fade-in och fade-up-animationer när chattbotens fönster öppnas. Dessa subtila animationer ger en touch av sofistikering till gränssnittet, vilket gör interaktionen mer dynamisk och engagerande. Vi har också implementerat en skrivindikator, en liten men avgörande funktion som låter användare veta att chattboten bearbetar deras fråga och skapar ett svar. Denna visuella ledtråd hjälper till att hantera användarens förväntningar och upprätthåller en känsla av aktiv kommunikation.

<video autoplay loop muted playsinline>
  <source src="/videos/ai-chatbot-animation.mp4" type="video/mp4">
</video>


Vi inser att vissa konversationer kan kräva mer skärmutrymme, så vi har lagt till möjligheten att öppna konversationen i ett större fönster. Denna funktion är särskilt användbar för längre utbyten eller när man granskar detaljerad information, vilket ger användare flexibiliteten att anpassa chattboten efter sina behov.

Bakom kulisserna har vi implementerat intelligent bearbetning för att optimera chattbotens svar. Vårt system parserar automatiskt AI:s svar för att ta bort referenser till våra interna dokument, vilket säkerställer att den information som presenteras är ren, relevant och fokuserad enbart på att adressera användarens fråga.
För att förbättra läsbarheten och möjliggöra mer nyanserad kommunikation har vi integrerat markdown-stöd med hjälp av 'marked'-biblioteket. Denna funktion gör det möjligt för vår AI att tillhandahålla rikt formattext, inklusive fet och kursiv betoning, strukturerade listor och till och med kodsnuttar när det är nödvändigt. Det är som att få ett välformat, skräddarsytt mini-dokument som svar på dina frågor.

Sist men inte minst har vi prioriterat säkerhet i vår implementation. Genom att använda DOMPurify-biblioteket sanerar vi HTML som genereras från markdown-parsing. Detta avgörande steg säkerställer att eventuella potentiellt skadliga skript eller kod tas bort innan innehållet visas för dig. Det är vårt sätt att garantera att den hjälpsamma information du får är både informativ och säker att konsumera.


## Framtida Utvecklingar

Så detta är bara början, vi har några spännande saker på väg för denna funktion. 

En av våra kommande funktioner är möjligheten att strömma svar i realtid. Snart kommer du att se chattbotens svar dyka upp tecken för tecken, vilket gör konversationerna mer naturliga och dynamiska. Det är som att se AI tänka, vilket skapar en mer engagerande och interaktiv upplevelse som håller dig informerad varje steg på vägen.

För våra värderade Blue-användare arbetar vi med personalisering. Chattboten kommer att känna igen när du är inloggad och anpassa sina svar baserat på din kontoinformation, användningshistorik och preferenser. Tänk dig en chattbot som inte bara svarar på dina frågor utan också förstår din specifika kontext inom Blue-ekosystemet, och erbjuder mer relevant och personlig hjälp.

Vi förstår att du kanske arbetar med flera projekt eller har olika frågor. Därför utvecklar vi möjligheten att upprätthålla flera distinkta konversationstrådar med vår chattbot. Denna funktion kommer att göra det möjligt för dig att växla mellan olika ämnen sömlöst, utan att tappa kontexten – precis som att ha flera flikar öppna i din webbläsare.

För att göra dina interaktioner ännu mer produktiva skapar vi en funktion som kommer att erbjuda föreslagna uppföljningsfrågor baserat på din aktuella konversation. Detta kommer att hjälpa dig att utforska ämnen djupare och upptäcka relaterad information som du kanske inte har tänkt på att fråga om, vilket gör varje chatt-session mer omfattande och värdefull.

Vi är också glada över att skapa en svit av specialiserade AI-assistenter, var och en skräddarsydd för specifika behov. Oavsett om du vill svara på frågor före försäljning, sätta upp ett nytt projekt eller felsöka avancerade funktioner, kommer du att kunna välja den assistent som bäst passar dina aktuella behov. Det är som att ha ett team av Blue-experter vid dina fingertoppar, var och en specialiserad på olika aspekter av vår plattform.

Slutligen arbetar vi på att låta dig ladda upp skärmdumpar direkt till chatten. AI:n kommer att analysera bilden och ge förklaringar eller felsökningssteg baserat på vad den ser. Denna funktion kommer att göra det enklare än någonsin att få hjälp med specifika problem du stöter på när du använder Blue, och överbrygga klyftan mellan visuell information och textuell hjälp.

## Slutsats

Vi hoppas att denna djupdykning i vår AI-chattbotutvecklingsprocess har gett några värdefulla insikter i vårt produktutvecklingstänkande på Blue. Vår resa från att identifiera behovet av en chattbot till att bygga vår egen lösning visar hur vi närmar oss beslutsfattande och innovation.

![](/insights/ai-chatbot-modal.png)

På Blue väger vi noggrant alternativen mellan att bygga och köpa, alltid med ett öga på vad som bäst tjänar våra användare och stämmer överens med våra långsiktiga mål. I det här fallet identifierade vi en betydande lucka på marknaden för en kostnadseffektiv men visuellt tilltalande chattbot som kunde möta våra specifika behov. Medan vi generellt förespråkar att utnyttja befintliga lösningar istället för att återuppfinna hjulet, är ibland den bästa vägen framåt att skapa något skräddarsytt för dina unika krav.

Vårt beslut att bygga vår egen chattbot togs inte lättvindigt. Det var resultatet av noggrann marknadsforskning, en tydlig förståelse för våra behov och ett engagemang för att ge den bästa möjliga upplevelsen för våra användare. Genom att utveckla internt kunde vi skapa en lösning som inte bara möter våra nuvarande behov utan också lägger grunden för framtida förbättringar och integrationer.

Detta projekt exemplifierar vår inställning på Blue: vi är inte rädda för att rulla upp ärmarna och bygga något från grunden när det är rätt val för vår produkt och våra användare. Det är denna vilja att gå den extra milen som gör att vi kan leverera innovativa lösningar som verkligen möter våra kunders behov.
Vi är entusiastiska över framtiden för vår AI-chattbot och det värde den kommer att tillföra både potentiella och befintliga Blue-användare. När vi fortsätter att förfina och utöka dess kapabiliteter förblir vi engagerade i att tänja på gränserna för vad som är möjligt inom projektledning och kundinteraktion.

Tack för att du följde med oss på denna resa genom vår utvecklingsprocess. Vi hoppas att det har gett dig en inblick i det genomtänkta, användarcentrerade tillvägagångssätt vi tar med varje aspekt av Blue. Håll utkik efter fler uppdateringar när vi fortsätter att utveckla och förbättra vår plattform för att bättre tjäna dig.

Om du är intresserad kan du hitta länken till källkoden för detta projekt här:

- **[ChatWidget](https://gitlab.com/bloohq/blue-website/-/blob/main/components/ChatWidget.vue)**: Detta är en Vue-komponent som driver chattwidgeten själv. 
- **[Chat API](https://gitlab.com/bloohq/blue-website/-/blob/main/server/api/chat.post.ts)**: Detta är mellanprogrammet som fungerar mellan chattkomponenten och OpenAI Assistants API.
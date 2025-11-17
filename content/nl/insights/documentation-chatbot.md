---
title: Waarom We Onze Eigen AI Documentatie Chatbot Hebben Gebouwd
description: We hebben onze eigen documentatie AI Chatbot gebouwd die is getraind op de documentatie van het Blue-platform.
category: "Product Updates"
date: 2024-07-09
---



Bij Blue zijn we altijd op zoek naar manieren om het leven van onze klanten gemakkelijker te maken. We hebben [uitgebreide documentatie van elke functie](https://documentation.blue.cc), [YouTube-video's](https://www.youtube.com/@workwithblue), [Tips & Tricks](/insights/tips-tricks) en [verschillende ondersteuningskanalen](/support). 

We hebben de ontwikkeling van AI (kunstmatige intelligentie) nauwlettend in de gaten gehouden, omdat we erg geïnteresseerd zijn in [projectmanagementautomatiseringen](/platform/features/automations). We hebben ook functies zoals [AI Auto Categorization](/insights/ai-auto-categorization) en [AI Samenvattingen](/insights/ai-content-summarization) uitgebracht om het werk gemakkelijker te maken voor onze duizenden klanten. 

Eén ding is duidelijk: AI is hier om te blijven, en het zal een ongelooflijke impact hebben op de meeste industrieën, en projectmanagement is daar geen uitzondering op. Dus vroegen we ons af hoe we AI verder konden benutten om de volledige levenscyclus van een klant te ondersteunen, van ontdekking, pre-sales, onboarding en ook bij doorlopende vragen.

Het antwoord was vrij duidelijk: **We hadden een AI-chatbot nodig die was getraind op onze documentatie.**

Laten we eerlijk zijn: *elke* organisatie zou waarschijnlijk een chatbot moeten hebben. Het zijn geweldige manieren voor klanten om directe antwoorden te krijgen op typische vragen, zonder door pagina's met dichte documentatie of uw website te hoeven bladeren. Het belang van chatbots op moderne marketingwebsites kan niet genoeg worden benadrukt. 

![](/insights/ai-chatbot-regular.png)

Voor softwarebedrijven in het bijzonder, moet men de marketingwebsite niet beschouwen als een afzonderlijk "ding" — het *is* een onderdeel van uw product. Dit komt omdat het past in de typische klantlevenscyclus:

- **Bewustzijn** (Ontdekking): Dit is waar potentiële klanten voor het eerst stuiten op uw geweldige product. Uw chatbot kan hun vriendelijke gids zijn, die hen meteen naar belangrijke functies en voordelen wijst.
- **Overweging** (Educatie): Nu zijn ze nieuwsgierig en willen ze meer leren. Uw chatbot wordt hun persoonlijke tutor, die informatie verstrekt die is afgestemd op hun specifieke behoeften en vragen.
- **Aankoop/Conversie**: Dit is het moment van de waarheid - wanneer een prospect besluit de sprong te wagen en klant te worden. Uw chatbot kan eventuele laatste hick-ups gladstrijken, die "net voordat ik koop" vragen beantwoorden en misschien zelfs een aantrekkelijke deal aanbieden om de deal te sluiten.
- **Onboarding**: Ze hebben gekocht, en nu? Uw chatbot transformeert in een behulpzame sidekick, die nieuwe gebruikers door de setup leidt, hen de kneepjes van het vak leert en ervoor zorgt dat ze zich niet verloren voelen in het wonderland van uw product.
- **Retentie**: Klanten tevreden houden is het doel. Uw chatbot is 24/7 beschikbaar, klaar om problemen op te lossen, tips en trucs aan te bieden, en ervoor te zorgen dat uw klanten zich gewaardeerd voelen.
- **Uitbreiding**: Tijd om naar een hoger niveau te tillen! Uw chatbot kan subtiel nieuwe functies, upsells of cross-sells voorstellen die aansluiten bij hoe de klant uw product al gebruikt. Het is alsof je een zeer slimme, niet-opdringerige verkoper altijd standby hebt.
- **Advocacy**: Tevreden klanten worden uw grootste aanhangers. Uw chatbot kan tevreden gebruikers aanmoedigen om het woord te verspreiden, beoordelingen achter te laten of deel te nemen aan doorverwijsprogramma's. Het is alsof je een hype-machine direct in je product hebt ingebouwd!

## Beslissing: Bouwen of Kopen

Zodra we besloten om een AI-chatbot te implementeren, was de volgende grote vraag: bouwen of kopen? Als een klein team dat zich volledig richt op ons kernproduct, geven we over het algemeen de voorkeur aan "as-a-service" oplossingen of populaire open-source platforms. We zijn tenslotte niet in de business van het opnieuw uitvinden van het wiel voor elk onderdeel van onze techstack.
Dus, we hebben onze mouwen opgestroopt en zijn de markt ingedoken, op zoek naar zowel betaalde als open-source AI-chatbotoplossingen. 

Onze vereisten waren eenvoudig, maar niet onderhandelbaar:

- **Ongebrandde Ervaring**: Deze chatbot is niet zomaar een leuk widget; hij komt op onze marketingwebsite en uiteindelijk in ons product. We zijn niet enthousiast over het adverteren van iemand anders zijn merk in onze eigen digitale ruimte.
- **Geweldige UX**: Voor veel potentiële klanten kan deze chatbot hun eerste contactpunt met Blue zijn. Het zet de toon voor hun perceptie van ons bedrijf. Laten we eerlijk zijn: als we geen goede chatbot op onze website kunnen neerzetten, hoe kunnen we dan verwachten dat klanten ons vertrouwen met hun cruciale projecten en processen?
- **Redelijke Kosten**: Met een grote gebruikersbasis en plannen om de chatbot in ons kernproduct te integreren, hadden we een oplossing nodig die de bank niet zou breken naarmate het gebruik toeneemt. Idealiter wilden we een **BYOK (Bring Your Own Key) optie**. Dit zou ons in staat stellen om onze eigen OpenAI of andere AI-service sleutel te gebruiken, en alleen te betalen voor directe variabele kosten in plaats van een opslagprijs aan een derde partij die de modellen niet daadwerkelijk draait.
- **Compatibel met OpenAI Assistants API**: Als we met open-source software zouden gaan, wilden we niet de rompslomp hebben van het beheren van een pijplijn voor documentinname, indexering, vector databases, en al dat soort dingen. We wilden de [OpenAI Assistants API](https://platform.openai.com/docs/assistants/overview) gebruiken die alle complexiteit achter een API zou abstraheren. Eerlijk gezegd - dit is echt goed gedaan. 
- **Schaalbaarheid**: We willen deze chatbot op meerdere plaatsen hebben, met mogelijk tienduizenden gebruikers per jaar. We verwachten aanzienlijk gebruik, en we willen niet vastzitten aan een oplossing die niet kan opschalen met onze behoeften.

## Commerciële AI Chatbots

De chatbots die we hebben beoordeeld, hadden over het algemeen een betere UX dan open-source oplossingen — zoals helaas vaak het geval is. Er is waarschijnlijk een aparte discussie te voeren over *waarom* veel open-source oplossingen het belang van UX negeren of onderbelichten. 

We zullen hier een lijst geven, voor het geval je op zoek bent naar enkele solide commerciële aanbiedingen:

- **[Chatbase](https://chatbase.co):** Chatbase stelt je in staat om een aangepaste AI-chatbot te bouwen die is getraind op je kennisbasis en deze toe te voegen aan je website of ermee te communiceren via hun API. Het biedt functies zoals betrouwbare antwoorden, leadgeneratie, geavanceerde analyses en de mogelijkheid om verbinding te maken met meerdere gegevensbronnen. Voor ons voelde dit als een van de meest gepolijste commerciële aanbiedingen die er zijn. 
- **[DocsBot AI](https://docsbot.ai/):** DocsBot AI creëert aangepaste ChatGPT-bots die zijn getraind op je documentatie en inhoud voor ondersteuning, pre-sales, onderzoek en meer. Het biedt inbedbare widgets om de chatbot eenvoudig aan je website toe te voegen, de mogelijkheid om automatisch op ondersteuningsverzoeken te reageren, en een krachtige API voor integratie.
- **[CustomGPT.ai](https://customgpt.ai):** CustomGPT.ai creëert een persoonlijke chatbotervaring door je bedrijfsgegevens in te nemen, inclusief website-inhoud, helpdesk, kennisbases, documenten en meer. Het stelt leads in staat om vragen te stellen en directe antwoorden te krijgen op basis van je inhoud, zonder te hoeven zoeken. Interessant genoeg, ze [beweren OpenAI te verslaan in RAG (Retrieval Augmented Generation)!](https://customgpt.ai/customgpt-beats-open-ai-in-rag-benchmark/)
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Dit is een interessante commerciële aanbieding, omdat het *ook* open-source software blijkt te zijn. Het lijkt een beetje in de vroege fase te zijn, en de prijsstelling voelde niet realistisch ($27/maand voor onbeperkte berichten zal commercieel nooit werken voor hen).

We hebben ook gekeken naar [InterCom Fin](https://www.intercom.com/fin) dat deel uitmaakt van hun klantenondersteuningssoftware. Dit zou hebben betekend dat we zouden moeten overstappen van [HelpScout](https://wwww.helpscout.com) dat we sinds de start van Blue gebruiken. Dit zou mogelijk zijn geweest, maar InterCom Fin heeft een krankzinnige prijsstelling die het simpelweg uitsloot uit overweging.

En dit is eigenlijk het probleem met veel van de commerciële aanbiedingen. InterCom Fin rekent $0,99 per klantenondersteuningsverzoek dat wordt afgehandeld, en ChatBase rekent $399/maand voor 40.000 berichten. Dat is bijna $5k per jaar voor een eenvoudige chatwidget. 

Gezien het feit dat de prijzen voor AI-inferentie als gek dalen. OpenAI heeft hun prijzen behoorlijk drastisch verlaagd:

- De oorspronkelijke GPT-4 (8k context) was geprijsd op $0,03 per 1K prompt tokens.
- De GPT-4 Turbo (128k context) was geprijsd op $0,01 per 1K prompt tokens, een vermindering van 50% ten opzichte van de oorspronkelijke GPT-4.
- Het GPT-4o-model is geprijsd op $0,005 per 1K tokens, wat een verdere vermindering van 50% is ten opzichte van de prijsstelling van GPT-4 Turbo.

Dat is een reductie van 83% in kosten, en we verwachten niet dat dit stil blijft staan. 

Gezien het feit dat we op zoek waren naar een schaalbare oplossing die door tienduizenden gebruikers per jaar zou worden gebruikt met een aanzienlijk aantal berichten, is het logisch om rechtstreeks naar de bron te gaan en de API-kosten rechtstreeks te betalen, en geen commerciële versie te gebruiken die de kosten verhoogt.

## Open Source AI Chatbots

Zoals eerder vermeld, waren de open-source opties die we hebben beoordeeld over het algemeen teleurstellend met betrekking tot de vereiste "Geweldige UX". 

We hebben gekeken naar:

- **[Deepchat](https://deepchat.dev/)**: Dit is een framework-onafhankelijke chatcomponent voor AI-diensten, die verbinding maakt met verschillende AI-API's, waaronder OpenAI. Het heeft ook de mogelijkheid voor gebruikers om een AI-model te downloaden dat rechtstreeks in de browser draait. We hebben hier mee gespeeld en kregen een werkende versie, maar de geïmplementeerde OpenAI Assistants API voelde behoorlijk buggy aan met verschillende problemen. Dit is echter een veelbelovend project, en hun playground is echt mooi gedaan. 
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Als we dit opnieuw bekijken vanuit een open-source perspectief, zou dit vereisen dat we behoorlijk wat infrastructuur opzetten, iets dat we niet wilden doen, omdat we zoveel mogelijk op de Assistants API van OpenAI wilden vertrouwen. 


## Onze Eigen ChatBot Bouwen

En dus, zonder iets te kunnen vinden dat aan al onze vereisten voldeed, besloten we onze eigen AI-chatbot te bouwen die kon communiceren met de OpenAI Assistants API. Dit bleek uiteindelijk relatief pijnloos te zijn! 

Onze website gebruikt [Nuxt3](https://nuxt.com), [Vue3](https://vuejs.org/) (wat hetzelfde framework is als het Blue-platform), en [TailwindUI](https://tailwindui.com/).

De eerste stap was om een API (Application Programming Interface) in Nuxt3 te creëren die kan "communiceren" met de OpenAI Assistants API. Dit was noodzakelijk omdat we niet alles aan de voorkant wilden doen, omdat dit onze OpenAI API-sleutel aan de wereld zou blootstellen, met het potentieel voor misbruik. 

Onze backend API fungeert als een veilige tussenpersoon tussen de browser van de gebruiker en OpenAI. Dit is wat het doet:

- **Gespreksbeheer:** Het creëert en beheert "threads" voor elk gesprek. Denk aan een thread als een unieke chatsessie die zich alles herinnert wat je hebt gezegd.
- **Berichtverwerking:** Wanneer je een bericht verzendt, voegt onze API het toe aan de juiste thread en vraagt de assistent van OpenAI om een antwoord te formuleren.
- **Slim Wachten:** In plaats van je naar een laadscherm te laten staren, controleert onze API elke seconde bij OpenAI of je antwoord klaar is. Het is alsof je een ober hebt die je bestelling in de gaten houdt zonder de chef elke twee seconden te storen.
- **Veiligheid Voorop:** Door dit alles op de server te verwerken, houden we je gegevens en onze API-sleutels veilig en wel.

Daarna was er de front-end en gebruikerservaring. Zoals eerder besproken, was dit *cruciaal* belangrijk, omdat we geen tweede kans krijgen om een eerste indruk te maken! 

Bij het ontwerpen van onze chatbot hebben we met grote zorgvuldigheid aandacht besteed aan de gebruikerservaring, zodat elke interactie soepel, intuïtief en een afspiegeling is van Blue's toewijding aan kwaliteit. De chatbotinterface begint met een eenvoudige, elegante blauwe cirkel, waarbij we [HeroIcons voor onze iconen](https://heroicons.com/) gebruiken (die we door de hele Blue-website gebruiken) om als ons chatbot-opening-widget te fungeren. Deze ontwerpkeuze zorgt voor visuele consistentie en onmiddellijke merkherkenning.


![](/insights/ai-chatbot-circle.png)

We begrijpen dat gebruikers soms extra ondersteuning of meer diepgaande informatie nodig hebben. Daarom hebben we handige links binnen de chatbotinterface opgenomen. Een e-maillink voor ondersteuning is gemakkelijk beschikbaar, zodat gebruikers rechtstreeks contact kunnen opnemen met ons team als ze meer persoonlijke hulp nodig hebben. Daarnaast hebben we een documentatielink opgenomen, die gemakkelijke toegang biedt tot meer uitgebreide bronnen voor degenen die dieper in de aanbiedingen van Blue willen duiken.

De gebruikerservaring wordt verder verbeterd door smaakvolle fade-in en fade-up animaties bij het openen van het chatbotvenster. Deze subtiele animaties voegen een vleugje verfijning toe aan de interface, waardoor de interactie dynamischer en boeiender aanvoelt. We hebben ook een typindicator geïmplementeerd, een kleine maar cruciale functie die gebruikers laat weten dat de chatbot hun vraag aan het verwerken is en een antwoord aan het formuleren is. Deze visuele aanwijzing helpt de verwachtingen van de gebruiker te beheren en behoudt een gevoel van actieve communicatie.

<video autoplay loop muted playsinline>
  <source src="/videos/ai-chatbot-animation.mp4" type="video/mp4">
</video>


Erkenning dat sommige gesprekken meer schermruimte nodig hebben, hebben we de mogelijkheid toegevoegd om het gesprek in een groter venster te openen. Deze functie is bijzonder nuttig voor langere uitwisselingen of wanneer gedetailleerde informatie wordt bekeken, waardoor gebruikers de flexibiliteit hebben om de chatbot aan hun behoeften aan te passen.

Achter de schermen hebben we enkele intelligente verwerkingen geïmplementeerd om de antwoorden van de chatbot te optimaliseren. Ons systeem parseert automatisch de antwoorden van de AI om verwijzingen naar onze interne documenten te verwijderen, zodat de gepresenteerde informatie schoon, relevant en uitsluitend gericht is op het beantwoorden van de vraag van de gebruiker.
Om de leesbaarheid te verbeteren en meer genuanceerde communicatie mogelijk te maken, hebben we markdown-ondersteuning geïntegreerd met behulp van de 'marked' bibliotheek. Deze functie stelt onze AI in staat om rijkelijk opgemaakte tekst te bieden, inclusief vetgedrukte en cursieve nadruk, gestructureerde lijsten en zelfs codefragmenten wanneer dat nodig is. Het is alsof je een goed opgemaakt, op maat gemaakt mini-document ontvangt als antwoord op je vragen.

Last but certainly not least, hebben we veiligheid prioriteit gegeven in onze implementatie. Met behulp van de DOMPurify-bibliotheek saneren we de HTML die is gegenereerd uit markdown-parsing. Deze cruciale stap zorgt ervoor dat eventuele potentieel schadelijke scripts of code worden verwijderd voordat de inhoud aan jou wordt weergegeven. Het is onze manier om te garanderen dat de nuttige informatie die je ontvangt niet alleen informatief is, maar ook veilig om te consumeren.


## Toekomstige Ontwikkelingen

Dus dit is nog maar het begin, we hebben enkele spannende dingen op de roadmap voor deze functie. 

Een van onze aankomende functies is de mogelijkheid om antwoorden in realtime te streamen. Binnenkort zie je de antwoorden van de chatbot letter voor letter verschijnen, waardoor gesprekken natuurlijker en dynamischer aanvoelen. Het is alsof je de AI ziet denken, wat een meer boeiende en interactieve ervaring creëert die je bij elke stap op de hoogte houdt.

Voor onze gewaardeerde Blue-gebruikers werken we aan personalisatie. De chatbot zal herkennen wanneer je bent ingelogd, en zijn antwoorden afstemmen op basis van je accountinformatie, gebruiksgeschiedenis en voorkeuren. Stel je een chatbot voor die niet alleen je vragen beantwoordt, maar ook je specifieke context binnen het Blue-ecosysteem begrijpt, waardoor relevantere en gepersonaliseerde hulp wordt geboden.

We begrijpen dat je mogelijk aan meerdere projecten werkt of verschillende vragen hebt. Daarom ontwikkelen we de mogelijkheid om verschillende afzonderlijke gespreksthreads met onze chatbot te onderhouden. Deze functie stelt je in staat om naadloos tussen verschillende onderwerpen te schakelen, zonder de context te verliezen – net als het openen van meerdere tabbladen in je browser.

Om je interacties nog productiever te maken, creëren we een functie die voorgestelde vervolgvragen biedt op basis van je huidige gesprek. Dit zal je helpen om onderwerpen dieper te verkennen en gerelateerde informatie te ontdekken die je misschien niet had gedacht te vragen, waardoor elke chatsessie uitgebreider en waardevoller wordt.

We zijn ook enthousiast over het creëren van een suite van gespecialiseerde AI-assistenten, elk afgestemd op specifieke behoeften. Of je nu op zoek bent naar antwoorden op pre-sales vragen, een nieuw project wilt opzetten of geavanceerde functies wilt oplossen, je kunt de assistent kiezen die het beste past bij je huidige behoeften. Het is alsof je een team van Blue-experts binnen handbereik hebt, die elk gespecialiseerd zijn in verschillende aspecten van ons platform.

Tot slot werken we eraan om je de mogelijkheid te geven om screenshots rechtstreeks naar de chat te uploaden. De AI zal de afbeelding analyseren en uitleg of probleemoplossingsstappen bieden op basis van wat het ziet. Deze functie maakt het gemakkelijker dan ooit om hulp te krijgen bij specifieke problemen die je tegenkomt tijdens het gebruik van Blue, en overbrugt de kloof tussen visuele informatie en tekstuele ondersteuning.

## Conclusie

We hopen dat deze diepgaande blik op ons ontwikkelingsproces van de AI-chatbot waardevolle inzichten heeft gegeven in onze productontwikkelingsdenkwijze bij Blue. Onze reis van het identificeren van de behoefte aan een chatbot tot het bouwen van onze eigen oplossing toont aan hoe we beslissingen nemen en innoveren.

![](/insights/ai-chatbot-modal.png)

Bij Blue wegen we zorgvuldig de opties van bouwen versus kopen af, altijd met het oog op wat het beste onze gebruikers zal dienen en in lijn is met onze langetermijndoelen. In dit geval identificeerden we een significante kloof in de markt voor een kosteneffectieve maar visueel aantrekkelijke chatbot die aan onze specifieke behoeften kon voldoen. Hoewel we over het algemeen pleiten voor het benutten van bestaande oplossingen in plaats van het wiel opnieuw uit te vinden, is soms de beste weg vooruit om iets te creëren dat is afgestemd op jouw unieke vereisten.

Onze beslissing om onze eigen chatbot te bouwen werd niet lichtvaardig genomen. Het was het resultaat van grondig marktonderzoek, een duidelijk begrip van onze behoeften en een toewijding om de best mogelijke ervaring voor onze gebruikers te bieden. Door in-house te ontwikkelen, konden we een oplossing creëren die niet alleen aan onze huidige behoeften voldoet, maar ook de basis legt voor toekomstige verbeteringen en integraties.

Dit project is een voorbeeld van onze aanpak bij Blue: we zijn niet bang om onze mouwen op te stropen en iets vanaf nul te bouwen wanneer het de juiste keuze is voor ons product en onze gebruikers. Het is deze bereidheid om een stapje verder te gaan die ons in staat stelt om innovatieve oplossingen te leveren die echt voldoen aan de behoeften van onze klanten.
We zijn enthousiast over de toekomst van onze AI-chatbot en de waarde die het zal brengen voor zowel potentiële als bestaande Blue-gebruikers. Terwijl we blijven verfijnen en de mogelijkheden uitbreiden, blijven we ons inzetten om de grenzen van wat mogelijk is in projectmanagement en klantinteractie te verleggen.

Bedankt dat je met ons op deze reis door ons ontwikkelingsproces bent gegaan. We hopen dat het je een kijkje heeft gegeven in de doordachte, gebruikersgerichte benadering die we bij elk aspect van Blue hanteren. Blijf op de hoogte voor meer updates terwijl we blijven evolueren en ons platform verbeteren om jou beter van dienst te zijn.

Als je geïnteresseerd bent, kun je de link naar de broncode voor dit project hier vinden:

- **[ChatWidget](https://gitlab.com/bloohq/blue-website/-/blob/main/components/ChatWidget.vue)**: Dit is een Vue-component die de chatwidget zelf aandrijft. 
- **[Chat API](https://gitlab.com/bloohq/blue-website/-/blob/main/server/api/chat.post.ts)**: Dit is de middleware die werkt tussen de chatcomponent en de OpenAI Assistants API.
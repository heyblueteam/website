---
title: Zoekfunctie in real-time
description: Blue onthult een nieuwe razendsnelle zoekmachine die resultaten uit al uw projecten in milliseconden teruggeeft, zodat u in een oogwenk van context kunt wisselen.
category: "Product Updates"
date: 2024-03-01
---



We zijn verheugd om de lancering van onze nieuwe zoekmachine aan te kondigen, ontworpen om te revolutioneren hoe u informatie binnen Blue vindt. Efficiënte zoekfunctionaliteit is cruciaal voor naadloos projectbeheer, en onze laatste update zorgt ervoor dat u uw gegevens sneller dan ooit kunt benaderen.

Onze nieuwe zoekmachine stelt u in staat om te zoeken naar alle opmerkingen, bestanden, records, aangepaste velden, beschrijvingen en checklists. Of u nu een specifieke opmerking over een project wilt vinden, snel een bestand wilt lokaliseren of zoekt naar een bepaald record of veld, onze zoekmachine biedt bliksemsnelle resultaten.

Naarmate tools een responsiviteit van 50-100 ms benaderen, vervagen ze en mengen ze zich in de achtergrond, wat een naadloze en bijna onzichtbare gebruikerservaring biedt. Ter context: een menselijke knippering duurt ongeveer 60-120 ms, dus 50 ms is eigenlijk sneller dan een knippering van het oog! Dit niveau van responsiviteit stelt u in staat om met Blue te interageren zonder zelfs maar te beseffen dat het er is, waardoor u zich kunt concentreren op het werk dat voorhanden is. Door gebruik te maken van dit prestatieniveau, zorgt onze nieuwe zoekmachine ervoor dat u snel toegang heeft tot de informatie die u nodig heeft, zonder dat dit ooit in de weg staat van uw workflow.

Om ons doel van razendsnelle zoekopdrachten te bereiken, hebben we de nieuwste open-source technologieën benut. Onze zoekmachine is gebouwd op MeiliSearch, een populaire open-source zoek-as-een-service die natuurlijke taalverwerking en vectorzoektechnologie gebruikt om snel relevante resultaten te vinden. Bovendien hebben we in-memory opslag geïmplementeerd, waarmee we vaak geraadpleegde gegevens in RAM kunnen opslaan, waardoor de tijd die nodig is om zoekresultaten terug te geven, wordt verminderd. Deze combinatie van MeiliSearch en in-memory opslag stelt onze zoekmachine in staat om resultaten in milliseconden te leveren, waardoor het mogelijk is voor u om snel te vinden wat u nodig heeft zonder ooit na te hoeven denken over de onderliggende technologie.

De nieuwe zoekbalk is handig geplaatst op de navigatiebalk, zodat u direct kunt beginnen met zoeken. Voor een meer gedetailleerde zoekervaring drukt u eenvoudig op de Tab-toets terwijl u zoekt om de volledige zoekpagina te openen. Bovendien kunt u de zoekfunctie snel vanaf elke plek activeren met de CMD/Ctrl+K sneltoets, waardoor het nog gemakkelijker wordt om te vinden wat u nodig heeft.

<video autoplay loop muted playsinline>
  <source src="/videos/search-demo.mp4" type="video/mp4">
</video>


## Toekomstige Ontwikkelingen

Dit is nog maar het begin. Nu we een next-generation zoekinfrastructuur hebben, kunnen we in de toekomst enkele echt interessante dingen doen.

De volgende stap is semantische zoekopdracht, wat een aanzienlijke verbetering is ten opzichte van de typische zoekopdracht op basis van trefwoorden. Laat ons dit uitleggen.

Deze functie stelt de zoekmachine in staat om de context van uw zoekopdrachten te begrijpen. Bijvoorbeeld, zoeken naar "zee" zal relevante documenten opleveren, zelfs als de exacte zin niet wordt gebruikt. U denkt misschien "maar ik typte 'ocean' in plaats daarvan!" - en dat klopt. De zoekmachine zal ook de gelijkenis tussen "zee" en "ocean" begrijpen en relevante documenten teruggeven, zelfs als de exacte zin niet wordt gebruikt. Deze functie is bijzonder nuttig bij het zoeken naar documenten die technische termen, acroniemen of gewoon veelvoorkomende woorden bevatten die meerdere variaties of typfouten hebben.

Een andere aankomende functie is de mogelijkheid om naar afbeeldingen te zoeken op basis van hun inhoud. Om dit te bereiken, zullen we elke afbeelding in uw project verwerken en een embedding voor elke afbeelding creëren. In algemene termen is een embedding een wiskundige set coördinaten die overeenkomt met de betekenis van een afbeelding. Dit betekent dat alle afbeeldingen kunnen worden doorzocht op basis van wat ze bevatten, ongeacht hun bestandsnaam of metadata. Stel je voor dat je zoekt naar "stroomschema" en alle afbeeldingen vindt die verband houden met stroomschema's, *ongeacht hun bestandsnamen.*
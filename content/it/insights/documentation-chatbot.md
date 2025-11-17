---
title: Perché abbiamo costruito il nostro chatbot AI per la documentazione
description: Abbiamo costruito il nostro chatbot AI per la documentazione, addestrato sulla documentazione della piattaforma Blue.
category: "Product Updates"
date: 2024-07-09
---


At Blue, siamo sempre alla ricerca di modi per semplificare la vita ai nostri clienti. Abbiamo [documentazione approfondita su ogni funzionalità](https://documentation.blue.cc), [video su YouTube](https://www.youtube.com/@workwithblue), [consigli e trucchi](/insights/tips-tricks) e [vari canali di supporto](/support).

Abbiamo tenuto d'occhio lo sviluppo dell'AI (Intelligenza Artificiale) poiché siamo molto interessati alle [automazioni nella gestione dei progetti](/platform/features/automations). Abbiamo anche rilasciato funzionalità come [AI Auto Categorization](/insights/ai-auto-categorization) e [AI Summaries](/insights/ai-content-summarization) per semplificare il lavoro ai nostri migliaia di clienti.

Una cosa è chiara: l'AI è qui per restare e avrà un effetto incredibile su gran parte delle industrie, e la gestione dei progetti non fa eccezione. Così ci siamo chiesti come potessimo sfruttare ulteriormente l'AI per aiutare l'intero ciclo di vita di un cliente, dalla scoperta, pre-vendita, onboarding e anche con domande in corso.

La risposta era piuttosto chiara: **avevamo bisogno di un chatbot AI addestrato sulla nostra documentazione.**

Affrontiamolo: *ogni* organizzazione dovrebbe probabilmente avere un chatbot. Sono ottimi strumenti per i clienti per ottenere risposte immediate a domande tipiche, senza dover setacciare pagine di documentazione densa o il tuo sito web. L'importanza dei chatbot nei siti web di marketing moderni non può essere sottovalutata.

![](/insights/ai-chatbot-regular.png)

Per le aziende software in particolare, non si dovrebbe considerare il sito web di marketing come una "cosa" separata — *è* parte del tuo prodotto. Questo perché si inserisce nella vita tipica del cliente:

- **Consapevolezza** (Scoperta): Qui è dove i potenziali clienti si imbattono per la prima volta nel tuo fantastico prodotto. Il tuo chatbot può essere la loro guida amichevole, indirizzandoli verso funzionalità e vantaggi chiave fin da subito.
- **Considerazione** (Educazione): Ora sono curiosi e vogliono saperne di più. Il tuo chatbot diventa il loro tutor personale, fornendo informazioni su misura per le loro esigenze e domande specifiche.
- **Acquisto/Conversione**: Questo è il momento della verità - quando un potenziale cliente decide di tuffarsi e diventare un cliente. Il tuo chatbot può risolvere eventuali intoppi dell'ultimo minuto, rispondere a quelle domande "proprio prima di acquistare" e magari anche offrire un affare interessante per chiudere la vendita.
- **Onboarding**: Hanno effettuato l'acquisto, e ora? Il tuo chatbot si trasforma in un aiutante utile, guidando i nuovi utenti attraverso la configurazione, mostrando loro le basi e assicurandosi che non si sentano persi nel meraviglioso mondo del tuo prodotto.
- **Retention**: Mantenere i clienti felici è il nome del gioco. Il tuo chatbot è disponibile 24/7, pronto a risolvere problemi, offrire consigli e trucchi, e assicurarsi che i tuoi clienti si sentano apprezzati.
- **Espansione**: È tempo di fare un salto di qualità! Il tuo chatbot può suggerire sottilmente nuove funzionalità, upsell o cross-sell che si allineano con l'uso attuale del cliente del tuo prodotto. È come avere un venditore molto intelligente e non invadente sempre a disposizione.
- **Advocacy**: I clienti soddisfatti diventano i tuoi più grandi sostenitori. Il tuo chatbot può incoraggiare gli utenti soddisfatti a spargere la voce, lasciare recensioni o partecipare a programmi di referral. È come avere una macchina di entusiasmo integrata direttamente nel tuo prodotto!

## Decisione Build vs Buy

Una volta deciso di implementare un chatbot AI, la prossima grande domanda era: costruire o comprare? Come piccolo team concentrato sul nostro prodotto principale, preferiamo generalmente soluzioni "as-a-service" o piattaforme open-source popolari. Non siamo nel business di reinventare la ruota per ogni parte del nostro stack tecnologico, dopotutto. 
Così, ci siamo rimboccati le maniche e ci siamo tuffati nel mercato, cercando sia soluzioni di chatbot AI a pagamento che open-source.

I nostri requisiti erano semplici, ma non negoziabili:

- **Esperienza senza marchio**: Questo chatbot non è solo un widget carino; andrà sul nostro sito web di marketing e infine nel nostro prodotto. Non siamo interessati a pubblicizzare il marchio di qualcun altro nel nostro spazio digitale.
- **Ottima UX**: Per molti potenziali clienti, questo chatbot potrebbe essere il loro primo punto di contatto con Blue. Imposta il tono per la loro percezione della nostra azienda. Diciamolo: se non riusciamo a realizzare un chatbot adeguato sul nostro sito web, come possiamo aspettarci che i clienti si fidino di noi per i loro progetti e processi critici?
- **Costo ragionevole**: Con una grande base di utenti e piani per integrare il chatbot nel nostro prodotto principale, avevamo bisogno di una soluzione che non ci facesse spendere una fortuna man mano che l'uso cresce. Idealmente, volevamo un'opzione **BYOK (Bring Your Own Key)**. Questo ci avrebbe permesso di utilizzare la nostra chiave di servizio OpenAI o di altro tipo, pagando solo i costi variabili diretti invece di un sovrapprezzo a un fornitore terzo che non gestisce effettivamente i modelli.
- **Compatibilità con OpenAI Assistants API**: Se avessimo optato per un software open-source, non volevamo avere il fastidio di gestire un pipeline per l'ingestione dei documenti, indicizzazione, database vettoriali e tutto il resto. Volevamo utilizzare l'[OpenAI Assistants API](https://platform.openai.com/docs/assistants/overview) che avrebbe astratto tutta la complessità dietro un'API. Onestamente — è davvero ben fatto.
- **Scalabilità**: Vogliamo avere questo chatbot in più posti, con potenzialmente decine di migliaia di utenti all'anno. Ci aspettiamo un utilizzo significativo e non vogliamo essere bloccati in una soluzione che non può scalare con le nostre esigenze.

## Chatbot AI Commerciali

Quelli che abbiamo esaminato tendevano ad avere una UX migliore rispetto alle soluzioni open-source — come purtroppo spesso accade. Probabilmente ci sarà una discussione separata da fare un giorno su *perché* molte soluzioni open-source ignorano o sottovalutano l'importanza della UX.

Forniremo qui un elenco, nel caso tu stia cercando alcune offerte commerciali solide:

- **[Chatbase](https://chatbase.co):** Chatbase ti consente di costruire un chatbot AI personalizzato addestrato sulla tua base di conoscenza e aggiungerlo al tuo sito web o interagire con esso tramite la loro API. Offre funzionalità come risposte affidabili, generazione di lead, analisi avanzate e la possibilità di connettersi a più fonti di dati. Per noi, questo sembrava una delle offerte commerciali più rifinite disponibili.
- **[DocsBot AI](https://docsbot.ai/):** DocsBot AI crea bot ChatGPT personalizzati addestrati sulla tua documentazione e contenuti per supporto, pre-vendite, ricerca e altro. Fornisce widget incorporabili per aggiungere facilmente il chatbot al tuo sito web, la possibilità di rispondere automaticamente ai ticket di supporto e una potente API per integrazione.
- **[CustomGPT.ai](https://customgpt.ai):** CustomGPT.ai crea un'esperienza di chatbot personale ingerendo i dati della tua azienda, inclusi contenuti del sito web, helpdesk, basi di conoscenza, documenti e altro. Consente ai lead di porre domande e ottenere risposte immediate basate sui tuoi contenuti, senza bisogno di cercare. Interessante, affermano anche di [battere OpenAI in RAG (Retrieval Augmented Generation)!](https://customgpt.ai/customgpt-beats-open-ai-in-rag-benchmark/)
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Questa è un'offerta commerciale interessante, perché *è* anche software open-source. Sembra un po' in fase iniziale e i prezzi non sembravano realistici ($27/mese per messaggi illimitati non funzioneranno mai commercialmente per loro).

Abbiamo anche esaminato [InterCom Fin](https://www.intercom.com/fin) che fa parte del loro software di supporto clienti. Questo avrebbe significato passare da [HelpScout](https://wwww.helpscout.com) che utilizziamo da quando abbiamo iniziato Blue. Questo sarebbe stato possibile, ma InterCom Fin ha dei prezzi esorbitanti che semplicemente lo escludono dalla considerazione.

E questo è effettivamente il problema con molte delle offerte commerciali. InterCom Fin addebita $0.99 per ogni richiesta di supporto gestita, e ChatBase addebita $399/mese per 40.000 messaggi. Sono quasi $5k all'anno per un semplice widget di chat.

Considerando che i prezzi per l'inferenza AI stanno scendendo in modo vertiginoso. OpenAI ha ridotto i suoi prezzi in modo piuttosto drammatico:

- Il GPT-4 originale (8k contesto) era prezzato a $0.03 per 1K token di prompt.
- Il GPT-4 Turbo (128k contesto) era prezzato a $0.01 per 1K token di prompt, una riduzione del 50% rispetto al GPT-4 originale.
- Il modello GPT-4o è prezzato a $0.005 per 1K token, che è un'ulteriore riduzione del 50% rispetto al prezzo del GPT-4 Turbo.

Si tratta di una riduzione dei costi dell'83%, e non ci aspettiamo che rimanga stagnante.

Considerando che stavamo cercando una soluzione scalabile che sarebbe stata utilizzata da decine di migliaia di utenti all'anno con un numero significativo di messaggi, ha senso andare direttamente alla fonte e pagare i costi API direttamente, non utilizzare una versione commerciale che aumenta i costi.

## Chatbot AI Open Source

Come accennato, le opzioni open source che abbiamo esaminato erano per lo più deludenti riguardo al requisito "Ottima UX".

Abbiamo esaminato:

- **[Deepchat](https://deepchat.dev/)**: Questo è un componente di chat indipendente dal framework per servizi AI, che si connette a varie API AI, inclusa OpenAI. Ha anche la possibilità per gli utenti di scaricare un modello AI che funziona direttamente nel browser. Abbiamo provato a lavorarci e abbiamo ottenuto una versione funzionante, ma l'API OpenAI Assistants implementata sembrava piuttosto buggata con diversi problemi. Tuttavia, questo è un progetto molto promettente e il loro playground è davvero ben fatto.
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Guardando di nuovo a questo da una prospettiva open-source, questo richiederebbe di creare un bel po' di infrastruttura, qualcosa che non volevamo fare, perché volevamo fare affidamento il più possibile sull'API Assistants di OpenAI.

## Costruire il nostro ChatBot

E così, senza riuscire a trovare qualcosa che soddisfacesse tutti i nostri requisiti, abbiamo deciso di costruire il nostro chatbot AI che potesse interfacciarsi con l'API Assistants di OpenAI. Questo, alla fine, si è rivelato relativamente indolore!

Il nostro sito web utilizza [Nuxt3](https://nuxt.com), [Vue3](https://vuejs.org/) (che è lo stesso framework della piattaforma Blue) e [TailwindUI](https://tailwindui.com/).

Il primo passo è stato creare un'API (Application Programming Interface) in Nuxt3 che possa "parlare" con l'API Assistants di OpenAI. Questo era necessario poiché non volevamo fare tutto sul front-end, poiché questo avrebbe esposto la nostra chiave API OpenAI al mondo, con il potenziale di abuso.

La nostra API backend funge da intermediario sicuro tra il browser dell'utente e OpenAI. Ecco cosa fa:

- **Gestione delle conversazioni:** Crea e gestisce "thread" per ogni conversazione. Pensa a un thread come a una sessione di chat unica che ricorda tutto ciò che hai detto.
- **Gestione dei messaggi:** Quando invii un messaggio, la nostra API lo aggiunge al thread giusto e chiede all'assistente di OpenAI di elaborare una risposta.
- **Attesa intelligente:** Invece di farti fissare uno schermo di caricamento, la nostra API controlla OpenAI ogni secondo per vedere se la tua risposta è pronta. È come avere un cameriere che tiene d'occhio il tuo ordine senza disturbare lo chef ogni due secondi.
- **Sicurezza prima di tutto:** Gestendo tutto questo sul server, manteniamo i tuoi dati e le nostre chiavi API al sicuro.

Poi, c'era il front-end e l'esperienza utente. Come discusso in precedenza, questo era *criticamente* importante, perché non abbiamo una seconda possibilità di fare una prima impressione!

Nel progettare il nostro chatbot, abbiamo prestato particolare attenzione all'esperienza utente, assicurandoci che ogni interazione fosse fluida, intuitiva e riflettesse l'impegno di Blue per la qualità. L'interfaccia del chatbot inizia con un semplice ed elegante cerchio blu, utilizzando [HeroIcons per le nostre icone](https://heroicons.com/) (che utilizziamo in tutto il sito web di Blue) come widget di apertura del nostro chatbot. Questa scelta di design garantisce coerenza visiva e immediata riconoscibilità del marchio.

![](/insights/ai-chatbot-circle.png)

Comprendiamo che a volte gli utenti potrebbero aver bisogno di ulteriore supporto o informazioni più dettagliate. Ecco perché abbiamo incluso collegamenti convenienti all'interno dell'interfaccia del chatbot. Un collegamento email per il supporto è prontamente disponibile, consentendo agli utenti di contattare direttamente il nostro team se necessitano di assistenza più personalizzata. Inoltre, abbiamo incorporato un collegamento alla documentazione, fornendo un facile accesso a risorse più complete per coloro che vogliono approfondire le offerte di Blue.

L'esperienza utente è ulteriormente migliorata da eleganti animazioni di fade-in e fade-up quando si apre la finestra del chatbot. Queste sottili animazioni aggiungono un tocco di sofisticazione all'interfaccia, rendendo l'interazione più dinamica e coinvolgente. Abbiamo anche implementato un indicatore di digitazione, una piccola ma cruciale funzionalità che informa gli utenti che il chatbot sta elaborando la loro richiesta e creando una risposta. Questo segnale visivo aiuta a gestire le aspettative degli utenti e mantiene un senso di comunicazione attiva.

<video autoplay loop muted playsinline>
  <source src="/videos/ai-chatbot-animation.mp4" type="video/mp4">
</video>

Riconoscendo che alcune conversazioni potrebbero richiedere più spazio sullo schermo, abbiamo aggiunto la possibilità di aprire la conversazione in una finestra più grande. Questa funzionalità è particolarmente utile per scambi più lunghi o quando si rivedono informazioni dettagliate, dando agli utenti la flessibilità di adattare il chatbot alle loro esigenze.

Dietro le quinte, abbiamo implementato un'elaborazione intelligente per ottimizzare le risposte del chatbot. Il nostro sistema analizza automaticamente le risposte dell'AI per rimuovere riferimenti ai nostri documenti interni, assicurando che le informazioni presentate siano pulite, pertinenti e focalizzate esclusivamente sull'affrontare la richiesta dell'utente. 
Per migliorare la leggibilità e consentire una comunicazione più sfumata, abbiamo incorporato il supporto markdown utilizzando la libreria 'marked'. Questa funzionalità consente alla nostra AI di fornire testo riccamente formattato, inclusi enfasi in grassetto e corsivo, elenchi strutturati e persino frammenti di codice quando necessario. È come ricevere un mini-documento ben formattato e personalizzato in risposta alle tue domande.

Ultimo ma non meno importante, abbiamo dato priorità alla sicurezza nella nostra implementazione. Utilizzando la libreria DOMPurify, sanitizziamo l'HTML generato dal parsing markdown. Questo passaggio cruciale garantisce che eventuali script o codici potenzialmente dannosi vengano rimossi prima che il contenuto venga visualizzato per te. È il nostro modo di garantire che le informazioni utili che ricevi siano non solo informative, ma anche sicure da consumare.

## Sviluppi Futuri

Quindi questo è solo l'inizio, abbiamo alcune cose entusiasmanti in programma per questa funzionalità.

Una delle nostre prossime funzionalità è la possibilità di trasmettere risposte in tempo reale. Presto vedrai le risposte del chatbot apparire carattere per carattere, rendendo le conversazioni più naturali e dinamiche. È come vedere l'AI pensare, creando un'esperienza più coinvolgente e interattiva che ti tiene informato ad ogni passo.

Per i nostri preziosi utenti di Blue, stiamo lavorando sulla personalizzazione. Il chatbot riconoscerà quando sei connesso, adattando le sue risposte in base alle informazioni del tuo account, alla cronologia degli utilizzi e alle preferenze. Immagina un chatbot che non solo risponde alle tue domande, ma comprende il tuo contesto specifico all'interno dell'ecosistema Blue, fornendo assistenza più pertinente e personalizzata.

Comprendiamo che potresti lavorare su più progetti o avere varie domande. Ecco perché stiamo sviluppando la possibilità di mantenere diversi thread di conversazione distinti con il nostro chatbot. Questa funzionalità ti permetterà di passare senza problemi tra diversi argomenti, senza perdere il contesto – proprio come avere più schede aperte nel tuo browser.

Per rendere le tue interazioni ancora più produttive, stiamo creando una funzionalità che offrirà domande di follow-up suggerite in base alla tua conversazione attuale. Questo ti aiuterà a esplorare argomenti più in profondità e scoprire informazioni correlate che potresti non aver pensato di chiedere, rendendo ogni sessione di chat più completa e preziosa.

Siamo anche entusiasti di creare una suite di assistenti AI specializzati, ciascuno progettato per esigenze specifiche. Che tu stia cercando di rispondere a domande pre-vendita, impostare un nuovo progetto o risolvere funzionalità avanzate, potrai scegliere l'assistente che meglio si adatta alle tue esigenze attuali. È come avere un team di esperti di Blue a portata di mano, ciascuno specializzato in diversi aspetti della nostra piattaforma.

Infine, stiamo lavorando per consentirti di caricare screenshot direttamente nella chat. L'AI analizzerà l'immagine e fornirà spiegazioni o passaggi di risoluzione dei problemi in base a ciò che vede. Questa funzionalità renderà più facile che mai ricevere aiuto con problemi specifici che incontri mentre utilizzi Blue, colmando il divario tra informazioni visive e assistenza testuale.

## Conclusione

Speriamo che questo approfondimento sul nostro processo di sviluppo del chatbot AI abbia fornito alcune intuizioni preziose sul nostro modo di pensare allo sviluppo del prodotto in Blue. Il nostro viaggio dall'identificazione della necessità di un chatbot alla costruzione della nostra soluzione mostra come affrontiamo il processo decisionale e l'innovazione.

![](/insights/ai-chatbot-modal.png)

In Blue, valutiamo attentamente le opzioni di costruire rispetto a comprare, sempre con un occhio a ciò che servirà meglio i nostri utenti e si allinea con i nostri obiettivi a lungo termine. In questo caso, abbiamo identificato un significativo gap nel mercato per un chatbot economico ma visivamente accattivante che potesse soddisfare le nostre esigenze specifiche. Anche se generalmente sosteniamo di sfruttare soluzioni esistenti piuttosto che reinventare la ruota, a volte il miglior percorso da seguire è creare qualcosa su misura per i tuoi requisiti unici.

La nostra decisione di costruire il nostro chatbot non è stata presa alla leggera. È stata il risultato di una ricerca di mercato approfondita, di una chiara comprensione delle nostre esigenze e di un impegno a fornire la migliore esperienza possibile per i nostri utenti. Sviluppando internamente, siamo stati in grado di creare una soluzione che non solo soddisfa le nostre esigenze attuali, ma getta anche le basi per futuri miglioramenti e integrazioni.

Questo progetto esemplifica il nostro approccio in Blue: non abbiamo paura di rimboccarci le maniche e costruire qualcosa da zero quando è la scelta giusta per il nostro prodotto e i nostri utenti. È questa disponibilità a fare un passo in più che ci consente di offrire soluzioni innovative che soddisfano veramente le esigenze dei nostri clienti.
Siamo entusiasti del futuro del nostro chatbot AI e del valore che porterà sia agli utenti potenziali che a quelli esistenti di Blue. Mentre continuiamo a perfezionare ed espandere le sue capacità, rimaniamo impegnati a spingere i confini di ciò che è possibile nella gestione dei progetti e nell'interazione con i clienti.

Grazie per averci accompagnato in questo viaggio attraverso il nostro processo di sviluppo. Speriamo che ti abbia dato un'idea dell'approccio attento e incentrato sull'utente che adottiamo in ogni aspetto di Blue. Rimanete sintonizzati per ulteriori aggiornamenti mentre continuiamo a evolvere e migliorare la nostra piattaforma per servirti meglio.

Se sei interessato, puoi trovare il link al codice sorgente per questo progetto qui:

- **[ChatWidget](https://gitlab.com/bloohq/blue-website/-/blob/main/components/ChatWidget.vue)**: Questo è un componente Vue che alimenta il widget di chat stesso.
- **[Chat API](https://gitlab.com/bloohq/blue-website/-/blob/main/server/api/chat.post.ts)**: Questo è il middleware che funziona tra il componente di chat e l'API Assistants di OpenAI.
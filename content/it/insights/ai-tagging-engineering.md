---
title: Categorizzazione Automatica con AI (Approfondimento Tecnico)
category: "Engineering"
description: Scoprite il dietro le quinte con il team di ingegneria di Blue mentre spiegano come hanno costruito una funzionalità di categorizzazione e tagging automatico basata sull'AI.
date: 2024-12-07
---

Abbiamo recentemente rilasciato la [Categorizzazione Automatica con AI](/insights/ai-auto-categorization) a tutti gli utenti di Blue. Si tratta di una funzionalità AI inclusa nell'abbonamento principale di Blue, senza costi aggiuntivi. In questo post, approfondiamo l'ingegneria dietro la realizzazione di questa funzionalità.

---
In Blue, il nostro approccio allo sviluppo delle funzionalità si basa su una profonda comprensione delle esigenze degli utenti e delle tendenze di mercato, unita all'impegno di mantenere la semplicità e la facilità d'uso che definiscono la nostra piattaforma. Questo è ciò che guida la nostra [roadmap](/platform/roadmap), e ciò che ci ha [permesso di rilasciare costantemente nuove funzionalità ogni mese per anni](/platform/changelog).

L'introduzione del tagging automatico basato su AI in Blue è un perfetto esempio di questa filosofia in azione. Prima di immergerci nei dettagli tecnici di come abbiamo costruito questa funzionalità, è fondamentale comprendere il problema che stavamo risolvendo e l'attenta considerazione che è stata dedicata al suo sviluppo.

Il panorama del project management si sta evolvendo rapidamente, con le capacità AI che stanno diventando sempre più centrali nelle aspettative degli utenti. I nostri clienti, in particolare quelli che gestiscono [progetti](/platform) su larga scala con milioni di [record](/platform/features/records), erano stati espliciti nel loro desiderio di modi più intelligenti ed efficienti per organizzare e categorizzare i loro dati.

Tuttavia, in Blue, non aggiungiamo semplicemente funzionalità perché sono di tendenza o richieste. La nostra filosofia è che ogni nuova aggiunta deve dimostrare il suo valore, con la risposta predefinita che è un deciso *"no"* fino a quando una funzionalità non dimostra una forte domanda e una chiara utilità.

Per comprendere veramente la profondità del problema e il potenziale del tagging automatico con AI, abbiamo condotto ampie interviste con i clienti, concentrandoci su utenti di lunga data che gestiscono progetti complessi e ricchi di dati in diversi domini.

Queste conversazioni hanno rivelato un filo conduttore comune: *mentre il tagging era prezioso per l'organizzazione e la ricercabilità, la natura manuale del processo stava diventando un collo di bottiglia, specialmente per i team che gestiscono alti volumi di record.*

Ma abbiamo visto oltre la semplice risoluzione del problema immediato del tagging manuale.

Abbiamo immaginato un futuro in cui il tagging basato su AI potrebbe diventare la base per flussi di lavoro più intelligenti e automatizzati.

Il vero potere di questa funzionalità, ci siamo resi conto, risiedeva nel suo potenziale di integrazione con il nostro [sistema di automazione del project management](/platform/features/automations). Immaginate uno strumento di project management che non solo categorizza le informazioni in modo intelligente, ma utilizza anche quelle categorie per instradare i task, attivare azioni e adattare i flussi di lavoro in tempo reale.

Questa visione si allineava perfettamente con il nostro obiettivo di mantenere Blue semplice ma potente.

Inoltre, abbiamo riconosciuto il potenziale per estendere questa capacità oltre i confini della nostra piattaforma. Sviluppando un robusto sistema di tagging AI, stavamo gettando le basi per una "API di categorizzazione" che potrebbe funzionare immediatamente, aprendo potenzialmente nuove strade per come i nostri utenti interagiscono con Blue e lo sfruttano nei loro ecosistemi tecnologici più ampi.

Questa funzionalità, quindi, non riguardava solo l'aggiunta di una casella di controllo AI alla nostra lista di funzionalità.

Si trattava di fare un passo significativo verso una piattaforma di project management più intelligente e adattiva, rimanendo fedeli alla nostra filosofia fondamentale di semplicità e centralità dell'utente.

Nelle sezioni seguenti, approfondiremo le sfide tecniche che abbiamo affrontato nel dare vita a questa visione, l'architettura che abbiamo progettato per supportarla e le soluzioni che abbiamo implementato. Esploreremo anche le possibilità future che questa funzionalità apre, mostrando come un'aggiunta attentamente considerata possa aprire la strada a cambiamenti trasformativi nel project management.

---
## Il Problema

Come discusso sopra, il tagging manuale dei record di progetto può essere dispendioso in termini di tempo e incoerente.

Ci siamo proposti di risolvere questo problema sfruttando l'AI per suggerire automaticamente tag basati sul contenuto dei record.

Le sfide principali erano:

1. Scegliere un modello AI appropriato
2. Elaborare efficacemente grandi volumi di record
3. Garantire privacy e sicurezza dei dati
4. Integrare la funzionalità senza problemi nella nostra architettura esistente

## Selezione del Modello AI

Abbiamo valutato diverse piattaforme AI, tra cui [OpenAI](https://openai.com), modelli open-source su [HuggingFace](https://huggingface.co/) e [Replicate](https://replicate.com).

I nostri criteri includevano:

- Rapporto costo-efficacia
- Accuratezza nella comprensione del contesto
- Capacità di aderire a formati di output specifici
- Garanzie di privacy dei dati

Dopo test approfonditi, abbiamo scelto [GPT-3.5 Turbo](https://platform.openai.com/docs/models/gpt-3-5-turbo) di OpenAI. Mentre [GPT-4](https://softgist.com/the-ultimate-guide-to-prompt-engineering) potrebbe offrire miglioramenti marginali nell'accuratezza, i nostri test hanno mostrato che le prestazioni di GPT-3.5 erano più che adeguate per le nostre esigenze di auto-tagging. L'equilibrio tra costo-efficacia e forti capacità di categorizzazione ha reso GPT-3.5 la scelta ideale per questa funzionalità.

Il costo più elevato di GPT-4 ci avrebbe costretto a offrire la funzionalità come componente aggiuntivo a pagamento, in conflitto con il nostro obiettivo di **includere l'AI nel nostro prodotto principale senza costi aggiuntivi per gli utenti finali.**

Al momento della nostra implementazione, i prezzi per GPT-3.5 Turbo sono:

- $0.0005 per 1K token di input (o $0.50 per 1M token di input)
- $0.0015 per 1K token di output (o $1.50 per 1M token di output)

Facciamo alcune ipotesi su un record medio in Blue:

- **Titolo**: ~10 token
- **Descrizione**: ~50 token
- **2 commenti**: ~30 token ciascuno
- **5 campi personalizzati**: ~10 token ciascuno
- **Nome lista, data di scadenza e altri metadati**: ~20 token
- **Prompt di sistema e tag disponibili**: ~50 token

Token di input totali per record: 10 + 50 + (30 * 2) + (10 * 5) + 20 + 50 ≈ 240 token

Per l'output, supponiamo una media di 3 tag suggeriti per record, che potrebbero totalizzare circa 20 token di output inclusa la formattazione JSON.

Per 1 milione di record:

- Costo input: (240 * 1.000.000 / 1.000.000) * $0.50 = $120
- Costo output: (20 * 1.000.000 / 1.000.000) * $1.50 = $30

**Costo totale per l'auto-tagging di 1 milione di record: $120 + $30 = $150**

## Prestazioni di GPT3.5 Turbo

La categorizzazione è un compito in cui i modelli linguistici di grandi dimensioni (LLM) come GPT-3.5 Turbo eccellono, rendendoli particolarmente adatti per la nostra funzionalità di auto-tagging. Gli LLM sono addestrati su vaste quantità di dati testuali, permettendo loro di comprendere contesto, semantica e relazioni tra concetti. Questa ampia base di conoscenze consente loro di eseguire compiti di categorizzazione con alta precisione in un'ampia gamma di domini.

Per il nostro caso d'uso specifico di tagging per project management, GPT-3.5 Turbo dimostra diversi punti di forza chiave:

- **Comprensione Contestuale:** Può cogliere il contesto generale di un record di progetto, considerando non solo le singole parole ma il significato trasmesso dall'intera descrizione, commenti e altri campi.
- **Flessibilità:** Può adattarsi a vari tipi di progetti e settori senza richiedere estesa riprogrammazione.
- **Gestione dell'Ambiguità:** Può valutare molteplici fattori per prendere decisioni sfumate.
- **Apprendimento dagli Esempi:** Può comprendere e applicare rapidamente nuovi schemi di categorizzazione senza addestramento aggiuntivo.
- **Classificazione Multi-etichetta:** Può suggerire più tag rilevanti per un singolo record, il che era cruciale per i nostri requisiti.

GPT-3.5 Turbo si è distinto anche per la sua affidabilità nell'aderire al nostro formato di output JSON richiesto, che era *cruciale* per l'integrazione perfetta con i nostri sistemi esistenti. I modelli open-source, sebbene promettenti, spesso aggiungevano commenti extra o deviavano dal formato previsto, il che avrebbe richiesto ulteriore post-processing. Questa coerenza nel formato di output è stata un fattore chiave nella nostra decisione, poiché ha semplificato notevolmente la nostra implementazione e ridotto i potenziali punti di errore.

Optare per GPT-3.5 Turbo con il suo output JSON coerente ci ha permesso di implementare una soluzione più diretta, affidabile e manutenibile.

Se avessimo scelto un modello con formattazione meno affidabile, avremmo affrontato una cascata di complicazioni: la necessità di una logica di parsing robusta per gestire vari formati di output, ampia gestione degli errori per output incoerenti, potenziali impatti sulle prestazioni dovuti all'elaborazione aggiuntiva, maggiore complessità dei test per coprire tutte le variazioni di output e un maggiore onere di manutenzione a lungo termine.

Gli errori di parsing potrebbero portare a tagging errati, impattando negativamente sull'esperienza utente. Evitando queste insidie, siamo stati in grado di concentrare i nostri sforzi ingegneristici su aspetti critici come l'ottimizzazione delle prestazioni e il design dell'interfaccia utente, piuttosto che combattere con output AI imprevedibili.

## Architettura del Sistema

La nostra funzionalità di auto-tagging AI è costruita su un'architettura robusta e scalabile progettata per gestire alti volumi di richieste in modo efficiente fornendo un'esperienza utente fluida. Come per tutti i nostri sistemi, abbiamo progettato questa funzionalità per supportare un ordine di grandezza in più di traffico rispetto a quello che attualmente sperimentiamo. Questo approccio, sebbene apparentemente sovradimensionato per le esigenze attuali, è una best practice che ci permette di gestire senza problemi picchi improvvisi di utilizzo e ci dà ampio margine per la crescita senza importanti revisioni architetturali. Altrimenti, dovremmo riprogettare tutti i nostri sistemi ogni 18 mesi — qualcosa che abbiamo imparato a nostre spese in passato!

Analizziamo i componenti e il flusso del nostro sistema:

- **Interazione Utente:** Il processo inizia quando un utente preme il pulsante "Autotag" nell'interfaccia di Blue. Questa azione attiva il flusso di lavoro di auto-tagging.
- **Chiamata API Blue:** L'azione dell'utente viene tradotta in una chiamata API al nostro backend Blue. Questo endpoint API è progettato per gestire le richieste di auto-tagging.
- **Gestione della Coda:** Invece di elaborare immediatamente la richiesta, il che potrebbe portare a problemi di prestazioni sotto carico elevato, aggiungiamo la richiesta di tagging a una coda. Utilizziamo Redis per questo meccanismo di accodamento, che ci permette di gestire efficacemente il carico e garantire la scalabilità del sistema.
- **Servizio in Background:** Abbiamo implementato un servizio in background che monitora continuamente la coda per nuove richieste. Questo servizio è responsabile dell'elaborazione delle richieste in coda.
- **Integrazione API OpenAI:** Il servizio in background prepara i dati necessari e effettua chiamate API al modello GPT-3.5 di OpenAI. È qui che avviene il tagging effettivo basato su AI. Inviamo dati di progetto rilevanti e riceviamo in cambio tag suggeriti.
- **Elaborazione dei Risultati:** Il servizio in background elabora i risultati ricevuti da OpenAI. Questo comporta il parsing della risposta dell'AI e la preparazione dei dati per l'applicazione al progetto.
- **Applicazione dei Tag:** I risultati elaborati vengono utilizzati per applicare i nuovi tag agli elementi rilevanti nel progetto. Questo passaggio aggiorna il nostro database con i tag suggeriti dall'AI.
- **Riflessione nell'Interfaccia Utente:** Infine, i nuovi tag appaiono nella vista del progetto dell'utente, completando il processo di auto-tagging dal punto di vista dell'utente.

Questa architettura offre diversi vantaggi chiave che migliorano sia le prestazioni del sistema che l'esperienza utente. Utilizzando una coda e l'elaborazione in background, abbiamo raggiunto una scalabilità impressionante, permettendoci di gestire numerose richieste simultaneamente senza sovraccaricare il nostro sistema o raggiungere i limiti di frequenza dell'API di OpenAI. L'implementazione di questa architettura ha richiesto un'attenta considerazione di vari fattori per garantire prestazioni e affidabilità ottimali. Per la gestione delle code, abbiamo scelto Redis, sfruttando la sua velocità e affidabilità nella gestione di code distribuite.

Questo approccio contribuisce anche alla reattività complessiva della funzionalità. Gli utenti ricevono un feedback immediato che la loro richiesta è in elaborazione, anche se il tagging effettivo richiede del tempo, creando una sensazione di interazione in tempo reale. La tolleranza ai guasti dell'architettura è un altro vantaggio cruciale. Se qualsiasi parte del processo incontra problemi, come interruzioni temporanee dell'API di OpenAI, possiamo riprovare con grazia o gestire l'errore senza impattare l'intero sistema.

Questa robustezza, combinata con l'apparizione in tempo reale dei tag, migliora l'esperienza utente, dando l'impressione della "magia" dell'AI al lavoro.

## Dati e Prompt

Un passaggio cruciale nel nostro processo di auto-tagging è la preparazione dei dati da inviare al modello GPT-3.5. Questo passaggio ha richiesto un'attenta considerazione per bilanciare la fornitura di contesto sufficiente per un tagging accurato mantenendo l'efficienza e proteggendo la privacy degli utenti. Ecco uno sguardo dettagliato al nostro processo di preparazione dei dati.

Per ogni record, compiliamo le seguenti informazioni:

- **Nome Lista**: Fornisce contesto sulla categoria più ampia o fase del progetto.
- **Titolo Record**: Spesso contiene informazioni chiave sullo scopo o contenuto del record.
- **Campi Personalizzati**: Includiamo [campi personalizzati](/platform/features/custom-fields) basati su testo e numeri, che spesso contengono informazioni cruciali specifiche del progetto.
- **Descrizione**: Tipicamente contiene le informazioni più dettagliate sul record.
- **Commenti**: Possono fornire contesto aggiuntivo o aggiornamenti che potrebbero essere rilevanti per il tagging.
- **Data di Scadenza**: Informazioni temporali che potrebbero influenzare la selezione dei tag.

È interessante notare che non inviamo dati di tag esistenti a GPT-3.5, e lo facciamo per evitare di influenzare il modello.

Il cuore della nostra funzionalità di auto-tagging risiede nel modo in cui interagiamo con il modello GPT-3.5 e processiamo le sue risposte. Questa sezione della nostra pipeline ha richiesto un design attento per garantire un tagging accurato, coerente ed efficiente.

Utilizziamo un prompt di sistema attentamente elaborato per istruire l'AI sul suo compito. Ecco una scomposizione del nostro prompt e la logica dietro ogni componente:

```
You will be provided with record data, and your task is to choose the tags that are relevant to the record.
You can respond with an empty array if you are unsure.
Available tags: ${tags}.
Today: ${currentDate}.
Please respond in JSON using the following format:
{ "tags": ["tag-1", "tag-2"] }
```

- **Definizione del Compito:** Dichiariamo chiaramente il compito dell'AI per garantire risposte focalizzate.
- **Gestione dell'Incertezza:** Permettiamo esplicitamente risposte vuote, prevenendo il tagging forzato quando l'AI è incerta.
- **Tag Disponibili:** Forniamo una lista di tag validi (${tags}) per vincolare le scelte dell'AI ai tag di progetto esistenti.
- **Data Corrente:** Includere ${currentDate} aiuta l'AI a comprendere il contesto temporale, che può essere cruciale per certi tipi di progetti.
- **Formato di Risposta:** Specifichiamo un formato JSON per un parsing facile e controllo degli errori.

Questo prompt è il risultato di test e iterazioni estensive. Abbiamo scoperto che essere espliciti sul compito, le opzioni disponibili e il formato di output desiderato ha migliorato significativamente l'accuratezza e la coerenza delle risposte dell'AI — la semplicità è la chiave!

La lista dei tag disponibili viene generata lato server e validata prima dell'inclusione nel prompt. Implementiamo limiti rigorosi di caratteri sui nomi dei tag per prevenire prompt sovradimensionati.

Come menzionato sopra, non abbiamo avuto problemi con GPT-3.5 Turbo nel ricevere la risposta JSON pura nel formato corretto il 100% delle volte.

Quindi in sintesi,

- Combiniamo il prompt di sistema con i dati del record preparati.
- Questo prompt combinato viene inviato al modello GPT-3.5 tramite l'API di OpenAI.
- Utilizziamo un'impostazione di temperatura di 0.3 per bilanciare creatività e coerenza nelle risposte dell'AI.
- La nostra chiamata API include un parametro max_tokens per limitare la dimensione della risposta e controllare i costi.

Una volta ricevuta la risposta dell'AI, passiamo attraverso diversi passaggi per elaborare e applicare i tag suggeriti:

* **Parsing JSON**: Tentiamo di fare il parsing della risposta come JSON. Se il parsing fallisce, registriamo l'errore e saltiamo il tagging per quel record.
* **Validazione dello Schema**: Validiamo il JSON parsato contro il nostro schema previsto (un oggetto con un array "tags"). Questo cattura qualsiasi deviazione strutturale nella risposta dell'AI.
* **Validazione dei Tag**: Incrociamo i tag suggeriti con la nostra lista di tag di progetto validi. Questo passaggio filtra qualsiasi tag che non esiste nel progetto, il che potrebbe verificarsi se l'AI ha frainteso o se i tag del progetto sono cambiati tra la creazione del prompt e l'elaborazione della risposta.
* **Deduplicazione**: Rimuoviamo qualsiasi tag duplicato dal suggerimento dell'AI per evitare tagging ridondanti.
* **Applicazione**: I tag validati e deduplicati vengono quindi applicati al record nel nostro database.
* **Logging e Analytics**: Registriamo i tag finali applicati. Questi dati sono preziosi per monitorare le prestazioni del sistema e migliorarlo nel tempo.

## Sfide

L'implementazione dell'auto-tagging basato su AI in Blue ha presentato diverse sfide uniche, ognuna richiedente soluzioni innovative per garantire una funzionalità robusta, efficiente e user-friendly.

### Annullare Operazioni in Blocco

La funzionalità di AI Tagging può essere eseguita sia su singoli record che in blocco. Il problema con l'operazione in blocco è che se all'utente non piace il risultato, dovrebbe passare manualmente attraverso migliaia di record e annullare il lavoro dell'AI. Chiaramente, questo è inaccettabile.

Per risolvere questo, abbiamo implementato un innovativo sistema di sessioni di tagging. Ad ogni operazione di tagging in blocco viene assegnato un ID sessione univoco, che è associato a tutti i tag applicati durante quella sessione. Questo ci permette di gestire efficientemente le operazioni di annullamento semplicemente eliminando tutti i tag associati a un particolare ID sessione. Rimuoviamo anche le tracce di audit correlate, assicurando che le operazioni annullate non lascino traccia nel sistema. Questo approccio dà agli utenti la fiducia di sperimentare con il tagging AI, sapendo che possono facilmente annullare le modifiche se necessario.

### Privacy dei Dati

La privacy dei dati era un'altra sfida critica che abbiamo affrontato.

I nostri utenti ci affidano i loro dati di progetto, ed era fondamentale garantire che queste informazioni non fossero conservate o utilizzate per l'addestramento del modello da OpenAI. Abbiamo affrontato questo su più fronti.

In primo luogo, abbiamo formato un accordo con OpenAI che proibisce esplicitamente l'uso dei nostri dati per l'addestramento del modello. Inoltre, OpenAI elimina i dati dopo l'elaborazione, fornendo un ulteriore livello di protezione della privacy.

Da parte nostra, abbiamo preso la precauzione di escludere informazioni sensibili, come i dettagli degli assegnatari, dai dati inviati all'AI, quindi questo garantisce che i nomi specifici degli individui non vengano inviati a terze parti insieme ad altri dati.

Questo approccio completo ci permette di sfruttare le capacità dell'AI mantenendo i più alti standard di privacy e sicurezza dei dati.

### Limiti di Frequenza e Gestione degli Errori

Una delle nostre preoccupazioni principali era la scalabilità e i limiti di frequenza. Le chiamate API dirette a OpenAI per ogni record sarebbero state inefficienti e avrebbero potuto rapidamente raggiungere i limiti di frequenza, specialmente per progetti grandi o durante i periodi di picco di utilizzo. Per affrontare questo, abbiamo sviluppato un'architettura di servizi in background che ci permette di raggruppare le richieste e implementare il nostro sistema di accodamento. Questo approccio ci aiuta a gestire la frequenza delle chiamate API e consente un'elaborazione più efficiente di grandi volumi di record, garantendo prestazioni fluide anche sotto carico pesante.

La natura delle interazioni AI significava che dovevamo anche prepararci per errori occasionali o output inaspettati. C'erano casi in cui l'AI poteva produrre JSON non valido o output che non corrispondevano al nostro formato previsto. Per gestire questo, abbiamo implementato una robusta gestione degli errori e logica di parsing in tutto il nostro sistema. Se la risposta dell'AI non è JSON valido o non contiene la chiave "tags" prevista, il nostro sistema è progettato per trattarla come se nessun tag fosse stato suggerito, piuttosto che tentare di elaborare dati potenzialmente corrotti. Questo garantisce che anche di fronte all'imprevedibilità dell'AI, il nostro sistema rimanga stabile e affidabile.

## Sviluppi Futuri

Crediamo che le funzionalità, e il prodotto Blue nel suo complesso, non siano mai "finiti" — c'è sempre spazio per miglioramenti.

Ci sono state alcune funzionalità che abbiamo considerato nella build iniziale che non hanno superato la fase di scoping, ma è interessante notarle poiché probabilmente implementeremo qualche versione di esse in futuro.

La prima è l'aggiunta della descrizione dei tag. Questo permetterebbe agli utenti finali di dare ai tag non solo un nome e un colore, ma anche una descrizione opzionale. Questa verrebbe anche passata all'AI per aiutare a fornire ulteriore contesto e potenzialmente migliorare l'accuratezza.

Mentre il contesto aggiuntivo potrebbe essere prezioso, siamo consapevoli della potenziale complessità che potrebbe introdurre. C'è un delicato equilibrio da trovare tra fornire informazioni utili e sopraffare gli utenti con troppi dettagli. Mentre sviluppiamo questa funzionalità, ci concentreremo sul trovare quel punto ottimale dove il contesto aggiunto migliora piuttosto che complicare l'esperienza utente.

Forse il miglioramento più eccitante all'orizzonte è l'integrazione dell'auto-tagging AI con il nostro [sistema di automazione del project management](/platform/features/automations).

Questo significa che la funzionalità di tagging AI potrebbe essere sia un trigger che un'azione da un'automazione. Questo potrebbe essere enorme poiché potrebbe trasformare questa funzionalità di categorizzazione AI relativamente semplice in un sistema di instradamento del lavoro basato su AI.

Immaginate un'automazione che afferma:

Quando l'AI tagga un record come "Critico" -> Assegna a "Manager" e Invia un'Email Personalizzata

Questo significa che quando fate l'AI-tag di un record, se l'AI decide che è un problema critico, allora può automaticamente assegnare il project manager e inviargli un'email personalizzata. Questo estende i [benefici del nostro sistema di automazione del project management](/platform/features/automations) da un sistema puramente basato su regole a un vero sistema AI flessibile.

Esplorando continuamente le frontiere dell'AI nel project management, miriamo a fornire ai nostri utenti strumenti che non solo soddisfano le loro esigenze attuali ma anticipano e plasmano il futuro del lavoro. Come sempre, svilupperemo queste funzionalità in stretta collaborazione con la nostra comunità di utenti, assicurando che ogni miglioramento aggiunga valore reale e pratico al processo di project management.

## Conclusione

Ecco fatto!

Questa è stata una funzionalità divertente da implementare, e uno dei nostri primi passi nell'AI, insieme al [Riepilogo dei Contenuti con AI](/insights/ai-content-summarization) che abbiamo precedentemente lanciato. Sappiamo che l'AI avrà un ruolo sempre più grande nel project management in futuro, e non vediamo l'ora di lanciare più funzionalità innovative sfruttando LLM (Large Language Models) avanzati.

C'è stato parecchio da pensare durante l'implementazione, e siamo particolarmente entusiasti di come possiamo sfruttare questa funzionalità in futuro con il [motore di automazione del project management](/insights/benefits-project-management-automation) esistente di Blue.

Speriamo anche che sia stata una lettura interessante, e che vi dia uno sguardo su come pensiamo all'ingegnerizzazione delle funzionalità che usate ogni giorno.

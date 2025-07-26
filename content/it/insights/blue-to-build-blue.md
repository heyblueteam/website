---
title:  Come utilizziamo Blue per costruire Blue. 
description: Scopri come utilizziamo la nostra piattaforma di gestione progetti per costruire la nostra piattaforma di gestione progetti!
category: "CEO Blog"
date: 2024-08-07
---


Stai per ricevere un tour esclusivo su come Blue costruisce Blue.

In Blue, consumiamo il nostro stesso prodotto.

Questo significa che utilizziamo Blue per *costruire* Blue.

Questo termine che può sembrare strano, spesso chiamato "dogfooding", è attribuito a Paul Maritz, un manager di Microsoft negli anni '80. Si dice che abbia inviato un'email con l'oggetto *"Mangiamo il nostro stesso cibo per cani"* per incoraggiare i dipendenti Microsoft a utilizzare i prodotti dell'azienda.

L'idea di utilizzare i propri strumenti per costruire i propri strumenti porta a un ciclo di feedback positivo.

L'idea di utilizzare i propri strumenti per costruire i propri strumenti porta a un ciclo di feedback positivo, creando numerosi vantaggi:

- **Ci aiuta a identificare rapidamente i problemi di usabilità nel mondo reale.** Poiché utilizziamo Blue quotidianamente, ci imbattiamo nelle stesse sfide che i nostri utenti potrebbero affrontare, permettendoci di affrontarle in modo proattivo.
- **Accelera la scoperta di bug.** L'uso interno rivela spesso bug prima che raggiungano i nostri clienti, migliorando la qualità complessiva del prodotto.
- **Aumenta la nostra empatia per gli utenti finali.** Il nostro team acquisisce esperienza diretta sui punti di forza e di debolezza di Blue, aiutandoci a prendere decisioni più incentrate sull'utente.
- **Promuove una cultura della qualità all'interno della nostra organizzazione.** Quando tutti utilizzano il prodotto, c'è un interesse condiviso nella sua eccellenza.
- **Favorisce l'innovazione.** L'uso regolare spesso genera idee per nuove funzionalità o miglioramenti, mantenendo Blue all'avanguardia.

[Abbiamo già parlato del motivo per cui non abbiamo un team di testing dedicato](/insights/open-beta) e questo è un ulteriore motivo.

Se ci sono bug nel nostro sistema, li troviamo quasi sempre nel nostro costante utilizzo quotidiano della piattaforma. E questo crea anche una funzione di costrizione per risolverli, poiché ovviamente li troveremmo molto fastidiosi essendo probabilmente uno dei principali utenti di Blue!

Questo approccio dimostra il nostro impegno verso il prodotto. Facendo affidamento su Blue, mostriamo ai nostri clienti che crediamo veramente in ciò che stiamo costruendo. Non è solo un prodotto che vendiamo – è uno strumento di cui dipendiamo ogni giorno.

## Processo Principale

Abbiamo un progetto in Blue, opportunamente chiamato "Prodotto".

**Tutto** ciò che riguarda lo sviluppo del nostro prodotto è tracciato qui. Feedback dei clienti, bug, idee per funzionalità, lavoro in corso, e così via. L'idea di avere un progetto in cui tracciamo tutto è che [promuove una migliore collaborazione.](/insights/great-teamwork)

Ogni record è una funzionalità o parte di una funzionalità. Questo è il modo in cui passiamo da "non sarebbe bello se..." a "guarda questa fantastica nuova funzionalità!"

Il progetto ha le seguenti liste:

- **Idee/Feedback**: Questa è una lista di idee del team o feedback dei clienti basati su chiamate o scambi di email. Sentiti libero di aggiungere qui qualsiasi idea! In questa lista, non abbiamo ancora deciso di costruire nessuna di queste funzionalità, ma le esaminiamo regolarmente per idee che vogliamo esplorare ulteriormente.
- **Backlog (Lungo Termine)**: Qui vanno le funzionalità della lista Idee/Feedback se decidiamo che sarebbero un buon aggiunta a Blue.
- **{Trimestre Corrente}**: Questa è tipicamente strutturata come "Qx YYYY" e mostra le nostre priorità trimestrali.
- **Bug**: Questa è una lista di bug noti segnalati dal team o dai clienti. I bug aggiunti qui avranno automaticamente l'etichetta "Bug".
- **Specifiche**: Queste funzionalità sono attualmente in fase di specifica. Non ogni funzionalità richiede una specifica o un design; dipende dalla dimensione prevista della funzionalità e dal livello di fiducia che abbiamo riguardo ai casi limite e alla complessità.
- **Backlog Design**: Questo è il backlog per i designer; ogni volta che hanno finito qualcosa che è in corso, possono scegliere qualsiasi elemento da questa lista.
- **Design in Corso**: Queste sono le funzionalità attuali che i designer stanno progettando.
- **Revisione Design**: Qui si trovano le funzionalità i cui design sono attualmente in fase di revisione.
- **Backlog (Breve Termine)**: Questa è una lista di funzionalità su cui probabilmente inizieremo a lavorare nelle prossime settimane. Qui avvengono le assegnazioni. Il CEO e il Responsabile Ingegneria decidono quali funzionalità sono assegnate a quale ingegnere in base all'esperienza precedente e al carico di lavoro. [I membri del team possono poi trasferirle in In Corso](/insights/push-vs-pull-kanban) una volta completato il loro lavoro attuale.
- **In Corso**: Queste sono funzionalità che sono attualmente in fase di sviluppo.
- **Revisione Codice**: Una volta che una funzionalità ha terminato lo sviluppo, viene sottoposta a revisione del codice. Poi verrà spostata di nuovo in "In Corso" se sono necessarie modifiche o distribuita nell'ambiente di Sviluppo.
- **Dev**: Queste sono tutte le funzionalità attualmente nell'ambiente di Sviluppo. Altri membri del team e alcuni clienti possono esaminarle.
- **Beta**: Queste sono tutte le funzionalità attualmente nell'[ambiente Beta](https://beta.app.blue.cc). Molti clienti utilizzano questo come la loro piattaforma Blue quotidiana e forniranno anche feedback.
- **Produzione**: Quando una funzionalità raggiunge la produzione, è considerata completata.

A volte, mentre sviluppiamo una funzionalità, ci rendiamo conto che alcuni sotto-funzionalità sono più difficili da implementare di quanto inizialmente previsto, e potremmo scegliere di non farle nella versione iniziale che distribuiamo ai clienti. In questo caso, possiamo creare un nuovo record con un nome seguendo il formato "{NomeFunzionalità} V2" e includere tutte le sotto-funzionalità come elementi di checklist.

## Etichette

- **Mobile**: Questo significa che la funzionalità è specifica per le nostre app iOS, Android o iPad.
- **{NomeClienteEnterprise}**: Una funzionalità è specificamente costruita per un cliente enterprise. Il tracciamento è importante poiché ci sono tipicamente ulteriori accordi commerciali per ogni funzionalità.
- **Bug**: Questo significa che si tratta di un bug che richiede una correzione.
- **Fast-Track**: Questo significa che si tratta di una modifica Fast-Track che non deve passare attraverso l'intero ciclo di rilascio come descritto sopra.
- **Principale**: Questa è una grande sviluppo di funzionalità. È tipicamente riservata per lavori infrastrutturali importanti, grandi aggiornamenti di dipendenze e nuovi moduli significativi all'interno di Blue.
- **AI**: Questa funzionalità contiene un componente di intelligenza artificiale.
- **Sicurezza**: Questo significa che deve essere esaminata un'implicazione di sicurezza o è necessario un patch.

L'etichetta fast-track è particolarmente interessante. Questa è riservata per aggiornamenti più piccoli e meno complessi che non richiedono il nostro intero ciclo di rilascio e che vogliamo spedire ai clienti entro 24-48 ore.

Le modifiche fast-track sono tipicamente aggiustamenti minori che possono migliorare significativamente l'esperienza dell'utente senza alterare la funzionalità principale. Pensa a correggere errori di battitura nell'interfaccia utente, modificare il padding dei pulsanti o aggiungere nuove icone per una migliore guida visiva. Questi sono il tipo di modifiche che, sebbene piccole, possono fare una grande differenza in come gli utenti percepiscono e interagiscono con il nostro prodotto. Sono anche fastidiose se ci vogliono secoli per essere spedite!

Il nostro processo fast-track è semplice.

Inizia creando un nuovo branch da main, implementando le modifiche e poi creando richieste di merge per ogni branch di destinazione - Dev, Beta e Produzione. Generiamo un link di anteprima per la revisione, assicurandoci che anche queste piccole modifiche soddisfino i nostri standard di qualità. Una volta approvate, le modifiche vengono unite simultaneamente in tutti i branch, mantenendo i nostri ambienti sincronizzati.

## Campi Personalizzati

Non abbiamo molti campi personalizzati nel nostro progetto Prodotto.

- **Specifiche**: Questo si collega a un documento Blue che ha la specifica per quella particolare funzionalità. Questo non è sempre fatto, poiché dipende dalla complessità della funzionalità.
- **MR**: Questo è il link alla Merge Request in [Gitlab](https://gitlab.com) dove ospitiamo il nostro codice.
- **Link Anteprima**: Per le funzionalità che cambiano principalmente il front-end, possiamo creare un URL unico che ha quelle modifiche per ogni commit, così possiamo facilmente rivedere le modifiche.
- **Lead**: Questo campo ci dice quale ingegnere senior sta seguendo la revisione del codice. Garantisce che ogni funzionalità riceva l'attenzione esperta che merita e c'è sempre una persona di riferimento chiara per domande o preoccupazioni.

## Checklist

Durante le nostre dimostrazioni settimanali, inseriremo il feedback discusso in una checklist chiamata "feedback" e ci sarà anche un'altra checklist che contiene il principale [WBS (Work Breakdown Scope)](/insights/simple-work-breakdown-structure) della funzionalità, così possiamo facilmente capire cosa è fatto e cosa deve ancora essere fatto.

## Conclusione

E questo è tutto!

Pensiamo che a volte le persone siano sorprese da quanto sia semplice il nostro processo, ma crediamo che processi semplici siano spesso di gran lunga superiori a processi eccessivamente complessi che non puoi facilmente comprendere.

Questa semplicità è intenzionale. Ci consente di rimanere agili, rispondere rapidamente alle esigenze dei clienti e mantenere allineato tutto il nostro team.

Utilizzando Blue per costruire Blue, non stiamo solo sviluppando un prodotto – lo stiamo vivendo.

Quindi, la prossima volta che utilizzi Blue, ricorda: non stai solo utilizzando un prodotto che abbiamo costruito. Stai utilizzando un prodotto di cui ci affidiamo personalmente ogni singolo giorno.

E questo fa tutta la differenza.
---
title: Automazione della gestione dei progetti — email per gli stakeholder.
description: Spesso, vuoi avere il controllo delle tue automazioni nella gestione dei progetti.
category: "Product Updates"
date: 2024-07-08
---


Abbiamo già trattato come [creare automazioni email in precedenza.](/insights/email-automations)

Tuttavia, spesso ci sono stakeholder nei progetti che devono essere avvisati solo quando c'è qualcosa di *veramente* importante.

Non sarebbe bello avere un'automazione nella gestione dei progetti in cui tu, come project manager, potessi controllare *esattamente* quando notificare un importante stakeholder con la pressione di un pulsante?

Beh, si scopre che con Blue puoi fare proprio questo!

Oggi impareremo a creare un'automazione nella gestione dei progetti davvero utile:

Una casella di controllo che notifica automaticamente uno o più stakeholder chiave, fornendo loro tutto il contesto fondamentale riguardo a ciò di cui li stai avvisando. Come punto bonus, impareremo anche come limitare questa capacità affinché solo alcuni membri del tuo progetto possano attivare questa notifica email.

Questo apparirà più o meno così una volta completato:

![](/insights/checkbox-email-automation.png)

E semplicemente selezionando questa casella di controllo, sarai in grado di attivare un'automazione nella gestione dei progetti per inviare un'email di notifica personalizzata agli stakeholder.

Procediamo passo dopo passo.

## 1. Crea il tuo campo personalizzato casella di controllo

È molto semplice, puoi consultare la nostra [documentazione dettagliata](https://documentation.blue.cc/custom-fields/introduction#creating-custom-fields) su come creare campi personalizzati.

Assicurati di dare a questo campo un nome ovvio che ricorderai, come "notifica gestione" o "notifica stakeholder".

## 2. Crea il tuo attivatore per l'automazione nella gestione dei progetti.

Nella vista dei record del tuo progetto, fai clic sul piccolo robot in alto a destra per aprire le impostazioni di automazione:

<video autoplay loop muted playsinline>
  <source src="/videos/notify-stakeholders-automation-setup.mp4" type="video/mp4">
</video>

## 3. Crea la tua azione per l'automazione nella gestione dei progetti.

In questo caso, la nostra azione sarà inviare una notifica email personalizzata a uno o più indirizzi email. È importante notare qui che queste persone **non** devono essere in Blue per ricevere queste email, puoi inviare email a *qualsiasi* indirizzo email.

Puoi saperne di più nella nostra [guida di documentazione dettagliata su come impostare automazioni email](https://documentation.blue.cc/automations/actions/email-automations)

Il tuo risultato finale dovrebbe apparire più o meno così:

![](/insights/email-automation-example.png)

## 4. Bonus: Limita l'accesso alla casella di controllo.

Puoi utilizzare [ruoli utente personalizzati in Blue](/platform/features/user-permissions) per limitare l'accesso ai campi personalizzati delle caselle di controllo, assicurando che solo i membri autorizzati del team possano attivare le notifiche email.

Blue consente agli Amministratori di Progetto di definire ruoli e assegnare permessi ai gruppi di utenti. Questo sistema è cruciale per mantenere il controllo su chi può interagire con elementi specifici del tuo progetto, inclusi campi personalizzati come la casella di controllo per le notifiche.

1. Naviga nella sezione Gestione Utenti in Blue e seleziona "Ruoli Utente Personalizzati."
2. Crea un nuovo ruolo fornendo un nome descrittivo e una descrizione opzionale.
3. All'interno dei permessi del ruolo, individua la sezione per l'accesso ai Campi Personalizzati.
4. Specifica se il ruolo può visualizzare o modificare il campo personalizzato della casella di controllo. Ad esempio, limita l'accesso alla modifica a ruoli come "Amministratore di Progetto" mentre consenti a un ruolo personalizzato appena creato di gestire questo campo.
5. Assegna il ruolo appena creato agli utenti o ai gruppi di utenti appropriati. Questo assicura che solo le persone designate abbiano la possibilità di interagire con la casella di controllo per le notifiche.

[Leggi di più sul nostro sito di documentazione ufficiale.](https://documentation.blue.cc/user-management/roles/custom-user-roles)

Implementando questi ruoli personalizzati, migliori la sicurezza e l'integrità dei tuoi processi di gestione dei progetti. Solo i membri autorizzati del team possono attivare notifiche email critiche, assicurando che gli stakeholder ricevano aggiornamenti importanti senza avvisi non necessari.

## Conclusione

Implementando l'automazione nella gestione dei progetti descritta sopra, ottieni un controllo preciso su quando e come notificare gli stakeholder chiave. Questo approccio garantisce che gli aggiornamenti importanti vengano comunicati in modo efficace, senza sopraffare i tuoi stakeholder con informazioni non necessarie. Utilizzando le funzionalità di campo personalizzato e automazione di Blue, puoi semplificare il tuo processo di gestione dei progetti, migliorare la comunicazione e mantenere un alto livello di efficienza.

Con una semplice casella di controllo, puoi attivare notifiche email personalizzate su misura per le esigenze del tuo progetto, assicurando che le persone giuste siano informate al momento giusto. Inoltre, la possibilità di limitare questa funzionalità a membri specifici del team aggiunge un ulteriore livello di controllo e sicurezza.

Inizia a sfruttare questa potente funzionalità in Blue oggi stesso per tenere informati i tuoi stakeholder e far funzionare i tuoi progetti senza intoppi. Per ulteriori passaggi dettagliati e opzioni di personalizzazione aggiuntive, fai riferimento ai link di documentazione forniti. Buona automazione!
---
title: Crea Dashboard
description: Crea un nuovo dashboard per la visualizzazione dei dati e la reportistica in Blue
---

## Crea un Dashboard

La mutazione `createDashboard` consente di creare un nuovo dashboard all'interno della tua azienda o progetto. I dashboard sono potenti strumenti di visualizzazione che aiutano i team a monitorare le metriche, seguire i progressi e prendere decisioni basate sui dati.

### Esempio Base

```graphql
mutation CreateDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      title: "Sales Performance Dashboard"
    }
  ) {
    id
    title
    createdBy {
      id
      email
      firstName
      lastName
    }
    createdAt
  }
}
```

### Dashboard Specifico del Progetto

Crea un dashboard associato a un progetto specifico:

```graphql
mutation CreateProjectDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      projectId: "proj_xyz789"
      title: "Q4 Project Metrics"
    }
  ) {
    id
    title
    project {
      id
      name
    }
    createdBy {
      id
      email
    }
    dashboardUsers {
      id
      user {
        id
        email
      }
      role
    }
    createdAt
  }
}
```

## Parametri di Input

### CreateDashboardInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `companyId` | String! | ✅ Sì | L'ID dell'azienda in cui verrà creato il dashboard |
| `title` | String! | ✅ Sì | Il nome del dashboard. Deve essere una stringa non vuota |
| `projectId` | String | No | ID opzionale di un progetto da associare a questo dashboard |

## Campi di Risposta

La mutazione restituisce un oggetto completo `Dashboard`:

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il dashboard creato |
| `title` | String! | Il titolo del dashboard come fornito |
| `companyId` | String! | L'azienda a cui appartiene questo dashboard |
| `projectId` | String | L'ID del progetto associato (se fornito) |
| `project` | Project | L'oggetto progetto associato (se è stato fornito projectId) |
| `createdBy` | User! | L'utente che ha creato il dashboard (tu) |
| `dashboardUsers` | [DashboardUser!]! | Elenco degli utenti con accesso (inizialmente solo il creatore) |
| `createdAt` | DateTime! | Timestamp di quando è stato creato il dashboard |
| `updatedAt` | DateTime! | Timestamp dell'ultima modifica (stesso di createdAt per i nuovi dashboard) |

### Campi DashboardUser

Quando un dashboard viene creato, il creatore viene automaticamente aggiunto come utente del dashboard:

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per la relazione utente-dashboard |
| `user` | User! | L'oggetto utente con accesso al dashboard |
| `role` | DashboardRole! | Il ruolo dell'utente (il creatore ottiene accesso completo) |
| `dashboard` | Dashboard! | Riferimento al dashboard |

## Permessi Richiesti

Qualsiasi utente autenticato che appartiene all'azienda specificata può creare dashboard. Non ci sono requisiti di ruolo speciali.

| Stato Utente | Può Creare Dashboard |
|--------------|---------------------|
| Company Member | ✅ Sì |
| Non-Membro dell'Azienda | ❌ No |
| Unauthenticated | ❌ No |

## Risposte di Errore

### Azienda Non Valida
```json
{
  "errors": [{
    "message": "Company not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Utente Non in Azienda
```json
{
  "errors": [{
    "message": "You don't have access to this company",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Progetto Non Valido
```json
{
  "errors": [{
    "message": "Project not found or doesn't belong to the specified company",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Titolo Vuoto
```json
{
  "errors": [{
    "message": "Dashboard title cannot be empty",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Note Importanti

- **Proprietà automatica**: L'utente che crea il dashboard diventa automaticamente il suo proprietario con permessi completi
- **Associazione al progetto**: Se fornisci un `projectId`, deve appartenere alla stessa azienda
- **Permessi iniziali**: Solo il creatore ha accesso inizialmente. Usa `editDashboard` per aggiungere più utenti
- **Requisiti del titolo**: I titoli dei dashboard devono essere stringhe non vuote. Non c'è requisito di unicità
- **Appartenenza all'azienda**: Devi essere un membro dell'azienda per creare dashboard in essa

## Flusso di Creazione del Dashboard

1. **Crea il dashboard** utilizzando questa mutazione
2. **Configura grafici e widget** utilizzando l'interfaccia del costruttore di dashboard
3. **Aggiungi membri del team** utilizzando la mutazione `editDashboard` con `dashboardUsers`
4. **Imposta filtri e intervalli di date** tramite l'interfaccia del dashboard
5. **Condividi o incorpora** il dashboard utilizzando il suo ID unico

## Casi d'Uso

1. **Dashboard esecutivi**: Crea panoramiche di alto livello delle metriche aziendali
2. **Monitoraggio dei progetti**: Costruisci dashboard specifici per i progetti per monitorare i progressi
3. **Performance del team**: Monitora la produttività del team e le metriche di realizzazione
4. **Reportistica per i clienti**: Crea dashboard per report rivolti ai clienti
5. **Monitoraggio in tempo reale**: Imposta dashboard per dati operativi dal vivo

## Migliori Pratiche

1. **Convenzioni di denominazione**: Usa titoli chiari e descrittivi che indicano lo scopo del dashboard
2. **Associazione al progetto**: Collega i dashboard ai progetti quando sono specifici per il progetto
3. **Gestione degli accessi**: Aggiungi i membri del team immediatamente dopo la creazione per la collaborazione
4. **Organizzazione**: Crea una gerarchia di dashboard utilizzando schemi di denominazione coerenti

## Operazioni Correlate

- [Elenca Dashboard](/api/dashboards/) - Recupera tutti i dashboard per un'azienda o un progetto
- [Modifica Dashboard](/api/dashboards/rename-dashboard) - Rinomina il dashboard o gestisci gli utenti
- [Copia Dashboard](/api/dashboards/copy-dashboard) - Duplica un dashboard esistente
- [Elimina Dashboard](/api/dashboards/delete-dashboard) - Rimuovi un dashboard
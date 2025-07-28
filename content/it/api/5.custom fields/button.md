---
title: Campo Personalizzato Pulsante
description: Crea campi pulsante interattivi che attivano automazioni quando vengono cliccati
---

I campi personalizzati pulsante forniscono elementi UI interattivi che attivano automazioni quando vengono cliccati. A differenza di altri tipi di campi personalizzati che memorizzano dati, i campi pulsante fungono da attivatori di azioni per eseguire flussi di lavoro configurati.

## Esempio di Base

Crea un semplice campo pulsante che attiva un'automazione:

```graphql
mutation CreateButtonField {
  createCustomField(input: {
    name: "Send Invoice"
    type: BUTTON
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un pulsante con requisiti di conferma:

```graphql
mutation CreateButtonWithConfirmation {
  createCustomField(input: {
    name: "Delete All Attachments"
    type: BUTTON
    projectId: "proj_123"
    buttonType: "hardConfirmation"
    buttonConfirmText: "DELETE"
    description: "Permanently removes all attachments from this task"
  }) {
    id
    name
    type
    buttonType
    buttonConfirmText
    description
  }
}
```

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del pulsante |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `BUTTON` |
| `projectId` | String! | ✅ Sì | ID del progetto in cui verrà creato il campo |
| `buttonType` | String | No | Comportamento di conferma (vedi Tipi di Pulsante qui sotto) |
| `buttonConfirmText` | String | No | Testo che gli utenti devono digitare per la conferma rigorosa |
| `description` | String | No | Testo di aiuto mostrato agli utenti |
| `required` | Boolean | No | Se il campo è richiesto (di default è falso) |
| `isActive` | Boolean | No | Se il campo è attivo (di default è vero) |

### Tipo di Campo Pulsante

Il campo `buttonType` è una stringa libera che può essere utilizzata dai client UI per determinare il comportamento di conferma. I valori comuni includono:

- `""` (vuoto) - Nessuna conferma
- `"soft"` - Dialogo di conferma semplice
- `"hard"` - Richiedere di digitare il testo di conferma

**Nota**: Questi sono solo suggerimenti UI. L'API non convalida né applica valori specifici.

## Attivazione dei Clic sui Pulsanti

Per attivare un clic su un pulsante ed eseguire automazioni associate:

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### Parametri di Input per il Clic

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del compito contenente il pulsante |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato pulsante |

### Importante: Comportamento dell'API

**Tutti i clic sui pulsanti tramite l'API vengono eseguiti immediatamente** indipendentemente da eventuali impostazioni `buttonType` o `buttonConfirmText`. Questi campi sono memorizzati per i client UI per implementare dialoghi di conferma, ma l'API stessa:

- Non convalida il testo di conferma
- Non applica alcun requisito di conferma
- Esegue immediatamente l'azione del pulsante quando viene chiamata

La conferma è puramente una funzionalità di sicurezza dell'interfaccia utente lato client.

### Esempio: Cliccare su Diversi Tipi di Pulsanti

```graphql
# Button with no confirmation
mutation ClickSimpleButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "simple_button_id"
  })
}

# Button with soft confirmation (API call is the same!)
mutation ClickSoftConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "soft_confirm_button_id"
  })
}

# Button with hard confirmation (API call is still the same!)
mutation ClickHardConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "hard_confirm_button_id"
  })
}
```

Tutte e tre le mutazioni sopra eseguiranno immediatamente l'azione del pulsante quando chiamate tramite l'API, bypassando eventuali requisiti di conferma.

## Campi di Risposta

### Risposta Campo Personalizzato

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il campo personalizzato |
| `name` | String! | Nome visualizzato del pulsante |
| `type` | CustomFieldType! | Sempre `BUTTON` per i campi pulsante |
| `buttonType` | String | Impostazione del comportamento di conferma |
| `buttonConfirmText` | String | Testo di conferma richiesto (se si utilizza la conferma rigorosa) |
| `description` | String | Testo di aiuto per gli utenti |
| `required` | Boolean! | Se il campo è richiesto |
| `isActive` | Boolean! | Se il campo è attualmente attivo |
| `projectId` | String! | ID del progetto a cui appartiene questo campo |
| `createdAt` | DateTime! | Quando è stato creato il campo |
| `updatedAt` | DateTime! | Quando è stato modificato per l'ultima volta il campo |

## Come Funzionano i Campi Pulsante

### Integrazione con l'Automazione

I campi pulsante sono progettati per funzionare con il sistema di automazione di Blue:

1. **Crea il campo pulsante** utilizzando la mutazione sopra
2. **Configura le automazioni** che ascoltano gli eventi `CUSTOM_FIELD_BUTTON_CLICKED`
3. **Gli utenti cliccano sul pulsante** nell'interfaccia utente
4. **Le automazioni eseguono** le azioni configurate

### Flusso degli Eventi

Quando un pulsante viene cliccato:

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### Nessuna Memorizzazione dei Dati

Importante: I campi pulsante non memorizzano alcun valore di dati. Servono puramente come attivatori di azioni. Ogni clic:
- Genera un evento
- Attiva automazioni associate
- Registra un'azione nella cronologia del compito
- Non modifica alcun valore di campo

## Permessi Richiesti

Gli utenti hanno bisogno di ruoli di progetto appropriati per creare e utilizzare i campi pulsante:

| Azione | Ruolo Richiesto |
|--------|-------------------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## Risposte di Errore

### Permesso Negato
```json
{
  "errors": [{
    "message": "You don't have permission to edit this custom field",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Campo Personalizzato Non Trovato
```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**Nota**: L'API non restituisce errori specifici per automazioni mancanti o discrepanze di conferma.

## Migliori Pratiche

### Convenzioni di Nominazione
- Usa nomi orientati all'azione: "Invia Fattura", "Crea Rapporto", "Notifica Team"
- Sii specifico su cosa fa il pulsante
- Evita nomi generici come "Pulsante 1" o "Clicca Qui"

### Impostazioni di Conferma
- Lascia `buttonType` vuoto per azioni sicure e reversibili
- Imposta `buttonType` per suggerire il comportamento di conferma ai client UI
- Usa `buttonConfirmText` per specificare cosa gli utenti dovrebbero digitare nelle conferme UI
- Ricorda: Questi sono solo suggerimenti UI - le chiamate API vengono sempre eseguite immediatamente

### Progettazione dell'Automazione
- Mantieni le azioni del pulsante focalizzate su un singolo flusso di lavoro
- Fornisci un feedback chiaro su cosa è successo dopo il clic
- Considera di aggiungere un testo descrittivo per spiegare lo scopo del pulsante

## Casi d'Uso Comuni

1. **Transizioni di Flusso di Lavoro**
   - "Segna come Completo"
   - "Invia per Approvazione"
   - "Archivia Compito"

2. **Integrazioni Esterne**
   - "Sincronizza con CRM"
   - "Genera Fattura"
   - "Invia Aggiornamento Email"

3. **Operazioni di Gruppo**
   - "Aggiorna Tutti i Sottocompiti"
   - "Copia nei Progetti"
   - "Applica Modello"

4. **Azioni di Reporting**
   - "Genera Rapporto"
   - "Esporta Dati"
   - "Crea Riepilogo"

## Limitazioni

- I pulsanti non possono memorizzare o visualizzare valori di dati
- Ogni pulsante può solo attivare automazioni, non chiamate API dirette (tuttavia, le automazioni possono includere azioni di richiesta HTTP per chiamare API esterne o le API di Blue)
- La visibilità del pulsante non può essere controllata in modo condizionale
- Massimo di un'esecuzione di automazione per clic (anche se quell'automazione può attivare più azioni)

## Risorse Correlate

- [API Automazioni](/api/automations/index) - Configura azioni attivate dai pulsanti
- [Panoramica dei Campi Personalizzati](/custom-fields/list-custom-fields) - Concetti generali sui campi personalizzati
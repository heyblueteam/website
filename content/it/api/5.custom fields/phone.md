---
title: Campo Personalizzato per il Telefono
description: Crea campi telefonici per memorizzare e convalidare numeri di telefono con formattazione internazionale
---

I campi personalizzati per il telefono ti consentono di memorizzare numeri di telefono nei record con convalida integrata e formattazione internazionale. Sono ideali per tenere traccia delle informazioni di contatto, dei contatti di emergenza o di qualsiasi dato relativo al telefono nei tuoi progetti.

## Esempio di Base

Crea un semplice campo telefonico:

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un campo telefonico con descrizione:

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
  }) {
    id
    name
    type
    description
  }
}
```

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo telefonico |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `PHONE` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

**Nota**: I campi personalizzati sono automaticamente associati al progetto in base al contesto del progetto corrente dell'utente. Non è richiesto alcun parametro `projectId`.

## Impostazione dei Valori Telefonici

Per impostare o aggiornare un valore telefonico su un record:

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### Parametri di SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato per il telefono |
| `text` | String | No | Numero di telefono con prefisso internazionale |
| `regionCode` | String | No | Prefisso internazionale (rilevato automaticamente) |

**Nota**: Anche se `text` è opzionale nello schema, è richiesto un numero di telefono affinché il campo abbia significato. Quando si utilizza `setTodoCustomField`, non viene eseguita alcuna convalida - puoi memorizzare qualsiasi valore di testo e regionCode. La rilevazione automatica avviene solo durante la creazione del record.

## Creazione di Record con Valori Telefonici

Quando crei un nuovo record con valori telefonici:

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      text
      regionCode
    }
  }
}
```

## Campi di Risposta

### TodoCustomField Risposta

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `text` | String | Il numero di telefono formattato (formato internazionale) |
| `regionCode` | String | Il prefisso internazionale (ad es., "US", "GB", "CA") |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato l'ultimo valore |

## Convalida del Numero di Telefono

**Importante**: La convalida e la formattazione del numero di telefono avvengono solo quando si creano nuovi record tramite `createTodo`. Quando si aggiornano valori telefonici esistenti utilizzando `setTodoCustomField`, non viene eseguita alcuna convalida e i valori vengono memorizzati come forniti.

### Formati Accettati (Durante la Creazione del Record)
I numeri di telefono devono includere un prefisso internazionale in uno di questi formati:

- **Formato E.164 (preferito)**: `+12345678900`
- **Formato internazionale**: `+1 234 567 8900`
- **Internazionale con punteggiatura**: `+1 (234) 567-8900`
- **Prefisso internazionale con trattini**: `+1-234-567-8900`

**Nota**: I formati nazionali senza prefisso internazionale (come `(234) 567-8900`) verranno rifiutati durante la creazione del record.

### Regole di Convalida (Durante la Creazione del Record)
- Utilizza libphonenumber-js per il parsing e la convalida
- Accetta vari formati di numeri di telefono internazionali
- Rileva automaticamente il paese dal numero
- Formatta il numero nel formato di visualizzazione internazionale (ad es., `+1 234 567 8900`)
- Estrae e memorizza il prefisso internazionale separatamente (ad es., `US`)

### Esempi di Numeri di Telefono Validi
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### Esempi di Numeri di Telefono Non Validi
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## Formato di Archiviazione

Quando si creano record con numeri di telefono:
- **text**: Memorizzato in formato internazionale (ad es., `+1 234 567 8900`) dopo la convalida
- **regionCode**: Memorizzato come codice paese ISO (ad es., `US`, `GB`, `CA`) rilevato automaticamente

Quando si aggiorna tramite `setTodoCustomField`:
- **text**: Memorizzato esattamente come fornito (senza formattazione)
- **regionCode**: Memorizzato esattamente come fornito (senza convalida)

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## Risposte di Errore

### Formato Telefonico Non Valido
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Campo Non Trovato
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

### Codice Paese Mancante
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Migliori Pratiche

### Inserimento Dati
- Includi sempre il codice paese nei numeri di telefono
- Utilizza il formato E.164 per coerenza
- Convalida i numeri prima di memorizzarli per operazioni importanti
- Considera le preferenze regionali per la formattazione di visualizzazione

### Qualità dei Dati
- Memorizza i numeri in formato internazionale per compatibilità globale
- Utilizza regionCode per funzionalità specifiche del paese
- Convalida i numeri di telefono prima di operazioni critiche (SMS, chiamate)
- Considera le implicazioni del fuso orario per il timing dei contatti

### Considerazioni Internazionali
- Il codice paese viene rilevato e memorizzato automaticamente
- I numeri sono formattati secondo lo standard internazionale
- Le preferenze di visualizzazione regionali possono utilizzare regionCode
- Considera le convenzioni di composizione locale quando visualizzi

## Casi d'Uso Comuni

1. **Gestione dei Contatti**
   - Numeri di telefono dei clienti
   - Informazioni di contatto dei fornitori
   - Numeri di telefono dei membri del team
   - Dettagli di contatto per il supporto

2. **Contatti di Emergenza**
   - Numeri di contatto di emergenza
   - Informazioni di contatto per il servizio di guardia
   - Contatti per la risposta alle crisi
   - Numeri di telefono per l'escalation

3. **Supporto Clienti**
   - Numeri di telefono dei clienti
   - Numeri di richiamo per il supporto
   - Numeri di telefono per la verifica
   - Numeri di contatto per il follow-up

4. **Vendite e Marketing**
   - Numeri di telefono dei lead
   - Liste di contatto per campagne
   - Informazioni di contatto dei partner
   - Numeri di telefono delle fonti di riferimento

## Funzionalità di Integrazione

### Con Automazioni
- Attiva azioni quando i campi telefonici vengono aggiornati
- Invia notifiche SMS ai numeri di telefono memorizzati
- Crea attività di follow-up basate su modifiche telefoniche
- Instrada le chiamate in base ai dati del numero di telefono

### Con Ricerche
- Riferisci i dati telefonici da altri record
- Aggrega liste telefoniche da più fonti
- Trova record per numero di telefono
- Incrocia informazioni di contatto

### Con Moduli
- Convalida telefonica automatica
- Verifica del formato internazionale
- Rilevamento del codice paese
- Feedback sul formato in tempo reale

## Limitazioni

- Richiede il codice paese per tutti i numeri
- Nessuna funzionalità SMS o di chiamata integrata
- Nessuna verifica del numero di telefono oltre al controllo del formato
- Nessun archiviazione dei metadati telefonici (operatore, tipo, ecc.)
- I numeri in formato nazionale senza codice paese vengono rifiutati
- Nessuna formattazione automatica del numero di telefono nell'interfaccia utente oltre allo standard internazionale

## Risorse Correlate

- [Campi di Testo](/api/custom-fields/text-single) - Per dati di testo non telefonici
- [Campi Email](/api/custom-fields/email) - Per indirizzi email
- [Campi URL](/api/custom-fields/url) - Per indirizzi web
- [Panoramica dei Campi Personalizzati](/custom-fields/list-custom-fields) - Concetti generali
---
title: Campo Personalizzato Email
description: Crea campi email per memorizzare e convalidare indirizzi email
---

I campi personalizzati email ti consentono di memorizzare indirizzi email nei record con convalida integrata. Sono ideali per tenere traccia delle informazioni di contatto, delle email degli assegnatari o di qualsiasi dato relativo alle email nei tuoi progetti.

## Esempio Base

Crea un semplice campo email:

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un campo email con descrizione:

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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

| Parametro | Tipo | Obbligatorio | Descrizione |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo email |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `EMAIL` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

## Impostazione dei Valori Email

Per impostare o aggiornare un valore email su un record:

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### Parametri SetTodoCustomFieldInput

| Parametro | Tipo | Obbligatorio | Descrizione |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo email personalizzato |
| `text` | String | No | Indirizzo email da memorizzare |

## Creazione di Record con Valori Email

Quando crei un nuovo record con valori email:

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## Campi di Risposta

### Risposta CustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | ID! | Identificatore unico per il campo personalizzato |
| `name` | String! | Nome visualizzato del campo email |
| `type` | CustomFieldType! | Il tipo di campo (EMAIL) |
| `description` | String | Testo di aiuto per il campo |
| `value` | JSON | Contiene il valore email (vedi sotto) |
| `createdAt` | DateTime! | Quando è stato creato il campo |
| `updatedAt` | DateTime! | Quando il campo è stato modificato per l'ultima volta |

**Importante**: I valori email sono accessibili attraverso il campo `customField.value.text`, non direttamente nella risposta.

## Interrogazione dei Valori Email

Quando interroghi record con campi email personalizzati, accedi all'email attraverso il percorso `customField.value.text`:

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

La risposta includerà l'email nella struttura nidificata:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## Convalida Email

### Convalida del Modulo
Quando i campi email sono utilizzati nei moduli, convalidano automaticamente il formato email:
- Utilizza regole standard di convalida email
- Rimuove gli spazi bianchi dall'input
- Rifiuta formati email non validi

### Regole di Convalida
- Deve contenere un simbolo `@`
- Deve avere un formato di dominio valido
- Gli spazi bianchi iniziali/finali vengono rimossi automaticamente
- Formati email comuni sono accettati

### Esempi di Email Valide
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### Esempi di Email Non Valide
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## Note Importanti

### API Diretta vs Moduli
- **Moduli**: La convalida email automatica è applicata
- **API Diretta**: Nessuna convalida - qualsiasi testo può essere memorizzato
- **Raccomandazione**: Utilizza i moduli per l'input degli utenti per garantire la convalida

### Formato di Memorizzazione
- Gli indirizzi email sono memorizzati come testo semplice
- Nessuna formattazione o parsing speciale
- Sensibilità al caso: i campi email personalizzati sono memorizzati in modo sensibile al caso (a differenza delle email di autenticazione utente che sono normalizzate in minuscolo)
- Nessun limite di lunghezza massima oltre ai vincoli del database (limite di 16MB)

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## Risposte di Errore

### Formato Email Non Valido (Solo Moduli)
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
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
      "code": "NOT_FOUND"
    }
  }]
}
```

## Migliori Pratiche

### Inserimento Dati
- Convalida sempre gli indirizzi email nella tua applicazione
- Utilizza i campi email solo per indirizzi email reali
- Considera di utilizzare i moduli per l'input degli utenti per ottenere una convalida automatica

### Qualità dei Dati
- Rimuovi gli spazi bianchi prima di memorizzare
- Considera la normalizzazione del caso (tipicamente in minuscolo)
- Convalida il formato email prima di operazioni importanti

### Considerazioni sulla Privacy
- Gli indirizzi email sono memorizzati come testo semplice
- Considera le normative sulla privacy dei dati (GDPR, CCPA)
- Implementa controlli di accesso appropriati

## Casi d'Uso Comuni

1. **Gestione Contatti**
   - Indirizzi email dei clienti
   - Informazioni di contatto dei fornitori
   - Email dei membri del team
   - Dettagli di contatto per il supporto

2. **Gestione Progetti**
   - Email degli stakeholder
   - Email di contatto per le approvazioni
   - Destinatari delle notifiche
   - Email dei collaboratori esterni

3. **Supporto Clienti**
   - Indirizzi email dei clienti
   - Contatti per i ticket di supporto
   - Contatti per le escalation
   - Indirizzi email per il feedback

4. **Vendite e Marketing**
   - Indirizzi email dei lead
   - Liste di contatto per le campagne
   - Informazioni di contatto dei partner
   - Email delle fonti di riferimento

## Funzionalità di Integrazione

### Con Automazioni
- Attiva azioni quando i campi email vengono aggiornati
- Invia notifiche agli indirizzi email memorizzati
- Crea attività di follow-up basate sulle modifiche delle email

### Con Ricerche
- Riferisci dati email da altri record
- Aggrega liste email da più fonti
- Trova record per indirizzo email

### Con Moduli
- Convalida email automatica
- Controllo del formato email
- Rimozione degli spazi bianchi

## Limitazioni

- Nessuna verifica o convalida email integrata oltre al controllo del formato
- Nessuna funzionalità UI specifica per email (come link email cliccabili)
- Memorizzato come testo semplice senza crittografia
- Nessuna capacità di composizione o invio email
- Nessuna memorizzazione dei metadati email (nome visualizzato, ecc.)
- Le chiamate API dirette bypassano la convalida (solo i moduli convalidano)

## Risorse Correlate

- [Campi di Testo](/api/custom-fields/text-single) - Per dati di testo non email
- [Campi URL](/api/custom-fields/url) - Per indirizzi di siti web
- [Campi Telefonici](/api/custom-fields/phone) - Per numeri di telefono
- [Panoramica dei Campi Personalizzati](/api/custom-fields/list-custom-fields) - Concetti generali
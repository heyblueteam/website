---
title: Campo Personalizzato Formula
description: Crea campi calcolati che calcolano automaticamente valori basati su altri dati
---

I campi personalizzati formula sono utilizzati per calcoli di grafici e dashboard all'interno di Blue. Definiscono funzioni di aggregazione (SOMMA, MEDIA, CONTEGGIO, ecc.) che operano sui dati dei campi personalizzati per visualizzare metriche calcolate nei grafici. Le formule non vengono calcolate a livello di singolo todo, ma aggregano i dati su più record per scopi di visualizzazione.

## Esempio di Base

Crea un campo formula per calcoli di grafico:

```graphql
mutation CreateFormulaField {
  createCustomField(input: {
    name: "Budget Total"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Budget)"
        html: "<span>SUM(Budget)</span>"
      }
      display: {
        type: NUMBER
        precision: 2
        function: SUM
      }
    }
  }) {
    id
    name
    type
    formula
  }
}
```

## Esempio Avanzato

Crea una formula di valuta con calcoli complessi:

```graphql
mutation CreateCurrencyFormula {
  createCustomField(input: {
    name: "Profit Margin"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Revenue) - SUM(Costs)"
        html: "<span>SUM(Revenue) - SUM(Costs)</span>"
      }
      display: {
        type: CURRENCY
        currency: {
          code: "USD"
          name: "US Dollar"
        }
        precision: 2
      }
    }
    description: "Automatically calculates profit by subtracting costs from revenue"
  }) {
    id
    name
    type
    formula
  }
}
```

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo formula |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `FORMULA` |
| `projectId` | String! | ✅ Sì | L'ID del progetto in cui verrà creato questo campo |
| `formula` | JSON | No | Definizione della formula per calcoli di grafico |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

### Struttura della Formula

```json
{
  "logic": {
    "text": "Display text for the formula",
    "html": "HTML formatted display text"
  },
  "display": {
    "type": "NUMBER|CURRENCY|PERCENTAGE",
    "currency": {
      "code": "USD",
      "name": "US Dollar"  
    },
    "precision": 2,
    "function": "SUM|AVERAGE|AVERAGEA|COUNT|COUNTA|MAX|MIN"
  }
}
```

## Funzioni Supportate

### Funzioni di Aggregazione per Grafici

I campi formula supportano le seguenti funzioni di aggregazione per calcoli di grafico:

| Funzione | Descrizione | Enum ChartFunction |
|----------|-------------|-------------------|
| `SUM` | Somma di tutti i valori | `SUM` |
| `AVERAGE` | Media dei valori numerici | `AVERAGE` |
| `AVERAGEA` | Media escludendo zeri e null | `AVERAGEA` |
| `COUNT` | Conteggio dei valori | `COUNT` |
| `COUNTA` | Conteggio escludendo zeri e null | `COUNTA` |
| `MAX` | Valore massimo | `MAX` |
| `MIN` | Valore minimo | `MIN` |

**Nota**: Queste funzioni sono utilizzate nel campo `display.function` e operano su dati aggregati per visualizzazioni di grafico. Espressioni matematiche complesse o calcoli a livello di campo non sono supportati.

## Tipi di Visualizzazione

### Visualizzazione Numero

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

Risultato: `1250.75`

### Visualizzazione Valuta

```json
{
  "display": {
    "type": "CURRENCY",
    "currency": {
      "code": "USD",
      "name": "US Dollar"
    },
    "precision": 2
  }
}
```

Risultato: `$1,250.75`

### Visualizzazione Percentuale

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

Risultato: `87.5%`

## Modifica dei Campi Formula

Aggiorna i campi formula esistenti:

```graphql
mutation EditFormulaField {
  editCustomField(input: {
    customFieldId: "field_456"
    formula: {
      logic: {
        text: "AVERAGE(Score)"
        html: "<span>AVERAGE(Score)</span>"
      }
      display: {
        type: PERCENTAGE
        precision: 1
      }
    }
  }) {
    id
    formula
  }
}
```

## Elaborazione della Formula

### Contesto di Calcolo del Grafico

I campi formula vengono elaborati nel contesto dei segmenti di grafico e delle dashboard:
- I calcoli avvengono quando i grafici vengono renderizzati o aggiornati
- I risultati sono memorizzati in `ChartSegment.formulaResult` come valori decimali
- L'elaborazione è gestita tramite una coda BullMQ dedicata chiamata 'formula'
- Gli aggiornamenti vengono pubblicati agli abbonati della dashboard per aggiornamenti in tempo reale

### Formattazione della Visualizzazione

La funzione `getFormulaDisplayValue` formatta i risultati calcolati in base al tipo di visualizzazione:
- **NUMERO**: Visualizza come numero semplice con precisione opzionale
- **PERCENTUALE**: Aggiunge il suffisso % con precisione opzionale  
- **VALUTA**: Formattta utilizzando il codice valuta specificato

## Archiviazione dei Risultati della Formula

I risultati sono memorizzati nel campo `formulaResult`:

```json
{
  "number": 1250.75,
  "formulaResult": {
    "number": 1250.75,
    "display": {
      "type": "CURRENCY",
      "currency": {
        "code": "USD",
        "name": "US Dollar"
      },
      "precision": 2
    }
  }
}
```

## Campi di Risposta

### Risposta TodoCustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore univoco per il valore del campo |
| `customField` | CustomField! | La definizione del campo formula |
| `number` | Float | Risultato numerico calcolato |
| `formulaResult` | JSON | Risultato completo con formattazione di visualizzazione |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato calcolato l'ultimo valore |

## Contesto dei Dati

### Origine Dati del Grafico

I campi formula operano nel contesto dell'origine dati del grafico:
- Le formule aggregano i valori dei campi personalizzati tra i todo in un progetto
- La funzione di aggregazione specificata in `display.function` determina il calcolo
- I risultati sono calcolati utilizzando funzioni di aggregazione SQL (avg, sum, count, ecc.)
- I calcoli vengono eseguiti a livello di database per efficienza

## Esempi Comuni di Formula

### Budget Totale (Visualizzazione Grafico)

```json
{
  "logic": {
    "text": "Total Budget",
    "html": "<span>Total Budget</span>"
  },
  "display": {
    "type": "CURRENCY",
    "currency": { "code": "USD", "name": "US Dollar" },
    "precision": 2,
    "function": "SUM"
  }
}
```

### Punteggio Medio (Visualizzazione Grafico)

```json
{
  "logic": {
    "text": "Average Quality Score",
    "html": "<span>Average Quality Score</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 1,
    "function": "AVERAGE"
  }
}
```

### Conteggio Attività (Visualizzazione Grafico)

```json
{
  "logic": {
    "text": "Total Tasks",
    "html": "<span>Total Tasks</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 0,
    "function": "COUNT"
  }
}
```

## Permessi Richiesti

Le operazioni sui campi personalizzati seguono i permessi standard basati sui ruoli:

| Azione | Ruolo Richiesto |
|--------|-----------------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**Nota**: I ruoli specifici richiesti dipendono dalla configurazione del ruolo personalizzato del tuo progetto. Non ci sono costanti di permesso speciali come CUSTOM_FIELDS_CREATE.

## Gestione degli Errori

### Errore di Validazione
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Campo Personalizzato Non Trovato
```json
{
  "errors": [{
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Migliori Pratiche

### Progettazione della Formula
- Utilizza nomi chiari e descrittivi per i campi formula
- Aggiungi descrizioni che spiegano la logica di calcolo
- Testa le formule con dati di esempio prima del deployment
- Mantieni le formule semplici e leggibili

### Ottimizzazione delle Prestazioni
- Evita dipendenze di formula profondamente annidate
- Usa riferimenti specifici ai campi piuttosto che jolly
- Considera strategie di caching per calcoli complessi
- Monitora le prestazioni delle formule in progetti di grandi dimensioni

### Qualità dei Dati
- Valida i dati sorgente prima di utilizzarli nelle formule
- Gestisci valori vuoti o null in modo appropriato
- Usa precisione appropriata per i tipi di visualizzazione
- Considera i casi limite nei calcoli

## Casi d'Uso Comuni

1. **Monitoraggio Finanziario**
   - Calcoli di budget
   - Dichiarazioni di profitto/perdita
   - Analisi dei costi
   - Proiezioni di entrate

2. **Gestione Progetti**
   - Percentuali di completamento
   - Utilizzo delle risorse
   - Calcoli delle tempistiche
   - Metriche di prestazione

3. **Controllo Qualità**
   - Punteggi medi
   - Tassi di pass/fallimento
   - Metriche di qualità
   - Monitoraggio della conformità

4. **Intelligenza Aziendale**
   - Calcoli KPI
   - Analisi delle tendenze
   - Metriche comparative
   - Valori della dashboard

## Limitazioni

- Le formule sono solo per aggregazioni di grafico/dashboard, non per calcoli a livello di todo
- Limitate alle sette funzioni di aggregazione supportate (SOMMA, MEDIA, ecc.)
- Nessuna espressione matematica complessa o calcoli campo-a-campo
- Non è possibile fare riferimento a più campi in una singola formula
- I risultati sono visibili solo in grafici e dashboard
- Il campo `logic` è solo per testo di visualizzazione, non per logica di calcolo reale

## Risorse Correlate

- [Campi Numero](/api/5.custom%20fields/number) - Per valori numerici statici
- [Campi Valuta](/api/5.custom%20fields/currency) - Per valori monetari
- [Campi di Riferimento](/api/5.custom%20fields/reference) - Per dati tra progetti
- [Campi di Ricerca](/api/5.custom%20fields/lookup) - Per dati aggregati
- [Panoramica dei Campi Personalizzati](/api/5.custom%20fields/2.list-custom-fields) - Concetti generali
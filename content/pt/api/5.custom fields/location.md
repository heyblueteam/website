---
title: Campo Personalizado de Localização
description: Crie campos de localização para armazenar coordenadas geográficas para registros
---

Os campos personalizados de localização armazenam coordenadas geográficas (latitude e longitude) para registros. Eles suportam armazenamento preciso de coordenadas, consultas geoespaciais e filtragem eficiente baseada em localização.

## Exemplo Básico

Crie um campo de localização simples:

```graphql
mutation CreateLocationField {
  createCustomField(input: {
    name: "Meeting Location"
    type: LOCATION
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de localização com descrição:

```graphql
mutation CreateDetailedLocationField {
  createCustomField(input: {
    name: "Office Location"
    type: LOCATION
    projectId: "proj_123"
    description: "Primary office location coordinates"
  }) {
    id
    name
    type
    description
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de localização |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `LOCATION` |
| `description` | String | Não | Texto de ajuda exibido aos usuários |

## Definindo Valores de Localização

Os campos de localização armazenam coordenadas de latitude e longitude:

```graphql
mutation SetLocationValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    latitude: 40.7128
    longitude: -74.0060
  })
}
```

### Parâmetros de SetTodoCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de localização |
| `latitude` | Float | Não | Coordenada de latitude (-90 a 90) |
| `longitude` | Float | Não | Coordenada de longitude (-180 a 180) |

**Nota**: Embora ambos os parâmetros sejam opcionais no esquema, ambas as coordenadas são necessárias para uma localização válida. Se apenas uma for fornecida, a localização será inválida.

## Validação de Coordenadas

### Faixas Válidas

| Coordenada | Faixa | Descrição |
|------------|-------|-----------|
| Latitude | -90 to 90 | Posição Norte/Sul |
| Longitude | -180 to 180 | Posição Leste/Oeste |

### Coordenadas de Exemplo

| Localização | Latitude | Longitude |
|-------------|----------|-----------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## Criando Registros com Valores de Localização

Ao criar um novo registro com dados de localização:

```graphql
mutation CreateRecordWithLocation {
  createTodo(input: {
    title: "Site Visit"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "location_field_id"
      value: "40.7128,-74.0060"  # Format: "latitude,longitude"
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
      latitude
      longitude
    }
  }
}
```

### Formato de Entrada para Criação

Ao criar registros, os valores de localização usam o formato separado por vírgulas:

| Formato | Exemplo | Descrição |
|---------|---------|-----------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | Formato de coordenadas padrão |
| `"51.5074,-0.1278"` | London coordinates | Sem espaços ao redor da vírgula |
| `"-33.8688,151.2093"` | Sydney coordinates | Valores negativos permitidos |

## Campos de Resposta

### TodoCustomField Resposta

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo personalizado |
| `latitude` | Float | Coordenada de latitude |
| `longitude` | Float | Coordenada de longitude |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

## Limitações Importantes

### Sem Geocodificação Integrada

Os campos de localização armazenam apenas coordenadas - eles **não** incluem:
- Conversão de endereço para coordenadas
- Geocodificação reversa (coordenadas para endereço)
- Validação ou pesquisa de endereço
- Integração com serviços de mapeamento
- Busca de nome de lugar

### Serviços Externos Necessários

Para funcionalidade de endereço, você precisará integrar serviços externos:
- **Google Maps API** para geocodificação
- **OpenStreetMap Nominatim** para geocodificação gratuita
- **MapBox** para mapeamento e geocodificação
- **Here API** para serviços de localização

### Exemplo de Integração

```javascript
// Client-side geocoding example (not part of Blue API)
async function geocodeAddress(address) {
  const response = await fetch(
    `&key=${API_KEY}`
  );
  const data = await response.json();
  
  if (data.results.length > 0) {
    const { lat, lng } = data.results[0].geometry.location;
    
    // Now set the location field in Blue
    await setTodoCustomField({
      todoId: "todo_123",
      customFieldId: "location_field_456",
      latitude: lat,
      longitude: lng
    });
  }
}
```

## Permissões Necessárias

| Ação | Função Necessária |
|------|-------------------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## Respostas de Erro

### Coordenadas Inválidas
```json
{
  "errors": [{
    "message": "Invalid coordinates: latitude must be between -90 and 90",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Longitude Inválida
```json
{
  "errors": [{
    "message": "Invalid coordinates: longitude must be between -180 and 180",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Melhores Práticas

### Coleta de Dados
- Use coordenadas GPS para locais precisos
- Valide coordenadas antes de armazenar
- Considere as necessidades de precisão das coordenadas (6 casas decimais ≈ 10cm de precisão)
- Armazene coordenadas em graus decimais (não em graus/minutos/segundos)

### Experiência do Usuário
- Forneça interfaces de mapa para seleção de coordenadas
- Mostre pré-visualizações de localização ao exibir coordenadas
- Valide coordenadas do lado do cliente antes das chamadas da API
- Considere as implicações de fuso horário para dados de localização

### Desempenho
- Use índices espaciais para consultas eficientes
- Limite a precisão das coordenadas à precisão necessária
- Considere o cache para locais acessados com frequência
- Agrupe atualizações de localização quando possível

## Casos de Uso Comuns

1. **Operações de Campo**
   - Localizações de equipamentos
   - Endereços de chamadas de serviço
   - Locais de inspeção
   - Locais de entrega

2. **Gerenciamento de Eventos**
   - Locais de eventos
   - Locais de reuniões
   - Locais de conferências
   - Locais de workshops

3. **Rastreamento de Ativos**
   - Posições de equipamentos
   - Localizações de instalações
   - Rastreamento de veículos
   - Locais de inventário

4. **Análise Geográfica**
   - Áreas de cobertura de serviço
   - Distribuição de clientes
   - Análise de mercado
   - Gerenciamento de território

## Recursos de Integração

### Com Pesquisas
- Referenciar dados de localização de outros registros
- Encontrar registros por proximidade geográfica
- Agregar dados baseados em localização
- Cruzar coordenadas

### Com Automação
- Acionar ações com base em mudanças de localização
- Criar notificações geofenced
- Atualizar registros relacionados quando as localizações mudam
- Gerar relatórios baseados em localização

### Com Fórmulas
- Calcular distâncias entre locais
- Determinar centros geográficos
- Analisar padrões de localização
- Criar métricas baseadas em localização

## Limitações

- Sem geocodificação integrada ou conversão de endereço
- Sem interface de mapeamento fornecida
- Requer serviços externos para funcionalidade de endereço
- Limitado ao armazenamento de coordenadas apenas
- Sem validação automática de localização além da verificação de faixa

## Recursos Relacionados

- [Visão Geral de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceitos gerais
- [Google Maps API](https://developers.google.com/maps) - Serviços de geocodificação
- [OpenStreetMap Nominatim](https://nominatim.org/) - Geocodificação gratuita
- [MapBox API](https://docs.mapbox.com/) - Serviços de mapeamento e geocodificação
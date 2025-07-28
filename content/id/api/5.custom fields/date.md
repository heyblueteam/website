---
title: Bidang Kustom Tanggal
description: Buat bidang tanggal untuk melacak tanggal tunggal atau rentang tanggal dengan dukungan zona waktu
---

Bidang kustom tanggal memungkinkan Anda untuk menyimpan tanggal tunggal atau rentang tanggal untuk catatan. Mereka mendukung penanganan zona waktu, pemformatan cerdas, dan dapat digunakan untuk melacak tenggat waktu, tanggal acara, atau informasi berbasis waktu lainnya.

## Contoh Dasar

Buat bidang tanggal sederhana:

```graphql
mutation CreateDateField {
  createCustomField(input: {
    name: "Deadline"
    type: DATE
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang tanggal jatuh tempo dengan deskripsi:

```graphql
mutation CreateDueDateField {
  createCustomField(input: {
    name: "Contract Expiration"
    type: DATE
    isDueDate: true
    description: "When the contract expires and needs renewal"
  }) {
    id
    name
    type
    isDueDate
    description
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang tanggal |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `DATE` |
| `isDueDate` | Boolean | Tidak | Apakah bidang ini mewakili tanggal jatuh tempo |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

**Catatan**: Bidang kustom secara otomatis diasosiasikan dengan proyek berdasarkan konteks proyek pengguna saat ini. Tidak ada `projectId` parameter yang diperlukan.

## Mengatur Nilai Tanggal

Bidang tanggal dapat menyimpan baik tanggal tunggal atau rentang tanggal:

### Tanggal Tunggal

```graphql
mutation SetSingleDate {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T10:00:00Z"
    endDate: "2025-01-15T10:00:00Z"
    timezone: "America/New_York"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Rentang Tanggal

```graphql
mutation SetDateRange {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-01T09:00:00Z"
    endDate: "2025-01-31T17:00:00Z"
    timezone: "Europe/London"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Acara Sepanjang Hari

```graphql
mutation SetAllDayEvent {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T00:00:00Z"
    endDate: "2025-01-15T23:59:59Z"
    timezone: "Asia/Tokyo"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Parameter SetTodoCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom tanggal |
| `startDate` | DateTime | Tidak | Tanggal/waktu mulai dalam format ISO 8601 |
| `endDate` | DateTime | Tidak | Tanggal/waktu akhir dalam format ISO 8601 |
| `timezone` | String | Tidak | Pengenal zona waktu (misalnya, "America/New_York") |

**Catatan**: Jika hanya `startDate` yang diberikan, `endDate` secara otomatis akan default ke nilai yang sama.

## Format Tanggal

### Format ISO 8601
Semua tanggal harus disediakan dalam format ISO 8601:
- `2025-01-15T14:30:00Z` - waktu UTC
- `2025-01-15T14:30:00+05:00` - Dengan offset zona waktu
- `2025-01-15T14:30:00.123Z` - Dengan milidetik

### Pengenal Zona Waktu
Gunakan pengenal zona waktu standar:
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

Jika tidak ada zona waktu yang diberikan, sistem akan default ke zona waktu yang terdeteksi oleh pengguna.

## Membuat Catatan dengan Nilai Tanggal

Saat membuat catatan baru dengan nilai tanggal:

```graphql
mutation CreateRecordWithDate {
  createTodo(input: {
    title: "Project Milestone"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "date_field_id"
      value: "2025-02-15"  # Simple date format
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Date values are accessed here
      }
    }
  }
}
```

### Format Input yang Didukung

Saat membuat catatan, tanggal dapat disediakan dalam berbagai format:

| Format | Contoh | Hasil |
|--------|--------|-------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## Bidang Respons

### Respons TodoCustomField

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | ID! | Pengenal unik untuk nilai bidang |
| `uid` | String! | String pengenal unik |
| `customField` | CustomField! | Definisi bidang kustom (berisi nilai tanggal) |
| `todo` | Todo! | Catatan yang dimiliki nilai ini |
| `createdAt` | DateTime! | Kapan nilai dibuat |
| `updatedAt` | DateTime! | Kapan nilai terakhir dimodifikasi |

**Penting**: Nilai tanggal (`startDate`, `endDate`, `timezone`) diakses melalui bidang `customField.value`, bukan langsung pada TodoCustomField.

### Struktur Objek Nilai

Nilai tanggal dikembalikan melalui bidang `customField.value` sebagai objek JSON:

```json
{
  "customField": {
    "value": {
      "startDate": "2025-01-15T10:00:00.000Z",
      "endDate": "2025-01-15T17:00:00.000Z",
      "timezone": "America/New_York"
    }
  }
}
```

**Catatan**: Bidang `value` berada pada tipe `CustomField`, bukan pada `TodoCustomField`.

## Menanyakan Nilai Tanggal

Saat menanyakan catatan dengan bidang kustom tanggal, akses nilai tanggal melalui bidang `customField.value`:

```graphql
query GetRecordWithDateField {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For DATE type, contains { startDate, endDate, timezone }
      }
    }
  }
}
```

Respons akan mencakup nilai tanggal dalam bidang `value`:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Deadline",
          "type": "DATE",
          "value": {
            "startDate": "2025-01-15T10:00:00.000Z",
            "endDate": "2025-01-15T10:00:00.000Z",
            "timezone": "America/New_York"
          }
        }
      }]
    }
  }
}
```

## Kecerdasan Tampilan Tanggal

Sistem secara otomatis memformat tanggal berdasarkan rentang:

| Skenario | Format Tampilan |
|----------|-----------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (tidak ada waktu yang ditampilkan) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**Deteksi sepanjang hari**: Acara dari 00:00 hingga 23:59 secara otomatis terdeteksi sebagai acara sepanjang hari.

## Penanganan Zona Waktu

### Penyimpanan
- Semua tanggal disimpan dalam UTC di database
- Informasi zona waktu disimpan secara terpisah
- Konversi terjadi saat ditampilkan

### Praktik Terbaik
- Selalu berikan zona waktu untuk akurasi
- Gunakan zona waktu yang konsisten dalam proyek
- Pertimbangkan lokasi pengguna untuk tim global

### Zona Waktu Umum

| Wilayah | ID Zona Waktu | Offset UTC |
|---------|---------------|------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## Penyaringan dan Penanyaan

Bidang tanggal mendukung penyaringan kompleks:

```graphql
query FilterByDateRange {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      dateRange: {
        startDate: "2025-01-01T00:00:00Z"
        endDate: "2025-12-31T23:59:59Z"
      }
      operator: EQ  # Returns todos whose dates overlap with this range
    }]
  }) {
    id
    title
  }
}
```

### Memeriksa Bidang Tanggal Kosong

```graphql
query FilterEmptyDates {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      values: null
      operator: IS  # Returns todos with no date set
    }]
  }) {
    id
    title
  }
}
```

### Operator yang Didukung

| Operator | Penggunaan | Deskripsi |
|----------|------------|-----------|
| `EQ` | Dengan dateRange | Tanggal tumpang tindih dengan rentang yang ditentukan (setiap persimpangan) |
| `NE` | Dengan dateRange | Tanggal tidak tumpang tindih dengan rentang |
| `IS` | Dengan `values: null` | Bidang tanggal kosong (startDate atau endDate adalah null) |
| `NOT` | Dengan `values: null` | Bidang tanggal memiliki nilai (kedua tanggal tidak null) |

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|----------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## Respons Kesalahan

### Format Tanggal Tidak Valid
```json
{
  "errors": [{
    "message": "Invalid date format. Use ISO 8601 format",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Bidang Tidak Ditemukan
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


## Batasan

- Tidak ada dukungan tanggal berulang (gunakan otomatisasi untuk acara berulang)
- Tidak dapat mengatur waktu tanpa tanggal
- Tidak ada perhitungan hari kerja bawaan
- Rentang tanggal tidak memvalidasi akhir > awal secara otomatis
- Presisi maksimum adalah hingga detik (tidak ada penyimpanan milidetik)

## Sumber Daya Terkait

- [Ikhtisar Bidang Kustom](/api/custom-fields/list-custom-fields) - Konsep umum bidang kustom
- [API Automasi](/api/automations/index) - Buat otomatisasi berbasis tanggal
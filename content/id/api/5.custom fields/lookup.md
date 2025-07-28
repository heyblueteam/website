---
title: Cari Bidang Kustom
description: Buat bidang pencarian yang secara otomatis menarik data dari catatan yang dirujuk
---

Bidang kustom pencarian secara otomatis menarik data dari catatan yang dirujuk oleh [Bidang Referensi](/api/custom-fields/reference), menampilkan informasi dari catatan yang terhubung tanpa penyalinan manual. Mereka diperbarui secara otomatis ketika data yang dirujuk berubah.

## Contoh Dasar

Buat bidang pencarian untuk menampilkan tag dari catatan yang dirujuk:

```graphql
mutation CreateLookupField {
  createCustomField(input: {
    name: "Related Todo Tags"
    type: LOOKUP
    lookupOption: {
      referenceId: "reference_field_id"
      lookupType: TODO_TAG
    }
    description: "Tags from related todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Contoh Lanjutan

Buat bidang pencarian untuk mengekstrak nilai bidang kustom dari catatan yang dirujuk:

```graphql
mutation CreateCustomFieldLookup {
  createCustomField(input: {
    name: "Referenced Budget Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "project_reference_field_id"
      lookupId: "budget_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
    description: "Budget values from referenced todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang pencarian |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ Ya | Konfigurasi pencarian |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

## Konfigurasi Pencarian

### CustomFieldLookupOptionInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `referenceId` | String! | ✅ Ya | ID dari bidang referensi untuk menarik data |
| `lookupId` | String | Tidak | ID dari bidang kustom tertentu untuk dicari (diperlukan untuk tipe TODO_CUSTOM_FIELD) |
| `lookupType` | CustomFieldLookupType! | ✅ Ya | Tipe data yang akan diekstrak dari catatan yang dirujuk |

## Tipe Pencarian

### Nilai CustomFieldLookupType

| Tipe | Deskripsi | Mengembalikan |
|------|-----------|---------------|
| `TODO_DUE_DATE` | Tanggal jatuh tempo dari todo yang dirujuk | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | Tanggal pembuatan dari todo yang dirujuk | Array of creation timestamps |
| `TODO_UPDATED_AT` | Tanggal terakhir diperbarui dari todo yang dirujuk | Array of update timestamps |
| `TODO_TAG` | Tag dari todo yang dirujuk | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | Penugasan dari todo yang dirujuk | Array of user objects |
| `TODO_DESCRIPTION` | Deskripsi dari todo yang dirujuk | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | Nama daftar todo dari todo yang dirujuk | Array of list titles |
| `TODO_CUSTOM_FIELD` | Nilai bidang kustom dari todo yang dirujuk | Array of values based on the field type |

## Bidang Respons

### Respons CustomField (untuk bidang pencarian)

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk bidang |
| `name` | String! | Nama tampilan dari bidang pencarian |
| `type` | CustomFieldType! | Akan menjadi `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | Konfigurasi dan hasil pencarian |
| `createdAt` | DateTime! | Ketika bidang dibuat |
| `updatedAt` | DateTime! | Ketika bidang terakhir diperbarui |

### Struktur CustomFieldLookupOption

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `lookupType` | CustomFieldLookupType! | Tipe pencarian yang dilakukan |
| `lookupResult` | JSON | Data yang diekstrak dari catatan yang dirujuk |
| `reference` | CustomField | Bidang referensi yang digunakan sebagai sumber |
| `lookup` | CustomField | Bidang spesifik yang dicari (untuk TODO_CUSTOM_FIELD) |
| `parentCustomField` | CustomField | Bidang pencarian induk |
| `parentLookup` | CustomField | Pencarian induk dalam rantai (untuk pencarian bersarang) |

## Cara Kerja Pencarian

1. **Ekstraksi Data**: Pencarian mengekstrak data spesifik dari semua catatan yang terhubung melalui bidang referensi
2. **Pembaruan Otomatis**: Ketika catatan yang dirujuk berubah, nilai pencarian diperbarui secara otomatis
3. **Hanya Baca**: Bidang pencarian tidak dapat diedit langsung - mereka selalu mencerminkan data yang dirujuk saat ini
4. **Tanpa Perhitungan**: Pencarian mengekstrak dan menampilkan data apa adanya tanpa agregasi atau perhitungan

## Pencarian TODO_CUSTOM_FIELD

Saat menggunakan `TODO_CUSTOM_FIELD` tipe, Anda harus menentukan bidang kustom mana yang akan diekstrak menggunakan parameter `lookupId`:

```graphql
mutation CreateCustomFieldValueLookup {
  createCustomField(input: {
    name: "Project Status Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "linked_projects_reference_field"
      lookupId: "status_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
  }) {
    id
  }
}
```

Ini mengekstrak nilai dari bidang kustom yang ditentukan dari semua catatan yang dirujuk.

## Mengquery Data Pencarian

```graphql
query GetLookupValues {
  todo(id: "todo_123") {
    customFields {
      id
      customField {
        name
        type
        customFieldLookupOption {
          lookupType
          lookupResult
          reference {
            id
            name
          }
          lookup {
            id
            name
            type
          }
        }
      }
    }
  }
}
```

## Contoh Hasil Pencarian

### Hasil Pencarian Tag
```json
{
  "lookupResult": [
    {
      "id": "tag_123",
      "title": "urgent",
      "color": "#ff0000"
    },
    {
      "id": "tag_456",
      "title": "development",
      "color": "#00ff00"
    }
  ]
}
```

### Hasil Pencarian Penugasan
```json
{
  "lookupResult": [
    {
      "id": "user_123",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### Hasil Pencarian Bidang Kustom
Hasil bervariasi berdasarkan tipe bidang kustom yang dicari. Misalnya, pencarian bidang mata uang mungkin mengembalikan:
```json
{
  "lookupResult": [
    {
      "value": 1000,
      "currency": "USD"
    },
    {
      "value": 2500,
      "currency": "EUR"
    }
  ]
}
```

## Izin yang Diperlukan

| Aksi | Izin yang Diperlukan |
|------|---------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**Penting**: Pengguna harus memiliki izin tampilan pada proyek saat ini dan proyek yang dirujuk untuk melihat hasil pencarian.

## Respons Kesalahan

### Bidang Referensi Tidak Valid
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

### Pencarian Sirkuler Terdeteksi
```json
{
  "errors": [{
    "message": "Circular lookup detected",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### ID Pencarian Hilang untuk TODO_CUSTOM_FIELD
```json
{
  "errors": [{
    "message": "lookupId is required when lookupType is TODO_CUSTOM_FIELD",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

## Praktik Terbaik

1. **Penamaan yang Jelas**: Gunakan nama deskriptif yang menunjukkan data apa yang dicari
2. **Tipe yang Sesuai**: Pilih tipe pencarian yang sesuai dengan kebutuhan data Anda
3. **Kinerja**: Pencarian memproses semua catatan yang dirujuk, jadi perhatikan bidang referensi dengan banyak tautan
4. **Izin**: Pastikan pengguna memiliki akses ke proyek yang dirujuk agar pencarian dapat berfungsi

## Kasus Penggunaan Umum

### Visibilitas Lintas Proyek
Tampilkan tag, penugasan, atau status dari proyek terkait tanpa sinkronisasi manual.

### Pelacakan Ketergantungan
Tampilkan tanggal jatuh tempo atau status penyelesaian tugas yang bergantung pada pekerjaan saat ini.

### Ikhtisar Sumber Daya
Tampilkan semua anggota tim yang ditugaskan ke tugas yang dirujuk untuk perencanaan sumber daya.

### Agregasi Status
Kumpulkan semua status unik dari tugas terkait untuk melihat kesehatan proyek dengan cepat.

## Batasan

- Bidang pencarian bersifat hanya baca dan tidak dapat diedit langsung
- Tidak ada fungsi agregasi (SUM, COUNT, AVG) - pencarian hanya mengekstrak data
- Tidak ada opsi penyaringan - semua catatan yang dirujuk disertakan
- Rantai pencarian sirkuler dicegah untuk menghindari loop tak berujung
- Hasil mencerminkan data saat ini dan diperbarui secara otomatis

## Sumber Daya Terkait

- [Bidang Referensi](/api/custom-fields/reference) - Buat tautan ke catatan untuk sumber pencarian
- [Nilai Bidang Kustom](/api/custom-fields/custom-field-values) - Atur nilai pada bidang kustom yang dapat diedit
- [Daftar Bidang Kustom](/api/custom-fields/list-custom-fields) - Kuery semua bidang kustom dalam proyek
---
title: Bidang Kustom Negara
description: Buat bidang pemilihan negara dengan validasi kode negara ISO
---

Bidang kustom negara memungkinkan Anda untuk menyimpan dan mengelola informasi negara untuk catatan. Bidang ini mendukung nama negara dan kode negara ISO Alpha-2.

**Penting**: Perilaku validasi dan konversi negara berbeda secara signifikan antara mutasi:
- **createTodo**: Secara otomatis memvalidasi dan mengonversi nama negara menjadi kode ISO
- **setTodoCustomField**: Menerima nilai apa pun tanpa validasi

## Contoh Dasar

Buat bidang negara sederhana:

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang negara dengan deskripsi:

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-------------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang negara |
| `type` | CustomFieldType! | ✅ Ya | Harus `COUNTRY` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

**Catatan**: `projectId` tidak diteruskan dalam input tetapi ditentukan oleh konteks GraphQL (biasanya dari header permintaan atau otentikasi).

## Mengatur Nilai Negara

Bidang negara menyimpan data dalam dua bidang basis data:
- **`countryCodes`**: Menyimpan kode negara ISO Alpha-2 sebagai string yang dipisahkan koma dalam basis data (dikembalikan sebagai array melalui API)
- **`text`**: Menyimpan teks tampilan atau nama negara sebagai string

### Memahami Parameter

Mutasi `setTodoCustomField` menerima dua parameter opsional untuk bidang negara:

| Parameter | Tipe | Diperlukan | Deskripsi | Apa yang dilakukannya |
|-----------|------|------------|-------------|--------------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui | - |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom negara | - |
| `countryCodes` | [String!] | Tidak | Array dari kode negara ISO Alpha-2 | Stored in the `countryCodes` field |
| `text` | String | Tidak | Teks tampilan atau nama negara | Stored in the `text` field |

**Penting**: 
- Dalam `setTodoCustomField`: Kedua parameter bersifat opsional dan disimpan secara independen
- Dalam `createTodo`: Sistem secara otomatis mengatur kedua bidang berdasarkan input Anda (Anda tidak dapat mengontrolnya secara independen)

### Opsi 1: Menggunakan Hanya Kode Negara

Simpan kode ISO yang telah divalidasi tanpa teks tampilan:

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

Hasil: `countryCodes` = `["US"]`, `text` = `null`

### Opsi 2: Menggunakan Hanya Teks

Simpan teks tampilan tanpa kode yang telah divalidasi:

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

Hasil: `countryCodes` = `null`, `text` = `"United States"`

**Catatan**: Saat menggunakan `setTodoCustomField`, tidak ada validasi yang terjadi terlepas dari parameter mana yang Anda gunakan. Nilai disimpan persis seperti yang diberikan.

### Opsi 3: Menggunakan Keduanya (Direkomendasikan)

Simpan baik kode yang telah divalidasi dan teks tampilan:

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

Hasil: `countryCodes` = `["US"]`, `text` = `"United States"`

### Beberapa Negara

Simpan beberapa negara menggunakan array:

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## Membuat Catatan dengan Nilai Negara

Saat membuat catatan, mutasi `createTodo` **secara otomatis memvalidasi dan mengonversi** nilai negara. Ini adalah satu-satunya mutasi yang melakukan validasi negara:

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      countryCodes
    }
  }
}
```

### Format Input yang Diterima

| Tipe Input | Contoh | Hasil |
|------------|---------|---------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **Tidak didukung** - diperlakukan sebagai nilai tidak valid tunggal |
| Mixed format | `"United States, CA"` | **Tidak didukung** - diperlakukan sebagai nilai tidak valid tunggal |

## Bidang Respons

### Respons TodoCustomField

| Bidang | Tipe | Deskripsi |
|-------|------|-------------|
| `id` | String! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang kustom |
| `text` | String | Teks tampilan (nama negara) |
| `countryCodes` | [String!] | Array dari kode negara ISO Alpha-2 |
| `todo` | Todo! | Catatan yang menjadi milik nilai ini |
| `createdAt` | DateTime! | Saat nilai dibuat |
| `updatedAt` | DateTime! | Saat nilai terakhir dimodifikasi |

## Standar Negara

Blue menggunakan standar **ISO 3166-1 Alpha-2** untuk kode negara:

- Kode negara dua huruf (misalnya, US, GB, FR, DE)
- Validasi menggunakan pustaka `i18n-iso-countries` **hanya terjadi di createTodo**
- Mendukung semua negara yang diakui secara resmi

### Contoh Kode Negara

| Negara | Kode ISO |
|---------|----------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

Untuk daftar resmi lengkap kode negara ISO 3166-1 alpha-2, kunjungi [Platform Penelusuran Online ISO](https://www.iso.org/obp/ui/#search/code/).

## Validasi

**Validasi hanya terjadi di mutasi `createTodo`**:

1. **Kode ISO Valid**: Menerima kode ISO Alpha-2 yang valid
2. **Nama Negara**: Secara otomatis mengonversi nama negara yang dikenali menjadi kode
3. **Input Tidak Valid**: Menghasilkan `CustomFieldValueParseError` untuk nilai yang tidak dikenali

**Catatan**: Mutasi `setTodoCustomField` tidak melakukan validasi dan menerima nilai string apa pun.

### Contoh Kesalahan

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Fitur Integrasi

### Bidang Lookup
Bidang negara dapat dirujuk oleh bidang kustom LOOKUP, memungkinkan Anda untuk menarik data negara dari catatan terkait.

### Automasi
Gunakan nilai negara dalam kondisi automasi:
- Filter tindakan berdasarkan negara tertentu
- Kirim notifikasi berdasarkan negara
- Rute tugas berdasarkan wilayah geografis

### Formulir
Bidang negara dalam formulir secara otomatis memvalidasi input pengguna dan mengonversi nama negara menjadi kode.

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|-------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## Respons Kesalahan

### Nilai Negara Tidak Valid
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Ketidakcocokan Tipe Bidang
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## Praktik Terbaik

### Penanganan Input
- Gunakan `createTodo` untuk validasi dan konversi otomatis
- Gunakan `setTodoCustomField` dengan hati-hati karena ini melewati validasi
- Pertimbangkan untuk memvalidasi input di aplikasi Anda sebelum menggunakan `setTodoCustomField`
- Tampilkan nama negara lengkap di UI untuk kejelasan

### Kualitas Data
- Validasi input negara di titik masuk
- Gunakan format yang konsisten di seluruh sistem Anda
- Pertimbangkan pengelompokan regional untuk pelaporan

### Beberapa Negara
- Gunakan dukungan array dalam `setTodoCustomField` untuk beberapa negara
- Beberapa negara dalam `createTodo` **tidak didukung** melalui bidang nilai
- Simpan kode negara sebagai array dalam `setTodoCustomField` untuk penanganan yang tepat

## Kasus Penggunaan Umum

1. **Manajemen Pelanggan**
   - Lokasi kantor pusat pelanggan
   - Tujuan pengiriman
   - Yurisdiksi pajak

2. **Pelacakan Proyek**
   - Lokasi proyek
   - Lokasi anggota tim
   - Target pasar

3. **Kepatuhan & Hukum**
   - Yurisdiksi regulasi
   - Persyaratan tempat tinggal data
   - Kontrol ekspor

4. **Penjualan & Pemasaran**
   - Penugasan wilayah
   - Segmentasi pasar
   - Penargetan kampanye

## Batasan

- Hanya mendukung kode ISO 3166-1 Alpha-2 (kode 2 huruf)
- Tidak ada dukungan bawaan untuk subdivisi negara (negara bagian/provinsi)
- Tidak ada ikon bendera negara otomatis (hanya berbasis teks)
- Tidak dapat memvalidasi kode negara historis
- Tidak ada pengelompokan wilayah atau benua bawaan
- **Validasi hanya berfungsi dalam `createTodo`, tidak dalam `setTodoCustomField`**
- **Beberapa negara tidak didukung dalam bidang nilai `createTodo`**
- **Kode negara disimpan sebagai string yang dipisahkan koma, bukan array yang sebenarnya**

## Sumber Daya Terkait

- [Ikhtisar Bidang Kustom](/custom-fields/list-custom-fields) - Konsep umum bidang kustom
- [Bidang Lookup](/api/custom-fields/lookup) - Referensi data negara dari catatan lain
- [API Formulir](/api/forms) - Sertakan bidang negara dalam formulir kustom
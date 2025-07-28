---
title: Bidang Kustom Teks Satu Baris
description: Buat bidang teks satu baris untuk nilai teks pendek seperti nama, judul, dan label
---

Bidang kustom teks satu baris memungkinkan Anda menyimpan nilai teks pendek yang dimaksudkan untuk input satu baris. Mereka ideal untuk nama, judul, label, atau data teks apa pun yang harus ditampilkan dalam satu baris.

## Contoh Dasar

Buat bidang teks satu baris yang sederhana:

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang teks satu baris dengan deskripsi:

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
  }) {
    id
    name
    type
    description
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang teks |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `TEXT_SINGLE` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

**Catatan**: Konteks proyek ditentukan secara otomatis dari header otentikasi Anda. Tidak ada parameter `projectId` yang diperlukan.

## Mengatur Nilai Teks

Untuk mengatur atau memperbarui nilai teks satu baris pada sebuah catatan:

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom teks |
| `text` | String | Tidak | Konten teks satu baris untuk disimpan |

## Membuat Catatan dengan Nilai Teks

Saat membuat catatan baru dengan nilai teks satu baris:

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## Bidang Respons

### TodoCustomField Respons

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | ID! | Pengenal unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang kustom (mengandung nilai teks) |
| `todo` | Todo! | Catatan yang nilai ini miliki |
| `createdAt` | DateTime! | Ketika nilai dibuat |
| `updatedAt` | DateTime! | Ketika nilai terakhir dimodifikasi |

**Penting**: Nilai teks diakses melalui bidang `customField.value.text`, bukan langsung pada TodoCustomField.

## Mengquery Nilai Teks

Saat mengquery catatan dengan bidang kustom teks, akses teks melalui jalur `customField.value.text`:

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

Respons akan mencakup teks dalam struktur bersarang:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## Validasi Teks

### Validasi Formulir
Saat bidang teks satu baris digunakan dalam formulir:
- Spasi di awal dan akhir secara otomatis dihapus
- Validasi diperlukan diterapkan jika bidang ditandai sebagai diperlukan
- Tidak ada validasi format khusus yang diterapkan

### Aturan Validasi
- Menerima konten string apa pun termasuk pemisah baris (meskipun tidak disarankan)
- Tidak ada batas panjang karakter (hingga batas database)
- Mendukung karakter Unicode dan simbol khusus
- Pemisah baris dipertahankan tetapi tidak dimaksudkan untuk jenis bidang ini

### Contoh Teks Tipikal
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## Catatan Penting

### Kapasitas Penyimpanan
- Disimpan menggunakan tipe `MediumText` MySQL
- Mendukung hingga 16MB konten teks
- Penyimpanan identik dengan bidang teks multi-baris
- Pengkodean UTF-8 untuk karakter internasional

### API Langsung vs Formulir
- **Formulir**: Pemangkasan spasi otomatis dan validasi diperlukan
- **API Langsung**: Teks disimpan persis seperti yang diberikan
- **Rekomendasi**: Gunakan formulir untuk input pengguna untuk memastikan format yang konsisten

### TEXT_SINGLE vs TEXT_MULTI
- **TEXT_SINGLE**: Input teks satu baris, ideal untuk nilai pendek
- **TEXT_MULTI**: Input area teks multi-baris, ideal untuk konten yang lebih panjang
- **Backend**: Keduanya menggunakan penyimpanan dan validasi yang identik
- **Frontend**: Komponen UI yang berbeda untuk entri data
- **Niat**: TEXT_SINGLE secara semantik dimaksudkan untuk nilai satu baris

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|----------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

## Respons Kesalahan

### Validasi Bidang Diperlukan (Hanya Formulir)
```json
{
  "errors": [{
    "message": "This field is required",
    "extensions": {
      "code": "VALIDATION_ERROR"
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

## Praktik Terbaik

### Pedoman Konten
- Jaga teks tetap ringkas dan sesuai untuk satu baris
- Hindari pemisah baris untuk tampilan satu baris yang dimaksudkan
- Gunakan format yang konsisten untuk tipe data yang serupa
- Pertimbangkan batas karakter berdasarkan kebutuhan UI Anda

### Entri Data
- Berikan deskripsi bidang yang jelas untuk membimbing pengguna
- Gunakan formulir untuk input pengguna untuk memastikan validasi
- Validasi format konten dalam aplikasi Anda jika diperlukan
- Pertimbangkan untuk menggunakan dropdown untuk nilai yang distandarisasi

### Pertimbangan Kinerja
- Bidang teks satu baris ringan dan berkinerja baik
- Pertimbangkan pengindeksan untuk bidang yang sering dicari
- Gunakan lebar tampilan yang sesuai di UI Anda
- Pantau panjang konten untuk tujuan tampilan

## Penyaringan dan Pencarian

### Pencarian Mengandung
Bidang teks satu baris mendukung pencarian substring:

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### Kemampuan Pencarian
- Pencocokan substring tanpa memperhatikan huruf besar/kecil
- Mendukung pencocokan kata sebagian
- Pencocokan nilai yang tepat
- Tidak ada pencarian teks penuh atau peringkat

## Kasus Penggunaan Umum

1. **Pengenal dan Kode**
   - SKU produk
   - Nomor pesanan
   - Kode referensi
   - Nomor versi

2. **Nama dan Judul**
   - Nama klien
   - Judul proyek
   - Nama produk
   - Label kategori

3. **Deskripsi Pendek**
   - Ringkasan singkat
   - Label status
   - Indikator prioritas
   - Tag klasifikasi

4. **Referensi Eksternal**
   - Nomor tiket
   - Referensi faktur
   - ID sistem eksternal
   - Nomor dokumen

## Fitur Integrasi

### Dengan Pencarian
- Referensi data teks dari catatan lain
- Temukan catatan berdasarkan konten teks
- Tampilkan informasi teks terkait
- Agregasi nilai teks dari beberapa sumber

### Dengan Formulir
- Pemangkasan spasi otomatis
- Validasi bidang yang diperlukan
- UI input teks satu baris
- Tampilan batas karakter (jika dikonfigurasi)

### Dengan Impor/Ekspor
- Pemetaan kolom CSV langsung
- Penugasan nilai teks otomatis
- Dukungan impor data massal
- Ekspor ke format spreadsheet

## Batasan

### Pembatasan Automasi
- Tidak tersedia langsung sebagai bidang pemicu automasi
- Tidak dapat digunakan dalam pembaruan bidang automasi
- Dapat dirujuk dalam kondisi automasi
- Tersedia dalam template email dan webhook

### Batasan Umum
- Tidak ada format atau gaya teks bawaan
- Tidak ada validasi otomatis di luar bidang yang diperlukan
- Tidak ada penegakan keunikan bawaan
- Tidak ada kompresi konten untuk teks yang sangat besar
- Tidak ada versi atau pelacakan perubahan
- Kemampuan pencarian terbatas (tidak ada pencarian teks penuh)

## Sumber Daya Terkait

- [Bidang Teks Multi-Baris](/api/custom-fields/text-multi) - Untuk konten teks yang lebih panjang
- [Bidang Email](/api/custom-fields/email) - Untuk alamat email
- [Bidang URL](/api/custom-fields/url) - Untuk alamat situs web
- [Bidang ID Unik](/api/custom-fields/unique-id) - Untuk pengenal yang dihasilkan secara otomatis
- [Ikhtisar Bidang Kustom](/api/custom-fields/list-custom-fields) - Konsep umum
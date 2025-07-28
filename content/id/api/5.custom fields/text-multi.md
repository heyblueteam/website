---
title: Bidang Kustom Teks Multi-Lini
description: Buat bidang teks multi-lini untuk konten yang lebih panjang seperti deskripsi, catatan, dan komentar
---

Bidang kustom teks multi-lini memungkinkan Anda menyimpan konten teks yang lebih panjang dengan pemisah baris dan format. Mereka ideal untuk deskripsi, catatan, komentar, atau data teks apa pun yang memerlukan beberapa baris.

## Contoh Dasar

Buat bidang teks multi-lini sederhana:

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang teks multi-lini dengan deskripsi:

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
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
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `TEXT_MULTI` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

**Catatan:** `projectId` diteruskan sebagai argumen terpisah ke mutasi, bukan sebagai bagian dari objek input. Sebagai alternatif, konteks proyek dapat ditentukan dari header `X-Bloo-Project-ID` dalam permintaan GraphQL Anda.

## Mengatur Nilai Teks

Untuk mengatur atau memperbarui nilai teks multi-lini pada sebuah catatan:

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom teks |
| `text` | String | Tidak | Konten teks multi-lini yang akan disimpan |

## Membuat Catatan dengan Nilai Teks

Saat membuat catatan baru dengan nilai teks multi-lini:

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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
    }
  }
}
```

## Bidang Respons

### TodoCustomField Respons

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang kustom |
| `text` | String | Konten teks multi-lini yang disimpan |
| `todo` | Todo! | Catatan yang nilai ini miliki |
| `createdAt` | DateTime! | Waktu nilai dibuat |
| `updatedAt` | DateTime! | Waktu nilai terakhir dimodifikasi |

## Validasi Teks

### Validasi Form
Saat bidang teks multi-lini digunakan dalam formulir:
- Spasi di awal dan akhir secara otomatis dipangkas
- Validasi yang diperlukan diterapkan jika bidang ditandai sebagai diperlukan
- Tidak ada validasi format tertentu yang diterapkan

### Aturan Validasi
- Menerima konten string apa pun termasuk pemisah baris
- Tidak ada batasan panjang karakter (hingga batas database)
- Mendukung karakter Unicode dan simbol khusus
- Pemisah baris dipertahankan dalam penyimpanan

### Contoh Teks yang Valid
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## Catatan Penting

### Kapasitas Penyimpanan
- Disimpan menggunakan tipe MySQL `MediumText`
- Mendukung hingga 16MB konten teks
- Pemisah baris dan format dipertahankan
- Pengkodean UTF-8 untuk karakter internasional

### API Langsung vs Formulir
- **Formulir**: Pemangkasan spasi otomatis dan validasi yang diperlukan
- **API Langsung**: Teks disimpan persis seperti yang diberikan
- **Rekomendasi**: Gunakan formulir untuk input pengguna untuk memastikan format yang konsisten

### TEXT_MULTI vs TEXT_SINGLE
- **TEXT_MULTI**: Input area teks multi-lini, ideal untuk konten yang lebih panjang
- **TEXT_SINGLE**: Input teks satu-lini, ideal untuk nilai pendek
- **Backend**: Kedua jenis identik - bidang penyimpanan, validasi, dan pemrosesan yang sama
- **Frontend**: Komponen UI yang berbeda untuk entri data (area teks vs bidang input)
- **Penting**: Perbedaan antara TEXT_MULTI dan TEXT_SINGLE ada semata-mata untuk tujuan UI

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|---------------------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Praktik Terbaik

### Organisasi Konten
- Gunakan format yang konsisten untuk konten terstruktur
- Pertimbangkan menggunakan sintaks mirip markdown untuk keterbacaan
- Pecah konten panjang menjadi bagian logis
- Gunakan pemisah baris untuk meningkatkan keterbacaan

### Entri Data
- Berikan deskripsi bidang yang jelas untuk membimbing pengguna
- Gunakan formulir untuk input pengguna untuk memastikan validasi
- Pertimbangkan batasan karakter berdasarkan kasus penggunaan Anda
- Validasi format konten dalam aplikasi Anda jika diperlukan

### Pertimbangan Kinerja
- Konten teks yang sangat panjang dapat mempengaruhi kinerja kueri
- Pertimbangkan paginasi untuk menampilkan bidang teks besar
- Pertimbangan indeks untuk fungsionalitas pencarian
- Pantau penggunaan penyimpanan untuk bidang dengan konten besar

## Penyaringan dan Pencarian

### Pencarian Mengandung
Bidang teks multi-lini mendukung pencarian substring melalui filter bidang kustom:

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### Kemampuan Pencarian
- Pencocokan substring dalam bidang teks menggunakan operator `CONTAINS`
- Pencarian tidak sensitif terhadap huruf besar menggunakan operator `NCONTAINS`
- Pencocokan tepat menggunakan operator `IS`
- Pencocokan negatif menggunakan operator `NOT`
- Pencarian di seluruh baris teks
- Mendukung pencocokan kata sebagian

## Kasus Penggunaan Umum

1. **Manajemen Proyek**
   - Deskripsi tugas
   - Persyaratan proyek
   - Catatan rapat
   - Pembaruan status

2. **Dukungan Pelanggan**
   - Deskripsi masalah
   - Catatan penyelesaian
   - Umpan balik pelanggan
   - Log komunikasi

3. **Manajemen Konten**
   - Konten artikel
   - Deskripsi produk
   - Komentar pengguna
   - Detail ulasan

4. **Dokumentasi**
   - Deskripsi proses
   - Instruksi
   - Pedoman
   - Materi referensi

## Fitur Integrasi

### Dengan Automasi
- Memicu tindakan saat konten teks berubah
- Mengekstrak kata kunci dari konten teks
- Membuat ringkasan atau notifikasi
- Memproses konten teks dengan layanan eksternal

### Dengan Pencarian
- Referensi data teks dari catatan lain
- Mengagregasi konten teks dari beberapa sumber
- Menemukan catatan berdasarkan konten teks
- Menampilkan informasi teks terkait

### Dengan Formulir
- Pemangkasan spasi otomatis
- Validasi bidang yang diperlukan
- UI area teks multi-lini
- Tampilan jumlah karakter (jika dikonfigurasi)

## Batasan

- Tidak ada pemformatan teks bawaan atau pengeditan teks kaya
- Tidak ada deteksi atau konversi tautan otomatis
- Tidak ada pemeriksaan ejaan atau validasi tata bahasa
- Tidak ada analisis atau pemrosesan teks bawaan
- Tidak ada versi atau pelacakan perubahan
- Kemampuan pencarian terbatas (tidak ada pencarian teks penuh)
- Tidak ada kompresi konten untuk teks yang sangat besar

## Sumber Daya Terkait

- [Bidang Teks Satu-Lini](/api/custom-fields/text-single) - Untuk nilai teks pendek
- [Bidang Email](/api/custom-fields/email) - Untuk alamat email
- [Bidang URL](/api/custom-fields/url) - Untuk alamat situs web
- [Ikhtisar Bidang Kustom](/api/custom-fields/2.list-custom-fields) - Konsep umum
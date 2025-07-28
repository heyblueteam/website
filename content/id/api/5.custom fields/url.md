---
title: Bidang Kustom URL
description: Buat bidang URL untuk menyimpan alamat situs web dan tautan
---

Bidang kustom URL memungkinkan Anda untuk menyimpan alamat situs web dan tautan dalam catatan Anda. Mereka ideal untuk melacak situs web proyek, tautan referensi, URL dokumentasi, atau sumber daya berbasis web lainnya yang terkait dengan pekerjaan Anda.

## Contoh Dasar

Buat bidang URL sederhana:

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang URL dengan deskripsi:

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
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
| `name` | String! | ✅ Ya | Nama tampilan dari bidang URL |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `URL` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

**Catatan:** `projectId` diteruskan sebagai argumen terpisah ke mutasi, bukan sebagai bagian dari objek input.

## Mengatur Nilai URL

Untuk mengatur atau memperbarui nilai URL pada catatan:

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom URL |
| `text` | String! | ✅ Ya | Alamat URL yang akan disimpan |

## Membuat Catatan dengan Nilai URL

Saat membuat catatan baru dengan nilai URL:

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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
| `text` | String | Alamat URL yang disimpan |
| `todo` | Todo! | Catatan yang dimiliki nilai ini |
| `createdAt` | DateTime! | Waktu nilai dibuat |
| `updatedAt` | DateTime! | Waktu nilai terakhir dimodifikasi |

## Validasi URL

### Implementasi Saat Ini
- **API Langsung**: Tidak ada validasi format URL yang diterapkan saat ini
- **Formulir**: Validasi URL direncanakan tetapi saat ini tidak aktif
- **Penyimpanan**: Setiap nilai string dapat disimpan dalam bidang URL

### Validasi yang Direncanakan
Versi mendatang akan mencakup:
- Validasi protokol HTTP/HTTPS
- Pemeriksaan format URL yang valid
- Validasi nama domain
- Penambahan prefiks protokol otomatis

### Format URL yang Direkomendasikan
Meskipun saat ini tidak diterapkan, gunakan format standar ini:

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## Catatan Penting

### Format Penyimpanan
- URL disimpan sebagai teks biasa tanpa modifikasi
- Tidak ada penambahan protokol otomatis (http://, https://)
- Sensitivitas huruf besar dan kecil dipertahankan seperti yang dimasukkan
- Tidak ada pengkodean/dekodean URL yang dilakukan

### API Langsung vs Formulir
- **Formulir**: Validasi URL yang direncanakan (tidak aktif saat ini)
- **API Langsung**: Tidak ada validasi - teks apa pun dapat disimpan
- **Rekomendasi**: Validasi URL dalam aplikasi Anda sebelum menyimpan

### URL vs Bidang Teks
- **URL**: Secara semantik ditujukan untuk alamat web
- **TEXT_SINGLE**: Teks umum satu baris
- **Backend**: Saat ini penyimpanan dan validasi identik
- **Frontend**: Komponen UI yang berbeda untuk entri data

## Izin yang Diperlukan

Operasi bidang kustom menggunakan izin berbasis peran:

| Tindakan | Peran yang Diperlukan |
|----------|-----------------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**Catatan:** Izin diperiksa berdasarkan peran pengguna dalam proyek, bukan konstanta izin spesifik.

## Respons Kesalahan

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

## Praktik Terbaik

### Standar Format URL
- Selalu sertakan protokol (http:// atau https://)
- Gunakan HTTPS jika memungkinkan untuk keamanan
- Uji URL sebelum menyimpan untuk memastikan dapat diakses
- Pertimbangkan menggunakan URL pendek untuk tujuan tampilan

### Kualitas Data
- Validasi URL dalam aplikasi Anda sebelum menyimpan
- Periksa kesalahan umum (protokol yang hilang, domain yang salah)
- Standarisasi format URL di seluruh organisasi Anda
- Pertimbangkan aksesibilitas dan ketersediaan URL

### Pertimbangan Keamanan
- Berhati-hatilah dengan URL yang diberikan pengguna
- Validasi domain jika membatasi ke situs tertentu
- Pertimbangkan pemindaian URL untuk konten berbahaya
- Gunakan URL HTTPS saat menangani data sensitif

## Penyaringan dan Pencarian

### Pencarian Mengandung
Bidang URL mendukung pencarian substring:

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### Kemampuan Pencarian
- Pencocokan substring yang tidak sensitif huruf besar/kecil
- Pencocokan domain sebagian
- Pencarian jalur dan parameter
- Tidak ada penyaringan spesifik protokol

## Kasus Penggunaan Umum

1. **Manajemen Proyek**
   - Situs web proyek
   - Tautan dokumentasi
   - URL repositori
   - Situs demo

2. **Manajemen Konten**
   - Materi referensi
   - Tautan sumber
   - Sumber daya media
   - Artikel eksternal

3. **Dukungan Pelanggan**
   - Situs web pelanggan
   - Dokumentasi dukungan
   - Artikel basis pengetahuan
   - Tutorial video

4. **Penjualan & Pemasaran**
   - Situs web perusahaan
   - Halaman produk
   - Materi pemasaran
   - Profil media sosial

## Fitur Integrasi

### Dengan Pencarian
- Tautan URL dari catatan lain
- Temukan catatan berdasarkan domain atau pola URL
- Tampilkan sumber daya web terkait
- Agregasi tautan dari beberapa sumber

### Dengan Formulir
- Komponen input khusus URL
- Validasi yang direncanakan untuk format URL yang benar
- Kemampuan pratayang tautan (frontend)
- Tampilan URL yang dapat diklik

### Dengan Pelaporan
- Lacak penggunaan dan pola URL
- Pantau tautan yang rusak atau tidak dapat diakses
- Kategorikan berdasarkan domain atau protokol
- Ekspor daftar URL untuk analisis

## Batasan

### Batasan Saat Ini
- Tidak ada validasi format URL yang aktif
- Tidak ada penambahan protokol otomatis
- Tidak ada verifikasi tautan atau pemeriksaan aksesibilitas
- Tidak ada pemendekan atau perluasan URL
- Tidak ada pembuatan favicon atau pratayang

### Pembatasan Otomasi
- Tidak tersedia sebagai bidang pemicu otomatisasi
- Tidak dapat digunakan dalam pembaruan bidang otomatisasi
- Dapat dirujuk dalam kondisi otomatisasi
- Tersedia dalam template email dan webhook

### Kendala Umum
- Tidak ada fungsionalitas pratayang tautan bawaan
- Tidak ada pemendekan URL otomatis
- Tidak ada pelacakan klik atau analitik
- Tidak ada pemeriksaan kedaluwarsa URL
- Tidak ada pemindaian URL berbahaya

## Peningkatan Masa Depan

### Fitur yang Direncanakan
- Validasi protokol HTTP/HTTPS
- Pola validasi regex kustom
- Penambahan prefiks protokol otomatis
- Pemeriksaan aksesibilitas URL

### Peningkatan Potensial
- Pembuatan pratayang tautan
- Tampilan favicon
- Integrasi pemendekan URL
- Kemampuan pelacakan klik
- Deteksi tautan yang rusak

## Sumber Daya Terkait

- [Bidang Teks](/api/custom-fields/text-single) - Untuk data teks non-URL
- [Bidang Email](/api/custom-fields/email) - Untuk alamat email
- [Ikhtisar Bidang Kustom](/api/custom-fields/2.list-custom-fields) - Konsep umum

## Migrasi dari Bidang Teks

Jika Anda bermigrasi dari bidang teks ke bidang URL:

1. **Buat bidang URL** dengan nama dan konfigurasi yang sama
2. **Ekspor nilai teks yang ada** untuk memverifikasi bahwa mereka adalah URL yang valid
3. **Perbarui catatan** untuk menggunakan bidang URL yang baru
4. **Hapus bidang teks lama** setelah migrasi berhasil
5. **Perbarui aplikasi** untuk menggunakan komponen UI khusus URL

### Contoh Migrasi
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```
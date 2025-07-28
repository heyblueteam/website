---
title: Bidang Kustom ID Unik
description: Buat bidang pengidentifikasi unik yang dihasilkan secara otomatis dengan penomoran berurutan dan format kustom
---

Bidang kustom ID unik secara otomatis menghasilkan pengidentifikasi unik yang berurutan untuk catatan Anda. Mereka sempurna untuk membuat nomor tiket, ID pesanan, nomor faktur, atau sistem pengidentifikasi berurutan lainnya dalam alur kerja Anda.

## Contoh Dasar

Buat bidang ID unik sederhana dengan pengurutan otomatis:

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## Contoh Lanjutan

Buat bidang ID unik yang diformat dengan awalan dan padding nol:

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## Parameter Input

### CreateCustomFieldInput (UNIQUE_ID)

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang ID unik |
| `type` | CustomFieldType! | ✅ Ya | Harus `UNIQUE_ID` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |
| `useSequenceUniqueId` | Boolean | Tidak | Aktifkan pengurutan otomatis (default: false) |
| `prefix` | String | Tidak | Teks awalan untuk ID yang dihasilkan (misalnya, "TUGAS-") |
| `sequenceDigits` | Int | Tidak | Jumlah digit untuk padding nol |
| `sequenceStartingNumber` | Int | Tidak | Nomor awal untuk urutan |

## Opsi Konfigurasi

### Pengurutan Otomatis (`useSequenceUniqueId`)
- **true**: Secara otomatis menghasilkan ID berurutan saat catatan dibuat
- **false** atau **undefined**: Masukan manual diperlukan (berfungsi seperti bidang teks)

### Awalan (`prefix`)
- Awalan teks opsional yang ditambahkan ke semua ID yang dihasilkan
- Contoh: "TUGAS-", "ORD-", "BUG-", "REQ-"
- Tidak ada batasan panjang, tetapi tetap wajar untuk tampilan

### Digit Urutan (`sequenceDigits`)
- Jumlah digit untuk padding nol pada nomor urutan
- Contoh: `sequenceDigits: 3` menghasilkan `001`, `002`, `003`
- Jika tidak ditentukan, tidak ada padding yang diterapkan

### Nomor Awal (`sequenceStartingNumber`)
- Nomor pertama dalam urutan
- Contoh: `sequenceStartingNumber: 1000` mulai dari 1000, 1001, 1002...
- Jika tidak ditentukan, mulai dari 1 (perilaku default)

## Format ID yang Dihasilkan

Format ID akhir menggabungkan semua opsi konfigurasi:

```
{prefix}{paddedSequenceNumber}
```

### Contoh Format

| Konfigurasi | ID yang Dihasilkan |
|--------------|--------------------|
| Tanpa opsi | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## Membaca Nilai ID Unik

### Kuery Catatan dengan ID Unik
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### Format Respons
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## Generasi ID Otomatis

### Kapan ID Dihasilkan
- **Pembuatan Catatan**: ID secara otomatis ditugaskan saat catatan baru dibuat
- **Penambahan Bidang**: Saat menambahkan bidang UNIQUE_ID ke catatan yang ada, pekerjaan latar belakang dijadwalkan (implementasi pekerja pending)
- **Pemrosesan Latar Belakang**: Generasi ID untuk catatan baru terjadi secara sinkron melalui pemicu database

### Proses Generasi
1. **Pemicu**: Catatan baru dibuat atau bidang UNIQUE_ID ditambahkan
2. **Pencarian Urutan**: Sistem menemukan nomor urutan berikutnya yang tersedia
3. **Penugasan ID**: Nomor urutan ditugaskan ke catatan
4. **Pembaruan Penghitung**: Penghitung urutan ditingkatkan untuk catatan mendatang
5. **Pemformatan**: ID diformat dengan awalan dan padding saat ditampilkan

### Jaminan Keunikan
- **Keterbatasan Database**: Pembatasan unik pada ID urutan dalam setiap bidang
- **Operasi Atomik**: Generasi urutan menggunakan kunci database untuk mencegah duplikasi
- **Lingkup Proyek**: Urutan bersifat independen per proyek
- **Perlindungan Kondisi Balapan**: Permintaan bersamaan ditangani dengan aman

## Mode Manual vs Otomatis

### Mode Otomatis (`useSequenceUniqueId: true`)
- ID dihasilkan secara otomatis melalui pemicu database
- Penomoran berurutan dijamin
- Generasi urutan atomik mencegah duplikasi
- ID yang diformat menggabungkan awalan + nomor urutan yang dipadati

### Mode Manual (`useSequenceUniqueId: false` atau `undefined`)
- Berfungsi seperti bidang teks biasa
- Pengguna dapat memasukkan nilai kustom melalui `setTodoCustomField` dengan parameter `text`
- Tidak ada generasi otomatis
- Tidak ada penegakan keunikan di luar batasan database

## Menetapkan Nilai Manual (Hanya Mode Manual)

Ketika `useSequenceUniqueId` adalah false, Anda dapat menetapkan nilai secara manual:

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## Bidang Respons

### Respons TodoCustomField (UNIQUE_ID)

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang kustom |
| `sequenceId` | Int | Nomor urutan yang dihasilkan (diisi untuk bidang UNIQUE_ID) |
| `text` | String | Nilai teks yang diformat (menggabungkan awalan + urutan yang dipadati) |
| `todo` | Todo! | Catatan yang dimiliki nilai ini |
| `createdAt` | DateTime! | Kapan nilai dibuat |
| `updatedAt` | DateTime! | Kapan nilai terakhir diperbarui |

### Respons CustomField (UNIQUE_ID)

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `useSequenceUniqueId` | Boolean | Apakah pengurutan otomatis diaktifkan |
| `prefix` | String | Teks awalan untuk ID yang dihasilkan |
| `sequenceDigits` | Int | Jumlah digit untuk padding nol |
| `sequenceStartingNumber` | Int | Nomor awal untuk urutan |

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|---------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## Respons Kesalahan

### Kesalahan Konfigurasi Bidang
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Kesalahan Izin
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Catatan Penting

### ID yang Dihasilkan Secara Otomatis
- **Hanya Baca**: ID yang dihasilkan secara otomatis tidak dapat diedit secara manual
- **Permanen**: Setelah ditugaskan, ID urutan tidak berubah
- **Kronologis**: ID mencerminkan urutan pembuatan
- **Terbatas**: Urutan bersifat independen per proyek

### Pertimbangan Kinerja
- Generasi ID untuk catatan baru bersifat sinkron melalui pemicu database
- Generasi urutan menggunakan `FOR UPDATE` kunci untuk operasi atomik
- Sistem pekerjaan latar belakang ada tetapi implementasi pekerja masih pending
- Pertimbangkan nomor awal urutan untuk proyek dengan volume tinggi

### Migrasi dan Pembaruan
- Menambahkan pengurutan otomatis ke catatan yang ada menjadwalkan pekerjaan latar belakang (pekerja pending)
- Mengubah pengaturan urutan hanya mempengaruhi catatan mendatang
- ID yang ada tetap tidak berubah saat pembaruan konfigurasi
- Penghitung urutan melanjutkan dari maksimum saat ini

## Praktik Terbaik

### Desain Konfigurasi
- Pilih awalan deskriptif yang tidak akan bertentangan dengan sistem lain
- Gunakan padding digit yang sesuai untuk volume yang diharapkan
- Tetapkan nomor awal yang wajar untuk menghindari konflik
- Uji konfigurasi dengan data sampel sebelum penerapan

### Pedoman Awalan
- Jaga agar awalan tetap pendek dan mudah diingat (2-5 karakter)
- Gunakan huruf kapital untuk konsistensi
- Sertakan pemisah (tanda hubung, garis bawah) untuk keterbacaan
- Hindari karakter khusus yang mungkin menyebabkan masalah di URL atau sistem

### Perencanaan Urutan
- Perkirakan volume catatan Anda untuk memilih padding digit yang sesuai
- Pertimbangkan pertumbuhan di masa depan saat menetapkan nomor awal
- Rencanakan rentang urutan yang berbeda untuk jenis catatan yang berbeda
- Dokumentasikan skema ID Anda untuk referensi tim

## Kasus Penggunaan Umum

1. **Sistem Dukungan**
   - Nomor tiket: `TICK-001`, `TICK-002`
   - ID kasus: `CASE-2024-001`
   - Permintaan dukungan: `SUP-001`

2. **Manajemen Proyek**
   - ID tugas: `TASK-001`, `TASK-002`
   - Item sprint: `SPRINT-001`
   - Nomor hasil: `DEL-001`

3. **Operasi Bisnis**
   - Nomor pesanan: `ORD-2024-001`
   - ID faktur: `INV-001`
   - Pesanan pembelian: `PO-001`

4. **Manajemen Kualitas**
   - Laporan bug: `BUG-001`
   - ID kasus uji: `TEST-001`
   - Nomor tinjauan: `REV-001`

## Fitur Integrasi

### Dengan Automasi
- Memicu tindakan saat ID unik ditugaskan
- Gunakan pola ID dalam aturan automasi
- Referensikan ID dalam template email dan notifikasi

### Dengan Pencarian
- Referensikan ID unik dari catatan lain
- Temukan catatan berdasarkan ID unik
- Tampilkan pengidentifikasi catatan terkait

### Dengan Pelaporan
- Kelompokkan dan filter berdasarkan pola ID
- Lacak tren penugasan ID
- Pantau penggunaan dan celah urutan

## Batasan

- **Hanya Berurutan**: ID ditugaskan dalam urutan kronologis
- **Tanpa Celah**: Catatan yang dihapus meninggalkan celah dalam urutan
- **Tanpa Penggunaan Kembali**: Nomor urutan tidak pernah digunakan kembali
- **Lingkup Proyek**: Tidak dapat berbagi urutan antar proyek
- **Keterbatasan Format**: Opsi pemformatan terbatas
- **Tanpa Pembaruan Massal**: Tidak dapat memperbarui ID urutan yang ada secara massal
- **Tanpa Logika Kustom**: Tidak dapat menerapkan aturan generasi ID kustom

## Sumber Daya Terkait

- [Bidang Teks](/api/custom-fields/text-single) - Untuk pengidentifikasi teks manual
- [Bidang Angka](/api/custom-fields/number) - Untuk urutan numerik
- [Ikhtisar Bidang Kustom](/api/custom-fields/2.list-custom-fields) - Konsep umum
- [Automasi](/api/automations) - Untuk aturan automasi berbasis ID
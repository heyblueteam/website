---
title: Bidang Kustom Penilaian
description: Buat bidang penilaian untuk menyimpan penilaian numerik dengan skala dan validasi yang dapat dikonfigurasi
---

Bidang kustom penilaian memungkinkan Anda menyimpan penilaian numerik dalam catatan dengan nilai minimum dan maksimum yang dapat dikonfigurasi. Mereka ideal untuk penilaian kinerja, skor kepuasan, tingkat prioritas, atau data berbasis skala numerik lainnya dalam proyek Anda.

## Contoh Dasar

Buat bidang penilaian sederhana dengan skala default 0-5:

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## Contoh Lanjutan

Buat bidang penilaian dengan skala dan deskripsi kustom:

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang penilaian |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `RATING` |
| `projectId` | String! | ✅ Ya | ID proyek tempat bidang ini akan dibuat |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |
| `min` | Float | Tidak | Nilai penilaian minimum (tanpa default) |
| `max` | Float | Tidak | Nilai penilaian maksimum |

## Mengatur Nilai Penilaian

Untuk mengatur atau memperbarui nilai penilaian pada catatan:

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID catatan untuk diperbarui |
| `customFieldId` | String! | ✅ Ya | ID bidang kustom penilaian |
| `value` | String! | ✅ Ya | Nilai penilaian sebagai string (dalam rentang yang dikonfigurasi) |

## Membuat Catatan dengan Nilai Penilaian

Saat membuat catatan baru dengan nilai penilaian:

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        min
        max
      }
      value
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
| `value` | Float | Nilai penilaian yang disimpan (diakses melalui customField.value) |
| `todo` | Todo! | Catatan yang dimiliki nilai ini |
| `createdAt` | DateTime! | Waktu nilai dibuat |
| `updatedAt` | DateTime! | Waktu nilai terakhir dimodifikasi |

**Catatan**: Nilai penilaian sebenarnya diakses melalui `customField.value.number` dalam kueri.

### CustomField Respons

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk bidang |
| `name` | String! | Nama tampilan dari bidang penilaian |
| `type` | CustomFieldType! | Selalu `RATING` |
| `min` | Float | Nilai penilaian minimum yang diizinkan |
| `max` | Float | Nilai penilaian maksimum yang diizinkan |
| `description` | String | Teks bantuan untuk bidang |

## Validasi Penilaian

### Pembatasan Nilai
- Nilai penilaian harus numerik (tipe Float)
- Nilai harus berada dalam rentang min/max yang dikonfigurasi
- Jika tidak ada minimum yang ditentukan, tidak ada nilai default
- Nilai maksimum bersifat opsional tetapi disarankan

### Aturan Validasi
**Penting**: Validasi hanya terjadi saat mengirimkan formulir, tidak saat menggunakan `setTodoCustomField` secara langsung.

- Input diparsing sebagai angka float (saat menggunakan formulir)
- Harus lebih besar dari atau sama dengan nilai minimum (saat menggunakan formulir)
- Harus kurang dari atau sama dengan nilai maksimum (saat menggunakan formulir)
- `setTodoCustomField` menerima nilai string apa pun tanpa validasi

### Contoh Penilaian Valid
Untuk bidang dengan min=1, max=5:
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### Contoh Penilaian Tidak Valid
Untuk bidang dengan min=1, max=5:
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## Opsi Konfigurasi

### Pengaturan Skala Penilaian
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### Skala Penilaian Umum
- **1-5 Bintang**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 Kinerja**: `min: 1, max: 10`
- **0-100 Persentase**: `min: 0, max: 100`
- **Skala Kustom**: Rentang numerik apa pun

## Izin yang Diperlukan

Operasi bidang kustom mengikuti izin berbasis peran standar:

| Tindakan | Peran Diperlukan |
|----------|------------------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**Catatan**: Peran spesifik yang diperlukan bergantung pada konfigurasi peran kustom proyek Anda dan izin tingkat bidang.

## Respons Kesalahan

### Kesalahan Validasi (Hanya Formulir)
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

**Penting**: Validasi nilai penilaian (pembatasan min/max) hanya terjadi saat mengirimkan formulir, tidak saat menggunakan `setTodoCustomField` secara langsung.

### Bidang Kustom Tidak Ditemukan
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

## Praktik Terbaik

### Desain Skala
- Gunakan skala penilaian yang konsisten di seluruh bidang yang serupa
- Pertimbangkan familiaritas pengguna (1-5 bintang, 0-10 NPS)
- Tetapkan nilai minimum yang sesuai (0 vs 1)
- Definisikan makna yang jelas untuk setiap tingkat penilaian

### Kualitas Data
- Validasi nilai penilaian sebelum menyimpan
- Gunakan presisi desimal dengan tepat
- Pertimbangkan pembulatan untuk tujuan tampilan
- Berikan panduan yang jelas tentang makna penilaian

### Pengalaman Pengguna
- Tampilkan skala penilaian secara visual (bintang, bilah kemajuan)
- Tampilkan nilai saat ini dan batas skala
- Berikan konteks untuk makna penilaian
- Pertimbangkan nilai default untuk catatan baru

## Kasus Penggunaan Umum

1. **Manajemen Kinerja**
   - Penilaian kinerja karyawan
   - Skor kualitas proyek
   - Penilaian penyelesaian tugas
   - Penilaian tingkat keterampilan

2. **Umpan Balik Pelanggan**
   - Penilaian kepuasan
   - Skor kualitas produk
   - Penilaian pengalaman layanan
   - Skor Net Promoter (NPS)

3. **Prioritas dan Kepentingan**
   - Tingkat prioritas tugas
   - Penilaian urgensi
   - Skor penilaian risiko
   - Penilaian dampak

4. **Jaminan Kualitas**
   - Penilaian tinjauan kode
   - Skor kualitas pengujian
   - Kualitas dokumentasi
   - Penilaian kepatuhan proses

## Fitur Integrasi

### Dengan Automasi
- Memicu tindakan berdasarkan ambang penilaian
- Mengirim notifikasi untuk penilaian rendah
- Membuat tugas tindak lanjut untuk penilaian tinggi
- Mengarahkan pekerjaan berdasarkan nilai penilaian

### Dengan Pencarian
- Menghitung rata-rata penilaian di seluruh catatan
- Menemukan catatan berdasarkan rentang penilaian
- Mengacu pada data penilaian dari catatan lain
- Mengagregasi statistik penilaian

### Dengan Antarmuka Blue
- Validasi rentang otomatis dalam konteks formulir
- Kontrol input penilaian visual
- Umpan balik validasi waktu nyata
- Opsi input bintang atau slider

## Pelacakan Aktivitas

Perubahan bidang penilaian secara otomatis dilacak:
- Nilai penilaian lama dan baru dicatat
- Aktivitas menunjukkan perubahan numerik
- Cap waktu untuk semua pembaruan penilaian
- Atribusi pengguna untuk perubahan

## Batasan

- Hanya nilai numerik yang didukung
- Tidak ada tampilan penilaian visual bawaan (bintang, dll.)
- Presisi desimal tergantung pada konfigurasi basis data
- Tidak ada penyimpanan metadata penilaian (komentar, konteks)
- Tidak ada agregasi atau statistik penilaian otomatis
- Tidak ada konversi penilaian bawaan antara skala
- **Kritis**: Validasi min/max hanya berfungsi di formulir, tidak melalui `setTodoCustomField`

## Sumber Daya Terkait

- [Bidang Angka](/api/5.custom%20fields/number) - Untuk data numerik umum
- [Bidang Persentase](/api/5.custom%20fields/percent) - Untuk nilai persentase
- [Bidang Pilih](/api/5.custom%20fields/select-single) - Untuk penilaian pilihan diskrit
- [Ikhtisar Bidang Kustom](/api/5.custom%20fields/2.list-custom-fields) - Konsep umum
---
title: Bidang Kustom Angka
description: Buat bidang angka untuk menyimpan nilai numerik dengan batasan min/max opsional dan format awalan
---

Bidang kustom angka memungkinkan Anda untuk menyimpan nilai numerik untuk catatan. Mereka mendukung batasan validasi, presisi desimal, dan dapat digunakan untuk kuantitas, skor, pengukuran, atau data numerik lainnya yang tidak memerlukan format khusus.

## Contoh Dasar

Buat bidang angka sederhana:

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang angka dengan batasan dan awalan:

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang angka |
| `type` | CustomFieldType! | ✅ Ya | Harus `NUMBER` |
| `projectId` | String! | ✅ Ya | ID proyek untuk membuat bidang ini |
| `min` | Float | Tidak | Batasan nilai minimum (hanya panduan UI) |
| `max` | Float | Tidak | Batasan nilai maksimum (hanya panduan UI) |
| `prefix` | String | Tidak | Awalan tampilan (misalnya, "#", "~", "$") |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

## Menetapkan Nilai Angka

Bidang angka menyimpan nilai desimal dengan validasi opsional:

### Nilai Angka Sederhana

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### Nilai Bulat

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID catatan untuk diperbarui |
| `customFieldId` | String! | ✅ Ya | ID bidang kustom angka |
| `number` | Float | Tidak | Nilai numerik yang akan disimpan |

## Batasan Nilai

### Batasan Min/Maks (Panduan UI)

**Penting**: Batasan min/maks disimpan tetapi TIDAK ditegakkan di sisi server. Mereka berfungsi sebagai panduan UI untuk aplikasi frontend.

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**Validasi Sisi Klien Diperlukan**: Aplikasi frontend harus menerapkan logika validasi untuk menegakkan batasan min/maks.

### Tipe Nilai yang Didukung

| Tipe | Contoh | Deskripsi |
|------|--------|-----------|
| Integer | `42` | Angka bulat |
| Decimal | `42.5` | Angka dengan tempat desimal |
| Negative | `-10` | Nilai negatif (jika tidak ada batasan min) |
| Zero | `0` | Nilai nol |

**Catatan**: Batasan min/maks TIDAK divalidasi di sisi server. Nilai di luar rentang yang ditentukan akan diterima dan disimpan.

## Membuat Catatan dengan Nilai Angka

Saat membuat catatan baru dengan nilai angka:

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
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
        prefix
      }
      number
      value
    }
  }
}
```

### Format Input yang Didukung

Saat membuat catatan, gunakan parameter `number` (bukan `value`) dalam array bidang kustom:

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## Bidang Respon

### Respon TodoCustomField

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang kustom |
| `number` | Float | Nilai numerik |
| `todo` | Todo! | Catatan yang dimiliki nilai ini |
| `createdAt` | DateTime! | Kapan nilai dibuat |
| `updatedAt` | DateTime! | Kapan nilai terakhir dimodifikasi |

### Respon CustomField

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk definisi bidang |
| `name` | String! | Nama tampilan dari bidang |
| `type` | CustomFieldType! | Selalu `NUMBER` |
| `min` | Float | Nilai minimum yang diizinkan |
| `max` | Float | Nilai maksimum yang diizinkan |
| `prefix` | String | Awalan tampilan |
| `description` | String | Teks bantuan |

**Catatan**: Jika nilai angka tidak diatur, bidang `number` akan `null`.

## Penyaringan dan Kuery

Bidang angka mendukung penyaringan numerik yang komprehensif:

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### Operator yang Didukung

| Operator | Deskripsi | Contoh |
|----------|-----------|--------|
| `EQ` | Sama dengan | `number = 42` |
| `NE` | Tidak sama dengan | `number ≠ 42` |
| `GT` | Lebih besar dari | `number > 42` |
| `GTE` | Lebih besar dari atau sama dengan | `number ≥ 42` |
| `LT` | Kurang dari | `number < 42` |
| `LTE` | Kurang dari atau sama dengan | `number ≤ 42` |
| `IN` | Dalam array | `number in [1, 2, 3]` |
| `NIN` | Tidak dalam array | `number not in [1, 2, 3]` |
| `IS` | Adalah null/tidak null | `number is null` |

### Penyaringan Rentang

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## Format Tampilan

### Dengan Awalan

Jika awalan diatur, itu akan ditampilkan:

| Nilai | Awalan | Tampilan |
|-------|--------|----------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### Presisi Desimal

Angka mempertahankan presisi desimal mereka:

| Input | Disimpan | Ditampilkan |
|-------|----------|-------------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|----------------------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## Respon Kesalahan

### Format Angka Tidak Valid
```json
{
  "errors": [{
    "message": "Invalid number format",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**Catatan**: Kesalahan validasi min/maks TIDAK terjadi di sisi server. Validasi batasan harus diterapkan dalam aplikasi frontend Anda.

### Bukan Angka
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Praktik Terbaik

### Desain Batasan
- Tetapkan nilai min/maks yang realistis untuk panduan UI
- Terapkan validasi sisi klien untuk menegakkan batasan
- Gunakan batasan untuk memberikan umpan balik kepada pengguna dalam formulir
- Pertimbangkan apakah nilai negatif valid untuk kasus penggunaan Anda

### Presisi Nilai
- Gunakan presisi desimal yang sesuai untuk kebutuhan Anda
- Pertimbangkan pembulatan untuk tujuan tampilan
- Konsisten dengan presisi di seluruh bidang terkait

### Peningkatan Tampilan
- Gunakan awalan yang bermakna untuk konteks
- Pertimbangkan satuan dalam nama bidang (misalnya, "Berat (kg)")
- Berikan deskripsi yang jelas untuk aturan validasi

## Kasus Penggunaan Umum

1. **Sistem Penilaian**
   - Penilaian kinerja
   - Skor kualitas
   - Tingkat prioritas
   - Penilaian kepuasan pelanggan

2. **Pengukuran**
   - Kuantitas dan jumlah
   - Dimensi dan ukuran
   - Durasi (dalam format numerik)
   - Kapasitas dan batas

3. **Metrik Bisnis**
   - Angka pendapatan
   - Tingkat konversi
   - Alokasi anggaran
   - Angka target

4. **Data Teknis**
   - Nomor versi
   - Nilai konfigurasi
   - Metrik kinerja
   - Pengaturan ambang batas

## Fitur Integrasi

### Dengan Grafik dan Dasbor
- Gunakan bidang ANGKA dalam perhitungan grafik
- Buat visualisasi numerik
- Lacak tren dari waktu ke waktu

### Dengan Automasi
- Picu tindakan berdasarkan ambang angka
- Perbarui bidang terkait berdasarkan perubahan angka
- Kirim notifikasi untuk nilai tertentu

### Dengan Pencarian
- Agregasi angka dari catatan terkait
- Hitung total dan rata-rata
- Temukan nilai min/maks di seluruh hubungan

### Dengan Grafik
- Buat visualisasi numerik
- Lacak tren dari waktu ke waktu
- Bandingkan nilai di seluruh catatan

## Batasan

- **Tidak ada validasi sisi server** dari batasan min/maks
- **Validasi sisi klien diperlukan** untuk penegakan batasan
- Tidak ada format mata uang bawaan (gunakan tipe MATA UANG sebagai gantinya)
- Tidak ada simbol persentase otomatis (gunakan tipe PERSENTASE sebagai gantinya)
- Tidak ada kemampuan konversi satuan
- Presisi desimal dibatasi oleh tipe Desimal database
- Tidak ada evaluasi rumus matematis di dalam bidang itu sendiri

## Sumber Daya Terkait

- [Ikhtisar Bidang Kustom](/api/custom-fields/1.index) - Konsep umum bidang kustom
- [Bidang Kustom Mata Uang](/api/custom-fields/currency) - Untuk nilai moneter
- [Bidang Kustom Persentase](/api/custom-fields/percent) - Untuk nilai persentase
- [API Automasi](/api/automations/1.index) - Buat automasi berbasis angka
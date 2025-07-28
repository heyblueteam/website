---
title: Bidang Kustom Persentase
description: Buat bidang persentase untuk menyimpan nilai numerik dengan penanganan simbol % otomatis dan format tampilan
---

Bidang kustom persentase memungkinkan Anda untuk menyimpan nilai persentase untuk catatan. Mereka secara otomatis menangani simbol % untuk input dan tampilan, sambil menyimpan nilai numerik mentah secara internal. Sempurna untuk tingkat penyelesaian, tingkat keberhasilan, atau metrik berbasis persentase lainnya.

## Contoh Dasar

Buat bidang persentase sederhana:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang persentase dengan deskripsi:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
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
| `name` | String! | ✅ Ya | Nama tampilan dari bidang persentase |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `PERCENT` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

**Catatan**: Konteks proyek ditentukan secara otomatis dari header otentikasi Anda. Tidak ada parameter `projectId` yang diperlukan.

**Catatan**: Bidang PERCENT tidak mendukung batasan min/max atau format awalan seperti bidang NUMBER.

## Mengatur Nilai Persentase

Bidang persentase menyimpan nilai numerik dengan penanganan simbol % otomatis:

### Dengan Simbol Persentase

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### Nilai Numerik Langsung

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom persentase |
| `number` | Float | Tidak | Nilai persentase numerik (misalnya, 75.5 untuk 75.5%) |

## Penyimpanan dan Tampilan Nilai

### Format Penyimpanan
- **Penyimpanan internal**: Nilai numerik mentah (misalnya, 75.5)
- **Database**: Disimpan sebagai `Decimal` dalam kolom `number`
- **GraphQL**: Dikembalikan sebagai tipe `Float`

### Format Tampilan
- **Antarmuka pengguna**: Aplikasi klien harus menambahkan simbol % (misalnya, "75.5%")
- **Grafik**: Menampilkan dengan simbol % ketika tipe output adalah PERSENTASE
- **Respon API**: Nilai numerik mentah tanpa simbol % (misalnya, 75.5)

## Membuat Catatan dengan Nilai Persentase

Ketika membuat catatan baru dengan nilai persentase:

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### Format Input yang Didukung

| Format | Contoh | Hasil |
|--------|--------|-------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**Catatan**: Simbol % secara otomatis dihapus dari input dan ditambahkan kembali selama tampilan.

## Menanyakan Nilai Persentase

Ketika menanyakan catatan dengan bidang kustom persentase, akses nilai melalui jalur `customField.value.number`:

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

Respon akan mencakup persentase sebagai angka mentah:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## Bidang Respon

### Respon TodoCustomField

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | ID! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang kustom (berisi nilai persentase) |
| `todo` | Todo! | Catatan yang nilai ini miliki |
| `createdAt` | DateTime! | Ketika nilai dibuat |
| `updatedAt` | DateTime! | Ketika nilai terakhir dimodifikasi |

**Penting**: Nilai persentase diakses melalui bidang `customField.value.number`. Simbol % tidak termasuk dalam nilai yang disimpan dan harus ditambahkan oleh aplikasi klien untuk tampilan.

## Penyaringan dan Penanyaan

Bidang persentase mendukung penyaringan yang sama seperti bidang NUMBER:

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
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
| `EQ` | Sama dengan | `percentage = 75` |
| `NE` | Tidak sama dengan | `percentage ≠ 75` |
| `GT` | Lebih besar dari | `percentage > 75` |
| `GTE` | Lebih besar dari atau sama dengan | `percentage ≥ 75` |
| `LT` | Kurang dari | `percentage < 75` |
| `LTE` | Kurang dari atau sama dengan | `percentage ≤ 75` |
| `IN` | Nilai dalam daftar | `percentage in [50, 75, 100]` |
| `NIN` | Nilai tidak dalam daftar | `percentage not in [0, 25]` |
| `IS` | Periksa null dengan `values: null` | `percentage is null` |
| `NOT` | Periksa tidak null dengan `values: null` | `percentage is not null` |

### Penyaringan Rentang

Untuk penyaringan rentang, gunakan beberapa operator:

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## Rentang Nilai Persentase

### Rentang Umum

| Rentang | Deskripsi | Kasus Penggunaan |
|---------|-----------|------------------|
| `0-100` | Persentase standar | Completion rates, success rates |
| `0-∞` | Persentase tidak terbatas | Growth rates, performance metrics |
| `-∞-∞` | Nilai apa pun | Change rates, variance |

### Nilai Contoh

| Input | Disimpan | Tampilan |
|-------|----------|----------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## Agregasi Grafik

Bidang persentase mendukung agregasi dalam grafik dasbor dan laporan. Fungsi yang tersedia termasuk:

- `AVERAGE` - Nilai persentase rata-rata
- `COUNT` - Jumlah catatan dengan nilai
- `MIN` - Nilai persentase terendah
- `MAX` - Nilai persentase tertinggi 
- `SUM` - Total dari semua nilai persentase

Agregasi ini tersedia saat membuat grafik dan dasbor, tidak dalam kueri GraphQL langsung.

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|----------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## Respon Kesalahan

### Format Persentase Tidak Valid
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

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

### Entri Nilai
- Izinkan pengguna untuk memasukkan dengan atau tanpa simbol %
- Validasi rentang yang wajar untuk kasus penggunaan Anda
- Berikan konteks yang jelas tentang apa yang diwakili 100%

### Tampilan
- Selalu tampilkan simbol % di antarmuka pengguna
- Gunakan presisi desimal yang sesuai
- Pertimbangkan pengkodean warna untuk rentang (merah/kuning/hijau)

### Interpretasi Data
- Dokumentasikan apa arti 100% dalam konteks Anda
- Tangani nilai di atas 100% dengan tepat
- Pertimbangkan apakah nilai negatif valid

## Kasus Penggunaan Umum

1. **Manajemen Proyek**
   - Tingkat penyelesaian tugas
   - Kemajuan proyek
   - Pemanfaatan sumber daya
   - Kecepatan sprint

2. **Pelacakan Kinerja**
   - Tingkat keberhasilan
   - Tingkat kesalahan
   - Metrik efisiensi
   - Skor kualitas

3. **Metrik Keuangan**
   - Tingkat pertumbuhan
   - Margin keuntungan
   - Jumlah diskon
   - Persentase perubahan

4. **Analitik**
   - Tingkat konversi
   - Tingkat klik
   - Metrik keterlibatan
   - Indikator kinerja

## Fitur Integrasi

### Dengan Rumus
- Referensi bidang PERCENT dalam perhitungan
- Format simbol % otomatis dalam output rumus
- Gabungkan dengan bidang numerik lainnya

### Dengan Automasi
- Trigger tindakan berdasarkan ambang persentase
- Kirim notifikasi untuk persentase tonggak
- Perbarui status berdasarkan tingkat penyelesaian

### Dengan Pencarian
- Agregasi persentase dari catatan terkait
- Hitung rata-rata tingkat keberhasilan
- Temukan item dengan kinerja tertinggi/terendah

### Dengan Grafik
- Buat visualisasi berbasis persentase
- Lacak kemajuan dari waktu ke waktu
- Bandingkan metrik kinerja

## Perbedaan dari Bidang NUMBER

### Apa yang Berbeda
- **Penanganan input**: Secara otomatis menghapus simbol %
- **Tampilan**: Secara otomatis menambahkan simbol %
- **Batasan**: Tidak ada validasi min/max
- **Format**: Tidak ada dukungan awalan

### Apa yang Sama
- **Penyimpanan**: Kolom dan tipe database yang sama
- **Penyaringan**: Operator kueri yang sama
- **Agregasi**: Fungsi agregasi yang sama
- **Izin**: Model izin yang sama

## Batasan

- Tidak ada batasan nilai min/max
- Tidak ada opsi format awalan
- Tidak ada validasi otomatis untuk rentang 0-100%
- Tidak ada konversi antara format persentase (misalnya, 0.75 ↔ 75%)
- Nilai di atas 100% diperbolehkan

## Sumber Daya Terkait

- [Ikhtisar Bidang Kustom](/api/custom-fields/list-custom-fields) - Konsep bidang kustom umum
- [Bidang Kustom Angka](/api/custom-fields/number) - Untuk nilai numerik mentah
- [API Automasi](/api/automations/index) - Buat automasi berbasis persentase
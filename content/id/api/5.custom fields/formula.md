---
title: Bidang Kustom Formula
description: Buat bidang yang dihitung yang secara otomatis menghitung nilai berdasarkan data lain
---

Bidang kustom formula digunakan untuk perhitungan grafik dan dasbor dalam Blue. Mereka mendefinisikan fungsi agregasi (JUMLAH, RATA-RATA, HITUNG, dll.) yang beroperasi pada data bidang kustom untuk menampilkan metrik yang dihitung dalam grafik. Formula tidak dihitung pada tingkat todo individu tetapi mengagregasi data dari beberapa catatan untuk tujuan visualisasi.

## Contoh Dasar

Buat bidang formula untuk perhitungan grafik:

```graphql
mutation CreateFormulaField {
  createCustomField(input: {
    name: "Budget Total"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Budget)"
        html: "<span>SUM(Budget)</span>"
      }
      display: {
        type: NUMBER
        precision: 2
        function: SUM
      }
    }
  }) {
    id
    name
    type
    formula
  }
}
```

## Contoh Lanjutan

Buat formula mata uang dengan perhitungan kompleks:

```graphql
mutation CreateCurrencyFormula {
  createCustomField(input: {
    name: "Profit Margin"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Revenue) - SUM(Costs)"
        html: "<span>SUM(Revenue) - SUM(Costs)</span>"
      }
      display: {
        type: CURRENCY
        currency: {
          code: "USD"
          name: "US Dollar"
        }
        precision: 2
      }
    }
    description: "Automatically calculates profit by subtracting costs from revenue"
  }) {
    id
    name
    type
    formula
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang formula |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `FORMULA` |
| `projectId` | String! | ✅ Ya | ID proyek tempat bidang ini akan dibuat |
| `formula` | JSON | Tidak | Definisi formula untuk perhitungan grafik |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

### Struktur Formula

```json
{
  "logic": {
    "text": "Display text for the formula",
    "html": "HTML formatted display text"
  },
  "display": {
    "type": "NUMBER|CURRENCY|PERCENTAGE",
    "currency": {
      "code": "USD",
      "name": "US Dollar"  
    },
    "precision": 2,
    "function": "SUM|AVERAGE|AVERAGEA|COUNT|COUNTA|MAX|MIN"
  }
}
```

## Fungsi yang Didukung

### Fungsi Agregasi Grafik

Bidang formula mendukung fungsi agregasi berikut untuk perhitungan grafik:

| Fungsi | Deskripsi | Enum ChartFunction |
|--------|-----------|--------------------|
| `SUM` | Jumlah dari semua nilai | `SUM` |
| `AVERAGE` | Rata-rata nilai numerik | `AVERAGE` |
| `AVERAGEA` | Rata-rata tanpa nol dan null | `AVERAGEA` |
| `COUNT` | Hitung nilai | `COUNT` |
| `COUNTA` | Hitung tanpa nol dan null | `COUNTA` |
| `MAX` | Nilai maksimum | `MAX` |
| `MIN` | Nilai minimum | `MIN` |

**Catatan**: Fungsi-fungsi ini digunakan dalam bidang `display.function` dan beroperasi pada data yang teragregasi untuk visualisasi grafik. Ekspresi matematis kompleks atau perhitungan tingkat bidang tidak didukung.

## Jenis Tampilan

### Tampilan Angka

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

Hasil: `1250.75`

### Tampilan Mata Uang

```json
{
  "display": {
    "type": "CURRENCY",
    "currency": {
      "code": "USD",
      "name": "US Dollar"
    },
    "precision": 2
  }
}
```

Hasil: `$1,250.75`

### Tampilan Persentase

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

Hasil: `87.5%`

## Mengedit Bidang Formula

Perbarui bidang formula yang ada:

```graphql
mutation EditFormulaField {
  editCustomField(input: {
    customFieldId: "field_456"
    formula: {
      logic: {
        text: "AVERAGE(Score)"
        html: "<span>AVERAGE(Score)</span>"
      }
      display: {
        type: PERCENTAGE
        precision: 1
      }
    }
  }) {
    id
    formula
  }
}
```

## Pemrosesan Formula

### Konteks Perhitungan Grafik

Bidang formula diproses dalam konteks segmen grafik dan dasbor:
- Perhitungan terjadi saat grafik dirender atau diperbarui
- Hasil disimpan dalam `ChartSegment.formulaResult` sebagai nilai desimal
- Pemrosesan ditangani melalui antrean BullMQ khusus bernama 'formula'
- Pembaruan diterbitkan kepada pelanggan dasbor untuk pembaruan waktu nyata

### Pemformatan Tampilan

Fungsi `getFormulaDisplayValue` memformat hasil yang dihitung berdasarkan jenis tampilan:
- **ANGKA**: Menampilkan sebagai angka biasa dengan presisi opsional
- **PERSENTASE**: Menambahkan akhiran % dengan presisi opsional  
- **MATA UANG**: Memformat menggunakan kode mata uang yang ditentukan

## Penyimpanan Hasil Formula

Hasil disimpan dalam bidang `formulaResult`:

```json
{
  "number": 1250.75,
  "formulaResult": {
    "number": 1250.75,
    "display": {
      "type": "CURRENCY",
      "currency": {
        "code": "USD",
        "name": "US Dollar"
      },
      "precision": 2
    }
  }
}
```

## Bidang Respons

### Respons TodoCustomField

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang formula |
| `number` | Float | Hasil numerik yang dihitung |
| `formulaResult` | JSON | Hasil lengkap dengan pemformatan tampilan |
| `todo` | Todo! | Catatan yang dimiliki nilai ini |
| `createdAt` | DateTime! | Kapan nilai dibuat |
| `updatedAt` | DateTime! | Kapan nilai terakhir dihitung |

## Konteks Data

### Sumber Data Grafik

Bidang formula beroperasi dalam konteks sumber data grafik:
- Formula mengagregasi nilai bidang kustom di seluruh todo dalam proyek
- Fungsi agregasi yang ditentukan dalam `display.function` menentukan perhitungan
- Hasil dihitung menggunakan fungsi agregasi SQL (avg, sum, count, dll.)
- Perhitungan dilakukan di tingkat basis data untuk efisiensi

## Contoh Formula Umum

### Total Anggaran (Tampilan Grafik)

```json
{
  "logic": {
    "text": "Total Budget",
    "html": "<span>Total Budget</span>"
  },
  "display": {
    "type": "CURRENCY",
    "currency": { "code": "USD", "name": "US Dollar" },
    "precision": 2,
    "function": "SUM"
  }
}
```

### Rata-rata Skor (Tampilan Grafik)

```json
{
  "logic": {
    "text": "Average Quality Score",
    "html": "<span>Average Quality Score</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 1,
    "function": "AVERAGE"
  }
}
```

### Hitung Tugas (Tampilan Grafik)

```json
{
  "logic": {
    "text": "Total Tasks",
    "html": "<span>Total Tasks</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 0,
    "function": "COUNT"
  }
}
```

## Izin yang Diperlukan

Operasi bidang kustom mengikuti izin berbasis peran standar:

| Tindakan | Peran yang Diperlukan |
|----------|-----------------------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**Catatan**: Peran spesifik yang diperlukan tergantung pada konfigurasi peran kustom proyek Anda. Tidak ada konstanta izin khusus seperti CUSTOM_FIELDS_CREATE.

## Penanganan Kesalahan

### Kesalahan Validasi
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

### Desain Formula
- Gunakan nama yang jelas dan deskriptif untuk bidang formula
- Tambahkan deskripsi yang menjelaskan logika perhitungan
- Uji formula dengan data sampel sebelum penerapan
- Jaga agar formula tetap sederhana dan mudah dibaca

### Optimisasi Kinerja
- Hindari ketergantungan formula yang terlalu dalam
- Gunakan referensi bidang spesifik daripada wildcard
- Pertimbangkan strategi caching untuk perhitungan kompleks
- Pantau kinerja formula dalam proyek besar

### Kualitas Data
- Validasi data sumber sebelum digunakan dalam formula
- Tangani nilai kosong atau null dengan tepat
- Gunakan presisi yang sesuai untuk jenis tampilan
- Pertimbangkan kasus tepi dalam perhitungan

## Kasus Penggunaan Umum

1. **Pelacakan Keuangan**
   - Perhitungan anggaran
   - Laporan laba/rugi
   - Analisis biaya
   - Proyeksi pendapatan

2. **Manajemen Proyek**
   - Persentase penyelesaian
   - Pemanfaatan sumber daya
   - Perhitungan garis waktu
   - Metrik kinerja

3. **Kontrol Kualitas**
   - Skor rata-rata
   - Tingkat lulus/gagal
   - Metrik kualitas
   - Pelacakan kepatuhan

4. **Intelijen Bisnis**
   - Perhitungan KPI
   - Analisis tren
   - Metrik komparatif
   - Nilai dasbor

## Batasan

- Formula hanya untuk agregasi grafik/dasbor, bukan perhitungan tingkat todo
- Terbatas pada tujuh fungsi agregasi yang didukung (JUMLAH, RATA-RATA, dll.)
- Tidak ada ekspresi matematis kompleks atau perhitungan antar bidang
- Tidak dapat merujuk ke beberapa bidang dalam satu formula
- Hasil hanya terlihat dalam grafik dan dasbor
- Bidang `logic` hanya untuk teks tampilan, bukan logika perhitungan yang sebenarnya

## Sumber Daya Terkait

- [Bidang Angka](/api/5.custom%20fields/number) - Untuk nilai numerik statis
- [Bidang Mata Uang](/api/5.custom%20fields/currency) - Untuk nilai moneter
- [Bidang Referensi](/api/5.custom%20fields/reference) - Untuk data lintas proyek
- [Bidang Lookup](/api/5.custom%20fields/lookup) - Untuk data teragregasi
- [Ikhtisar Bidang Kustom](/api/5.custom%20fields/2.list-custom-fields) - Konsep umum
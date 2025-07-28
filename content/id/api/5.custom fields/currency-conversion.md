---
title: Bidang Kustom Konversi Mata Uang
description: Buat bidang yang secara otomatis mengonversi nilai mata uang menggunakan nilai tukar waktu nyata
---

Bidang kustom Konversi Mata Uang secara otomatis mengonversi nilai dari bidang CURRENCY sumber ke mata uang target yang berbeda menggunakan nilai tukar waktu nyata. Bidang ini diperbarui secara otomatis setiap kali nilai mata uang sumber berubah.

Nilai tukar konversi disediakan oleh [Frankfurter API](https://github.com/hakanensari/frankfurter), sebuah layanan sumber terbuka yang melacak nilai tukar referensi yang diterbitkan oleh [Bank Sentral Eropa](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html). Ini memastikan konversi mata uang yang akurat, dapat diandalkan, dan terkini untuk kebutuhan bisnis internasional Anda.

## Contoh Dasar

Buat bidang konversi mata uang yang sederhana:

```graphql
mutation CreateCurrencyConversionField {
  createCustomField(input: {
    name: "Price in EUR"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_currency_field_id"
    conversionDateType: "currentDate"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
  }
}
```

## Contoh Lanjutan

Buat bidang konversi dengan tanggal tertentu untuk nilai tukar historis:

```graphql
mutation CreateHistoricalConversionField {
  createCustomField(input: {
    name: "Q1 Budget in Local Currency"
    type: CURRENCY_CONVERSION
    currencyFieldId: "budget_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2024-01-01T00:00:00Z"
    description: "Budget converted at Q1 exchange rates"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
    conversionDate
  }
}
```

## Proses Pengaturan Lengkap

Mengatur bidang konversi mata uang memerlukan tiga langkah:

### Langkah 1: Buat Bidang CURRENCY Sumber

```graphql
mutation CreateSourceCurrencyField {
  createCustomField(input: {
    name: "Contract Value"
    type: CURRENCY
    currency: "USD"
  }) {
    id  # Save this ID for Step 2
    name
    type
  }
}
```

### Langkah 2: Buat Bidang CURRENCY_CONVERSION

```graphql
mutation CreateConversionField {
  createCustomField(input: {
    name: "Contract Value (Local Currency)"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id_from_step_1"
    conversionDateType: "currentDate"
  }) {
    id  # Save this ID for Step 3
    name
    type
  }
}
```

### Langkah 3: Buat Opsi Konversi

```graphql
mutation CreateConversionOptions {
  createCustomFieldOptions(input: {
    customFieldId: "conversion_field_id_from_step_2"
    customFieldOptions: [
      {
        title: "USD to EUR"
        currencyConversionFrom: "USD"
        currencyConversionTo: "EUR"
      },
      {
        title: "USD to GBP"
        currencyConversionFrom: "USD"
        currencyConversionTo: "GBP"
      },
      {
        title: "Any to JPY"
        currencyConversionFrom: "Any"
        currencyConversionTo: "JPY"
      }
    ]
  }) {
    id
    title
    currencyConversionFrom
    currencyConversionTo
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang konversi |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | Tidak | ID dari bidang CURRENCY sumber untuk dikonversi |
| `conversionDateType` | String | Tidak | Strategi tanggal untuk nilai tukar (lihat di bawah) |
| `conversionDate` | String | Tidak | String tanggal untuk konversi (berdasarkan conversionDateType) |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

**Catatan**: Bidang kustom secara otomatis diasosiasikan dengan proyek berdasarkan konteks proyek pengguna saat ini. Tidak ada `projectId` parameter yang diperlukan.

### Jenis Tanggal Konversi

| Tipe | Deskripsi | Parameter conversionDate |
|------|----------|-------------------------|
| `currentDate` | Menggunakan nilai tukar waktu nyata | Tidak diperlukan |
| `specificDate` | Menggunakan nilai tukar dari tanggal tetap | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | Menggunakan tanggal dari bidang lain | "todoDueDate" or DATE field ID |

## Membuat Opsi Konversi

Opsi konversi mendefinisikan pasangan mata uang mana yang dapat dikonversi:

### CreateCustomFieldOptionInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `customFieldId` | String! | ✅ Ya | ID dari bidang CURRENCY_CONVERSION |
| `title` | String! | ✅ Ya | Nama tampilan untuk opsi konversi ini |
| `currencyConversionFrom` | String! | ✅ Ya | Kode mata uang sumber atau "Any" |
| `currencyConversionTo` | String! | ✅ Ya | Kode mata uang target |

### Menggunakan "Any" sebagai Sumber

Nilai khusus "Any" sebagai `currencyConversionFrom` menciptakan opsi cadangan:

```graphql
mutation CreateUniversalConversion {
  createCustomFieldOption(input: {
    customFieldId: "conversion_field_id"
    title: "Any currency to EUR"
    currencyConversionFrom: "Any"
    currencyConversionTo: "EUR"
  }) {
    id
  }
}
```

Opsi ini akan digunakan ketika tidak ada pasangan mata uang tertentu yang ditemukan.

## Cara Kerja Konversi Otomatis

1. **Pembaruan Nilai**: Ketika nilai diatur dalam bidang CURRENCY sumber
2. **Pencocokan Opsi**: Sistem menemukan opsi konversi yang cocok berdasarkan mata uang sumber
3. **Pengambilan Nilai Tukar**: Mengambil nilai tukar dari Frankfurter API
4. **Perhitungan**: Mengalikan jumlah sumber dengan nilai tukar
5. **Penyimpanan**: Menyimpan nilai yang telah dikonversi dengan kode mata uang target

### Alur Contoh

```graphql
# 1. Set value in source CURRENCY field
mutation SetSourceValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "source_currency_field_id"
    number: 1000
    currency: "USD"
  })
}

# 2. CURRENCY_CONVERSION fields automatically update
# If you have USD→EUR and USD→GBP options configured,
# both conversion fields will calculate and store their values
```

## Konversi Berdasarkan Tanggal

### Menggunakan Tanggal Saat Ini

```graphql
mutation CreateRealtimeConversion {
  createCustomField(input: {
    name: "Current EUR Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "currentDate"
  })
}
```

Konversi diperbarui dengan nilai tukar saat ini setiap kali nilai sumber berubah.

### Menggunakan Tanggal Tertentu

```graphql
mutation CreateFixedDateConversion {
  createCustomField(input: {
    name: "Year-End 2023 Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2023-12-31T00:00:00Z"
  })
}
```

Selalu menggunakan nilai tukar dari tanggal yang ditentukan.

### Menggunakan Tanggal dari Bidang

```graphql
mutation CreateDateFieldConversion {
  createCustomField(input: {
    name: "Value at Contract Date"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "fromDateField"
    conversionDate: "contract_date_field_id"  # ID of a DATE custom field
  })
}
```

Menggunakan tanggal dari bidang lain (baik tanggal jatuh tempo todo atau bidang kustom DATE).

## Bidang Respons

### Respons TodoCustomField

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang konversi |
| `number` | Float | Jumlah yang telah dikonversi |
| `currency` | String | Kode mata uang target |
| `todo` | Todo! | Rekor yang dimiliki nilai ini |
| `createdAt` | DateTime! | Ketika nilai dibuat |
| `updatedAt` | DateTime! | Ketika nilai terakhir diperbarui |

## Sumber Nilai Tukar

Blue menggunakan **Frankfurter API** untuk nilai tukar:
- API sumber terbuka yang dihosting oleh Bank Sentral Eropa
- Memperbarui setiap hari dengan nilai tukar resmi
- Mendukung nilai tukar historis sejak 1999
- Gratis dan dapat diandalkan untuk penggunaan bisnis

## Penanganan Kesalahan

### Kegagalan Konversi

Ketika konversi gagal (kesalahan API, mata uang tidak valid, dll.):
- Nilai yang dikonversi diatur menjadi `0`
- Kode mata uang target tetap disimpan
- Tidak ada kesalahan yang ditampilkan kepada pengguna

### Skenario Umum

| Skenario | Hasil |
|----------|-------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| Tidak ada opsi yang cocok | Uses "Any" option if available |
| Missing source value | Tidak ada konversi yang dilakukan |

## Izin yang Diperlukan

Manajemen bidang kustom memerlukan akses tingkat proyek:

| Peran | Dapat Membuat/Memperbarui Bidang |
|-------|---------------------------------|
| `OWNER` | ✅ Ya |
| `ADMIN` | ✅ Ya |
| `MEMBER` | ❌ Tidak |
| `CLIENT` | ❌ Tidak |

Izin tampilan untuk nilai yang dikonversi mengikuti aturan akses rekaman standar.

## Praktik Terbaik

### Konfigurasi Opsi
- Buat pasangan mata uang tertentu untuk konversi umum
- Tambahkan opsi cadangan "Any" untuk fleksibilitas
- Gunakan judul deskriptif untuk opsi

### Pemilihan Strategi Tanggal
- Gunakan `currentDate` untuk pelacakan keuangan langsung
- Gunakan `specificDate` untuk pelaporan historis
- Gunakan `fromDateField` untuk nilai tukar spesifik transaksi

### Pertimbangan Kinerja
- Beberapa bidang konversi diperbarui secara paralel
- Panggilan API hanya dilakukan ketika nilai sumber berubah
- Konversi dengan mata uang yang sama melewatkan panggilan API

## Kasus Penggunaan Umum

1. **Proyek Multi-Mata Uang**
   - Lacak biaya proyek dalam mata uang lokal
   - Laporkan total anggaran dalam mata uang perusahaan
   - Bandingkan nilai di berbagai wilayah

2. **Penjualan Internasional**
   - Konversi nilai kesepakatan ke mata uang pelaporan
   - Lacak pendapatan dalam beberapa mata uang
   - Konversi historis untuk kesepakatan yang ditutup

3. **Pelaporan Keuangan**
   - Konversi mata uang akhir periode
   - Laporan keuangan konsolidasi
   - Anggaran vs. aktual dalam mata uang lokal

4. **Manajemen Kontrak**
   - Konversi nilai kontrak pada tanggal penandatanganan
   - Lacak jadwal pembayaran dalam beberapa mata uang
   - Penilaian risiko mata uang

## Batasan

- Tidak ada dukungan untuk konversi cryptocurrency
- Tidak dapat mengatur nilai yang dikonversi secara manual (selalu dihitung)
- Presisi tetap 2 tempat desimal untuk semua jumlah yang dikonversi
- Tidak ada dukungan untuk nilai tukar kustom
- Tidak ada caching nilai tukar (panggilan API segar untuk setiap konversi)
- Bergantung pada ketersediaan Frankfurter API

## Sumber Daya Terkait

- [Bidang Mata Uang](/api/custom-fields/currency) - Bidang sumber untuk konversi
- [Bidang Tanggal](/api/custom-fields/date) - Untuk konversi berbasis tanggal
- [Bidang Formula](/api/custom-fields/formula) - Perhitungan alternatif
- [Ikhtisar Bidang Kustom](/custom-fields/list-custom-fields) - Konsep umum
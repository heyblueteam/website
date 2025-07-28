---
title: Bidang Kustom Mata Uang
description: Buat bidang mata uang untuk melacak nilai moneter dengan format dan validasi yang tepat
---

Bidang kustom mata uang memungkinkan Anda untuk menyimpan dan mengelola nilai moneter dengan kode mata uang yang terkait. Bidang ini mendukung 72 mata uang yang berbeda termasuk mata uang fiat utama dan cryptocurrency, dengan format otomatis dan batasan min/max opsional.

## Contoh Dasar

Buat bidang mata uang sederhana:

```graphql
mutation CreateCurrencyField {
  createCustomField(input: {
    name: "Budget"
    type: CURRENCY
    projectId: "proj_123"
    currency: "USD"
  }) {
    id
    name
    type
    currency
  }
}
```

## Contoh Lanjutan

Buat bidang mata uang dengan batasan validasi:

```graphql
mutation CreateConstrainedCurrencyField {
  createCustomField(input: {
    name: "Deal Value"
    type: CURRENCY
    projectId: "proj_123"
    currency: "EUR"
    min: 0
    max: 1000000
    description: "Estimated deal value in euros"
    isActive: true
  }) {
    id
    name
    type
    currency
    min
    max
    description
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang mata uang |
| `type` | CustomFieldType! | ✅ Ya | Harus `CURRENCY` |
| `currency` | String | Tidak | Kode mata uang default (kode ISO 3 huruf) |
| `min` | Float | Tidak | Nilai minimum yang diizinkan (disimpan tetapi tidak diterapkan pada pembaruan) |
| `max` | Float | Tidak | Nilai maksimum yang diizinkan (disimpan tetapi tidak diterapkan pada pembaruan) |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

**Catatan**: Konteks proyek ditentukan secara otomatis dari autentikasi Anda. Anda harus memiliki akses ke proyek tempat Anda membuat bidang.

## Mengatur Nilai Mata Uang

Untuk mengatur atau memperbarui nilai mata uang pada sebuah catatan:

```graphql
mutation SetCurrencyValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 1500.50
    currency: "USD"
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom mata uang |
| `number` | Float! | ✅ Ya | Jumlah moneter |
| `currency` | String! | ✅ Ya | Kode mata uang 3 huruf |

## Membuat Catatan dengan Nilai Mata Uang

Saat membuat catatan baru dengan nilai mata uang:

```graphql
mutation CreateRecordWithCurrency {
  createTodo(input: {
    title: "Q4 Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "currency_field_id"
      value: "25000.00"
      currency: "GBP"
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
      number
      currency
    }
  }
}
```

### Format Input untuk Membuat

Saat membuat catatan, nilai mata uang diteruskan dengan cara yang berbeda:

| Parameter | Tipe | Deskripsi |
|-----------|------|-----------|
| `customFieldId` | String! | ID dari bidang mata uang |
| `value` | String! | Jumlah sebagai string (misalnya, "1500.50") |
| `currency` | String! | Kode mata uang 3 huruf |

## Mata Uang yang Didukung

Blue mendukung 72 mata uang termasuk 70 mata uang fiat dan 2 cryptocurrency:

### Mata Uang Fiat

#### Amerika
| Mata Uang | Kode | Nama |
|-----------|------|------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | Bolívar Soberano Venezuela |
| Boliviano Bolivia | `BOB` | Boliviano Bolivia |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### Eropa
| Mata Uang | Kode | Nama |
|-----------|------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| Krone Norwegia | `NOK` | Krone Norwegia |
| Danish Krone | `DKK` | Danish Krone |
| Polish Złoty | `PLN` | Polish Złoty |
| Czech Koruna | `CZK` | Czech Koruna |
| Hungarian Forint | `HUF` | Hungarian Forint |
| Romanian Leu | `RON` | Romanian Leu |
| Bulgarian Lev | `BGN` | Bulgarian Lev |
| Turkish Lira | `TRY` | Turkish Lira |
| Ukrainian Hryvnia | `UAH` | Ukrainian Hryvnia |
| Russian Ruble | `RUB` | Russian Ruble |
| Georgian Lari | `GEL` | Georgian Lari |
| Icelandic króna | `ISK` | Icelandic króna |
| Bosnia-Herzegovina Mark | `BAM` | Bosnia-Herzegovina Convertible Mark |

#### Asia-Pasifik
| Mata Uang | Kode | Nama |
|-----------|------|------|
| Japanese Yen | `JPY` | Yen |
| Chinese Yuan | `CNY` | Yuan |
| Hong Kong Dollar | `HKD` | Hong Kong Dollar |
| Singapore Dollar | `SGD` | Singapore Dollar |
| Australian Dollar | `AUD` | Australian Dollar |
| New Zealand Dollar | `NZD` | New Zealand Dollar |
| South Korean Won | `KRW` | South Korean Won |
| Indian Rupee | `INR` | Indian Rupee |
| Indonesian Rupiah | `IDR` | Indonesian Rupiah |
| Thai Baht | `THB` | Thai Baht |
| Malaysian Ringgit | `MYR` | Malaysian Ringgit |
| Philippine Peso | `PHP` | Philippine Peso |
| Vietnamese Dong | `VND` | Vietnamese Dong |
| Taiwanese Dollar | `TWD` | New Taiwan Dollar |
| Pakistani Rupee | `PKR` | Pakistani Rupee |
| Sri Lankan Rupee | `LKR` | Sri Lankan Rupee |
| Cambodian Riel | `KHR` | Cambodian Riel |
| Kazakhstani Tenge | `KZT` | Kazakhstani Tenge |

#### Timur Tengah & Afrika
| Mata Uang | Kode | Nama |
|-----------|------|------|
| UAE Dirham | `AED` | UAE Dirham |
| Saudi Riyal | `SAR` | Saudi Riyal |
| Kuwaiti Dinar | `KWD` | Kuwaiti Dinar |
| Bahraini Dinar | `BHD` | Bahraini Dinar |
| Qatari Riyal | `QAR` | Qatari Riyal |
| Israeli Shekel | `ILS` | Israeli New Shekel |
| Egyptian Pound | `EGP` | Egyptian Pound |
| Moroccan Dirham | `MAD` | Moroccan Dirham |
| Tunisian Dinar | `TND` | Tunisian Dinar |
| South African Rand | `ZAR` | South African Rand |
| Kenyan Shilling | `KES` | Kenyan Shilling |
| Nigerian Naira | `NGN` | Nigerian Naira |
| Ghanaian Cedi | `GHS` | Ghanaian Cedi |
| Zambian Kwacha | `ZMW` | Zambian Kwacha |
| Malagasy Ariary | `MGA` | Malagasy Ariary |

### Cryptocurrency
| Mata Uang | Kode |
|-----------|------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## Bidang Respons

### Respon TodoCustomField

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang kustom |
| `number` | Float | Jumlah moneter |
| `currency` | String | Kode mata uang 3 huruf |
| `todo` | Todo! | Catatan yang nilai ini miliki |
| `createdAt` | DateTime! | Ketika nilai dibuat |
| `updatedAt` | DateTime! | Ketika nilai terakhir dimodifikasi |

## Format Mata Uang

Sistem secara otomatis memformat nilai mata uang berdasarkan lokal:

- **Penempatan simbol**: Menempatkan simbol mata uang dengan benar (sebelum/setelah)
- **Pemisah desimal**: Menggunakan pemisah khusus lokal (. atau ,)
- **Pemisah ribuan**: Menerapkan pengelompokan yang sesuai
- **Tempat desimal**: Menampilkan 0-2 tempat desimal berdasarkan jumlah
- **Penanganan khusus**: USD/CAD menunjukkan awalan kode mata uang untuk kejelasan

### Contoh Format

| Nilai | Mata Uang | Tampilan |
|-------|-----------|----------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## Validasi

### Validasi Jumlah
- Harus merupakan angka yang valid
- Batasan min/max disimpan dengan definisi bidang tetapi tidak diterapkan selama pembaruan nilai
- Mendukung hingga 2 tempat desimal untuk tampilan (presisi penuh disimpan secara internal)

### Validasi Kode Mata Uang
- Harus salah satu dari 72 kode mata uang yang didukung
- Sensitif terhadap huruf besar (gunakan huruf kapital)
- Kode yang tidak valid mengembalikan kesalahan

## Fitur Integrasi

### Formula
Bidang mata uang dapat digunakan dalam bidang kustom FORMULA untuk perhitungan:
- Menjumlahkan beberapa bidang mata uang
- Menghitung persentase
- Melakukan operasi aritmatika

### Konversi Mata Uang
Gunakan bidang CURRENCY_CONVERSION untuk secara otomatis mengonversi antara mata uang (lihat [Bidang Konversi Mata Uang](/api/custom-fields/currency-conversion))

### Automasi
Nilai mata uang dapat memicu automasi berdasarkan:
- Ambang jumlah
- Jenis mata uang
- Perubahan nilai

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|---------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**Catatan**: Meskipun anggota proyek mana pun dapat membuat bidang kustom, kemampuan untuk mengatur nilai tergantung pada izin berbasis peran yang dikonfigurasi untuk setiap bidang.

## Respons Kesalahan

### Nilai Mata Uang Tidak Valid
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

Kesalahan ini terjadi ketika:
- Kode mata uang bukan salah satu dari 72 kode yang didukung
- Format angka tidak valid
- Nilai tidak dapat diparsing dengan benar

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

### Pemilihan Mata Uang
- Tetapkan mata uang default yang sesuai dengan pasar utama Anda
- Gunakan kode mata uang ISO 4217 secara konsisten
- Pertimbangkan lokasi pengguna saat memilih default

### Batasan Nilai
- Tetapkan nilai min/max yang wajar untuk mencegah kesalahan entri data
- Gunakan 0 sebagai minimum untuk bidang yang tidak boleh negatif
- Pertimbangkan kasus penggunaan Anda saat menetapkan maksimum

### Proyek Multi-Mata Uang
- Gunakan mata uang dasar yang konsisten untuk pelaporan
- Implementasikan bidang CURRENCY_CONVERSION untuk konversi otomatis
- Dokumentasikan mata uang mana yang harus digunakan untuk setiap bidang

## Kasus Penggunaan Umum

1. **Penganggaran Proyek**
   - Pelacakan anggaran proyek
   - Perkiraan biaya
   - Pelacakan pengeluaran

2. **Penjualan & Kesepakatan**
   - Nilai kesepakatan
   - Jumlah kontrak
   - Pelacakan pendapatan

3. **Perencanaan Keuangan**
   - Jumlah investasi
   - Putaran pendanaan
   - Target keuangan

4. **Bisnis Internasional**
   - Penetapan harga multi-mata uang
   - Pelacakan valuta asing
   - Transaksi lintas batas

## Batasan

- Maksimum 2 tempat desimal untuk tampilan (meskipun lebih banyak presisi disimpan)
- Tidak ada konversi mata uang bawaan dalam bidang CURRENCY standar
- Tidak dapat mencampur mata uang dalam satu nilai bidang
- Tidak ada pembaruan otomatis untuk nilai tukar (gunakan CURRENCY_CONVERSION untuk ini)
- Simbol mata uang tidak dapat disesuaikan

## Sumber Daya Terkait

- [Bidang Konversi Mata Uang](/api/custom-fields/currency-conversion) - Konversi mata uang otomatis
- [Bidang Angka](/api/custom-fields/number) - Untuk nilai numerik non-moneter
- [Bidang Formula](/api/custom-fields/formula) - Hitung dengan nilai mata uang
- [Bidang Kustom Daftar](/api/custom-fields/list-custom-fields) - Kuery semua bidang kustom dalam sebuah proyek
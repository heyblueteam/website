---
title: Bidang Kustom Telepon
description: Buat bidang telepon untuk menyimpan dan memvalidasi nomor telepon dengan format internasional
---

Bidang kustom telepon memungkinkan Anda untuk menyimpan nomor telepon dalam catatan dengan validasi bawaan dan format internasional. Mereka ideal untuk melacak informasi kontak, kontak darurat, atau data terkait telepon lainnya dalam proyek Anda.

## Contoh Dasar

Buat bidang telepon sederhana:

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang telepon dengan deskripsi:

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
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
| `name` | String! | ✅ Ya | Nama tampilan dari bidang telepon |
| `type` | CustomFieldType! | ✅ Ya | Harus `PHONE` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

**Catatan**: Bidang kustom secara otomatis diasosiasikan dengan proyek berdasarkan konteks proyek pengguna saat ini. Tidak ada parameter `projectId` yang diperlukan.

## Mengatur Nilai Telepon

Untuk mengatur atau memperbarui nilai telepon pada catatan:

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom telepon |
| `text` | String | Tidak | Nomor telepon dengan kode negara |
| `regionCode` | String | Tidak | Kode negara (secara otomatis terdeteksi) |

**Catatan**: Meskipun `text` bersifat opsional dalam skema, nomor telepon diperlukan agar bidang tersebut bermakna. Saat menggunakan `setTodoCustomField`, tidak ada validasi yang dilakukan - Anda dapat menyimpan nilai teks apa pun dan regionCode. Deteksi otomatis hanya terjadi selama pembuatan catatan.

## Membuat Catatan dengan Nilai Telepon

Saat membuat catatan baru dengan nilai telepon:

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
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
      regionCode
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
| `text` | String | Nomor telepon yang diformat (format internasional) |
| `regionCode` | String | Kode negara (misalnya, "US", "GB", "CA") |
| `todo` | Todo! | Catatan yang nilai ini miliki |
| `createdAt` | DateTime! | Kapan nilai tersebut dibuat |
| `updatedAt` | DateTime! | Kapan nilai tersebut terakhir dimodifikasi |

## Validasi Nomor Telepon

**Penting**: Validasi dan format nomor telepon hanya terjadi saat membuat catatan baru melalui `createTodo`. Saat memperbarui nilai telepon yang ada menggunakan `setTodoCustomField`, tidak ada validasi yang dilakukan dan nilai disimpan sesuai yang diberikan.

### Format yang Diterima (Selama Pembuatan Catatan)
Nomor telepon harus menyertakan kode negara dalam salah satu format ini:

- **Format E.164 (diutamakan)**: `+12345678900`
- **Format internasional**: `+1 234 567 8900`
- **Internasional dengan tanda baca**: `+1 (234) 567-8900`
- **Kode negara dengan tanda hubung**: `+1-234-567-8900`

**Catatan**: Format nasional tanpa kode negara (seperti `(234) 567-8900`) akan ditolak selama pembuatan catatan.

### Aturan Validasi (Selama Pembuatan Catatan)
- Menggunakan libphonenumber-js untuk parsing dan validasi
- Menerima berbagai format nomor telepon internasional
- Secara otomatis mendeteksi negara dari nomor
- Memformat nomor dalam format tampilan internasional (misalnya, `+1 234 567 8900`)
- Mengekstrak dan menyimpan kode negara secara terpisah (misalnya, `US`)

### Contoh Telepon yang Valid
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### Contoh Telepon yang Tidak Valid
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## Format Penyimpanan

Saat membuat catatan dengan nomor telepon:
- **text**: Disimpan dalam format internasional (misalnya, `+1 234 567 8900`) setelah validasi
- **regionCode**: Disimpan sebagai kode negara ISO (misalnya, `US`, `GB`, `CA`) yang terdeteksi secara otomatis

Saat memperbarui melalui `setTodoCustomField`:
- **text**: Disimpan persis seperti yang diberikan (tanpa format)
- **regionCode**: Disimpan persis seperti yang diberikan (tanpa validasi)

## Izin yang Diperlukan

| Aksi | Izin yang Diperlukan |
|------|----------------------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## Respons Kesalahan

### Format Telepon Tidak Valid
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
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
    "message": "Custom field not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Kode Negara Hilang
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Praktik Terbaik

### Entri Data
- Selalu sertakan kode negara dalam nomor telepon
- Gunakan format E.164 untuk konsistensi
- Validasi nomor sebelum menyimpan untuk operasi penting
- Pertimbangkan preferensi regional untuk format tampilan

### Kualitas Data
- Simpan nomor dalam format internasional untuk kompatibilitas global
- Gunakan regionCode untuk fitur spesifik negara
- Validasi nomor telepon sebelum operasi kritis (SMS, panggilan)
- Pertimbangkan implikasi zona waktu untuk penjadwalan kontak

### Pertimbangan Internasional
- Kode negara secara otomatis terdeteksi dan disimpan
- Nomor diformat dalam standar internasional
- Preferensi tampilan regional dapat menggunakan regionCode
- Pertimbangkan konvensi dial lokal saat menampilkan

## Kasus Penggunaan Umum

1. **Manajemen Kontak**
   - Nomor telepon klien
   - Informasi kontak vendor
   - Nomor telepon anggota tim
   - Detail kontak dukungan

2. **Kontak Darurat**
   - Nomor kontak darurat
   - Informasi kontak yang siap sedia
   - Kontak respons krisis
   - Nomor telepon eskalasi

3. **Dukungan Pelanggan**
   - Nomor telepon pelanggan
   - Nomor panggilan balik dukungan
   - Nomor telepon verifikasi
   - Nomor kontak tindak lanjut

4. **Penjualan & Pemasaran**
   - Nomor telepon prospek
   - Daftar kontak kampanye
   - Informasi kontak mitra
   - Telepon sumber rujukan

## Fitur Integrasi

### Dengan Automasi
- Memicu tindakan saat bidang telepon diperbarui
- Mengirim notifikasi SMS ke nomor telepon yang disimpan
- Membuat tugas tindak lanjut berdasarkan perubahan telepon
- Mengarahkan panggilan berdasarkan data nomor telepon

### Dengan Pencarian
- Mengacu pada data telepon dari catatan lain
- Mengagregasi daftar telepon dari berbagai sumber
- Menemukan catatan berdasarkan nomor telepon
- Melakukan cross-reference informasi kontak

### Dengan Formulir
- Validasi telepon otomatis
- Pemeriksaan format internasional
- Deteksi kode negara
- Umpan balik format waktu nyata

## Batasan

- Memerlukan kode negara untuk semua nomor
- Tidak ada kemampuan SMS atau panggilan bawaan
- Tidak ada verifikasi nomor telepon di luar pemeriksaan format
- Tidak ada penyimpanan metadata telepon (operator, tipe, dll.)
- Nomor format nasional tanpa kode negara ditolak
- Tidak ada pemformatan nomor telepon otomatis di UI di luar standar internasional

## Sumber Daya Terkait

- [Bidang Teks](/api/custom-fields/text-single) - Untuk data teks non-telepon
- [Bidang Email](/api/custom-fields/email) - Untuk alamat email
- [Bidang URL](/api/custom-fields/url) - Untuk alamat situs web
- [Ikhtisar Bidang Kustom](/custom-fields/list-custom-fields) - Konsep umum
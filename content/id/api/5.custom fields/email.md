---
title: Bidang Kustom Email
description: Buat bidang email untuk menyimpan dan memvalidasi alamat email
---

Bidang kustom email memungkinkan Anda untuk menyimpan alamat email dalam catatan dengan validasi bawaan. Mereka ideal untuk melacak informasi kontak, email penugasan, atau data terkait email lainnya dalam proyek Anda.

## Contoh Dasar

Buat bidang email sederhana:

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang email dengan deskripsi:

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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
| `name` | String! | ✅ Ya | Nama tampilan dari bidang email |
| `type` | CustomFieldType! | ✅ Ya | Harus `EMAIL` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

## Mengatur Nilai Email

Untuk mengatur atau memperbarui nilai email pada catatan:

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom email |
| `text` | String | Tidak | Alamat email yang akan disimpan |

## Membuat Catatan dengan Nilai Email

Saat membuat catatan baru dengan nilai email:

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## Bidang Respons

### CustomField Response

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | ID! | Pengidentifikasi unik untuk bidang kustom |
| `name` | String! | Nama tampilan dari bidang email |
| `type` | CustomFieldType! | Tipe bidang (EMAIL) |
| `description` | String | Teks bantuan untuk bidang |
| `value` | JSON | Berisi nilai email (lihat di bawah) |
| `createdAt` | DateTime! | Ketika bidang dibuat |
| `updatedAt` | DateTime! | Ketika bidang terakhir dimodifikasi |

**Penting**: Nilai email diakses melalui bidang `customField.value.text`, bukan langsung pada respons.

## Menanyakan Nilai Email

Saat menanyakan catatan dengan bidang kustom email, akses email melalui jalur `customField.value.text`:

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

Respons akan menyertakan email dalam struktur bersarang:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## Validasi Email

### Validasi Formulir
Saat bidang email digunakan dalam formulir, mereka secara otomatis memvalidasi format email:
- Menggunakan aturan validasi email standar
- Memangkas spasi dari input
- Menolak format email yang tidak valid

### Aturan Validasi
- Harus mengandung simbol `@`
- Harus memiliki format domain yang valid
- Spasi di awal/akhir secara otomatis dihapus
- Format email umum diterima

### Contoh Email Valid
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### Contoh Email Tidak Valid
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## Catatan Penting

### API Langsung vs Formulir
- **Formulir**: Validasi email otomatis diterapkan
- **API Langsung**: Tidak ada validasi - teks apa pun dapat disimpan
- **Rekomendasi**: Gunakan formulir untuk input pengguna untuk memastikan validasi

### Format Penyimpanan
- Alamat email disimpan sebagai teks biasa
- Tidak ada pemformatan atau penguraian khusus
- Sensitivitas huruf: Bidang kustom EMAIL disimpan dengan sensitif huruf (berbeda dengan email otentikasi pengguna yang dinormalisasi menjadi huruf kecil)
- Tidak ada batasan panjang maksimum di luar batasan basis data (batas 16MB)

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|----------------------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## Respons Kesalahan

### Format Email Tidak Valid (Hanya Formulir)
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
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
      "code": "NOT_FOUND"
    }
  }]
}
```

## Praktik Terbaik

### Entri Data
- Selalu validasi alamat email dalam aplikasi Anda
- Gunakan bidang email hanya untuk alamat email yang sebenarnya
- Pertimbangkan menggunakan formulir untuk input pengguna agar mendapatkan validasi otomatis

### Kualitas Data
- Pangkas spasi sebelum menyimpan
- Pertimbangkan normalisasi huruf (biasanya huruf kecil)
- Validasi format email sebelum operasi penting

### Pertimbangan Privasi
- Alamat email disimpan sebagai teks biasa
- Pertimbangkan regulasi privasi data (GDPR, CCPA)
- Terapkan kontrol akses yang sesuai

## Kasus Penggunaan Umum

1. **Manajemen Kontak**
   - Alamat email klien
   - Informasi kontak vendor
   - Email anggota tim
   - Detail kontak dukungan

2. **Manajemen Proyek**
   - Email pemangku kepentingan
   - Email kontak persetujuan
   - Penerima notifikasi
   - Email kolaborator eksternal

3. **Dukungan Pelanggan**
   - Alamat email pelanggan
   - Kontak tiket dukungan
   - Kontak eskalasi
   - Alamat email umpan balik

4. **Penjualan & Pemasaran**
   - Alamat email prospek
   - Daftar kontak kampanye
   - Informasi kontak mitra
   - Email sumber rujukan

## Fitur Integrasi

### Dengan Automasi
- Memicu tindakan saat bidang email diperbarui
- Mengirim notifikasi ke alamat email yang disimpan
- Membuat tugas tindak lanjut berdasarkan perubahan email

### Dengan Pencarian
- Referensi data email dari catatan lain
- Mengagregasi daftar email dari beberapa sumber
- Mencari catatan berdasarkan alamat email

### Dengan Formulir
- Validasi email otomatis
- Pemeriksaan format email
- Pemangkasan spasi

## Batasan

- Tidak ada verifikasi atau validasi email bawaan di luar pemeriksaan format
- Tidak ada fitur UI khusus email (seperti tautan email yang dapat diklik)
- Disimpan sebagai teks biasa tanpa enkripsi
- Tidak ada kemampuan komposisi atau pengiriman email
- Tidak ada penyimpanan metadata email (nama tampilan, dll.)
- Panggilan API langsung melewati validasi (hanya formulir yang divalidasi)

## Sumber Daya Terkait

- [Bidang Teks](/api/custom-fields/text-single) - Untuk data teks non-email
- [Bidang URL](/api/custom-fields/url) - Untuk alamat situs web
- [Bidang Telepon](/api/custom-fields/phone) - Untuk nomor telepon
- [Ikhtisar Bidang Kustom](/api/custom-fields/list-custom-fields) - Konsep umum
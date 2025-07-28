---
title: Bidang Kustom Checkbox
description: Buat bidang checkbox boolean untuk data ya/tidak atau benar/salah
---

Bidang kustom checkbox menyediakan input boolean sederhana (benar/salah) untuk tugas. Mereka sempurna untuk pilihan biner, indikator status, atau melacak apakah sesuatu telah diselesaikan.

## Contoh Dasar

Buat bidang checkbox sederhana:

```graphql
mutation CreateCheckboxField {
  createCustomField(input: {
    name: "Reviewed"
    type: CHECKBOX
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang checkbox dengan deskripsi dan validasi:

```graphql
mutation CreateDetailedCheckbox {
  createCustomField(input: {
    name: "Customer Approved"
    type: CHECKBOX
    description: "Check this box when the customer has approved the work"
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
| `name` | String! | ✅ Ya | Nama tampilan dari checkbox |
| `type` | CustomFieldType! | ✅ Ya | Harus `CHECKBOX` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

## Mengatur Nilai Checkbox

Untuk mengatur atau memperbarui nilai checkbox pada tugas:

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

Untuk menghapus centang pada checkbox:

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### Parameter SetTodoCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID tugas yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID bidang kustom checkbox |
| `checked` | Boolean | Tidak | Benar untuk mencentang, salah untuk menghapus centang |

## Membuat Tugas dengan Nilai Checkbox

Saat membuat tugas baru dengan nilai checkbox:

```graphql
mutation CreateTaskWithCheckbox {
  createTodo(input: {
    title: "Review contract"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "checkbox_field_id"
      value: "true"  # Pass as string
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
      checked
    }
  }
}
```

### Nilai String yang Diterima

Saat membuat tugas, nilai checkbox harus disampaikan sebagai string:

| Nilai String | Hasil |
|--------------|-------|
| `"true"` | ✅ Dicentang (case-sensitive) |
| `"1"` | ✅ Dicentang |
| `"checked"` | ✅ Dicentang (case-sensitive) |
| Any other value | ❌ Tidak dicentang |

**Catatan**: Perbandingan string selama pembuatan tugas bersifat case-sensitive. Nilai harus cocok persis dengan `"true"`, `"1"`, atau `"checked"` untuk menghasilkan status dicentang.

## Bidang Respons

### TodoCustomField Response

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | ID! | Pengidentifikasi unik untuk nilai bidang |
| `uid` | String! | Pengidentifikasi unik alternatif |
| `customField` | CustomField! | Definisi bidang kustom |
| `checked` | Boolean | Status checkbox (benar/salah/null) |
| `todo` | Todo! | Tugas yang nilai ini miliki |
| `createdAt` | DateTime! | Kapan nilai tersebut dibuat |
| `updatedAt` | DateTime! | Kapan nilai tersebut terakhir dimodifikasi |

## Integrasi Otomatisasi

Bidang checkbox memicu berbagai peristiwa otomatisasi berdasarkan perubahan status:

| Aksi | Peristiwa yang Dipicu | Deskripsi |
|------|-----------------------|-----------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | Dipicu ketika checkbox dicentang |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | Dipicu ketika checkbox tidak dicentang |

Ini memungkinkan Anda untuk membuat otomatisasi yang merespons perubahan status checkbox, seperti:
- Mengirim pemberitahuan ketika item disetujui
- Memindahkan tugas ketika checkbox tinjauan dicentang
- Memperbarui bidang terkait berdasarkan status checkbox

## Impor/Ekspor Data

### Mengimpor Nilai Checkbox

Saat mengimpor data melalui CSV atau format lainnya:
- `"true"`, `"yes"` → Dicentang (case-insensitive)
- Nilai lainnya (termasuk `"false"`, `"no"`, `"0"`, kosong) → Tidak dicentang

### Mengekspor Nilai Checkbox

Saat mengekspor data:
- Kotak yang dicentang diekspor sebagai `"X"`
- Kotak yang tidak dicentang diekspor sebagai string kosong `""`

## Izin yang Diperlukan

| Aksi | Izin yang Diperlukan |
|------|----------------------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## Respons Kesalahan

### Tipe Nilai Tidak Valid
```json
{
  "errors": [{
    "message": "Invalid value type for checkbox field",
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

## Praktik Terbaik

### Konvensi Penamaan
- Gunakan nama yang jelas dan berorientasi tindakan: "Disetujui", "Ditinjau", "Selesai"
- Hindari nama negatif yang membingungkan pengguna: lebih baik "Aktif" daripada "Tidak Aktif"
- Spesifik tentang apa yang diwakili oleh checkbox

### Kapan Menggunakan Checkbox
- **Pilihan biner**: Ya/Tidak, Benar/Salah, Selesai/Tidak Selesai
- **Indikator status**: Disetujui, Ditinjau, Diterbitkan
- **Bendera fitur**: Memiliki Dukungan Prioritas, Memerlukan Tanda Tangan
- **Pelacakan sederhana**: Email Terkirim, Faktur Dibayar, Item Dikirim

### Kapan TIDAK Menggunakan Checkbox
- Ketika Anda memerlukan lebih dari dua opsi (gunakan SELECT_SINGLE sebagai gantinya)
- Untuk data numerik atau teks (gunakan bidang NUMBER atau TEXT)
- Ketika Anda perlu melacak siapa yang mencentang atau kapan (gunakan log audit)

## Kasus Penggunaan Umum

1. **Alur Kerja Persetujuan**
   - "Manajer Disetujui"
   - "Persetujuan Klien"
   - "Tinjauan Hukum Selesai"

2. **Manajemen Tugas**
   - "Terhambat"
   - "Siap untuk Ditinjau"
   - "Prioritas Tinggi"

3. **Kontrol Kualitas**
   - "QA Lulus"
   - "Dokumentasi Selesai"
   - "Tes Ditulis"

4. **Bendera Administratif**
   - "Faktur Dikirim"
   - "Kontrak Ditandatangani"
   - "Tindak Lanjut Diperlukan"

## Batasan

- Bidang checkbox hanya dapat menyimpan nilai benar/salah (tidak ada tri-state atau null setelah pengaturan awal)
- Tidak ada konfigurasi nilai default (selalu dimulai sebagai null hingga diatur)
- Tidak dapat menyimpan metadata tambahan seperti siapa yang mencentang atau kapan
- Tidak ada visibilitas bersyarat berdasarkan nilai bidang lain

## Sumber Daya Terkait

- [Ikhtisar Bidang Kustom](/api/custom-fields/list-custom-fields) - Konsep umum bidang kustom
- [API Otomatisasi](/api/automations) - Buat otomatisasi yang dipicu oleh perubahan checkbox
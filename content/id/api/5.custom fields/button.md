---
title: Bidang Kustom Tombol
description: Buat bidang tombol interaktif yang memicu otomatisasi saat diklik
---

Bidang kustom tombol menyediakan elemen UI interaktif yang memicu otomatisasi saat diklik. Berbeda dengan jenis bidang kustom lainnya yang menyimpan data, bidang tombol berfungsi sebagai pemicu aksi untuk mengeksekusi alur kerja yang telah dikonfigurasi.

## Contoh Dasar

Buat bidang tombol sederhana yang memicu otomatisasi:

```graphql
mutation CreateButtonField {
  createCustomField(input: {
    name: "Send Invoice"
    type: BUTTON
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat tombol dengan persyaratan konfirmasi:

```graphql
mutation CreateButtonWithConfirmation {
  createCustomField(input: {
    name: "Delete All Attachments"
    type: BUTTON
    projectId: "proj_123"
    buttonType: "hardConfirmation"
    buttonConfirmText: "DELETE"
    description: "Permanently removes all attachments from this task"
  }) {
    id
    name
    type
    buttonType
    buttonConfirmText
    description
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan tombol |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `BUTTON` |
| `projectId` | String! | ✅ Ya | ID proyek tempat bidang akan dibuat |
| `buttonType` | String | Tidak | Perilaku konfirmasi (lihat Jenis Tombol di bawah) |
| `buttonConfirmText` | String | Tidak | Teks yang harus diketik pengguna untuk konfirmasi keras |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |
| `required` | Boolean | Tidak | Apakah bidang diperlukan (defaultnya false) |
| `isActive` | Boolean | Tidak | Apakah bidang aktif (defaultnya true) |

### Jenis Bidang Tombol

Bidang `buttonType` adalah string bebas yang dapat digunakan oleh klien UI untuk menentukan perilaku konfirmasi. Nilai umum termasuk:

- `""` (kosong) - Tidak ada konfirmasi
- `"soft"` - Dialog konfirmasi sederhana
- `"hard"` - Memerlukan teks konfirmasi yang diketik

**Catatan**: Ini hanya petunjuk UI. API tidak memvalidasi atau menegakkan nilai tertentu.

## Memicu Klik Tombol

Untuk memicu klik tombol dan mengeksekusi otomatisasi terkait:

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### Parameter Input Klik

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID tugas yang berisi tombol |
| `customFieldId` | String! | ✅ Ya | ID bidang kustom tombol |

### Penting: Perilaku API

**Semua klik tombol melalui API dieksekusi segera** terlepas dari pengaturan `buttonType` atau `buttonConfirmText`. Bidang ini disimpan untuk klien UI untuk menerapkan dialog konfirmasi, tetapi API itu sendiri:

- Tidak memvalidasi teks konfirmasi
- Tidak menegakkan persyaratan konfirmasi
- Menjalankan aksi tombol segera saat dipanggil

Konfirmasi murni merupakan fitur keamanan UI di sisi klien.

### Contoh: Mengklik Jenis Tombol yang Berbeda

```graphql
# Button with no confirmation
mutation ClickSimpleButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "simple_button_id"
  })
}

# Button with soft confirmation (API call is the same!)
mutation ClickSoftConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "soft_confirm_button_id"
  })
}

# Button with hard confirmation (API call is still the same!)
mutation ClickHardConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "hard_confirm_button_id"
  })
}
```

Ketiga mutasi di atas akan mengeksekusi aksi tombol segera saat dipanggil melalui API, melewati persyaratan konfirmasi apa pun.

## Bidang Respons

### Respons Bidang Kustom

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk bidang kustom |
| `name` | String! | Nama tampilan tombol |
| `type` | CustomFieldType! | Selalu `BUTTON` untuk bidang tombol |
| `buttonType` | String | Pengaturan perilaku konfirmasi |
| `buttonConfirmText` | String | Teks konfirmasi yang diperlukan (jika menggunakan konfirmasi keras) |
| `description` | String | Teks bantuan untuk pengguna |
| `required` | Boolean! | Apakah bidang diperlukan |
| `isActive` | Boolean! | Apakah bidang saat ini aktif |
| `projectId` | String! | ID proyek yang dimiliki bidang ini |
| `createdAt` | DateTime! | Waktu saat bidang dibuat |
| `updatedAt` | DateTime! | Waktu saat bidang terakhir dimodifikasi |

## Cara Kerja Bidang Tombol

### Integrasi Otomatisasi

Bidang tombol dirancang untuk bekerja dengan sistem otomatisasi Blue:

1. **Buat bidang tombol** menggunakan mutasi di atas
2. **Konfigurasi otomatisasi** yang mendengarkan peristiwa `CUSTOM_FIELD_BUTTON_CLICKED`
3. **Pengguna mengklik tombol** di UI
4. **Otomatisasi mengeksekusi** aksi yang telah dikonfigurasi

### Alur Peristiwa

Saat tombol diklik:

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### Tidak Ada Penyimpanan Data

Penting: Bidang tombol tidak menyimpan data nilai apa pun. Mereka murni berfungsi sebagai pemicu aksi. Setiap klik:
- Menghasilkan peristiwa
- Memicu otomatisasi terkait
- Mencatat aksi dalam riwayat tugas
- Tidak memodifikasi nilai bidang apa pun

## Izin yang Diperlukan

Pengguna memerlukan peran proyek yang sesuai untuk membuat dan menggunakan bidang tombol:

| Aksi | Peran yang Diperlukan |
|------|-----------------------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## Respons Kesalahan

### Izin Ditolak
```json
{
  "errors": [{
    "message": "You don't have permission to edit this custom field",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Bidang Kustom Tidak Ditemukan
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

**Catatan**: API tidak mengembalikan kesalahan spesifik untuk otomatisasi yang hilang atau ketidakcocokan konfirmasi.

## Praktik Terbaik

### Konvensi Penamaan
- Gunakan nama yang berorientasi aksi: "Kirim Faktur", "Buat Laporan", "Beri Tahu Tim"
- Spesifik tentang apa yang dilakukan tombol
- Hindari nama generik seperti "Tombol 1" atau "Klik Di Sini"

### Pengaturan Konfirmasi
- Biarkan `buttonType` kosong untuk aksi yang aman dan dapat dibalik
- Atur `buttonType` untuk menyarankan perilaku konfirmasi kepada klien UI
- Gunakan `buttonConfirmText` untuk menentukan apa yang harus diketik pengguna dalam konfirmasi UI
- Ingat: Ini hanya petunjuk UI - panggilan API selalu dieksekusi segera

### Desain Otomatisasi
- Pertahankan aksi tombol terfokus pada satu alur kerja
- Berikan umpan balik yang jelas tentang apa yang terjadi setelah mengklik
- Pertimbangkan untuk menambahkan teks deskripsi untuk menjelaskan tujuan tombol

## Kasus Penggunaan Umum

1. **Transisi Alur Kerja**
   - "Tandai sebagai Lengkap"
   - "Kirim untuk Persetujuan"
   - "Arsipkan Tugas"

2. **Integrasi Eksternal**
   - "Sinkronkan ke CRM"
   - "Hasilkan Faktur"
   - "Kirim Pembaruan Email"

3. **Operasi Batch**
   - "Perbarui Semua Subtugas"
   - "Salin ke Proyek"
   - "Terapkan Template"

4. **Tindakan Pelaporan**
   - "Hasilkan Laporan"
   - "Ekspor Data"
   - "Buat Ringkasan"

## Batasan

- Tombol tidak dapat menyimpan atau menampilkan nilai data
- Setiap tombol hanya dapat memicu otomatisasi, bukan panggilan API langsung (namun, otomatisasi dapat mencakup aksi permintaan HTTP untuk memanggil API eksternal atau API Blue sendiri)
- Visibilitas tombol tidak dapat dikendalikan secara kondisional
- Maksimal satu eksekusi otomatisasi per klik (meskipun otomatisasi tersebut dapat memicu beberapa aksi)

## Sumber Daya Terkait

- [API Otomatisasi](/api/automations/index) - Konfigurasi aksi yang dipicu oleh tombol
- [Ikhtisar Bidang Kustom](/custom-fields/list-custom-fields) - Konsep umum bidang kustom
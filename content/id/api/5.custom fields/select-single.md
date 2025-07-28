---
title: Bidang Kustom Pilihan Tunggal
description: Buat bidang pilihan tunggal untuk memungkinkan pengguna memilih satu opsi dari daftar yang telah ditentukan
---

Bidang kustom pilihan tunggal memungkinkan pengguna untuk memilih tepat satu opsi dari daftar yang telah ditentukan. Mereka ideal untuk bidang status, kategori, prioritas, atau skenario apa pun di mana hanya satu pilihan yang harus dibuat dari sekumpulan opsi yang terkontrol.

## Contoh Dasar

Buat bidang pilihan tunggal yang sederhana:

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang pilihan tunggal dengan opsi yang telah ditentukan:

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang pilihan tunggal |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `SELECT_SINGLE` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | Tidak | Opsi awal untuk bidang |

### CreateCustomFieldOptionInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `title` | String! | ✅ Ya | Teks tampilan untuk opsi |
| `color` | String | Tidak | Kode warna hex untuk opsi |

## Menambahkan Opsi ke Bidang yang Ada

Tambahkan opsi baru ke bidang pilihan tunggal yang ada:

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## Mengatur Nilai Pilihan Tunggal

Untuk mengatur opsi yang dipilih pada sebuah catatan:

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom pilihan tunggal |
| `customFieldOptionId` | String | Tidak | ID dari opsi yang dipilih (diutamakan untuk pilihan tunggal) |
| `customFieldOptionIds` | [String!] | Tidak | Array dari ID opsi (menggunakan elemen pertama untuk pilihan tunggal) |

## Menanyakan Nilai Pilihan Tunggal

Tanyakan nilai pilihan tunggal dari sebuah catatan:

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

Bidang `value` mengembalikan objek JSON dengan rincian opsi yang dipilih.

## Membuat Catatan dengan Nilai Pilihan Tunggal

Saat membuat catatan baru dengan nilai pilihan tunggal:

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
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
      value  # Contains the selected option object
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
| `value` | JSON | Berisi objek opsi yang dipilih dengan id, judul, warna, posisi |
| `todo` | Todo! | Catatan yang menjadi milik nilai ini |
| `createdAt` | DateTime! | Saat nilai tersebut dibuat |
| `updatedAt` | DateTime! | Saat nilai tersebut terakhir dimodifikasi |

### CustomFieldOption Respons

| Bidang | Tipe | Deskripsi |
|-------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk opsi |
| `title` | String! | Teks tampilan untuk opsi |
| `color` | String | Kode warna hex untuk representasi visual |
| `position` | Float | Urutan sortir untuk opsi |
| `customField` | CustomField! | Bidang kustom yang menjadi milik opsi ini |

### CustomField Respons

| Bidang | Tipe | Deskripsi |
|-------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk bidang |
| `name` | String! | Nama tampilan dari bidang pilihan tunggal |
| `type` | CustomFieldType! | Selalu `SELECT_SINGLE` |
| `description` | String | Teks bantuan untuk bidang |
| `customFieldOptions` | [CustomFieldOption!] | Semua opsi yang tersedia |

## Format Nilai

### Format Input
- **Parameter API**: Gunakan `customFieldOptionId` untuk ID opsi tunggal
- **Alternatif**: Gunakan array `customFieldOptionIds` (mengambil elemen pertama)
- **Menghapus Pilihan**: Hilangkan kedua bidang atau kirim nilai kosong

### Format Output
- **Respons GraphQL**: Objek JSON dalam bidang `value` yang berisi {id, title, color, position}
- **Log Aktivitas**: Judul opsi sebagai string
- **Data Otomatisasi**: Judul opsi sebagai string

## Perilaku Pemilihan

### Pemilihan Eksklusif
- Mengatur opsi baru secara otomatis menghapus pilihan sebelumnya
- Hanya satu opsi yang dapat dipilih pada satu waktu
- Mengatur `null` atau nilai kosong menghapus pilihan

### Logika Cadangan
- Jika array `customFieldOptionIds` disediakan, hanya opsi pertama yang digunakan
- Ini memastikan kompatibilitas dengan format input multi-pilihan
- Array kosong atau nilai null menghapus pilihan

## Mengelola Opsi

### Memperbarui Properti Opsi
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
  }) {
    id
    title
    color
  }
}
```

### Menghapus Opsi
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

**Catatan**: Menghapus opsi akan menghapusnya dari semua catatan di mana opsi tersebut dipilih.

### Mengurutkan Ulang Opsi
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## Aturan Validasi

### Validasi Opsi
- ID opsi yang diberikan harus ada
- Opsi harus menjadi milik bidang kustom yang ditentukan
- Hanya satu opsi yang dapat dipilih (ditegakkan secara otomatis)
- Nilai null/kosong adalah valid (tidak ada pilihan)

### Validasi Bidang
- Harus memiliki setidaknya satu opsi yang ditentukan agar dapat digunakan
- Judul opsi harus unik dalam bidang
- Kode warna harus dalam format hex yang valid (jika disediakan)

## Izin yang Diperlukan

| Aksi | Izin yang Diperlukan |
|------|----------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## Respons Kesalahan

### ID Opsi Tidak Valid
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### Opsi Tidak Milik Bidang
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
    "extensions": {
      "code": "VALIDATION_ERROR"
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

### Tidak Dapat Mengurai Nilai
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

## Praktik Terbaik

### Desain Opsi
- Gunakan judul opsi yang jelas dan deskriptif
- Terapkan pengkodean warna yang bermakna
- Jaga daftar opsi tetap fokus dan relevan
- Urutkan opsi secara logis (berdasarkan prioritas, frekuensi, dll.)

### Pola Bidang Status
- Gunakan alur kerja status yang konsisten di seluruh proyek
- Pertimbangkan perkembangan alami dari opsi
- Sertakan status akhir yang jelas (Selesai, Dibatalkan, dll.)
- Gunakan warna yang mencerminkan makna opsi

### Manajemen Data
- Tinjau dan bersihkan opsi yang tidak terpakai secara berkala
- Gunakan konvensi penamaan yang konsisten
- Pertimbangkan dampak penghapusan opsi pada catatan yang ada
- Rencanakan pembaruan dan migrasi opsi

## Kasus Penggunaan Umum

1. **Status dan Alur Kerja**
   - Status tugas (Untuk Dilakukan, Dalam Proses, Selesai)
   - Status persetujuan (Menunggu, Disetujui, Ditolak)
   - Fase proyek (Perencanaan, Pengembangan, Pengujian, Dirilis)
   - Status penyelesaian masalah

2. **Klasifikasi dan Kategorisasi**
   - Tingkat prioritas (Rendah, Sedang, Tinggi, Kritis)
   - Jenis tugas (Bug, Fitur, Peningkatan, Dokumentasi)
   - Kategori proyek (Internal, Klien, Penelitian)
   - Penugasan departemen

3. **Kualitas dan Penilaian**
   - Status tinjauan (Belum Dimulai, Dalam Tinjauan, Disetujui)
   - Penilaian kualitas (Buruk, Cukup, Baik, Sangat Baik)
   - Tingkat risiko (Rendah, Sedang, Tinggi)
   - Tingkat kepercayaan

4. **Penugasan dan Kepemilikan**
   - Penugasan tim
   - Kepemilikan departemen
   - Penugasan berbasis peran
   - Penugasan regional

## Fitur Integrasi

### Dengan Otomatisasi
- Memicu tindakan saat opsi tertentu dipilih
- Mengarahkan pekerjaan berdasarkan kategori yang dipilih
- Mengirim pemberitahuan untuk perubahan status
- Membuat alur kerja bersyarat berdasarkan pilihan

### Dengan Pencarian
- Menyaring catatan berdasarkan opsi yang dipilih
- Merujuk data opsi dari catatan lain
- Membuat laporan berdasarkan pilihan opsi
- Mengelompokkan catatan berdasarkan nilai yang dipilih

### Dengan Formulir
- Kontrol input dropdown
- Antarmuka tombol radio
- Validasi dan penyaringan opsi
- Tampilan bidang bersyarat berdasarkan pilihan

## Pelacakan Aktivitas

Perubahan bidang pilihan tunggal secara otomatis dilacak:
- Menunjukkan pilihan opsi lama dan baru
- Menampilkan judul opsi dalam log aktivitas
- Stempel waktu untuk semua perubahan pilihan
- Atribusi pengguna untuk modifikasi

## Perbedaan dari Multi-Pilih

| Fitur | Pilihan Tunggal | Pilihan Multi |
|-------|-----------------|----------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## Pembatasan

- Hanya satu opsi yang dapat dipilih pada satu waktu
- Tidak ada struktur opsi hierarkis atau bersarang
- Opsi dibagikan di semua catatan yang menggunakan bidang
- Tidak ada analitik atau pelacakan penggunaan opsi bawaan
- Kode warna hanya untuk tampilan, tidak ada dampak fungsional
- Tidak dapat menetapkan izin yang berbeda per opsi

## Sumber Daya Terkait

- [Bidang Multi-Pilih](/api/custom-fields/select-multi) - Untuk pilihan ganda
- [Bidang Kotak Centang](/api/custom-fields/checkbox) - Untuk pilihan boolean sederhana
- [Bidang Teks](/api/custom-fields/text-single) - Untuk input teks bebas
- [Ikhtisar Bidang Kustom](/api/custom-fields/1.index) - Konsep umum
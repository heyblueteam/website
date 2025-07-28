---
title: Bidang Kustom Multi-Pilih
description: Buat bidang multi-pilih untuk memungkinkan pengguna memilih beberapa opsi dari daftar yang telah ditentukan
---

Bidang kustom multi-pilih memungkinkan pengguna untuk memilih beberapa opsi dari daftar yang telah ditentukan. Mereka ideal untuk kategori, tag, keterampilan, fitur, atau skenario apa pun di mana beberapa pilihan diperlukan dari sekumpulan opsi yang terkontrol.

## Contoh Dasar

Buat bidang multi-pilih sederhana:

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang multi-pilih dan kemudian tambahkan opsi secara terpisah:

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang multi-pilih |
| `type` | CustomFieldType! | ✅ Ya | Harus `SELECT_MULTI` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |
| `projectId` | String! | ✅ Ya | ID proyek untuk bidang ini |

### CreateCustomFieldOptionInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom |
| `title` | String! | ✅ Ya | Teks tampilan untuk opsi |
| `color` | String | Tidak | Warna untuk opsi (string apa pun) |
| `position` | Float | Tidak | Urutan sortir untuk opsi |

## Menambahkan Opsi ke Bidang yang Ada

Tambahkan opsi baru ke bidang multi-pilih yang sudah ada:

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## Mengatur Nilai Multi-Pilih

Untuk mengatur beberapa opsi yang dipilih pada sebuah catatan:

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom multi-pilih |
| `customFieldOptionIds` | [String!] | ✅ Ya | Array dari ID opsi yang akan dipilih |

## Membuat Catatan dengan Nilai Multi-Pilih

Saat membuat catatan baru dengan nilai multi-pilih:

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
    }
  }
}
```

## Bidang Respon

### Respon TodoCustomField

| Bidang | Tipe | Deskripsi |
|-------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang kustom |
| `selectedOptions` | [CustomFieldOption!] | Array dari opsi yang dipilih |
| `todo` | Todo! | Catatan yang dimiliki nilai ini |
| `createdAt` | DateTime! | Ketika nilai dibuat |
| `updatedAt` | DateTime! | Ketika nilai terakhir dimodifikasi |

### Respon CustomFieldOption

| Bidang | Tipe | Deskripsi |
|-------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk opsi |
| `title` | String! | Teks tampilan untuk opsi |
| `color` | String | Kode warna hex untuk representasi visual |
| `position` | Float | Urutan sortir untuk opsi |
| `customField` | CustomField! | Bidang kustom yang dimiliki opsi ini |

### Respon CustomField

| Bidang | Tipe | Deskripsi |
|-------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk bidang |
| `name` | String! | Nama tampilan dari bidang multi-pilih |
| `type` | CustomFieldType! | Selalu `SELECT_MULTI` |
| `description` | String | Teks bantuan untuk bidang |
| `customFieldOptions` | [CustomFieldOption!] | Semua opsi yang tersedia |

## Format Nilai

### Format Input
- **Parameter API**: Array dari ID opsi (`["option1", "option2", "option3"]`)
- **Format String**: ID opsi yang dipisahkan koma (`"option1,option2,option3"`)

### Format Output
- **Respon GraphQL**: Array dari objek CustomFieldOption
- **Log Aktivitas**: Judul opsi yang dipisahkan koma
- **Data Automasi**: Array dari judul opsi

## Mengelola Opsi

### Memperbarui Properti Opsi
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
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

### Mengatur Ulang Urutan Opsi
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## Aturan Validasi

### Validasi Opsi
- Semua ID opsi yang diberikan harus ada
- Opsi harus milik bidang kustom yang ditentukan
- Hanya bidang SELECT_MULTI yang dapat memiliki beberapa opsi yang dipilih
- Array kosong adalah valid (tidak ada pilihan)

### Validasi Bidang
- Harus memiliki setidaknya satu opsi yang didefinisikan agar dapat digunakan
- Judul opsi harus unik dalam bidang
- Bidang warna menerima nilai string apa pun (tanpa validasi hex)

## Izin yang Diperlukan

| Aksi | Izin yang Diperlukan |
|------|---------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## Respon Kesalahan

### ID Opsi Tidak Valid
```json
{
  "errors": [{
    "message": "Custom field option not found",
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
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Beberapa Opsi pada Bidang Non-Multi
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Praktik Terbaik

### Desain Opsi
- Gunakan judul opsi yang deskriptif dan ringkas
- Terapkan skema pengkodean warna yang konsisten
- Jaga daftar opsi tetap dapat dikelola (biasanya 3-20 opsi)
- Urutkan opsi secara logis (secara alfabet, berdasarkan frekuensi, dll.)

### Manajemen Data
- Tinjau dan bersihkan opsi yang tidak digunakan secara berkala
- Gunakan konvensi penamaan yang konsisten di seluruh proyek
- Pertimbangkan penggunaan kembali opsi saat membuat bidang
- Rencanakan untuk pembaruan dan migrasi opsi

### Pengalaman Pengguna
- Berikan deskripsi bidang yang jelas
- Gunakan warna untuk meningkatkan perbedaan visual
- Kelompokkan opsi terkait bersama-sama
- Pertimbangkan pilihan default untuk kasus umum

## Kasus Penggunaan Umum

1. **Manajemen Proyek**
   - Kategori dan tag tugas
   - Tingkat dan jenis prioritas
   - Penugasan anggota tim
   - Indikator status

2. **Manajemen Konten**
   - Kategori dan topik artikel
   - Tipe dan format konten
   - Saluran publikasi
   - Alur kerja persetujuan

3. **Dukungan Pelanggan**
   - Kategori dan jenis masalah
   - Produk atau layanan yang terpengaruh
   - Metode penyelesaian
   - Segmen pelanggan

4. **Pengembangan Produk**
   - Kategori fitur
   - Persyaratan teknis
   - Lingkungan pengujian
   - Saluran rilis

## Fitur Integrasi

### Dengan Automasi
- Memicu tindakan saat opsi tertentu dipilih
- Mengarahkan pekerjaan berdasarkan kategori yang dipilih
- Mengirim notifikasi untuk pilihan prioritas tinggi
- Membuat tugas tindak lanjut berdasarkan kombinasi opsi

### Dengan Pencarian
- Menyaring catatan berdasarkan opsi yang dipilih
- Mengagregasi data dari pilihan opsi
- Merujuk data opsi dari catatan lain
- Membuat laporan berdasarkan kombinasi opsi

### Dengan Formulir
- Kontrol input multi-pilih
- Validasi dan penyaringan opsi
- Memuat opsi secara dinamis
- Menampilkan bidang bersyarat

## Pelacakan Aktivitas

Perubahan bidang multi-pilih secara otomatis dilacak:
- Menunjukkan opsi yang ditambahkan dan dihapus
- Menampilkan judul opsi dalam log aktivitas
- Stempel waktu untuk semua perubahan pilihan
- Atribusi pengguna untuk modifikasi

## Batasan

- Batas praktis maksimum opsi tergantung pada kinerja UI
- Tidak ada struktur opsi hierarkis atau bersarang
- Opsi dibagikan di semua catatan yang menggunakan bidang
- Tidak ada analitik atau pelacakan penggunaan opsi bawaan
- Bidang warna menerima string apa pun (tanpa validasi hex)
- Tidak dapat mengatur izin berbeda per opsi
- Opsi harus dibuat secara terpisah, tidak inline dengan pembuatan bidang
- Tidak ada mutasi urutan ulang khusus (gunakan editCustomFieldOption dengan posisi)

## Sumber Daya Terkait

- [Bidang Pilihan Tunggal](/api/custom-fields/select-single) - Untuk pilihan tunggal
- [Bidang Checkbox](/api/custom-fields/checkbox) - Untuk pilihan boolean sederhana
- [Bidang Teks](/api/custom-fields/text-single) - Untuk input teks bebas
- [Ikhtisar Bidang Kustom](/api/custom-fields/2.list-custom-fields) - Konsep umum
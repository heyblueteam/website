---
title: Bidang Kustom File
description: Buat bidang file untuk melampirkan dokumen, gambar, dan file lainnya ke catatan
---

Bidang kustom file memungkinkan Anda untuk melampirkan beberapa file ke catatan. File disimpan dengan aman di AWS S3 dengan pelacakan metadata yang komprehensif, validasi jenis file, dan kontrol akses yang tepat.

## Contoh Dasar

Buat bidang file sederhana:

```graphql
mutation CreateFileField {
  createCustomField(input: {
    name: "Attachments"
    type: FILE
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang file dengan deskripsi:

```graphql
mutation CreateDetailedFileField {
  createCustomField(input: {
    name: "Project Documents"
    type: FILE
    description: "Upload project-related documents, images, and files"
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
| `name` | String! | ✅ Ya | Nama tampilan dari bidang file |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `FILE` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

**Catatan**: Bidang kustom secara otomatis diasosiasikan dengan proyek berdasarkan konteks proyek pengguna saat ini. Tidak ada parameter `projectId` yang diperlukan.

## Proses Unggah File

### Langkah 1: Unggah File

Pertama, unggah file untuk mendapatkan UID file:

```graphql
mutation UploadFile {
  uploadFile(input: {
    file: $file  # File upload variable
    companyId: "company_123"
    projectId: "proj_123"
  }) {
    id
    uid
    name
    size
    type
    extension
    status
  }
}
```

### Langkah 2: Lampirkan File ke Catatan

Kemudian lampirkan file yang diunggah ke catatan:

```graphql
mutation AttachFileToRecord {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "file_field_456"
    fileUid: "file_uid_from_upload"
  }) {
    id
    file {
      uid
      name
      size
      type
    }
  }
}
```

## Mengelola Lampiran File

### Menambahkan File Tunggal

```graphql
mutation AddFileToField {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  }) {
    id
    position
    file {
      uid
      name
      size
      type
      extension
    }
  }
}
```

### Menghapus File

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### Operasi File Massal

Perbarui beberapa file sekaligus menggunakan customFieldOptionIds:

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## Parameter Input Unggah File

### UploadFileInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `file` | Upload! | ✅ Ya | File yang akan diunggah |
| `companyId` | String! | ✅ Ya | ID perusahaan untuk penyimpanan file |
| `projectId` | String | Tidak | ID proyek untuk file spesifik proyek |

### Parameter Input Manajemen File

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom file |
| `fileUid` | String! | ✅ Ya | Pengidentifikasi unik dari file yang diunggah |

## Penyimpanan File dan Batasan

### Batas Ukuran File

| Jenis Batas | Ukuran |
|-------------|--------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### Jenis File yang Didukung

#### Gambar
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### Video
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### Audio
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### Dokumen
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### Arsip
- `zip`, `rar`, `7z`, `tar`, `gz`

#### Kode/Teks
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### Arsitektur Penyimpanan

- **Penyimpanan**: AWS S3 dengan struktur folder yang terorganisir
- **Format Jalur**: `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **Keamanan**: URL yang ditandatangani untuk akses yang aman
- **Cadangan**: Redundansi S3 otomatis

## Bidang Respon

### Respon File

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | ID! | ID database |
| `uid` | String! | Pengidentifikasi file unik |
| `name` | String! | Nama file asli |
| `size` | Float! | Ukuran file dalam byte |
| `type` | String! | Tipe MIME |
| `extension` | String! | Ekstensi file |
| `status` | FileStatus | PENDING atau CONFIRMED (nullable) |
| `shared` | Boolean! | Apakah file dibagikan |
| `createdAt` | DateTime! | Timestamp unggah |

### Respon TodoCustomFieldFile

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | ID! | ID catatan persimpangan |
| `uid` | String! | Pengidentifikasi unik |
| `position` | Float! | Urutan tampilan |
| `file` | File! | Objek file yang terkait |
| `todoCustomField` | TodoCustomField! | Bidang kustom induk |
| `createdAt` | DateTime! | Kapan file dilampirkan |

## Membuat Catatan dengan File

Saat membuat catatan, Anda dapat melampirkan file menggunakan UID mereka:

```graphql
mutation CreateRecordWithFiles {
  createTodo(input: {
    title: "Project deliverables"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "file_field_id"
      customFieldOptionIds: ["file_uid_1", "file_uid_2"]
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
      todoCustomFieldFiles {
        id
        position
        file {
          uid
          name
          size
          type
        }
      }
    }
  }
}
```

## Validasi dan Keamanan File

### Validasi Unggah

- **Pemeriksaan Tipe MIME**: Memvalidasi terhadap jenis yang diizinkan
- **Validasi Ekstensi File**: Cadangan untuk `application/octet-stream`
- **Batas Ukuran**: Diterapkan saat waktu unggah
- **Sanitasi Nama File**: Menghapus karakter khusus

### Kontrol Akses

- **Izin Unggah**: Keanggotaan proyek/perusahaan diperlukan
- **Asosiasi File**: Peran ADMIN, OWNER, MEMBER, CLIENT
- **Akses File**: Diwarisi dari izin proyek/perusahaan
- **URL Aman**: URL yang ditandatangani dengan batas waktu untuk akses file

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|---------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## Respon Kesalahan

### File Terlalu Besar
```json
{
  "errors": [{
    "message": "File \"filename.pdf\": Size exceeds maximum limit of 256MB",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### File Tidak Ditemukan
```json
{
  "errors": [{
    "message": "File not found",
    "extensions": {
      "code": "FILE_NOT_FOUND"
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

### Manajemen File
- Unggah file sebelum melampirkannya ke catatan
- Gunakan nama file yang deskriptif
- Organisir file berdasarkan proyek/tujuan
- Bersihkan file yang tidak digunakan secara berkala

### Kinerja
- Unggah file dalam batch jika memungkinkan
- Gunakan format file yang sesuai untuk jenis konten
- Kompres file besar sebelum diunggah
- Pertimbangkan persyaratan pratinjau file

### Keamanan
- Validasi konten file, bukan hanya ekstensi
- Gunakan pemindaian virus untuk file yang diunggah
- Terapkan kontrol akses yang tepat
- Pantau pola unggah file

## Kasus Penggunaan Umum

1. **Manajemen Dokumen**
   - Spesifikasi proyek
   - Kontrak dan perjanjian
   - Catatan rapat dan presentasi
   - Dokumentasi teknis

2. **Manajemen Aset**
   - File desain dan mockup
   - Aset merek dan logo
   - Materi pemasaran
   - Gambar produk

3. **Kepatuhan dan Catatan**
   - Dokumen hukum
   - Jejak audit
   - Sertifikat dan lisensi
   - Catatan keuangan

4. **Kolaborasi**
   - Sumber daya bersama
   - Dokumen yang dikendalikan versi
   - Umpan balik dan anotasi
   - Materi referensi

## Fitur Integrasi

### Dengan Automasi
- Memicu tindakan saat file ditambahkan/dihapus
- Memproses file berdasarkan jenis atau metadata
- Mengirim notifikasi untuk perubahan file
- Mengarsipkan file berdasarkan kondisi

### Dengan Gambar Sampul
- Gunakan bidang file sebagai sumber gambar sampul
- Pemrosesan gambar otomatis dan thumbnail
- Pembaruan sampul dinamis saat file berubah

### Dengan Pencarian
- Referensi file dari catatan lain
- Menghitung dan mengagregasi ukuran file
- Temukan catatan berdasarkan metadata file
- Referensi silang lampiran file

## Batasan

- Maksimum 256MB per file
- Bergantung pada ketersediaan S3
- Tidak ada versi file bawaan
- Tidak ada konversi file otomatis
- Kemampuan pratinjau file terbatas
- Tidak ada pengeditan kolaboratif waktu nyata

## Sumber Daya Terkait

- [API Unggah File](/api/upload-files) - Titik akhir unggah file
- [Ikhtisar Bidang Kustom](/api/custom-fields/list-custom-fields) - Konsep umum
- [API Automasi](/api/automations) - Automasi berbasis file
- [Dokumentasi AWS S3](https://docs.aws.amazon.com/s3/) - Backend penyimpanan
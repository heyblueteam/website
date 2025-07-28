---
title: Referensi Bidang Kustom
description: Buat bidang referensi yang menghubungkan ke catatan di proyek lain untuk hubungan lintas proyek
---

Bidang kustom referensi memungkinkan Anda untuk membuat tautan antara catatan di proyek yang berbeda, memungkinkan hubungan lintas proyek dan berbagi data. Mereka menyediakan cara yang kuat untuk menghubungkan pekerjaan terkait di seluruh struktur proyek organisasi Anda.

## Contoh Dasar

Buat bidang referensi sederhana:

```graphql
mutation CreateReferenceField {
  createCustomField(input: {
    name: "Related Project"
    type: REFERENCE
    referenceProjectId: "proj_456"
    description: "Link to related project records"
  }) {
    id
    name
    type
    referenceProjectId
  }
}
```

## Contoh Lanjutan

Buat bidang referensi dengan penyaringan dan pemilihan ganda:

```graphql
mutation CreateFilteredReferenceField {
  createCustomField(input: {
    name: "Dependencies"
    type: REFERENCE
    referenceProjectId: "proj_456"
    referenceMultiple: true
    referenceFilter: {
      status: ACTIVE
      tags: ["dependency"]
    }
    description: "Select multiple dependency records from the project"
  }) {
    id
    name
    type
    referenceProjectId
    referenceMultiple
    referenceFilter
  }
}
```

## Parameter Input

### CreateCustomFieldInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang referensi |
| `type` | CustomFieldType! | ✅ Ya | Harus `REFERENCE` |
| `referenceProjectId` | String | Tidak | ID proyek yang dirujuk |
| `referenceMultiple` | Boolean | Tidak | Izinkan pemilihan catatan ganda (default: false) |
| `referenceFilter` | TodoFilterInput | Tidak | Kriteria penyaringan untuk catatan yang dirujuk |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

**Catatan**: Bidang kustom secara otomatis diasosiasikan dengan proyek berdasarkan konteks proyek pengguna saat ini.

## Konfigurasi Referensi

### Referensi Tunggal vs Ganda

**Referensi Tunggal (default):**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- Pengguna dapat memilih satu catatan dari proyek yang dirujuk
- Mengembalikan satu objek Todo

**Referensi Ganda:**
```graphql
{
  referenceMultiple: true
}
```
- Pengguna dapat memilih beberapa catatan dari proyek yang dirujuk
- Mengembalikan array objek Todo

### Penyaringan Referensi

Gunakan `referenceFilter` untuk membatasi catatan mana yang dapat dipilih:

```graphql
{
  referenceFilter: {
    assigneeIds: ["user_123"]
    tagIds: ["tag_123"]
    dueStart: "2024-01-01"
    dueEnd: "2024-12-31"
    showCompleted: false
  }
}
```

## Mengatur Nilai Referensi

### Referensi Tunggal

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### Referensi Ganda

```graphql
mutation SetMultipleReferences {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: [
      "referenced_todo_789",
      "referenced_todo_012",
      "referenced_todo_345"
    ]
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom referensi |
| `customFieldReferenceTodoIds` | [String!] | ✅ Ya | Array dari ID catatan yang dirujuk |

## Membuat Catatan dengan Referensi

```graphql
mutation CreateRecordWithReference {
  createTodo(input: {
    title: "Implementation Task"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "reference_field_id"
      value: "referenced_todo_789"
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
      selectedTodos {
        id
        title
        status
      }
    }
  }
}
```

## Bidang Respons

### TodoCustomField Respons

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | ID! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang referensi |
| `todo` | Todo! | Catatan yang nilai ini miliki |
| `createdAt` | DateTime! | Saat nilai dibuat |
| `updatedAt` | DateTime! | Saat nilai terakhir dimodifikasi |

**Catatan**: Todo yang dirujuk diakses melalui `customField.selectedTodos`, bukan langsung di TodoCustomField.

### Bidang Todo yang Dirujuk

Setiap Todo yang dirujuk mencakup:

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | ID! | Pengidentifikasi unik dari catatan yang dirujuk |
| `title` | String! | Judul dari catatan yang dirujuk |
| `status` | TodoStatus! | Status saat ini (AKTIF, SELESAI, dll.) |
| `description` | String | Deskripsi dari catatan yang dirujuk |
| `dueDate` | DateTime | Tanggal jatuh tempo jika ditetapkan |
| `assignees` | [User!] | Pengguna yang ditugaskan |
| `tags` | [Tag!] | Tag yang terkait |
| `project` | Project! | Proyek yang berisi catatan yang dirujuk |

## Menanyakan Data Referensi

### Kuery Dasar

```graphql
query GetRecordsWithReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        selectedTodos {
          id
          title
          status
          project {
            id
            name
          }
        }
      }
    }
  }
}
```

### Kuery Lanjutan dengan Data Bersarang

```graphql
query GetDetailedReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        referenceProjectId
        referenceMultiple
      }
      selectedTodos {
        id
        title
        description
        status
        dueDate
        assignees {
          id
          name
          email
        }
        tags {
          id
          name
          color
        }
        project {
          id
          name
        }
      }
    }
  }
}
```

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|----------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**Penting**: Pengguna harus memiliki izin tampilan pada proyek yang dirujuk untuk melihat catatan yang terhubung.

## Akses Lintas Proyek

### Visibilitas Proyek

- Pengguna hanya dapat merujuk catatan dari proyek yang mereka akses
- Catatan yang dirujuk menghormati izin proyek asli
- Perubahan pada catatan yang dirujuk muncul secara real-time
- Menghapus catatan yang dirujuk menghilangkannya dari bidang referensi

### Warisan Izin

- Bidang referensi mewarisi izin dari kedua proyek
- Pengguna perlu akses tampilan ke proyek yang dirujuk
- Izin edit didasarkan pada aturan proyek saat ini
- Data yang dirujuk bersifat hanya-baca dalam konteks bidang referensi

## Respons Kesalahan

### Proyek Referensi Tidak Valid

```json
{
  "errors": [{
    "message": "Project not found",
    "extensions": {
      "code": "PROJECT_NOT_FOUND"
    }
  }]
}
```

### Catatan yang Dirujuk Tidak Ditemukan

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

### Izin Ditolak

```json
{
  "errors": [{
    "message": "Forbidden",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

## Praktik Terbaik

### Desain Bidang

1. **Penamaan yang Jelas** - Gunakan nama deskriptif yang menunjukkan hubungan
2. **Penyaringan yang Tepat** - Atur filter untuk menampilkan hanya catatan yang relevan
3. **Pertimbangkan izin** - Pastikan pengguna memiliki akses ke proyek yang dirujuk
4. **Dokumentasikan hubungan** - Berikan deskripsi yang jelas tentang koneksi

### Pertimbangan Kinerja

1. **Batasi ruang lingkup referensi** - Gunakan filter untuk mengurangi jumlah catatan yang dapat dipilih
2. **Hindari nesting dalam** - Jangan buat rantai referensi yang kompleks
3. **Pertimbangkan caching** - Data yang dirujuk disimpan dalam cache untuk kinerja
4. **Pantau penggunaan** - Lacak bagaimana referensi digunakan di seluruh proyek

### Integritas Data

1. **Tangani penghapusan** - Rencanakan untuk saat catatan yang dirujuk dihapus
2. **Validasi izin** - Pastikan pengguna dapat mengakses proyek yang dirujuk
3. **Perbarui ketergantungan** - Pertimbangkan dampak saat mengubah catatan yang dirujuk
4. **Jejak audit** - Lacak hubungan referensi untuk kepatuhan

## Kasus Penggunaan Umum

### Ketergantungan Proyek

```graphql
# Link to prerequisite tasks in other projects
{
  name: "Prerequisites"
  type: REFERENCE
  referenceProjectId: "infrastructure_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: true
    tagIds: ["prerequisite_tag_id"]
  }
}
```

### Persyaratan Klien

```graphql
# Reference client requirements from a requirements project
{
  name: "Client Requirements"
  type: REFERENCE
  referenceProjectId: "requirements_project"
  referenceFilter: {
    assigneeIds: ["client_user_id"]
    showCompleted: false
  }
}
```

### Alokasi Sumber Daya

```graphql
# Link to resource records in a resource management project
{
  name: "Assigned Resources"
  type: REFERENCE
  referenceProjectId: "resources_project"
  referenceMultiple: true
  referenceFilter: {
    tagIds: ["available_tag_id"]
  }
}
```

### Jaminan Kualitas

```graphql
# Reference QA test cases from a testing project
{
  name: "Test Cases"
  type: REFERENCE
  referenceProjectId: "qa_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: false
    tagIds: ["test_case_tag_id"]
  }
}
```

## Integrasi dengan Pencarian

Bidang referensi bekerja dengan [Bidang Pencarian](/api/custom-fields/lookup) untuk menarik data dari catatan yang dirujuk. Bidang pencarian dapat mengekstrak nilai dari catatan yang dipilih di bidang referensi, tetapi mereka hanya extractor data (tidak ada fungsi agregasi seperti SUM yang didukung).

```graphql
# Reference field links to records
{
  name: "Related Tasks"
  type: REFERENCE
  referenceProjectId: "other_project"
}

# Lookup field extracts data from referenced records
{
  name: "Task Status"
  type: LOOKUP
  lookupOption: {
    customFieldId: "related_tasks_field_id"
    targetField: "status"
  }
}
```

## Batasan

- Proyek yang dirujuk harus dapat diakses oleh pengguna
- Perubahan pada izin proyek yang dirujuk mempengaruhi akses bidang referensi
- Nesting dalam referensi dapat mempengaruhi kinerja
- Tidak ada validasi bawaan untuk referensi melingkar
- Tidak ada pembatasan otomatis yang mencegah referensi proyek yang sama
- Validasi filter tidak diterapkan saat mengatur nilai referensi

## Sumber Daya Terkait

- [Bidang Pencarian](/api/custom-fields/lookup) - Ekstrak data dari catatan yang dirujuk
- [API Proyek](/api/projects) - Mengelola proyek yang berisi referensi
- [API Catatan](/api/records) - Bekerja dengan catatan yang memiliki referensi
- [Ikhtisar Bidang Kustom](/api/custom-fields/list-custom-fields) - Konsep umum
---
title: Bidang Kustom Durasi Waktu
description: Buat bidang durasi waktu yang dihitung yang melacak waktu antara peristiwa dalam alur kerja Anda
---

Bidang kustom Durasi Waktu secara otomatis menghitung dan menampilkan durasi antara dua peristiwa dalam alur kerja Anda. Mereka ideal untuk melacak waktu pemrosesan, waktu respons, waktu siklus, atau metrik berbasis waktu lainnya dalam proyek Anda.

## Contoh Dasar

Buat bidang durasi waktu sederhana yang melacak berapa lama tugas diselesaikan:

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## Contoh Lanjutan

Buat bidang durasi waktu kompleks yang melacak waktu antara perubahan bidang kustom dengan target SLA:

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## Parameter Input

### CreateCustomFieldInput (TIME_DURATION)

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Ya | Nama tampilan dari bidang durasi |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `TIME_DURATION` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ Ya | Cara menampilkan durasi |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ Ya | Konfigurasi peristiwa mulai |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ Ya | Konfigurasi peristiwa akhir |
| `timeDurationTargetTime` | Float | Tidak | Durasi target dalam detik untuk pemantauan SLA |

### CustomFieldTimeDurationInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `type` | CustomFieldTimeDurationType! | ✅ Ya | Tipe peristiwa yang dilacak |
| `condition` | CustomFieldTimeDurationCondition! | ✅ Ya | `FIRST` atau `LAST` kejadian |
| `customFieldId` | String | Conditional | Diperlukan untuk tipe `TODO_CUSTOM_FIELD` |
| `customFieldOptionIds` | [String!] | Conditional | Diperlukan untuk perubahan bidang pilihan |
| `todoListId` | String | Conditional | Diperlukan untuk tipe `TODO_MOVED` |
| `tagId` | String | Conditional | Diperlukan untuk tipe `TODO_TAG_ADDED` |
| `assigneeId` | String | Conditional | Diperlukan untuk tipe `TODO_ASSIGNEE_ADDED` |

### Nilai CustomFieldTimeDurationType

| Nilai | Deskripsi |
|-------|-----------|
| `TODO_CREATED_AT` | Ketika catatan dibuat |
| `TODO_CUSTOM_FIELD` | Ketika nilai bidang kustom berubah |
| `TODO_DUE_DATE` | Ketika tanggal jatuh tempo ditetapkan |
| `TODO_MARKED_AS_COMPLETE` | Ketika catatan ditandai selesai |
| `TODO_MOVED` | Ketika catatan dipindahkan ke daftar yang berbeda |
| `TODO_TAG_ADDED` | Ketika tag ditambahkan ke catatan |
| `TODO_ASSIGNEE_ADDED` | Ketika penugasan ditambahkan ke catatan |

### Nilai CustomFieldTimeDurationCondition

| Nilai | Deskripsi |
|-------|-----------|
| `FIRST` | Gunakan kejadian pertama dari peristiwa |
| `LAST` | Gunakan kejadian terakhir dari peristiwa |

### Nilai CustomFieldTimeDurationDisplayType

| Nilai | Deskripsi | Contoh |
|-------|-----------|---------|
| `FULL_DATE` | Format Hari:Jam:Menit:Detik | `"01:02:03:04"` |
| `FULL_DATE_STRING` | Ditulis lengkap dalam kata-kata | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | Numerik dengan satuan | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | Durasi dalam hari saja | `"2.5"` (2.5 days) |
| `HOURS` | Durasi dalam jam saja | `"60"` (60 hours) |
| `MINUTES` | Durasi dalam menit saja | `"3600"` (3600 minutes) |
| `SECONDS` | Durasi dalam detik saja | `"216000"` (216000 seconds) |

## Bidang Respons

### Respons TodoCustomField

| Field | Tipe | Deskripsi |
|-------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang kustom |
| `number` | Float | Durasi dalam detik |
| `value` | Float | Alias untuk angka (durasi dalam detik) |
| `todo` | Todo! | Catatan yang dimiliki nilai ini |
| `createdAt` | DateTime! | Ketika nilai dibuat |
| `updatedAt` | DateTime! | Ketika nilai terakhir diperbarui |

### Respons CustomField (TIME_DURATION)

| Field | Tipe | Deskripsi |
|-------|------|-----------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | Format tampilan untuk durasi |
| `timeDurationStart` | CustomFieldTimeDuration | Konfigurasi peristiwa mulai |
| `timeDurationEnd` | CustomFieldTimeDuration | Konfigurasi peristiwa akhir |
| `timeDurationTargetTime` | Float | Durasi target dalam detik (untuk pemantauan SLA) |

## Perhitungan Durasi

### Cara Kerjanya
1. **Peristiwa Mulai**: Sistem memantau peristiwa mulai yang ditentukan
2. **Peristiwa Akhir**: Sistem memantau peristiwa akhir yang ditentukan
3. **Perhitungan**: Durasi = Waktu Akhir - Waktu Mulai
4. **Penyimpanan**: Durasi disimpan dalam detik sebagai angka
5. **Tampilan**: Diformat sesuai dengan pengaturan `timeDurationDisplay`

### Pemicu Pembaruan
Nilai durasi secara otomatis dihitung ulang ketika:
- Catatan dibuat atau diperbarui
- Nilai bidang kustom berubah
- Tag ditambahkan atau dihapus
- Penugasan ditambahkan atau dihapus
- Catatan dipindahkan antara daftar
- Catatan ditandai selesai/tidak selesai

## Membaca Nilai Durasi

### Kuery Bidang Durasi
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### Nilai Tampilan Terformat
Nilai durasi secara otomatis diformat berdasarkan pengaturan `timeDurationDisplay`:

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## Contoh Konfigurasi Umum

### Waktu Penyelesaian Tugas
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### Durasi Perubahan Status
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### Waktu dalam Daftar Tertentu
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### Waktu Respons Penugasan
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## Izin yang Diperlukan

| Tindakan | Izin yang Diperlukan |
|----------|---------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## Respons Kesalahan

### Konfigurasi Tidak Valid
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Bidang yang Direferensikan Tidak Ditemukan
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

### Opsi yang Diperlukan Hilang
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Catatan Penting

### Perhitungan Otomatis
- Bidang durasi adalah **hanya-baca** - nilai dihitung secara otomatis
- Anda tidak dapat mengatur nilai durasi secara manual melalui API
- Perhitungan terjadi secara asinkron melalui pekerjaan latar belakang
- Nilai diperbarui secara otomatis ketika peristiwa pemicu terjadi

### Pertimbangan Kinerja
- Perhitungan durasi dimasukkan dalam antrean dan diproses secara asinkron
- Jumlah besar bidang durasi dapat mempengaruhi kinerja
- Pertimbangkan frekuensi peristiwa pemicu saat merancang bidang durasi
- Gunakan kondisi spesifik untuk menghindari perhitungan ulang yang tidak perlu

### Nilai Null
Bidang durasi akan menunjukkan `null` ketika:
- Peristiwa mulai belum terjadi
- Peristiwa akhir belum terjadi
- Konfigurasi merujuk pada entitas yang tidak ada
- Perhitungan mengalami kesalahan

## Praktik Terbaik

### Desain Konfigurasi
- Gunakan tipe peristiwa spesifik daripada yang umum jika memungkinkan
- Pilih kondisi `FIRST` vs `LAST` yang sesuai berdasarkan alur kerja Anda
- Uji perhitungan durasi dengan data contoh sebelum penerapan
- Dokumentasikan logika bidang durasi Anda untuk anggota tim

### Format Tampilan
- Gunakan `FULL_DATE_SUBSTRING` untuk format yang paling mudah dibaca
- Gunakan `FULL_DATE` untuk tampilan lebar yang kompak dan konsisten
- Gunakan `FULL_DATE_STRING` untuk laporan dan dokumen formal
- Gunakan `DAYS`, `HOURS`, `MINUTES`, atau `SECONDS` untuk tampilan numerik sederhana
- Pertimbangkan batasan ruang UI Anda saat memilih format

### Pemantauan SLA dengan Waktu Target
Saat menggunakan `timeDurationTargetTime`:
- Tetapkan durasi target dalam detik
- Bandingkan durasi aktual dengan target untuk kepatuhan SLA
- Gunakan dalam dasbor untuk menyoroti item yang terlambat
- Contoh: SLA respons 24 jam = 86400 detik

### Integrasi Alur Kerja
- Rancang bidang durasi untuk mencocokkan proses bisnis Anda yang sebenarnya
- Gunakan data durasi untuk perbaikan dan optimisasi proses
- Pantau tren durasi untuk mengidentifikasi kemacetan alur kerja
- Siapkan peringatan untuk ambang durasi jika diperlukan

## Kasus Penggunaan Umum

1. **Kinerja Proses**
   - Waktu penyelesaian tugas
   - Waktu siklus tinjauan
   - Waktu pemrosesan persetujuan
   - Waktu respons

2. **Pemantauan SLA**
   - Waktu untuk respons pertama
   - Waktu penyelesaian
   - Kerangka waktu eskalasi
   - Kepatuhan tingkat layanan

3. **Analitik Alur Kerja**
   - Identifikasi kemacetan
   - Optimisasi proses
   - Metrik kinerja tim
   - Penjadwalan jaminan kualitas

4. **Manajemen Proyek**
   - Durasi fase
   - Penjadwalan tonggak
   - Waktu alokasi sumber daya
   - Kerangka waktu pengiriman

## Batasan

- Bidang durasi adalah **hanya-baca** dan tidak dapat diatur secara manual
- Nilai dihitung secara asinkron dan mungkin tidak tersedia segera
- Membutuhkan pemicu peristiwa yang tepat untuk diatur dalam alur kerja Anda
- Tidak dapat menghitung durasi untuk peristiwa yang belum terjadi
- Terbatas pada pelacakan waktu antara peristiwa diskrit (bukan pelacakan waktu terus-menerus)
- Tidak ada peringatan atau notifikasi SLA bawaan
- Tidak dapat mengagregasi beberapa perhitungan durasi ke dalam satu bidang

## Sumber Daya Terkait

- [Bidang Angka](/api/custom-fields/number) - Untuk nilai numerik manual
- [Bidang Tanggal](/api/custom-fields/date) - Untuk pelacakan tanggal tertentu
- [Ikhtisar Bidang Kustom](/api/custom-fields/list-custom-fields) - Konsep umum
- [Automasi](/api/automations) - Untuk memicu tindakan berdasarkan ambang durasi
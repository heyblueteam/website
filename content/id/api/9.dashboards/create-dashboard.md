---
title: Buat Dashboard
description: Buat dashboard baru untuk visualisasi data dan pelaporan di Blue
---

## Buat Dashboard

Mutasi `createDashboard` memungkinkan Anda untuk membuat dashboard baru dalam perusahaan atau proyek Anda. Dashboard adalah alat visualisasi yang kuat yang membantu tim melacak metrik, memantau kemajuan, dan membuat keputusan berbasis data.

### Contoh Dasar

```graphql
mutation CreateDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      title: "Sales Performance Dashboard"
    }
  ) {
    id
    title
    createdBy {
      id
      email
      firstName
      lastName
    }
    createdAt
  }
}
```

### Dashboard Khusus Proyek

Buat dashboard yang terkait dengan proyek tertentu:

```graphql
mutation CreateProjectDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      projectId: "proj_xyz789"
      title: "Q4 Project Metrics"
    }
  ) {
    id
    title
    project {
      id
      name
    }
    createdBy {
      id
      email
    }
    dashboardUsers {
      id
      user {
        id
        email
      }
      role
    }
    createdAt
  }
}
```

## Parameter Input

### CreateDashboardInput

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `companyId` | String! | ✅ Ya | ID perusahaan tempat dashboard akan dibuat |
| `title` | String! | ✅ Ya | Nama dashboard. Harus berupa string yang tidak kosong |
| `projectId` | String | Tidak | ID opsional dari proyek untuk dikaitkan dengan dashboard ini |

## Field Respons

Mutasi mengembalikan objek `Dashboard` yang lengkap:

| Field | Tipe | Deskripsi |
|-------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk dashboard yang dibuat |
| `title` | String! | Judul dashboard sesuai yang diberikan |
| `companyId` | String! | Perusahaan yang dimiliki dashboard ini |
| `projectId` | String | ID proyek yang terkait (jika diberikan) |
| `project` | Project | Objek proyek yang terkait (jika projectId diberikan) |
| `createdBy` | User! | Pengguna yang membuat dashboard (Anda) |
| `dashboardUsers` | [DashboardUser!]! | Daftar pengguna dengan akses (awalnya hanya pencipta) |
| `createdAt` | DateTime! | Timestamp saat dashboard dibuat |
| `updatedAt` | DateTime! | Timestamp modifikasi terakhir (sama dengan createdAt untuk dashboard baru) |

### Field DashboardUser

Saat dashboard dibuat, pencipta secara otomatis ditambahkan sebagai pengguna dashboard:

| Field | Tipe | Deskripsi |
|-------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk hubungan pengguna dashboard |
| `user` | User! | Objek pengguna dengan akses ke dashboard |
| `role` | DashboardRole! | Peran pengguna (pencipta mendapatkan akses penuh) |
| `dashboard` | Dashboard! | Referensi kembali ke dashboard |

## Izin yang Diperlukan

Setiap pengguna yang terautentikasi yang merupakan anggota perusahaan yang ditentukan dapat membuat dashboard. Tidak ada persyaratan peran khusus.

| Status Pengguna | Dapat Membuat Dashboard |
|------------------|-------------------------|
| Company Member | ✅ Ya |
| Anggota Non-Perusahaan | ❌ Tidak |
| Unauthenticated | ❌ Tidak |

## Respons Kesalahan

### Perusahaan Tidak Valid
```json
{
  "errors": [{
    "message": "Company not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Pengguna Tidak Dalam Perusahaan
```json
{
  "errors": [{
    "message": "You don't have access to this company",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Proyek Tidak Valid
```json
{
  "errors": [{
    "message": "Project not found or doesn't belong to the specified company",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Judul Kosong
```json
{
  "errors": [{
    "message": "Dashboard title cannot be empty",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Catatan Penting

- **Kepemilikan otomatis**: Pengguna yang membuat dashboard secara otomatis menjadi pemiliknya dengan izin penuh
- **Asosiasi proyek**: Jika Anda memberikan `projectId`, itu harus milik perusahaan yang sama
- **Izin awal**: Hanya pencipta yang memiliki akses pada awalnya. Gunakan `editDashboard` untuk menambahkan lebih banyak pengguna
- **Persyaratan judul**: Judul dashboard harus berupa string yang tidak kosong. Tidak ada persyaratan keunikan
- **Keanggotaan perusahaan**: Anda harus menjadi anggota perusahaan untuk membuat dashboard di dalamnya

## Alur Kerja Pembuatan Dashboard

1. **Buat dashboard** menggunakan mutasi ini
2. **Konfigurasi grafik dan widget** menggunakan antarmuka pembangun dashboard
3. **Tambahkan anggota tim** menggunakan mutasi `editDashboard` dengan `dashboardUsers`
4. **Atur filter dan rentang tanggal** melalui antarmuka dashboard
5. **Bagikan atau sematkan** dashboard menggunakan ID uniknya

## Kasus Penggunaan

1. **Dashboard eksekutif**: Buat gambaran tingkat tinggi dari metrik perusahaan
2. **Pelacakan proyek**: Bangun dashboard khusus proyek untuk memantau kemajuan
3. **Kinerja tim**: Lacak produktivitas tim dan metrik pencapaian
4. **Pelaporan klien**: Buat dashboard untuk laporan yang dihadapi klien
5. **Pemantauan waktu nyata**: Siapkan dashboard untuk data operasional langsung

## Praktik Terbaik

1. **Konvensi penamaan**: Gunakan judul yang jelas dan deskriptif yang menunjukkan tujuan dashboard
2. **Asosiasi proyek**: Tautkan dashboard ke proyek saat mereka khusus proyek
3. **Manajemen akses**: Tambahkan anggota tim segera setelah pembuatan untuk kolaborasi
4. **Organisasi**: Buat hierarki dashboard menggunakan pola penamaan yang konsisten

## Operasi Terkait

- [Daftar Dashboard](/api/dashboards/) - Ambil semua dashboard untuk perusahaan atau proyek
- [Edit Dashboard](/api/dashboards/rename-dashboard) - Ganti nama dashboard atau kelola pengguna
- [Salin Dashboard](/api/dashboards/copy-dashboard) - Duplikasi dashboard yang ada
- [Hapus Dashboard](/api/dashboards/delete-dashboard) - Hapus dashboard
---
title: Bidang Kustom Lokasi
description: Buat bidang lokasi untuk menyimpan koordinat geografis untuk catatan
---

Bidang kustom lokasi menyimpan koordinat geografis (lintang dan bujur) untuk catatan. Mereka mendukung penyimpanan koordinat yang tepat, kueri geospasial, dan penyaringan berbasis lokasi yang efisien.

## Contoh Dasar

Buat bidang lokasi sederhana:

```graphql
mutation CreateLocationField {
  createCustomField(input: {
    name: "Meeting Location"
    type: LOCATION
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Contoh Lanjutan

Buat bidang lokasi dengan deskripsi:

```graphql
mutation CreateDetailedLocationField {
  createCustomField(input: {
    name: "Office Location"
    type: LOCATION
    projectId: "proj_123"
    description: "Primary office location coordinates"
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
| `name` | String! | ✅ Ya | Nama tampilan dari bidang lokasi |
| `type` | CustomFieldType! | ✅ Ya | Harus berupa `LOCATION` |
| `description` | String | Tidak | Teks bantuan yang ditampilkan kepada pengguna |

## Mengatur Nilai Lokasi

Bidang lokasi menyimpan koordinat lintang dan bujur:

```graphql
mutation SetLocationValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    latitude: 40.7128
    longitude: -74.0060
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Tipe | Diperlukan | Deskripsi |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Ya | ID dari catatan yang akan diperbarui |
| `customFieldId` | String! | ✅ Ya | ID dari bidang kustom lokasi |
| `latitude` | Float | Tidak | Koordinat lintang (-90 hingga 90) |
| `longitude` | Float | Tidak | Koordinat bujur (-180 hingga 180) |

**Catatan**: Meskipun kedua parameter bersifat opsional dalam skema, kedua koordinat diperlukan untuk lokasi yang valid. Jika hanya satu yang diberikan, lokasi akan menjadi tidak valid.

## Validasi Koordinat

### Rentang Valid

| Koordinat | Rentang | Deskripsi |
|-----------|---------|-----------|
| Latitude | -90 to 90 | Posisi Utara/Selatan |
| Longitude | -180 to 180 | Posisi Timur/Barat |

### Contoh Koordinat

| Lokasi | Lintang | Bujur |
|--------|---------|-------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## Membuat Catatan dengan Nilai Lokasi

Saat membuat catatan baru dengan data lokasi:

```graphql
mutation CreateRecordWithLocation {
  createTodo(input: {
    title: "Site Visit"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "location_field_id"
      value: "40.7128,-74.0060"  # Format: "latitude,longitude"
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
      latitude
      longitude
    }
  }
}
```

### Format Input untuk Pembuatan

Saat membuat catatan, nilai lokasi menggunakan format yang dipisahkan dengan koma:

| Format | Contoh | Deskripsi |
|--------|--------|-----------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | Format koordinat standar |
| `"51.5074,-0.1278"` | London coordinates | Tidak ada spasi di sekitar koma |
| `"-33.8688,151.2093"` | Sydney coordinates | Nilai negatif diperbolehkan |

## Bidang Respons

### TodoCustomField Response

| Bidang | Tipe | Deskripsi |
|--------|------|-----------|
| `id` | String! | Pengidentifikasi unik untuk nilai bidang |
| `customField` | CustomField! | Definisi bidang kustom |
| `latitude` | Float | Koordinat lintang |
| `longitude` | Float | Koordinat bujur |
| `todo` | Todo! | Catatan yang nilai ini miliki |
| `createdAt` | DateTime! | Saat nilai dibuat |
| `updatedAt` | DateTime! | Saat nilai terakhir dimodifikasi |

## Pembatasan Penting

### Tidak Ada Geocoding Bawaan

Bidang lokasi hanya menyimpan koordinat - mereka **tidak** mencakup:
- Konversi alamat ke koordinat
- Geocoding terbalik (koordinat ke alamat)
- Validasi atau pencarian alamat
- Integrasi dengan layanan pemetaan
- Pencarian nama tempat

### Layanan Eksternal Diperlukan

Untuk fungsionalitas alamat, Anda perlu mengintegrasikan layanan eksternal:
- **Google Maps API** untuk geocoding
- **OpenStreetMap Nominatim** untuk geocoding gratis
- **MapBox** untuk pemetaan dan geocoding
- **Here API** untuk layanan lokasi

### Contoh Integrasi

```javascript
// Client-side geocoding example (not part of Blue API)
async function geocodeAddress(address) {
  const response = await fetch(
    `https://maps.googleapis.com/maps/api/geocode/json?address=${encodeURIComponent(address)}&key=${API_KEY}`
  );
  const data = await response.json();
  
  if (data.results.length > 0) {
    const { lat, lng } = data.results[0].geometry.location;
    
    // Now set the location field in Blue
    await setTodoCustomField({
      todoId: "todo_123",
      customFieldId: "location_field_456",
      latitude: lat,
      longitude: lng
    });
  }
}
```

## Izin yang Diperlukan

| Tindakan | Peran yang Diperlukan |
|----------|-----------------------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## Respons Kesalahan

### Koordinat Tidak Valid
```json
{
  "errors": [{
    "message": "Invalid coordinates: latitude must be between -90 and 90",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Bujur Tidak Valid
```json
{
  "errors": [{
    "message": "Invalid coordinates: longitude must be between -180 and 180",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Praktik Terbaik

### Pengumpulan Data
- Gunakan koordinat GPS untuk lokasi yang tepat
- Validasi koordinat sebelum menyimpan
- Pertimbangkan kebutuhan presisi koordinat (6 tempat desimal ≈ 10cm akurasi)
- Simpan koordinat dalam derajat desimal (bukan derajat/menit/detik)

### Pengalaman Pengguna
- Sediakan antarmuka peta untuk pemilihan koordinat
- Tampilkan pratinjau lokasi saat menampilkan koordinat
- Validasi koordinat di sisi klien sebelum panggilan API
- Pertimbangkan implikasi zona waktu untuk data lokasi

### Kinerja
- Gunakan indeks spasial untuk kueri yang efisien
- Batasi presisi koordinat pada akurasi yang dibutuhkan
- Pertimbangkan caching untuk lokasi yang sering diakses
- Batch pembaruan lokasi jika memungkinkan

## Kasus Penggunaan Umum

1. **Operasi Lapangan**
   - Lokasi peralatan
   - Alamat panggilan layanan
   - Lokasi inspeksi
   - Lokasi pengiriman

2. **Manajemen Acara**
   - Tempat acara
   - Lokasi pertemuan
   - Lokasi konferensi
   - Lokasi lokakarya

3. **Pelacakan Aset**
   - Posisi peralatan
   - Lokasi fasilitas
   - Pelacakan kendaraan
   - Lokasi inventaris

4. **Analisis Geografis**
   - Area cakupan layanan
   - Distribusi pelanggan
   - Analisis pasar
   - Manajemen wilayah

## Fitur Integrasi

### Dengan Pencarian
- Referensi data lokasi dari catatan lain
- Temukan catatan berdasarkan kedekatan geografis
- Agregasi data berbasis lokasi
- Referensi silang koordinat

### Dengan Automasi
- Memicu tindakan berdasarkan perubahan lokasi
- Buat notifikasi geofenced
- Perbarui catatan terkait saat lokasi berubah
- Hasilkan laporan berbasis lokasi

### Dengan Rumus
- Hitung jarak antara lokasi
- Tentukan pusat geografis
- Analisis pola lokasi
- Buat metrik berbasis lokasi

## Pembatasan

- Tidak ada geocoding bawaan atau konversi alamat
- Tidak ada antarmuka pemetaan yang disediakan
- Memerlukan layanan eksternal untuk fungsionalitas alamat
- Terbatas pada penyimpanan koordinat saja
- Tidak ada validasi lokasi otomatis di luar pemeriksaan rentang

## Sumber Daya Terkait

- [Ikhtisar Bidang Kustom](/api/custom-fields/list-custom-fields) - Konsep umum
- [Google Maps API](https://developers.google.com/maps) - Layanan geocoding
- [OpenStreetMap Nominatim](https://nominatim.org/) - Geocoding gratis
- [MapBox API](https://docs.mapbox.com/) - Layanan pemetaan dan geocoding
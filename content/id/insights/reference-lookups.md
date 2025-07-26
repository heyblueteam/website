---
title: Referensi & pencarian bidang kustom
description: Buat proyek yang saling terhubung dengan mudah di Blue, mengubahnya menjadi sumber kebenaran tunggal untuk bisnis Anda dengan Bidang Referensi dan Pencarian yang baru.
category: "Product Updates"
date: 2023-11-01
---


Proyek di Blue sudah menjadi cara yang kuat untuk mengelola data bisnis Anda dan memajukan pekerjaan.

Hari ini, kami mengambil langkah logis berikutnya dan memungkinkan Anda untuk menghubungkan data Anda *antara* proyek untuk fleksibilitas dan kekuatan yang maksimal.

Menghubungkan proyek dalam Blue mengubahnya menjadi sumber kebenaran tunggal untuk bisnis Anda. Kemampuan ini memungkinkan pembuatan dataset yang komprehensif dan saling terhubung, memungkinkan aliran data yang mulus dan meningkatkan visibilitas di seluruh proyek. Dengan menghubungkan proyek, tim dapat mencapai pandangan yang terpadu tentang operasi, meningkatkan pengambilan keputusan dan efisiensi operasional.

## Contoh

Pertimbangkan Perusahaan ACME, yang menggunakan bidang kustom Referensi dan Pencarian Blue untuk menciptakan ekosistem data yang saling terhubung di seluruh proyek Pelanggan, Penjualan, dan Inventaris. Rekaman pelanggan di proyek Pelanggan dihubungkan melalui bidang Referensi ke transaksi penjualan di proyek Penjualan. Penghubungan ini memungkinkan bidang Pencarian untuk menarik detail pelanggan terkait, seperti nomor telepon dan status akun, langsung ke setiap catatan penjualan. Selain itu, item inventaris yang terjual ditampilkan dalam catatan penjualan melalui bidang Pencarian yang merujuk pada data Jumlah Terjual dari proyek Inventaris. Akhirnya, penarikan inventaris terhubung ke penjualan yang relevan melalui bidang Referensi di Inventaris, yang mengarah kembali ke catatan Penjualan. Pengaturan ini memberikan visibilitas penuh tentang penjualan mana yang memicu penghapusan inventaris, menciptakan pandangan terintegrasi 360 derajat di seluruh proyek.

## Cara Kerja Bidang Referensi

Bidang kustom Referensi memungkinkan Anda untuk membuat hubungan antara catatan di berbagai proyek di Blue. Saat membuat bidang Referensi, Administrator Proyek memilih proyek tertentu yang akan menyediakan daftar catatan referensi. Opsi konfigurasi termasuk:

* **Pilih Tunggal**: Memungkinkan memilih satu catatan referensi.
* **Pilih Beberapa**: Memungkinkan memilih beberapa catatan referensi.
* **Penyaringan**: Atur filter untuk memungkinkan pengguna memilih hanya catatan yang sesuai dengan kriteria filter.

Setelah diatur, pengguna dapat memilih catatan tertentu dari menu dropdown dalam bidang Referensi, membangun tautan antara proyek.

## Memperluas bidang referensi menggunakan pencarian

Bidang kustom Pencarian digunakan untuk mengimpor data dari catatan di proyek lain, menciptakan visibilitas satu arah. Mereka selalu hanya baca dan terhubung ke bidang kustom Referensi tertentu. Ketika seorang pengguna memilih satu atau lebih catatan menggunakan bidang kustom Referensi, bidang kustom Pencarian akan menampilkan data dari catatan tersebut. Pencarian dapat menampilkan data seperti:

* Dibuat pada
* Diperbarui pada
* Tanggal Jatuh Tempo
* Deskripsi
* Daftar
* Tag
* Penugasan
* Bidang kustom yang didukung dari catatan yang dirujuk — termasuk bidang pencarian lainnya!

Sebagai contoh, bayangkan skenario di mana Anda memiliki tiga proyek: **Proyek A** adalah proyek penjualan, **Proyek B** adalah proyek manajemen inventaris, dan **Proyek C** adalah proyek hubungan pelanggan. Di Proyek A, Anda memiliki bidang kustom Referensi yang menghubungkan catatan penjualan ke catatan pelanggan yang sesuai di Proyek C. Di Proyek B, Anda memiliki bidang kustom Pencarian yang mengimpor informasi dari Proyek A, seperti jumlah yang terjual. Dengan cara ini, ketika catatan penjualan dibuat di Proyek A, informasi pelanggan yang terkait dengan penjualan tersebut secara otomatis ditarik dari Proyek C, dan jumlah yang terjual secara otomatis ditarik dari Proyek B. Ini memungkinkan Anda untuk menyimpan semua informasi relevan di satu tempat dan melihatnya tanpa harus membuat data duplikat atau memperbarui catatan secara manual di seluruh proyek.

Contoh nyata dari ini adalah perusahaan e-commerce yang menggunakan Blue untuk mengelola penjualan, inventaris, dan hubungan pelanggan mereka. Di proyek **Penjualan** mereka, mereka memiliki bidang kustom Referensi yang menghubungkan setiap catatan penjualan ke catatan **Pelanggan** yang sesuai di proyek **Pelanggan** mereka. Di proyek **Inventaris** mereka, mereka memiliki bidang kustom Pencarian yang mengimpor informasi dari proyek Penjualan, seperti jumlah yang terjual, dan menampilkannya dalam catatan item inventaris. Ini memungkinkan mereka untuk dengan mudah melihat penjualan mana yang mendorong penghapusan inventaris dan menjaga tingkat inventaris mereka tetap terkini tanpa harus memperbarui catatan secara manual di seluruh proyek.

## Kesimpulan

Bayangkan dunia di mana data proyek Anda tidak terisolasi tetapi mengalir bebas antara proyek, memberikan wawasan komprehensif dan mendorong efisiensi. Itulah kekuatan bidang Referensi dan Pencarian Blue. Dengan memungkinkan koneksi data yang mulus dan memberikan visibilitas waktu nyata di seluruh proyek, fitur-fitur ini mengubah cara tim berkolaborasi dan membuat keputusan. Apakah Anda mengelola hubungan pelanggan, melacak penjualan, atau mengawasi inventaris, bidang Referensi dan Pencarian di Blue memberdayakan tim Anda untuk bekerja lebih cerdas, lebih cepat, dan lebih efektif. Masuki dunia yang saling terhubung di Blue dan saksikan produktivitas Anda meningkat.

[Periksa dokumentasi](https://documentation.blue.cc/custom-fields/reference) atau [daftar dan coba sendiri.](https://app.blue.cc)
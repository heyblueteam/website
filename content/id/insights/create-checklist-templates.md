---
title: Membuat checklist yang dapat digunakan kembali menggunakan automasi
description: Pelajari cara membuat automasi manajemen proyek untuk checklist yang dapat digunakan kembali.
category: "Best Practices"
date: 2024-07-08
---


Dalam banyak proyek dan proses, Anda mungkin perlu menggunakan checklist yang sama di berbagai catatan atau tugas.

Namun, tidak efisien untuk mengetik ulang checklist secara manual setiap kali Anda ingin menambahkannya ke sebuah catatan. Di sinilah Anda dapat memanfaatkan [automasi manajemen proyek yang kuat](/platform/features/automations) untuk melakukan ini secara otomatis untuk Anda!

Sebagai pengingat, automasi di Blue memerlukan dua hal kunci:

1. Pemicu — Apa yang harus terjadi untuk memulai automasi. Ini bisa terjadi ketika sebuah catatan diberikan tag tertentu, ketika berpindah ke sebuah 
2. Satu atau lebih Tindakan — Dalam hal ini, itu akan menjadi pembuatan otomatis satu atau lebih checklist.

Mari kita mulai dengan tindakan terlebih dahulu, kemudian diskusikan pemicu yang mungkin Anda gunakan.

## Tindakan Automasi Checklist

Anda dapat membuat automasi baru, dan Anda dapat mengatur satu atau lebih checklist untuk dibuat, seperti contoh di bawah ini:

![](/insights/checklist-automation.png)

Ini akan menjadi checklist yang ingin Anda buat setiap kali Anda mengambil tindakan.

## Pemicu Automasi Checklist

Ada beberapa cara Anda dapat memicu pembuatan checklist yang dapat digunakan kembali. Berikut adalah beberapa opsi populer:

- **Menambahkan Tag Tertentu:** Anda dapat mengatur automasi untuk dipicu ketika tag tertentu ditambahkan ke sebuah catatan. Misalnya, ketika tag "Proyek Baru" ditambahkan, itu bisa secara otomatis membuat checklist inisiasi proyek Anda.
- **Penugasan Catatan:** Pembuatan checklist dapat dipicu ketika sebuah catatan ditugaskan kepada individu tertentu atau kepada siapa saja. Ini berguna untuk checklist onboarding atau prosedur spesifik tugas.
- **Pindah ke Daftar Tertentu:** Ketika sebuah catatan dipindahkan ke daftar tertentu di papan proyek Anda, itu dapat memicu pembuatan checklist yang relevan. Misalnya, memindahkan item ke daftar "Jaminan Kualitas" dapat memicu checklist QA.
- **Bidang Checkbox Kustom:** Buat bidang checkbox kustom dan atur automasi untuk dipicu ketika kotak ini dicentang. Ini memberi Anda kontrol manual atas kapan menambahkan checklist.
- **Bidang Kustom Pilihan Tunggal atau Multi-Pilih:** Anda dapat membuat bidang kustom pilihan tunggal atau multi-pilih dengan berbagai opsi. Setiap opsi dapat dihubungkan ke template checklist tertentu melalui automasi terpisah. Ini memungkinkan kontrol yang lebih rinci dan kemampuan untuk memiliki beberapa template checklist siap untuk berbagai skenario.

Untuk meningkatkan kontrol atas siapa yang dapat memicu automasi ini, Anda dapat menyembunyikan bidang kustom ini dari pengguna tertentu menggunakan peran pengguna kustom. Ini memastikan bahwa hanya admin proyek atau personel berwenang lainnya yang dapat memicu opsi ini.

Ingat, kunci untuk penggunaan efektif checklist yang dapat digunakan kembali dengan automasi adalah merancang pemicu Anda dengan bijaksana. Pertimbangkan alur kerja tim Anda, jenis proyek yang Anda tangani, dan siapa yang seharusnya memiliki kemampuan untuk memulai proses yang berbeda. Dengan automasi yang direncanakan dengan baik, Anda dapat secara signifikan memperlancar manajemen proyek Anda dan memastikan konsistensi di seluruh operasi Anda.

## Sumber Daya Berguna

- [Dokumentasi Automasi Manajemen Proyek](https://documentation.blue.cc/automations)
- [Dokumentasi Peran Pengguna Kustom](https://documentation.blue.cc/user-management/roles/custom-user-roles)
- [Dokumentasi Bidang Kustom](https://documentation.blue.cc/custom-fields)
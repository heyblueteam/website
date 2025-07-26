---
title: Automatisasi manajemen proyek — email kepada pemangku kepentingan.
description: Seringkali, Anda ingin mengendalikan otomatisasi manajemen proyek Anda.
category: "Product Updates"
date: 2024-07-08
---


Kami telah membahas cara [membuat otomatisasi email sebelumnya.](/insights/email-automations)

Namun, sering kali ada pemangku kepentingan dalam proyek yang hanya perlu diberi tahu ketika ada sesuatu yang *sangat* penting.

Tidakkah akan menyenangkan jika ada otomatisasi manajemen proyek di mana Anda, sebagai manajer proyek, dapat mengendalikan *persis* kapan memberi tahu pemangku kepentingan kunci dengan menekan tombol?

Nah, ternyata dengan Blue, Anda dapat melakukan hal ini dengan tepat!

Hari ini kita akan belajar cara membuat otomatisasi manajemen proyek yang sangat berguna:

Sebuah kotak centang yang secara otomatis memberi tahu satu atau lebih pemangku kepentingan kunci, memberikan mereka semua konteks penting tentang apa yang Anda beri tahu. Sebagai poin bonus, kita juga akan belajar cara membatasi kemampuan ini sehingga hanya anggota tertentu dari proyek Anda yang dapat memicu pemberitahuan email ini.

Ini akan terlihat seperti ini setelah Anda selesai:

![](/insights/checkbox-email-automation.png)

Dan hanya dengan mencentang kotak ini, Anda akan dapat memicu otomatisasi manajemen proyek untuk mengirim email pemberitahuan kustom kepada pemangku kepentingan.

Mari kita lakukan langkah demi langkah.

## 1. Buat bidang kustom kotak centang Anda

Ini sangat mudah, Anda dapat melihat [dokumentasi rinci kami](https://documentation.blue.cc/custom-fields/introduction#creating-custom-fields) tentang cara membuat bidang kustom.

Pastikan Anda memberi nama bidang ini sesuatu yang jelas yang akan Anda ingat seperti “beri tahu manajemen” atau “beri tahu pemangku kepentingan”.

## 2. Buat pemicu otomatisasi manajemen proyek Anda.

Di tampilan catatan dalam proyek Anda, klik pada robot kecil di kanan atas untuk membuka pengaturan otomatisasi:

<video autoplay loop muted playsinline>
  <source src="/videos/notify-stakeholders-automation-setup.mp4" type="video/mp4">
</video>

## 3. Buat tindakan otomatisasi manajemen proyek Anda.

Dalam hal ini, tindakan kita adalah mengirim pemberitahuan email kustom ke satu atau lebih alamat email. Penting untuk dicatat di sini bahwa orang-orang ini **tidak** harus berada di Blue untuk menerima email ini, Anda dapat mengirim email ke *alamat email mana pun*.

Anda dapat mempelajari lebih lanjut di [panduan dokumentasi rinci kami tentang cara mengatur otomatisasi email](https://documentation.blue.cc/automations/actions/email-automations)

Hasil akhir Anda harus terlihat seperti ini:

![](/insights/email-automation-example.png)

## 4. Bonus: Batasi akses ke kotak centang.

Anda dapat menggunakan [peran pengguna kustom di Blue](/platform/features/user-permissions) untuk membatasi akses ke bidang kustom kotak centang, memastikan bahwa hanya anggota tim yang berwenang yang dapat memicu pemberitahuan email.

Blue memungkinkan Administrator Proyek untuk mendefinisikan peran dan menetapkan izin kepada kelompok pengguna. Sistem ini sangat penting untuk mempertahankan kontrol atas siapa yang dapat berinteraksi dengan elemen tertentu dari proyek Anda, termasuk bidang kustom seperti kotak centang pemberitahuan.

1. Arahkan ke bagian Manajemen Pengguna di Blue dan pilih "Peran Pengguna Kustom."
2. Buat peran baru dengan memberikan nama deskriptif dan deskripsi opsional.
3. Dalam izin peran, temukan bagian untuk Akses Bidang Kustom.
4. Tentukan apakah peran dapat melihat atau mengedit bidang kustom kotak centang. Misalnya, batasi akses pengeditan untuk peran seperti "Administrator Proyek" sambil memungkinkan peran kustom yang baru dibuat untuk mengelola bidang ini.
5. Tetapkan peran yang baru dibuat kepada pengguna atau kelompok pengguna yang sesuai. Ini memastikan bahwa hanya individu yang ditunjuk yang memiliki kemampuan untuk berinteraksi dengan kotak centang pemberitahuan.

[Baca lebih lanjut di situs dokumentasi resmi kami.](https://documentation.blue.cc/user-management/roles/custom-user-roles)

Dengan menerapkan peran kustom ini, Anda meningkatkan keamanan dan integritas proses manajemen proyek Anda. Hanya anggota tim yang berwenang yang dapat memicu pemberitahuan email penting, memastikan bahwa pemangku kepentingan menerima pembaruan penting tanpa pemberitahuan yang tidak perlu.

## Kesimpulan

Dengan menerapkan otomatisasi manajemen proyek yang dijelaskan di atas, Anda mendapatkan kontrol yang tepat atas kapan dan bagaimana memberi tahu pemangku kepentingan kunci. Pendekatan ini memastikan bahwa pembaruan penting dikomunikasikan secara efektif, tanpa membanjiri pemangku kepentingan Anda dengan informasi yang tidak perlu. Dengan memanfaatkan fitur bidang kustom dan otomatisasi Blue, Anda dapat menyederhanakan proses manajemen proyek Anda, meningkatkan komunikasi, dan mempertahankan tingkat efisiensi yang tinggi.

Hanya dengan kotak centang sederhana, Anda dapat memicu pemberitahuan email kustom yang disesuaikan dengan kebutuhan proyek Anda, memastikan bahwa orang yang tepat diinformasikan pada waktu yang tepat. Selain itu, kemampuan untuk membatasi fungsionalitas ini kepada anggota tim tertentu menambah lapisan kontrol dan keamanan tambahan.

Mulailah memanfaatkan fitur kuat ini di Blue hari ini untuk menjaga pemangku kepentingan Anda tetap terinformasi dan proyek Anda berjalan lancar. Untuk langkah-langkah lebih rinci dan opsi kustomisasi tambahan, lihat tautan dokumentasi yang disediakan. Selamat mengotomatisasi!
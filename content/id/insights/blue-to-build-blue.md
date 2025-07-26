---
title:  Bagaimana kami menggunakan Blue untuk membangun Blue. 
description: Pelajari bagaimana kami menggunakan platform manajemen proyek kami sendiri untuk membangun platform manajemen proyek kami!
category: "CEO Blog"
date: 2024-08-07
---


Anda akan mendapatkan tur dari dalam tentang bagaimana Blue membangun Blue.

Di Blue, kami menggunakan produk kami sendiri.

Ini berarti bahwa kami menggunakan Blue untuk *membangun* Blue.

Istilah yang terdengar aneh ini, yang sering disebut sebagai "dogfooding", sering dikaitkan dengan Paul Maritz, seorang manajer di Microsoft pada tahun 1980-an. Dia dilaporkan mengirim email dengan subjek *"Eating our own dog food"* untuk mendorong karyawan Microsoft menggunakan produk perusahaan.

Ide menggunakan alat Anda sendiri untuk membangun alat Anda adalah bahwa itu mengarah pada siklus umpan balik yang positif.

Ide menggunakan alat Anda sendiri untuk membangun alat Anda mengarah pada siklus umpan balik yang positif, menciptakan berbagai manfaat:

- **Ini membantu kami mengidentifikasi masalah kegunaan di dunia nyata dengan cepat.** Saat kami menggunakan Blue setiap hari, kami menghadapi tantangan yang sama yang mungkin dihadapi pengguna kami, memungkinkan kami untuk mengatasinya secara proaktif.
- **Ini mempercepat penemuan bug.** Penggunaan internal sering mengungkap bug sebelum mencapai pelanggan kami, meningkatkan kualitas produk secara keseluruhan.
- **Ini meningkatkan empati kami terhadap pengguna akhir.** Tim kami mendapatkan pengalaman langsung tentang kekuatan dan kelemahan Blue, membantu kami membuat keputusan yang lebih berfokus pada pengguna.
- **Ini mendorong budaya kualitas di dalam organisasi kami.** Ketika semua orang menggunakan produk, ada kepentingan bersama dalam keunggulannya.
- **Ini mendorong inovasi.** Penggunaan reguler sering memicu ide untuk fitur baru atau perbaikan, menjaga Blue di garis depan.

[Kami telah berbicara sebelumnya tentang mengapa kami tidak memiliki tim pengujian yang khusus](/insights/open-beta) dan ini adalah alasan lain. 

Jika ada bug dalam sistem kami, kami hampir selalu menemukannya dalam penggunaan platform kami yang konstan setiap hari. Dan ini juga menciptakan fungsi paksaan untuk memperbaikinya, karena kami jelas akan merasa sangat terganggu saat menemukannya, karena kami mungkin salah satu pengguna utama Blue!

Pendekatan ini menunjukkan komitmen kami terhadap produk. Dengan mengandalkan Blue sendiri, kami menunjukkan kepada pelanggan kami bahwa kami benar-benar percaya pada apa yang kami bangun. Ini bukan hanya produk yang kami jual – ini adalah alat yang kami andalkan setiap hari.

## Proses Utama

Kami memiliki satu proyek di Blue, yang dinamakan "Produk".

**Segala sesuatu** yang terkait dengan pengembangan produk kami dilacak di sini. Umpan balik pelanggan, bug, ide fitur, pekerjaan yang sedang berlangsung, dan sebagainya. Ide memiliki satu proyek di mana kami melacak semuanya adalah bahwa itu [mendorong kerja tim yang lebih baik.](/insights/great-teamwork)

Setiap catatan adalah fitur atau bagian dari fitur. Inilah cara kami bergerak dari "bukankah akan keren jika..." ke "lihat fitur baru yang luar biasa ini!"

Proyek ini memiliki daftar berikut:

- **Ide/Umpan Balik**: Ini adalah daftar ide tim atau umpan balik pelanggan berdasarkan panggilan atau pertukaran email. Silakan tambahkan ide apa pun di sini! Dalam daftar ini, kami belum memutuskan bahwa kami akan membangun salah satu fitur ini, tetapi kami secara teratur meninjaunya untuk ide yang ingin kami eksplorasi lebih lanjut.
- **Backlog (Jangka Panjang)**: Ini adalah tempat fitur dari daftar Ide/Umpan Balik pergi jika kami memutuskan bahwa itu akan menjadi tambahan yang baik untuk Blue.
- **{Kuartal Saat Ini}**: Ini biasanya disusun sebagai "Qx YYYY" dan menunjukkan prioritas kuartal kami.
- **Bug**: Ini adalah daftar bug yang diketahui yang dilaporkan oleh tim atau pelanggan. Bug yang ditambahkan di sini akan secara otomatis memiliki tag "Bug" ditambahkan.
- **Spesifikasi**: Fitur-fitur ini saat ini sedang dispesifikasikan. Tidak setiap fitur memerlukan spesifikasi atau desain; itu tergantung pada ukuran yang diharapkan dari fitur dan tingkat kepercayaan yang kami miliki terkait kasus tepi dan kompleksitas.
- **Backlog Desain**: Ini adalah backlog untuk desainer, setiap kali mereka menyelesaikan sesuatu yang sedang berlangsung, mereka dapat memilih item dari daftar ini.
- **Desain Sedang Berlangsung**: Ini adalah fitur-fitur yang saat ini sedang dirancang oleh desainer.
- **Tinjauan Desain**: Ini adalah tempat fitur-fitur yang desainnya sedang ditinjau.
- **Backlog (Jangka Pendek)**: Ini adalah daftar fitur yang kemungkinan besar akan kami mulai kerjakan dalam beberapa minggu ke depan. Di sinilah penugasan dilakukan. CEO dan Kepala Teknik memutuskan fitur mana yang ditugaskan kepada insinyur mana berdasarkan pengalaman sebelumnya dan beban kerja. [Anggota tim kemudian dapat menarik ini ke dalam In Progress](/insights/push-vs-pull-kanban) setelah mereka menyelesaikan pekerjaan mereka saat ini.
- **Sedang Berlangsung**: Ini adalah fitur-fitur yang saat ini sedang dikembangkan.
- **Tinjauan Kode**: Setelah fitur selesai dikembangkan, ia menjalani tinjauan kode. Kemudian akan dipindahkan kembali ke "Sedang Berlangsung" jika penyesuaian diperlukan atau dikerahkan ke lingkungan Pengembangan.
- **Dev**: Ini adalah semua fitur yang saat ini ada di lingkungan Pengembangan. Anggota tim lain dan pelanggan tertentu dapat meninjau ini.
- **Beta**: Ini adalah semua fitur yang saat ini ada di [lingkungan Beta](https://beta.app.blue.cc). Banyak pelanggan menggunakan ini sebagai platform Blue harian mereka dan juga akan memberikan umpan balik.
- **Produksi**: Ketika fitur mencapai produksi, maka dianggap selesai.

Kadang-kadang, saat kami mengembangkan fitur, kami menyadari bahwa sub-fitur tertentu lebih sulit untuk diimplementasikan daripada yang diharapkan, dan kami mungkin memilih untuk tidak melakukannya dalam versi awal yang kami terapkan kepada pelanggan. Dalam hal ini, kami dapat membuat catatan baru dengan nama mengikuti format "{FeatureName} V2" dan menyertakan semua sub-fitur sebagai item daftar periksa.

## Tag

- **Mobile**: Ini berarti bahwa fitur tersebut spesifik untuk aplikasi iOS, Android, atau iPad kami.
- **{NamaPelangganPerusahaan}**: Fitur sedang dibangun khusus untuk pelanggan perusahaan. Pelacakan penting karena biasanya ada perjanjian komersial tambahan untuk setiap fitur.
- **Bug**: Ini berarti bahwa ini adalah bug yang memerlukan perbaikan.
- **Fast-Track**: Ini berarti bahwa ini adalah Perubahan Fast-Track yang tidak perlu melalui siklus rilis penuh seperti yang dijelaskan di atas.
- **Main**: Ini adalah pengembangan fitur utama. Ini biasanya diperuntukkan bagi pekerjaan infrastruktur besar, peningkatan ketergantungan besar, dan modul baru yang signifikan dalam Blue.
- **AI**: Fitur ini mengandung komponen kecerdasan buatan.
- **Keamanan**: Ini berarti implikasi keamanan harus ditinjau atau patch diperlukan.

Tag fast-track sangat menarik. Ini diperuntukkan bagi pembaruan yang lebih kecil dan kurang kompleks yang tidak memerlukan siklus rilis penuh kami, dan yang ingin kami kirimkan kepada pelanggan dalam waktu 24-48 jam.

Perubahan fast-track biasanya adalah penyesuaian kecil yang dapat secara signifikan meningkatkan pengalaman pengguna tanpa mengubah fungsionalitas inti. Pikirkan tentang memperbaiki kesalahan ketik di UI, menyesuaikan padding tombol, atau menambahkan ikon baru untuk panduan visual yang lebih baik. Ini adalah jenis perubahan yang, meskipun kecil, dapat membuat perbedaan besar dalam bagaimana pengguna memandang dan berinteraksi dengan produk kami. Mereka juga mengganggu jika memakan waktu lama untuk dikirim!

Proses fast-track kami sederhana.

Ini dimulai dengan membuat cabang baru dari cabang utama, menerapkan perubahan, dan kemudian membuat permintaan penggabungan untuk setiap cabang target - Dev, Beta, dan Produksi. Kami menghasilkan tautan pratinjau untuk ditinjau, memastikan bahwa bahkan perubahan kecil ini memenuhi standar kualitas kami. Setelah disetujui, perubahan digabungkan secara bersamaan ke dalam semua cabang, menjaga lingkungan kami tetap sinkron.

## Bidang Kustom

Kami tidak memiliki banyak bidang kustom dalam proyek Produk kami.

- **Spesifikasi**: Ini menghubungkan ke dokumen Blue yang memiliki spesifikasi untuk fitur tertentu tersebut. Ini tidak selalu dilakukan, karena tergantung pada kompleksitas fitur.
- **MR**: Ini adalah tautan ke Permintaan Penggabungan di [Gitlab](https://gitlab.com) tempat kami menyimpan kode kami.
- **Tautan Pratinjau**: Untuk fitur yang terutama mengubah antarmuka depan, kami dapat membuat URL unik yang memiliki perubahan tersebut untuk setiap komit, sehingga kami dapat dengan mudah meninjau perubahan.
- **Pemimpin**: Bidang ini memberi tahu kami insinyur senior mana yang memimpin tinjauan kode. Ini memastikan bahwa setiap fitur mendapatkan perhatian ahli yang layak, dan selalu ada orang yang jelas untuk pertanyaan atau kekhawatiran.

## Daftar Periksa

Selama demo mingguan kami, kami akan mencantumkan umpan balik yang dibahas dalam daftar periksa yang disebut "umpan balik" dan akan ada juga daftar periksa lain yang berisi [WBS (Work Breakdown Scope)](/insights/simple-work-breakdown-structure) utama dari fitur tersebut, sehingga kami dapat dengan mudah mengetahui apa yang sudah selesai dan apa yang belum dilakukan.

## Kesimpulan

Dan itu saja!

Kami pikir kadang-kadang orang terkejut dengan betapa sederhana proses kami, tetapi kami percaya bahwa proses yang sederhana sering kali jauh lebih unggul daripada proses yang terlalu kompleks yang tidak dapat dengan mudah dipahami.

Kesederhanaan ini disengaja. Ini memungkinkan kami untuk tetap gesit, merespons dengan cepat terhadap kebutuhan pelanggan, dan menjaga seluruh tim kami selaras.

Dengan menggunakan Blue untuk membangun Blue, kami tidak hanya mengembangkan produk – kami menghidupkannya.

Jadi, lain kali Anda menggunakan Blue, ingatlah: Anda tidak hanya menggunakan produk yang telah kami bangun. Anda menggunakan produk yang kami andalkan setiap hari.

Dan itu membuat semua perbedaan.
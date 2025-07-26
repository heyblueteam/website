---
title:  Mengapa Blue Memiliki Beta Terbuka
description: Pelajari mengapa sistem manajemen proyek kami memiliki beta terbuka yang sedang berlangsung.
category: "Engineering"
date: 2024-08-03
---


Banyak startup B2B SaaS yang diluncurkan dalam Beta, dan dengan alasan yang baik. Ini adalah bagian dari moto tradisional Silicon Valley *“bergerak cepat dan merusak hal-hal”*.

Menempelkan stiker “beta” pada produk mengurangi ekspektasi.

Ada yang rusak? Oh ya, ini hanya beta.

Sistemnya lambat? Oh ya, ini hanya beta.

[Dokumentasi](https://blue.cc/docs) tidak ada? Oh ya… Anda mengerti maksudnya.

Dan ini *sebenarnya* adalah hal yang baik. Reid Hoffman, pendiri LinkedIn, terkenal mengatakan:

> Jika Anda tidak merasa malu dengan versi pertama produk Anda, Anda telah meluncurkan terlalu lambat.

Dan stiker beta juga baik untuk pelanggan. Ini membantu mereka untuk memilih sendiri.

Pelanggan yang mencoba produk beta adalah mereka yang berada di tahap awal dari Siklus Adopsi Teknologi, yang juga dikenal sebagai Kurva Adopsi Produk.

Siklus Adopsi Teknologi biasanya dibagi menjadi lima segmen utama:

1. Inovator
2. Pengadopsi Awal
3. Mayoritas Awal
4. Mayoritas Akhir
5. Pengikut

![](/insights/technology-adoption-lifecycle-graph.png)

Namun, pada akhirnya produk harus matang, dan pelanggan mengharapkan produk yang stabil dan berfungsi. Mereka tidak ingin mengakses lingkungan “beta” di mana segala sesuatunya bisa rusak.

Atau apakah mereka?

*Inilah* pertanyaan yang kami ajukan pada diri kami sendiri.

Kami percaya kami mengajukan pertanyaan ini karena sifat bagaimana Blue awalnya dibangun. [Blue dimulai sebagai cabang dari agensi desain yang sibuk](/insights/agency-success-playbook), dan kami bekerja *di dalam* kantor bisnis yang secara aktif menggunakan Blue untuk menjalankan semua proyek mereka.

Ini berarti bahwa selama bertahun-tahun, kami dapat mengamati bagaimana *manusia nyata* — yang duduk tepat di sebelah kami! — menggunakan Blue dalam kehidupan sehari-hari mereka.

Dan karena mereka menggunakan Blue sejak hari-hari awal, tim ini selalu menggunakan Blue Beta!

Oleh karena itu, adalah hal yang wajar bagi kami untuk membiarkan semua pelanggan kami yang lain menggunakannya juga.

**Dan inilah mengapa kami tidak memiliki tim pengujian yang khusus.**

Itu benar.

Tidak ada orang di Blue yang memiliki tanggung jawab *sendiri* untuk memastikan bahwa platform kami berjalan dengan baik dan stabil.

Ini karena beberapa alasan.

Alasan pertama adalah basis biaya yang lebih rendah.

Tidak memiliki tim pengujian penuh waktu secara signifikan mengurangi biaya kami, dan kami dapat meneruskan penghematan ini kepada pelanggan kami dengan harga terendah di industri.

Untuk memberikan perspektif, kami menawarkan set fitur tingkat perusahaan yang dikenakan biaya $30-$55/user/bulan oleh pesaing kami hanya dengan $7/bulan.

Ini tidak terjadi secara kebetulan, ini *dirancang*.

Namun, bukan strategi yang baik untuk menjual produk yang lebih murah jika tidak berfungsi.

Jadi *pertanyaan sebenarnya adalah*, bagaimana kami berhasil menciptakan platform yang stabil yang dapat digunakan oleh ribuan pelanggan tanpa tim pengujian yang khusus?

Tentu saja, pendekatan kami untuk memiliki Beta terbuka sangat penting untuk ini, tetapi sebelum kami membahasnya, kami ingin menyentuh tanggung jawab pengembang.

Kami membuat keputusan awal di Blue bahwa kami tidak akan pernah membagi tanggung jawab untuk teknologi front-end dan back-end. Kami hanya akan mempekerjakan atau melatih pengembang full stack.

Alasan kami membuat keputusan ini adalah untuk memastikan bahwa seorang pengembang akan sepenuhnya memiliki fitur yang mereka kerjakan. Jadi tidak akan ada mentalitas *“lempar masalah ke pagar kebun”* yang kadang-kadang muncul ketika ada tanggung jawab bersama untuk fitur.

Dan ini juga berlaku untuk pengujian fitur, memahami kasus penggunaan pelanggan dan permintaan, serta membaca dan mengomentari spesifikasi.

Dengan kata lain, setiap pengembang membangun pemahaman yang mendalam dan intuitif tentang fitur yang mereka bangun.

Baiklah, sekarang mari kita bicarakan tentang beta terbuka kami.

Ketika kami mengatakan itu “terbuka” — kami benar-benar maksudkan. Setiap pelanggan dapat mencobanya hanya dengan menambahkan “beta” di depan URL aplikasi web kami.

Jadi “app.blue.cc” menjadi “beta.app.blue.cc”

Ketika mereka melakukan ini, mereka dapat melihat data biasa mereka, karena baik lingkungan Beta maupun Produksi berbagi database yang sama, tetapi mereka juga akan dapat melihat fitur baru.

Pelanggan dapat dengan mudah bekerja bahkan jika mereka memiliki beberapa anggota tim di Produksi dan beberapa yang penasaran di Beta.

Kami biasanya memiliki beberapa ratus pelanggan yang menggunakan Beta pada satu waktu, dan kami memposting pratinjau fitur di forum komunitas kami sehingga mereka dapat melihat apa yang baru dan mencobanya.

Dan inilah intinya: kami memiliki *beberapa ratus* penguji!

Semua pelanggan ini akan mencoba fitur dalam alur kerja mereka, dan cukup vokal jika ada sesuatu yang tidak berjalan dengan baik, karena mereka *sudah* menerapkan fitur tersebut di dalam bisnis mereka!

Umpan balik yang paling umum adalah perubahan kecil tetapi sangat berguna yang menangani kasus pinggiran yang tidak kami pertimbangkan.

Kami meninggalkan fitur baru di Beta antara 2-4 minggu. Setiap kali kami merasa bahwa mereka stabil, maka kami merilisnya ke produksi.

Kami juga memiliki kemampuan untuk melewati Beta jika diperlukan, menggunakan bendera fast-track. Ini biasanya dilakukan untuk perbaikan bug yang tidak ingin kami tahan selama 2-4 minggu sebelum dikirim ke produksi.

Hasilnya?

Mendorong ke produksi terasa… yah membosankan! Seperti tidak ada — itu tidak menjadi masalah besar bagi kami.

Dan ini berarti bahwa ini memperlancar jadwal rilis kami, yang telah memungkinkan kami untuk [mengirim fitur setiap bulan seperti jam selama enam tahun terakhir.](/changelog).

Namun, seperti pilihan lainnya, ada beberapa trade-off.

Dukungan pelanggan sedikit lebih kompleks, karena kami harus mendukung pelanggan di dua versi platform kami. Terkadang ini dapat menyebabkan kebingungan bagi pelanggan yang memiliki anggota tim yang menggunakan dua versi yang berbeda.

Titik nyeri lainnya adalah bahwa pendekatan ini kadang-kadang dapat memperlambat jadwal rilis keseluruhan ke produksi. Ini terutama berlaku untuk fitur yang lebih besar yang dapat “terjebak” di Beta jika ada fitur terkait lain yang mengalami masalah dan memerlukan beberapa pekerjaan lebih lanjut.

Tetapi secara keseluruhan, kami berpikir bahwa trade-off ini sebanding dengan manfaat dari basis biaya yang lebih rendah dan keterlibatan pelanggan yang lebih besar.

Kami adalah salah satu dari sedikit perusahaan perangkat lunak yang mengadopsi pendekatan ini, tetapi sekarang ini adalah bagian fundamental dari pendekatan pengembangan produk kami.
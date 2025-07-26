---
title: Pencarian waktu nyata
description: Blue meluncurkan mesin pencari baru yang sangat cepat yang mengembalikan hasil di semua proyek Anda dalam milidetik, memberdayakan Anda untuk beralih konteks dalam sekejap mata.
category: "Product Updates"
date: 2024-03-01
---


Kami sangat senang mengumumkan peluncuran mesin pencari baru kami, yang dirancang untuk merevolusi cara Anda menemukan informasi di dalam Blue. Fungsionalitas pencarian yang efisien sangat penting untuk manajemen proyek yang lancar, dan pembaruan terbaru kami memastikan bahwa Anda dapat mengakses data Anda lebih cepat dari sebelumnya.

Mesin pencari baru kami memungkinkan Anda untuk mencari semua komentar, file, catatan, bidang kustom, deskripsi, dan daftar periksa. Apakah Anda perlu menemukan komentar tertentu yang dibuat pada suatu proyek, dengan cepat menemukan file, atau mencari catatan atau bidang tertentu, mesin pencari kami memberikan hasil yang sangat cepat.

Ketika alat mendekati responsivitas 50-100ms, mereka cenderung memudar dan menyatu dengan latar belakang, memberikan pengalaman pengguna yang mulus dan hampir tidak terlihat. Sebagai konteks, kedipan manusia memakan waktu sekitar 60-120ms, jadi 50ms sebenarnya lebih cepat daripada kedipan mata! Tingkat responsivitas ini memungkinkan Anda berinteraksi dengan Blue tanpa menyadari keberadaannya, membebaskan Anda untuk fokus pada pekerjaan yang sebenarnya. Dengan memanfaatkan tingkat kinerja ini, mesin pencari baru kami memastikan bahwa Anda dapat dengan cepat mengakses informasi yang Anda butuhkan, tanpa mengganggu alur kerja Anda.

Untuk mencapai tujuan pencarian yang sangat cepat, kami memanfaatkan teknologi open-source terbaru. Mesin pencari kami dibangun di atas MeiliSearch, layanan pencarian open-source populer yang menggunakan pemrosesan bahasa alami dan pencarian vektor untuk dengan cepat menemukan hasil yang relevan. Selain itu, kami menerapkan penyimpanan dalam memori, yang memungkinkan kami menyimpan data yang sering diakses di RAM, mengurangi waktu yang dibutuhkan untuk mengembalikan hasil pencarian. Kombinasi MeiliSearch dan penyimpanan dalam memori memungkinkan mesin pencari kami memberikan hasil dalam milidetik, sehingga Anda dapat dengan cepat menemukan apa yang Anda butuhkan tanpa harus memikirkan teknologi yang mendasarinya.

Bilah pencarian baru terletak dengan nyaman di bilah navigasi, memungkinkan Anda untuk mulai mencari segera. Untuk pengalaman pencarian yang lebih mendetail, cukup tekan tombol Tab saat mencari untuk mengakses halaman pencarian penuh. Selain itu, Anda dapat dengan cepat mengaktifkan fungsi pencarian dari mana saja menggunakan pintasan CMD/Ctrl+K, membuatnya lebih mudah untuk menemukan apa yang Anda butuhkan.

<video autoplay loop muted playsinline>
  <source src="/videos/search-demo.mp4" type="video/mp4">
</video>


## Perkembangan Mendatang

Ini baru permulaan. Sekarang kami memiliki infrastruktur pencarian generasi berikutnya, kami dapat melakukan beberapa hal yang sangat menarik di masa depan.

Selanjutnya adalah pencarian semantik, yang merupakan peningkatan signifikan dari pencarian kata kunci biasa. Izinkan kami menjelaskan.

Fitur ini akan memungkinkan mesin pencari untuk memahami konteks dari kueri Anda. Misalnya, mencari "laut" akan mengambil dokumen yang relevan meskipun frasa yang tepat tidak digunakan. Anda mungkin berpikir "tapi saya mengetik 'samudera' sebagai gantinya!" - dan Anda benar. Mesin pencari juga akan memahami kesamaan antara "laut" dan "samudera", dan mengembalikan dokumen yang relevan meskipun frasa yang tepat tidak digunakan. Fitur ini sangat berguna saat mencari dokumen yang mengandung istilah teknis, akronim, atau hanya kata-kata umum yang memiliki berbagai variasi atau kesalahan ketik.

Fitur lain yang akan datang adalah kemampuan untuk mencari gambar berdasarkan kontennya. Untuk mencapai ini, kami akan memproses setiap gambar di proyek Anda, membuat embedding untuk masing-masing. Dalam istilah tingkat tinggi, embedding adalah sekumpulan koordinat matematis yang sesuai dengan makna sebuah gambar. Ini berarti bahwa semua gambar dapat dicari berdasarkan apa yang mereka miliki, terlepas dari nama file atau metadata mereka. Bayangkan mencari "diagram alur" dan menemukan semua gambar yang terkait dengan diagram alur, *tanpa mempedulikan nama file mereka.*
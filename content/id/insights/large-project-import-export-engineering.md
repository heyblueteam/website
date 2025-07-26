---
title:  Meningkatkan Impor & Ekspor CSV hingga 250.000+ Rekor
description: Temukan bagaimana Blue meningkatkan impor dan ekspor CSV 10x menggunakan Rust dan arsitektur yang dapat diskalakan serta pilihan teknologi strategis dalam B2B SaaS.
category: "Engineering"
date: 2024-07-18
---


Di Blue, kami [terus mendorong batasan](/platform/roadmap) dari apa yang mungkin dalam perangkat lunak manajemen proyek. Selama bertahun-tahun, kami telah [meluncurkan ratusan fitur](/platform/changelog)

Prestasi rekayasa terbaru kami? 

Perombakan total dari sistem [impor CSV](https://documentation.blue.cc/integrations/csv-import) dan [ekspor](https://documentation.blue.cc/integrations/csv-export) kami, yang secara dramatis meningkatkan kinerja dan skalabilitas. 

Posting ini membawa Anda ke belakang layar tentang bagaimana kami mengatasi tantangan ini, teknologi yang kami gunakan, dan hasil mengesankan yang kami capai.

Hal yang paling menarik di sini adalah bahwa kami harus melangkah keluar dari [tumpukan teknologi](https://sop.blue.cc/product/technology-stack) kami yang biasa untuk mencapai hasil yang kami inginkan. Ini adalah keputusan yang harus diambil dengan hati-hati, karena dampak jangka panjangnya bisa sangat parah dalam hal utang teknologi dan overhead pemeliharaan jangka panjang. 

<video autoplay loop muted playsinline>
  <source src="/videos/import-export-video.mp4" type="video/mp4">
</video>

## Skala untuk Kebutuhan Perusahaan

Perjalanan kami dimulai dengan permintaan dari pelanggan perusahaan di industri acara. Klien ini menggunakan Blue sebagai pusat utama mereka untuk mengelola daftar acara, tempat, dan pembicara yang sangat besar, mengintegrasikannya secara mulus dengan situs web mereka. 

Bagi mereka, Blue bukan hanya alat — ini adalah sumber kebenaran tunggal untuk seluruh operasi mereka.

Meskipun kami selalu bangga mendengar bahwa pelanggan menggunakan kami untuk kebutuhan yang sangat penting, ada juga tanggung jawab besar di pihak kami untuk memastikan sistem yang cepat dan andal.

Saat pelanggan ini meningkatkan operasi mereka, mereka menghadapi hambatan yang signifikan: **mengimpor dan mengekspor file CSV besar yang berisi 100.000 hingga 200.000+ rekaman.**

Ini berada di luar kemampuan sistem kami pada saat itu. Faktanya, sistem impor/ekspor kami sebelumnya sudah kesulitan dengan impor dan ekspor yang berisi lebih dari 10.000 hingga 20.000 rekaman! Jadi 200.000+ rekaman adalah hal yang tidak mungkin. 

Pengguna mengalami waktu tunggu yang sangat lama, dan dalam beberapa kasus, impor atau ekspor akan *gagal diselesaikan sama sekali.* Ini sangat mempengaruhi operasi mereka karena mereka bergantung pada impor dan ekspor harian untuk mengelola aspek tertentu dari operasi mereka. 

> Multi-tenancy adalah arsitektur di mana satu instance perangkat lunak melayani beberapa pelanggan (penyewa). Meskipun efisien, ini memerlukan manajemen sumber daya yang hati-hati untuk memastikan tindakan satu penyewa tidak berdampak negatif pada penyewa lainnya.

Dan keterbatasan ini tidak hanya mempengaruhi klien tertentu ini. 

Karena arsitektur multi-penyewa kami—di mana beberapa pelanggan berbagi infrastruktur yang sama—impor atau ekspor yang memerlukan sumber daya intensif dapat memperlambat operasi untuk pengguna lain, yang dalam praktiknya sering terjadi. 

Seperti biasa, kami melakukan analisis build vs buy, untuk memahami apakah kami harus menghabiskan waktu untuk meningkatkan sistem kami sendiri atau membeli sistem dari orang lain. Kami melihat berbagai kemungkinan.

Vendor yang menonjol adalah penyedia SaaS bernama [Flatfile](https://flatfile.com/). Sistem dan kemampuannya terlihat persis seperti yang kami butuhkan. 

Namun, setelah meninjau [harga](https://flatfile.com/pricing/) mereka, kami memutuskan bahwa ini akan menjadi solusi yang sangat mahal untuk aplikasi skala kami — *$2/file mulai bertambah dengan sangat cepat!* —dan lebih baik untuk memperluas mesin impor/ekspor CSV bawaan kami. 

Untuk mengatasi tantangan ini, kami membuat keputusan berani: memperkenalkan Rust ke dalam tumpukan teknologi Javascript utama kami. Bahasa pemrograman sistem ini, yang dikenal karena kinerja dan keamanannya, adalah alat yang sempurna untuk kebutuhan pemrosesan CSV dan pemetaan data yang kritis untuk kinerja kami.

Berikut adalah bagaimana kami mendekati solusi ini.

### Memperkenalkan Layanan Latar Belakang

Dasar dari solusi kami adalah pengenalan layanan latar belakang untuk menangani tugas-tugas yang memerlukan sumber daya intensif. Pendekatan ini memungkinkan kami untuk mengalihkan pemrosesan berat dari server utama kami, secara signifikan meningkatkan kinerja keseluruhan sistem.
Arsitektur layanan latar belakang kami dirancang dengan mempertimbangkan skalabilitas. Seperti semua komponen infrastruktur kami, layanan ini secara otomatis diskalakan berdasarkan permintaan. 

Ini berarti bahwa selama waktu puncak, ketika beberapa impor atau ekspor besar diproses secara bersamaan, sistem secara otomatis mengalokasikan lebih banyak sumber daya untuk menangani beban yang meningkat. Sebaliknya, selama periode yang lebih tenang, ia menyesuaikan diri untuk mengoptimalkan penggunaan sumber daya.

Arsitektur layanan latar belakang yang dapat diskalakan ini telah menguntungkan Blue tidak hanya untuk impor & ekspor CSV. Seiring waktu, kami telah memindahkan sejumlah besar fitur ke dalam layanan latar belakang untuk mengurangi beban pada server utama kami:

- **[Perhitungan Formula](https://documentation.blue.cc/custom-fields/formula)**: Mengalihkan operasi matematika kompleks untuk memastikan pembaruan cepat dari bidang turunan tanpa mempengaruhi kinerja server utama.
- **[Dasbor/Grafik](/platform/features/dashboards)**: Memproses dataset besar di latar belakang untuk menghasilkan visualisasi terkini tanpa memperlambat antarmuka pengguna.
- **[Indeks Pencarian](https://documentation.blue.cc/projects/search)**: Secara terus-menerus memperbarui indeks pencarian di latar belakang, memastikan hasil pencarian yang cepat dan akurat tanpa mempengaruhi kinerja sistem.
- **[Menyalin Proyek](https://documentation.blue.cc/projects/copying-projects)**: Menangani replikasi proyek besar dan kompleks di latar belakang, memungkinkan pengguna untuk terus bekerja sementara salinan sedang dibuat.
- **[Automasi Manajemen Proyek](/platform/features/automations)**: Menjalankan alur kerja otomatis yang ditentukan pengguna di latar belakang, memastikan tindakan tepat waktu tanpa memblokir operasi lainnya.
- **[Rekaman Berulang](https://documentation.blue.cc/records/repeat)**: Menghasilkan tugas atau acara berulang di latar belakang, menjaga akurasi jadwal tanpa membebani aplikasi utama.
- **[Bidang Kustom Durasi Waktu](https://documentation.blue.cc/custom-fields/duration)**: Secara terus-menerus menghitung dan memperbarui selisih waktu antara dua peristiwa di Blue, memberikan data durasi waktu secara real-time tanpa mempengaruhi responsivitas sistem.

## Modul Rust Baru untuk Pemrosesan Data

Inti dari solusi pemrosesan CSV kami adalah modul Rust kustom. Meskipun ini menandai usaha pertama kami di luar tumpukan teknologi inti kami yang berbasis Javascript, keputusan untuk menggunakan Rust didorong oleh kinerjanya yang luar biasa dalam operasi bersamaan dan tugas pemrosesan file.

Kekuatan Rust sangat sesuai dengan tuntutan pemrosesan CSV dan pemetaan data. Abstraksi tanpa biaya memungkinkan pemrograman tingkat tinggi tanpa mengorbankan kinerja, sementara model kepemilikannya memastikan keamanan memori tanpa perlu pengumpulan sampah. Fitur-fitur ini membuat Rust sangat mahir dalam menangani dataset besar secara efisien dan aman.

Untuk pemrosesan CSV, kami memanfaatkan crate csv Rust, yang menawarkan pembacaan dan penulisan data CSV berkinerja tinggi. Kami menggabungkan ini dengan logika pemetaan data kustom untuk memastikan integrasi yang mulus dengan struktur data Blue.

Kurva pembelajaran untuk Rust cukup curam tetapi dapat dikelola. Tim kami menghabiskan sekitar dua minggu untuk belajar intensif tentang ini.

Perbaikan yang kami capai sangat mengesankan:

![](/insights/import-export.png)

Sistem baru kami dapat memproses jumlah rekaman yang sama yang dapat diproses oleh sistem lama kami dalam 15 menit dalam waktu sekitar 30 detik. 

## Interaksi Server Web dan Database

Untuk komponen server web dari implementasi Rust kami, kami memilih Rocket sebagai kerangka kerja kami. Rocket menonjol karena kombinasi kinerja dan fitur yang ramah pengembang. Pengetikan statis dan pemeriksaan waktu kompilasi sangat sesuai dengan prinsip keamanan Rust, membantu kami menangkap masalah potensial lebih awal dalam proses pengembangan.
Di sisi database, kami memilih SQLx. Perpustakaan SQL async ini untuk Rust menawarkan beberapa keuntungan yang menjadikannya ideal untuk kebutuhan kami:

- SQL yang aman tipe: SQLx memungkinkan kami untuk menulis SQL mentah dengan kueri yang diperiksa pada waktu kompilasi, memastikan keamanan tipe tanpa mengorbankan kinerja.
- Dukungan async: Ini sangat sesuai dengan Rocket dan kebutuhan kami untuk operasi database yang efisien dan tidak memblokir.
- Agnostik database: Meskipun kami terutama menggunakan [AWS Aurora](https://aws.amazon.com/rds/aurora/), yang kompatibel dengan MySQL, dukungan SQLx untuk beberapa database memberi kami fleksibilitas untuk masa depan jika kami memutuskan untuk berubah. 

## Optimisasi Pengelompokan

Perjalanan kami menuju konfigurasi pengelompokan yang optimal adalah salah satu pengujian ketat dan analisis yang hati-hati. Kami menjalankan benchmark ekstensif dengan berbagai kombinasi transaksi bersamaan dan ukuran chunk, mengukur tidak hanya kecepatan mentah tetapi juga pemanfaatan sumber daya dan stabilitas sistem.

Proses ini melibatkan pembuatan dataset uji dengan ukuran dan kompleksitas yang bervariasi, mensimulasikan pola penggunaan dunia nyata. Kami kemudian menjalankan dataset ini melalui sistem kami, menyesuaikan jumlah transaksi bersamaan dan ukuran chunk untuk setiap percobaan.

Setelah menganalisis hasilnya, kami menemukan bahwa memproses 5 transaksi bersamaan dengan ukuran chunk 500 rekaman memberikan keseimbangan terbaik antara kecepatan dan pemanfaatan sumber daya. Konfigurasi ini memungkinkan kami untuk mempertahankan throughput tinggi tanpa membebani database kami atau mengonsumsi memori yang berlebihan.

Menariknya, kami menemukan bahwa meningkatkan concurrency di atas 5 transaksi tidak menghasilkan peningkatan kinerja yang signifikan dan kadang-kadang menyebabkan peningkatan kontensi database. Demikian pula, ukuran chunk yang lebih besar meningkatkan kecepatan mentah tetapi dengan biaya penggunaan memori yang lebih tinggi dan waktu respons yang lebih lama untuk impor/ekspor kecil hingga menengah.

## Ekspor CSV melalui Tautan Email

Bagian terakhir dari solusi kami menangani tantangan mengirimkan file ekspor besar kepada pengguna. Alih-alih menyediakan unduhan langsung dari aplikasi web kami, yang dapat menyebabkan masalah timeout dan meningkatkan beban server, kami mengimplementasikan sistem tautan unduhan yang dikirim melalui email.

Ketika seorang pengguna memulai ekspor besar, sistem kami memproses permintaan di latar belakang. Setelah selesai, alih-alih mempertahankan koneksi terbuka atau menyimpan file di server web kami, kami mengunggah file ke lokasi penyimpanan sementara yang aman. Kami kemudian menghasilkan tautan unduhan unik dan aman dan mengirimkannya melalui email kepada pengguna.

Tautan unduhan ini berlaku selama 2 jam, menciptakan keseimbangan antara kenyamanan pengguna dan keamanan informasi. Kerangka waktu ini memberi pengguna kesempatan yang cukup untuk mengambil data mereka sambil memastikan bahwa informasi sensitif tidak dibiarkan dapat diakses selamanya.

Keamanan tautan unduhan ini adalah prioritas utama dalam desain kami. Setiap tautan adalah:

- Unik dan dihasilkan secara acak, membuatnya hampir tidak mungkin untuk ditebak
- Berlaku hanya selama 2 jam
- Enkripsi saat transit, memastikan keamanan data saat diunduh

Pendekatan ini menawarkan beberapa manfaat:

- Mengurangi beban pada server web kami, karena mereka tidak perlu menangani unduhan file besar secara langsung
- Meningkatkan pengalaman pengguna, terutama bagi pengguna dengan koneksi internet yang lebih lambat yang mungkin menghadapi masalah timeout browser dengan unduhan langsung
- Menyediakan solusi yang lebih andal untuk ekspor yang sangat besar yang mungkin melebihi batas waktu web yang biasa

Umpan balik pengguna tentang fitur ini sangat positif, dengan banyak yang menghargai fleksibilitas yang ditawarkannya dalam mengelola ekspor data besar.

## Mengekspor Data yang Difilter

Perbaikan lain yang jelas adalah memungkinkan pengguna untuk hanya mengekspor data yang sudah difilter dalam tampilan proyek mereka. Ini berarti jika ada tag aktif "prioritas", maka hanya rekaman yang memiliki tag ini yang akan masuk ke ekspor CSV. Ini berarti lebih sedikit waktu untuk memanipulasi data di Excel untuk menyaring hal-hal yang tidak penting, dan juga membantu kami mengurangi jumlah baris yang diproses.

## Melihat ke Depan

Meskipun kami tidak memiliki rencana segera untuk memperluas penggunaan Rust kami, proyek ini telah menunjukkan potensi teknologi ini untuk operasi yang kritis terhadap kinerja. Ini adalah opsi menarik yang sekarang kami miliki dalam toolkit kami untuk kebutuhan optimasi di masa depan. Perombakan impor dan ekspor CSV ini selaras sempurna dengan komitmen Blue terhadap skalabilitas. 

Kami berkomitmen untuk menyediakan platform yang tumbuh bersama pelanggan kami, menangani kebutuhan data mereka yang berkembang tanpa mengorbankan kinerja.

Keputusan untuk memperkenalkan Rust ke dalam tumpukan teknologi kami tidak diambil dengan ringan. Ini mengangkat pertanyaan penting yang dihadapi banyak tim rekayasa: Kapan tepat untuk melangkah keluar dari tumpukan teknologi inti Anda, dan kapan Anda harus tetap menggunakan alat yang sudah dikenal?

Tidak ada jawaban satu ukuran untuk semua, tetapi di Blue, kami telah mengembangkan kerangka kerja untuk membuat keputusan penting ini:

- **Pendekatan Masalah Pertama:** Kami selalu mulai dengan mendefinisikan masalah yang ingin kami selesaikan dengan jelas. Dalam hal ini, kami perlu secara dramatis meningkatkan kinerja impor dan ekspor CSV untuk dataset besar.
- **Mengeksplorasi Solusi yang Ada:** Sebelum melihat di luar tumpukan inti kami, kami secara menyeluruh mengeksplorasi apa yang dapat dicapai dengan teknologi yang ada. Ini sering melibatkan profiling, optimasi, dan memikirkan kembali pendekatan kami dalam batasan yang sudah dikenal.
- **Mengkuantifikasi Potensi Keuntungan:** Jika kami mempertimbangkan teknologi baru, kami perlu dapat menjelaskan dengan jelas dan, idealnya, mengkuantifikasi manfaatnya. Untuk proyek CSV kami, kami memproyeksikan peningkatan kecepatan pemrosesan yang signifikan.
- **Menilai Biaya:** Memperkenalkan teknologi baru bukan hanya tentang proyek yang segera. Kami mempertimbangkan biaya jangka panjang:
  - Kurva pembelajaran untuk tim
  - Pemeliharaan dan dukungan yang berkelanjutan
  - Potensi komplikasi dalam penerapan dan operasi
  - Dampak pada perekrutan dan komposisi tim
- **Pembatasan dan Integrasi:** Jika kami memperkenalkan teknologi baru, kami bertujuan untuk membatasinya pada bagian sistem kami yang spesifik dan terdefinisi dengan baik. Kami juga memastikan kami memiliki rencana yang jelas tentang bagaimana itu akan terintegrasi dengan tumpukan yang ada.
- **Mempersiapkan Masa Depan:** Kami mempertimbangkan apakah pilihan teknologi ini membuka peluang di masa depan atau jika itu mungkin menjebak kami dalam situasi yang sulit.

Salah satu risiko utama dari sering mengadopsi teknologi baru adalah berakhir dengan apa yang kami sebut *"zoo teknologi"* - ekosistem yang terfragmentasi di mana bagian-bagian berbeda dari aplikasi Anda ditulis dalam bahasa atau kerangka kerja yang berbeda, memerlukan berbagai keterampilan khusus untuk dipelihara.


## Kesimpulan

Proyek ini mencerminkan pendekatan rekayasa Blue: *kami tidak takut untuk melangkah keluar dari zona nyaman kami dan mengadopsi teknologi baru ketika itu berarti memberikan pengalaman yang jauh lebih baik bagi pengguna kami.* 

Dengan membayangkan kembali proses impor dan ekspor CSV kami, kami tidak hanya menyelesaikan kebutuhan mendesak untuk satu klien perusahaan tetapi juga meningkatkan pengalaman bagi semua pengguna kami yang berurusan dengan dataset besar.

Saat kami terus mendorong batasan dari apa yang mungkin dalam [perangkat lunak manajemen proyek](/solutions/use-case/project-management), kami bersemangat untuk menghadapi lebih banyak tantangan seperti ini. 

Tetap disini untuk lebih banyak [penjelasan mendalam tentang rekayasa yang mendukung Blue!](/insights/engineering-blog)
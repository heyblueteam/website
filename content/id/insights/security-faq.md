---
title: FAQ tentang Keamanan Blue
description: Ini adalah daftar pertanyaan yang paling sering diajukan tentang protokol dan praktik keamanan di Blue.
category: "FAQ"
date: 2024-07-19
---


Misi kami adalah mengorganisir pekerjaan di seluruh dunia dengan membangun platform manajemen proyek terbaik di planet ini.

Pusat untuk mencapai misi ini adalah memastikan bahwa platform kami aman, dapat diandalkan, dan terpercaya. Kami memahami bahwa untuk menjadi sumber kebenaran tunggal Anda, Blue harus melindungi data bisnis sensitif Anda dari ancaman eksternal, kehilangan data, dan waktu henti.

Ini berarti bahwa kami menganggap serius keamanan di Blue.

Ketika kami memikirkan tentang keamanan, kami mempertimbangkan pendekatan holistik yang berfokus pada tiga area kunci:

1.  **Keamanan Infrastruktur & Jaringan**: Memastikan bahwa sistem fisik dan virtual kami terlindungi dari ancaman eksternal dan akses yang tidak sah.
2.  **Keamanan Perangkat Lunak**: Berfokus pada keamanan kode itu sendiri, termasuk praktik pengkodean yang aman, tinjauan kode secara berkala, dan manajemen kerentanan.
3.  **Keamanan Platform**: Termasuk fitur-fitur dalam Blue, seperti [kontrol akses yang canggih](/platform/features/user-permissions), memastikan bahwa proyek bersifat pribadi secara default, dan langkah-langkah lain untuk melindungi data dan privasi pengguna.


## Seberapa skalabel Blue?

Ini adalah pertanyaan penting, karena Anda menginginkan sistem yang dapat *tumbuh* bersama Anda. Anda tidak ingin harus beralih platform manajemen proyek dan proses dalam enam atau dua belas bulan.

Kami memilih penyedia platform dengan hati-hati, untuk memastikan bahwa mereka dapat menangani beban kerja yang menuntut dari pelanggan kami. Kami menggunakan layanan cloud dari beberapa penyedia cloud terkemuka di dunia yang mendukung perusahaan seperti [Spotify](https://spotify.com) dan [Netflix](https://netflix.com), yang memiliki lalu lintas beberapa urutan magnitudo lebih banyak daripada kami.

Penyedia cloud utama yang kami gunakan adalah:

- **[Cloudflare](https://cloudflare.com)**: Kami mengelola DNS (Domain Name Service) melalui Cloudflare serta situs web pemasaran kami yang berjalan di [Cloudflare Pages](https://pages.cloudflare.com/).
- **[Amazon Web Services](https://aws.amazon.com/)**: Kami menggunakan AWS untuk basis data kami, yang adalah [Aurora](https://aws.amazon.com/rds/aurora/), untuk penyimpanan file melalui [Simple Storage Service (S3)](https://aws.amazon.com/s3/), dan juga untuk mengirim email melalui [Simple Email Service (SES)](https://aws.amazon.com/ses/)
- **[Render](https://render.com)**: Kami menggunakan Render untuk server front-end kami, server aplikasi/API, layanan latar belakang kami, sistem antrian, dan basis data Redis. Menariknya, Render sebenarnya dibangun *di atas* AWS! 


## Seberapa aman file di Blue? 

Mari kita mulai dengan penyimpanan data. File kami dihosting di [AWS S3](https://aws.amazon.com/s3/), yang merupakan penyimpanan objek cloud paling populer di dunia dengan skalabilitas, ketersediaan data, keamanan, dan kinerja yang terdepan di industri.

Kami memiliki ketersediaan file 99,99% dan daya tahan tinggi 99,999999999%.

Mari kita uraikan apa artinya ini.

Ketersediaan mengacu pada jumlah waktu data beroperasi dan dapat diakses. Ketersediaan file 99,99% berarti bahwa kami dapat mengharapkan file tidak tersedia tidak lebih dari sekitar 8,76 jam per tahun.

Daya tahan mengacu pada kemungkinan bahwa data tetap utuh dan tidak korup seiring waktu. Tingkat daya tahan ini berarti kami dapat mengharapkan untuk kehilangan tidak lebih dari satu file dari 10 miliar file yang diunggah, berkat redundansi yang luas dan replikasi data di beberapa pusat data.

Kami menggunakan [S3 Intelligent-Tiering](https://aws.amazon.com/s3/storage-classes/intelligent-tiering/) untuk secara otomatis memindahkan file ke kelas penyimpanan yang berbeda berdasarkan frekuensi akses. Berdasarkan pola aktivitas dari ratusan ribu proyek, kami memperhatikan bahwa sebagian besar file diakses dalam pola yang menyerupai kurva exponential backoff. Ini berarti bahwa sebagian besar file diakses sangat sering dalam beberapa hari pertama, dan kemudian diakses semakin jarang. Ini memungkinkan kami untuk memindahkan file yang lebih tua ke penyimpanan yang lebih lambat, tetapi jauh lebih murah, tanpa mempengaruhi pengalaman pengguna dengan cara yang berarti.

Penghematan biaya untuk ini sangat signifikan. S3 Standard-Infrequent Access (S3 Standard-IA) sekitar 1,84 kali lebih murah daripada S3 Standard. Ini berarti bahwa untuk setiap dolar yang kami habiskan untuk S3 Standard, kami hanya menghabiskan sekitar 54 sen untuk S3 Standard-IA untuk jumlah data yang sama disimpan.

| Fitur                    | S3 Standard             | S3 Standard-IA       |
|--------------------------|-------------------------|-----------------------|
| Biaya Penyimpanan        | $0.023 - $0.021 per GB  | $0.0125 per GB        |
| Biaya Permintaan (PUT, dll.) | $0.005 per 1.000 permintaan | $0.01 per 1.000 permintaan |
| Biaya Permintaan (GET)   | $0.0004 per 1.000 permintaan | $0.001 per 1.000 permintaan |
| Biaya Pengambilan Data   | $0.00 per GB            | $0.01 per GB          |


File yang Anda unggah melalui Blue dienkripsi baik dalam perjalanan maupun saat istirahat. Data yang ditransfer ke dan dari Amazon S3 diamankan menggunakan [Transport Layer Security (TLS)](https://www.internetsociety.org/deploy360/tls/basics), melindungi dari [penyadapan](https://en.wikipedia.org/wiki/Network_eavesdropping) dan [serangan man-in-the-middle](https://en.wikipedia.org/wiki/Man-in-the-middle_attack). Untuk enkripsi saat istirahat, Amazon S3 menggunakan Server-Side Encryption (SSE-S3), yang secara otomatis mengenkripsi semua unggahan baru dengan enkripsi AES-256, dengan Amazon mengelola kunci enkripsi. Ini memastikan data Anda tetap aman sepanjang siklus hidupnya.

## Bagaimana dengan data non-file? 

Basis data kami didukung oleh [AWS Aurora](https://aws.amazon.com/rds/aurora/), layanan basis data relasional modern yang memastikan kinerja tinggi, ketersediaan, dan keamanan untuk data Anda.

Data di Aurora dienkripsi baik dalam perjalanan maupun saat istirahat. Kami menggunakan SSL (AES-256) untuk mengamankan koneksi antara instance basis data Anda dan aplikasi Anda, melindungi data selama transfer. Untuk enkripsi saat istirahat, Aurora menggunakan kunci yang dikelola melalui AWS Key Management Service (KMS), memastikan bahwa semua data yang disimpan, termasuk cadangan otomatis, snapshot, dan replika, dienkripsi dan dilindungi.

Aurora memiliki sistem penyimpanan yang terdistribusi, toleran kesalahan, dan dapat menyembuhkan diri sendiri. Sistem ini terpisah dari sumber daya komputasi dan dapat otomatis menskalakan hingga 128 TiB per instance basis data. Data direplikasi di tiga [Availability Zones](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) (AZ), memberikan ketahanan terhadap kehilangan data dan memastikan ketersediaan tinggi. Dalam kasus kerusakan basis data, Aurora mengurangi waktu pemulihan menjadi kurang dari 60 detik, memastikan gangguan minimal.

Blue secara terus-menerus mencadangkan basis data kami ke Amazon S3, memungkinkan pemulihan pada titik waktu tertentu. Ini berarti kami dapat mengembalikan basis data master Blue ke waktu tertentu dalam lima menit terakhir, memastikan bahwa data Anda selalu dapat dipulihkan. Kami juga mengambil snapshot reguler dari basis data untuk periode retensi cadangan yang lebih lama.

Sebagai layanan yang sepenuhnya dikelola, Aurora mengotomatiskan tugas administrasi yang memakan waktu seperti penyediaan perangkat keras, pengaturan basis data, patching, dan cadangan. Ini mengurangi beban operasional dan memastikan bahwa basis data kami selalu diperbarui dengan patch keamanan dan perbaikan kinerja terbaru.

Jika kami lebih efisien, kami dapat meneruskan penghematan biaya kami kepada pelanggan kami dengan [harga terdepan di industri](/pricing).

Aurora mematuhi berbagai standar industri seperti HIPAA, GDPR, dan SOC 2, memastikan bahwa praktik manajemen data Anda memenuhi persyaratan regulasi yang ketat. Audit keamanan reguler dan integrasi dengan [Amazon GuardDuty](https://aws.amazon.com/guardduty/) membantu mendeteksi dan mengurangi potensi ancaman keamanan.

## Bagaimana Blue memastikan keamanan login?

Blue menggunakan [tautan ajaib melalui email](https://documentation.blue.cc/user-management/magic-links) untuk memberikan akses yang aman dan nyaman ke akun Anda, menghilangkan kebutuhan akan kata sandi tradisional.

Pendekatan ini secara signifikan meningkatkan keamanan dengan mengurangi ancaman umum yang terkait dengan login berbasis kata sandi. Dengan menghilangkan kata sandi, tautan ajaib melindungi dari serangan phishing dan pencurian kata sandi, *karena tidak ada kata sandi untuk dicuri atau dieksploitasi.*

Setiap tautan ajaib hanya berlaku untuk satu sesi login, mengurangi risiko akses yang tidak sah. Selain itu, tautan ini kedaluwarsa setelah 15 menit, memastikan bahwa tautan yang tidak digunakan tidak dapat dieksploitasi, lebih meningkatkan keamanan.

Kenyamanan yang ditawarkan oleh tautan ajaib juga patut dicatat. Tautan ajaib memberikan pengalaman login yang tanpa repot, memungkinkan Anda mengakses akun Anda *tanpa* perlu mengingat kata sandi yang rumit.

Ini tidak hanya menyederhanakan proses login tetapi juga mencegah pelanggaran keamanan yang terjadi ketika kata sandi digunakan kembali di berbagai layanan. Banyak pengguna cenderung menggunakan kata sandi yang sama di berbagai platform, yang berarti pelanggaran keamanan pada satu layanan dapat membahayakan akun mereka di layanan lain, termasuk Blue. Dengan menggunakan tautan ajaib, keamanan Blue tidak bergantung pada praktik keamanan layanan lain, memberikan lapisan perlindungan yang lebih kuat dan independen bagi pengguna kami.

Ketika Anda meminta untuk login ke akun Blue Anda, URL login unik dikirim ke email Anda. Mengklik tautan ini akan langsung masuk ke akun Anda. Tautan ini dirancang untuk kedaluwarsa setelah satu kali penggunaan atau setelah 15 menit, mana yang lebih dulu, menambahkan lapisan keamanan tambahan. Dengan menggunakan tautan ajaib, Blue memastikan bahwa proses login Anda aman dan ramah pengguna, memberikan ketenangan pikiran dan kenyamanan.

## Bagaimana saya bisa memeriksa keandalan dan waktu aktif Blue?

Di Blue, kami berkomitmen untuk mempertahankan tingkat keandalan dan transparansi yang tinggi bagi pengguna kami. Untuk memberikan visibilitas ke dalam kinerja platform kami, kami menawarkan [halaman status sistem khusus](https://status.blue.cc) yang juga terhubung dari footer di setiap halaman situs web kami.

![](/insights/status-page.png)

Halaman ini menampilkan data ketersediaan historis kami, memungkinkan Anda melihat seberapa konsisten layanan kami tersedia seiring waktu. Selain itu, halaman status mencakup laporan insiden terperinci, memberikan transparansi tentang masalah masa lalu, dampaknya, dan langkah-langkah yang telah kami ambil untuk menyelesaikannya dan mencegah terjadinya di masa depan.
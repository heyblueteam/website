---
title: Membangun Mesin Izin Kustom Blue
description: Lihat di balik layar bersama tim teknik Blue saat mereka menjelaskan cara membangun fitur pengkategorian dan penandaan otomatis yang didukung AI.
category: "Engineering"
date: 2024-07-25
---


Manajemen proyek dan proses yang efektif sangat penting bagi organisasi dari semua ukuran.

Di Blue, [kami telah menjadikan misi kami](/about) untuk mengorganisir pekerjaan dunia dengan membangun platform manajemen proyek terbaik di planet ini—sederhana, kuat, fleksibel, dan terjangkau untuk semua.

Ini berarti bahwa platform kami harus dapat beradaptasi dengan kebutuhan unik setiap tim. Hari ini, kami sangat senang untuk membongkar salah satu fitur paling kuat kami: Izin Kustom.

Alat manajemen proyek adalah tulang punggung alur kerja modern, menyimpan data sensitif, komunikasi penting, dan rencana strategis. Oleh karena itu, kemampuan untuk mengontrol akses ke informasi ini dengan cermat bukan hanya sebuah kemewahan—ini adalah kebutuhan.

<video autoplay loop muted playsinline>
  <source src="/videos/user-roles.mp4" type="video/mp4">
</video>

Izin kustom memainkan peran penting dalam platform B2B SaaS, terutama dalam alat manajemen proyek, di mana keseimbangan antara kolaborasi dan keamanan dapat menentukan keberhasilan suatu proyek.

Namun, di sinilah Blue mengambil pendekatan yang berbeda: **kami percaya bahwa fitur tingkat perusahaan tidak seharusnya hanya diperuntukkan bagi anggaran berukuran perusahaan.**

Di era di mana AI memberdayakan tim kecil untuk beroperasi pada skala yang belum pernah terjadi sebelumnya, mengapa keamanan yang kuat dan kustomisasi harus di luar jangkauan?

Dalam pandangan di balik layar ini, kami akan menjelajahi bagaimana kami mengembangkan fitur Izin Kustom kami, menantang status quo dari tingkat harga SaaS, dan membawa opsi keamanan yang kuat dan fleksibel kepada bisnis dari semua ukuran.

Apakah Anda sebuah startup dengan impian besar atau pemain mapan yang ingin mengoptimalkan proses Anda, izin kustom dapat memungkinkan kasus penggunaan baru yang bahkan tidak Anda ketahui mungkin.

## Memahami Izin Pengguna Kustom

Sebelum kita menyelami perjalanan kami dalam mengembangkan izin kustom untuk Blue, mari kita luangkan waktu sejenak untuk memahami apa itu izin pengguna kustom dan mengapa mereka sangat penting dalam perangkat lunak manajemen proyek.

Izin pengguna kustom mengacu pada kemampuan untuk menyesuaikan hak akses untuk pengguna individu atau kelompok dalam sistem perangkat lunak. Alih-alih bergantung pada peran yang telah ditentukan sebelumnya dengan set izin tetap, izin kustom memungkinkan administrator untuk membuat profil akses yang sangat spesifik yang selaras dengan struktur organisasi dan kebutuhan alur kerja mereka.

Dalam konteks perangkat lunak manajemen proyek seperti Blue, izin kustom mencakup:

* **Kontrol akses yang terperinci**: Menentukan siapa yang dapat melihat, mengedit, atau menghapus jenis data proyek tertentu.
* **Pembatasan berbasis fitur**: Mengaktifkan atau menonaktifkan fitur tertentu untuk pengguna atau tim tertentu.
* **Tingkat sensitivitas data**: Mengatur tingkat akses yang bervariasi untuk informasi sensitif dalam proyek.
* **Izin spesifik alur kerja**: Menyelaraskan kemampuan pengguna dengan tahap atau aspek tertentu dari alur kerja proyek Anda.

Pentingnya izin kustom dalam manajemen proyek tidak dapat dilebih-lebihkan:

* **Keamanan yang ditingkatkan**: Dengan memberikan pengguna hanya akses yang mereka butuhkan, Anda mengurangi risiko pelanggaran data atau perubahan yang tidak sah.
* **Kepatuhan yang lebih baik**: Izin kustom membantu organisasi memenuhi persyaratan regulasi spesifik industri dengan mengontrol akses data.
* **Kolaborasi yang lebih efisien**: Tim dapat bekerja lebih efisien ketika setiap anggota memiliki tingkat akses yang tepat untuk menjalankan perannya tanpa batasan yang tidak perlu atau hak istimewa yang berlebihan.
* **Fleksibilitas untuk organisasi yang kompleks**: Seiring perusahaan tumbuh dan berkembang, izin kustom memungkinkan perangkat lunak untuk beradaptasi dengan struktur dan proses organisasi yang berubah.

## Mendapatkan YA

[Kami telah menulis sebelumnya](/insights/value-proposition-blue), bahwa setiap fitur di Blue harus menjadi **YA** yang **keras** sebelum kami memutuskan untuk membangunnya. Kami tidak memiliki kemewahan ratusan insinyur dan membuang waktu serta uang untuk membangun hal-hal yang tidak dibutuhkan pelanggan.

Dan jadi, jalan untuk menerapkan izin kustom di Blue bukanlah garis lurus. Seperti banyak fitur kuat lainnya, itu dimulai dengan kebutuhan yang jelas dari pengguna kami dan berkembang melalui pertimbangan dan perencanaan yang cermat.

Selama bertahun-tahun, pelanggan kami telah meminta kontrol yang lebih terperinci atas izin pengguna. Ketika organisasi dari semua ukuran mulai menangani proyek yang semakin kompleks dan sensitif, batasan dari kontrol akses berbasis peran standar kami menjadi jelas.

Startup kecil yang bekerja dengan klien eksternal, perusahaan menengah dengan proses persetujuan yang rumit, dan perusahaan besar dengan persyaratan kepatuhan yang ketat semuanya menyuarakan kebutuhan yang sama:

Lebih banyak fleksibilitas dalam mengelola akses pengguna.

Meskipun ada permintaan yang jelas, kami awalnya ragu untuk terjun ke pengembangan izin kustom.

Mengapa?

Kami memahami kompleksitas yang terlibat!

Izin kustom menyentuh setiap bagian dari sistem manajemen proyek, mulai dari antarmuka pengguna hingga struktur basis data. Kami tahu bahwa menerapkan fitur ini akan memerlukan perubahan signifikan pada arsitektur inti kami dan pertimbangan cermat terhadap implikasi kinerja.

Saat kami memeriksa lanskap, kami memperhatikan bahwa sangat sedikit pesaing kami yang mencoba menerapkan mesin izin kustom yang kuat seperti yang diminta pelanggan kami. Mereka yang melakukannya sering kali menyimpannya untuk rencana perusahaan tingkat tertinggi mereka.

Jelas mengapa: upaya pengembangan sangat besar, dan taruhannya tinggi.

Menerapkan izin kustom secara tidak benar dapat memperkenalkan bug kritis atau kerentanan keamanan, yang dapat membahayakan seluruh sistem. Kesadaran ini menegaskan besarnya tantangan yang kami pertimbangkan.

### Menantang Status Quo

Namun, saat kami terus tumbuh dan berkembang, kami mencapai kesadaran penting:

**Model SaaS tradisional yang menyimpan fitur kuat untuk pelanggan perusahaan tidak lagi masuk akal di lanskap bisnis saat ini.**

Pada tahun 2024, dengan kekuatan AI dan alat canggih, tim kecil dapat beroperasi pada skala dan kompleksitas yang menyaingi organisasi yang jauh lebih besar. Sebuah startup mungkin menangani data klien sensitif di berbagai negara. Sebuah agensi pemasaran kecil dapat mengelola puluhan proyek klien dengan berbagai persyaratan kerahasiaan. Bisnis-bisnis ini membutuhkan tingkat keamanan dan kustomisasi yang sama seperti *perusahaan besar mana pun*.

Kami bertanya pada diri sendiri: Mengapa ukuran tenaga kerja atau anggaran perusahaan harus menentukan kemampuan mereka untuk menjaga data mereka tetap aman dan proses mereka tetap efisien?

### Fitur Tingkat Perusahaan untuk Semua

Kesadaran ini membawa kami pada filosofi inti yang sekarang mendorong banyak pengembangan kami di Blue: Fitur tingkat perusahaan harus dapat diakses oleh bisnis dari semua ukuran.

Kami percaya bahwa:

- **Keamanan tidak seharusnya menjadi kemewahan.** Setiap perusahaan, terlepas dari ukuran, berhak mendapatkan alat untuk melindungi data dan proses mereka.
- **Fleksibilitas mendorong inovasi.** Dengan memberikan semua pengguna kami alat yang kuat, kami memungkinkan mereka untuk menciptakan alur kerja dan sistem yang mendorong industri mereka maju.
- **Pertumbuhan tidak seharusnya memerlukan perubahan platform.** Saat pelanggan kami tumbuh, alat mereka harus tumbuh dengan mereka tanpa hambatan.

Dengan pola pikir ini, kami memutuskan untuk menghadapi tantangan izin kustom secara langsung, berkomitmen untuk membuatnya tersedia bagi semua pengguna kami, bukan hanya mereka yang berada di rencana tingkat lebih tinggi.

Keputusan ini menempatkan kami pada jalur desain yang cermat, pengembangan iteratif, dan umpan balik pengguna yang berkelanjutan yang akhirnya mengarah pada fitur izin kustom yang kami banggakan untuk tawarkan hari ini.

Di bagian berikutnya, kami akan menyelami bagaimana kami mendekati proses desain dan pengembangan untuk mewujudkan fitur kompleks ini.

### Desain dan Pengembangan

Ketika kami memutuskan untuk menangani izin kustom, kami dengan cepat menyadari bahwa kami menghadapi tugas yang sangat besar.

Pada pandangan pertama, "izin kustom" mungkin terdengar sederhana, tetapi ini adalah fitur yang secara menipu kompleks yang menyentuh setiap aspek sistem kami.

Tantangannya sangat menakutkan: kami perlu menerapkan izin bertingkat, memungkinkan pengeditan secara langsung, melakukan perubahan signifikan pada skema basis data, dan memastikan fungsionalitas yang mulus di seluruh ekosistem kami – aplikasi web, Mac, Windows, iOS, dan Android, serta API dan webhook kami.

Kompleksitasnya cukup untuk membuat bahkan pengembang yang paling berpengalaman berhenti sejenak.

Pendekatan kami berpusat pada dua prinsip kunci:

1. Memecah fitur menjadi versi yang dapat dikelola
2. Mengadopsi pengiriman bertahap.

Menghadapi kompleksitas izin kustom skala penuh, kami bertanya pada diri sendiri pertanyaan penting:

> Apa versi pertama dari fitur ini yang paling sederhana mungkin?

Pendekatan ini sejalan dengan prinsip agile untuk memberikan Produk Minimum Layak (MVP) dan beriterasi berdasarkan umpan balik.

Jawaban kami sangat sederhana:

1. Memperkenalkan toggle untuk menyembunyikan tab aktivitas proyek
2. Menambahkan toggle lain untuk menyembunyikan tab formulir

**Itu saja.**

Tidak ada lonceng dan peluit, tidak ada matriks izin yang kompleks—hanya dua saklar sederhana on/off.

Meskipun mungkin tampak kurang mengesankan pada pandangan pertama, pendekatan ini menawarkan beberapa keuntungan signifikan:

* **Implementasi Cepat**: Saklar sederhana ini dapat dikembangkan dan diuji dengan cepat, memungkinkan kami untuk mendapatkan versi dasar izin kustom ke tangan pengguna dengan cepat.
* **Nilai Pengguna yang Jelas**: Bahkan dengan hanya dua opsi ini, kami memberikan nilai yang nyata. Beberapa tim mungkin ingin menyembunyikan umpan aktivitas dari klien, sementara yang lain mungkin perlu membatasi akses ke formulir untuk kelompok pengguna tertentu.
* **Fondasi untuk Pertumbuhan**: Awal yang sederhana ini meletakkan dasar untuk izin yang lebih kompleks. Ini memungkinkan kami untuk menyiapkan infrastruktur dasar untuk izin kustom tanpa terjebak dalam kompleksitas sejak awal.
* **Umpan Balik Pengguna**: Dengan merilis versi sederhana ini, kami dapat mengumpulkan umpan balik dunia nyata tentang bagaimana pengguna berinteraksi dengan izin kustom, yang menginformasikan pengembangan kami di masa depan.
* **Pembelajaran Teknis**: Implementasi awal ini memberikan tim pengembangan kami pengalaman praktis dalam memodifikasi izin di seluruh platform kami, mempersiapkan kami untuk iterasi yang lebih kompleks.

Dan Anda tahu, sebenarnya cukup merendahkan untuk memiliki visi besar untuk sesuatu, dan kemudian mengirimkan sesuatu yang merupakan persentase kecil dari visi itu.

Setelah mengirimkan dua saklar pertama ini, kami memutuskan untuk menangani sesuatu yang lebih canggih. Kami mendarat pada dua izin peran pengguna kustom baru.

Yang pertama adalah kemampuan untuk membatasi pengguna hanya untuk melihat catatan yang telah ditugaskan secara spesifik kepada mereka. Ini sangat berguna jika Anda memiliki klien dalam sebuah proyek dan Anda hanya ingin mereka melihat catatan yang secara spesifik ditugaskan kepada mereka alih-alih semua yang Anda kerjakan untuk mereka.

Yang kedua adalah opsi bagi administrator proyek untuk memblokir kelompok pengguna agar tidak dapat mengundang pengguna lain. Ini baik jika Anda memiliki proyek sensitif yang ingin Anda pastikan tetap pada basis "perlu dilihat".

Setelah kami mengirimkan ini, kami mendapatkan lebih banyak kepercayaan diri dan untuk versi ketiga kami menangani izin tingkat kolom, yang berarti dapat memutuskan kolom kustom mana yang dapat dilihat atau diedit oleh kelompok pengguna tertentu.

Ini sangat kuat. Bayangkan Anda memiliki proyek CRM, dan Anda memiliki data di sana yang tidak hanya terkait dengan jumlah yang akan dibayar pelanggan, tetapi juga biaya dan margin keuntungan Anda. Anda mungkin tidak ingin kolom biaya dan kolom rumus margin proyek Anda terlihat oleh staf junior, dan izin kustom memungkinkan Anda mengunci kolom tersebut sehingga tidak ditampilkan.

Selanjutnya, kami beralih ke membuat izin berbasis daftar, di mana administrator proyek dapat memutuskan apakah kelompok pengguna dapat melihat, mengedit, dan menghapus daftar tertentu. Jika mereka menyembunyikan daftar, semua catatan di dalam daftar tersebut juga menjadi tersembunyi, yang sangat baik karena berarti Anda dapat menyembunyikan bagian tertentu dari proses Anda dari anggota tim atau klien Anda.

Ini adalah hasil akhirnya:

<video autoplay loop muted playsinline>
  <source src="/videos/custom-user-roles.mp4" type="video/mp4">
</video>

## Pertimbangan Teknis

Di jantung arsitektur teknis Blue terletak GraphQL, pilihan penting yang telah secara signifikan memengaruhi kemampuan kami untuk menerapkan fitur kompleks seperti izin kustom. Tetapi sebelum kita menyelami spesifiknya, mari kita mundur sejenak dan memahami apa itu GraphQL dan bagaimana ia berbeda dari pendekatan REST API yang lebih tradisional.
GraphQL vs REST API: Penjelasan yang Mudah Dipahami

Bayangkan Anda berada di sebuah restoran. Dengan REST API, ini seperti memesan dari menu tetap. Anda meminta hidangan tertentu (endpoint), dan Anda mendapatkan semua yang menyertainya, apakah Anda ingin semuanya atau tidak. Jika Anda ingin menyesuaikan makanan Anda, Anda mungkin perlu melakukan beberapa pesanan (panggilan API) atau meminta hidangan yang disiapkan khusus (endpoint kustom).

GraphQL, di sisi lain, seperti berbicara dengan koki yang dapat menyiapkan apa pun. Anda memberi tahu koki persis bahan apa yang Anda inginkan (bidang data), dan dalam jumlah berapa. Koki kemudian menyiapkan hidangan yang persis seperti yang Anda minta - tidak lebih, tidak kurang. Inilah yang dilakukan GraphQL - memungkinkan klien meminta tepat data yang mereka butuhkan, dan server memberikan hanya itu.

### Makan Siang yang Penting

Sekitar enam minggu setelah pengembangan awal Blue, insinyur utama dan CEO kami keluar untuk makan siang.

Topik diskusinya?

Apakah akan beralih dari REST API ke GraphQL. Ini bukan keputusan yang bisa dianggap remeh - mengadopsi GraphQL berarti membuang enam minggu pekerjaan awal.

Dalam perjalanan kembali ke kantor, CEO mengajukan pertanyaan penting kepada insinyur utama: "Apakah kita akan menyesal tidak melakukannya lima tahun dari sekarang?"

Jawabannya menjadi jelas: GraphQL adalah jalan ke depan.

Kami mengenali potensi teknologi ini sejak awal, melihat bagaimana ia dapat mendukung visi kami untuk platform manajemen proyek yang fleksibel dan kuat.

Pandangan jauh ke depan kami dalam mengadopsi GraphQL membuahkan hasil ketika menerapkan izin kustom. Dengan REST API, kami akan membutuhkan endpoint yang berbeda untuk setiap konfigurasi izin kustom yang mungkin - pendekatan yang dengan cepat akan menjadi tidak praktis dan sulit untuk dipelihara.

GraphQL, bagaimanapun, memungkinkan kami menangani izin kustom secara dinamis. Inilah cara kerjanya:

- **Pemeriksaan Izin Secara Langsung**: Ketika klien membuat permintaan, server GraphQL kami dapat memeriksa izin pengguna langsung dari basis data kami.
- **Pengambilan Data yang Tepat**: Berdasarkan izin ini, GraphQL mengembalikan hanya data yang diminta yang sesuai dengan hak akses pengguna.
- **Kueri Fleksibel**: Saat izin berubah, kami tidak perlu membuat endpoint baru atau mengubah yang sudah ada. Kueri GraphQL yang sama dapat beradaptasi dengan pengaturan izin yang berbeda.
- **Pengambilan Data yang Efisien**: GraphQL memungkinkan klien meminta tepat apa yang mereka butuhkan. Ini berarti kami tidak mengambil data berlebihan, yang dapat mengekspos informasi yang seharusnya tidak diakses pengguna.

Fleksibilitas ini sangat penting untuk fitur yang kompleks seperti izin kustom. Ini memungkinkan kami menawarkan kontrol terperinci *tanpa* mengorbankan kinerja atau pemeliharaan.

## Tantangan

Menerapkan izin kustom di Blue membawa tantangannya sendiri, masing-masing mendorong kami untuk berinovasi dan menyempurnakan pendekatan kami. Optimasi kinerja dengan cepat muncul sebagai perhatian kritis. Saat kami menambahkan lebih banyak pemeriksaan izin terperinci, kami berisiko memperlambat sistem kami, terutama untuk proyek besar dengan banyak pengguna dan pengaturan izin yang kompleks. Untuk mengatasi ini, kami menerapkan strategi caching bertingkat, mengoptimalkan kueri basis data kami, dan memanfaatkan kemampuan GraphQL untuk meminta hanya data yang diperlukan. Pendekatan ini memungkinkan kami mempertahankan waktu respons yang cepat bahkan saat proyek berkembang dan kompleksitas izin meningkat.

Antarmuka pengguna untuk izin kustom menghadirkan hambatan signifikan lainnya. Kami perlu membuat antarmuka intuitif dan dapat dikelola untuk administrator, bahkan saat kami menambahkan lebih banyak opsi dan meningkatkan kompleksitas sistem.

Solusi kami melibatkan beberapa putaran pengujian pengguna dan desain iteratif.

Kami memperkenalkan matriks izin visual yang memungkinkan administrator dengan cepat melihat dan memodifikasi izin di berbagai peran dan area proyek.

Memastikan konsistensi lintas platform menghadirkan tantangan tersendiri. Kami perlu menerapkan izin kustom secara seragam di aplikasi web, desktop, dan mobile kami, masing-masing dengan antarmuka dan pertimbangan pengalaman pengguna yang unik. Ini sangat rumit untuk aplikasi mobile kami, yang harus menyembunyikan dan menampilkan fitur secara dinamis berdasarkan izin pengguna. Kami mengatasi ini dengan memusatkan logika izin kami di lapisan API, memastikan bahwa semua platform menerima data izin yang konsisten.

Kemudian, kami mengembangkan kerangka UI yang fleksibel yang dapat beradaptasi dengan perubahan izin ini secara real-time, memberikan pengalaman yang mulus terlepas dari platform yang digunakan.

Pendidikan dan adopsi pengguna menghadirkan hambatan terakhir dalam perjalanan izin kustom kami. Memperkenalkan fitur yang begitu kuat berarti kami perlu membantu pengguna kami memahami dan memanfaatkan izin kustom secara efektif.

Kami awalnya meluncurkan izin kustom kepada subset basis pengguna kami, dengan hati-hati memantau pengalaman mereka dan mengumpulkan wawasan. Pendekatan ini memungkinkan kami untuk menyempurnakan fitur dan materi pendidikan kami berdasarkan penggunaan dunia nyata sebelum meluncurkan kepada seluruh basis pengguna kami.

Peluncuran bertahap terbukti sangat berharga, membantu kami mengidentifikasi dan menangani masalah kecil serta titik kebingungan pengguna yang tidak kami duga, yang pada akhirnya mengarah pada fitur yang lebih halus dan ramah pengguna untuk semua pengguna kami.

Pendekatan ini meluncurkan kepada subset pengguna, serta periode "Beta" kami yang biasanya 2-3 minggu di Beta publik kami membantu kami tidur nyenyak di malam hari. :)

## Melihat ke Depan

Seperti semua fitur, tidak ada yang pernah *"selesai"*.

Visi jangka panjang kami untuk fitur izin kustom membentang di seluruh tag, filter kolom kustom, navigasi proyek yang dapat disesuaikan, dan kontrol komentar.

Mari kita selami setiap aspek.

### Izin Tag

Kami pikir akan sangat luar biasa jika dapat membuat izin berdasarkan apakah sebuah catatan memiliki satu atau lebih tag. Kasus penggunaan yang paling jelas adalah Anda membuat peran pengguna kustom yang disebut "Pelanggan" dan hanya mengizinkan pengguna dalam peran itu untuk melihat catatan yang memiliki tag "Pelanggan".

Ini memberi Anda pandangan sekilas tentang apakah catatan dapat atau tidak dapat dilihat oleh pelanggan Anda.

Ini bisa menjadi lebih kuat dengan kombinator AND/OR, di mana Anda dapat menentukan aturan yang lebih kompleks. Misalnya, Anda dapat mengatur aturan yang memungkinkan akses ke catatan yang diberi tag baik "Pelanggan" DAN "Publik", atau catatan yang diberi tag baik "Internal" ATAU "Kerahasiaan". Tingkat fleksibilitas ini akan memungkinkan pengaturan izin yang sangat nuansa, memenuhi bahkan struktur organisasi dan alur kerja yang paling kompleks.

Aplikasi potensialnya sangat luas. Manajer proyek dapat dengan mudah memisahkan informasi sensitif, tim penjualan dapat memiliki akses otomatis ke data klien yang relevan, dan kolaborator eksternal dapat diintegrasikan ke dalam bagian tertentu dari proyek tanpa risiko mengekspos informasi internal yang sensitif.

### Filter Kolom Kustom

Visi kami untuk Filter Kolom Kustom mewakili lompatan signifikan dalam kontrol akses terperinci. Fitur ini akan memberdayakan administrator proyek untuk mendefinisikan catatan mana yang dapat dilihat oleh kelompok pengguna tertentu berdasarkan nilai kolom kustom. Ini tentang menciptakan batasan yang dinamis dan berbasis data untuk akses informasi.

Bayangkan dapat mengatur izin seperti ini:

- Tampilkan hanya catatan di mana dropdown "Status Proyek" diatur ke "Publik"
- Batasi visibilitas untuk item di mana kolom multi-select "Departemen" mencakup "Pemasaran"
- Izinkan akses ke tugas di mana kotak centang "Prioritas" dicentang
- Tampilkan proyek di mana kolom angka "Anggaran" di atas ambang tertentu

### Navigasi Proyek yang Dapat Disesuaikan

Ini hanyalah perpanjangan dari saklar yang sudah kami miliki. Alih-alih hanya memiliki saklar untuk "aktivitas" dan "formulir", kami ingin memperluasnya ke setiap bagian dari navigasi proyek. Dengan cara ini, administrator proyek dapat membuat antarmuka yang terfokus dan menghapus alat yang tidak mereka butuhkan.

### Kontrol Komentar

Di masa depan, kami ingin kreatif dalam cara kami memungkinkan pelanggan kami memutuskan siapa yang dapat dan tidak dapat melihat komentar. Kami mungkin mengizinkan beberapa area komentar bertab di bawah satu catatan, dan masing-masing dapat terlihat atau tidak terlihat oleh kelompok pengguna yang berbeda.

Selain itu, kami mungkin juga mengizinkan fitur di mana hanya komentar di mana pengguna *secara spesifik* disebutkan yang terlihat, dan tidak ada yang lain. Ini akan memungkinkan tim yang memiliki klien di proyek untuk memastikan bahwa hanya komentar yang ingin mereka lihat yang terlihat oleh klien.

## Kesimpulan

Jadi, di sinilah kami, inilah cara kami mendekati pembangunan salah satu fitur yang paling menarik dan kuat! [Seperti yang Anda lihat di alat perbandingan manajemen proyek kami](/compare), sangat sedikit sistem manajemen proyek yang memiliki pengaturan matriks izin yang kuat seperti ini, dan yang memilikinya biasanya menyimpannya untuk rencana perusahaan mereka yang paling mahal, membuatnya tidak dapat diakses oleh perusahaan kecil atau menengah yang tipikal.

Dengan Blue, Anda memiliki *semua* fitur yang tersedia dengan rencana kami — kami tidak percaya bahwa fitur tingkat perusahaan harus diperuntukkan bagi pelanggan perusahaan!
---
title: AI Auto Categorization (Engineering Deep Dive)
category: "Engineering"
description: Ikuti tim engineering Blue saat mereka menjelaskan bagaimana mereka membangun fitur auto-categorization dan tagging bertenaga AI.
date: 2024-12-07
---

Kami baru saja meluncurkan [AI Auto Categorization](/insights/ai-auto-categorization) untuk semua pengguna Blue. Ini adalah fitur AI yang disertakan dalam langganan inti Blue, tanpa biaya tambahan. Dalam tulisan ini, kami akan menyelami engineering di balik pembuatan fitur ini.

---
Di Blue, pendekatan kami terhadap pengembangan fitur berakar pada pemahaman mendalam tentang kebutuhan pengguna dan tren pasar, dipadukan dengan komitmen untuk mempertahankan kesederhanaan dan kemudahan penggunaan yang menjadi ciri khas platform kami. Inilah yang mendorong [roadmap](/platform/roadmap) kami, dan yang [memungkinkan kami untuk secara konsisten meluncurkan fitur setiap bulan selama bertahun-tahun](/platform/changelog).

Pengenalan auto-tagging bertenaga AI ke Blue adalah contoh sempurna dari filosofi ini dalam praktik. Sebelum kami menyelami detail teknis tentang bagaimana kami membangun fitur ini, penting untuk memahami masalah yang kami pecahkan dan pertimbangan matang yang masuk ke dalam pengembangannya.

Lanskap manajemen proyek berkembang pesat, dengan kemampuan AI menjadi semakin sentral terhadap ekspektasi pengguna. Pelanggan kami, terutama mereka yang mengelola [proyek](/platform) skala besar dengan jutaan [records](/platform/features/records), telah vokal tentang keinginan mereka untuk cara yang lebih cerdas dan efisien dalam mengorganisir dan mengkategorikan data mereka.

Namun, di Blue, kami tidak sekadar menambahkan fitur karena sedang tren atau diminta. Filosofi kami adalah bahwa setiap penambahan baru harus membuktikan nilainya, dengan jawaban default yang tegas *"tidak"* sampai sebuah fitur menunjukkan permintaan yang kuat dan utilitas yang jelas.

Untuk benar-benar memahami kedalaman masalah dan potensi AI auto-tagging, kami melakukan wawancara pelanggan yang ekstensif, berfokus pada pengguna lama yang mengelola proyek kompleks dan kaya data di berbagai domain.

Percakapan ini mengungkap benang merah: *sementara tagging sangat berharga untuk organisasi dan kemampuan pencarian, sifat manual dari proses tersebut menjadi hambatan, terutama bagi tim yang berurusan dengan volume records yang tinggi.*

Tapi kami melihat lebih dari sekadar memecahkan masalah langsung dari tagging manual.

Kami membayangkan masa depan di mana tagging bertenaga AI dapat menjadi fondasi untuk alur kerja yang lebih cerdas dan otomatis.

Kekuatan sebenarnya dari fitur ini, kami sadari, terletak pada potensi integrasinya dengan [sistem otomasi manajemen proyek](/platform/features/automations) kami. Bayangkan sebuah alat manajemen proyek yang tidak hanya mengkategorikan informasi secara cerdas tetapi juga menggunakan kategori tersebut untuk mengarahkan tugas, memicu tindakan, dan menyesuaikan alur kerja secara real-time.

Visi ini selaras sempurna dengan tujuan kami untuk menjaga Blue tetap sederhana namun kuat.

Lebih lanjut, kami mengenali potensi untuk memperluas kemampuan ini di luar batas platform kami. Dengan mengembangkan sistem AI tagging yang kuat, kami sedang meletakkan dasar untuk "categorization API" yang dapat bekerja langsung, berpotensi membuka jalan baru tentang bagaimana pengguna kami berinteraksi dengan dan memanfaatkan Blue dalam ekosistem teknologi mereka yang lebih luas.

Oleh karena itu, fitur ini bukan hanya tentang menambahkan kotak centang AI ke daftar fitur kami.

Ini tentang mengambil langkah signifikan menuju platform manajemen proyek yang lebih cerdas dan adaptif sambil tetap setia pada filosofi inti kami tentang kesederhanaan dan fokus pada pengguna.

Dalam bagian-bagian berikut, kami akan menyelami tantangan teknis yang kami hadapi dalam mewujudkan visi ini, arsitektur yang kami rancang untuk mendukungnya, dan solusi yang kami implementasikan. Kami juga akan mengeksplorasi kemungkinan masa depan yang dibuka oleh fitur ini, menunjukkan bagaimana penambahan yang dipertimbangkan dengan cermat dapat membuka jalan bagi perubahan transformatif dalam manajemen proyek.

---
## Masalah

Seperti yang dibahas di atas, tagging manual untuk records proyek bisa memakan waktu dan tidak konsisten.

Kami berangkat untuk memecahkan ini dengan memanfaatkan AI untuk secara otomatis menyarankan tag berdasarkan konten record.

Tantangan utamanya adalah:

1. Memilih model AI yang tepat
2. Memproses volume records yang besar secara efisien
3. Memastikan privasi dan keamanan data
4. Mengintegrasikan fitur dengan mulus ke dalam arsitektur kami yang ada

## Memilih Model AI

Kami mengevaluasi beberapa platform AI, termasuk [OpenAI](https://openai.com), model open-source di [HuggingFace](https://huggingface.co/), dan [Replicate](https://replicate.com).

Kriteria kami termasuk:

- Efektivitas biaya
- Akurasi dalam memahami konteks
- Kemampuan untuk mematuhi format output tertentu
- Jaminan privasi data

Setelah pengujian menyeluruh, kami memilih [GPT-3.5 Turbo](https://platform.openai.com/docs/models/gpt-3-5-turbo) dari OpenAI. Meskipun [GPT-4](https://softgist.com/the-ultimate-guide-to-prompt-engineering) mungkin menawarkan peningkatan marginal dalam akurasi, tes kami menunjukkan bahwa kinerja GPT-3.5 lebih dari memadai untuk kebutuhan auto-tagging kami. Keseimbangan efektivitas biaya dan kemampuan kategorisasi yang kuat membuat GPT-3.5 pilihan ideal untuk fitur ini.

Biaya yang lebih tinggi dari GPT-4 akan memaksa kami untuk menawarkan fitur ini sebagai add-on berbayar, bertentangan dengan tujuan kami untuk **menggabungkan AI dalam produk utama kami tanpa biaya tambahan bagi pengguna akhir.**

Pada saat implementasi kami, harga untuk GPT-3.5 Turbo adalah:

- $0.0005 per 1K token input (atau $0.50 per 1M token input)
- $0.0015 per 1K token output (atau $1.50 per 1M token output)

Mari kita buat beberapa asumsi tentang record rata-rata di Blue:

- **Judul**: ~10 token
- **Deskripsi**: ~50 token
- **2 komentar**: ~30 token masing-masing
- **5 custom fields**: ~10 token masing-masing
- **Nama list, tanggal jatuh tempo, dan metadata lainnya**: ~20 token
- **System prompt dan tag yang tersedia**: ~50 token

Total token input per record: 10 + 50 + (30 * 2) + (10 * 5) + 20 + 50 ≈ 240 token

Untuk output, mari kita asumsikan rata-rata 3 tag yang disarankan per record, yang mungkin total sekitar 20 token output termasuk format JSON.

Untuk 1 juta records:

- Biaya input: (240 * 1.000.000 / 1.000.000) * $0.50 = $120
- Biaya output: (20 * 1.000.000 / 1.000.000) * $1.50 = $30

**Total biaya untuk auto-tagging 1 juta records: $120 + $30 = $150**

## Kinerja GPT3.5 Turbo

Kategorisasi adalah tugas yang sangat dikuasai oleh large language models (LLMs) seperti GPT-3.5 Turbo, membuat mereka sangat cocok untuk fitur auto-tagging kami. LLMs dilatih pada data teks dalam jumlah besar, memungkinkan mereka untuk memahami konteks, semantik, dan hubungan antara konsep. Basis pengetahuan yang luas ini memungkinkan mereka untuk melakukan tugas kategorisasi dengan akurasi tinggi di berbagai domain.

Untuk kasus penggunaan spesifik kami dalam tagging manajemen proyek, GPT-3.5 Turbo menunjukkan beberapa kekuatan utama:

- **Pemahaman Kontekstual:** Dapat memahami konteks keseluruhan dari record proyek, mempertimbangkan tidak hanya kata-kata individual tetapi makna yang disampaikan oleh seluruh deskripsi, komentar, dan field lainnya.
- **Fleksibilitas:** Dapat beradaptasi dengan berbagai jenis proyek dan industri tanpa memerlukan pemrograman ulang yang ekstensif.
- **Menangani Ambiguitas:** Dapat mempertimbangkan berbagai faktor untuk membuat keputusan yang bernuansa.
- **Belajar dari Contoh:** Dapat dengan cepat memahami dan menerapkan skema kategorisasi baru tanpa pelatihan tambahan.
- **Klasifikasi Multi-label:** Dapat menyarankan beberapa tag yang relevan untuk satu record, yang sangat penting untuk kebutuhan kami.

GPT-3.5 Turbo juga menonjol karena keandalannya dalam mematuhi format output JSON yang kami butuhkan, yang *sangat penting* untuk integrasi mulus dengan sistem kami yang ada. Model open-source, meskipun menjanjikan, sering menambahkan komentar tambahan atau menyimpang dari format yang diharapkan, yang akan memerlukan post-processing tambahan. Konsistensi dalam format output ini adalah faktor kunci dalam keputusan kami, karena secara signifikan menyederhanakan implementasi kami dan mengurangi titik kegagalan potensial.

Memilih GPT-3.5 Turbo dengan output JSON yang konsisten memungkinkan kami untuk mengimplementasikan solusi yang lebih sederhana, andal, dan dapat dipelihara.

Seandainya kami memilih model dengan format yang kurang andal, kami akan menghadapi rangkaian komplikasi: kebutuhan akan logika parsing yang kuat untuk menangani berbagai format output, penanganan error yang ekstensif untuk output yang tidak konsisten, dampak kinerja potensial dari pemrosesan tambahan, kompleksitas pengujian yang meningkat untuk mencakup semua variasi output, dan beban pemeliharaan jangka panjang yang lebih besar.

Error parsing dapat menyebabkan tagging yang salah, berdampak negatif pada pengalaman pengguna. Dengan menghindari jebakan ini, kami dapat memfokuskan upaya engineering kami pada aspek kritis seperti optimasi kinerja dan desain antarmuka pengguna, daripada bergulat dengan output AI yang tidak dapat diprediksi.

## Arsitektur Sistem

Fitur AI auto-tagging kami dibangun di atas arsitektur yang kuat dan dapat diskalakan yang dirancang untuk menangani volume permintaan yang tinggi secara efisien sambil memberikan pengalaman pengguna yang mulus. Seperti halnya semua sistem kami, kami telah merancang fitur ini untuk mendukung satu orde besaran lebih banyak lalu lintas daripada yang kami alami saat ini. Pendekatan ini, meskipun tampak terlalu direkayasa untuk kebutuhan saat ini, adalah praktik terbaik yang memungkinkan kami untuk menangani lonjakan penggunaan yang tiba-tiba dengan mulus dan memberi kami ruang yang cukup untuk pertumbuhan tanpa perombakan arsitektur besar. Jika tidak, kami harus merekayasa ulang semua sistem kami setiap 18 bulan — sesuatu yang telah kami pelajari dengan cara yang sulit di masa lalu!

Mari kita uraikan komponen dan alur sistem kami:

- **Interaksi Pengguna:** Proses dimulai ketika pengguna menekan tombol "Autotag" di antarmuka Blue. Tindakan ini memicu alur kerja auto-tagging.
- **Panggilan API Blue:** Tindakan pengguna diterjemahkan ke dalam panggilan API ke backend Blue kami. Endpoint API ini dirancang untuk menangani permintaan auto-tagging.
- **Manajemen Antrian:** Alih-alih memproses permintaan segera, yang dapat menyebabkan masalah kinerja di bawah beban tinggi, kami menambahkan permintaan tagging ke antrian. Kami menggunakan Redis untuk mekanisme antrian ini, yang memungkinkan kami untuk mengelola beban secara efektif dan memastikan skalabilitas sistem.
- **Layanan Background:** Kami mengimplementasikan layanan background yang terus memantau antrian untuk permintaan baru. Layanan ini bertanggung jawab untuk memproses permintaan yang mengantri.
- **Integrasi API OpenAI:** Layanan background menyiapkan data yang diperlukan dan membuat panggilan API ke model GPT-3.5 OpenAI. Di sinilah tagging bertenaga AI yang sebenarnya terjadi. Kami mengirim data proyek yang relevan dan menerima tag yang disarankan sebagai balasannya.
- **Pemrosesan Hasil:** Layanan background memproses hasil yang diterima dari OpenAI. Ini melibatkan parsing respons AI dan menyiapkan data untuk aplikasi ke proyek.
- **Aplikasi Tag:** Hasil yang diproses digunakan untuk menerapkan tag baru ke item yang relevan dalam proyek. Langkah ini memperbarui database kami dengan tag yang disarankan AI.
- **Refleksi Antarmuka Pengguna:** Akhirnya, tag baru muncul di tampilan proyek pengguna, menyelesaikan proses auto-tagging dari perspektif pengguna.

Arsitektur ini menawarkan beberapa manfaat utama yang meningkatkan kinerja sistem dan pengalaman pengguna. Dengan memanfaatkan antrian dan pemrosesan background, kami telah mencapai skalabilitas yang mengesankan, memungkinkan kami untuk menangani banyak permintaan secara bersamaan tanpa membebani sistem kami atau mencapai batas rate dari API OpenAI. Mengimplementasikan arsitektur ini memerlukan pertimbangan cermat dari berbagai faktor untuk memastikan kinerja dan keandalan yang optimal. Untuk manajemen antrian, kami memilih Redis, memanfaatkan kecepatan dan keandalannya dalam menangani antrian terdistribusi.

Pendekatan ini juga berkontribusi pada responsivitas keseluruhan fitur. Pengguna menerima umpan balik segera bahwa permintaan mereka sedang diproses, bahkan jika tagging yang sebenarnya memakan waktu, menciptakan rasa interaksi real-time. Toleransi kesalahan arsitektur adalah keuntungan penting lainnya. Jika bagian mana pun dari proses mengalami masalah, seperti gangguan API OpenAI sementara, kami dapat mencoba lagi dengan baik atau menangani kegagalan tanpa mempengaruhi seluruh sistem.

Ketangguhan ini, dikombinasikan dengan kemunculan tag secara real-time, meningkatkan pengalaman pengguna, memberikan kesan "keajaiban" AI yang sedang bekerja.

## Data & Prompts

Langkah penting dalam proses auto-tagging kami adalah menyiapkan data untuk dikirim ke model GPT-3.5. Langkah ini memerlukan pertimbangan cermat untuk menyeimbangkan penyediaan konteks yang cukup untuk tagging yang akurat sambil mempertahankan efisiensi dan melindungi privasi pengguna. Berikut adalah tampilan terperinci dari proses persiapan data kami.

Untuk setiap record, kami mengumpulkan informasi berikut:

- **Nama List**: Memberikan konteks tentang kategori atau fase proyek yang lebih luas.
- **Judul Record**: Sering berisi informasi kunci tentang tujuan atau konten record.
- **Custom Fields**: Kami menyertakan [custom fields](/platform/features/custom-fields) berbasis teks dan angka, yang sering berisi informasi khusus proyek yang penting.
- **Deskripsi**: Biasanya berisi informasi paling rinci tentang record.
- **Komentar**: Dapat memberikan konteks tambahan atau pembaruan yang mungkin relevan untuk tagging.
- **Tanggal Jatuh Tempo**: Informasi temporal yang mungkin mempengaruhi pemilihan tag.

Menariknya, kami tidak mengirim data tag yang ada ke GPT-3.5, dan kami melakukan ini untuk menghindari bias pada model.

Inti dari fitur auto-tagging kami terletak pada bagaimana kami berinteraksi dengan model GPT-3.5 dan memproses responsnya. Bagian dari pipeline kami ini memerlukan desain yang cermat untuk memastikan tagging yang akurat, konsisten, dan efisien.

Kami menggunakan system prompt yang dibuat dengan cermat untuk menginstruksikan AI tentang tugasnya. Berikut adalah rincian prompt kami dan alasan di balik setiap komponen:

```
You will be provided with record data, and your task is to choose the tags that are relevant to the record.
You can respond with an empty array if you are unsure.
Available tags: ${tags}.
Today: ${currentDate}.
Please respond in JSON using the following format:
{ "tags": ["tag-1", "tag-2"] }
```

- **Definisi Tugas:** Kami dengan jelas menyatakan tugas AI untuk memastikan respons yang terfokus.
- **Penanganan Ketidakpastian:** Kami secara eksplisit mengizinkan respons kosong, mencegah tagging yang dipaksakan ketika AI tidak yakin.
- **Tag yang Tersedia:** Kami menyediakan daftar tag yang valid (${tags}) untuk membatasi pilihan AI pada tag proyek yang ada.
- **Tanggal Saat Ini:** Menyertakan ${currentDate} membantu AI memahami konteks temporal, yang bisa menjadi penting untuk jenis proyek tertentu.
- **Format Respons:** Kami menentukan format JSON untuk parsing yang mudah dan pemeriksaan error.

Prompt ini adalah hasil dari pengujian dan iterasi yang ekstensif. Kami menemukan bahwa eksplisit tentang tugas, opsi yang tersedia, dan format output yang diinginkan secara signifikan meningkatkan akurasi dan konsistensi respons AI — kesederhanaan adalah kunci!

Daftar tag yang tersedia dihasilkan di sisi server dan divalidasi sebelum dimasukkan dalam prompt. Kami menerapkan batas karakter yang ketat pada nama tag untuk mencegah prompt yang terlalu besar.

Seperti yang disebutkan di atas, kami tidak memiliki masalah dengan GPT-3.5 Turbo dalam mendapatkan kembali respons JSON murni dalam format yang benar 100% dari waktu.

Jadi secara ringkas,

- Kami menggabungkan system prompt dengan data record yang disiapkan.
- Prompt gabungan ini dikirim ke model GPT-3.5 melalui API OpenAI.
- Kami menggunakan pengaturan temperature 0.3 untuk menyeimbangkan kreativitas dan konsistensi dalam respons AI.
- Panggilan API kami menyertakan parameter max_tokens untuk membatasi ukuran respons dan mengontrol biaya.

Setelah kami menerima respons AI, kami melalui beberapa langkah untuk memproses dan menerapkan tag yang disarankan:

* **Parsing JSON**: Kami mencoba untuk mem-parse respons sebagai JSON. Jika parsing gagal, kami mencatat error dan melewatkan tagging untuk record tersebut.
* **Validasi Schema**: Kami memvalidasi JSON yang di-parse terhadap schema yang kami harapkan (objek dengan array "tags"). Ini menangkap penyimpangan struktural apa pun dalam respons AI.
* **Validasi Tag**: Kami merujuk silang tag yang disarankan terhadap daftar tag proyek yang valid. Langkah ini menyaring tag apa pun yang tidak ada dalam proyek, yang dapat terjadi jika AI salah paham atau jika tag proyek berubah antara pembuatan prompt dan pemrosesan respons.
* **Deduplikasi**: Kami menghapus tag duplikat apa pun dari saran AI untuk menghindari tagging yang berlebihan.
* **Aplikasi**: Tag yang divalidasi dan dideduplikasi kemudian diterapkan ke record di database kami.
* **Logging dan Analytics**: Kami mencatat tag akhir yang diterapkan. Data ini berharga untuk memantau kinerja sistem dan meningkatkannya dari waktu ke waktu.

## Tantangan

Mengimplementasikan auto-tagging bertenaga AI di Blue menghadirkan beberapa tantangan unik, masing-masing memerlukan solusi inovatif untuk memastikan fitur yang kuat, efisien, dan ramah pengguna.

### Undo Operasi Bulk

Fitur AI Tagging dapat dilakukan baik pada record individual maupun secara bulk. Masalah dengan operasi bulk adalah jika pengguna tidak menyukai hasilnya, mereka harus secara manual melalui ribuan record dan membatalkan pekerjaan AI. Jelas, itu tidak dapat diterima.

Untuk mengatasi ini, kami mengimplementasikan sistem sesi tagging yang inovatif. Setiap operasi tagging bulk diberikan ID sesi unik, yang dikaitkan dengan semua tag yang diterapkan selama sesi tersebut. Ini memungkinkan kami untuk mengelola operasi undo secara efisien dengan hanya menghapus semua tag yang terkait dengan ID sesi tertentu. Kami juga menghapus jejak audit terkait, memastikan bahwa operasi yang dibatalkan tidak meninggalkan jejak dalam sistem. Pendekatan ini memberi pengguna kepercayaan diri untuk bereksperimen dengan AI tagging, mengetahui mereka dapat dengan mudah mengembalikan perubahan jika diperlukan.

### Privasi Data

Privasi data adalah tantangan kritis lain yang kami hadapi.

Pengguna kami mempercayai kami dengan data proyek mereka, dan sangat penting untuk memastikan informasi ini tidak disimpan atau digunakan untuk pelatihan model oleh OpenAI. Kami mengatasi ini di berbagai front.

Pertama, kami membentuk perjanjian dengan OpenAI yang secara eksplisit melarang penggunaan data kami untuk pelatihan model. Selain itu, OpenAI menghapus data setelah pemrosesan, memberikan lapisan perlindungan privasi tambahan.

Di sisi kami, kami mengambil tindakan pencegahan untuk mengecualikan informasi sensitif, seperti detail penerima tugas, dari data yang dikirim ke AI sehingga ini memastikan bahwa nama individu tertentu tidak dikirim ke pihak ketiga bersama dengan data lainnya.

Pendekatan komprehensif ini memungkinkan kami untuk memanfaatkan kemampuan AI sambil mempertahankan standar privasi dan keamanan data tertinggi.

### Rate Limits dan Menangkap Error

Salah satu perhatian utama kami adalah skalabilitas dan pembatasan rate. Panggilan API langsung ke OpenAI untuk setiap record akan tidak efisien dan dapat dengan cepat mencapai batas rate, terutama untuk proyek besar atau selama waktu penggunaan puncak. Untuk mengatasi ini, kami mengembangkan arsitektur layanan background yang memungkinkan kami untuk mem-batch permintaan dan mengimplementasikan sistem antrian kami sendiri. Pendekatan ini membantu kami mengelola frekuensi panggilan API dan memungkinkan pemrosesan volume records yang besar lebih efisien, memastikan kinerja yang lancar bahkan di bawah beban berat.

Sifat interaksi AI berarti kami juga harus bersiap untuk error sesekali atau output yang tidak terduga. Ada contoh di mana AI mungkin menghasilkan JSON yang tidak valid atau output yang tidak cocok dengan format yang kami harapkan. Untuk menangani ini, kami mengimplementasikan penanganan error yang kuat dan logika parsing di seluruh sistem kami. Jika respons AI bukan JSON yang valid atau tidak berisi kunci "tags" yang diharapkan, sistem kami dirancang untuk memperlakukannya seolah-olah tidak ada tag yang disarankan, daripada mencoba memproses data yang berpotensi rusak. Ini memastikan bahwa bahkan dalam menghadapi ketidakpastian AI, sistem kami tetap stabil dan andal.

## Pengembangan Masa Depan

Kami percaya bahwa fitur, dan produk Blue secara keseluruhan, tidak pernah "selesai" — selalu ada ruang untuk perbaikan.

Ada beberapa fitur yang kami pertimbangkan dalam build awal yang tidak lolos fase scoping, tetapi menarik untuk dicatat karena kami kemungkinan akan mengimplementasikan beberapa versi dari mereka di masa depan.

Yang pertama adalah menambahkan deskripsi tag. Ini akan memungkinkan pengguna akhir untuk tidak hanya memberi tag nama dan warna, tetapi juga deskripsi opsional. Ini juga akan diteruskan ke AI untuk membantu memberikan konteks lebih lanjut dan berpotensi meningkatkan akurasi.

Meskipun konteks tambahan bisa berharga, kami sadar akan kompleksitas potensial yang mungkin diperkenalkannya. Ada keseimbangan halus yang harus dicapai antara menyediakan informasi yang berguna dan membuat pengguna kewalahan dengan terlalu banyak detail. Saat kami mengembangkan fitur ini, kami akan fokus pada menemukan titik sweet spot di mana konteks tambahan meningkatkan daripada memperumit pengalaman pengguna.

Mungkin peningkatan paling menarik di horizon kami adalah integrasi AI auto-tagging dengan [sistem otomasi manajemen proyek](/platform/features/automations) kami.

Ini berarti bahwa fitur AI tagging bisa menjadi trigger, atau action dari automation. Ini bisa sangat besar karena bisa mengubah fitur kategorisasi AI yang cukup sederhana ini menjadi sistem routing berbasis AI untuk pekerjaan.

Bayangkan automation yang menyatakan:

Ketika AI menandai record sebagai "Critical" -> Ditugaskan ke "Manager" dan Kirim Email Kustom

Ini berarti bahwa ketika Anda AI-tag record, jika AI memutuskan itu adalah masalah kritis, maka dapat secara otomatis menugaskan manajer proyek dan mengirim mereka email kustom. Ini memperluas [manfaat dari sistem otomasi manajemen proyek](/platform/features/automations) kami dari sistem berbasis aturan murni menjadi sistem AI yang benar-benar fleksibel.

Dengan terus mengeksplorasi batas AI dalam manajemen proyek, kami bertujuan untuk menyediakan alat kepada pengguna kami yang tidak hanya memenuhi kebutuhan mereka saat ini tetapi mengantisipasi dan membentuk masa depan pekerjaan. Seperti biasa, kami akan mengembangkan fitur-fitur ini dalam kolaborasi erat dengan komunitas pengguna kami, memastikan bahwa setiap peningkatan menambah nilai praktis yang nyata pada proses manajemen proyek.

## Kesimpulan

Jadi itu dia!

Ini adalah fitur yang menyenangkan untuk diimplementasikan, dan salah satu langkah pertama kami ke AI, bersama dengan [AI Content Summarization](/insights/ai-content-summarization) yang telah kami luncurkan sebelumnya. Kami tahu bahwa AI akan memainkan peran yang semakin besar dalam manajemen proyek di masa depan, dan kami tidak sabar untuk meluncurkan lebih banyak fitur inovatif yang memanfaatkan LLMs (Large Language Models) canggih.

Ada cukup banyak hal yang harus dipikirkan saat mengimplementasikan ini, dan kami sangat bersemangat tentang bagaimana kami dapat memanfaatkan fitur ini di masa depan dengan [mesin otomasi manajemen proyek](/insights/benefits-project-management-automation) Blue yang ada.

Kami juga berharap ini telah menjadi bacaan yang menarik, dan memberi Anda sekilas tentang bagaimana kami memikirkan engineering fitur yang Anda gunakan setiap hari.

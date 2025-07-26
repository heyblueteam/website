---
title: Mengapa Kami Membangun Chatbot Dokumentasi AI Kami Sendiri
description: Kami membangun chatbot AI dokumentasi kami sendiri yang dilatih pada dokumentasi platform Blue.
category: "Product Updates"
date: 2024-07-09
---


Di Blue, kami selalu mencari cara untuk mempermudah hidup pelanggan kami. Kami memiliki [dokumentasi mendalam tentang setiap fitur](https://documentation.blue.cc), [Video YouTube](https://www.youtube.com/@HeyBlueTeam), [Tips & Trik](/insights/tips-tricks), dan [berbagai saluran dukungan](/support).

Kami telah mengawasi perkembangan AI (Kecerdasan Buatan) dengan seksama karena kami sangat tertarik pada [automasi manajemen proyek](/platform/features/automations). Kami juga merilis fitur seperti [Kategorisasi Otomatis AI](/insights/ai-auto-categorization) dan [Ringkasan AI](/insights/ai-content-summarization) untuk mempermudah pekerjaan ribuan pelanggan kami.

Satu hal yang jelas adalah bahwa AI ada untuk tinggal, dan akan memiliki efek luar biasa di sebagian besar industri, dan manajemen proyek bukanlah pengecualian. Jadi kami bertanya pada diri sendiri bagaimana kami bisa lebih memanfaatkan AI untuk membantu seluruh siklus hidup pelanggan, mulai dari penemuan, pra-penjualan, onboarding, dan juga dengan pertanyaan yang berkelanjutan.

Jawabannya cukup jelas: **Kami membutuhkan chatbot AI yang dilatih pada dokumentasi kami.**

Mari kita hadapi kenyataan: *setiap* organisasi seharusnya memiliki chatbot. Mereka adalah cara yang bagus bagi pelanggan untuk mendapatkan jawaban instan untuk pertanyaan umum, tanpa harus menggali halaman-halaman dokumentasi yang padat atau situs web Anda. Pentingnya chatbot di situs web pemasaran modern tidak bisa diabaikan.

![](/insights/ai-chatbot-regular.png)

Untuk perusahaan perangkat lunak secara khusus, seseorang tidak boleh menganggap situs web pemasaran sebagai "hal" terpisah — itu *adalah* bagian dari produk Anda. Ini karena itu sesuai dengan kehidupan pelanggan yang khas:

- **Kesadaran** (Penemuan): Ini adalah saat pelanggan potensial pertama kali menemukan produk luar biasa Anda. Chatbot Anda bisa menjadi pemandu ramah mereka, mengarahkan mereka ke fitur dan manfaat kunci sejak awal.
- **Pertimbangan** (Pendidikan): Sekarang mereka penasaran dan ingin belajar lebih banyak. Chatbot Anda menjadi tutor pribadi mereka, memberikan informasi yang disesuaikan dengan kebutuhan dan pertanyaan spesifik mereka.
- **Pembelian/konversi**: Ini adalah momen kebenaran - ketika seorang prospek memutuskan untuk melangkah dan menjadi pelanggan. Chatbot Anda dapat menghaluskan masalah menit terakhir, menjawab pertanyaan "sebelum saya membeli", dan mungkin bahkan memberikan penawaran menarik untuk menyegel kesepakatan.
- **Onboarding**: Mereka sudah membeli, sekarang apa? Chatbot Anda bertransformasi menjadi asisten yang membantu, membimbing pengguna baru melalui pengaturan, menunjukkan cara menggunakan produk, dan memastikan mereka tidak merasa tersesat di dunia produk Anda.
- **Retensi**: Menjaga pelanggan tetap bahagia adalah nama permainannya. Chatbot Anda siap sedia 24/7, siap untuk memecahkan masalah, menawarkan tips dan trik, dan memastikan pelanggan Anda merasa diperhatikan.
- **Ekspansi**: Saatnya untuk naik level! Chatbot Anda dapat secara halus menyarankan fitur baru, upsell, atau cross-sell yang sesuai dengan cara pelanggan sudah menggunakan produk Anda. Ini seperti memiliki tenaga penjual yang sangat pintar, tidak memaksa, selalu siap sedia.
- **Advokasi**: Pelanggan yang bahagia menjadi pendukung terbesar Anda. Chatbot Anda dapat mendorong pengguna yang puas untuk menyebarkan berita, meninggalkan ulasan, atau berpartisipasi dalam program rujukan. Ini seperti memiliki mesin promosi yang terintegrasi langsung ke dalam produk Anda!

## Keputusan Membangun vs Membeli

Setelah kami memutuskan untuk menerapkan chatbot AI, pertanyaan besar berikutnya adalah: membangun atau membeli? Sebagai tim kecil yang fokus pada produk inti kami, kami umumnya lebih memilih solusi "sebagai layanan" atau platform open-source populer. Kami tidak berada dalam bisnis menciptakan kembali roda untuk setiap bagian dari tumpukan teknologi kami, setelah semua.
Jadi, kami menggulung lengan baju dan menyelam ke pasar, berburu solusi chatbot AI berbayar dan open-source.

Persyaratan kami sederhana, tetapi tidak dapat dinegosiasikan:

- **Pengalaman Tanpa Merek**: Chatbot ini bukan hanya widget yang bagus untuk dimiliki; ini akan ada di situs web pemasaran kami dan akhirnya di produk kami. Kami tidak tertarik untuk mengiklankan merek orang lain di real estate digital kami sendiri.
- **UX yang Hebat**: Bagi banyak pelanggan potensial, chatbot ini mungkin menjadi titik kontak pertama mereka dengan Blue. Ini menetapkan nada untuk persepsi mereka tentang perusahaan kami. Mari kita hadapi kenyataan: jika kami tidak bisa membuat chatbot yang tepat di situs web kami, bagaimana kami bisa mengharapkan pelanggan mempercayai kami dengan proyek dan proses kritis mereka?
- **Biaya yang Wajar**: Dengan basis pengguna yang besar dan rencana untuk mengintegrasikan chatbot ke dalam produk inti kami, kami membutuhkan solusi yang tidak akan menguras anggaran saat penggunaan meningkat. Idealnya, kami menginginkan opsi **BYOK (Bring Your Own Key)**. Ini akan memungkinkan kami untuk menggunakan kunci layanan OpenAI atau AI lainnya sendiri, membayar biaya variabel langsung alih-alih markup kepada vendor pihak ketiga yang sebenarnya tidak menjalankan model.
- **Kompatibel dengan API Asisten OpenAI**: Jika kami akan menggunakan perangkat lunak open-source, kami tidak ingin repot mengelola pipeline untuk pengambilan dokumen, pengindeksan, basis data vektor, dan semua itu. Kami ingin menggunakan [API Asisten OpenAI](https://platform.openai.com/docs/assistants/overview) yang akan mengabstraksi semua kompleksitas di balik API. Sejujurnya — ini sangat baik dilakukan.
- **Skalabilitas**: Kami ingin memiliki chatbot ini di beberapa tempat, dengan potensi puluhan ribu pengguna per tahun. Kami mengharapkan penggunaan yang signifikan, dan kami tidak ingin terjebak dalam solusi yang tidak dapat berkembang sesuai kebutuhan kami.

## Chatbot AI Komersial

Yang kami tinjau cenderung memiliki UX yang lebih baik daripada solusi open-source — seperti yang sering terjadi. Mungkin ada diskusi terpisah yang perlu dilakukan suatu hari tentang *mengapa* banyak solusi open-source mengabaikan atau meremehkan pentingnya UX.

Kami akan memberikan daftar di sini, jika Anda mencari beberapa penawaran komersial yang solid:

- **[Chatbase](https://chatbase.co):** Chatbase memungkinkan Anda membangun chatbot AI kustom yang dilatih pada basis pengetahuan Anda dan menambahkannya ke situs web Anda atau berinteraksi dengannya melalui API mereka. Ini menawarkan fitur seperti jawaban yang dapat dipercaya, generasi prospek, analitik lanjutan, dan kemampuan untuk terhubung ke berbagai sumber data. Bagi kami, ini terasa seperti salah satu penawaran komersial yang paling halus di luar sana.
- **[DocsBot AI](https://docsbot.ai/):** DocsBot AI membuat bot ChatGPT kustom yang dilatih pada dokumentasi dan konten Anda untuk dukungan, pra-penjualan, penelitian, dan lainnya. Ini menyediakan widget yang dapat disematkan untuk menambahkan chatbot ke situs web Anda dengan mudah, kemampuan untuk secara otomatis membalas tiket dukungan, dan API yang kuat untuk integrasi.
- **[CustomGPT.ai](https://customgpt.ai):** CustomGPT.ai menciptakan pengalaman chatbot pribadi dengan mengambil data bisnis Anda, termasuk konten situs web, helpdesk, basis pengetahuan, dokumen, dan lainnya. Ini memungkinkan prospek untuk mengajukan pertanyaan dan mendapatkan jawaban instan berdasarkan konten Anda, tanpa perlu mencari. Menariknya, mereka juga [mengklaim mengalahkan OpenAI di RAG (Retrieval Augmented Generation)!](https://customgpt.ai/customgpt-beats-open-ai-in-rag-benchmark/)
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Ini adalah penawaran komersial yang menarik, karena ini *juga* kebetulan merupakan perangkat lunak open-source. Sepertinya masih dalam tahap awal, dan harga tidak terasa realistis ($27/bulan untuk pesan tanpa batas tidak akan pernah berhasil secara komersial untuk mereka).

Kami juga melihat [InterCom Fin](https://www.intercom.com/fin) yang merupakan bagian dari perangkat lunak dukungan pelanggan mereka. Ini akan berarti beralih dari [HelpScout](https://wwww.helpscout.com) yang telah kami gunakan sejak kami memulai Blue. Ini mungkin bisa dilakukan, tetapi InterCom Fin memiliki harga yang gila yang membuatnya tidak dipertimbangkan.

Dan ini sebenarnya adalah masalah dengan banyak penawaran komersial. InterCom Fin mengenakan biaya $0,99 per permintaan dukungan pelanggan yang ditangani, dan ChatBase mengenakan biaya $399/bulan untuk 40.000 pesan. Itu hampir $5k setahun untuk widget chat sederhana.

Mengingat bahwa harga untuk inferensi AI turun dengan cepat. OpenAI mengurangi harga mereka secara dramatis:

- GPT-4 asli (8k konteks) dihargai $0,03 per 1K token prompt.
- GPT-4 Turbo (128k konteks) dihargai $0,01 per 1K token prompt, pengurangan 50% dari GPT-4 asli.
- Model GPT-4o dihargai $0,005 per 1K token, yang merupakan pengurangan lebih lanjut 50% dari harga GPT-4 Turbo.

Itu adalah pengurangan biaya sebesar 83%, dan kami tidak mengharapkan itu akan tetap stagnan.

Mengingat bahwa kami mencari solusi yang dapat diskalakan yang akan digunakan oleh puluhan ribu pengguna per tahun dengan jumlah pesan yang signifikan, masuk akal untuk langsung pergi ke sumbernya dan membayar biaya API secara langsung, tidak menggunakan versi komersial yang menandai biaya.

## Chatbot AI Open Source

Seperti yang disebutkan, opsi open source yang kami tinjau sebagian besar mengecewakan terkait dengan persyaratan "UX yang Hebat".

Kami melihat:

- **[Deepchat](https://deepchat.dev/)**: Ini adalah komponen chat yang tidak terikat pada framework untuk layanan AI, yang terhubung ke berbagai API AI termasuk OpenAI. Ini juga memiliki kemampuan bagi pengguna untuk mengunduh model AI yang berjalan langsung di browser. Kami mencoba ini dan mendapatkan versi yang berfungsi, tetapi API Asisten OpenAI yang diimplementasikan terasa cukup bermasalah dengan beberapa masalah. Namun, ini adalah proyek yang sangat menjanjikan, dan playground mereka sangat baik.
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Melihat ini lagi dari perspektif open-source, ini akan membutuhkan kami untuk membangun cukup banyak infrastruktur, sesuatu yang tidak ingin kami lakukan, karena kami ingin bergantung sebanyak mungkin pada API Asisten OpenAI.

## Membangun ChatBot Kami Sendiri

Dan jadi, tanpa dapat menemukan sesuatu yang sesuai dengan semua persyaratan kami, kami memutuskan untuk membangun chatbot AI kami sendiri yang dapat berinteraksi dengan API Asisten OpenAI. Ini, pada akhirnya, ternyata relatif tanpa rasa sakit!

Situs web kami menggunakan [Nuxt3](https://nuxt.com), [Vue3](https://vuejs.org/) (yang merupakan framework yang sama dengan Platform Blue), dan [TailwindUI](https://tailwindui.com/).

Langkah pertama adalah membuat API (Antarmuka Pemrograman Aplikasi) di Nuxt3 yang dapat "berbicara" dengan API Asisten OpenAI. Ini diperlukan karena kami tidak ingin melakukan semuanya di front-end, karena ini akan mengekspos kunci API OpenAI kami ke dunia, dengan potensi penyalahgunaan.

API backend kami bertindak sebagai perantara yang aman antara browser pengguna dan OpenAI. Berikut adalah apa yang dilakukannya:

- **Manajemen Percakapan:** Ini membuat dan mengelola "thread" untuk setiap percakapan. Anggap saja thread sebagai sesi chat unik yang mengingat semua yang telah Anda katakan.
- **Penanganan Pesan:** Ketika Anda mengirim pesan, API kami menambahkannya ke thread yang tepat dan meminta asisten OpenAI untuk menyusun respons.
- **Menunggu Cerdas:** Alih-alih membuat Anda menatap layar loading, API kami memeriksa OpenAI setiap detik untuk melihat apakah respons Anda sudah siap. Ini seperti memiliki pelayan yang mengawasi pesanan Anda tanpa mengganggu koki setiap dua detik.
- **Keamanan Pertama:** Dengan menangani semua ini di server, kami menjaga data Anda dan kunci API kami aman dan terjamin.

Kemudian, ada front-end dan pengalaman pengguna. Seperti yang dibahas sebelumnya, ini sangat *penting*, karena kami tidak mendapatkan kesempatan kedua untuk membuat kesan pertama!

Dalam merancang chatbot kami, kami memperhatikan pengalaman pengguna dengan cermat, memastikan bahwa setiap interaksi berjalan lancar, intuitif, dan mencerminkan komitmen Blue terhadap kualitas. Antarmuka chatbot dimulai dengan lingkaran Blue yang sederhana dan elegan, menggunakan [HeroIcons untuk ikon kami](https://heroicons.com/) (yang kami gunakan di seluruh situs web Blue) untuk bertindak sebagai widget pembuka chatbot kami. Pilihan desain ini memastikan konsistensi visual dan pengenalan merek yang segera.

![](/insights/ai-chatbot-circle.png)

Kami memahami bahwa terkadang pengguna mungkin memerlukan dukungan tambahan atau informasi yang lebih mendalam. Itulah sebabnya kami telah menyertakan tautan yang nyaman dalam antarmuka chatbot. Tautan email untuk dukungan tersedia dengan mudah, memungkinkan pengguna untuk menghubungi tim kami langsung jika mereka memerlukan bantuan yang lebih personal. Selain itu, kami telah mengintegrasikan tautan dokumentasi, memberikan akses mudah ke sumber daya yang lebih komprehensif bagi mereka yang ingin menyelami lebih dalam tentang penawaran Blue.

Pengalaman pengguna semakin ditingkatkan dengan animasi fade-in dan fade-up yang menarik saat membuka jendela chatbot. Animasi halus ini menambahkan sentuhan keanggunan pada antarmuka, membuat interaksi terasa lebih dinamis dan menarik. Kami juga telah menerapkan indikator pengetikan, fitur kecil tetapi penting yang memberi tahu pengguna bahwa chatbot sedang memproses pertanyaan mereka dan menyusun respons. Petunjuk visual ini membantu mengelola ekspektasi pengguna dan mempertahankan rasa komunikasi yang aktif.

<video autoplay loop muted playsinline>
  <source src="/videos/ai-chatbot-animation.mp4" type="video/mp4">
</video>

Menyadari bahwa beberapa percakapan mungkin memerlukan lebih banyak ruang layar, kami telah menambahkan kemampuan untuk membuka percakapan dalam jendela yang lebih besar. Fitur ini sangat berguna untuk pertukaran yang lebih panjang atau saat meninjau informasi terperinci, memberikan pengguna fleksibilitas untuk menyesuaikan chatbot dengan kebutuhan mereka.

Di balik layar, kami telah menerapkan beberapa pemrosesan cerdas untuk mengoptimalkan respons chatbot. Sistem kami secara otomatis mem-parsing balasan AI untuk menghapus referensi ke dokumen internal kami, memastikan bahwa informasi yang disajikan bersih, relevan, dan fokus hanya pada menjawab pertanyaan pengguna.
Untuk meningkatkan keterbacaan dan memungkinkan komunikasi yang lebih nuansa, kami telah mengintegrasikan dukungan markdown menggunakan pustaka 'marked'. Fitur ini memungkinkan AI kami untuk memberikan teks yang diformat dengan kaya, termasuk penekanan tebal dan miring, daftar terstruktur, dan bahkan potongan kode jika diperlukan. Ini seperti menerima mini-dokumen yang diformat dengan baik dan disesuaikan sebagai respons terhadap pertanyaan Anda.

Terakhir tetapi tidak kalah penting, kami telah memprioritaskan keamanan dalam implementasi kami. Menggunakan pustaka DOMPurify, kami menyaring HTML yang dihasilkan dari parsing markdown. Langkah penting ini memastikan bahwa skrip atau kode yang berpotensi berbahaya dihapus sebelum konten ditampilkan kepada Anda. Ini adalah cara kami menjamin bahwa informasi bermanfaat yang Anda terima tidak hanya informatif tetapi juga aman untuk dikonsumsi.

## Perkembangan Masa Depan

Jadi ini baru permulaan, kami memiliki beberapa hal menarik dalam roadmap untuk fitur ini.

Salah satu fitur mendatang kami adalah kemampuan untuk melakukan streaming respons secara real-time. Segera, Anda akan melihat balasan chatbot muncul karakter demi karakter, membuat percakapan terasa lebih alami dan dinamis. Ini seperti melihat AI berpikir, menciptakan pengalaman yang lebih menarik dan interaktif yang membuat Anda tetap terlibat di setiap langkah.

Untuk pengguna Blue yang kami hargai, kami sedang mengerjakan personalisasi. Chatbot akan mengenali ketika Anda masuk, menyesuaikan responsnya berdasarkan informasi akun Anda, riwayat penggunaan, dan preferensi. Bayangkan chatbot yang tidak hanya menjawab pertanyaan Anda tetapi juga memahami konteks spesifik Anda dalam ekosistem Blue, memberikan bantuan yang lebih relevan dan personal.

Kami memahami bahwa Anda mungkin sedang mengerjakan beberapa proyek atau memiliki berbagai pertanyaan. Itulah sebabnya kami sedang mengembangkan kemampuan untuk mempertahankan beberapa thread percakapan yang berbeda dengan chatbot kami. Fitur ini akan memungkinkan Anda untuk beralih antara topik yang berbeda dengan lancar, tanpa kehilangan konteks – seperti memiliki beberapa tab terbuka di browser Anda.

Untuk membuat interaksi Anda bahkan lebih produktif, kami sedang menciptakan fitur yang akan menawarkan pertanyaan tindak lanjut yang disarankan berdasarkan percakapan Anda saat ini. Ini akan membantu Anda menjelajahi topik lebih dalam dan menemukan informasi terkait yang mungkin tidak Anda pikirkan untuk ditanyakan, membuat setiap sesi chat lebih komprehensif dan berharga.

Kami juga bersemangat untuk menciptakan rangkaian asisten AI khusus, masing-masing disesuaikan untuk kebutuhan tertentu. Apakah Anda ingin menjawab pertanyaan pra-penjualan, mengatur proyek baru, atau memecahkan masalah fitur lanjutan, Anda akan dapat memilih asisten yang paling sesuai dengan kebutuhan Anda saat ini. Ini seperti memiliki tim ahli Blue di ujung jari Anda, masing-masing mengkhususkan diri dalam berbagai aspek platform kami.

Terakhir, kami sedang mengerjakan kemampuan untuk memungkinkan Anda mengunggah tangkapan layar langsung ke chat. AI akan menganalisis gambar dan memberikan penjelasan atau langkah pemecahan masalah berdasarkan apa yang dilihatnya. Fitur ini akan memudahkan Anda mendapatkan bantuan dengan masalah spesifik yang Anda temui saat menggunakan Blue, menjembatani kesenjangan antara informasi visual dan bantuan tekstual.

## Kesimpulan

Kami berharap penjelasan mendalam tentang proses pengembangan chatbot AI kami telah memberikan wawasan berharga tentang pemikiran pengembangan produk kami di Blue. Perjalanan kami dari mengidentifikasi kebutuhan akan chatbot hingga membangun solusi kami sendiri menunjukkan bagaimana kami mendekati pengambilan keputusan dan inovasi.

![](/insights/ai-chatbot-modal.png)

Di Blue, kami dengan hati-hati mempertimbangkan opsi membangun versus membeli, selalu dengan memperhatikan apa yang akan paling melayani pengguna kami dan sejalan dengan tujuan jangka panjang kami. Dalam hal ini, kami mengidentifikasi celah signifikan di pasar untuk chatbot yang hemat biaya namun menarik secara visual yang dapat memenuhi kebutuhan spesifik kami. Meskipun kami umumnya menganjurkan untuk memanfaatkan solusi yang ada daripada menciptakan kembali roda, terkadang jalan terbaik ke depan adalah menciptakan sesuatu yang disesuaikan dengan kebutuhan unik Anda.

Keputusan kami untuk membangun chatbot kami sendiri tidak diambil dengan ringan. Itu adalah hasil dari penelitian pasar yang menyeluruh, pemahaman yang jelas tentang kebutuhan kami, dan komitmen untuk memberikan pengalaman terbaik bagi pengguna kami. Dengan mengembangkan secara internal, kami dapat menciptakan solusi yang tidak hanya memenuhi kebutuhan kami saat ini tetapi juga meletakkan dasar untuk peningkatan dan integrasi di masa depan.

Proyek ini mencerminkan pendekatan kami di Blue: kami tidak takut untuk menggulung lengan baju dan membangun sesuatu dari awal ketika itu adalah pilihan yang tepat untuk produk dan pengguna kami. Kesiapan ini untuk melakukan lebih banyak usaha memungkinkan kami untuk memberikan solusi inovatif yang benar-benar memenuhi kebutuhan pelanggan kami.
Kami bersemangat tentang masa depan chatbot AI kami dan nilai yang akan dibawanya bagi pengguna Blue yang potensial dan yang sudah ada. Saat kami terus menyempurnakan dan memperluas kemampuannya, kami tetap berkomitmen untuk mendorong batasan apa yang mungkin dalam manajemen proyek dan interaksi pelanggan.

Terima kasih telah bergabung dengan kami dalam perjalanan ini melalui proses pengembangan kami. Kami berharap ini memberi Anda gambaran tentang pendekatan yang penuh perhatian dan berfokus pada pengguna yang kami ambil dengan setiap aspek Blue. Nantikan pembaruan lebih lanjut saat kami terus mengembangkan dan meningkatkan platform kami untuk melayani Anda dengan lebih baik.

Jika Anda tertarik, Anda dapat menemukan tautan ke kode sumber untuk proyek ini di sini:

- **[ChatWidget](https://gitlab.com/bloohq/blue-website/-/blob/main/components/ChatWidget.vue)**: Ini adalah Komponen Vue yang menggerakkan widget chat itu sendiri.
- **[Chat API](https://gitlab.com/bloohq/blue-website/-/blob/main/server/api/chat.post.ts)**: Ini adalah middleware yang bekerja di antara komponen chat dan API Asisten OpenAI.
---
title: FAQ dotyczące bezpieczeństwa Blue
description: To jest lista najczęściej zadawanych pytań dotyczących protokołów i praktyk bezpieczeństwa w Blue.
category: "FAQ"
date: 2024-07-19
---



Naszą misją jest zorganizowanie pracy na świecie poprzez stworzenie najlepszej platformy do zarządzania projektami na planecie.

Centralnym elementem osiągnięcia tej misji jest zapewnienie, że nasza platforma jest bezpieczna, niezawodna i godna zaufania. Rozumiemy, że aby być twoim jedynym źródłem prawdy, Blue musi chronić twoje wrażliwe dane biznesowe przed zagrożeniami zewnętrznymi, utratą danych i przestojami.

Oznacza to, że traktujemy bezpieczeństwo poważnie w Blue.

Kiedy myślimy o bezpieczeństwie, rozważamy holistyczne podejście, które koncentruje się na trzech kluczowych obszarach:

1.  **Bezpieczeństwo infrastruktury i sieci**: Zapewnia, że nasze systemy fizyczne i wirtualne są chronione przed zagrożeniami zewnętrznymi i nieautoryzowanym dostępem.
2.  **Bezpieczeństwo oprogramowania**: Skupia się na bezpieczeństwie samego kodu, w tym na praktykach bezpiecznego kodowania, regularnych przeglądach kodu i zarządzaniu podatnościami.
3.  **Bezpieczeństwo platformy**: Obejmuje funkcje w Blue, takie jak [zaawansowane kontrole dostępu](/platform/features/user-permissions), zapewniając, że projekty są domyślnie prywatne, oraz inne środki mające na celu ochronę danych użytkowników i prywatności.


## Jak skalowalny jest Blue?

To ważne pytanie, ponieważ chcesz mieć system, który może *rosnąć* razem z tobą. Nie chcesz musieć zmieniać swojej platformy do zarządzania projektami i procesami za sześć lub dwanaście miesięcy.

Starannie wybieramy dostawców platform, aby upewnić się, że mogą obsługiwać wymagające obciążenia naszych klientów. Korzystamy z usług chmurowych od niektórych z najlepszych dostawców chmurowych na świecie, którzy obsługują firmy takie jak [Spotify](https://spotify.com) i [Netflix](https://netflix.com), które mają kilka rzędów wielkości więcej ruchu niż my.

Główni dostawcy chmurowi, z których korzystamy, to:

- **[Cloudflare](https://cloudflare.com)**: Zarządzamy DNS (usługą nazw domen) za pośrednictwem Cloudflare, a także naszą stroną marketingową, która działa na [Cloudflare Pages](https://pages.cloudflare.com/).
- **[Amazon Web Services](https://aws.amazon.com/)**: Używamy AWS do naszej bazy danych, która jest [Aurora](https://aws.amazon.com/rds/aurora/), do przechowywania plików za pośrednictwem [Simple Storage Service (S3)](https://aws.amazon.com/s3/), a także do wysyłania e-maili za pośrednictwem [Simple Email Service (SES)](https://aws.amazon.com/ses/)
- **[Render](https://render.com)**: Używamy Render do naszych serwerów front-end, serwerów aplikacji/API, naszych usług w tle, systemu kolejkowania i bazy danych Redis. Co ciekawe, Render jest faktycznie zbudowany *na bazie* AWS! 


## Jak bezpieczne są pliki w Blue?

Zacznijmy od przechowywania danych. Nasze pliki są hostowane na [AWS S3](https://aws.amazon.com/s3/), które jest najpopularniejszym na świecie obiektem przechowywania w chmurze, oferującym wiodącą w branży skalowalność, dostępność danych, bezpieczeństwo i wydajność.

Mamy 99,99% dostępności plików i 99,999999999% wysokiej trwałości.

Rozłóżmy to, co to oznacza.

Dostępność odnosi się do ilości czasu, w którym dane są operacyjne i dostępne. 99,99% dostępności plików oznacza, że możemy oczekiwać, że pliki będą niedostępne nie dłużej niż około 8,76 godziny rocznie.

Trwałość odnosi się do prawdopodobieństwa, że dane pozostaną nienaruszone i nieuszkodzone w czasie. Ten poziom trwałości oznacza, że możemy oczekiwać, że stracimy nie więcej niż jeden plik na 10 miliardów przesłanych plików, dzięki rozległej redundancji i replikacji danych w wielu centrach danych.

Używamy [S3 Intelligent-Tiering](https://aws.amazon.com/s3/storage-classes/intelligent-tiering/), aby automatycznie przenosić pliki do różnych klas przechowywania w zależności od częstotliwości dostępu. Na podstawie wzorców aktywności setek tysięcy projektów zauważamy, że większość plików jest dostępna w wzorze przypominającym krzywą eksponencjalnego wycofania. Oznacza to, że większość plików jest bardzo często dostępna w pierwszych kilku dniach, a następnie szybko dostępność maleje. Pozwala to na przeniesienie starszych plików do wolniejszego, ale znacznie tańszego przechowywania, bez wpływu na doświadczenia użytkowników w znaczący sposób.

Oszczędności kosztów z tego tytułu są znaczące. S3 Standard-Infrequent Access (S3 Standard-IA) jest około 1,84 razy tańszy niż S3 Standard. Oznacza to, że za każdy dolar, który wydalibyśmy na S3 Standard, wydajemy tylko około 54 centów na S3 Standard-IA za tę samą ilość przechowywanych danych.

| Funkcja                  | S3 Standard             | S3 Standard-IA       |
|--------------------------|-------------------------|-----------------------|
| Koszt przechowywania     | $0.023 - $0.021 za GB   | $0.0125 za GB         |
| Koszt żądania (PUT itp.) | $0.005 za 1 000 żądań   | $0.01 za 1 000 żądań  |
| Koszt żądania (GET)      | $0.0004 za 1 000 żądań  | $0.001 za 1 000 żądań |
| Koszt odzyskiwania danych | $0.00 za GB            | $0.01 za GB           |


Pliki, które przesyłasz przez Blue, są szyfrowane zarówno w tranzycie, jak i w spoczynku. Dane przesyłane do i z Amazon S3 są zabezpieczone za pomocą [Transport Layer Security (TLS)](https://www.internetsociety.org/deploy360/tls/basics), chroniąc przed [podsłuchiwaniem](https://en.wikipedia.org/wiki/Network_eavesdropping) i [atakami typu man-in-the-middle](https://en.wikipedia.org/wiki/Man-in-the-middle_attack). W przypadku szyfrowania w spoczynku Amazon S3 używa szyfrowania po stronie serwera (SSE-S3), które automatycznie szyfruje wszystkie nowe przesyłki za pomocą szyfrowania AES-256, a Amazon zarządza kluczami szyfrującymi. To zapewnia, że twoje dane pozostają bezpieczne przez cały okres ich życia.

## A co z danymi, które nie są plikami?

Nasza baza danych jest zasilana przez [AWS Aurora](https://aws.amazon.com/rds/aurora/), nowoczesną usługę relacyjnej bazy danych, która zapewnia wysoką wydajność, dostępność i bezpieczeństwo twoich danych.

Dane w Aurorze są szyfrowane zarówno w tranzycie, jak i w spoczynku. Używamy SSL (AES-256), aby zabezpieczyć połączenia między twoją instancją bazy danych a twoją aplikacją, chroniąc dane podczas transferu. W przypadku szyfrowania w spoczynku Aurora używa kluczy zarządzanych przez AWS Key Management Service (KMS), zapewniając, że wszystkie przechowywane dane, w tym automatyczne kopie zapasowe, migawki i repliki, są szyfrowane i chronione.

Aurora ma rozproszony, odporny na błędy i samonaprawiający się system przechowywania. System ten jest oddzielony od zasobów obliczeniowych i może automatycznie skalować się do 128 TiB na instancję bazy danych. Dane są replikowane w trzech [Strefach dostępności](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) (AZ), co zapewnia odporność na utratę danych i zapewnia wysoką dostępność. W przypadku awarii bazy danych Aurora skraca czasy odzyskiwania do mniej niż 60 sekund, zapewniając minimalne zakłócenia.

Blue nieustannie tworzy kopie zapasowe naszej bazy danych do Amazon S3, umożliwiając odzyskiwanie w określonym czasie. Oznacza to, że możemy przywrócić główną bazę danych Blue do dowolnego konkretnego momentu w ciągu ostatnich pięciu minut, zapewniając, że twoje dane są zawsze możliwe do odzyskania. Regularnie wykonujemy również migawki bazy danych na dłuższe okresy przechowywania kopii zapasowych.

Jako w pełni zarządzana usługa, Aurora automatyzuje czasochłonne zadania administracyjne, takie jak dostarczanie sprzętu, konfiguracja bazy danych, aktualizacje i kopie zapasowe. To zmniejsza obciążenie operacyjne i zapewnia, że nasza baza danych jest zawsze aktualna z najnowszymi poprawkami bezpieczeństwa i ulepszeniami wydajności.

Jeśli jesteśmy bardziej efektywni, możemy przekazać nasze oszczędności kosztów naszym klientom dzięki naszym [wiodącym w branży cenom](/pricing).

Aurora jest zgodna z różnymi standardami branżowymi, takimi jak HIPAA, GDPR i SOC 2, zapewniając, że twoje praktyki zarządzania danymi spełniają rygorystyczne wymagania regulacyjne. Regularne audyty bezpieczeństwa i integracja z [Amazon GuardDuty](https://aws.amazon.com/guardduty/) pomagają wykrywać i łagodzić potencjalne zagrożenia bezpieczeństwa.

## Jak Blue zapewnia bezpieczeństwo logowania?

Blue używa [magicznych linków za pośrednictwem e-maila](https://documentation.blue.cc/user-management/magic-links), aby zapewnić bezpieczny i wygodny dostęp do twojego konta, eliminując potrzebę tradycyjnych haseł.

To podejście znacznie zwiększa bezpieczeństwo, łagodząc powszechne zagrożenia związane z logowaniem opartym na hasłach. Eliminując hasła, magiczne linki chronią przed atakami phishingowymi i kradzieżą haseł, *ponieważ nie ma hasła do kradzieży ani wykorzystania.*

Każdy magiczny link jest ważny tylko na jedną sesję logowania, co zmniejsza ryzyko nieautoryzowanego dostępu. Dodatkowo, te linki wygasają po 15 minutach, zapewniając, że wszelkie niewykorzystane linki nie mogą być wykorzystane, co dodatkowo zwiększa bezpieczeństwo.

Warto również zauważyć wygodę oferowaną przez magiczne linki. Magiczne linki zapewniają bezproblemowe doświadczenie logowania, pozwalając ci uzyskać dostęp do swojego konta *bez* potrzeby zapamiętywania skomplikowanych haseł.

To nie tylko upraszcza proces logowania, ale także zapobiega naruszeniom bezpieczeństwa, które występują, gdy hasła są używane w wielu usługach. Wiele osób ma tendencję do używania tego samego hasła w różnych platformach, co oznacza, że naruszenie bezpieczeństwa w jednej usłudze może zagrozić ich kontom w innych usługach, w tym Blue. Używając magicznych linków, bezpieczeństwo Blue nie zależy od praktyk bezpieczeństwa innych usług, zapewniając bardziej solidną i niezależną warstwę ochrony dla naszych użytkowników.

Kiedy żądasz logowania do swojego konta Blue, unikalny adres URL logowania jest wysyłany na twój e-mail. Kliknięcie tego linku natychmiast zaloguje cię do twojego konta. Link jest zaprojektowany tak, aby wygasał po jednokrotnym użyciu lub po 15 minutach, w zależności od tego, co nastąpi wcześniej, co dodaje dodatkową warstwę bezpieczeństwa. Dzięki użyciu magicznych linków, Blue zapewnia, że proces logowania jest zarówno bezpieczny, jak i przyjazny dla użytkownika, zapewniając spokój ducha i wygodę.

## Jak mogę sprawdzić niezawodność i dostępność Blue?

W Blue zobowiązujemy się do utrzymania wysokiego poziomu niezawodności i przejrzystości dla naszych użytkowników. Aby zapewnić widoczność wydajności naszej platformy, oferujemy [dedykowaną stronę statusu systemu](https://status.blue.cc), która jest również linkowana z naszego stopki na każdej stronie naszej witryny.

![](/insights/status-page.png)

Ta strona wyświetla nasze historyczne dane dotyczące dostępności, pozwalając ci zobaczyć, jak konsekwentnie nasze usługi były dostępne w czasie. Dodatkowo, strona statusu zawiera szczegółowe raporty o incydentach, zapewniając przejrzystość na temat wszelkich przeszłych problemów, ich wpływu oraz kroków, które podjęliśmy, aby je rozwiązać i zapobiec ich wystąpieniu w przyszłości.
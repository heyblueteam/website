---
title: Wyszukiwanie w czasie rzeczywistym
description: Blue wprowadza nową, błyskawicznie szybką wyszukiwarkę, która zwraca wyniki we wszystkich Twoich projektach w milisekundach, umożliwiając Ci zmianę kontekstu w mgnieniu oka.
category: "Product Updates"
date: 2024-03-01
---



Z radością ogłaszamy uruchomienie naszej nowej wyszukiwarki, zaprojektowanej w celu zrewolucjonizowania sposobu, w jaki znajdujesz informacje w Blue. Efektywna funkcjonalność wyszukiwania jest kluczowa dla płynnego zarządzania projektami, a nasza najnowsza aktualizacja zapewnia, że możesz uzyskać dostęp do swoich danych szybciej niż kiedykolwiek.

Nasza nowa wyszukiwarka pozwala na wyszukiwanie wszystkich komentarzy, plików, rekordów, pól niestandardowych, opisów i list kontrolnych. Niezależnie od tego, czy musisz znaleźć konkretny komentarz dotyczący projektu, szybko zlokalizować plik, czy wyszukać określony rekord lub pole, nasza wyszukiwarka dostarcza błyskawiczne wyniki.

Gdy narzędzia osiągają responsywność na poziomie 50-100 ms, mają tendencję do znikania i stapiania się w tle, zapewniając płynne i niemal niewidoczne doświadczenie użytkownika. Dla kontekstu, mrugnięcie oka trwa około 60-120 ms, więc 50 ms jest w rzeczywistości szybsze niż mrugnięcie! Taki poziom responsywności pozwala Ci na interakcję z Blue, nawet nie zdając sobie sprawy, że to tam jest, uwalniając Cię od konieczności skupiania się na rzeczywistej pracy. Dzięki wykorzystaniu tego poziomu wydajności, nasza nowa wyszukiwarka zapewnia, że możesz szybko uzyskać dostęp do potrzebnych informacji, nie przeszkadzając w Twoim przepływie pracy.

Aby osiągnąć nasz cel błyskawicznego wyszukiwania, wykorzystaliśmy najnowsze technologie open-source. Nasza wyszukiwarka oparta jest na MeiliSearch, popularnej usłudze wyszukiwania jako usługi, która wykorzystuje przetwarzanie języka naturalnego i wyszukiwanie wektorowe, aby szybko znajdować odpowiednie wyniki. Dodatkowo, wdrożyliśmy pamięć operacyjną, co pozwala nam przechowywać często używane dane w RAM, skracając czas potrzebny na zwrócenie wyników wyszukiwania. Ta kombinacja MeiliSearch i pamięci operacyjnej umożliwia naszej wyszukiwarce dostarczanie wyników w milisekundach, co pozwala Ci szybko znaleźć to, czego potrzebujesz, nie myśląc o technologii, która za tym stoi.

Nowy pasek wyszukiwania znajduje się wygodnie na pasku nawigacyjnym, co pozwala na natychmiastowe rozpoczęcie wyszukiwania. Aby uzyskać bardziej szczegółowe doświadczenie wyszukiwania, wystarczy nacisnąć klawisz Tab podczas wyszukiwania, aby uzyskać dostęp do pełnej strony wyszukiwania. Dodatkowo, możesz szybko aktywować funkcję wyszukiwania z dowolnego miejsca, używając skrótu CMD/Ctrl+K, co jeszcze bardziej ułatwia znalezienie tego, czego potrzebujesz.

<video autoplay loop muted playsinline>
  <source src="/videos/search-demo.mp4" type="video/mp4">
</video>


## Przyszłe rozwinięcia

To dopiero początek. Teraz, gdy mamy infrastrukturę wyszukiwania nowej generacji, możemy w przyszłości robić naprawdę interesujące rzeczy.

Następną funkcją będzie wyszukiwanie semantyczne, które jest znaczącym ulepszeniem typowego wyszukiwania opartego na słowach kluczowych. Pozwól, że wyjaśnię.

Ta funkcja pozwoli wyszukiwarce zrozumieć kontekst Twoich zapytań. Na przykład, wyszukiwanie "morze" zwróci odpowiednie dokumenty, nawet jeśli dokładna fraza nie jest używana. Możesz pomyśleć "ale wpisałem 'ocean' zamiast!" - i masz rację. Wyszukiwarka również zrozumie podobieństwo między "morze" a "ocean" i zwróci odpowiednie dokumenty, nawet jeśli dokładna fraza nie jest używana. Ta funkcja jest szczególnie przydatna przy wyszukiwaniu dokumentów zawierających terminy techniczne, akronimy lub po prostu powszechne słowa, które mają wiele wariantów lub literówek.

Kolejną nadchodzącą funkcją jest możliwość wyszukiwania obrazów według ich zawartości. Aby to osiągnąć, będziemy przetwarzać każdy obraz w Twoim projekcie, tworząc dla każdego z nich osadzenie. W dużym uproszczeniu, osadzenie to matematyczny zbiór współrzędnych, który odpowiada znaczeniu obrazu. Oznacza to, że wszystkie obrazy mogą być wyszukiwane na podstawie tego, co zawierają, niezależnie od ich nazwy pliku czy metadanych. Wyobraź sobie wyszukiwanie "schemat blokowy" i znajdowanie wszystkich obrazów związanych ze schematami blokowymi, *niezależnie od ich nazw plików.*
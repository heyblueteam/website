---
title: Referencje i pola wyszukiwania niestandardowe
description: Bez wysiłku twórz powiązane projekty w Blue, przekształcając go w jedno źródło prawdy dla Twojego biznesu dzięki nowym polom Referencji i Wyszukiwania.
category: "Product Updates"
date: 2023-11-01
---



Projekty w Blue są już potężnym sposobem zarządzania danymi biznesowymi i posuwania pracy do przodu.

Dziś podejmujemy kolejny logiczny krok, pozwalając Ci połączyć dane *między* projektami, aby uzyskać maksymalną elastyczność i moc.

Łączenie projektów w Blue przekształca go w jedno źródło prawdy dla Twojego biznesu. Ta funkcjonalność umożliwia stworzenie kompleksowego i powiązanego zbioru danych, co pozwala na płynny przepływ informacji i zwiększoną widoczność w projektach. Dzięki łączeniu projektów zespoły mogą osiągnąć jednolity obraz operacji, co poprawia podejmowanie decyzji i efektywność operacyjną.

## Przykład

Rozważ firmę ACME, która wykorzystuje pola Referencji i Wyszukiwania w Blue do stworzenia powiązanego ekosystemu danych w swoich projektach Klient, Sprzedaż i Magazyn. Rekordy klientów w projekcie Klient są powiązane za pomocą pól Referencji z transakcjami sprzedaży w projekcie Sprzedaż. To powiązanie pozwala polom Wyszukiwania na pobieranie powiązanych szczegółów klientów, takich jak numery telefonów i statusy kont, bezpośrednio do każdego rekordu sprzedaży. Dodatkowo, sprzedawane pozycje magazynowe są wyświetlane w rekordzie sprzedaży za pomocą pola Wyszukiwania odnoszącego się do danych Ilość Sprzedana z projektu Magazyn. Na koniec, wycofania z magazynu są powiązane z odpowiednimi sprzedażami za pomocą pola Referencji w Magazynie, wskazującego na rekordy Sprzedaży. Ta konfiguracja zapewnia pełną widoczność tego, która sprzedaż spowodowała usunięcie z magazynu, tworząc zintegrowany widok 360 stopni w projektach.

## Jak działają pola Referencji

Pola niestandardowe Referencji umożliwiają tworzenie relacji między rekordami w różnych projektach w Blue. Podczas tworzenia pola Referencji, Administrator Projektu wybiera konkretny projekt, który dostarczy listę rekordów referencyjnych. Opcje konfiguracyjne obejmują:

* **Wybór pojedynczy**: Pozwala na wybór jednego rekordu referencyjnego.
* **Wybór wielokrotny**: Pozwala na wybór wielu rekordów referencyjnych.
* **Filtrowanie**: Ustaw filtry, aby umożliwić użytkownikom wybór tylko rekordów, które spełniają kryteria filtru.

Po skonfigurowaniu użytkownicy mogą wybierać konkretne rekordy z menu rozwijanego w polu Referencji, ustanawiając link między projektami.

## Rozszerzanie pól referencyjnych za pomocą wyszukiwań

Pola niestandardowe Wyszukiwania są używane do importowania danych z rekordów w innych projektach, tworząc jednokierunkową widoczność. Zawsze są tylko do odczytu i są powiązane z konkretnym polem niestandardowym Referencji. Gdy użytkownik wybiera jeden lub więcej rekordów za pomocą pola niestandardowego Referencji, pole niestandardowe Wyszukiwania pokaże dane z tych rekordów. Wyszukiwania mogą wyświetlać dane takie jak:

* Utworzone w
* Zaktualizowane w
* Termin
* Opis
* Lista
* Etykieta
* Osoba przypisana
* Jakiekolwiek wspierane pole niestandardowe z rekordu referencyjnego — w tym inne pola wyszukiwania!


Na przykład, wyobraź sobie scenariusz, w którym masz trzy projekty: **Projekt A** to projekt sprzedażowy, **Projekt B** to projekt zarządzania magazynem, a **Projekt C** to projekt relacji z klientami. W Projekcie A masz pole niestandardowe Referencji, które łączy rekordy sprzedaży z odpowiadającymi rekordami klientów w Projekcie C. W Projekcie B masz pole niestandardowe Wyszukiwania, które importuje informacje z Projektu A, takie jak ilość sprzedana. W ten sposób, gdy rekord sprzedaży jest tworzony w Projekcie A, informacje o kliencie związane z tą sprzedażą są automatycznie pobierane z Projektu C, a ilość sprzedana jest automatycznie pobierana z Projektu B. To pozwala na utrzymanie wszystkich istotnych informacji w jednym miejscu i przeglądanie ich bez konieczności tworzenia zduplikowanych danych lub ręcznego aktualizowania rekordów w różnych projektach.

Przykładem tego w rzeczywistości jest firma e-commerce, która używa Blue do zarządzania swoimi sprzedażami, magazynem i relacjami z klientami. W ich projekcie **Sprzedaż** mają pole niestandardowe Referencji, które łączy każdy rekord sprzedaży z odpowiadającym rekordem **Klienta** w ich projekcie **Klienci**. W ich projekcie **Magazyn** mają pole niestandardowe Wyszukiwania, które importuje informacje z projektu Sprzedaż, takie jak ilość sprzedana, i wyświetla je w rekordzie pozycji magazynowej. To pozwala im łatwo zobaczyć, które sprzedaże powodują usunięcia z magazynu i utrzymać poziomy zapasów na bieżąco bez konieczności ręcznego aktualizowania rekordów w różnych projektach.

## Podsumowanie

Wyobraź sobie świat, w którym Twoje dane projektowe nie są izolowane, ale swobodnie przepływają między projektami, dostarczając kompleksowych informacji i zwiększając efektywność. To moc pól Referencji i Wyszukiwania w Blue. Umożliwiając płynne połączenia danych i zapewniając widoczność w czasie rzeczywistym w projektach, te funkcje zmieniają sposób, w jaki zespoły współpracują i podejmują decyzje. Niezależnie od tego, czy zarządzasz relacjami z klientami, śledzisz sprzedaż, czy nadzorujesz magazyn, pola Referencji i Wyszukiwania w Blue umożliwiają Twojemu zespołowi pracę mądrzej, szybciej i skuteczniej. Zanurz się w powiązanym świecie Blue i obserwuj, jak Twoja produktywność rośnie.

[Sprawdź dokumentację](https://documentation.blue.cc/custom-fields/reference) lub [zarejestruj się i wypróbuj to samodzielnie.](https://app.blue.cc)
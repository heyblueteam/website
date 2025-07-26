---
title:  Dlaczego Blue ma otwartą betę
description: Dowiedz się, dlaczego nasz system zarządzania projektami ma trwającą otwartą betę.
category: "Engineering"
date: 2024-08-03
---



Wiele startupów B2B SaaS uruchamia swoje produkty w wersji beta, i to z dobrych powodów. To część tradycyjnego motta Doliny Krzemowej *„działaj szybko i łam zasady”*.

Umieszczenie naklejki „beta” na produkcie obniża oczekiwania.

Coś jest zepsute? Cóż, to tylko beta.

System działa wolno? Cóż, to tylko beta.

[Dokumentacja](https://blue.cc/docs) nie istnieje? Cóż… rozumiesz o co chodzi.

I to jest *właściwie* dobra rzecz. Reid Hoffman, założyciel LinkedIn, słynnie powiedział:

> Jeśli nie wstydzisz się pierwszej wersji swojego produktu, uruchomiłeś go za późno.

A naklejka beta jest również dobra dla klientów. Pomaga im w samodzielnym wyborze.

Klienci, którzy próbują produktów beta, to ci, którzy znajdują się na wczesnych etapach Cyklu Przyjmowania Technologii, znanego również jako Krzywa Przyjmowania Produktu.

Cykl Przyjmowania Technologii jest zazwyczaj podzielony na pięć głównych segmentów:

1. Innowatorzy
2. Wczesni Adopterzy
3. Wczesna Większość
4. Późna Większość
5. Opóźnieni

![](/insights/technology-adoption-lifecycle-graph.png)


Jednak w końcu produkt musi dojrzeć, a klienci oczekują stabilnego, działającego produktu. Nie chcą mieć dostępu do środowiska „beta”, w którym rzeczy się psują.

Czyżby?

*To* jest pytanie, które zadaliśmy sobie.

Wierzymy, że zadaliśmy sobie to pytanie z powodu natury, w jakiej Blue zostało początkowo zbudowane. [Blue zaczęło jako odgałęzienie zajętej agencji projektowej](/insights/agency-success-playbook), więc pracowaliśmy *wewnątrz* biura firmy, która aktywnie korzystała z Blue do zarządzania wszystkimi swoimi projektami.

Oznacza to, że przez lata mogliśmy obserwować, jak *prawdziwi* ludzie — siedzący tuż obok nas! — używali Blue w swoim codziennym życiu.

A ponieważ używali Blue od pierwszych dni, ten zespół zawsze korzystał z Blue Beta!

Dlatego naturalne było dla nas, aby umożliwić wszystkim naszym innym klientom korzystanie z tego samego.

**I dlatego nie mamy dedykowanego zespołu testowego.**

Zgadza się.

Nikt w Blue nie ma *wyłącznej* odpowiedzialności za zapewnienie, że nasza platforma działa dobrze i stabilnie.

Jest to spowodowane kilkoma powodami.

Pierwszym jest niższa baza kosztowa.

Brak pełnoetatowego zespołu testowego znacznie obniża nasze koszty, a my możemy przekazać te oszczędności naszym klientom, oferując najniższe ceny w branży.

Aby to zobrazować, oferujemy zestawy funkcji na poziomie przedsiębiorstwa, za które nasza konkurencja pobiera od 30 do 55 USD/użytkownika/miesiąc, za jedyne 7 USD/miesiąc.

To nie dzieje się przypadkowo, to *jest zaprojektowane*.

Jednak nie jest dobrą strategią sprzedawać tańszy produkt, jeśli nie działa.

Więc *prawdziwe pytanie brzmi*, jak udaje nam się stworzyć stabilną platformę, z której mogą korzystać tysiące klientów bez dedykowanego zespołu testowego?

Oczywiście, nasze podejście do posiadania otwartej bety jest kluczowe, ale zanim w to wejdziemy, chcemy poruszyć odpowiedzialność dewelopera.

Podjęliśmy wczesną decyzję w Blue, że nigdy nie podzielimy odpowiedzialności za technologie front-end i back-end. Zatrudnialibyśmy lub szkolili tylko deweloperów full stack.

Powód, dla którego podjęliśmy tę decyzję, był taki, aby zapewnić, że deweloper w pełni odpowiada za funkcję, nad którą pracuje. Tak więc nie byłoby mentalności *„rzucania problemu przez płot”*, którą czasami można spotkać, gdy istnieją wspólne odpowiedzialności za funkcje.

I to rozszerza się na testowanie funkcji, zrozumienie przypadków użycia klientów i ich próśb oraz czytanie i komentowanie specyfikacji.

Innymi słowy, każdy deweloper buduje głębokie i intuicyjne zrozumienie funkcji, którą tworzy.

Dobrze, porozmawiajmy teraz o naszej otwartej becie.

Kiedy mówimy, że jest „otwarta” — mamy na myśli to. Każdy klient może spróbować, po prostu dodając „beta” przed adresem URL naszej aplikacji internetowej.

Więc „app.blue.cc” staje się „beta.app.blue.cc”

Kiedy to robią, mogą zobaczyć swoje zwykłe dane, ponieważ zarówno środowiska Beta, jak i Produkcyjne dzielą tę samą bazę danych, ale będą mogli również zobaczyć nowe funkcje.

Klienci mogą łatwo pracować, nawet jeśli niektórzy członkowie zespołu są w Produkcji, a inni ciekawscy są w Becie.

Zazwyczaj mamy kilka setek klientów korzystających z Bety w danym momencie, a na naszych forach społecznościowych publikujemy zapowiedzi funkcji, aby mogli sprawdzić, co nowego i wypróbować to.

I to jest kluczowe: mamy *kilkaset* testerów!

Wszyscy ci klienci przetestują funkcje w swoich przepływach pracy i będą dość głośni, jeśli coś nie będzie w porządku, ponieważ *już* wdrażają tę funkcję w swojej firmie!

Najczęstsze opinie to małe, ale bardzo przydatne zmiany, które dotyczą przypadków brzegowych, których nie wzięliśmy pod uwagę.

Nowe funkcje pozostawiamy w Becie na okres od 2 do 4 tygodni. Kiedy czujemy, że są stabilne, wówczas wprowadzamy je do produkcji.

Mamy również możliwość ominięcia Bety, jeśli zajdzie taka potrzeba, używając flagi szybkiej ścieżki. Zazwyczaj robi się to w przypadku poprawek błędów, których nie chcemy trzymać przez 2-4 tygodnie przed wprowadzeniem do produkcji.

A jaki jest rezultat?

Wprowadzanie do produkcji wydaje się… cóż, nudne! Jak nic — to po prostu nie jest dla nas wielka sprawa.

I oznacza to, że to wygładza nasz harmonogram wydania, co umożliwiło nam [wprowadzanie funkcji co miesiąc jak w zegarku przez ostatnie sześć lat.](/changelog).

Jednak, jak każda decyzja, są pewne kompromisy.

Wsparcie klientów jest nieco bardziej skomplikowane, ponieważ musimy wspierać klientów w dwóch wersjach naszej platformy. Czasami może to powodować zamieszanie wśród klientów, którzy mają członków zespołu korzystających z dwóch różnych wersji.

Innym punktem bólu jest to, że takie podejście może czasami spowolnić ogólny harmonogram wydania do produkcji. Dotyczy to szczególnie większych funkcji, które mogą utknąć w Becie, jeśli istnieje inna powiązana funkcja, która ma problemy i wymaga dalszej pracy.

Jednak w równowadze uważamy, że te kompromisy są warte korzyści z niższej bazy kosztowej i większego zaangażowania klientów.

Jesteśmy jedną z nielicznych firm programistycznych, które przyjęły to podejście, ale teraz jest to fundamentalna część naszego podejścia do rozwoju produktu.
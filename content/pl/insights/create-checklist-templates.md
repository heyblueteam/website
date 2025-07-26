---
title: Tworzenie wielokrotnego użytku list kontrolnych za pomocą automatyzacji
description: Dowiedz się, jak tworzyć automatyzacje zarządzania projektami dla wielokrotnego użytku list kontrolnych.
category: "Best Practices"
date: 2024-07-08
---



W wielu projektach i procesach może być konieczne użycie tej samej listy kontrolnej w wielu rekordach lub zadaniach.

Jednak ręczne przepisywanie listy kontrolnej za każdym razem, gdy chcesz ją dodać do rekordu, nie jest zbyt efektywne. W tym miejscu możesz wykorzystać [potężne automatyzacje zarządzania projektami](/platform/features/automations), aby zrobić to automatycznie!

Przypominamy, że automatyzacje w Blue wymagają dwóch kluczowych elementów:

1. Wyzwalacz — Co powinno się wydarzyć, aby uruchomić automatyzację. Może to być, gdy rekord otrzyma konkretną etykietę, gdy przejdzie do konkretnego 
2. Jedna lub więcej Akcji — W tym przypadku byłoby to automatyczne utworzenie jednej lub więcej list kontrolnych.

Zacznijmy od akcji, a następnie omówimy możliwe wyzwalacze, które możesz wykorzystać.

## Akcja Automatyzacji Listy Kontrolnej

Możesz stworzyć nową automatyzację i skonfigurować jedną lub więcej list kontrolnych do utworzenia, zgodnie z poniższym przykładem:

![](/insights/checklist-automation.png)

To będą listy kontrolne, które chcesz, aby były tworzone za każdym razem, gdy podejmiesz tę akcję.

## Wyzwalacze Automatyzacji Listy Kontrolnej

Istnieje kilka sposobów, aby wyzwolić utworzenie wielokrotnego użytku list kontrolnych. Oto kilka popularnych opcji:

- **Dodanie konkretnej etykiety:** Możesz skonfigurować automatyzację, aby uruchamiała się, gdy do rekordu dodana zostanie konkretna etykieta. Na przykład, gdy dodana zostanie etykieta "Nowy Projekt", może to automatycznie utworzyć twoją listę kontrolną do inicjacji projektu.
- **Przypisanie rekordu:** Utworzenie listy kontrolnej może być wyzwolone, gdy rekord zostanie przypisany do konkretnej osoby lub do kogokolwiek. To jest przydatne dla list kontrolnych do onboardingu lub procedur specyficznych dla zadań.
- **Przeniesienie do konkretnej listy:** Gdy rekord zostanie przeniesiony do konkretnej listy na twojej tablicy projektowej, może to wyzwolić utworzenie odpowiedniej listy kontrolnej. Na przykład, przeniesienie elementu do listy "Kontrola Jakości" może wyzwolić listę kontrolną QA.
- **Niestandardowe pole wyboru:** Utwórz niestandardowe pole wyboru i ustaw automatyzację, aby uruchamiała się, gdy to pole zostanie zaznaczone. Daje to manualną kontrolę nad tym, kiedy dodać listę kontrolną.
- **Pojedyncze lub wielokrotne pola wyboru:** Możesz stworzyć pojedyncze lub wielokrotne pole wyboru z różnymi opcjami. Każda opcja może być powiązana z konkretnym szablonem listy kontrolnej przez oddzielne automatyzacje. To pozwala na bardziej szczegółową kontrolę i możliwość posiadania wielu szablonów list kontrolnych gotowych do różnych scenariuszy.

Aby zwiększyć kontrolę nad tym, kto może wyzwalać te automatyzacje, możesz ukryć te niestandardowe pola przed niektórymi użytkownikami, korzystając z niestandardowych ról użytkowników. Zapewnia to, że tylko administratorzy projektów lub inny uprawniony personel mogą uruchamiać te opcje.

Pamiętaj, że kluczem do skutecznego wykorzystania wielokrotnego użytku list kontrolnych z automatyzacjami jest staranne zaprojektowanie wyzwalaczy. Weź pod uwagę przepływ pracy swojego zespołu, rodzaje projektów, którymi się zajmujesz, oraz kto powinien mieć możliwość inicjowania różnych procesów. Dzięki dobrze zaplanowanym automatyzacjom możesz znacznie uprościć zarządzanie projektami i zapewnić spójność w swoich operacjach.

## Przydatne zasoby

- [Dokumentacja automatyzacji zarządzania projektami](https://documentation.blue.cc/automations)
- [Dokumentacja niestandardowych ról użytkowników](https://documentation.blue.cc/user-management/roles/custom-user-roles)
- [Dokumentacja pól niestandardowych](https://documentation.blue.cc/custom-fields)
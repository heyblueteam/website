--- 
title: Wiederverwendbare Checklisten mit Automatisierungen erstellen
category: "Best Practices"
description: Erfahren Sie, wie Sie Automatisierungen im Projektmanagement für wiederverwendbare Checklisten erstellen.
date: 2024-07-08
---

In vielen Projekten und Prozessen müssen Sie möglicherweise dieselbe Checkliste über mehrere Datensätze oder Aufgaben hinweg verwenden.

Es ist jedoch nicht sehr effizient, die Checkliste jedes Mal manuell neu einzugeben, wenn Sie sie zu einem Datensatz hinzufügen möchten. Hier können Sie [leistungsstarke Automatisierungen im Projektmanagement](/platform/features/automations) nutzen, um dies automatisch für Sie zu erledigen!

Zur Erinnerung: Automatisierungen in Blue erfordern zwei Dinge:

1. Einen Trigger — Was passieren soll, um die Automatisierung zu starten. Dies kann sein, wenn einem Datensatz ein bestimmtes Tag zugewiesen wird, wenn er in eine bestimmte 
2. Eine oder mehrere Aktionen — In diesem Fall wäre es die automatische Erstellung von einer oder mehreren Checklisten.

Lassen Sie uns zuerst mit der Aktion beginnen und dann die möglichen Trigger besprechen, die Sie verwenden können.

## Aktion zur Automatisierung von Checklisten

Sie können eine neue Automatisierung erstellen und eine oder mehrere Checklisten einrichten, die erstellt werden sollen, wie im folgenden Beispiel:

![](/insights/checklist-automation.png)

Dies wären die Checkliste(n), die jedes Mal erstellt werden sollen, wenn Sie die Aktion ausführen.

## Trigger für die Automatisierung von Checklisten

Es gibt mehrere Möglichkeiten, die Erstellung Ihrer wiederverwendbaren Checklisten auszulösen. Hier sind einige beliebte Optionen:

- **Hinzufügen eines bestimmten Tags:** Sie können die Automatisierung so einrichten, dass sie ausgelöst wird, wenn ein bestimmtes Tag zu einem Datensatz hinzugefügt wird. Zum Beispiel könnte beim Hinzufügen des Tags "Neues Projekt" automatisch Ihre Checkliste zur Projektinitiierung erstellt werden.
- **Zuweisung von Datensätzen:** Die Erstellung der Checkliste kann ausgelöst werden, wenn ein Datensatz einer bestimmten Person oder jemandem zugewiesen wird. Dies ist nützlich für Onboarding-Checklisten oder aufgabenbezogene Verfahren.
- **Verschieben in eine bestimmte Liste:** Wenn ein Datensatz in eine bestimmte Liste auf Ihrem Projektboard verschoben wird, kann dies die Erstellung einer relevanten Checkliste auslösen. Zum Beispiel könnte das Verschieben eines Elements in eine "Qualitätssicherung"-Liste eine QA-Checkliste auslösen.
- **Benutzerdefiniertes Kontrollkästchenfeld:** Erstellen Sie ein benutzerdefiniertes Kontrollkästchenfeld und setzen Sie die Automatisierung so, dass sie ausgelöst wird, wenn dieses Feld angekreuzt wird. Dies gibt Ihnen manuelle Kontrolle darüber, wann die Checkliste hinzugefügt wird.
- **Einzel- oder Mehrfachauswahl benutzerdefinierte Felder:** Sie können ein Einzel- oder Mehrfachauswahl benutzerdefiniertes Feld mit verschiedenen Optionen erstellen. Jede Option kann über separate Automatisierungen mit einer bestimmten Checklisten-Vorlage verknüpft werden. Dies ermöglicht eine granularere Kontrolle und die Möglichkeit, mehrere Checklisten-Vorlagen für verschiedene Szenarien bereitzuhalten.

Um die Kontrolle darüber zu verbessern, wer diese Automatisierungen auslösen kann, können Sie diese benutzerdefinierten Felder für bestimmte Benutzer mithilfe benutzerdefinierter Benutzerrollen ausblenden. Dies stellt sicher, dass nur Projektadministratoren oder andere autorisierte Personen diese Optionen auslösen können.

Denken Sie daran, dass der Schlüssel zur effektiven Nutzung von wiederverwendbaren Checklisten mit Automatisierungen darin besteht, Ihre Trigger durchdacht zu gestalten. Berücksichtigen Sie den Workflow Ihres Teams, die Arten von Projekten, die Sie bearbeiten, und wer die Möglichkeit haben sollte, verschiedene Prozesse zu initiieren. Mit gut geplanten Automatisierungen können Sie Ihr Projektmanagement erheblich optimieren und Konsistenz in Ihren Abläufen gewährleisten.

## Nützliche Ressourcen

- [Dokumentation zur Projektmanagement-Automatisierung](https://documentation.blue.cc/automations)
- [Dokumentation zu benutzerdefinierten Benutzerrollen](https://documentation.blue.cc/user-management/roles/custom-user-roles)
- [Dokumentation zu benutzerdefinierten Feldern](https://documentation.blue.cc/custom-fields)
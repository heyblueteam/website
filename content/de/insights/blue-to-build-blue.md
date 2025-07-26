---
title:  Wie wir Blue nutzen, um Blue zu bauen. 
category: "CEO Blog"
description: Erfahren Sie, wie wir unsere eigene Projektmanagement-Plattform nutzen, um unsere Projektmanagement-Plattform zu entwickeln! 
date: 2024-08-07
---

Sie stehen kurz davor, einen Insider-Blick darauf zu werfen, wie Blue Blue baut.

Bei Blue verwenden wir unser eigenes Produkt.

Das bedeutet, dass wir Blue nutzen, um *Blue* zu bauen.

Dieser seltsam klingende Begriff, oft als "Dogfooding" bezeichnet, wird häufig Paul Maritz, einem Manager bei Microsoft in den 1980er Jahren, zugeschrieben. Berichten zufolge hat er eine E-Mail mit dem Betreff *"Eating our own dog food"* gesendet, um die Mitarbeiter von Microsoft zu ermutigen, die Produkte des Unternehmens zu nutzen.

Die Idee, Ihre eigenen Werkzeuge zu verwenden, um Ihre Werkzeuge zu bauen, führt zu einem positiven Feedback-Zyklus.

Die Idee, Ihre eigenen Werkzeuge zu verwenden, um Ihre Werkzeuge zu bauen, führt zu einem positiven Feedback-Zyklus, der zahlreiche Vorteile schafft:

- **Es hilft uns, reale Usability-Probleme schnell zu identifizieren.** Da wir Blue täglich nutzen, begegnen wir denselben Herausforderungen, mit denen unsere Nutzer konfrontiert sein könnten, was es uns ermöglicht, proaktiv darauf zu reagieren.
- **Es beschleunigt die Entdeckung von Bugs.** Die interne Nutzung zeigt oft Bugs, bevor sie unsere Kunden erreichen, was die Gesamtqualität des Produkts verbessert.
- **Es stärkt unser Einfühlungsvermögen für Endbenutzer.** Unser Team sammelt direkte Erfahrungen mit den Stärken und Schwächen von Blue, was uns hilft, benutzerzentrierte Entscheidungen zu treffen.
- **Es fördert eine Kultur der Qualität innerhalb unserer Organisation.** Wenn jeder das Produkt nutzt, gibt es ein gemeinsames Interesse an dessen Exzellenz.
- **Es fördert Innovation.** Regelmäßige Nutzung regt oft Ideen für neue Funktionen oder Verbesserungen an, wodurch Blue an der Spitze bleibt.

[Wir haben bereits darüber gesprochen, warum wir kein dediziertes Testteam haben](/insights/open-beta), und dies ist ein weiterer Grund.

Wenn es Bugs in unserem System gibt, finden wir sie fast immer in unserem ständigen täglichen Gebrauch der Plattform. Und das schafft auch einen Zwang, sie zu beheben, da wir sie offensichtlich als sehr störend empfinden, da wir wahrscheinlich zu den Hauptnutzern von Blue gehören!

Dieser Ansatz zeigt unser Engagement für das Produkt. Indem wir uns auf Blue verlassen, zeigen wir unseren Kunden, dass wir wirklich an das glauben, was wir aufbauen. Es ist nicht nur ein Produkt, das wir verkaufen – es ist ein Werkzeug, auf das wir jeden Tag angewiesen sind.

## Hauptprozess

Wir haben ein Projekt in Blue, treffend "Produkt" genannt.

**Alles**, was mit unserer Produktentwicklung zu tun hat, wird hier verfolgt. Kundenfeedback, Bugs, Ideen für Funktionen, laufende Arbeiten usw. Die Idee, ein Projekt zu haben, in dem wir alles verfolgen, ist, dass es [besseres Teamwork fördert.](/insights/great-teamwork)

Jeder Eintrag ist eine Funktion oder ein Teil einer Funktion. So bewegen wir uns von "Wäre es nicht cool, wenn..." zu "Schau dir diese großartige neue Funktion an!"

Das Projekt hat die folgenden Listen:

- **Ideen/Feedback**: Dies ist eine Liste von Teamideen oder Kundenfeedback basierend auf Anrufen oder E-Mail-Austausch. Fühlen Sie sich frei, hier Ideen hinzuzufügen! In dieser Liste haben wir noch nicht entschieden, dass wir eine dieser Funktionen bauen werden, aber wir überprüfen dies regelmäßig auf Ideen, die wir weiter erkunden möchten.
- **Backlog (Langfristig)**: Hier kommen Funktionen aus der Ideen-/Feedbackliste hin, wenn wir entscheiden, dass sie eine gute Ergänzung für Blue wären.
- **{Aktuelles Quartal}**: Dies ist typischerweise als "Qx YYYY" strukturiert und zeigt unsere Quartalsprioritäten.
- **Bugs**: Dies ist eine Liste bekannter Bugs, die vom Team oder von Kunden gemeldet wurden. Hier hinzugefügte Bugs erhalten automatisch das Tag "Bug".
- **Spezifikationen**: Diese Funktionen werden derzeit spezifiziert. Nicht jede Funktion benötigt eine Spezifikation oder ein Design; es hängt von der erwarteten Größe der Funktion und dem Vertrauensniveau ab, das wir hinsichtlich Randfällen und Komplexität haben.
- **Design-Backlog**: Dies ist das Backlog für die Designer. Jedes Mal, wenn sie etwas abgeschlossen haben, das in Arbeit ist, können sie einen beliebigen Punkt aus dieser Liste auswählen.
- **In Arbeit Design**: Dies sind die aktuellen Funktionen, die die Designer entwerfen.
- **Design-Review**: Hier befinden sich die Funktionen, deren Designs derzeit überprüft werden.
- **Backlog (Kurzfristig)**: Dies ist eine Liste von Funktionen, an denen wir wahrscheinlich in den nächsten Wochen zu arbeiten beginnen werden. Hier finden die Zuweisungen statt. Der CEO und der Leiter der Technik entscheiden, welche Funktionen welchem Ingenieur basierend auf früheren Erfahrungen und Arbeitslast zugewiesen werden. [Teammitglieder können diese dann in den In Arbeit-Bereich ziehen](/insights/push-vs-pull-kanban), sobald sie ihre aktuelle Arbeit abgeschlossen haben.
- **In Arbeit**: Dies sind Funktionen, die derzeit entwickelt werden.
- **Code-Review**: Sobald eine Funktion die Entwicklung abgeschlossen hat, wird sie einer Code-Überprüfung unterzogen. Dann wird sie entweder zurück zu "In Arbeit" verschoben, wenn Anpassungen erforderlich sind, oder in die Entwicklungsumgebung bereitgestellt.
- **Dev**: Dies sind alle Funktionen, die sich derzeit in der Entwicklungsumgebung befinden. Andere Teammitglieder und bestimmte Kunden können diese überprüfen.
- **Beta**: Dies sind alle Funktionen, die sich derzeit in der [Beta-Umgebung](https://beta.app.blue.cc) befinden. Viele Kunden nutzen dies als ihre tägliche Blue-Plattform und geben ebenfalls Feedback.
- **Produktion**: Wenn eine Funktion die Produktion erreicht, wird sie als abgeschlossen betrachtet.

Manchmal stellen wir während der Entwicklung einer Funktion fest, dass bestimmte Unterfunktionen schwieriger zu implementieren sind als zunächst erwartet, und wir entscheiden uns möglicherweise, diese in der ersten Version, die wir an Kunden bereitstellen, nicht zu machen. In diesem Fall können wir einen neuen Eintrag mit einem Namen im Format "{FeatureName} V2" erstellen und alle Unterfunktionen als Checklistenpunkte hinzufügen.

## Tags

- **Mobile**: Dies bedeutet, dass die Funktion spezifisch für unsere iOS-, Android- oder iPad-Apps ist.
- **{EnterpriseCustomerName}**: Eine Funktion wird speziell für einen Unternehmenskunden entwickelt. Das Tracking ist wichtig, da es typischerweise zusätzliche kommerzielle Vereinbarungen für jede Funktion gibt.
- **Bug**: Dies bedeutet, dass es sich um einen Bug handelt, der behoben werden muss.
- **Fast-Track**: Dies bedeutet, dass dies eine Fast-Track-Änderung ist, die nicht den vollständigen Release-Zyklus durchlaufen muss, wie oben beschrieben.
- **Main**: Dies ist eine bedeutende Funktionentwicklung. Sie ist typischerweise für größere Infrastrukturarbeiten, große Abhängigkeitsaktualisierungen und bedeutende neue Module innerhalb von Blue reserviert.
- **AI**: Diese Funktion enthält eine Komponente der künstlichen Intelligenz.
- **Sicherheit**: Dies bedeutet, dass eine Sicherheitsauswirkung überprüft werden muss oder ein Patch erforderlich ist.

Das Fast-Track-Tag ist besonders interessant. Dies ist für kleinere, weniger komplexe Updates reserviert, die nicht unseren vollständigen Release-Zyklus erfordern und die wir innerhalb von 24-48 Stunden an Kunden ausliefern möchten.

Fast-Track-Änderungen sind typischerweise kleinere Anpassungen, die die Benutzererfahrung erheblich verbessern können, ohne die Kernfunktionalität zu verändern. Denken Sie an das Beheben von Tippfehlern in der Benutzeroberfläche, das Anpassen von Button-Padding oder das Hinzufügen neuer Icons für eine bessere visuelle Anleitung. Dies sind die Art von Änderungen, die, obwohl sie klein sind, einen großen Unterschied darin machen können, wie Benutzer unser Produkt wahrnehmen und damit interagieren. Sie sind auch ärgerlich, wenn sie lange dauern, um ausgeliefert zu werden!

Unser Fast-Track-Prozess ist unkompliziert.

Er beginnt mit der Erstellung eines neuen Branches von main, der Implementierung der Änderungen und dann der Erstellung von Merge-Anfragen für jeden Ziel-Branch - Dev, Beta und Produktion. Wir generieren einen Vorschau-Link zur Überprüfung, um sicherzustellen, dass selbst diese kleinen Änderungen unseren Qualitätsstandards entsprechen. Nach Genehmigung werden die Änderungen gleichzeitig in alle Branches zusammengeführt, um unsere Umgebungen synchron zu halten.

## Benutzerdefinierte Felder

Wir haben nicht viele benutzerdefinierte Felder in unserem Produktprojekt.

- **Spezifikationen**: Dies verweist auf ein Blue-Dokument, das die Spezifikation für diese bestimmte Funktion enthält. Dies wird nicht immer gemacht, da es von der Komplexität der Funktion abhängt.
- **MR**: Dies ist der Link zur Merge-Anfrage in [Gitlab](https://gitlab.com), wo wir unseren Code hosten.
- **Vorschau-Link**: Für Funktionen, die hauptsächlich das Frontend ändern, können wir eine eindeutige URL erstellen, die diese Änderungen für jedes Commit enthält, sodass wir die Änderungen leicht überprüfen können.
- **Lead**: Dieses Feld sagt uns, welcher leitende Ingenieur für die Code-Überprüfung zuständig ist. Es stellt sicher, dass jede Funktion die Expertenaufmerksamkeit erhält, die sie verdient, und es gibt immer eine klare Ansprechperson für Fragen oder Bedenken.

## Checklisten

Während unserer wöchentlichen Demos werden wir das besprochene Feedback in einer Checkliste namens "Feedback" festhalten, und es wird auch eine weitere Checkliste geben, die den Haupt-[WBS (Work Breakdown Scope)](/insights/simple-work-breakdown-structure) der Funktion enthält, sodass wir leicht erkennen können, was erledigt ist und was noch zu tun bleibt.

## Fazit

Und das war's!

Wir denken, dass die Leute manchmal überrascht sind, wie unkompliziert unser Prozess ist, aber wir glauben, dass einfache Prozesse oft weit überlegen sind gegenüber übermäßig komplexen Prozessen, die man nicht leicht verstehen kann.

Diese Einfachheit ist absichtlich. Sie ermöglicht es uns, agil zu bleiben, schnell auf Kundenbedürfnisse zu reagieren und unser gesamtes Team aufeinander abzustimmen.

Indem wir Blue nutzen, um Blue zu bauen, entwickeln wir nicht nur ein Produkt – wir leben es.

Also denken Sie das nächste Mal, wenn Sie Blue verwenden: Sie nutzen nicht nur ein Produkt, das wir gebaut haben. Sie nutzen ein Produkt, auf das wir persönlich jeden einzelnen Tag angewiesen sind.

Und das macht den ganzen Unterschied.
---
title: Warum Blue eine offene Beta hat
category: "Engineering"
description: Erfahren Sie, warum unser Projektmanagementsystem eine laufende offene Beta hat.
date: 2024-08-03
---

Viele B2B SaaS-Startups starten in der Beta-Phase, und das aus gutem Grund. Es ist Teil des traditionellen Silicon Valley-Mottos *„schnell handeln und Dinge kaputt machen“*.

Ein „Beta“-Aufkleber auf einem Produkt reduziert die Erwartungen.

Etwas ist kaputt? Na ja, es ist nur eine Beta.

Das System ist langsam? Na ja, es ist nur eine Beta.

[Die Dokumentation](https://blue.cc/docs) existiert nicht? Na ja… Sie verstehen schon.

Und das ist *tatsächlich* eine gute Sache. Reid Hoffman, der Gründer von LinkedIn, sagte einmal:

> Wenn Sie sich nicht für die erste Version Ihres Produkts schämen, haben Sie zu spät gestartet.

Und der Beta-Aufkleber ist auch gut für die Kunden. Er hilft ihnen, sich selbst auszuwählen.

Die Kunden, die Beta-Produkte ausprobieren, befinden sich in den frühen Phasen des Technologieakzeptanzzyklus, auch bekannt als die Produktakzeptanzkurve.

Der Technologieakzeptanzzyklus wird typischerweise in fünf Hauptsegmente unterteilt:

1. Innovatoren
2. Frühe Anwender
3. Frühe Mehrheit
4. Späte Mehrheit
5. Nachzügler

![](/insights/technology-adoption-lifecycle-graph.png)

Letztendlich muss das Produkt jedoch reifen, und die Kunden erwarten ein stabiles, funktionierendes Produkt. Sie möchten keinen Zugang zu einer „Beta“-Umgebung, in der Dinge kaputtgehen.

Oder etwa doch?

*Das* ist die Frage, die wir uns gestellt haben.

Wir glauben, dass wir uns diese Frage aufgrund der Art und Weise gestellt haben, wie Blue ursprünglich aufgebaut wurde. [Blue begann als Ableger einer geschäftigen Designagentur](/insights/agency-success-playbook), und so arbeiteten wir *innerhalb* des Büros eines Unternehmens, das Blue aktiv zur Verwaltung aller ihrer Projekte nutzte.

Das bedeutet, dass wir über Jahre hinweg beobachten konnten, wie *echte* Menschen — direkt neben uns! — Blue in ihrem täglichen Leben nutzten.

Und da sie Blue von den frühen Tagen an verwendeten, nutzte dieses Team immer Blue Beta!

Und so war es für uns natürlich, auch all unseren anderen Kunden die Nutzung zu ermöglichen.

**Und genau deshalb haben wir kein spezielles Testteam.**

Das stimmt.

Niemand bei Blue hat die *alleinige* Verantwortung dafür, dass unsere Plattform gut und stabil läuft.

Das hat mehrere Gründe.

Der erste ist eine niedrigere Kostenbasis.

Kein Vollzeit-Testteam zu haben, senkt unsere Kosten erheblich, und wir können diese Einsparungen an unsere Kunden mit den niedrigsten Preisen der Branche weitergeben.

Um das ins rechte Licht zu rücken: Wir bieten Unternehmensfunktionen an, für die unsere Konkurrenz 30-55 $/Nutzer/Monat verlangt, für nur 7 $/Monat.

Das geschieht nicht zufällig, es ist *absichtlich*.

Es ist jedoch keine gute Strategie, ein günstigeres Produkt zu verkaufen, wenn es nicht funktioniert.

Die *eigentliche Frage ist*, wie schaffen wir es, eine stabile Plattform zu schaffen, die Tausende von Kunden nutzen können, ohne ein spezielles Testteam?

Natürlich ist unser Ansatz, eine offene Beta zu haben, entscheidend dafür, aber bevor wir darauf eingehen, möchten wir die Verantwortung der Entwickler ansprechen.

Wir haben bei Blue früh entschieden, dass wir die Verantwortlichkeiten für Front-End- und Back-End-Technologien niemals aufteilen würden. Wir würden nur Full-Stack-Entwickler einstellen oder ausbilden.

Der Grund für diese Entscheidung war, sicherzustellen, dass ein Entwickler die Funktion, an der er arbeitet, vollständig besitzt. So gibt es keine *„das Problem über den Gartenzaun werfen“* Mentalität, die manchmal auftritt, wenn es gemeinsame Verantwortlichkeiten für Funktionen gibt.

Und das erstreckt sich auf das Testen der Funktion, das Verständnis der Kundenanwendungsfälle und -anfragen sowie das Lesen und Kommentieren der Spezifikationen.

Mit anderen Worten, jeder Entwickler entwickelt ein tiefes und intuitives Verständnis für die Funktion, die er erstellt.

Okay, lassen Sie uns jetzt über unsere offene Beta sprechen.

Wenn wir sagen, sie ist „offen“ — dann meinen wir das. Jeder Kunde kann sie einfach ausprobieren, indem er „beta“ vor unserer Webanwendungs-URL hinzufügt.

So wird „app.blue.cc“ zu „beta.app.blue.cc“.

Wenn sie dies tun, können sie ihre üblichen Daten sehen, da sowohl die Beta- als auch die Produktionsumgebungen dieselbe Datenbank teilen, aber sie werden auch in der Lage sein, neue Funktionen zu sehen.

Kunden können problemlos arbeiten, auch wenn einige Teammitglieder in der Produktion und einige Neugierige in der Beta sind.

Typischerweise haben wir zu jedem Zeitpunkt einige Hundert Kunden, die die Beta nutzen, und wir veröffentlichen Funktionsvorschauen in unseren Community-Foren, damit sie sehen können, was neu ist und es ausprobieren können.

Und das ist der Punkt: Wir haben *mehrere Hundert* Tester!

All diese Kunden werden Funktionen in ihren Arbeitsabläufen ausprobieren und ziemlich lautstark sein, wenn etwas nicht ganz richtig ist, weil sie die Funktion *bereits* in ihrem Unternehmen implementieren!

Das häufigste Feedback sind kleine, aber sehr nützliche Änderungen, die Randfälle ansprechen, die wir nicht berücksichtigt haben.

Wir lassen neue Funktionen zwischen 2-4 Wochen in der Beta. Wann immer wir das Gefühl haben, dass sie stabil sind, geben wir sie in die Produktion frei.

Wir haben auch die Möglichkeit, die Beta bei Bedarf zu umgehen, indem wir ein Fast-Track-Flag verwenden. Dies geschieht typischerweise für Fehlerbehebungen, die wir nicht 2-4 Wochen zurückhalten möchten, bevor wir sie in die Produktion bringen.

Das Ergebnis?

Das Pushen in die Produktion fühlt sich… nun ja, langweilig an! Wie nichts — es ist einfach kein großes Ding für uns.

Und es bedeutet, dass dies unseren Veröffentlichungszeitplan glättet, was es uns ermöglicht hat, [seit sechs Jahren monatlich wie am Schnürchen Funktionen auszuliefern.](/changelog).

Wie bei jeder Entscheidung gibt es jedoch einige Kompromisse.

Der Kundensupport ist etwas komplexer, da wir Kunden über zwei Versionen unserer Plattform unterstützen müssen. Manchmal kann dies zu Verwirrung bei Kunden führen, die Teammitglieder haben, die zwei verschiedene Versionen verwenden.

Ein weiterer Schmerzpunkt ist, dass dieser Ansatz manchmal den gesamten Veröffentlichungszeitplan für die Produktion verlangsamen kann. Dies gilt insbesondere für größere Funktionen, die in der Beta „stecken bleiben“ können, wenn es ein anderes verwandtes Feature gibt, das Probleme hat und weitere Arbeiten benötigt.

Aber insgesamt glauben wir, dass diese Kompromisse die Vorteile einer niedrigeren Kostenbasis und einer höheren Kundenbindung wert sind.

Wir gehören zu den wenigen Softwareunternehmen, die diesen Ansatz verfolgen, aber er ist mittlerweile ein grundlegender Bestandteil unseres Produktentwicklungsansatzes.
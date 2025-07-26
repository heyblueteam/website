---
title: Skalierung von CSV-Importen und -Exporten auf über 250.000 Datensätze
category: "Engineering"
description: Erfahren Sie, wie Blue CSV-Importe und -Exporte mit Rust und skalierbarer Architektur sowie strategischen Technologieentscheidungen im B2B SaaS um das 10-fache skaliert hat.
date: 2024-07-18
---
Bei Blue [verschieben wir ständig die Grenzen](/platform/roadmap) dessen, was in der Projektmanagement-Software möglich ist. Im Laufe der Jahre haben wir [hunderte von Funktionen veröffentlicht](/platform/changelog).

Unser neuester Ingenieurkunststück? 

Eine vollständige Überarbeitung unseres [CSV-Import](https://documentation.blue.cc/integrations/csv-import) und [Export](https://documentation.blue.cc/integrations/csv-export) Systems, die die Leistung und Skalierbarkeit dramatisch verbessert hat. 

Dieser Beitrag gibt Ihnen einen Einblick, wie wir diese Herausforderung angegangen sind, welche Technologien wir eingesetzt haben und welche beeindruckenden Ergebnisse wir erzielt haben.

Das Interessanteste daran ist, dass wir über unseren typischen [Technologiestack](https://sop.blue.cc/product/technology-stack) hinausgehen mussten, um die gewünschten Ergebnisse zu erzielen. Dies ist eine Entscheidung, die sorgfältig getroffen werden muss, da die langfristigen Auswirkungen in Bezug auf technische Schulden und langfristige Wartungskosten erheblich sein können.

<video autoplay loop muted playsinline>
  <source src="/videos/import-export-video.mp4" type="video/mp4">
</video>

## Skalierung für Unternehmensbedürfnisse

Unsere Reise begann mit einer Anfrage von einem Unternehmenskunden aus der Veranstaltungsbranche. Dieser Kunde nutzt Blue als zentrales Hub zur Verwaltung umfangreicher Listen von Veranstaltungen, Veranstaltungsorten und Rednern und integriert es nahtlos in seine Website. 

Für sie ist Blue nicht nur ein Werkzeug – es ist die einzige Quelle der Wahrheit für ihren gesamten Betrieb.

Während wir immer stolz sind zu hören, dass Kunden uns für solch geschäftskritische Bedürfnisse nutzen, liegt auch eine große Verantwortung auf unserer Seite, ein schnelles, zuverlässiges System zu gewährleisten.

Als dieser Kunde seine Operationen skalierte, sah er sich mit einem erheblichen Hindernis konfrontiert: **Importieren und Exportieren großer CSV-Dateien mit 100.000 bis 200.000+ Datensätzen.**

Das war zu diesem Zeitpunkt über die Fähigkeiten unseres Systems hinaus. Tatsächlich hatte unser vorheriges Import-/Export-System bereits Schwierigkeiten mit Importen und Exporten, die mehr als 10.000 bis 20.000 Datensätze enthielten! Daher waren 200.000+ Datensätze ausgeschlossen. 

Benutzer erlebten frustrierend lange Wartezeiten, und in einigen Fällen würden Importe oder Exporte *ganz ausbleiben.* Dies hatte erhebliche Auswirkungen auf ihre Abläufe, da sie auf tägliche Importe und Exporte angewiesen waren, um bestimmte Aspekte ihrer Operationen zu verwalten. 

> Multi-Tenancy ist eine Architektur, bei der eine einzelne Instanz von Software mehreren Kunden (Mandanten) dient. Während sie effizient ist, erfordert sie eine sorgfältige Ressourcenverwaltung, um sicherzustellen, dass die Aktionen eines Mandanten andere nicht negativ beeinflussen.

Und diese Einschränkung betraf nicht nur diesen speziellen Kunden. 

Aufgrund unserer Multi-Tenant-Architektur – bei der mehrere Kunden dieselbe Infrastruktur teilen – könnte ein einzelner ressourcenintensiver Import oder Export potenziell die Operationen anderer Benutzer verlangsamen, was in der Praxis oft geschah. 

Wie üblich führten wir eine Build-vs-Buy-Analyse durch, um zu verstehen, ob wir die Zeit investieren sollten, um unser eigenes System zu verbessern oder ein System von jemand anderem zu kaufen. Wir betrachteten verschiedene Möglichkeiten.

Der Anbieter, der herausstach, war ein SaaS-Anbieter namens [Flatfile](https://flatfile.com/). Ihr System und ihre Fähigkeiten schienen genau das zu sein, was wir benötigten. 

Aber nach der Überprüfung ihrer [Preise](https://flatfile.com/pricing/) entschieden wir, dass dies eine extrem teure Lösung für eine Anwendung unserer Größenordnung werden würde – *$2/pro Datei summiert sich wirklich schnell!* – und es besser wäre, unsere integrierte CSV-Import-/Export-Engine zu erweitern. 

Um diese Herausforderung anzugehen, trafen wir eine mutige Entscheidung: Rust in unseren primären Javascript-Technologiestack einzuführen. Diese Systemprogrammiersprache, bekannt für ihre Leistung und Sicherheit, war das perfekte Werkzeug für unsere leistungskritischen CSV-Parsing- und Datenmapping-Bedürfnisse.

So gingen wir an die Lösung heran.

### Einführung von Hintergrunddiensten

Die Grundlage unserer Lösung war die Einführung von Hintergrunddiensten zur Bearbeitung ressourcenintensiver Aufgaben. Dieser Ansatz ermöglichte es uns, schwere Verarbeitungen von unserem Hauptserver auszulagern, was die Gesamtleistung des Systems erheblich verbesserte.
Unsere Architektur der Hintergrunddienste ist mit Blick auf Skalierbarkeit konzipiert. Wie alle Komponenten unserer Infrastruktur skalieren diese Dienste automatisch basierend auf der Nachfrage. 

Das bedeutet, dass während der Spitzenzeiten, wenn mehrere große Importe oder Exporte gleichzeitig verarbeitet werden, das System automatisch mehr Ressourcen zuweist, um die erhöhte Last zu bewältigen. Umgekehrt wird es in ruhigeren Zeiten heruntergefahren, um die Ressourcennutzung zu optimieren.

Diese skalierbare Architektur der Hintergrunddienste hat Blue nicht nur für CSV-Importe und -Exporte zugutegekommen. Im Laufe der Zeit haben wir eine beträchtliche Anzahl von Funktionen in Hintergrunddienste verlagert, um die Last von unseren Hauptservern zu nehmen:

- **[Formelberechnungen](https://documentation.blue.cc/custom-fields/formula)**: Lagert komplexe mathematische Operationen aus, um schnelle Aktualisierungen abgeleiteter Felder zu gewährleisten, ohne die Leistung des Hauptservers zu beeinträchtigen.
- **[Dashboard/Diagramme](/platform/features/dashboards)**: Verarbeitet große Datensätze im Hintergrund, um aktuelle Visualisierungen zu erstellen, ohne die Benutzeroberfläche zu verlangsamen.
- **[Suchindex](https://documentation.blue.cc/projects/search)**: Aktualisiert kontinuierlich den Suchindex im Hintergrund und sorgt für schnelle und genaue Suchergebnisse, ohne die Systemleistung zu beeinträchtigen.
- **[Projekte kopieren](https://documentation.blue.cc/projects/copying-projects)**: Bearbeitet die Replikation großer, komplexer Projekte im Hintergrund, sodass Benutzer weiterarbeiten können, während die Kopie erstellt wird.
- **[Projektmanagement-Automatisierungen](/platform/features/automations)**: Führt benutzerdefinierte automatisierte Workflows im Hintergrund aus und gewährleistet rechtzeitige Aktionen, ohne andere Operationen zu blockieren.
- **[Wiederholende Datensätze](https://documentation.blue.cc/records/repeat)**: Generiert wiederkehrende Aufgaben oder Ereignisse im Hintergrund und sorgt für die Genauigkeit des Zeitplans, ohne die Hauptanwendung zu belasten.
- **[Zeitdauer benutzerdefinierte Felder](https://documentation.blue.cc/custom-fields/duration)**: Berechnet und aktualisiert kontinuierlich die Zeitdifferenz zwischen zwei Ereignissen in Blue und liefert Echtzeit-Dauerinformationen, ohne die Reaktionsfähigkeit des Systems zu beeinträchtigen.

## Neues Rust-Modul für die Datenverarbeitung

Das Herzstück unserer CSV-Verarbeitungslösung ist ein benutzerdefiniertes Rust-Modul. Während dies unser erster Schritt außerhalb unseres Kerntechnologiestacks von Javascript war, wurde die Entscheidung, Rust zu verwenden, durch seine außergewöhnliche Leistung bei gleichzeitigen Operationen und Datei-Verarbeitungsaufgaben motiviert.

Rusts Stärken passen perfekt zu den Anforderungen des CSV-Parsings und des Datenmappings. Seine Nullkosten-Abstraktionen ermöglichen hochgradiges Programmieren, ohne die Leistung zu opfern, während sein Ownership-Modell die Speichersicherheit ohne die Notwendigkeit von Garbage Collection gewährleistet. Diese Eigenschaften machen Rust besonders geeignet für die effiziente und sichere Verarbeitung großer Datensätze.

Für das CSV-Parsing haben wir Rusts csv crate genutzt, das eine hochleistungsfähige Lese- und Schreibfunktion für CSV-Daten bietet. Wir kombinierten dies mit benutzerdefinierter Datenmapping-Logik, um eine nahtlose Integration mit Blues Datenstrukturen sicherzustellen.

Die Lernkurve für Rust war steil, aber machbar. Unser Team widmete etwa zwei Wochen intensives Lernen dafür.

Die Verbesserungen waren beeindruckend:

![](/insights/import-export.png)

Unser neues System kann die gleiche Anzahl von Datensätzen verarbeiten, die unser altes System in 15 Minuten verarbeiten konnte, in etwa 30 Sekunden. 

## Interaktion zwischen Webserver und Datenbank

Für die Webserver-Komponente unserer Rust-Implementierung wählten wir Rocket als unser Framework. Rocket stach durch seine Kombination aus Leistung und entwicklerfreundlichen Funktionen hervor. Seine statische Typisierung und die Überprüfung zur Compile-Zeit passen gut zu Rusts Sicherheitsprinzipien und helfen uns, potenzielle Probleme früh im Entwicklungsprozess zu erkennen.
Auf der Datenbankseite entschieden wir uns für SQLx. Diese asynchrone SQL-Bibliothek für Rust bietet mehrere Vorteile, die sie ideal für unsere Bedürfnisse machen:

- Typ-sicheres SQL: SQLx ermöglicht es uns, rohes SQL mit zur Compile-Zeit überprüften Abfragen zu schreiben, was Typensicherheit gewährleistet, ohne die Leistung zu beeinträchtigen.
- Async-Unterstützung: Dies passt gut zu Rocket und unserem Bedarf an effizienten, nicht-blockierenden Datenbankoperationen.
- Datenbankunabhängig: Während wir hauptsächlich [AWS Aurora](https://aws.amazon.com/rds/aurora/) verwenden, das MySQL-kompatibel ist, bietet SQLx Unterstützung für mehrere Datenbanken, was uns Flexibilität für die Zukunft gibt, falls wir uns jemals entscheiden sollten, zu wechseln. 

## Optimierung des Batchings

Unser Weg zur optimalen Batching-Konfiguration war von rigorosen Tests und sorgfältiger Analyse geprägt. Wir führten umfangreiche Benchmarks mit verschiedenen Kombinationen von gleichzeitigen Transaktionen und Chunk-Größen durch und maßen nicht nur die rohe Geschwindigkeit, sondern auch die Ressourcennutzung und die Systemstabilität.

Der Prozess umfasste die Erstellung von Testdatensätzen unterschiedlicher Größe und Komplexität, die reale Nutzungsmuster simulierten. Wir ließen diese Datensätze dann durch unser System laufen und passten die Anzahl der gleichzeitigen Transaktionen und die Chunk-Größe für jeden Durchlauf an.

Nach der Analyse der Ergebnisse fanden wir heraus, dass die Verarbeitung von 5 gleichzeitigen Transaktionen mit einer Chunk-Größe von 500 Datensätzen das beste Gleichgewicht zwischen Geschwindigkeit und Ressourcennutzung bot. Diese Konfiguration ermöglicht es uns, eine hohe Durchsatzrate aufrechtzuerhalten, ohne unsere Datenbank zu überlasten oder übermäßigen Speicher zu verbrauchen.

Interessanterweise fanden wir heraus, dass eine Erhöhung der Gleichzeitigkeit über 5 Transaktionen keine signifikanten Leistungsgewinne brachte und manchmal zu einer erhöhten Datenbankkonkurrenz führte. Ähnlich verbesserten größere Chunk-Größen die rohe Geschwindigkeit, jedoch auf Kosten eines höheren Speicherverbrauchs und längerer Reaktionszeiten bei kleinen bis mittelgroßen Importen/Exporten.

## CSV-Exporte über E-Mail-Links

Das letzte Puzzlestück unserer Lösung befasst sich mit der Herausforderung, große exportierte Dateien an Benutzer zu liefern. Anstatt einen direkten Download von unserer Webanwendung anzubieten, der zu Timeout-Problemen und einer erhöhten Serverlast führen könnte, implementierten wir ein System von per E-Mail versendeten Download-Links.

Wenn ein Benutzer einen großen Export initiiert, verarbeitet unser System die Anfrage im Hintergrund. Sobald der Export abgeschlossen ist, laden wir die Datei an einen sicheren, temporären Speicherort hoch, anstatt die Verbindung offen zu halten oder die Datei auf unseren Webservern zu speichern. Anschließend generieren wir einen einzigartigen, sicheren Download-Link und senden ihn per E-Mail an den Benutzer.

Diese Download-Links sind 2 Stunden lang gültig und bieten ein Gleichgewicht zwischen Benutzerfreundlichkeit und Informationssicherheit. Dieser Zeitraum gibt den Benutzern ausreichend Gelegenheit, ihre Daten abzurufen, während sichergestellt wird, dass sensible Informationen nicht unbegrenzt zugänglich bleiben.

Die Sicherheit dieser Download-Links hatte in unserem Design oberste Priorität. Jeder Link ist:

- Einzigartig und zufällig generiert, was es praktisch unmöglich macht, ihn zu erraten
- Nur 2 Stunden gültig
- Verschlüsselt während der Übertragung, um die Sicherheit der Daten beim Herunterladen zu gewährleisten

Dieser Ansatz bietet mehrere Vorteile:

- Er reduziert die Last auf unseren Webservern, da sie keine großen Datei-Downloads direkt abwickeln müssen
- Er verbessert die Benutzererfahrung, insbesondere für Benutzer mit langsameren Internetverbindungen, die möglicherweise auf Timeout-Probleme mit direkten Downloads stoßen
- Er bietet eine zuverlässigere Lösung für sehr große Exporte, die die typischen Web-Timeout-Grenzen überschreiten könnten

Das Benutzerfeedback zu dieser Funktion war überwältigend positiv, viele schätzen die Flexibilität, die sie bei der Verwaltung großer Datenexporte bietet.

## Exportieren gefilterter Daten

Die andere offensichtliche Verbesserung bestand darin, den Benutzern zu ermöglichen, nur die Daten zu exportieren, die bereits in ihrer Projektansicht gefiltert waren. Das bedeutet, wenn es ein aktives Tag "Priorität" gibt, würden nur Datensätze, die dieses Tag haben, im CSV-Export landen. Das bedeutet weniger Zeit, um Daten in Excel zu manipulieren, um Dinge herauszufiltern, die nicht wichtig sind, und hilft uns auch, die Anzahl der zu verarbeitenden Zeilen zu reduzieren.

## Ausblick

Während wir keine unmittelbaren Pläne haben, unseren Einsatz von Rust zu erweitern, hat uns dieses Projekt das Potenzial dieser Technologie für leistungskritische Operationen aufgezeigt. Es ist eine aufregende Option, die wir jetzt in unserem Werkzeugkasten für zukünftige Optimierungsbedürfnisse haben. Diese Überarbeitung des CSV-Imports und -Exports passt perfekt zu Blues Engagement für Skalierbarkeit. 

Wir sind bestrebt, eine Plattform bereitzustellen, die mit unseren Kunden wächst und ihre sich erweiternden Datenbedürfnisse ohne Kompromisse bei der Leistung bewältigt.

Die Entscheidung, Rust in unseren Technologiestack einzuführen, wurde nicht leichtfertig getroffen. Sie stellte eine wichtige Frage, mit der viele Ingenieurteams konfrontiert sind: Wann ist es angemessen, über den eigenen Technologiestack hinauszugehen, und wann sollte man bei vertrauten Werkzeugen bleiben?

Es gibt keine universelle Antwort, aber bei Blue haben wir einen Rahmen entwickelt, um diese entscheidenden Entscheidungen zu treffen:

- **Problem-First-Ansatz:** Wir beginnen immer damit, das Problem klar zu definieren, das wir lösen möchten. In diesem Fall mussten wir die Leistung von CSV-Importen und -Exporten für große Datensätze dramatisch verbessern.
- **Ausschöpfung bestehender Lösungen:** Bevor wir außerhalb unseres Kernstacks suchen, erkunden wir gründlich, was mit unseren bestehenden Technologien erreicht werden kann. Dies umfasst oft Profiling, Optimierung und das Überdenken unseres Ansatzes innerhalb vertrauter Einschränkungen.
- **Quantifizierung des potenziellen Gewinns:** Wenn wir eine neue Technologie in Betracht ziehen, müssen wir in der Lage sein, die Vorteile klar zu artikulieren und idealerweise zu quantifizieren. Für unser CSV-Projekt haben wir Verbesserungen in der Verarbeitungszeit in der Größenordnung prognostiziert.
- **Bewertung der Kosten:** Die Einführung einer neuen Technologie betrifft nicht nur das unmittelbare Projekt. Wir berücksichtigen die langfristigen Kosten:
  - Lernkurve für das Team
  - Laufende Wartung und Unterstützung
  - Mögliche Komplikationen bei Bereitstellung und Betrieb
  - Auswirkungen auf Einstellung und Teamzusammensetzung
- **Eingrenzung und Integration:** Wenn wir eine neue Technologie einführen, streben wir an, sie auf einen spezifischen, gut definierten Teil unseres Systems zu beschränken. Wir stellen auch sicher, dass wir einen klaren Plan haben, wie sie in unseren bestehenden Stack integriert wird.
- **Zukunftssicherung:** Wir prüfen, ob diese Technologieentscheidung zukünftige Möglichkeiten eröffnet oder ob sie uns in eine Ecke drängt.

Eines der Hauptprobleme bei der häufigen Einführung neuer Technologien ist, dass man in das gerät, was wir als *"Technologie-Zoo"* bezeichnen – ein fragmentiertes Ökosystem, in dem verschiedene Teile Ihrer Anwendung in unterschiedlichen Sprachen oder Frameworks geschrieben sind und eine breite Palette spezialisierter Fähigkeiten zur Wartung erfordern.

## Fazit

Dieses Projekt exemplifiziert Blues Ansatz zur Ingenieurkunst: *Wir scheuen uns nicht, unsere Komfortzone zu verlassen und neue Technologien zu übernehmen, wenn es bedeutet, unseren Benutzern ein erheblich besseres Erlebnis zu bieten.* 

Durch die Neugestaltung unseres CSV-Import- und Exportprozesses haben wir nicht nur ein dringendes Bedürfnis eines Unternehmenskunden gelöst, sondern auch die Erfahrung für alle unsere Benutzer verbessert, die mit großen Datensätzen arbeiten.

Während wir weiterhin die Grenzen dessen, was in [Projektmanagement-Software](/solutions/use-case/project-management) möglich ist, verschieben, freuen wir uns darauf, weitere Herausforderungen wie diese anzugehen. 

Bleiben Sie dran für weitere [tiefe Einblicke in die Technik, die Blue antreibt!](/insights/engineering-blog)
---
title: AI Auto-Kategorisierung (Engineering Deep Dive)
category: "Engineering"
description: Werfen Sie einen Blick hinter die Kulissen des Blue Engineering-Teams, das erklärt, wie sie eine KI-gestützte Auto-Kategorisierungs- und Tagging-Funktion entwickelt haben.
date: 2024-12-07
---

Wir haben kürzlich die [AI Auto-Kategorisierung](/insights/ai-auto-categorization) für alle Blue-Nutzer freigegeben. Dies ist eine KI-Funktion, die ohne zusätzliche Kosten im Blue-Kernabonnement enthalten ist. In diesem Beitrag gehen wir auf die technischen Aspekte ein, die diese Funktion möglich gemacht haben.

---
Bei Blue basiert unser Ansatz für die Funktionsentwicklung auf einem tiefen Verständnis der Nutzerbedürfnisse und Markttrends, gepaart mit dem Engagement, die Einfachheit und Benutzerfreundlichkeit beizubehalten, die unsere Plattform auszeichnen. Dies treibt unsere [Roadmap](/platform/roadmap) an und hat es uns ermöglicht, [über Jahre hinweg konsequent jeden Monat neue Funktionen zu liefern](/platform/changelog).

Die Einführung des KI-gestützten Auto-Taggings in Blue ist ein perfektes Beispiel für diese Philosophie in Aktion. Bevor wir in die technischen Details eintauchen, wie wir diese Funktion entwickelt haben, ist es entscheidend, das Problem zu verstehen, das wir lösen wollten, und die sorgfältigen Überlegungen, die in die Entwicklung eingeflossen sind.

Die Projektmanagement-Landschaft entwickelt sich rasant weiter, wobei KI-Fähigkeiten zunehmend zentral für die Nutzererwartungen werden. Unsere Kunden, insbesondere diejenigen, die groß angelegte [Projekte](/platform) mit Millionen von [Datensätzen](/platform/features/records) verwalten, hatten deutlich ihren Wunsch nach intelligenteren, effizienteren Wegen zur Organisation und Kategorisierung ihrer Daten geäußert.

Bei Blue fügen wir jedoch nicht einfach Funktionen hinzu, nur weil sie im Trend liegen oder angefragt werden. Unsere Philosophie besagt, dass jede neue Ergänzung ihren Wert beweisen muss, wobei die Standardantwort ein klares *"Nein"* ist, bis eine Funktion eine starke Nachfrage und einen klaren Nutzen nachweist.

Um die Tiefe des Problems und das Potenzial des KI-Auto-Taggings wirklich zu verstehen, führten wir umfangreiche Kundeninterviews durch, wobei wir uns auf langjährige Nutzer konzentrierten, die komplexe, datenreiche Projekte über mehrere Bereiche hinweg verwalten.

Diese Gespräche offenbarten einen gemeinsamen roten Faden: *Während das Tagging für die Organisation und Durchsuchbarkeit von unschätzbarem Wert war, wurde die manuelle Natur des Prozesses zu einem Engpass, insbesondere für Teams, die mit großen Datenmengen arbeiten.*

Aber wir sahen über die bloße Lösung des unmittelbaren Schmerzpunktes des manuellen Taggings hinaus.

Wir stellten uns eine Zukunft vor, in der KI-gestütztes Tagging zur Grundlage für intelligentere, automatisierte Arbeitsabläufe werden könnte.

Die wahre Kraft dieser Funktion, so erkannten wir, lag in ihrem Potenzial zur Integration mit unserem [Projektmanagement-Automatisierungssystem](/platform/features/automations). Stellen Sie sich ein Projektmanagement-Tool vor, das Informationen nicht nur intelligent kategorisiert, sondern diese Kategorien auch nutzt, um Aufgaben zu routen, Aktionen auszulösen und Arbeitsabläufe in Echtzeit anzupassen.

Diese Vision passte perfekt zu unserem Ziel, Blue einfach und doch leistungsstark zu halten.

Darüber hinaus erkannten wir das Potenzial, diese Fähigkeit über die Grenzen unserer Plattform hinaus zu erweitern. Durch die Entwicklung eines robusten KI-Tagging-Systems legten wir den Grundstein für eine "Kategorisierungs-API", die sofort einsatzbereit funktionieren könnte und möglicherweise neue Wege eröffnet, wie unsere Nutzer mit Blue in ihren breiteren Tech-Ökosystemen interagieren und es nutzen können.

Diese Funktion ging also nicht nur darum, ein KI-Häkchen auf unsere Feature-Liste zu setzen.

Es ging darum, einen bedeutenden Schritt in Richtung einer intelligenteren, anpassungsfähigeren Projektmanagement-Plattform zu machen und dabei unserer Kernphilosophie der Einfachheit und Nutzerzentrierung treu zu bleiben.

In den folgenden Abschnitten werden wir uns mit den technischen Herausforderungen befassen, denen wir uns bei der Verwirklichung dieser Vision gegenübersahen, der Architektur, die wir zu ihrer Unterstützung entworfen haben, und den implementierten Lösungen. Wir werden auch die zukünftigen Möglichkeiten erkunden, die diese Funktion eröffnet, und zeigen, wie eine sorgfältig durchdachte Ergänzung den Weg für transformative Veränderungen im Projektmanagement ebnen kann.

---
## Das Problem

Wie oben besprochen, kann das manuelle Tagging von Projektdatensätzen zeitaufwändig und inkonsistent sein.

Wir machten uns daran, dies zu lösen, indem wir KI nutzten, um automatisch Tags basierend auf dem Inhalt der Datensätze vorzuschlagen.

Die Hauptherausforderungen waren:

1. Die Auswahl eines geeigneten KI-Modells
2. Die effiziente Verarbeitung großer Datenmengen
3. Die Gewährleistung von Datenschutz und Sicherheit
4. Die nahtlose Integration der Funktion in unsere bestehende Architektur

## Auswahl des KI-Modells

Wir evaluierten mehrere KI-Plattformen, darunter [OpenAI](https://openai.com), Open-Source-Modelle auf [HuggingFace](https://huggingface.co/) und [Replicate](https://replicate.com).

Unsere Kriterien umfassten:

- Kosteneffizienz
- Genauigkeit beim Verstehen von Kontext
- Fähigkeit zur Einhaltung spezifischer Ausgabeformate
- Datenschutzgarantien

Nach gründlichen Tests entschieden wir uns für OpenAIs [GPT-3.5 Turbo](https://platform.openai.com/docs/models/gpt-3-5-turbo). Während [GPT-4](https://softgist.com/the-ultimate-guide-to-prompt-engineering) möglicherweise marginale Verbesserungen in der Genauigkeit bietet, zeigten unsere Tests, dass die Leistung von GPT-3.5 für unsere Auto-Tagging-Anforderungen mehr als ausreichend war. Die Balance aus Kosteneffizienz und starken Kategorisierungsfähigkeiten machte GPT-3.5 zur idealen Wahl für diese Funktion.

Die höheren Kosten von GPT-4 hätten uns gezwungen, die Funktion als kostenpflichtiges Add-on anzubieten, was im Widerspruch zu unserem Ziel stand, **KI ohne zusätzliche Kosten für Endnutzer in unser Hauptprodukt zu integrieren.**

Zum Zeitpunkt unserer Implementierung betragen die Preise für GPT-3.5 Turbo:

- $0.0005 pro 1K Input-Token (oder $0.50 pro 1M Input-Token)
- $0.0015 pro 1K Output-Token (oder $1.50 pro 1M Output-Token)

Lassen Sie uns einige Annahmen über einen durchschnittlichen Datensatz in Blue treffen:

- **Titel**: ~10 Token
- **Beschreibung**: ~50 Token
- **2 Kommentare**: ~30 Token jeweils
- **5 benutzerdefinierte Felder**: ~10 Token jeweils
- **Listenname, Fälligkeitsdatum und andere Metadaten**: ~20 Token
- **System-Prompt und verfügbare Tags**: ~50 Token

Gesamte Input-Token pro Datensatz: 10 + 50 + (30 * 2) + (10 * 5) + 20 + 50 ≈ 240 Token

Für den Output nehmen wir durchschnittlich 3 vorgeschlagene Tags pro Datensatz an, was insgesamt etwa 20 Output-Token einschließlich JSON-Formatierung ergeben könnte.

Für 1 Million Datensätze:

- Input-Kosten: (240 * 1.000.000 / 1.000.000) * $0.50 = $120
- Output-Kosten: (20 * 1.000.000 / 1.000.000) * $1.50 = $30

**Gesamtkosten für das Auto-Tagging von 1 Million Datensätzen: $120 + $30 = $150**

## GPT3.5 Turbo Leistung

Kategorisierung ist eine Aufgabe, bei der große Sprachmodelle (LLMs) wie GPT-3.5 Turbo hervorragend sind, was sie besonders gut für unsere Auto-Tagging-Funktion geeignet macht. LLMs werden auf riesigen Mengen von Textdaten trainiert, was es ihnen ermöglicht, Kontext, Semantik und Beziehungen zwischen Konzepten zu verstehen. Diese breite Wissensbasis ermöglicht es ihnen, Kategorisierungsaufgaben mit hoher Genauigkeit über eine Vielzahl von Bereichen hinweg durchzuführen.

Für unseren spezifischen Anwendungsfall des Projektmanagement-Taggings zeigt GPT-3.5 Turbo mehrere Schlüsselstärken:

- **Kontextuelles Verständnis:** Kann den Gesamtkontext eines Projektdatensatzes erfassen und dabei nicht nur einzelne Wörter, sondern die Bedeutung berücksichtigen, die durch die gesamte Beschreibung, Kommentare und andere Felder vermittelt wird.
- **Flexibilität:** Kann sich an verschiedene Projekttypen und Branchen anpassen, ohne umfangreiche Neuprogrammierung zu erfordern.
- **Umgang mit Mehrdeutigkeit:** Kann mehrere Faktoren abwägen, um nuancierte Entscheidungen zu treffen.
- **Lernen aus Beispielen:** Kann neue Kategorisierungsschemata schnell verstehen und anwenden, ohne zusätzliches Training.
- **Multi-Label-Klassifikation:** Kann mehrere relevante Tags für einen einzelnen Datensatz vorschlagen, was für unsere Anforderungen entscheidend war.

GPT-3.5 Turbo zeichnete sich auch durch seine Zuverlässigkeit bei der Einhaltung unseres erforderlichen JSON-Ausgabeformats aus, was für die nahtlose Integration in unsere bestehenden Systeme *entscheidend* war. Open-Source-Modelle fügten oft zusätzliche Kommentare hinzu oder wichen vom erwarteten Format ab, was eine zusätzliche Nachbearbeitung erfordert hätte. Diese Konsistenz im Ausgabeformat war ein Schlüsselfaktor bei unserer Entscheidung, da sie unsere Implementierung erheblich vereinfachte und potenzielle Fehlerquellen reduzierte.

Die Entscheidung für GPT-3.5 Turbo mit seiner konsistenten JSON-Ausgabe ermöglichte es uns, eine unkompliziertere, zuverlässigere und wartbarere Lösung zu implementieren.

Hätten wir ein Modell mit weniger zuverlässiger Formatierung gewählt, hätten wir eine Kaskade von Komplikationen bewältigen müssen: die Notwendigkeit robuster Parsing-Logik zur Handhabung verschiedener Ausgabeformate, umfangreiche Fehlerbehandlung für inkonsistente Ausgaben, potenzielle Leistungseinbußen durch zusätzliche Verarbeitung, erhöhte Testkomplexität zur Abdeckung aller Ausgabevariationen und eine größere langfristige Wartungslast.

Parsing-Fehler könnten zu fehlerhaftem Tagging führen und die Nutzererfahrung negativ beeinflussen. Indem wir diese Fallstricke vermieden, konnten wir unsere Engineering-Bemühungen auf kritische Aspekte wie Leistungsoptimierung und Benutzeroberflächendesign konzentrieren, anstatt uns mit unvorhersehbaren KI-Ausgaben herumzuschlagen.

## Systemarchitektur

Unsere KI-Auto-Tagging-Funktion basiert auf einer robusten, skalierbaren Architektur, die darauf ausgelegt ist, große Anfragevolumen effizient zu bewältigen und gleichzeitig eine nahtlose Benutzererfahrung zu bieten. Wie bei all unseren Systemen haben wir diese Funktion so konzipiert, dass sie eine Größenordnung mehr Traffic unterstützt, als wir derzeit erleben. Dieser Ansatz, der für aktuelle Bedürfnisse scheinbar überentwickelt ist, ist eine Best Practice, die es uns ermöglicht, plötzliche Nutzungsspitzen nahtlos zu bewältigen und uns ausreichend Spielraum für Wachstum ohne größere architektonische Überarbeitungen zu geben. Andernfalls müssten wir alle unsere Systeme alle 18 Monate neu entwickeln - etwas, das wir in der Vergangenheit auf die harte Tour gelernt haben!

Lassen Sie uns die Komponenten und den Ablauf unseres Systems aufschlüsseln:

- **Benutzerinteraktion:** Der Prozess beginnt, wenn ein Benutzer die Schaltfläche "Autotag" in der Blue-Oberfläche drückt. Diese Aktion löst den Auto-Tagging-Workflow aus.
- **Blue API-Aufruf:** Die Benutzeraktion wird in einen API-Aufruf an unser Blue-Backend übersetzt. Dieser API-Endpunkt ist darauf ausgelegt, Auto-Tagging-Anfragen zu verarbeiten.
- **Warteschlangenverwaltung:** Anstatt die Anfrage sofort zu verarbeiten, was bei hoher Last zu Leistungsproblemen führen könnte, fügen wir die Tagging-Anfrage einer Warteschlange hinzu. Wir verwenden Redis für diesen Warteschlangenmechanismus, der es uns ermöglicht, die Last effektiv zu verwalten und die Skalierbarkeit des Systems sicherzustellen.
- **Hintergrunddienst:** Wir haben einen Hintergrunddienst implementiert, der die Warteschlange kontinuierlich auf neue Anfragen überwacht. Dieser Dienst ist für die Verarbeitung der in der Warteschlange befindlichen Anfragen verantwortlich.
- **OpenAI API-Integration:** Der Hintergrunddienst bereitet die erforderlichen Daten vor und macht API-Aufrufe an OpenAIs GPT-3.5-Modell. Hier findet das eigentliche KI-gestützte Tagging statt. Wir senden relevante Projektdaten und erhalten im Gegenzug vorgeschlagene Tags.
- **Ergebnisverarbeitung:** Der Hintergrunddienst verarbeitet die von OpenAI erhaltenen Ergebnisse. Dies umfasst das Parsen der KI-Antwort und die Vorbereitung der Daten für die Anwendung auf das Projekt.
- **Tag-Anwendung:** Die verarbeiteten Ergebnisse werden verwendet, um die neuen Tags auf die relevanten Elemente im Projekt anzuwenden. Dieser Schritt aktualisiert unsere Datenbank mit den von der KI vorgeschlagenen Tags.
- **Aktualisierung der Benutzeroberfläche:** Schließlich erscheinen die neuen Tags in der Projektansicht des Benutzers und schließen den Auto-Tagging-Prozess aus der Perspektive des Benutzers ab.

Diese Architektur bietet mehrere wichtige Vorteile, die sowohl die Systemleistung als auch die Benutzererfahrung verbessern. Durch die Nutzung einer Warteschlange und Hintergrundverarbeitung haben wir eine beeindruckende Skalierbarkeit erreicht, die es uns ermöglicht, zahlreiche Anfragen gleichzeitig zu verarbeiten, ohne unser System zu überlasten oder die Ratenlimits der OpenAI API zu erreichen. Die Implementierung dieser Architektur erforderte sorgfältige Überlegungen zu verschiedenen Faktoren, um optimale Leistung und Zuverlässigkeit zu gewährleisten. Für die Warteschlangenverwaltung wählten wir Redis und nutzten seine Geschwindigkeit und Zuverlässigkeit bei der Handhabung verteilter Warteschlangen.

Dieser Ansatz trägt auch zur allgemeinen Reaktionsfähigkeit der Funktion bei. Benutzer erhalten sofortiges Feedback, dass ihre Anfrage bearbeitet wird, auch wenn das eigentliche Tagging einige Zeit dauert, was ein Gefühl der Echtzeitinteraktion erzeugt. Die Fehlertoleranz der Architektur ist ein weiterer entscheidender Vorteil. Wenn ein Teil des Prozesses auf Probleme stößt, wie z.B. temporäre OpenAI API-Störungen, können wir elegant wiederholen oder den Fehler behandeln, ohne das gesamte System zu beeinträchtigen.

Diese Robustheit, kombiniert mit dem Echtzeit-Erscheinen von Tags, verbessert die Benutzererfahrung und vermittelt den Eindruck von KI-"Magie" bei der Arbeit.

## Daten & Prompts

Ein entscheidender Schritt in unserem Auto-Tagging-Prozess ist die Vorbereitung der Daten, die an das GPT-3.5-Modell gesendet werden. Dieser Schritt erforderte sorgfältige Überlegungen, um ein Gleichgewicht zwischen der Bereitstellung ausreichenden Kontexts für genaues Tagging und der Aufrechterhaltung von Effizienz und Datenschutz zu finden. Hier ist ein detaillierter Blick auf unseren Datenvorbereitungsprozess.

Für jeden Datensatz stellen wir folgende Informationen zusammen:

- **Listenname**: Bietet Kontext über die breitere Kategorie oder Phase des Projekts.
- **Datensatztitel**: Enthält oft Schlüsselinformationen über den Zweck oder Inhalt des Datensatzes.
- **Benutzerdefinierte Felder**: Wir schließen text- und zahlenbasierte [benutzerdefinierte Felder](/platform/features/custom-fields) ein, die oft entscheidende projektspezifische Informationen enthalten.
- **Beschreibung**: Enthält typischerweise die detailliertesten Informationen über den Datensatz.
- **Kommentare**: Können zusätzlichen Kontext oder Updates liefern, die für das Tagging relevant sein könnten.
- **Fälligkeitsdatum**: Zeitliche Informationen, die die Tag-Auswahl beeinflussen könnten.

Interessanterweise senden wir keine vorhandenen Tag-Daten an GPT-3.5, und wir tun dies, um eine Voreingenommenheit des Modells zu vermeiden.

Der Kern unserer Auto-Tagging-Funktion liegt darin, wie wir mit dem GPT-3.5-Modell interagieren und seine Antworten verarbeiten. Dieser Abschnitt unserer Pipeline erforderte sorgfältiges Design, um genaues, konsistentes und effizientes Tagging sicherzustellen.

Wir verwenden einen sorgfältig gestalteten System-Prompt, um die KI über ihre Aufgabe zu instruieren. Hier ist eine Aufschlüsselung unseres Prompts und die Begründung für jede Komponente:

```
You will be provided with record data, and your task is to choose the tags that are relevant to the record.
You can respond with an empty array if you are unsure.
Available tags: ${tags}.
Today: ${currentDate}.
Please respond in JSON using the following format:
{ "tags": ["tag-1", "tag-2"] }
```

- **Aufgabendefinition:** Wir geben die Aufgabe der KI klar an, um fokussierte Antworten zu gewährleisten.
- **Unsicherheitsbehandlung:** Wir erlauben explizit leere Antworten, um erzwungenes Tagging zu verhindern, wenn die KI unsicher ist.
- **Verfügbare Tags:** Wir stellen eine Liste gültiger Tags (${tags}) zur Verfügung, um die Auswahl der KI auf vorhandene Projekt-Tags zu beschränken.
- **Aktuelles Datum:** Die Einbeziehung von ${currentDate} hilft der KI, den zeitlichen Kontext zu verstehen, was für bestimmte Projekttypen entscheidend sein kann.
- **Antwortformat:** Wir spezifizieren ein JSON-Format für einfaches Parsen und Fehlerprüfung.

Dieser Prompt ist das Ergebnis umfangreicher Tests und Iterationen. Wir stellten fest, dass die explizite Angabe der Aufgabe, der verfügbaren Optionen und des gewünschten Ausgabeformats die Genauigkeit und Konsistenz der KI-Antworten erheblich verbesserte - Einfachheit ist der Schlüssel!

Die Liste der verfügbaren Tags wird serverseitig generiert und vor der Aufnahme in den Prompt validiert. Wir implementieren strenge Zeichenbegrenzungen für Tag-Namen, um übergroße Prompts zu verhindern.

Wie oben erwähnt, hatten wir keine Probleme mit GPT-3.5 Turbo, zu 100% der Zeit die reine JSON-Antwort im korrekten Format zurückzubekommen.

Zusammenfassend:

- Wir kombinieren den System-Prompt mit den vorbereiteten Datensatzdaten.
- Dieser kombinierte Prompt wird über die API von OpenAI an das GPT-3.5-Modell gesendet.
- Wir verwenden eine Temperatureinstellung von 0.3, um Kreativität und Konsistenz in den KI-Antworten auszubalancieren.
- Unser API-Aufruf enthält einen max_tokens-Parameter, um die Antwortgröße zu begrenzen und die Kosten zu kontrollieren.

Sobald wir die Antwort der KI erhalten, durchlaufen wir mehrere Schritte, um die vorgeschlagenen Tags zu verarbeiten und anzuwenden:

* **JSON-Parsing**: Wir versuchen, die Antwort als JSON zu parsen. Wenn das Parsen fehlschlägt, protokollieren wir den Fehler und überspringen das Tagging für diesen Datensatz.
* **Schema-Validierung**: Wir validieren das geparste JSON gegen unser erwartetes Schema (ein Objekt mit einem "tags"-Array). Dies fängt strukturelle Abweichungen in der KI-Antwort auf.
* **Tag-Validierung**: Wir gleichen die vorgeschlagenen Tags mit unserer Liste gültiger Projekt-Tags ab. Dieser Schritt filtert alle Tags heraus, die im Projekt nicht existieren, was auftreten könnte, wenn die KI missversteht oder wenn sich Projekt-Tags zwischen Prompt-Erstellung und Antwortverarbeitung ändern.
* **Deduplizierung**: Wir entfernen alle doppelten Tags aus dem KI-Vorschlag, um redundantes Tagging zu vermeiden.
* **Anwendung**: Die validierten und deduplizierten Tags werden dann auf den Datensatz in unserer Datenbank angewendet.
* **Protokollierung und Analyse**: Wir protokollieren die endgültig angewendeten Tags. Diese Daten sind wertvoll für die Überwachung der Systemleistung und deren Verbesserung im Laufe der Zeit.

## Herausforderungen

Die Implementierung des KI-gestützten Auto-Taggings in Blue stellte mehrere einzigartige Herausforderungen dar, die jeweils innovative Lösungen erforderten, um eine robuste, effiziente und benutzerfreundliche Funktion zu gewährleisten.

### Massen-Operation rückgängig machen

Die KI-Tagging-Funktion kann sowohl bei einzelnen Datensätzen als auch in Masse durchgeführt werden. Das Problem bei der Massenoperation ist, dass der Benutzer, wenn ihm das Ergebnis nicht gefällt, manuell durch Tausende von Datensätzen gehen und die Arbeit der KI rückgängig machen müsste. Das ist eindeutig inakzeptabel.

Um dies zu lösen, haben wir ein innovatives Tagging-Sitzungssystem implementiert. Jeder Massen-Tagging-Operation wird eine eindeutige Sitzungs-ID zugewiesen, die mit allen während dieser Sitzung angewendeten Tags verknüpft ist. Dies ermöglicht es uns, Rückgängig-Operationen effizient zu verwalten, indem wir einfach alle Tags löschen, die mit einer bestimmten Sitzungs-ID verknüpft sind. Wir entfernen auch verwandte Audit-Trails, um sicherzustellen, dass rückgängig gemachte Operationen keine Spuren im System hinterlassen. Dieser Ansatz gibt Benutzern das Vertrauen, mit KI-Tagging zu experimentieren, da sie wissen, dass sie Änderungen bei Bedarf leicht rückgängig machen können.

### Datenschutz

Datenschutz war eine weitere kritische Herausforderung, der wir uns stellten.

Unsere Nutzer vertrauen uns ihre Projektdaten an, und es war von größter Bedeutung sicherzustellen, dass diese Informationen nicht von OpenAI gespeichert oder für das Modelltraining verwendet werden. Wir haben dies an mehreren Fronten angegangen.

Zunächst haben wir eine Vereinbarung mit OpenAI getroffen, die die Verwendung unserer Daten für das Modelltraining ausdrücklich verbietet. Zusätzlich löscht OpenAI die Daten nach der Verarbeitung, was eine zusätzliche Schutzebene für den Datenschutz bietet.

Auf unserer Seite haben wir die Vorsichtsmaßnahme getroffen, sensible Informationen wie Zuweisungsdetails von den an die KI gesendeten Daten auszuschließen, so dass sichergestellt ist, dass spezifische Personennamen nicht zusammen mit anderen Daten an Dritte gesendet werden.

Dieser umfassende Ansatz ermöglicht es uns, KI-Fähigkeiten zu nutzen und gleichzeitig die höchsten Standards für Datenschutz und Sicherheit aufrechtzuerhalten.

### Ratenlimits und Fehlererkennung

Eine unserer Hauptsorgen waren Skalierbarkeit und Ratenbegrenzung. Direkte API-Aufrufe an OpenAI für jeden Datensatz wären ineffizient gewesen und könnten schnell Ratenlimits erreichen, insbesondere bei großen Projekten oder während Spitzenzeiten. Um dies zu adressieren, entwickelten wir eine Hintergrunddienst-Architektur, die es uns ermöglicht, Anfragen zu bündeln und unser eigenes Warteschlangensystem zu implementieren. Dieser Ansatz hilft uns, die API-Aufruffrequenz zu verwalten und ermöglicht eine effizientere Verarbeitung großer Datenmengen, wodurch eine reibungslose Leistung auch bei hoher Last gewährleistet wird.

Die Natur der KI-Interaktionen bedeutete, dass wir uns auch auf gelegentliche Fehler oder unerwartete Ausgaben vorbereiten mussten. Es gab Fälle, in denen die KI ungültiges JSON oder Ausgaben produzieren könnte, die nicht unserem erwarteten Format entsprachen. Um dies zu handhaben, implementierten wir robuste Fehlerbehandlung und Parsing-Logik in unserem gesamten System. Wenn die KI-Antwort kein gültiges JSON ist oder nicht den erwarteten "tags"-Schlüssel enthält, ist unser System darauf ausgelegt, dies so zu behandeln, als ob keine Tags vorgeschlagen wurden, anstatt zu versuchen, potenziell beschädigte Daten zu verarbeiten. Dies stellt sicher, dass unser System auch angesichts der KI-Unvorhersehbarkeit stabil und zuverlässig bleibt.

## Zukünftige Entwicklungen

Wir glauben, dass Funktionen und das Blue-Produkt als Ganzes niemals "fertig" sind - es gibt immer Raum für Verbesserungen.

Es gab einige Funktionen, die wir in der ersten Entwicklung in Betracht gezogen haben, die die Scoping-Phase nicht bestanden haben, aber es ist interessant, sie zu erwähnen, da wir wahrscheinlich in Zukunft eine Version davon implementieren werden.

Die erste ist das Hinzufügen von Tag-Beschreibungen. Dies würde es Endbenutzern ermöglichen, Tags nicht nur einen Namen und eine Farbe zu geben, sondern auch eine optionale Beschreibung. Diese würde auch an die KI weitergegeben werden, um weiteren Kontext zu liefern und möglicherweise die Genauigkeit zu verbessern.

Während zusätzlicher Kontext wertvoll sein könnte, sind wir uns der potenziellen Komplexität bewusst, die er einführen könnte. Es gibt ein empfindliches Gleichgewicht zwischen der Bereitstellung nützlicher Informationen und der Überforderung der Benutzer mit zu vielen Details. Bei der Entwicklung dieser Funktion werden wir uns darauf konzentrieren, diesen Sweet Spot zu finden, an dem zusätzlicher Kontext die Benutzererfahrung verbessert anstatt sie zu verkomplizieren.

Die vielleicht aufregendste Verbesserung an unserem Horizont ist die Integration des KI-Auto-Taggings mit unserem [Projektmanagement-Automatisierungssystem](/platform/features/automations).

Dies bedeutet, dass die KI-Tagging-Funktion entweder ein Auslöser oder eine Aktion einer Automatisierung sein könnte. Dies könnte enorm sein, da es diese relativ einfache KI-Kategorisierungsfunktion in ein KI-basiertes Routing-System für Arbeit verwandeln könnte.

Stellen Sie sich eine Automatisierung vor, die besagt:

Wenn KI einen Datensatz als "Kritisch" taggt -> Zuweisen an "Manager" und Senden einer benutzerdefinierten E-Mail

Dies bedeutet, dass wenn Sie einen Datensatz mit KI taggen und die KI entscheidet, dass es sich um ein kritisches Problem handelt, es automatisch dem Projektmanager zugewiesen und ihm eine benutzerdefinierte E-Mail gesendet werden kann. Dies erweitert die [Vorteile unseres Projektmanagement-Automatisierungssystems](/platform/features/automations) von einem rein regelbasierten System zu einem echten flexiblen KI-System.

Indem wir kontinuierlich die Grenzen der KI im Projektmanagement erkunden, wollen wir unseren Nutzern Tools zur Verfügung stellen, die nicht nur ihre aktuellen Bedürfnisse erfüllen, sondern die Zukunft der Arbeit antizipieren und gestalten. Wie immer werden wir diese Funktionen in enger Zusammenarbeit mit unserer Nutzer-Community entwickeln und sicherstellen, dass jede Verbesserung echten, praktischen Wert zum Projektmanagement-Prozess hinzufügt.

## Fazit

Das war's also!

Dies war eine unterhaltsame Funktion zu implementieren und einer unserer ersten Schritte in die KI, neben der [KI-Inhaltszusammenfassung](/insights/ai-content-summarization), die wir zuvor eingeführt haben. Wir wissen, dass KI in Zukunft eine immer größere Rolle im Projektmanagement spielen wird, und wir können es kaum erwarten, weitere innovative Funktionen mit fortschrittlichen LLMs (Large Language Models) einzuführen.

Es gab einiges zu bedenken bei der Implementierung, und wir sind besonders gespannt darauf, wie wir diese Funktion in Zukunft mit Blues bestehender [Projektmanagement-Automatisierungs-Engine](/insights/benefits-project-management-automation) nutzen können.

Wir hoffen auch, dass es eine interessante Lektüre war und Ihnen einen Einblick gibt, wie wir über die Entwicklung der Funktionen nachdenken, die Sie täglich nutzen.

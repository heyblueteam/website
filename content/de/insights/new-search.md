---
title: Echtzeit-Suche
category: "Product Updates"
description: Blue präsentiert eine neue blitzschnelle Suchmaschine, die Ergebnisse aus all Ihren Projekten in Millisekunden liefert und es Ihnen ermöglicht, den Kontext im Handumdrehen zu wechseln.
date: 2024-03-01
---

Wir freuen uns, die Einführung unserer neuen Suchmaschine bekannt zu geben, die darauf ausgelegt ist, wie Sie Informationen innerhalb von Blue finden, zu revolutionieren. Eine effiziente Suchfunktion ist entscheidend für ein nahtloses Projektmanagement, und unser neuestes Update stellt sicher, dass Sie schneller als je zuvor auf Ihre Daten zugreifen können.

Unsere neue Suchmaschine ermöglicht es Ihnen, nach allen Kommentaren, Dateien, Datensätzen, benutzerdefinierten Feldern, Beschreibungen und Checklisten zu suchen. Egal, ob Sie einen bestimmten Kommentar zu einem Projekt finden, schnell eine Datei lokalisieren oder nach einem bestimmten Datensatz oder Feld suchen müssen, unsere Suchmaschine liefert blitzschnelle Ergebnisse.

Wenn Tools eine Reaktionszeit von 50-100 ms erreichen, neigen sie dazu, in den Hintergrund zu treten und eine nahtlose und fast unsichtbare Benutzererfahrung zu bieten. Zum Kontext: Ein menschlicher Blinzeln dauert etwa 60-120 ms, also ist 50 ms tatsächlich schneller als ein Blinzeln! Dieses Maß an Reaktionsfähigkeit ermöglicht es Ihnen, mit Blue zu interagieren, ohne überhaupt zu bemerken, dass es da ist, und gibt Ihnen die Freiheit, sich auf die eigentliche Arbeit zu konzentrieren. Durch die Nutzung dieser Leistungsstufe stellt unsere neue Suchmaschine sicher, dass Sie schnell auf die Informationen zugreifen können, die Sie benötigen, ohne dass sie jemals Ihren Arbeitsablauf stört.

Um unser Ziel einer blitzschnellen Suche zu erreichen, haben wir die neuesten Open-Source-Technologien genutzt. Unsere Suchmaschine basiert auf MeiliSearch, einem beliebten Open-Source-Suchdienst, der natürliche Sprachverarbeitung und Vektorsuche verwendet, um schnell relevante Ergebnisse zu finden. Darüber hinaus haben wir eine In-Memory-Speicherung implementiert, die es uns ermöglicht, häufig abgerufene Daten im RAM zu speichern, wodurch die Zeit zum Zurückgeben von Suchergebnissen verkürzt wird. Diese Kombination aus MeiliSearch und In-Memory-Speicherung ermöglicht es unserer Suchmaschine, Ergebnisse in Millisekunden zu liefern, sodass Sie schnell finden können, was Sie benötigen, ohne jemals über die zugrunde liegende Technologie nachdenken zu müssen.

Die neue Suchleiste befindet sich bequem in der Navigationsleiste, sodass Sie sofort mit der Suche beginnen können. Für ein detaillierteres Sucherlebnis drücken Sie einfach die Tabulatortaste während der Suche, um auf die vollständige Suchseite zuzugreifen. Darüber hinaus können Sie die Suchfunktion von überall aus schnell mit der CMD/Ctrl+K-Tastenkombination aktivieren, was es noch einfacher macht, das zu finden, was Sie benötigen.

<video autoplay loop muted playsinline>
  <source src="/videos/search-demo.mp4" type="video/mp4">
</video>


## Zukünftige Entwicklungen

Das ist erst der Anfang. Jetzt, da wir eine Suchinfrastruktur der nächsten Generation haben, können wir in Zukunft wirklich interessante Dinge tun.

Als Nächstes wird die semantische Suche kommen, die eine erhebliche Verbesserung zur typischen Schlüsselwortsuche darstellt. Lassen Sie uns das erklären.

Diese Funktion wird es der Suchmaschine ermöglichen, den Kontext Ihrer Abfragen zu verstehen. Zum Beispiel wird die Suche nach "Meer" relevante Dokumente abrufen, selbst wenn der genaue Ausdruck nicht verwendet wird. Sie denken vielleicht: "Aber ich habe stattdessen 'Ozean' eingegeben!" - und Sie haben recht. Die Suchmaschine wird auch die Ähnlichkeit zwischen "Meer" und "Ozean" verstehen und relevante Dokumente zurückgeben, selbst wenn der genaue Ausdruck nicht verwendet wird. Diese Funktion ist besonders nützlich, wenn Sie nach Dokumenten suchen, die technische Begriffe, Akronyme oder einfach gängige Wörter enthalten, die mehrere Variationen oder Tippfehler haben.

Eine weitere kommende Funktion ist die Möglichkeit, nach Bildern anhand ihres Inhalts zu suchen. Um dies zu erreichen, werden wir jedes Bild in Ihrem Projekt verarbeiten und für jedes eine Einbettung erstellen. In einfachen Worten ist eine Einbettung ein mathematisches Koordinatensystem, das der Bedeutung eines Bildes entspricht. Das bedeutet, dass alle Bilder basierend auf ihrem Inhalt durchsucht werden können, unabhängig von ihrem Dateinamen oder ihren Metadaten. Stellen Sie sich vor, Sie suchen nach "Flussdiagramm" und finden alle Bilder, die mit Flussdiagrammen zu tun haben, *unabhängig von ihren Dateinamen.*
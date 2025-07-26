---
title: Référencer et rechercher des champs personnalisés
description: Créez facilement des projets interconnectés dans Blue, le transformant en une source unique de vérité pour votre entreprise grâce aux nouveaux champs de Référence et de Recherche.
category: "Product Updates"
date: 2023-11-01
---


Les projets dans Blue sont déjà un moyen puissant de gérer vos données commerciales et de faire avancer le travail.

Aujourd'hui, nous faisons le pas logique suivant en vous permettant d'interconnecter vos données *entre* les projets pour une flexibilité et une puissance ultimes.

L'interconnexion des projets au sein de Blue le transforme en une source unique de vérité pour votre entreprise. Cette capacité permet de créer un ensemble de données complet et interconnecté, facilitant le flux de données et améliorant la visibilité à travers les projets. En reliant les projets, les équipes peuvent obtenir une vue unifiée des opérations, améliorant ainsi la prise de décision et l'efficacité opérationnelle.

## Un exemple

Considérons la société ACME, qui utilise les champs personnalisés de Référence et de Recherche de Blue pour créer un écosystème de données interconnecté à travers ses projets Clients, Ventes et Inventaire. Les enregistrements clients dans le projet Clients sont liés via des champs de Référence aux transactions de vente dans le projet Ventes. Ce lien permet aux champs de Recherche d'extraire les détails associés aux clients, tels que les numéros de téléphone et les statuts de compte, directement dans chaque enregistrement de vente. De plus, les articles d'inventaire vendus sont affichés dans l'enregistrement de vente via un champ de Recherche faisant référence aux données de Quantité Vendue du projet Inventaire. Enfin, les retraits d'inventaire sont connectés aux ventes pertinentes via un champ de Référence dans Inventaire, pointant vers les enregistrements de Ventes. Cette configuration offre une visibilité complète sur quelle vente a déclenché le retrait d'inventaire, créant une vue intégrée à 360 degrés à travers les projets.

## Comment fonctionnent les champs de Référence

Les champs personnalisés de Référence vous permettent de créer des relations entre les enregistrements de différents projets dans Blue. Lors de la création d'un champ de Référence, l'Administrateur de Projet sélectionne le projet spécifique qui fournira la liste des enregistrements de référence. Les options de configuration incluent :

* **Sélection unique** : Permet de choisir un enregistrement de référence.
* **Sélection multiple** : Permet de choisir plusieurs enregistrements de référence.
* **Filtrage** : Définir des filtres pour permettre aux utilisateurs de sélectionner uniquement les enregistrements qui correspondent aux critères de filtrage.

Une fois configuré, les utilisateurs peuvent sélectionner des enregistrements spécifiques dans le menu déroulant du champ de Référence, établissant ainsi un lien entre les projets.

## Étendre les champs de référence à l'aide de recherches

Les champs personnalisés de Recherche sont utilisés pour importer des données à partir d'enregistrements dans d'autres projets, créant une visibilité unidirectionnelle. Ils sont toujours en lecture seule et sont connectés à un champ personnalisé de Référence spécifique. Lorsqu'un utilisateur sélectionne un ou plusieurs enregistrements à l'aide d'un champ personnalisé de Référence, le champ personnalisé de Recherche affichera les données de ces enregistrements. Les Recherches peuvent afficher des données telles que :

* Créé le
* Mis à jour le
* Date d'échéance
* Description
* Liste
* Étiquette
* Responsable
* Tout champ personnalisé pris en charge de l'enregistrement référencé — y compris d'autres champs de recherche !

Par exemple, imaginez un scénario où vous avez trois projets : **Projet A** est un projet de vente, **Projet B** est un projet de gestion des stocks, et **Projet C** est un projet de gestion des relations clients. Dans le Projet A, vous avez un champ personnalisé de Référence qui lie les enregistrements de vente aux enregistrements clients correspondants dans le Projet C. Dans le Projet B, vous avez un champ personnalisé de Recherche qui importe des informations du Projet A, telles que la quantité vendue. De cette manière, lorsqu'un enregistrement de vente est créé dans le Projet A, les informations client associées à cette vente sont automatiquement extraites du Projet C, et la quantité vendue est automatiquement extraite du Projet B. Cela vous permet de garder toutes les informations pertinentes au même endroit et de les visualiser sans avoir à créer des données en double ou à mettre à jour manuellement les enregistrements à travers les projets.

Un exemple concret de cela est une entreprise de commerce électronique qui utilise Blue pour gérer ses ventes, son inventaire et ses relations clients. Dans leur projet **Ventes**, ils ont un champ personnalisé de Référence qui lie chaque enregistrement de vente à l'enregistrement **Client** correspondant dans leur projet **Clients**. Dans leur projet **Inventaire**, ils ont un champ personnalisé de Recherche qui importe des informations du projet Ventes, telles que la quantité vendue, et les affiche dans l'enregistrement de l'article d'inventaire. Cela leur permet de voir facilement quelles ventes entraînent des retraits d'inventaire et de maintenir leurs niveaux d'inventaire à jour sans avoir à mettre à jour manuellement les enregistrements à travers les projets.

## Conclusion

Imaginez un monde où vos données de projet ne sont pas isolées mais circulent librement entre les projets, fournissant des informations complètes et favorisant l'efficacité. C'est la puissance des champs de Référence et de Recherche de Blue. En permettant des connexions de données sans faille et en offrant une visibilité en temps réel à travers les projets, ces fonctionnalités transforment la manière dont les équipes collaborent et prennent des décisions. Que vous gériez des relations clients, suiviez des ventes ou supervisiez des stocks, les champs de Référence et de Recherche dans Blue permettent à votre équipe de travailler plus intelligemment, plus rapidement et plus efficacement. Plongez dans le monde interconnecté de Blue et regardez votre productivité s'envoler.

[Consultez la documentation](https://documentation.blue.cc/custom-fields/reference) ou [inscrivez-vous et essayez-le par vous-même.](https://app.blue.cc)
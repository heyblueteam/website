---
title: Comment créer des automatisations d'email personnalisées
description: Les notifications par email personnalisées sont une fonctionnalité incroyablement puissante dans Blue qui peut aider à faire avancer le travail et à garantir que la communication est en pilote automatique.
category: "Product Updates"
---


Les automatisations d'email dans Blue sont une [automatisation puissante de la gestion de projet](/platform/features/automations) pour rationaliser la communication, garantir [un excellent travail d'équipe](/insights/great-teamwork) et faire avancer les projets. En tirant parti des données stockées dans vos enregistrements, vous pouvez envoyer automatiquement des emails personnalisés lorsque certains déclencheurs se produisent, comme la création d'un nouvel enregistrement ou le dépassement d'un délai pour une tâche.

Dans cet article, nous allons explorer comment configurer et utiliser les automatisations d'email dans Blue.

## Configuration des Automatisations d'Email

Créer une automatisation d'email dans Blue est un processus simple. Tout d'abord, sélectionnez le déclencheur qui initiera l'email automatisé. Certains déclencheurs courants incluent :

- Un nouvel enregistrement est créé
- Une étiquette est ajoutée à un enregistrement
- Un enregistrement est déplacé vers une autre liste

Ensuite, configurez les détails de l'email, y compris :

- Nom de l'expéditeur et adresse de réponse
- Adresse du destinataire (peut être statique ou extraite dynamiquement d'un champ personnalisé d'email)
- Adresses CC ou BCC (optionnel)

![](/insights/email-automations-image.webp)

L'un des principaux avantages des automatisations d'email dans Blue est la possibilité de personnaliser le contenu à l'aide de balises de fusion. Lorsque vous personnalisez l'objet et le corps de l'email, vous pouvez insérer des balises de fusion qui font référence à des données spécifiques de l'enregistrement, comme le nom de l'enregistrement ou les valeurs des champs personnalisés. Il suffit d'utiliser la syntaxe {accolades} pour insérer des balises de fusion.

Vous pouvez également inclure des pièces jointes en les faisant glisser et en les déposant dans l'email ou en utilisant l'icône de pièce jointe. Les fichiers provenant de champs personnalisés de fichiers peuvent être automatiquement attachés s'ils sont inférieurs à 10 Mo.

Avant de finaliser votre automatisation d'email, il est recommandé d'envoyer un email test à vous-même ou à un collègue pour vous assurer que tout fonctionne comme prévu.

## Cas d'Utilisation et Exemples

Les automatisations d'email dans Blue peuvent être utilisées à diverses fins. Voici quelques exemples :

1. Envoyer un email de confirmation lorsqu'un client soumet une demande via un formulaire d'admission. Configurez le déclencheur pour envoyer un email lorsqu'un nouvel enregistrement est créé via le formulaire, et assurez-vous d'inclure un champ email dans le formulaire pour capturer l'adresse du client.
2. Notifier un assigné lorsqu'une nouvelle tâche à haute priorité est créée. Configurez le déclencheur pour envoyer un email lorsqu'une étiquette "Priorité" est ajoutée à un enregistrement, et utilisez la balise de fusion {Assignee} pour envoyer automatiquement l'email à l'utilisateur assigné.
3. Envoyer un sondage à un client après qu'un ticket de support soit marqué comme résolu. Configurez le déclencheur pour envoyer un email lorsqu'un enregistrement est marqué comme complété et déplacé vers la liste "Fait". Incluez l'email du client dans un champ personnalisé et fournissez des informations détaillées sur le problème résolu dans le corps de l'email.
4. Automatiser un programme de recrutement en envoyant des emails de confirmation aux candidats. Configurez le déclencheur pour envoyer un email lorsqu'une candidature est soumise via un formulaire et ajoutée à la liste "Reçue". Capturez l'email du candidat dans le formulaire et utilisez-le pour envoyer une réponse de remerciement.

## Avantages des Automatisations d'Email

Les automatisations d'email dans Blue offrent plusieurs avantages clés :

- Communication personnalisée grâce à l'utilisation de balises de fusion et de données de champs personnalisés
- Notifications automatiques qui réduisent le travail manuel et garantissent des mises à jour en temps opportun
- Flux de travail structurés et basés sur les données qui font avancer les projets en fonction des données des enregistrements

## Conclusion

Les automatisations d'email dans Blue sont un outil précieux pour rationaliser la communication et maintenir les projets sur la bonne voie. En tirant parti des déclencheurs, des balises de fusion et des données de champs personnalisés, vous pouvez créer des emails automatisés et personnalisés qui améliorent la productivité de votre équipe et garantissent que les mises à jour importantes ne sont jamais manquées. Avec une large gamme de cas d'utilisation et une configuration facile, les automatisations d'email sont une fonctionnalité indispensable pour tout utilisateur de Blue cherchant à optimiser son flux de travail.
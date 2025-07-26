---
title: Webhooks
description: Blue introduit des webhooks granulaires pour permettre aux clients d'envoyer des données aux systèmes en quelques millisecondes.
category: "Product Updates"
date: 2023-06-01
---


Blue [dispose d'une API avec une couverture fonctionnelle de 100 % depuis des années.](/platform/api), vous permettant d'extraire des données telles que des listes de projets et des enregistrements, ou de publier de nouvelles informations dans Blue. Mais que faire si vous souhaitez que votre propre système reçoive des mises à jour lorsque quelque chose change dans Blue ? C'est là que les webhooks entrent en jeu.

Au lieu de vérifier constamment l'API de Blue pour des mises à jour, Blue peut désormais notifier proactivement votre plateforme lorsque de nouveaux événements se produisent.

Cependant, la mise en œuvre efficace des webhooks peut être un défi.

## Une Nouvelle Approche des Webhooks

De nombreuses plateformes proposent un webhook universel qui envoie des données pour tous les types d'événements, laissant à vous et votre équipe le soin de trier les informations et d'extraire ce qui est pertinent.

Chez Blue, nous nous sommes demandé : **Y a-t-il un meilleur moyen ? Comment pouvons-nous rendre nos webhooks aussi conviviaux que possible pour les développeurs ?**

Notre solution ?

Un contrôle précis !

Avec Blue, vous pouvez choisir *exactement* quels événements, ou *combinaisons* d'événements, déclencheront un webhook. Vous pouvez également spécifier dans quels projets, ou *combinaisons* de projets (même à travers différentes entreprises !), les événements doivent se produire.

Ce niveau de granularité est sans précédent, et il vous permet de recevoir uniquement les données dont vous avez besoin, au moment où vous en avez besoin.

## Fiabilité et Facilité d'Utilisation

Nous avons intégré de l'intelligence dans notre système de webhook pour garantir la fiabilité.

Blue surveille automatiquement la santé de vos connexions de webhook et utilise une logique de réessai intelligente, tentant la livraison jusqu'à cinq fois avant de désactiver un webhook. Cela aide à prévenir la perte de données et réduit le besoin d'intervention manuelle.

Configurer des webhooks dans Blue est simple.

Vous pouvez les configurer via notre API pour une configuration programmatique, ou utiliser notre application web pour une interface conviviale. Cette flexibilité permet aux développeurs comme aux utilisateurs non techniques de tirer parti de la puissance des webhooks.

## Données en Temps Réel, Possibilités Infinies

En tirant parti des webhooks de Blue, vous pouvez créer des intégrations en temps réel entre Blue et vos autres systèmes d'entreprise. Cela ouvre un monde de possibilités pour l'automatisation, la synchronisation des données et les flux de travail personnalisés. Que vous mettiez à jour un CRM, déclenchiez des alertes ou alimentiez des outils d'analyse, les webhooks de Blue fournissent la connexion en temps réel dont vous avez besoin.

Prêt à commencer avec les webhooks de Blue ? [Consultez notre documentation détaillée](https://documentation.blue.cc/integrations/webhooks) pour des guides de mise en œuvre, des meilleures pratiques et des cas d'utilisation d'exemple.

Si vous avez besoin d'aide, [notre équipe de support](/support) est toujours là pour vous aider à tirer le meilleur parti de cette fonctionnalité puissante.
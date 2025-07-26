---
title: FAQ sur la sécurité de Blue
description: Voici une liste des questions les plus fréquemment posées sur les protocoles et pratiques de sécurité chez Blue.
category: "FAQ"
date: 2024-07-19
---


Notre mission est d'organiser le travail dans le monde en construisant la meilleure plateforme de gestion de projet sur la planète.

Au cœur de cette mission se trouve la nécessité de garantir que notre plateforme est sécurisée, fiable et digne de confiance. Nous comprenons que pour être votre source unique de vérité, Blue doit protéger vos données commerciales sensibles contre les menaces extérieures, la perte de données et les temps d'arrêt.

Cela signifie que nous prenons la sécurité très au sérieux chez Blue.

Lorsque nous pensons à la sécurité, nous adoptons une approche holistique qui se concentre sur trois domaines clés :

1.  **Sécurité de l'infrastructure et du réseau** : Assure que nos systèmes physiques et virtuels sont protégés contre les menaces externes et l'accès non autorisé.
2.  **Sécurité des logiciels** : Se concentre sur la sécurité du code lui-même, y compris les pratiques de codage sécurisé, les revues de code régulières et la gestion des vulnérabilités.
3.  **Sécurité de la plateforme** : Inclut les fonctionnalités au sein de Blue, telles que [les contrôles d'accès sophistiqués](/platform/features/user-permissions), garantissant que les projets sont privés par défaut, ainsi que d'autres mesures pour protéger les données et la vie privée des utilisateurs.


## Quelle est la scalabilité de Blue ?

C'est une question importante, car vous souhaitez un système qui peut *croître* avec vous. Vous ne voulez pas avoir à changer votre plateforme de gestion de projet et de processus dans six ou douze mois.

Nous choisissons les fournisseurs de plateforme avec soin, pour nous assurer qu'ils peuvent gérer les charges de travail exigeantes de nos clients. Nous utilisons des services cloud de certains des meilleurs fournisseurs de cloud au monde qui alimentent des entreprises telles que [Spotify](https://spotify.com) et [Netflix](https://netflix.com), qui ont plusieurs ordres de grandeur de trafic par rapport à nous.

Les principaux fournisseurs de cloud que nous utilisons sont :

- **[Cloudflare](https://cloudflare.com)** : Nous gérons le DNS (Service de nom de domaine) via Cloudflare ainsi que notre site web marketing qui fonctionne sur [Cloudflare Pages](https://pages.cloudflare.com/).
- **[Amazon Web Services](https://aws.amazon.com/)** : Nous utilisons AWS pour notre base de données, qui est [Aurora](https://aws.amazon.com/rds/aurora/), pour le stockage de fichiers via [Simple Storage Service (S3)](https://aws.amazon.com/s3/), et également pour l'envoi d'emails via [Simple Email Service (SES)](https://aws.amazon.com/ses/)
- **[Render](https://render.com)** : Nous utilisons Render pour nos serveurs front-end, serveurs d'application/API, nos services en arrière-plan, notre système de mise en file d'attente et notre base de données Redis. Fait intéressant, Render est en fait construit *au-dessus* d'AWS !


## Quelle est la sécurité des fichiers dans Blue ?

Commençons par le stockage des données. Nos fichiers sont hébergés sur [AWS S3](https://aws.amazon.com/s3/), qui est le service de stockage d'objets cloud le plus populaire au monde, avec une scalabilité, une disponibilité des données, une sécurité et des performances de pointe.

Nous avons une disponibilité des fichiers de 99,99 % et une durabilité élevée de 99,999999999 %.

Décomposons ce que cela signifie.

La disponibilité fait référence à la durée pendant laquelle les données sont opérationnelles et accessibles. La disponibilité des fichiers de 99,99 % signifie que nous pouvons nous attendre à ce que les fichiers soient indisponibles pendant pas plus d'environ 8,76 heures par an.

La durabilité fait référence à la probabilité que les données restent intactes et non corrompues au fil du temps. Ce niveau de durabilité signifie que nous pouvons nous attendre à perdre pas plus d'un fichier sur 10 milliards de fichiers téléchargés, grâce à une redondance extensive et à une réplication des données à travers plusieurs centres de données.

Nous utilisons [S3 Intelligent-Tiering](https://aws.amazon.com/s3/storage-classes/intelligent-tiering/) pour déplacer automatiquement les fichiers vers différentes classes de stockage en fonction de la fréquence d'accès. En fonction des modèles d'activité de centaines de milliers de projets, nous remarquons que la plupart des fichiers sont accessibles selon un modèle qui ressemble à une courbe de recul exponentiel. Cela signifie que la plupart des fichiers sont très fréquemment accessibles dans les premiers jours, puis sont rapidement accédés de moins en moins fréquemment. Cela nous permet de déplacer les fichiers plus anciens vers un stockage plus lent, mais significativement moins cher, sans affecter l'expérience utilisateur de manière significative.

Les économies de coûts pour cela sont significatives. S3 Standard-Infrequent Access (S3 Standard-IA) est environ 1,84 fois moins cher que S3 Standard. Cela signifie que pour chaque dollar que nous aurions dépensé sur S3 Standard, nous ne dépensons qu'environ 54 cents sur S3 Standard-IA pour la même quantité de données stockées.

| Fonctionnalité           | S3 Standard             | S3 Standard-IA       |
|--------------------------|-------------------------|-----------------------|
| Coût de stockage          | 0,023 $ - 0,021 $ par Go | 0,0125 $ par Go       |
| Coût de requête (PUT, etc.) | 0,005 $ par 1 000 requêtes | 0,01 $ par 1 000 requêtes |
| Coût de requête (GET)    | 0,0004 $ par 1 000 requêtes | 0,001 $ par 1 000 requêtes |
| Coût de récupération de données | 0,00 $ par Go            | 0,01 $ par Go          |


Les fichiers que vous téléchargez via Blue sont chiffrés à la fois en transit et au repos. Les données transférées vers et depuis Amazon S3 sont sécurisées à l'aide de [Transport Layer Security (TLS)](https://www.internetsociety.org/deploy360/tls/basics), protégeant contre [l'écoute clandestine](https://en.wikipedia.org/wiki/Network_eavesdropping) et [les attaques de type homme du milieu](https://en.wikipedia.org/wiki/Man-in-the-middle_attack). Pour le chiffrement au repos, Amazon S3 utilise le chiffrement côté serveur (SSE-S3), qui chiffre automatiquement tous les nouveaux téléchargements avec un chiffrement AES-256, Amazon gérant les clés de chiffrement. Cela garantit que vos données restent sécurisées tout au long de leur cycle de vie.

## Qu'en est-il des données non-fichiers ?

Notre base de données est alimentée par [AWS Aurora](https://aws.amazon.com/rds/aurora/), un service de base de données relationnelle moderne qui garantit des performances, une disponibilité et une sécurité élevées pour vos données.

Les données dans Aurora sont chiffrées à la fois en transit et au repos. Nous utilisons SSL (AES-256) pour sécuriser les connexions entre votre instance de base de données et votre application, protégeant les données lors du transfert. Pour le chiffrement au repos, Aurora utilise des clés gérées via AWS Key Management Service (KMS), garantissant que toutes les données stockées, y compris les sauvegardes automatiques, les instantanés et les répliques, sont chiffrées et protégées.

Aurora dispose d'un système de stockage distribué, tolérant aux pannes et auto-réparateur. Ce système est découplé des ressources de calcul et peut s'auto-scaler jusqu'à 128 TiB par instance de base de données. Les données sont répliquées à travers trois [zones de disponibilité](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) (AZs), offrant une résilience contre la perte de données et garantissant une haute disponibilité. En cas de panne de base de données, Aurora réduit les temps de récupération à moins de 60 secondes, assurant une interruption minimale.

Blue sauvegarde continuellement notre base de données sur Amazon S3, permettant une récupération à un instant donné. Cela signifie que nous pouvons restaurer la base de données principale de Blue à tout moment spécifique dans les cinq dernières minutes, garantissant que vos données sont toujours récupérables. Nous prenons également des instantanés réguliers de la base de données pour des périodes de rétention de sauvegarde plus longues.

En tant que service entièrement géré, Aurora automatise les tâches d'administration chronophages telles que la provisionnement matériel, la configuration de la base de données, les mises à jour et les sauvegardes. Cela réduit la charge opérationnelle et garantit que notre base de données est toujours à jour avec les derniers correctifs de sécurité et améliorations de performances.

Si nous sommes plus efficaces, nous pouvons transmettre nos économies de coûts à nos clients avec notre [tarification de pointe](/pricing).

Aurora est conforme à diverses normes industrielles telles que HIPAA, GDPR et SOC 2, garantissant que vos pratiques de gestion des données respectent des exigences réglementaires strictes. Des audits de sécurité réguliers et une intégration avec [Amazon GuardDuty](https://aws.amazon.com/guardduty/) aident à détecter et à atténuer les menaces potentielles à la sécurité.

## Comment Blue garantit-elle la sécurité des connexions ?

Blue utilise [des liens magiques par email](https://documentation.blue.cc/user-management/magic-links) pour fournir un accès sécurisé et pratique à votre compte, éliminant ainsi le besoin de mots de passe traditionnels.

Cette approche améliore considérablement la sécurité en atténuant les menaces courantes associées aux connexions basées sur des mots de passe. En éliminant les mots de passe, les liens magiques protègent contre les attaques de phishing et le vol de mots de passe, *car il n'y a pas de mot de passe à voler ou à exploiter.*

Chaque lien magique est valide pour une seule session de connexion, réduisant le risque d'accès non autorisé. De plus, ces liens expirent après 15 minutes, garantissant que tout lien inutilisé ne peut pas être exploité, renforçant ainsi la sécurité.

La commodité offerte par les liens magiques est également remarquable. Les liens magiques offrent une expérience de connexion sans tracas, vous permettant d'accéder à votre compte *sans* avoir à vous souvenir de mots de passe complexes.

Cela simplifie non seulement le processus de connexion, mais empêche également les violations de sécurité qui se produisent lorsque les mots de passe sont réutilisés sur plusieurs services. De nombreux utilisateurs ont tendance à utiliser le même mot de passe sur diverses plateformes, ce qui signifie qu'une violation de sécurité sur un service pourrait compromettre leurs comptes sur d'autres services, y compris Blue. En utilisant des liens magiques, la sécurité de Blue ne dépend pas des pratiques de sécurité d'autres services, offrant une couche de protection plus robuste et indépendante pour nos utilisateurs.

Lorsque vous demandez à vous connecter à votre compte Blue, une URL de connexion unique est envoyée à votre email. En cliquant sur ce lien, vous serez instantanément connecté à votre compte. Le lien est conçu pour expirer après une seule utilisation ou après 15 minutes, selon la première éventualité, ajoutant une couche de sécurité supplémentaire. En utilisant des liens magiques, Blue garantit que votre processus de connexion est à la fois sécurisé et convivial, offrant tranquillité d'esprit et commodité.

## Comment puis-je vérifier la fiabilité et la disponibilité de Blue ?

Chez Blue, nous nous engageons à maintenir un niveau élevé de fiabilité et de transparence pour nos utilisateurs. Pour fournir une visibilité sur les performances de notre plateforme, nous offrons une [page de statut système dédiée](https://status.blue.cc) qui est également liée depuis notre pied de page sur chaque page de notre site web.

![](/insights/status-page.png)

Cette page affiche nos données historiques de disponibilité, vous permettant de voir à quelle fréquence nos services ont été disponibles au fil du temps. De plus, la page de statut comprend des rapports d'incidents détaillés, fournissant une transparence sur les problèmes passés, leur impact, et les mesures que nous avons prises pour les résoudre et prévenir de futures occurrences.
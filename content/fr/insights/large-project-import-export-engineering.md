---
title:  Mise à l'échelle des importations et exportations CSV à 250 000+ enregistrements
description: Découvrez comment Blue a multiplié par 10 les importations et exportations CSV en utilisant Rust, une architecture évolutive et des choix technologiques stratégiques dans le SaaS B2B.
category: "Engineering"
date: 2024-07-18
---


Chez Blue, nous [repoussons constamment les limites](/platform/roadmap) de ce qui est possible dans les logiciels de gestion de projet. Au fil des ans, nous avons [publié des centaines de fonctionnalités](/platform/changelog).

Notre dernière prouesse technique ?

Une refonte complète de notre système d'[importation CSV](https://documentation.blue.cc/integrations/csv-import) et d'[exportation](https://documentation.blue.cc/integrations/csv-export), améliorant considérablement les performances et la scalabilité.

Cet article vous emmène dans les coulisses de la manière dont nous avons relevé ce défi, les technologies que nous avons utilisées et les résultats impressionnants que nous avons obtenus.

La chose la plus intéressante ici est que nous avons dû sortir de notre [pile technologique](https://sop.blue.cc/product/technology-stack) habituelle pour atteindre les résultats que nous souhaitions. C'est une décision qui doit être prise avec soin, car les répercussions à long terme peuvent être sévères en termes de dette technologique et de frais de maintenance à long terme.

<video autoplay loop muted playsinline>
  <source src="/videos/import-export-video.mp4" type="video/mp4">
</video>

## Mise à l'échelle pour les besoins des entreprises

Notre parcours a commencé par une demande d'un client entreprise dans le secteur des événements. Ce client utilise Blue comme son hub central pour gérer d'énormes listes d'événements, de lieux et d'intervenants, l'intégrant parfaitement à son site web.

Pour eux, Blue n'est pas juste un outil — c'est la source unique de vérité pour l'ensemble de leur opération.

Bien que nous soyons toujours fiers d'entendre que nos clients nous utilisent pour des besoins aussi critiques, il y a aussi une grande responsabilité de notre côté pour garantir un système rapide et fiable.

Alors que ce client développait ses opérations, il a rencontré un obstacle significatif : **l'importation et l'exportation de grands fichiers CSV contenant 100 000 à 200 000+ enregistrements.**

Cela dépassait les capacités de notre système à l'époque. En fait, notre ancien système d'importation/exportation avait déjà du mal avec des importations et des exportations contenant plus de 10 000 à 20 000 enregistrements ! Donc, 200 000+ enregistrements étaient hors de question.

Les utilisateurs ont connu des temps d'attente frustrants, et dans certains cas, les importations ou exportations *ne parvenaient pas à se terminer*. Cela a considérablement affecté leurs opérations car ils comptaient sur des importations et des exportations quotidiennes pour gérer certains aspects de leurs opérations.

> La multi-location est une architecture où une seule instance de logiciel sert plusieurs clients (locataires). Bien que cela soit efficace, cela nécessite une gestion minutieuse des ressources pour garantir que les actions d'un locataire n'impactent pas négativement les autres.

Et cette limitation n'affectait pas seulement ce client particulier.

En raison de notre architecture multi-locataire — où plusieurs clients partagent la même infrastructure — une seule importation ou exportation gourmande en ressources pouvait potentiellement ralentir les opérations pour d'autres utilisateurs, ce qui se produisait souvent dans la pratique.

Comme d'habitude, nous avons effectué une analyse de construction contre achat, pour comprendre si nous devions passer du temps à mettre à niveau notre propre système ou acheter un système auprès de quelqu'un d'autre. Nous avons examiné diverses possibilités.

Le fournisseur qui s'est démarqué était un fournisseur SaaS appelé [Flatfile](https://flatfile.com/). Leur système et leurs capacités semblaient exactement ce dont nous avions besoin.

Mais, après avoir examiné leur [tarification](https://flatfile.com/pricing/), nous avons décidé que cela finirait par être une solution extrêmement coûteuse pour une application de notre taille — *2 $/fichier s'accumule très rapidement !* — et il valait mieux étendre notre moteur d'importation/exportation CSV intégré.

Pour relever ce défi, nous avons pris une décision audacieuse : introduire Rust dans notre pile technologique principalement Javascript. Ce langage de programmation système, connu pour ses performances et sa sécurité, était l'outil parfait pour nos besoins critiques de parsing CSV et de mapping de données.

Voici comment nous avons abordé la solution.

### Introduction de services en arrière-plan

La base de notre solution était l'introduction de services en arrière-plan pour gérer des tâches gourmandes en ressources. Cette approche nous a permis de décharger le traitement lourd de notre serveur principal, améliorant considérablement les performances globales du système. Notre architecture de services en arrière-plan est conçue avec la scalabilité à l'esprit. Comme tous les composants de notre infrastructure, ces services s'auto-scalent en fonction de la demande.

Cela signifie qu'en période de pointe, lorsque plusieurs grandes importations ou exportations sont traitées simultanément, le système alloue automatiquement plus de ressources pour gérer la charge accrue. Inversement, pendant les périodes plus calmes, il réduit la taille pour optimiser l'utilisation des ressources.

Cette architecture de services en arrière-plan évolutive a bénéficié à Blue non seulement pour les importations et exportations CSV. Au fil du temps, nous avons déplacé un nombre substantiel de fonctionnalités vers des services en arrière-plan pour alléger nos serveurs principaux :

- **[Calculs de formules](https://documentation.blue.cc/custom-fields/formula)** : Décharge les opérations mathématiques complexes pour garantir des mises à jour rapides des champs dérivés sans impacter les performances du serveur principal.
- **[Tableaux de bord/Graphiques](/platform/features/dashboards)** : Traite de grands ensembles de données en arrière-plan pour générer des visualisations à jour sans ralentir l'interface utilisateur.
- **[Index de recherche](https://documentation.blue.cc/projects/search)** : Met à jour en continu l'index de recherche en arrière-plan, garantissant des résultats de recherche rapides et précis sans impacter les performances du système.
- **[Copie de projets](https://documentation.blue.cc/projects/copying-projects)** : Gère la réplication de grands projets complexes en arrière-plan, permettant aux utilisateurs de continuer à travailler pendant que la copie est en cours de création.
- **[Automatisations de gestion de projet](/platform/features/automations)** : Exécute des flux de travail automatisés définis par l'utilisateur en arrière-plan, garantissant des actions en temps voulu sans bloquer d'autres opérations.
- **[Enregistrements récurrents](https://documentation.blue.cc/records/repeat)** : Génère des tâches ou événements récurrents en arrière-plan, maintenant l'exactitude des horaires sans alourdir l'application principale.
- **[Champs personnalisés de durée](https://documentation.blue.cc/custom-fields/duration)** : Calcule et met à jour en continu la différence de temps entre deux événements dans Blue, fournissant des données de durée en temps réel sans impacter la réactivité du système.

## Nouveau module Rust pour le parsing de données

Le cœur de notre solution de traitement CSV est un module Rust personnalisé. Bien que cela ait marqué notre première incursion en dehors de notre pile technologique principale de Javascript, la décision d'utiliser Rust a été motivée par ses performances exceptionnelles dans les opérations concurrentes et les tâches de traitement de fichiers.

Les points forts de Rust s'alignent parfaitement avec les exigences du parsing CSV et du mapping de données. Ses abstractions sans coût permettent une programmation de haut niveau sans sacrifier les performances, tandis que son modèle de propriété garantit la sécurité de la mémoire sans avoir besoin de collecte des déchets. Ces caractéristiques font de Rust un choix particulièrement adapté pour gérer efficacement et en toute sécurité de grands ensembles de données.

Pour le parsing CSV, nous avons utilisé la crate csv de Rust, qui offre une lecture et une écriture de données CSV à haute performance. Nous avons combiné cela avec une logique de mapping de données personnalisée pour garantir une intégration transparente avec les structures de données de Blue.

La courbe d'apprentissage pour Rust était raide mais gérable. Notre équipe a consacré environ deux semaines à un apprentissage intensif à ce sujet.

Les améliorations étaient impressionnantes :

![](/insights/import-export.png)

Notre nouveau système peut traiter la même quantité d'enregistrements que notre ancien système pouvait traiter en 15 minutes en environ 30 secondes.

## Interaction entre le serveur web et la base de données

Pour le composant serveur web de notre implémentation Rust, nous avons choisi Rocket comme notre framework. Rocket s'est démarqué par sa combinaison de performances et de fonctionnalités conviviales pour les développeurs. Son typage statique et sa vérification à la compilation s'alignent bien avec les principes de sécurité de Rust, nous aidant à détecter les problèmes potentiels tôt dans le processus de développement. 
Du côté de la base de données, nous avons opté pour SQLx. Cette bibliothèque SQL asynchrone pour Rust offre plusieurs avantages qui la rendent idéale pour nos besoins :

- SQL sécurisé par type : SQLx nous permet d'écrire du SQL brut avec des requêtes vérifiées à la compilation, garantissant la sécurité des types sans sacrifier les performances.
- Support asynchrone : Cela s'aligne bien avec Rocket et notre besoin d'opérations de base de données efficaces et non bloquantes.
- Indépendant de la base de données : Bien que nous utilisions principalement [AWS Aurora](https://aws.amazon.com/rds/aurora/), qui est compatible MySQL, le support de SQLx pour plusieurs bases de données nous donne de la flexibilité pour l'avenir au cas où nous déciderions de changer.

## Optimisation du traitement par lots

Notre parcours vers la configuration optimale du traitement par lots a été marqué par des tests rigoureux et une analyse minutieuse. Nous avons effectué des benchmarks étendus avec diverses combinaisons de transactions concurrentes et de tailles de lots, mesurant non seulement la vitesse brute mais aussi l'utilisation des ressources et la stabilité du système.

Le processus a impliqué la création de jeux de données de test de tailles et de complexités variées, simulant des modèles d'utilisation réels. Nous avons ensuite fait passer ces ensembles de données dans notre système, ajustant le nombre de transactions concurrentes et la taille des lots pour chaque exécution.

Après avoir analysé les résultats, nous avons constaté que le traitement de 5 transactions concurrentes avec une taille de lot de 500 enregistrements offrait le meilleur équilibre entre vitesse et utilisation des ressources. Cette configuration nous permet de maintenir un débit élevé sans surcharger notre base de données ou consommer une mémoire excessive.

Fait intéressant, nous avons constaté qu'augmenter la concurrence au-delà de 5 transactions ne produisait pas de gains de performance significatifs et entraînait parfois une contention accrue de la base de données. De même, des tailles de lots plus importantes amélioraient la vitesse brute mais au prix d'une utilisation de mémoire plus élevée et de temps de réponse plus longs pour les importations/exportations de petite à moyenne taille.

## Exportations CSV via des liens par e-mail

Le dernier élément de notre solution aborde le défi de la livraison de grands fichiers exportés aux utilisateurs. Au lieu de fournir un téléchargement direct depuis notre application web, ce qui pourrait entraîner des problèmes de délai d'attente et une charge accrue sur le serveur, nous avons mis en place un système de liens de téléchargement par e-mail.

Lorsqu'un utilisateur initie une grande exportation, notre système traite la demande en arrière-plan. Une fois terminé, plutôt que de maintenir la connexion ouverte ou de stocker le fichier sur nos serveurs web, nous téléchargeons le fichier dans un emplacement de stockage temporaire sécurisé. Nous générons ensuite un lien de téléchargement unique et sécurisé et l'envoyons par e-mail à l'utilisateur.

Ces liens de téléchargement sont valides pendant 2 heures, trouvant un équilibre entre la commodité pour l'utilisateur et la sécurité de l'information. Ce délai donne aux utilisateurs amplement l'occasion de récupérer leurs données tout en garantissant que les informations sensibles ne restent pas accessibles indéfiniment.

La sécurité de ces liens de téléchargement était une priorité absolue dans notre conception. Chaque lien est :

- Unique et généré aléatoirement, rendant pratiquement impossible de le deviner
- Valide pendant seulement 2 heures
- Chiffré en transit, garantissant la sécurité des données lors du téléchargement

Cette approche offre plusieurs avantages :

- Elle réduit la charge sur nos serveurs web, car ils n'ont pas besoin de gérer directement les téléchargements de fichiers volumineux
- Elle améliore l'expérience utilisateur, en particulier pour les utilisateurs ayant des connexions Internet plus lentes qui pourraient rencontrer des problèmes de délai d'attente avec des téléchargements directs
- Elle fournit une solution plus fiable pour les très grandes exportations qui pourraient dépasser les limites de délai d'attente web typiques

Les retours des utilisateurs sur cette fonctionnalité ont été extrêmement positifs, beaucoup appréciant la flexibilité qu'elle offre dans la gestion des grandes exportations de données.

## Exportation de données filtrées

L'autre amélioration évidente était de permettre aux utilisateurs d'exporter uniquement les données qui étaient déjà filtrées dans leur vue de projet. Cela signifie que s'il y a une étiquette active "priorité", alors seuls les enregistrements ayant cette étiquette se retrouveraient dans l'exportation CSV. Cela signifie moins de temps à manipuler les données dans Excel pour filtrer ce qui n'est pas important, et cela nous aide également à réduire le nombre de lignes à traiter.

## Perspectives d'avenir

Bien que nous n'ayons pas de plans immédiats pour étendre notre utilisation de Rust, ce projet nous a montré le potentiel de cette technologie pour des opérations critiques en termes de performances. C'est une option passionnante que nous avons désormais dans notre boîte à outils pour les besoins d'optimisation futurs. Cette refonte des importations et exportations CSV s'aligne parfaitement avec l'engagement de Blue envers la scalabilité.

Nous sommes déterminés à fournir une plateforme qui évolue avec nos clients, gérant leurs besoins de données croissants sans compromettre les performances.

La décision d'introduire Rust dans notre pile technologique n'a pas été prise à la légère. Elle a soulevé une question importante à laquelle de nombreuses équipes d'ingénierie sont confrontées : Quand est-il approprié de sortir de votre pile technologique principale, et quand devriez-vous vous en tenir à des outils familiers ?

Il n'y a pas de réponse unique, mais chez Blue, nous avons développé un cadre pour prendre ces décisions cruciales :

- **Approche axée sur le problème :** Nous commençons toujours par définir clairement le problème que nous essayons de résoudre. Dans ce cas, nous devions améliorer considérablement les performances des importations et exportations CSV pour de grands ensembles de données.
- **Épuisement des solutions existantes :** Avant de chercher en dehors de notre pile principale, nous explorons minutieusement ce qui peut être réalisé avec nos technologies existantes. Cela implique souvent du profilage, de l'optimisation et une reconsidération de notre approche dans des contraintes familières.
- **Quantification du gain potentiel :** Si nous envisageons une nouvelle technologie, nous devons être en mesure d'articuler clairement et, idéalement, de quantifier les avantages. Pour notre projet CSV, nous avons projeté des améliorations d'ordre de grandeur en termes de vitesse de traitement.
- **Évaluation des coûts :** L'introduction d'une nouvelle technologie ne concerne pas seulement le projet immédiat. Nous considérons les coûts à long terme :
  - Courbe d'apprentissage pour l'équipe
  - Maintenance et support continus
  - Complications potentielles dans le déploiement et les opérations
  - Impact sur le recrutement et la composition de l'équipe
- **Confinement et intégration :** Si nous introduisons une nouvelle technologie, nous visons à la contenir dans une partie spécifique et bien définie de notre système. Nous veillons également à avoir un plan clair sur la manière dont elle s'intégrera à notre pile existante.
- **Préparation pour l'avenir :** Nous considérons si ce choix technologique ouvre de futures opportunités ou s'il pourrait nous coincer.

L'un des principaux risques d'adopter fréquemment de nouvelles technologies est de se retrouver avec ce que nous appelons un *"zoo technologique"* - un écosystème fragmenté où différentes parties de votre application sont écrites dans différents langages ou frameworks, nécessitant une large gamme de compétences spécialisées pour maintenir.

## Conclusion

Ce projet illustre l'approche de Blue en matière d'ingénierie : *nous n'avons pas peur de sortir de notre zone de confort et d'adopter de nouvelles technologies lorsque cela signifie offrir une expérience significativement meilleure à nos utilisateurs.*

En réimaginant notre processus d'importation et d'exportation CSV, nous avons non seulement résolu un besoin pressant pour un client entreprise, mais amélioré l'expérience pour tous nos utilisateurs traitant de grands ensembles de données.

Alors que nous continuons à repousser les limites de ce qui est possible dans [les logiciels de gestion de projet](/solutions/use-case/project-management), nous sommes impatients de relever d'autres défis comme celui-ci.

Restez à l'écoute pour plus de [plongées approfondies dans l'ingénierie qui alimente Blue!](/insights/engineering-blog)
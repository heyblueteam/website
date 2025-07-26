---
title:  Création du moteur de permissions personnalisé de Blue
description: Découvrez les coulisses de l'équipe d'ingénierie de Blue alors qu'elle explique comment construire une fonctionnalité de catégorisation et de marquage automatique alimentée par l'IA.
category: "Engineering"
date: 2024-07-25
---


La gestion efficace des projets et des processus est cruciale pour les organisations de toutes tailles.

Chez Blue, [nous avons fait de notre mission](/about) d'organiser le travail dans le monde en construisant la meilleure plateforme de gestion de projet sur la planète : simple, puissante, flexible et abordable pour tous.

Cela signifie que notre plateforme doit s'adapter aux besoins uniques de chaque équipe. Aujourd'hui, nous sommes ravis de lever le voile sur l'une de nos fonctionnalités les plus puissantes : les Permissions Personnalisées.

Les outils de gestion de projet sont la colonne vertébrale des flux de travail modernes, abritant des données sensibles, des communications cruciales et des plans stratégiques. En tant que tel, la capacité à contrôler finement l'accès à ces informations n'est pas seulement un luxe, mais une nécessité.

<video autoplay loop muted playsinline>
  <source src="/videos/user-roles.mp4" type="video/mp4">
</video>

Les permissions personnalisées jouent un rôle critique dans les plateformes SaaS B2B, en particulier dans les outils de gestion de projet, où l'équilibre entre collaboration et sécurité peut faire ou défaire le succès d'un projet.

Mais voici où Blue adopte une approche différente : **nous croyons que les fonctionnalités de niveau entreprise ne devraient pas être réservées aux budgets de grande entreprise.**

À une époque où l'IA permet aux petites équipes d'opérer à des échelles sans précédent, pourquoi la sécurité robuste et la personnalisation devraient-elles être hors de portée ?

Dans ce regard en coulisses, nous allons explorer comment nous avons développé notre fonctionnalité de Permissions Personnalisées, défiant le statu quo des niveaux de prix SaaS et apportant des options de sécurité puissantes et flexibles aux entreprises de toutes tailles.

Que vous soyez une startup avec de grands rêves ou un acteur établi cherchant à optimiser vos processus, les permissions personnalisées peuvent permettre de nouveaux cas d'utilisation que vous n'auriez jamais cru possibles.

## Comprendre les Permissions Utilisateur Personnalisées

Avant de plonger dans notre parcours de développement des permissions personnalisées pour Blue, prenons un moment pour comprendre ce que sont les permissions utilisateur personnalisées et pourquoi elles sont si cruciales dans les logiciels de gestion de projet.

Les permissions utilisateur personnalisées font référence à la capacité d'adapter les droits d'accès pour des utilisateurs ou des groupes individuels au sein d'un système logiciel. Au lieu de s'appuyer sur des rôles prédéfinis avec des ensembles de permissions fixes, les permissions personnalisées permettent aux administrateurs de créer des profils d'accès très spécifiques qui s'alignent parfaitement avec la structure et les besoins de flux de travail de leur organisation.

Dans le contexte des logiciels de gestion de projet comme Blue, les permissions personnalisées incluent :

* **Contrôle d'accès granulaire** : Déterminer qui peut voir, modifier ou supprimer des types spécifiques de données de projet.
* **Restrictions basées sur les fonctionnalités** : Activer ou désactiver certaines fonctionnalités pour des utilisateurs ou des équipes particulières.
* **Niveaux de sensibilité des données** : Définir des niveaux d'accès variés à des informations sensibles au sein des projets.
* **Permissions spécifiques aux flux de travail** : Aligner les capacités des utilisateurs avec des étapes ou des aspects spécifiques de votre flux de travail de projet.

L'importance des permissions personnalisées dans la gestion de projet ne peut être sous-estimée :

* **Sécurité améliorée** : En fournissant aux utilisateurs uniquement l'accès dont ils ont besoin, vous réduisez le risque de violations de données ou de modifications non autorisées.
* **Conformité améliorée** : Les permissions personnalisées aident les organisations à répondre aux exigences réglementaires spécifiques à l'industrie en contrôlant l'accès aux données.
* **Collaboration rationalisée** : Les équipes peuvent travailler plus efficacement lorsque chaque membre a le bon niveau d'accès pour effectuer son rôle sans restrictions inutiles ou privilèges écrasants.
* **Flexibilité pour les organisations complexes** : À mesure que les entreprises grandissent et évoluent, les permissions personnalisées permettent au logiciel de s'adapter aux structures et processus organisationnels changeants.

## Obtenir un OUI

[Nous avons déjà écrit](/insights/value-proposition-blue) que chaque fonctionnalité de Blue doit être un **OUI** catégorique avant que nous décidions de la construire. Nous n'avons pas le luxe de centaines d'ingénieurs et de perdre du temps et de l'argent à construire des choses dont les clients n'ont pas besoin.

Ainsi, le chemin pour mettre en œuvre des permissions personnalisées dans Blue n'a pas été une ligne droite. Comme beaucoup de fonctionnalités puissantes, il a commencé par un besoin clair de nos utilisateurs et a évolué grâce à une réflexion et une planification minutieuses.

Pendant des années, nos clients avaient demandé un contrôle plus granulaire sur les permissions utilisateur. Alors que des organisations de toutes tailles commençaient à gérer des projets de plus en plus complexes et sensibles, les limitations de notre contrôle d'accès basé sur des rôles standard sont devenues évidentes.

Des petites startups travaillant avec des clients externes, des entreprises de taille intermédiaire avec des processus d'approbation complexes, et de grandes entreprises avec des exigences de conformité strictes ont tous exprimé le même besoin :

Plus de flexibilité dans la gestion de l'accès des utilisateurs.

Malgré la demande claire, nous avons d'abord hésité à nous plonger dans le développement des permissions personnalisées.

Pourquoi ?

Nous comprenions la complexité impliquée !

Les permissions personnalisées touchent chaque partie d'un système de gestion de projet, de l'interface utilisateur à la structure de la base de données. Nous savions que la mise en œuvre de cette fonctionnalité nécessiterait des changements significatifs dans notre architecture de base et une réflexion minutieuse sur les implications en matière de performance.

En examinant le paysage, nous avons remarqué que très peu de nos concurrents avaient tenté de mettre en œuvre un moteur de permissions personnalisées puissant comme celui que nos clients demandaient. Ceux qui l'ont fait le réservaient souvent à leurs plans d'entreprise de plus haut niveau.

Il est devenu clair pourquoi : l'effort de développement est substantiel, et les enjeux sont élevés.

Mettre en œuvre des permissions personnalisées de manière incorrecte pourrait introduire des bogues critiques ou des vulnérabilités de sécurité, compromettant potentiellement l'ensemble du système. Cette réalisation a souligné l'ampleur du défi que nous envisagions.

### Défier le Statu Quo

Cependant, alors que nous continuions à croître et à évoluer, nous avons atteint une réalisation décisive :

**Le modèle SaaS traditionnel qui réserve des fonctionnalités puissantes aux clients d'entreprise n'a plus de sens dans le paysage commercial d'aujourd'hui.**

En 2024, avec la puissance de l'IA et des outils avancés, les petites équipes peuvent opérer à une échelle et une complexité qui rivalisent avec celles d'organisations beaucoup plus grandes. Une startup pourrait gérer des données sensibles de clients dans plusieurs pays. Une petite agence de marketing pourrait jongler avec des dizaines de projets clients ayant des exigences de confidentialité variées. Ces entreprises ont besoin du même niveau de sécurité et de personnalisation que *n'importe quelle* grande entreprise.

Nous nous sommes demandé : pourquoi la taille de la main-d'œuvre ou du budget d'une entreprise devrait-elle déterminer sa capacité à garder ses données en sécurité et ses processus efficaces ?

### Fonctionnalités de Niveau Entreprise pour Tous

Cette réalisation nous a conduits à une philosophie fondamentale qui guide désormais une grande partie de notre développement chez Blue : les fonctionnalités de niveau entreprise devraient être accessibles aux entreprises de toutes tailles.

Nous croyons que :

- **La sécurité ne devrait pas être un luxe.** Chaque entreprise, quelle que soit sa taille, mérite les outils pour protéger ses données et ses processus.
- **La flexibilité stimule l'innovation.** En donnant à tous nos utilisateurs des outils puissants, nous leur permettons de créer des flux de travail et des systèmes qui font avancer leurs industries.
- **La croissance ne devrait pas nécessiter de changements de plateforme.** À mesure que nos clients grandissent, leurs outils devraient évoluer sans heurts avec eux.

Avec cet état d'esprit, nous avons décidé de relever le défi des permissions personnalisées de front, engagés à les rendre disponibles à tous nos utilisateurs, pas seulement à ceux des plans de niveau supérieur.

Cette décision nous a mis sur la voie d'une conception soignée, d'un développement itératif et d'un retour d'expérience continu des utilisateurs qui ont finalement conduit à la fonctionnalité de permissions personnalisées dont nous sommes fiers aujourd'hui.

Dans la section suivante, nous allons plonger dans la façon dont nous avons abordé le processus de conception et de développement pour donner vie à cette fonctionnalité complexe.

### Conception et Développement

Lorsque nous avons décidé de nous attaquer aux permissions personnalisées, nous avons rapidement réalisé que nous faisions face à une tâche colossale.

À première vue, "permissions personnalisées" peut sembler simple, mais c'est une fonctionnalité trompeusement complexe qui touche chaque aspect de notre système.

Le défi était redoutable : nous devions mettre en œuvre des permissions en cascade, permettre des modifications à la volée, apporter des changements significatifs au schéma de la base de données et garantir une fonctionnalité fluide à travers notre écosystème entier – applications web, Mac, Windows, iOS et Android, ainsi que notre API et nos webhooks.

La complexité était suffisante pour faire hésiter même les développeurs les plus expérimentés.

Notre approche était centrée sur deux principes clés :

1. Décomposer la fonctionnalité en versions gérables
2. Adopter un envoi incrémental.

Face à la complexité des permissions personnalisées à grande échelle, nous nous sommes posé une question cruciale :

> Quelle serait la première version la plus simple possible de cette fonctionnalité ?

Cette approche s'aligne avec le principe agile de livrer un Produit Minimum Viable (MVP) et d'itérer en fonction des retours.

Notre réponse était agréablement simple :

1. Introduire un interrupteur pour masquer l'onglet d'activité du projet
2. Ajouter un autre interrupteur pour masquer l'onglet des formulaires

**C'était tout.**

Pas de cloches ni de sifflets, pas de matrices de permissions complexes—juste deux simples interrupteurs marche/arrêt.

Bien que cela puisse sembler peu impressionnant à première vue, cette approche offrait plusieurs avantages significatifs :

* **Mise en œuvre rapide** : Ces simples interrupteurs pouvaient être développés et testés rapidement, nous permettant de mettre une version de base des permissions personnalisées entre les mains des utilisateurs rapidement.
* **Valeur utilisateur claire** : Même avec juste ces deux options, nous fournissions une valeur tangible. Certaines équipes pourraient vouloir masquer le fil d'activité des clients, tandis que d'autres pourraient avoir besoin de restreindre l'accès aux formulaires pour certains groupes d'utilisateurs.
* **Fondation pour la croissance** : Ce simple départ a posé les bases pour des permissions plus complexes. Cela nous a permis de mettre en place l'infrastructure de base pour les permissions personnalisées sans nous enliser dans la complexité dès le départ.
* **Retour d'expérience utilisateur** : En publiant cette version simple, nous pouvions recueillir des retours du monde réel sur la façon dont les utilisateurs interagissaient avec les permissions personnalisées, informant notre développement futur.
* **Apprentissage technique** : Cette mise en œuvre initiale a donné à notre équipe de développement une expérience pratique dans la modification des permissions à travers notre plateforme, nous préparant à des itérations plus complexes.

Et vous savez, c'est en fait assez humblant d'avoir une grande vision pour quelque chose, puis de livrer quelque chose qui est un si petit pourcentage de cette vision.

Après avoir livré ces deux premiers interrupteurs, nous avons décidé de nous attaquer à quelque chose de plus sophistiqué. Nous avons abouti à deux nouvelles permissions de rôle utilisateur personnalisées.

La première était la capacité de limiter les utilisateurs à ne voir que les enregistrements qui leur ont été spécifiquement assignés. Cela est très utile si vous avez un client dans un projet et que vous ne voulez que lui montrer les enregistrements qui lui sont spécifiquement assignés au lieu de tout ce sur quoi vous travaillez pour lui.

La seconde était une option pour les administrateurs de projet de bloquer les groupes d'utilisateurs pour qu'ils ne puissent pas inviter d'autres utilisateurs. Cela est bon si vous avez un projet sensible que vous souhaitez garantir qu'il reste sur une base de "besoin de voir".

Une fois que nous avons livré cela, nous avons gagné en confiance et pour notre troisième version, nous avons abordé les permissions au niveau des colonnes, ce qui signifie être capable de décider quels champs personnalisés un groupe d'utilisateurs spécifique peut voir ou modifier.

C'est extrêmement puissant. Imaginez que vous avez un projet CRM, et que vous avez des données qui ne sont pas seulement liées aux montants que le client paiera, mais aussi à vos coûts et marges bénéficiaires. Vous ne voudriez peut-être pas que vos champs de coûts et votre champ de formule de marge de projet soient visibles par le personnel junior, et les permissions personnalisées vous permettent de verrouiller ces champs pour qu'ils ne soient pas affichés.

Ensuite, nous sommes passés à la création de permissions basées sur des listes, où les administrateurs de projet peuvent décider si un groupe d'utilisateurs peut voir, modifier et supprimer une liste spécifique. S'ils masquent une liste, tous les enregistrements à l'intérieur de cette liste deviennent également masqués, ce qui est génial car cela signifie que vous pouvez cacher certaines parties de votre processus à vos membres d'équipe ou clients.

Voici le résultat final :

<video autoplay loop muted playsinline>
  <source src="/videos/custom-user-roles.mp4" type="video/mp4">
</video>

## Considérations Techniques

Au cœur de l'architecture technique de Blue se trouve GraphQL, un choix clé qui a considérablement influencé notre capacité à mettre en œuvre des fonctionnalités complexes comme les permissions personnalisées. Mais avant de plonger dans les détails, faisons un pas en arrière et comprenons ce qu'est GraphQL et comment il diffère de l'approche API REST plus traditionnelle.
GraphQL vs API REST : Une Explication Accessible

Imaginez que vous êtes dans un restaurant. Avec une API REST, c'est comme commander à partir d'un menu fixe. Vous demandez un plat spécifique (point de terminaison), et vous obtenez tout ce qui l'accompagne, que vous le vouliez ou non. Si vous souhaitez personnaliser votre repas, vous pourriez avoir besoin de passer plusieurs commandes (appels API) ou de demander un plat spécialement préparé (point de terminaison personnalisé).

GraphQL, en revanche, est comme avoir une conversation avec un chef qui peut préparer n'importe quoi. Vous dites au chef exactement quels ingrédients vous voulez (champs de données), et en quelles quantités. Le chef prépare alors un plat qui est précisément ce que vous avez demandé - ni plus, ni moins. C'est essentiellement ce que fait GraphQL - il permet au client de demander exactement les données dont il a besoin, et le serveur fournit juste cela.

### Un Déjeuner Important

Environ six semaines après le développement initial de Blue, notre ingénieur en chef et PDG est sorti déjeuner.

Le sujet de discussion ?

La question de savoir s'il fallait passer des API REST à GraphQL. Ce n'était pas une décision à prendre à la légère - adopter GraphQL signifierait abandonner six semaines de travail initial.

En marchant de retour au bureau, le PDG a posé une question cruciale à l'ingénieur en chef : "Regretterions-nous de ne pas avoir fait cela dans cinq ans ?"

La réponse est devenue claire : GraphQL était la voie à suivre.

Nous avons reconnu le potentiel de cette technologie dès le départ, voyant comment elle pouvait soutenir notre vision d'une plateforme de gestion de projet flexible et puissante.

Notre prévoyance dans l'adoption de GraphQL a porté ses fruits lorsqu'il s'est agi de mettre en œuvre des permissions personnalisées. Avec une API REST, nous aurions eu besoin d'un point de terminaison différent pour chaque configuration possible de permissions personnalisées - une approche qui serait rapidement devenue ingérable et difficile à maintenir.

GraphQL, en revanche, nous permet de gérer les permissions personnalisées de manière dynamique. Voici comment cela fonctionne :

- **Vérifications de permissions à la volée** : Lorsqu'un client fait une demande, notre serveur GraphQL peut vérifier les permissions de l'utilisateur directement depuis notre base de données.
- **Récupération de données précise** : En fonction de ces permissions, GraphQL ne renvoie que les données demandées qui correspondent aux droits d'accès de l'utilisateur.
- **Requêtes flexibles** : À mesure que les permissions changent, nous n'avons pas besoin de créer de nouveaux points de terminaison ou de modifier les existants. La même requête GraphQL peut s'adapter à différentes configurations de permissions.
- **Récupération de données efficace** : GraphQL permet aux clients de demander exactement ce dont ils ont besoin. Cela signifie que nous ne surchargeons pas les données, ce qui pourrait potentiellement exposer des informations auxquelles l'utilisateur ne devrait pas avoir accès.

Cette flexibilité est cruciale pour une fonctionnalité aussi complexe que les permissions personnalisées. Elle nous permet d'offrir un contrôle granulaire *sans* sacrifier la performance ou la maintenabilité.

## Défis

La mise en œuvre des permissions personnalisées dans Blue a apporté son lot de défis, chacun nous poussant à innover et à affiner notre approche. L'optimisation des performances est rapidement devenue une préoccupation critique. À mesure que nous ajoutions des vérifications de permissions plus granulaires, nous risquions de ralentir notre système, en particulier pour les grands projets avec de nombreux utilisateurs et des configurations de permissions complexes. Pour y remédier, nous avons mis en œuvre une stratégie de mise en cache multi-niveaux, optimisé nos requêtes de base de données et tiré parti de la capacité de GraphQL à demander uniquement les données nécessaires. Cette approche nous a permis de maintenir des temps de réponse rapides même à mesure que les projets se développaient et que la complexité des permissions augmentait.

L'interface utilisateur pour les permissions personnalisées a présenté un autre obstacle significatif. Nous devions rendre l'interface intuitive et gérable pour les administrateurs, même en ajoutant plus d'options et en augmentant la complexité du système.

Notre solution a impliqué plusieurs cycles de tests utilisateurs et de conception itérative.

Nous avons introduit une matrice de permissions visuelle qui permettait aux administrateurs de visualiser et de modifier rapidement les permissions à travers différents rôles et domaines de projet.

Assurer la cohérence entre les plateformes a présenté son propre ensemble de défis. Nous devions mettre en œuvre les permissions personnalisées de manière uniforme à travers nos applications web, de bureau et mobiles, chacune ayant ses propres considérations d'interface et d'expérience utilisateur. Cela était particulièrement délicat pour nos applications mobiles, qui devaient masquer et afficher dynamiquement des fonctionnalités en fonction des permissions de l'utilisateur. Nous avons abordé cela en centralisant notre logique de permissions dans la couche API, garantissant que toutes les plateformes reçoivent des données de permissions cohérentes.

Ensuite, nous avons développé un cadre UI flexible qui pouvait s'adapter à ces changements de permissions en temps réel, offrant une expérience fluide, quelle que soit la plateforme utilisée.

L'éducation et l'adoption des utilisateurs ont présenté le dernier obstacle dans notre parcours de permissions personnalisées. L'introduction d'une fonctionnalité aussi puissante signifiait que nous devions aider nos utilisateurs à comprendre et à exploiter efficacement les permissions personnalisées.

Nous avons initialement lancé les permissions personnalisées à un sous-ensemble de notre base d'utilisateurs, surveillant attentivement leurs expériences et collectant des informations. Cette approche nous a permis d'affiner la fonctionnalité et nos matériaux éducatifs en fonction de l'utilisation réelle avant de lancer à l'ensemble de notre base d'utilisateurs.

Le déploiement par phases s'est avéré inestimable, nous aidant à identifier et à résoudre des problèmes mineurs et des points de confusion des utilisateurs que nous n'avions pas anticipés, menant finalement à une fonctionnalité plus polie et conviviale pour tous nos utilisateurs.

Cette approche de lancement à un sous-ensemble d'utilisateurs, ainsi que notre période de "Bêta" typiquement de 2 à 3 semaines sur notre Bêta publique, nous aide à dormir sur nos deux oreilles. :)

## Perspectives

Comme pour toutes les fonctionnalités, rien n'est jamais *"terminé"*.

Notre vision à long terme pour la fonctionnalité de permissions personnalisées s'étend à des balises, des filtres de champs personnalisés, une navigation de projet personnalisable et des contrôles de commentaires.

Plongeons dans chaque aspect.

### Permissions de Balises

Nous pensons qu'il serait incroyable de pouvoir créer des permissions basées sur le fait qu'un enregistrement ait une ou plusieurs balises. Le cas d'utilisation le plus évident serait de créer un rôle utilisateur personnalisé appelé "Clients" et de ne permettre qu'aux utilisateurs de ce rôle de voir les enregistrements qui ont la balise "Clients".

Cela vous donne une vue d'ensemble de la possibilité ou non pour vos clients de voir un enregistrement.

Cela pourrait devenir encore plus puissant avec des combinatoires ET/OU, où vous pouvez spécifier des règles plus complexes. Par exemple, vous pourriez établir une règle qui permet l'accès aux enregistrements étiquetés à la fois "Clients" ET "Public", ou aux enregistrements étiquetés soit "Interne" SOIT "Confidentiel". Ce niveau de flexibilité permettrait des paramètres de permissions incroyablement nuancés, répondant même aux structures organisationnelles et aux flux de travail les plus complexes.

Les applications potentielles sont vastes. Les chefs de projet pourraient facilement séparer les informations sensibles, les équipes de vente pourraient avoir un accès automatique aux données clients pertinentes, et les collaborateurs externes pourraient être intégrés sans effort dans des parties spécifiques d'un projet sans risquer d'exposer des informations internes sensibles.

### Filtres de Champs Personnalisés

Notre vision pour les Filtres de Champs Personnalisés représente un saut significatif en avant dans le contrôle d'accès granulaire. Cette fonctionnalité permettra aux administrateurs de projet de définir quels enregistrements des groupes d'utilisateurs spécifiques peuvent voir en fonction des valeurs des champs personnalisés. Il s'agit de créer des limites dynamiques, basées sur les données, pour l'accès à l'information.

Imaginez pouvoir établir des permissions comme ceci :

- Afficher uniquement les enregistrements où le menu déroulant "Statut du Projet" est défini sur "Public"
- Restreindre la visibilité aux éléments où le champ multi-sélection "Département" inclut "Marketing"
- Autoriser l'accès aux tâches où la case à cocher "Priorité" est cochée
- Afficher les projets où le champ numérique "Budget" est au-dessus d'un certain seuil

### Navigation de Projet Personnalisable

Ceci est simplement une extension des interrupteurs que nous avons déjà. Au lieu d'avoir simplement des interrupteurs pour "activité" et "formulaires", nous voulons étendre cela à chaque partie de la navigation du projet. De cette façon, les administrateurs de projet peuvent créer des interfaces ciblées et supprimer des outils dont ils n'ont pas besoin.

### Contrôles de Commentaires

À l'avenir, nous voulons être créatifs dans la façon dont nous permettons à nos clients de décider qui peut et ne peut pas voir les commentaires. Nous pourrions permettre plusieurs zones de commentaires sous un même enregistrement, et chacune peut être visible ou non visible pour différents groupes d'utilisateurs.

De plus, nous pourrions également permettre une fonctionnalité où seuls les commentaires où un utilisateur est *spécifiquement* mentionné sont visibles, et rien d'autre ne l'est. Cela permettrait aux équipes ayant des clients sur des projets de s'assurer que seuls les commentaires qu'elles souhaitent que les clients voient sont visibles.

## Conclusion

Voilà, c'est ainsi que nous avons abordé la construction de l'une des fonctionnalités les plus intéressantes et puissantes ! [Comme vous pouvez le voir sur notre outil de comparaison de gestion de projet](/compare), très peu de systèmes de gestion de projet ont une configuration de matrice de permissions aussi puissante, et ceux qui le font la réservent à leurs plans d'entreprise les plus coûteux, la rendant inaccessible à une entreprise typique de petite ou moyenne taille.

Avec Blue, vous avez *toutes* les fonctionnalités disponibles avec notre plan — nous ne croyons pas que les fonctionnalités de niveau entreprise devraient être réservées aux clients d'entreprise !
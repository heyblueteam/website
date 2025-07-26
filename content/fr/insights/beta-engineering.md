---
title: Pourquoi Blue a une beta ouverte
description: Découvrez pourquoi notre système de gestion de projet a une beta ouverte en cours.
category: "Engineering"
date: 2024-08-03
---


De nombreuses startups B2B SaaS lancent en beta, et pour de bonnes raisons. C'est une partie du motto traditionnel de la Silicon Valley *“aller vite et casser des choses”*.

Mettre un autocollant “beta” sur un produit réduit les attentes.

Quelque chose ne fonctionne pas ? Oh eh bien, ce n'est qu'une beta.

Le système est lent ? Oh eh bien, ce n'est qu'une beta.

[La documentation](https://blue.cc/docs) est inexistante ? Oh eh bien…vous avez compris.

Et c'est *en fait* une bonne chose. Reid Hoffman, le fondateur de LinkedIn, a dit célèbrement :

> Si vous n'êtes pas gêné par la première version de votre produit, vous avez lancé trop tard.

Et l'autocollant beta est également bénéfique pour les clients. Cela les aide à se sélectionner eux-mêmes.

Les clients qui essaient des produits beta sont ceux qui se trouvent aux premiers stades du cycle d'adoption technologique, également connu sous le nom de courbe d'adoption des produits.

Le cycle d'adoption technologique est généralement divisé en cinq segments principaux :

1. Innovateurs
2. Premiers adopteurs
3. Majorité précoce
4. Majorité tardive
5. Retardataires

![](/insights/technology-adoption-lifecycle-graph.png)

Cependant, à un moment donné, le produit doit mûrir, et les clients s'attendent à un produit stable et fonctionnel. Ils ne veulent pas avoir accès à un environnement “beta” où les choses se cassent.

Ou peut-être le veulent-ils ?

*C'est* la question que nous nous sommes posée.

Nous pensons que nous nous sommes posés cette question en raison de la manière dont Blue a été initialement construit. [Blue a commencé comme une branche d'une agence de design très occupée](/insights/agency-success-playbook), et nous avons donc travaillé *à l'intérieur* du bureau d'une entreprise qui utilisait activement Blue pour gérer tous ses projets.

Cela signifie que pendant des années, nous avons pu observer comment de *vrais* êtres humains — assis juste à côté de nous ! — utilisaient Blue dans leur vie quotidienne.

Et parce qu'ils utilisaient Blue depuis les premiers jours, cette équipe a toujours utilisé Blue Beta !

Il était donc naturel pour nous de permettre à tous nos autres clients de l'utiliser également.

**Et c'est pourquoi nous n'avons pas d'équipe de test dédiée.**

C'est exact.

Personne chez Blue n'a la *seule* responsabilité de s'assurer que notre plateforme fonctionne bien et est stable.

C'est pour plusieurs raisons.

La première est une base de coûts plus faible.

Ne pas avoir une équipe de test à plein temps réduit considérablement nos coûts, et nous sommes en mesure de transmettre ces économies à nos clients avec les prix les plus bas du secteur.

Pour mettre cela en perspective, nous offrons des ensembles de fonctionnalités de niveau entreprise que notre concurrence facture entre 30 et 55 $/utilisateur/mois pour seulement 7 $/mois.

Cela ne se produit pas par accident, c'est *intentionnel*.

Cependant, ce n'est pas une bonne stratégie de vendre un produit moins cher s'il ne fonctionne pas.

Donc la *vraie question est*, comment parvenons-nous à créer une plateforme stable que des milliers de clients peuvent utiliser sans une équipe de test dédiée ?

Bien sûr, notre approche d'avoir une beta ouverte est cruciale pour cela, mais avant d'entrer dans le vif du sujet, nous voulons aborder la responsabilité des développeurs.

Nous avons pris la décision dès le départ chez Blue de ne jamais séparer les responsabilités pour les technologies front-end et back-end. Nous n'engagerions ou ne formerions que des développeurs full stack.

La raison pour laquelle nous avons pris cette décision était de garantir qu'un développeur possède pleinement la fonctionnalité sur laquelle il travaillait. Ainsi, il n'y aurait pas de mentalité *“jeter le problème par-dessus la clôture du jardin”* que l'on obtient parfois lorsque les responsabilités pour les fonctionnalités sont partagées.

Et cela s'étend à la test de la fonctionnalité, à la compréhension des cas d'utilisation des clients et des demandes, et à la lecture et aux commentaires sur les spécifications.

En d'autres termes, chaque développeur construit une compréhension profonde et intuitive de la fonctionnalité qu'il développe.

D'accord, parlons maintenant de notre beta ouverte.

Quand nous disons qu'elle est “ouverte” — nous le pensons. Tout client peut l'essayer simplement en ajoutant “beta” devant l'URL de notre application web.

Ainsi, “app.blue.cc” devient “beta.app.blue.cc”.

Lorsqu'ils font cela, ils peuvent voir leurs données habituelles, car les environnements Beta et Production partagent la même base de données, mais ils pourront également voir de nouvelles fonctionnalités.

Les clients peuvent facilement travailler même s'ils ont certains membres de l'équipe en Production et d'autres curieux en Beta.

Nous avons généralement quelques centaines de clients utilisant Beta à tout moment, et nous publions des aperçus de fonctionnalités sur nos forums communautaires afin qu'ils puissent découvrir ce qui est nouveau et l'essayer.

Et c'est le point : nous avons *plusieurs centaines* de testeurs !

Tous ces clients essaieront des fonctionnalités dans leurs flux de travail et seront assez vocaux si quelque chose ne va pas, car ils sont *déjà* en train d'implémenter la fonctionnalité dans leur entreprise !

Les retours les plus courants sont de petits mais très utiles changements qui traitent des cas particuliers que nous n'avions pas envisagés.

Nous laissons les nouvelles fonctionnalités en Beta entre 2 et 4 semaines. Chaque fois que nous estimons qu'elles sont stables, nous les publions en production.

Nous avons également la possibilité de contourner Beta si nécessaire, en utilisant un drapeau de voie rapide. Cela se fait généralement pour des corrections de bogues que nous ne voulons pas retenir pendant 2 à 4 semaines avant de les expédier en production.

Le résultat ?

Pousser en production semble… eh bien ennuyeux ! Comme rien — ce n'est tout simplement pas un gros problème pour nous.

Et cela signifie que cela lisse notre calendrier de publication, ce qui nous a permis de [livrer des fonctionnalités mensuellement comme une horloge pendant les six dernières années.](/changelog).

Cependant, comme tout choix, il y a des compromis.

Le support client est légèrement plus complexe, car nous devons soutenir les clients à travers deux versions de notre plateforme. Parfois, cela peut causer de la confusion chez les clients qui ont des membres d'équipe utilisant deux versions différentes.

Un autre point de douleur est que cette approche peut parfois ralentir l'ensemble du calendrier de publication en production. Cela est particulièrement vrai pour les fonctionnalités plus importantes qui peuvent se “bloquer” en Beta s'il y a une autre fonctionnalité connexe qui rencontre des problèmes et nécessite des travaux supplémentaires.

Mais dans l'ensemble, nous pensons que ces compromis valent les avantages d'une base de coûts plus faible et d'un plus grand engagement des clients.

Nous sommes l'une des rares entreprises de logiciels à adopter cette approche, mais elle fait désormais partie intégrante de notre approche de développement de produits.
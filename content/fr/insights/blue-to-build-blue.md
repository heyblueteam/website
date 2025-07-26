---
title: Comment nous utilisons Blue pour construire Blue.
description: Découvrez comment nous utilisons notre propre plateforme de gestion de projet pour construire notre plateforme de gestion de projet !
category: "CEO Blog"
date: 2024-08-07
---


Vous êtes sur le point de découvrir comment Blue construit Blue.

Chez Blue, nous utilisons notre propre produit.

Cela signifie que nous utilisons Blue pour *construire* Blue.

Ce terme étrange, souvent appelé "dogfooding", est souvent attribué à Paul Maritz, un responsable chez Microsoft dans les années 1980. Il aurait envoyé un e-mail avec l'objet *"Manger notre propre nourriture pour chien"* pour encourager les employés de Microsoft à utiliser les produits de l'entreprise.

L'idée d'utiliser vos propres outils pour construire vos outils est qu'elle conduit à un cycle de rétroaction positif.

L'idée d'utiliser vos propres outils pour construire vos outils conduit à un cycle de rétroaction positif, créant de nombreux avantages :

- **Cela nous aide à identifier rapidement les problèmes d'utilisabilité dans le monde réel.** Comme nous utilisons Blue quotidiennement, nous rencontrons les mêmes défis que nos utilisateurs, ce qui nous permet de les aborder de manière proactive.
- **Cela accélère la découverte de bugs.** L'utilisation interne révèle souvent des bugs avant qu'ils n'atteignent nos clients, améliorant ainsi la qualité globale du produit.
- **Cela renforce notre empathie pour les utilisateurs finaux.** Notre équipe acquiert une expérience directe des forces et des faiblesses de Blue, ce qui nous aide à prendre des décisions plus centrées sur l'utilisateur.
- **Cela favorise une culture de qualité au sein de notre organisation.** Lorsque tout le monde utilise le produit, il y a un intérêt commun pour son excellence.
- **Cela stimule l'innovation.** L'utilisation régulière suscite souvent des idées pour de nouvelles fonctionnalités ou améliorations, maintenant Blue à la pointe.

[Nous avons déjà parlé de pourquoi nous n'avons pas d'équipe de test dédiée](/insights/open-beta) et c'est encore une autre raison.

S'il y a des bugs dans notre système, nous les trouvons presque toujours dans notre utilisation quotidienne constante de la plateforme. Et cela crée également une fonction de contrainte pour les corriger, car nous les trouverons évidemment très ennuyeux étant donné que nous sommes probablement l'un des principaux utilisateurs de Blue !

Cette approche démontre notre engagement envers le produit. En nous appuyant sur Blue nous-mêmes, nous montrons à nos clients que nous croyons vraiment en ce que nous construisons. Ce n'est pas seulement un produit que nous vendons - c'est un outil dont nous dépendons chaque jour.

## Processus Principal

Nous avons un projet dans Blue, judicieusement nommé "Produit".

**Tout** ce qui concerne notre développement de produit est suivi ici. Feedback des clients, bugs, idées de fonctionnalités, travaux en cours, etc. L'idée d'avoir un projet où nous suivons tout est que cela [favorise un meilleur travail d'équipe.](/insights/great-teamwork)

Chaque enregistrement est une fonctionnalité ou une partie d'une fonctionnalité. C'est ainsi que nous passons de "ce serait génial si..." à "regardez cette nouvelle fonctionnalité incroyable !"

Le projet a les listes suivantes :

- **Idées/Feedback** : C'est une liste d'idées de l'équipe ou de feedback des clients basé sur des appels ou des échanges d'e-mails. N'hésitez pas à ajouter des idées ici ! Dans cette liste, nous n'avons pas encore décidé que nous allons construire l'une de ces fonctionnalités, mais nous la passons régulièrement en revue pour des idées que nous souhaitons explorer davantage.
- **Backlog (Long Terme)** : C'est là que les fonctionnalités de la liste Idées/Feedback vont si nous décidons qu'elles seraient un bon ajout à Blue.
- **{Trimestre Actuel}** : Cela est généralement structuré comme "Qx AAAA" et montre nos priorités trimestrielles.
- **Bugs** : C'est une liste de bugs connus signalés par l'équipe ou les clients. Les bugs ajoutés ici auront automatiquement le tag "Bug" ajouté.
- **Spécifications** : Ces fonctionnalités sont actuellement en cours de spécification. Toutes les fonctionnalités ne nécessitent pas de spécification ou de conception ; cela dépend de la taille attendue de la fonctionnalité et du niveau de confiance que nous avons concernant les cas limites et la complexité.
- **Backlog de Design** : C'est le backlog pour les designers, chaque fois qu'ils ont terminé quelque chose qui est en cours, ils peuvent choisir n'importe quel élément de cette liste.
- **Design en Cours** : Ce sont les fonctionnalités que les designers sont en train de concevoir.
- **Revue de Design** : C'est là que les fonctionnalités dont les conceptions sont actuellement en cours de révision.
- **Backlog (Court Terme)** : C'est une liste de fonctionnalités sur lesquelles nous allons probablement commencer à travailler dans les prochaines semaines. C'est là que les affectations ont lieu. Le PDG et le Responsable de l'Ingénierie décident quelles fonctionnalités sont attribuées à quel ingénieur en fonction de l'expérience précédente et de la charge de travail. [Les membres de l'équipe peuvent ensuite les intégrer dans le En Cours](/insights/push-vs-pull-kanban) une fois qu'ils ont terminé leur travail actuel.
- **En Cours** : Ce sont des fonctionnalités qui sont actuellement en cours de développement.
- **Revue de Code** : Une fois qu'une fonctionnalité a terminé son développement, elle subit une revue de code. Ensuite, elle sera soit renvoyée à "En Cours" si des ajustements sont nécessaires, soit déployée dans l'environnement de Développement.
- **Dev** : Ce sont toutes les fonctionnalités actuellement dans l'environnement de Développement. D'autres membres de l'équipe et certains clients peuvent les examiner.
- **Beta** : Ce sont toutes les fonctionnalités actuellement dans l'[environnement Beta](https://beta.app.blue.cc). De nombreux clients utilisent cela comme leur plateforme Blue quotidienne et fourniront également des retours.
- **Production** : Lorsqu'une fonctionnalité atteint la production, elle est alors considérée comme terminée.

Parfois, alors que nous développons une fonctionnalité, nous réalisons que certaines sous-fonctionnalités sont plus difficiles à mettre en œuvre que prévu initialement, et nous pouvons choisir de ne pas les inclure dans la version initiale que nous déployons aux clients. Dans ce cas, nous pouvons créer un nouvel enregistrement avec un nom suivant le format "{NomDeLaFonctionnalité} V2" et inclure toutes les sous-fonctionnalités en tant qu'éléments de liste de contrôle.

## Tags

- **Mobile** : Cela signifie que la fonctionnalité est spécifique à nos applications iOS, Android ou iPad.
- **{NomDuClientEntreprise}** : Une fonctionnalité est spécifiquement développée pour un client entreprise. Le suivi est important car il y a généralement des accords commerciaux supplémentaires pour chaque fonctionnalité.
- **Bug** : Cela signifie qu'il s'agit d'un bug qui nécessite une correction.
- **Fast-Track** : Cela signifie qu'il s'agit d'un changement Fast-Track qui n'a pas besoin de passer par le cycle de publication complet décrit ci-dessus.
- **Principal** : Il s'agit d'un développement de fonctionnalité majeur. Il est généralement réservé aux travaux d'infrastructure majeurs, aux grandes mises à jour de dépendances et aux nouveaux modules significatifs au sein de Blue.
- **IA** : Cette fonctionnalité contient un composant d'intelligence artificielle.
- **Sécurité** : Cela signifie qu'une implication en matière de sécurité doit être examinée ou qu'un correctif est nécessaire.

Le tag fast-track est particulièrement intéressant. Cela est réservé aux mises à jour plus petites et moins complexes qui ne nécessitent pas notre cycle de publication complet, et que nous voulons expédier aux clients dans les 24 à 48 heures.

Les changements fast-track sont généralement des ajustements mineurs qui peuvent améliorer considérablement l'expérience utilisateur sans altérer la fonctionnalité de base. Pensez à corriger des fautes de frappe dans l'interface utilisateur, à ajuster le rembourrage des boutons ou à ajouter de nouvelles icônes pour une meilleure orientation visuelle. Ce sont le genre de changements qui, bien que petits, peuvent faire une grande différence dans la façon dont les utilisateurs perçoivent et interagissent avec notre produit. Ils sont également ennuyeux s'ils prennent du temps à être expédiés !

Notre processus fast-track est simple.

Il commence par créer une nouvelle branche à partir de la branche principale, mettre en œuvre les changements, puis créer des demandes de fusion pour chaque branche cible - Dev, Beta et Production. Nous générons un lien de prévisualisation pour révision, garantissant que même ces petits changements respectent nos normes de qualité. Une fois approuvés, les changements sont fusionnés simultanément dans toutes les branches, maintenant nos environnements synchronisés.

## Champs Personnalisés

Nous n'avons pas beaucoup de champs personnalisés dans notre projet Produit.

- **Spécifications** : Cela renvoie à un document Blue qui contient la spécification pour cette fonctionnalité particulière. Cela n'est pas toujours fait, car cela dépend de la complexité de la fonctionnalité.
- **MR** : C'est le lien vers la Demande de Fusion dans [Gitlab](https://gitlab.com) où nous hébergeons notre code.
- **Lien de Prévisualisation** : Pour les fonctionnalités qui changent principalement le front-end, nous pouvons créer une URL unique qui contient ces changements pour chaque commit, afin que nous puissions facilement examiner les changements.
- **Lead** : Ce champ nous indique quel ingénieur senior prend en charge la révision du code. Cela garantit que chaque fonctionnalité reçoit l'attention experte qu'elle mérite, et qu'il y a toujours une personne de référence claire pour les questions ou préoccupations.

## Listes de Contrôle

Lors de nos démonstrations hebdomadaires, nous déposerons le feedback discuté dans une liste de contrôle appelée "feedback" et il y aura également une autre liste de contrôle contenant le principal [WBS (Work Breakdown Scope)](/insights/simple-work-breakdown-structure) de la fonctionnalité, afin que nous puissions facilement savoir ce qui est fait et ce qu'il reste à faire.

## Conclusion

Et c'est tout !

Nous pensons que parfois les gens sont surpris par la simplicité de notre processus, mais nous croyons que des processus simples sont souvent bien supérieurs à des processus trop complexes que l'on ne peut pas facilement comprendre.

Cette simplicité est intentionnelle. Elle nous permet de rester agiles, de répondre rapidement aux besoins des clients et de garder toute notre équipe alignée.

En utilisant Blue pour construire Blue, nous ne développons pas seulement un produit – nous le vivons.

Alors la prochaine fois que vous utilisez Blue, rappelez-vous : vous n'utilisez pas seulement un produit que nous avons construit. Vous utilisez un produit dont nous dépendons personnellement chaque jour.

Et cela fait toute la différence.
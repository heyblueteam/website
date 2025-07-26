---
title: Catégorisation automatique par IA (Analyse technique approfondie)
category: "Engineering"
description: Découvrez les coulisses avec l'équipe d'ingénierie de Blue qui explique comment elle a développé une fonctionnalité de catégorisation et d'étiquetage automatique alimentée par l'IA.
date: 2024-12-07
---

Nous avons récemment lancé la [Catégorisation automatique par IA](/insights/ai-auto-categorization) pour tous les utilisateurs de Blue. Il s'agit d'une fonctionnalité IA intégrée à l'abonnement principal de Blue, sans coûts supplémentaires. Dans cet article, nous explorons l'ingénierie derrière la réalisation de cette fonctionnalité.

---
Chez Blue, notre approche du développement de fonctionnalités est ancrée dans une compréhension approfondie des besoins des utilisateurs et des tendances du marché, associée à un engagement à maintenir la simplicité et la facilité d'utilisation qui définissent notre plateforme. C'est ce qui guide notre [feuille de route](/platform/roadmap), et ce qui nous a [permis de livrer régulièrement des fonctionnalités chaque mois pendant des années](/platform/changelog).

L'introduction de l'étiquetage automatique alimenté par l'IA dans Blue est un exemple parfait de cette philosophie en action. Avant de plonger dans les détails techniques de la construction de cette fonctionnalité, il est crucial de comprendre le problème que nous résolvions et la réflexion approfondie qui a guidé son développement.

Le paysage de la gestion de projet évolue rapidement, les capacités d'IA devenant de plus en plus centrales dans les attentes des utilisateurs. Nos clients, en particulier ceux qui gèrent des [projets](/platform) à grande échelle avec des millions d'[enregistrements](/platform/features/records), avaient exprimé clairement leur désir de disposer de moyens plus intelligents et plus efficaces pour organiser et catégoriser leurs données.

Cependant, chez Blue, nous n'ajoutons pas simplement des fonctionnalités parce qu'elles sont tendance ou demandées. Notre philosophie est que chaque nouvel ajout doit prouver sa valeur, avec une réponse par défaut ferme *"non"* jusqu'à ce qu'une fonctionnalité démontre une forte demande et une utilité claire.

Pour vraiment comprendre la profondeur du problème et le potentiel de l'étiquetage automatique par IA, nous avons mené des entretiens approfondis avec les clients, en nous concentrant sur les utilisateurs de longue date qui gèrent des projets complexes et riches en données dans plusieurs domaines.

Ces conversations ont révélé un fil conducteur commun : *bien que l'étiquetage soit inestimable pour l'organisation et la recherche, la nature manuelle du processus devenait un goulot d'étranglement, en particulier pour les équipes traitant de gros volumes d'enregistrements.*

Mais nous avons vu au-delà de la simple résolution du problème immédiat de l'étiquetage manuel.

Nous avons envisagé un avenir où l'étiquetage alimenté par l'IA pourrait devenir la base de flux de travail plus intelligents et automatisés.

La véritable puissance de cette fonctionnalité, avons-nous réalisé, résidait dans son potentiel d'intégration avec notre [système d'automatisation de gestion de projet](/platform/features/automations). Imaginez un outil de gestion de projet qui non seulement catégorise intelligemment les informations, mais utilise également ces catégories pour router les tâches, déclencher des actions et adapter les flux de travail en temps réel.

Cette vision s'alignait parfaitement avec notre objectif de garder Blue simple mais puissant.

De plus, nous avons reconnu le potentiel d'étendre cette capacité au-delà des limites de notre plateforme. En développant un système d'étiquetage IA robuste, nous posions les bases d'une "API de catégorisation" qui pourrait fonctionner immédiatement, ouvrant potentiellement de nouvelles voies pour la façon dont nos utilisateurs interagissent avec et exploitent Blue dans leurs écosystèmes technologiques plus larges.

Cette fonctionnalité, par conséquent, ne visait pas seulement à ajouter une case à cocher IA à notre liste de fonctionnalités.

Il s'agissait de faire un pas important vers une plateforme de gestion de projet plus intelligente et adaptative tout en restant fidèle à notre philosophie fondamentale de simplicité et de centrage sur l'utilisateur.

Dans les sections suivantes, nous plongerons dans les défis techniques auxquels nous avons été confrontés pour donner vie à cette vision, l'architecture que nous avons conçue pour la soutenir et les solutions que nous avons mises en œuvre. Nous explorerons également les possibilités futures qu'ouvre cette fonctionnalité, montrant comment un ajout soigneusement réfléchi peut ouvrir la voie à des changements transformateurs dans la gestion de projet.

---
## Le problème

Comme discuté ci-dessus, l'étiquetage manuel des enregistrements de projet peut être chronophage et incohérent.

Nous avons entrepris de résoudre ce problème en exploitant l'IA pour suggérer automatiquement des étiquettes basées sur le contenu des enregistrements.

Les principaux défis étaient :

1. Choisir un modèle d'IA approprié
2. Traiter efficacement de gros volumes d'enregistrements
3. Assurer la confidentialité et la sécurité des données
4. Intégrer la fonctionnalité de manière transparente dans notre architecture existante

## Sélection du modèle d'IA

Nous avons évalué plusieurs plateformes d'IA, y compris [OpenAI](https://openai.com), des modèles open source sur [HuggingFace](https://huggingface.co/), et [Replicate](https://replicate.com).

Nos critères incluaient :

- Rapport coût-efficacité
- Précision dans la compréhension du contexte
- Capacité à respecter des formats de sortie spécifiques
- Garanties de confidentialité des données

Après des tests approfondis, nous avons choisi [GPT-3.5 Turbo](https://platform.openai.com/docs/models/gpt-3-5-turbo) d'OpenAI. Bien que [GPT-4](https://softgist.com/the-ultimate-guide-to-prompt-engineering) puisse offrir des améliorations marginales en termes de précision, nos tests ont montré que les performances de GPT-3.5 étaient plus que suffisantes pour nos besoins d'étiquetage automatique. L'équilibre entre le rapport coût-efficacité et les fortes capacités de catégorisation a fait de GPT-3.5 le choix idéal pour cette fonctionnalité.

Le coût plus élevé de GPT-4 nous aurait obligés à proposer la fonctionnalité comme un module complémentaire payant, en conflit avec notre objectif d'**intégrer l'IA dans notre produit principal sans coût supplémentaire pour les utilisateurs finaux.**

Au moment de notre implémentation, la tarification pour GPT-3.5 Turbo est :

- 0,0005 $ par 1K tokens d'entrée (ou 0,50 $ par 1M tokens d'entrée)
- 0,0015 $ par 1K tokens de sortie (ou 1,50 $ par 1M tokens de sortie)

Faisons quelques hypothèses sur un enregistrement moyen dans Blue :

- **Titre** : ~10 tokens
- **Description** : ~50 tokens
- **2 commentaires** : ~30 tokens chacun
- **5 champs personnalisés** : ~10 tokens chacun
- **Nom de liste, date d'échéance et autres métadonnées** : ~20 tokens
- **Invite système et étiquettes disponibles** : ~50 tokens

Total de tokens d'entrée par enregistrement : 10 + 50 + (30 * 2) + (10 * 5) + 20 + 50 ≈ 240 tokens

Pour la sortie, supposons une moyenne de 3 étiquettes suggérées par enregistrement, ce qui pourrait totaliser environ 20 tokens de sortie, y compris le formatage JSON.

Pour 1 million d'enregistrements :

- Coût d'entrée : (240 * 1 000 000 / 1 000 000) * 0,50 $ = 120 $
- Coût de sortie : (20 * 1 000 000 / 1 000 000) * 1,50 $ = 30 $

**Coût total pour l'étiquetage automatique de 1 million d'enregistrements : 120 $ + 30 $ = 150 $**

## Performance de GPT3.5 Turbo

La catégorisation est une tâche dans laquelle les grands modèles de langage (LLM) comme GPT-3.5 Turbo excellent, ce qui les rend particulièrement adaptés à notre fonctionnalité d'étiquetage automatique. Les LLM sont entraînés sur de vastes quantités de données textuelles, ce qui leur permet de comprendre le contexte, la sémantique et les relations entre les concepts. Cette large base de connaissances leur permet d'effectuer des tâches de catégorisation avec une grande précision dans un large éventail de domaines.

Pour notre cas d'utilisation spécifique d'étiquetage de gestion de projet, GPT-3.5 Turbo démontre plusieurs forces clés :

- **Compréhension contextuelle :** Peut saisir le contexte global d'un enregistrement de projet, en considérant non seulement des mots individuels mais le sens véhiculé par l'ensemble de la description, des commentaires et d'autres champs.
- **Flexibilité :** Peut s'adapter à divers types de projets et industries sans nécessiter de reprogrammation extensive.
- **Gestion de l'ambiguïté :** Peut peser plusieurs facteurs pour prendre des décisions nuancées.
- **Apprentissage à partir d'exemples :** Peut rapidement comprendre et appliquer de nouveaux schémas de catégorisation sans formation supplémentaire.
- **Classification multi-étiquettes :** Peut suggérer plusieurs étiquettes pertinentes pour un seul enregistrement, ce qui était crucial pour nos exigences.

GPT-3.5 Turbo s'est également distingué par sa fiabilité à respecter notre format de sortie JSON requis, ce qui était *crucial* pour une intégration transparente avec nos systèmes existants. Les modèles open source, bien que prometteurs, ajoutaient souvent des commentaires supplémentaires ou s'écartaient du format attendu, ce qui aurait nécessité un post-traitement supplémentaire. Cette cohérence dans le format de sortie a été un facteur clé dans notre décision, car elle a considérablement simplifié notre implémentation et réduit les points de défaillance potentiels.

Opter pour GPT-3.5 Turbo avec sa sortie JSON cohérente nous a permis d'implémenter une solution plus simple, fiable et maintenable.

Si nous avions choisi un modèle avec un formatage moins fiable, nous aurions fait face à une cascade de complications : le besoin d'une logique d'analyse robuste pour gérer divers formats de sortie, une gestion extensive des erreurs pour les sorties incohérentes, des impacts potentiels sur les performances dus au traitement supplémentaire, une complexité accrue des tests pour couvrir toutes les variations de sortie, et une charge de maintenance à long terme plus importante.

Les erreurs d'analyse pourraient conduire à un étiquetage incorrect, impactant négativement l'expérience utilisateur. En évitant ces écueils, nous avons pu concentrer nos efforts d'ingénierie sur des aspects critiques comme l'optimisation des performances et la conception de l'interface utilisateur, plutôt que de lutter contre des sorties d'IA imprévisibles.

## Architecture du système

Notre fonctionnalité d'étiquetage automatique par IA est construite sur une architecture robuste et évolutive conçue pour gérer efficacement de gros volumes de requêtes tout en offrant une expérience utilisateur transparente. Comme pour tous nos systèmes, nous avons conçu cette fonctionnalité pour supporter un ordre de grandeur de trafic supérieur à celui que nous connaissons actuellement. Cette approche, bien qu'apparemment sur-conçue pour les besoins actuels, est une meilleure pratique qui nous permet de gérer facilement les pics soudains d'utilisation et nous donne une marge de croissance suffisante sans révisions architecturales majeures. Sinon, nous devrions réingénier tous nos systèmes tous les 18 mois — quelque chose que nous avons appris à nos dépens dans le passé !

Décomposons les composants et le flux de notre système :

- **Interaction utilisateur :** Le processus commence lorsqu'un utilisateur appuie sur le bouton "Étiquetage automatique" dans l'interface Blue. Cette action déclenche le flux de travail d'étiquetage automatique.
- **Appel API Blue :** L'action de l'utilisateur est traduite en un appel API vers notre backend Blue. Ce point de terminaison API est conçu pour gérer les demandes d'étiquetage automatique.
- **Gestion de file d'attente :** Au lieu de traiter la demande immédiatement, ce qui pourrait entraîner des problèmes de performance sous charge élevée, nous ajoutons la demande d'étiquetage à une file d'attente. Nous utilisons Redis pour ce mécanisme de file d'attente, ce qui nous permet de gérer efficacement la charge et d'assurer l'évolutivité du système.
- **Service en arrière-plan :** Nous avons implémenté un service en arrière-plan qui surveille continuellement la file d'attente pour de nouvelles demandes. Ce service est responsable du traitement des demandes en file d'attente.
- **Intégration API OpenAI :** Le service en arrière-plan prépare les données nécessaires et effectue des appels API au modèle GPT-3.5 d'OpenAI. C'est là que l'étiquetage alimenté par l'IA se produit réellement. Nous envoyons des données de projet pertinentes et recevons en retour des étiquettes suggérées.
- **Traitement des résultats :** Le service en arrière-plan traite les résultats reçus d'OpenAI. Cela implique d'analyser la réponse de l'IA et de préparer les données pour l'application au projet.
- **Application des étiquettes :** Les résultats traités sont utilisés pour appliquer les nouvelles étiquettes aux éléments pertinents du projet. Cette étape met à jour notre base de données avec les étiquettes suggérées par l'IA.
- **Reflet dans l'interface utilisateur :** Enfin, les nouvelles étiquettes apparaissent dans la vue du projet de l'utilisateur, complétant le processus d'étiquetage automatique du point de vue de l'utilisateur.

Cette architecture offre plusieurs avantages clés qui améliorent à la fois les performances du système et l'expérience utilisateur. En utilisant une file d'attente et un traitement en arrière-plan, nous avons atteint une évolutivité impressionnante, nous permettant de gérer de nombreuses demandes simultanément sans surcharger notre système ou atteindre les limites de débit de l'API OpenAI. La mise en œuvre de cette architecture a nécessité une considération minutieuse de divers facteurs pour assurer des performances et une fiabilité optimales. Pour la gestion de file d'attente, nous avons choisi Redis, tirant parti de sa vitesse et de sa fiabilité dans la gestion des files d'attente distribuées.

Cette approche contribue également à la réactivité globale de la fonctionnalité. Les utilisateurs reçoivent un retour immédiat indiquant que leur demande est en cours de traitement, même si l'étiquetage réel prend du temps, créant une sensation d'interaction en temps réel. La tolérance aux pannes de l'architecture est un autre avantage crucial. Si une partie du processus rencontre des problèmes, comme des perturbations temporaires de l'API OpenAI, nous pouvons réessayer gracieusement ou gérer l'échec sans affecter l'ensemble du système.

Cette robustesse, combinée à l'apparition en temps réel des étiquettes, améliore l'expérience utilisateur, donnant l'impression d'une "magie" IA à l'œuvre.

## Données et invites

Une étape cruciale dans notre processus d'étiquetage automatique est la préparation des données à envoyer au modèle GPT-3.5. Cette étape a nécessité une considération minutieuse pour équilibrer la fourniture de suffisamment de contexte pour un étiquetage précis tout en maintenant l'efficacité et en protégeant la confidentialité des utilisateurs. Voici un aperçu détaillé de notre processus de préparation des données.

Pour chaque enregistrement, nous compilons les informations suivantes :

- **Nom de la liste :** Fournit un contexte sur la catégorie ou la phase plus large du projet.
- **Titre de l'enregistrement :** Contient souvent des informations clés sur l'objectif ou le contenu de l'enregistrement.
- **Champs personnalisés :** Nous incluons des [champs personnalisés](/platform/features/custom-fields) basés sur du texte et des nombres, qui contiennent souvent des informations cruciales spécifiques au projet.
- **Description :** Contient généralement les informations les plus détaillées sur l'enregistrement.
- **Commentaires :** Peuvent fournir un contexte supplémentaire ou des mises à jour qui pourraient être pertinentes pour l'étiquetage.
- **Date d'échéance :** Les informations temporelles qui pourraient influencer la sélection des étiquettes.

Il est intéressant de noter que nous n'envoyons pas les données d'étiquettes existantes à GPT-3.5, et nous faisons cela pour éviter de biaiser le modèle.

Le cœur de notre fonctionnalité d'étiquetage automatique réside dans la façon dont nous interagissons avec le modèle GPT-3.5 et traitons ses réponses. Cette section de notre pipeline a nécessité une conception minutieuse pour assurer un étiquetage précis, cohérent et efficace.

Nous utilisons une invite système soigneusement conçue pour instruire l'IA sur sa tâche. Voici une décomposition de notre invite et la justification derrière chaque composant :

```
You will be provided with record data, and your task is to choose the tags that are relevant to the record.
You can respond with an empty array if you are unsure.
Available tags: ${tags}.
Today: ${currentDate}.
Please respond in JSON using the following format:
{ "tags": ["tag-1", "tag-2"] }
```

- **Définition de la tâche :** Nous énonçons clairement la tâche de l'IA pour garantir des réponses ciblées.
- **Gestion de l'incertitude :** Nous permettons explicitement des réponses vides, empêchant l'étiquetage forcé lorsque l'IA n'est pas sûre.
- **Étiquettes disponibles :** Nous fournissons une liste d'étiquettes valides (${tags}) pour contraindre les choix de l'IA aux étiquettes de projet existantes.
- **Date actuelle :** L'inclusion de ${currentDate} aide l'IA à comprendre le contexte temporel, ce qui peut être crucial pour certains types de projets.
- **Format de réponse :** Nous spécifions un format JSON pour une analyse facile et une vérification des erreurs.

Cette invite est le résultat de tests et d'itérations approfondis. Nous avons constaté qu'être explicite sur la tâche, les options disponibles et le format de sortie souhaité améliorait considérablement la précision et la cohérence des réponses de l'IA — la simplicité est la clé !

La liste des étiquettes disponibles est générée côté serveur et validée avant l'inclusion dans l'invite. Nous implémentons des limites strictes de caractères sur les noms d'étiquettes pour éviter des invites surdimensionnées.

Comme mentionné ci-dessus, nous n'avons eu aucun problème avec GPT-3.5 Turbo pour récupérer la réponse JSON pure dans le format correct 100 % du temps.

Donc en résumé,

- Nous combinons l'invite système avec les données d'enregistrement préparées.
- Cette invite combinée est envoyée au modèle GPT-3.5 via l'API d'OpenAI.
- Nous utilisons un paramètre de température de 0,3 pour équilibrer la créativité et la cohérence dans les réponses de l'IA.
- Notre appel API inclut un paramètre max_tokens pour limiter la taille de la réponse et contrôler les coûts.

Une fois que nous recevons la réponse de l'IA, nous passons par plusieurs étapes pour traiter et appliquer les étiquettes suggérées :

* **Analyse JSON** : Nous tentons d'analyser la réponse en tant que JSON. Si l'analyse échoue, nous enregistrons l'erreur et sautons l'étiquetage pour cet enregistrement.
* **Validation du schéma** : Nous validons le JSON analysé par rapport à notre schéma attendu (un objet avec un tableau "tags"). Cela détecte toute déviation structurelle dans la réponse de l'IA.
* **Validation des étiquettes** : Nous recoupons les étiquettes suggérées avec notre liste d'étiquettes de projet valides. Cette étape filtre toutes les étiquettes qui n'existent pas dans le projet, ce qui pourrait se produire si l'IA a mal compris ou si les étiquettes du projet ont changé entre la création de l'invite et le traitement de la réponse.
* **Dédoublonnage** : Nous supprimons toutes les étiquettes en double de la suggestion de l'IA pour éviter un étiquetage redondant.
* **Application** : Les étiquettes validées et dédoublonnées sont ensuite appliquées à l'enregistrement dans notre base de données.
* **Journalisation et analyse** : Nous enregistrons les étiquettes finales appliquées. Ces données sont précieuses pour surveiller les performances du système et l'améliorer au fil du temps.

## Défis

L'implémentation de l'étiquetage automatique alimenté par l'IA dans Blue a présenté plusieurs défis uniques, chacun nécessitant des solutions innovantes pour garantir une fonctionnalité robuste, efficace et conviviale.

### Annuler l'opération en masse

La fonctionnalité d'étiquetage IA peut être effectuée à la fois sur des enregistrements individuels ainsi qu'en masse. Le problème avec l'opération en masse est que si l'utilisateur n'aime pas le résultat, il devrait passer manuellement en revue des milliers d'enregistrements et annuler le travail de l'IA. Clairement, c'est inacceptable.

Pour résoudre ce problème, nous avons implémenté un système innovant de session d'étiquetage. Chaque opération d'étiquetage en masse se voit attribuer un ID de session unique, qui est associé à toutes les étiquettes appliquées pendant cette session. Cela nous permet de gérer efficacement les opérations d'annulation en supprimant simplement toutes les étiquettes associées à un ID de session particulier. Nous supprimons également les pistes d'audit liées, garantissant que les opérations annulées ne laissent aucune trace dans le système. Cette approche donne aux utilisateurs la confiance d'expérimenter avec l'étiquetage IA, sachant qu'ils peuvent facilement annuler les modifications si nécessaire.

### Confidentialité des données

La confidentialité des données était un autre défi critique auquel nous avons été confrontés.

Nos utilisateurs nous font confiance avec leurs données de projet, et il était primordial de s'assurer que ces informations n'étaient pas conservées ou utilisées pour l'entraînement du modèle par OpenAI. Nous avons abordé cela sur plusieurs fronts.

Premièrement, nous avons conclu un accord avec OpenAI qui interdit explicitement l'utilisation de nos données pour l'entraînement du modèle. De plus, OpenAI supprime les données après le traitement, offrant une couche supplémentaire de protection de la confidentialité.

De notre côté, nous avons pris la précaution d'exclure les informations sensibles, telles que les détails des personnes assignées, des données envoyées à l'IA afin de garantir que les noms spécifiques des individus ne soient pas envoyés à des tiers avec d'autres données.

Cette approche complète nous permet d'exploiter les capacités de l'IA tout en maintenant les normes les plus élevées de confidentialité et de sécurité des données.

### Limites de débit et capture des erreurs

L'une de nos principales préoccupations était l'évolutivité et la limitation du débit. Les appels API directs à OpenAI pour chaque enregistrement auraient été inefficaces et pourraient rapidement atteindre les limites de débit, en particulier pour les grands projets ou pendant les périodes de pointe. Pour résoudre ce problème, nous avons développé une architecture de service en arrière-plan qui nous permet de regrouper les demandes et d'implémenter notre propre système de file d'attente. Cette approche nous aide à gérer la fréquence des appels API et permet un traitement plus efficace de gros volumes d'enregistrements, garantissant des performances fluides même sous forte charge.

La nature des interactions IA signifiait que nous devions également nous préparer à des erreurs occasionnelles ou des sorties inattendues. Il y avait des cas où l'IA pouvait produire du JSON invalide ou des sorties qui ne correspondaient pas à notre format attendu. Pour gérer cela, nous avons implémenté une gestion robuste des erreurs et une logique d'analyse tout au long de notre système. Si la réponse de l'IA n'est pas un JSON valide ou ne contient pas la clé "tags" attendue, notre système est conçu pour la traiter comme si aucune étiquette n'était suggérée, plutôt que de tenter de traiter des données potentiellement corrompues. Cela garantit que même face à l'imprévisibilité de l'IA, notre système reste stable et fiable.

## Développements futurs

Nous croyons que les fonctionnalités, et le produit Blue dans son ensemble, ne sont jamais "terminés" — il y a toujours place à l'amélioration.

Il y avait certaines fonctionnalités que nous avons envisagées dans la construction initiale qui n'ont pas passé la phase de définition du périmètre, mais qu'il est intéressant de noter car nous en implémenterons probablement certaines versions à l'avenir.

La première est l'ajout de descriptions d'étiquettes. Cela permettrait aux utilisateurs finaux de donner non seulement un nom et une couleur aux étiquettes, mais aussi une description optionnelle. Cela serait également transmis à l'IA pour aider à fournir un contexte supplémentaire et potentiellement améliorer la précision.

Bien qu'un contexte supplémentaire puisse être précieux, nous sommes conscients de la complexité potentielle qu'il pourrait introduire. Il y a un équilibre délicat à trouver entre fournir des informations utiles et submerger les utilisateurs avec trop de détails. Au fur et à mesure que nous développons cette fonctionnalité, nous nous concentrerons sur la recherche de ce point idéal où le contexte ajouté améliore plutôt que complique l'expérience utilisateur.

L'amélioration la plus excitante à notre horizon est peut-être l'intégration de l'étiquetage automatique par IA avec notre [système d'automatisation de gestion de projet](/platform/features/automations).

Cela signifie que la fonctionnalité d'étiquetage IA pourrait être soit un déclencheur, soit une action d'une automatisation. Cela pourrait être énorme car cela pourrait transformer cette fonctionnalité de catégorisation IA plutôt simple en un système de routage basé sur l'IA pour le travail.

Imaginez une automatisation qui indique :

Quand l'IA étiquette un enregistrement comme "Critique" -> Assigné à "Manager" et Envoyer un email personnalisé

Cela signifie que lorsque vous étiquetez un enregistrement par IA, si l'IA décide qu'il s'agit d'un problème critique, alors elle peut automatiquement assigner le chef de projet et lui envoyer un email personnalisé. Cela étend les [avantages de notre système d'automatisation de gestion de projet](/platform/features/automations) d'un système purement basé sur des règles à un véritable système IA flexible.

En explorant continuellement les frontières de l'IA dans la gestion de projet, nous visons à fournir à nos utilisateurs des outils qui non seulement répondent à leurs besoins actuels mais anticipent et façonnent l'avenir du travail. Comme toujours, nous développerons ces fonctionnalités en étroite collaboration avec notre communauté d'utilisateurs, en veillant à ce que chaque amélioration ajoute une valeur réelle et pratique au processus de gestion de projet.

## Conclusion

Voilà !

C'était une fonctionnalité amusante à implémenter, et l'un de nos premiers pas dans l'IA, aux côtés du [Résumé de contenu par IA](/insights/ai-content-summarization) que nous avons précédemment lancé. Nous savons que l'IA jouera un rôle de plus en plus important dans la gestion de projet à l'avenir, et nous avons hâte de déployer des fonctionnalités plus innovantes exploitant des LLM (Large Language Models) avancés.

Il y avait beaucoup de choses à penser lors de l'implémentation de ceci, et nous sommes particulièrement enthousiastes quant à la façon dont nous pouvons tirer parti de cette fonctionnalité à l'avenir avec le [moteur d'automatisation de gestion de projet](/insights/benefits-project-management-automation) existant de Blue.

Nous espérons également que cela a été une lecture intéressante, et que cela vous donne un aperçu de la façon dont nous pensons à l'ingénierie des fonctionnalités que vous utilisez tous les jours.

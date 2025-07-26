---
title: Créer des listes de contrôle réutilisables à l'aide d'automatisations
description: Apprenez à créer des automatisations de gestion de projet pour des listes de contrôle réutilisables.
category: "Best Practices"
date: 2024-07-08
---


Dans de nombreux projets et processus, vous pourriez avoir besoin d'utiliser la même liste de contrôle pour plusieurs enregistrements ou tâches.

Cependant, il n'est pas très efficace de retaper manuellement la liste de contrôle chaque fois que vous souhaitez l'ajouter à un enregistrement. C'est ici que vous pouvez tirer parti des [automatisations de gestion de projet puissantes](/platform/features/automations) pour le faire automatiquement pour vous !

Pour rappel, les automatisations dans Blue nécessitent deux éléments clés :

1. Un Déclencheur — Ce qui doit se passer pour démarrer l'automatisation. Cela peut être lorsque qu'un enregistrement reçoit une étiquette spécifique, lorsqu'il passe à une liste spécifique.
2. Une ou plusieurs Actions — Dans ce cas, il s'agirait de la création automatique d'une ou plusieurs listes de contrôle.

Commençons par l'action d'abord, puis discutons des déclencheurs possibles que vous pouvez utiliser.

## Action d'Automatisation de Liste de Contrôle

Vous pouvez créer une nouvelle automatisation et configurer une ou plusieurs listes de contrôle à créer, comme dans l'exemple ci-dessous :

![](/insights/checklist-automation.png)

Ce seraient les listes de contrôle que vous souhaitez créer chaque fois que vous effectuez l'action.

## Déclencheurs d'Automatisation de Liste de Contrôle

Il existe plusieurs façons de déclencher la création de vos listes de contrôle réutilisables. Voici quelques options populaires :

- **Ajout d'une Étiquette Spécifique :** Vous pouvez configurer l'automatisation pour qu'elle se déclenche lorsqu'une étiquette particulière est ajoutée à un enregistrement. Par exemple, lorsque l'étiquette "Nouveau Projet" est ajoutée, cela pourrait automatiquement créer votre liste de contrôle de démarrage de projet.
- **Attribution d'Enregistrement :** La création de la liste de contrôle peut être déclenchée lorsqu'un enregistrement est attribué à une personne spécifique ou à quiconque. Cela est utile pour les listes de contrôle d'intégration ou les procédures spécifiques aux tâches.
- **Déplacement vers une Liste Spécifique :** Lorsqu'un enregistrement est déplacé vers une liste particulière dans votre tableau de projet, cela peut déclencher la création d'une liste de contrôle pertinente. Par exemple, déplacer un élément vers une liste "Assurance Qualité" pourrait déclencher une liste de contrôle QA.
- **Champ de Case à Cocher Personnalisé :** Créez un champ de case à cocher personnalisé et configurez l'automatisation pour qu'elle se déclenche lorsque cette case est cochée. Cela vous donne un contrôle manuel sur le moment d'ajouter la liste de contrôle.
- **Champs Personnalisés à Sélection Unique ou Multiple :** Vous pouvez créer un champ personnalisé à sélection unique ou multiple avec diverses options. Chaque option peut être liée à un modèle de liste de contrôle spécifique via des automatisations séparées. Cela permet un contrôle plus granulaire et la possibilité d'avoir plusieurs modèles de listes de contrôle prêts pour différents scénarios.

Pour améliorer le contrôle sur qui peut déclencher ces automatisations, vous pouvez masquer ces champs personnalisés à certains utilisateurs en utilisant des rôles d'utilisateur personnalisés. Cela garantit que seuls les administrateurs de projet ou d'autres personnes autorisées peuvent déclencher ces options.

N'oubliez pas, la clé d'une utilisation efficace des listes de contrôle réutilisables avec des automatisations est de concevoir vos déclencheurs de manière réfléchie. Considérez le flux de travail de votre équipe, les types de projets que vous gérez et qui devrait avoir la capacité d'initier différents processus. Avec des automatisations bien planifiées, vous pouvez considérablement rationaliser votre gestion de projet et garantir la cohérence de vos opérations.

## Ressources Utiles

- [Documentation sur l'Automatisation de la Gestion de Projet](https://documentation.blue.cc/automations)
- [Documentation sur les Rôles d'Utilisateur Personnalisés](https://documentation.blue.cc/user-management/roles/custom-user-roles)
- [Documentation sur les Champs Personnalisés](https://documentation.blue.cc/custom-fields)
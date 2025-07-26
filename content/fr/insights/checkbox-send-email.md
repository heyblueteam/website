---
title: Automatisation de la gestion de projet — emails aux parties prenantes.
description: Souvent, vous souhaitez contrôler vos automatisations de gestion de projet.
category: "Product Updates"
date: 2024-07-08
---


Nous avons déjà abordé comment [créer des automatisations par email auparavant.](/insights/email-automations)

Cependant, il y a souvent des parties prenantes dans les projets qui n'ont besoin d'être alertées que lorsqu'il y a quelque chose de *vraiment* important.

Ne serait-il pas agréable d'avoir une automatisation de gestion de projet où vous, en tant que chef de projet, pourriez contrôler *exactement* quand notifier une partie prenante clé d'une simple pression sur un bouton ?

Eh bien, il s'avère qu'avec Blue, vous pouvez faire précisément cela !

Aujourd'hui, nous allons apprendre à créer une automatisation de gestion de projet vraiment utile :

Une case à cocher qui notifie automatiquement une ou plusieurs parties prenantes clés, leur fournissant tout le contexte essentiel concernant ce que vous les notifiez. En bonus, nous apprendrons également comment restreindre cette capacité afin que seuls certains membres de votre projet puissent déclencher cette notification par email.

Cela ressemblera à quelque chose comme ça une fois que vous aurez terminé :

![](/insights/checkbox-email-automation.png)

Et en cochant simplement cette case, vous pourrez déclencher une automatisation de gestion de projet pour envoyer un email de notification personnalisé aux parties prenantes.

Allons-y étape par étape.

## 1. Créez votre champ personnalisé de case à cocher

C'est très facile, vous pouvez consulter notre [documentation détaillée](https://documentation.blue.cc/custom-fields/introduction#creating-custom-fields) sur la création de champs personnalisés.

Assurez-vous de nommer ce champ quelque chose d'évident que vous vous souviendrez, comme "notifier la direction" ou "notifier les parties prenantes".

## 2. Créez votre déclencheur d'automatisation de gestion de projet.

Dans la vue des enregistrements de votre projet, cliquez sur le petit robot en haut à droite pour ouvrir les paramètres d'automatisation :

<video autoplay loop muted playsinline>
  <source src="/videos/notify-stakeholders-automation-setup.mp4" type="video/mp4">
</video>

## 3. Créez votre action d'automatisation de gestion de projet.

Dans ce cas, notre action sera d'envoyer une notification par email personnalisée à une ou plusieurs adresses email. Il est bon de noter ici que ces personnes ne doivent **pas** être dans Blue pour recevoir ces emails, vous pouvez envoyer des emails à *n'importe* quelle adresse email.

Vous pouvez en apprendre davantage dans notre [guide de documentation détaillé sur la configuration des automatisations par email](https://documentation.blue.cc/automations/actions/email-automations)

Votre résultat final devrait ressembler à quelque chose comme ça :

![](/insights/email-automation-example.png)

## 4. Bonus : Restreindre l'accès à la case à cocher.

Vous pouvez utiliser [des rôles d'utilisateur personnalisés dans Blue](/platform/features/user-permissions) pour restreindre l'accès aux champs personnalisés de case à cocher, garantissant que seuls les membres autorisés de l'équipe peuvent déclencher des notifications par email.

Blue permet aux administrateurs de projet de définir des rôles et d'attribuer des autorisations à des groupes d'utilisateurs. Ce système est crucial pour maintenir le contrôle sur qui peut interagir avec des éléments spécifiques de votre projet, y compris des champs personnalisés comme la case à cocher de notification.

1. Accédez à la section Gestion des utilisateurs dans Blue et sélectionnez "Rôles d'utilisateur personnalisés".
2. Créez un nouveau rôle en fournissant un nom descriptif et une description optionnelle.
3. Dans les autorisations du rôle, localisez la section Accès aux champs personnalisés.
4. Spécifiez si le rôle peut voir ou modifier le champ personnalisé de case à cocher. Par exemple, restreindre l'accès à la modification aux rôles comme "Administrateur de projet" tout en permettant à un rôle personnalisé nouvellement créé de gérer ce champ.
5. Attribuez le rôle nouvellement créé aux utilisateurs ou groupes d'utilisateurs appropriés. Cela garantit que seules les personnes désignées ont la capacité d'interagir avec la case à cocher de notification.

[Lisez-en plus sur notre site de documentation officiel.](https://documentation.blue.cc/user-management/roles/custom-user-roles)

En mettant en œuvre ces rôles personnalisés, vous améliorez la sécurité et l'intégrité de vos processus de gestion de projet. Seuls les membres d'équipe autorisés peuvent déclencher des notifications par email critiques, garantissant que les parties prenantes reçoivent des mises à jour importantes sans alertes inutiles.

## Conclusion

En mettant en œuvre l'automatisation de gestion de projet décrite ci-dessus, vous gagnez un contrôle précis sur quand et comment notifier les parties prenantes clés. Cette approche garantit que les mises à jour importantes sont communiquées efficacement, sans submerger vos parties prenantes avec des informations inutiles. En utilisant les fonctionnalités de champ personnalisé et d'automatisation de Blue, vous pouvez rationaliser votre processus de gestion de projet, améliorer la communication et maintenir un haut niveau d'efficacité.

Avec juste une simple case à cocher, vous pouvez déclencher des notifications par email personnalisées adaptées aux besoins de votre projet, garantissant que les bonnes personnes sont informées au bon moment. De plus, la capacité de restreindre cette fonctionnalité à des membres spécifiques de l'équipe ajoute une couche supplémentaire de contrôle et de sécurité.

Commencez à tirer parti de cette fonctionnalité puissante dans Blue dès aujourd'hui pour tenir vos parties prenantes informées et faire en sorte que vos projets se déroulent sans accroc. Pour des étapes plus détaillées et des options de personnalisation supplémentaires, consultez les liens de documentation fournis. Bonne automatisation !
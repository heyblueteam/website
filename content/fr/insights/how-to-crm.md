---
title: Comment configurer Blue en tant que CRM
description: Apprenez à configurer Blue pour suivre vos clients et vos affaires de manière simple.
category: "Best Practices"
date: 2024-08-11
---


## Introduction

L'un des principaux avantages de l'utilisation de Blue est de ne pas l'utiliser pour un cas d'utilisation *spécifique*, mais de l'utiliser *dans* plusieurs cas d'utilisation. Cela signifie que vous n'avez pas à payer pour plusieurs outils, et vous avez également un endroit où vous pouvez facilement passer d'un projet à l'autre et gérer vos divers processus tels que le recrutement, les ventes, le marketing, et plus encore.

En aidant des milliers de clients à se configurer dans Blue au fil des ans, nous avons remarqué que la partie difficile n'est *pas* de configurer Blue lui-même, mais de réfléchir aux processus et de tirer le meilleur parti de notre plateforme.

Les éléments clés sont de penser au flux de travail étape par étape pour chacun de vos processus commerciaux que vous souhaitez suivre, ainsi qu'aux spécificités des données que vous souhaitez capturer, et comment cela se traduit par les champs personnalisés que vous configurez.

Aujourd'hui, nous allons vous guider à travers la création [d'un système CRM de vente facile à utiliser, mais puissant](/solutions/use-case/sales-crm) avec une base de données clients liée à un pipeline d'opportunités. Toutes ces données seront intégrées dans un tableau de bord où vous pourrez voir des données en temps réel sur vos ventes totales, vos ventes prévisionnelles, et plus encore.

## Base de données clients

La première chose à faire est de configurer un nouveau projet pour stocker vos données clients. Ces données seront ensuite croisées dans un autre projet où vous suivez des opportunités de vente spécifiques.

La raison pour laquelle nous séparons vos informations clients des opportunités est qu'elles ne se correspondent pas un à un.

Un client peut avoir plusieurs opportunités ou projets.

Par exemple, si vous êtes une agence de marketing et de design, vous pouvez d'abord vous engager avec un client pour son image de marque, puis réaliser un projet séparé pour son site web, et ensuite un autre pour sa gestion des réseaux sociaux.

Tous ceux-ci seraient des opportunités de vente distinctes nécessitant leur propre suivi et propositions, mais elles sont toutes liées à ce client unique.

L'avantage de séparer votre base de données clients dans un projet distinct est que si vous mettez à jour des détails dans votre base de données clients, toutes vos opportunités auront automatiquement les nouvelles données, ce qui signifie que vous avez maintenant une source unique de vérité dans votre entreprise ! Vous n'avez pas à revenir et à tout modifier manuellement !

Ainsi, la première chose à décider est si vous allez être centré sur l'entreprise ou centré sur la personne.

Cette décision dépend vraiment de ce que vous vendez et à qui vous vendez. Si vous vendez principalement à des entreprises, vous voudrez probablement que le nom de l'enregistrement soit le nom de l'entreprise. Cependant, si vous vendez principalement à des particuliers (c'est-à-dire que vous êtes un coach de santé personnel ou un consultant en image de marque personnelle), vous adopterez probablement une approche centrée sur la personne.

Ainsi, le champ du nom de l'enregistrement sera soit le nom de l'entreprise, soit le nom de la personne, selon votre choix. La raison en est que cela signifie que vous pouvez facilement identifier un client d'un coup d'œil, juste en regardant votre tableau ou votre base de données.

Ensuite, vous devez considérer quelles informations vous souhaitez capturer dans votre base de données clients. Ce sont celles qui deviendront vos champs personnalisés.

Les suspects habituels ici sont :

- Email
- Numéro de téléphone
- Site web
- Adresse
- Source (c'est-à-dire d'où vient ce client pour la première fois ?)
- Catégorie

Dans Blue, vous pouvez également supprimer tous les champs par défaut dont vous n'avez pas besoin. Pour cette base de données clients, nous vous recommandons généralement de supprimer la date d'échéance, l'assigné, les dépendances et les listes de contrôle. Vous voudrez peut-être garder notre champ de description par défaut disponible au cas où vous auriez des notes générales sur ce client qui ne sont pas spécifiques à une opportunité de vente.

Nous recommandons de garder le champ "Référencé par", car cela sera utile plus tard. Une fois que nous aurons configuré notre base de données d'opportunités, nous pourrons voir chaque enregistrement de vente qui est lié à ce client particulier ici.

En termes de listes, nous voyons généralement nos clients garder cela simple et avoir une liste appelée "Clients" et s'en tenir à cela. Il est préférable d'utiliser des étiquettes ou des champs personnalisés pour la catégorisation.

Ce qui est génial ici, c'est qu'une fois que vous avez cette configuration, vous pouvez facilement importer vos données d'autres systèmes ou feuilles Excel dans Blue via notre fonction d'importation CSV, et vous pouvez également créer un formulaire pour que de nouveaux clients potentiels soumettent leurs détails afin que vous puissiez **automatiquement** les capturer dans votre base de données.

## Base de données d'opportunités

Maintenant que nous avons notre base de données clients, nous devons créer un autre projet pour capturer nos véritables opportunités de vente. Vous pouvez appeler ce projet "CRM de vente" ou "Opportunités".

### Listes en tant qu'étapes de processus

Pour configurer votre processus de vente, vous devez réfléchir aux étapes habituelles par lesquelles une opportunité passe depuis le moment où vous recevez une demande d'un client jusqu'à la signature d'un contrat.

Chaque liste dans votre projet sera une étape de votre processus.

Quelle que soit votre processus spécifique, il y aura quelques listes communes que TOUS les CRM de vente devraient avoir :

- Non qualifié — Toutes les demandes entrantes, où vous n'avez pas encore qualifié un client.
- Gagné — Toutes les opportunités que vous avez remportées et transformées en ventes !
- Perdu — Toutes les opportunités pour lesquelles vous avez fait une offre à un client, et qu'il n'a pas acceptée.
- N/A — C'est ici que vous placez toutes les opportunités que vous n'avez pas gagnées, mais qui n'ont pas non plus été "perdues". Cela pourrait être celles que vous avez refusées, celles où le client, pour une raison quelconque, vous a ignoré, et ainsi de suite.

En réfléchissant à votre processus commercial CRM, vous devriez considérer le niveau de granularité que vous souhaitez. Nous ne recommandons pas d'avoir 20 ou 30 colonnes, cela devient généralement confus et vous empêche de voir la vue d'ensemble.

Cependant, il est également important de ne pas rendre chaque processus trop large, sinon les affaires seront "bloquées" à une étape spécifique pendant des semaines ou des mois, même lorsqu'elles avancent en réalité. Voici une approche typique recommandée :

- **Non qualifié** : Toutes les demandes entrantes, où vous n'avez pas encore qualifié un client.
- **Qualification** : C'est ici que vous prenez l'opportunité et commencez le processus de compréhension si c'est un bon ajustement pour votre entreprise.
- **Rédaction de la proposition** : C'est ici que vous commencez à transformer l'opportunité en une proposition pour votre entreprise. C'est un document que vous enverrez au client.
- **Proposition envoyée** : C'est ici que vous avez envoyé la proposition au client et attendez une réponse.
- **Négociations** : C'est ici que vous êtes en train de finaliser l'accord.
- **Contrat en attente de signature** : C'est ici que vous attendez simplement que le client signe le contrat.
- **Gagné** : C'est ici que vous avez remporté l'affaire et que vous travaillez maintenant sur le projet.
- **Perdu** : C'est ici que vous avez fait une offre au client, mais qu'il n'a pas accepté les conditions.
- **N/A** : C'est ici que vous placez toutes les opportunités que vous n'avez pas gagnées, mais qui n'ont pas non plus été "perdues". Cela pourrait être celles que vous avez refusées, celles où le client, pour une raison quelconque, vous a ignoré, et ainsi de suite.

### Étiquettes en tant que catégories de services
Parlons maintenant des étiquettes.

Nous recommandons d'utiliser des étiquettes pour les différents types de services que vous proposez. Donc, en revenant à notre exemple d'agence de marketing et de design, vous pourriez avoir des étiquettes pour "branding", "site web", "SEO", "gestion Facebook", et ainsi de suite.

Les avantages ici sont que vous pouvez facilement filtrer par service en un clic, ce qui peut vous donner un aperçu rapide des services les plus populaires, et cela peut également informer les futures embauches, car généralement, différents services nécessitent différents membres de l'équipe.

### Champs personnalisés du CRM de vente

Ensuite, nous devons considérer quels champs personnalisés nous voulons avoir.

Les champs typiques que nous voyons utilisés sont :

- **Montant** : C'est un champ de devise pour le montant du projet.
- **Coût** : Votre coût prévu pour réaliser la vente, également un champ de devise.
- **Profit** : Un champ de formule pour calculer le profit basé sur les champs de montant et de coût.
- **URL de la proposition** : Cela peut inclure un lien vers un document Google en ligne ou un document Word de votre proposition, afin que vous puissiez facilement cliquer et le consulter.
- **Fichiers reçus** : Cela peut être un champ de fichier personnalisé où vous pouvez déposer tous les fichiers reçus du client tels que des matériaux de recherche, des NDA, et ainsi de suite.
- **Contrats** : Un autre champ de fichier personnalisé où vous pouvez ajouter des contrats signés pour les conserver en sécurité.
- **Niveau de confiance** : Un champ personnalisé avec 5 étoiles, indiquant à quel point vous êtes confiant de gagner cette opportunité particulière. Cela peut être utilisé plus tard dans le tableau de bord pour les prévisions !
- **Date de clôture prévue** : Un champ de date pour estimer quand l'affaire est susceptible de se conclure.
- **Client** : Un champ de référence liant à la personne de contact principale dans la base de données clients.
- **Nom du client** : Un champ de recherche qui extrait le nom du client de l'enregistrement lié dans la base de données clients.
- **Email du client** : Un champ de recherche qui extrait l'email du client de l'enregistrement lié dans la base de données clients.
- **Source de l'affaire** : Un champ déroulant pour suivre d'où provient l'opportunité (par exemple, recommandation, site web, appel à froid, salon professionnel).
- **Raison de la perte** : Un champ déroulant (pour les affaires perdues) pour catégoriser pourquoi l'opportunité a été perdue.
- **Taille du client** : Un champ déroulant pour catégoriser les clients par taille (par exemple, petite, moyenne, grande entreprise).

Encore une fois, c'est vraiment **à vous** de décider précisément quels champs vous souhaitez avoir. Un mot d'avertissement : il est facile, lors de la configuration, d'ajouter beaucoup de champs à votre CRM de vente pour les données que vous souhaitez capturer. Cependant, vous devez être réaliste en termes de discipline et d'engagement en temps. Il n'y a aucun intérêt à avoir 30 champs dans votre CRM de vente si 90 % des enregistrements n'auront aucune donnée.

La grande chose à propos des champs personnalisés est qu'ils s'intègrent bien dans [les autorisations personnalisées](/platform/features/user-permissions). Cela signifie que vous pouvez décider exactement quels champs les membres de votre équipe peuvent voir ou modifier. Par exemple, vous voudrez peut-être masquer les informations de coût et de profit aux employés juniors.

### Automatisations

[Les automatisations du CRM de vente](/platform/features/automations) sont une fonctionnalité puissante dans Blue qui peut rationaliser votre processus de vente, garantir la cohérence et économiser du temps sur des tâches répétitives. En configurant des automatisations intelligentes, vous pouvez améliorer l'efficacité de votre CRM de vente et permettre à votre équipe de se concentrer sur ce qui compte le plus : conclure des affaires. Voici quelques automatisations clés à envisager pour votre CRM de vente :

- **Attribution de nouveaux leads** : Attribuez automatiquement de nouveaux leads aux représentants des ventes en fonction de critères prédéfinis tels que la localisation, la taille de l'affaire ou l'industrie. Cela garantit un suivi rapide et une répartition équilibrée de la charge de travail.
- **Rappels de suivi** : Configurez des rappels automatiques pour que les représentants des ventes suivent les prospects après une certaine période d'inactivité. Cela aide à éviter que des leads ne tombent dans l'oubli.
- **Notifications de progression d'étape** : Informez les membres de l'équipe concernés lorsqu'une affaire passe à une nouvelle étape dans le pipeline. Cela permet à chacun d'être informé des progrès et de permettre des interventions en temps opportun si nécessaire.
- **Alertes d'ancienneté des affaires** : Créez des alertes pour les affaires qui sont restées à une étape particulière plus longtemps que prévu. Cela aide à identifier les affaires bloquées qui peuvent nécessiter une attention supplémentaire.

## Lien entre clients et affaires

L'une des fonctionnalités les plus puissantes de Blue pour créer un système CRM efficace est la capacité de lier votre base de données clients avec vos opportunités de vente. Cette connexion vous permet de maintenir une source unique de vérité pour les informations clients tout en suivant plusieurs affaires associées à chaque client. Explorons comment configurer cela en utilisant des champs personnalisés de référence et de recherche.

### Configuration du champ de référence

1. Dans votre projet Opportunités (ou CRM de vente), créez un nouveau champ personnalisé.
2. Choisissez le type de champ "Référence".
3. Sélectionnez votre projet de base de données clients comme source pour la référence.
4. Configurez le champ pour permettre une sélection unique (car chaque opportunité est généralement associée à un seul client).
5. Nommez ce champ quelque chose comme "Client" ou "Entreprise associée".

Maintenant, lorsque vous créez ou modifiez une opportunité, vous pourrez sélectionner le client associé à partir d'un menu déroulant peuplé d'enregistrements de votre base de données clients.

### Amélioration avec des champs de recherche

Une fois que vous avez établi la connexion de référence, vous pouvez utiliser des champs de recherche pour apporter des informations clients pertinentes directement dans votre vue d'opportunités. Voici comment :

1. Dans votre projet Opportunités, créez un nouveau champ personnalisé.
2. Choisissez le type de champ "Recherche".
3. Sélectionnez le champ de référence que vous venez de créer ("Client" ou "Entreprise associée") comme source.
4. Choisissez quelles informations clients vous souhaitez afficher. Vous pourriez envisager des champs comme : Email, Numéro de téléphone, Catégorie de client, Responsable de compte.

Répétez ce processus pour chaque information client que vous souhaitez afficher dans votre vue d'opportunités.

Les avantages de cela sont :

- **Source unique de vérité** : Mettez à jour les informations clients une fois dans la base de données clients, et cela se reflète automatiquement dans toutes les opportunités liées.
- **Efficacité** : Accédez rapidement aux détails clients pertinents tout en travaillant sur des opportunités sans avoir à passer d'un projet à l'autre.
- **Intégrité des données** : Réduisez les erreurs de saisie manuelle en tirant automatiquement les informations clients.
- **Vue holistique** : Voir facilement toutes les opportunités associées à un client en utilisant le champ "Référencé par" dans votre base de données clients.

### Astuce avancée : Rechercher une recherche

Blue propose une fonctionnalité avancée appelée "Rechercher une recherche" qui peut être incroyablement utile pour des configurations CRM complexes. Cette fonctionnalité vous permet de créer des connexions entre plusieurs projets, vous permettant d'accéder à des informations à la fois de votre base de données clients et de votre projet Opportunités dans un troisième projet.

Par exemple, disons que vous avez un espace de travail "Projets" où vous gérez le travail réel pour vos clients. Vous souhaitez que cet espace de travail ait accès à la fois aux détails clients et aux informations sur les opportunités. Voici comment vous pouvez configurer cela :

Tout d'abord, créez un champ de référence dans votre espace de travail Projets qui se lie au projet Opportunités. Cela établit la connexion initiale. Ensuite, créez des champs de recherche basés sur cette référence pour tirer des détails spécifiques des opportunités, tels que la valeur de l'affaire ou la date de clôture prévue.

La véritable puissance réside dans l'étape suivante : vous pouvez créer des champs de recherche supplémentaires qui atteignent à travers la référence de l'opportunité vers la base de données clients. Cela vous permet de tirer des informations clients telles que les coordonnées ou le statut du compte directement dans votre espace de travail Projets.

Cette chaîne de connexions vous donne une vue complète dans votre espace de travail Projets, combinant des données à la fois de vos opportunités et de votre base de données clients. C'est un moyen puissant de garantir que vos équipes de projet disposent de toutes les informations pertinentes à portée de main sans avoir besoin de passer d'un projet à l'autre.

### Meilleures pratiques pour les systèmes CRM liés

Maintenez votre base de données clients comme la source unique de vérité pour toutes les informations clients. Chaque fois que vous devez mettre à jour les détails des clients, faites-le toujours d'abord dans la base de données clients. Cela garantit que les informations restent cohérentes dans tous les projets liés.

Lors de la création de champs de référence et de recherche, utilisez des noms clairs et significatifs. Cela aide à maintenir la clarté, surtout à mesure que votre système devient plus complexe.

Examinez régulièrement votre configuration pour vous assurer que vous tirez les informations les plus pertinentes. À mesure que les besoins de votre entreprise évoluent, vous pourriez avoir besoin d'ajouter de nouveaux champs de recherche ou de supprimer ceux qui ne sont plus utiles. Des examens périodiques aident à garder votre système rationalisé et efficace.

Envisagez de tirer parti des fonctionnalités d'automatisation de Blue pour garder vos données synchronisées et à jour à travers les projets. Par exemple, vous pourriez configurer une automatisation pour notifier les membres de l'équipe concernés lorsque des informations clés sur les clients sont mises à jour dans la base de données clients.

En mettant en œuvre efficacement ces stratégies et en utilisant pleinement les champs de référence et de recherche, vous pouvez créer un système CRM puissant et interconnecté dans Blue. Ce système vous fournira une vue complète à 360 degrés de vos relations clients et de votre pipeline de ventes, permettant une prise de décision plus éclairée et des opérations plus fluides dans toute votre organisation.

## Tableaux de bord

Les tableaux de bord sont un élément crucial de tout système CRM efficace, fournissant des aperçus instantanés de votre performance commerciale et de vos relations clients. La fonctionnalité de tableau de bord de Blue est particulièrement puissante car elle vous permet de combiner des données en temps réel provenant de plusieurs projets simultanément, vous offrant une vue complète de vos opérations de vente.

Lors de la configuration de votre tableau de bord CRM dans Blue, envisagez d'inclure plusieurs indicateurs clés. Le pipeline généré par mois montre la valeur totale des nouvelles opportunités ajoutées à votre pipeline, vous aidant à suivre la capacité de votre équipe à générer de nouvelles affaires. Les ventes par mois affichent vos affaires réellement conclues, vous permettant de surveiller la performance de votre équipe dans la conversion des opportunités en ventes.

Introduire le concept de remises sur le pipeline peut conduire à des prévisions plus précises. Par exemple, vous pourriez compter 90 % de la valeur des affaires dans la phase "Contrat en attente de signature", mais seulement 50 % des affaires dans la phase "Proposition envoyée". Cette approche pondérée fournit une prévision des ventes plus réaliste.

Suivre les nouvelles opportunités par mois vous aide à surveiller le nombre de nouvelles affaires potentielles entrant dans votre pipeline, ce qui est un bon indicateur des efforts de prospection de votre équipe de vente. Décomposer les ventes par type peut vous aider à identifier vos offres les plus réussies. Si vous configurez un projet de suivi des factures lié à vos opportunités, vous pouvez également suivre les revenus réels sur votre tableau de bord, fournissant une image complète de l'opportunité à l'encaissement.

Blue offre plusieurs fonctionnalités puissantes pour vous aider à créer un tableau de bord CRM informatif et interactif. La plateforme fournit trois types principaux de graphiques : cartes statistiques, graphiques à secteurs et graphiques à barres. Les cartes statistiques sont idéales pour afficher des indicateurs clés tels que la valeur totale du pipeline ou le nombre d'opportunités actives. Les graphiques à secteurs sont parfaits pour montrer la composition de vos ventes par type ou la distribution des affaires à travers différentes étapes. Les graphiques à barres excellent dans la comparaison des indicateurs au fil du temps, tels que les ventes mensuelles ou les nouvelles opportunités.

Les capacités de filtrage sophistiquées de Blue vous permettent de segmenter vos données par projet, liste, étiquette et période. Cela est particulièrement utile pour examiner des aspects spécifiques de vos données de vente ou comparer la performance entre différentes équipes ou produits. La plateforme consolide intelligemment les listes et les étiquettes portant le même nom à travers les projets, permettant une analyse fluide entre projets. Cela est inestimable pour une configuration CRM où vous pourriez avoir des projets séparés pour les clients, les opportunités et les factures.

La personnalisation est une force clé de la fonctionnalité de tableau de bord de Blue. La fonctionnalité de glisser-déposer et la flexibilité d'affichage vous permettent de créer un tableau de bord qui répond parfaitement à vos besoins. Vous pouvez facilement réorganiser les graphiques et choisir la visualisation la plus appropriée pour chaque indicateur.
Bien que les tableaux de bord soient actuellement réservés à un usage interne, vous pouvez facilement les partager avec les membres de l'équipe, en leur accordant des autorisations de consultation ou de modification. Cela garantit que tous les membres de votre équipe commerciale ont accès aux informations dont ils ont besoin.

En tirant parti de ces fonctionnalités et en incluant les indicateurs clés que nous avons discutés, vous pouvez créer un tableau de bord CRM complet dans Blue qui fournit des aperçus en temps réel de votre performance commerciale, de la santé de votre pipeline et de la croissance globale de votre entreprise. Ce tableau de bord deviendra un outil inestimable pour prendre des décisions basées sur les données et maintenir toute votre équipe alignée sur vos objectifs et vos progrès commerciaux.

## Conclusion

Configurer un CRM de vente complet dans Blue est un moyen puissant de rationaliser votre processus de vente et d'obtenir des aperçus précieux sur vos relations clients et votre performance commerciale. En suivant les étapes décrites dans ce guide, vous avez créé un système robuste qui intègre les informations clients, les opportunités de vente et les indicateurs de performance dans une plateforme unique et cohérente.

Nous avons commencé par créer une base de données clients, établissant une source unique de vérité pour toutes vos informations clients. Cette fondation vous permet de maintenir des enregistrements précis et à jour pour tous vos clients et prospects. Nous avons ensuite construit cela avec une base de données d'opportunités, vous permettant de suivre et de gérer efficacement votre pipeline de ventes.

L'un des principaux atouts de l'utilisation de Blue pour votre CRM est la capacité de lier ces bases de données à l'aide de champs de référence et de recherche. Cette intégration crée un système dynamique où les mises à jour des informations clients se reflètent instantanément dans toutes les opportunités liées, garantissant la cohérence des données et économisant du temps sur les mises à jour manuelles.
Nous avons exploré comment tirer parti des puissantes fonctionnalités d'automatisation de Blue pour rationaliser votre flux de travail, de l'attribution de nouveaux leads à l'envoi de rappels de suivi. Ces automatisations aident à garantir qu'aucune opportunité ne tombe dans l'oubli et que votre équipe peut se concentrer sur des activités à forte valeur ajoutée plutôt que sur des tâches administratives.

Enfin, nous avons approfondi la création de tableaux de bord qui fournissent des aperçus instantanés de votre performance commerciale. En combinant des données de vos bases de données clients et d'opportunités, ces tableaux de bord offrent une vue complète de votre pipeline de ventes, des affaires conclues et de la santé globale de votre entreprise.

N'oubliez pas que la clé pour tirer le meilleur parti de votre CRM est une utilisation cohérente et un raffinement régulier. Encouragez votre équipe à adopter pleinement le système, à revoir régulièrement vos processus et automatisations, et à continuer d'explorer de nouvelles façons de tirer parti des fonctionnalités de Blue pour soutenir vos efforts de vente.

Avec cette configuration de CRM de vente dans Blue, vous êtes bien équipé pour entretenir des relations clients, conclure plus d'affaires et faire avancer votre entreprise.
---
title: Pourquoi nous avons construit notre propre chatbot de documentation AI
description: Nous avons construit notre propre chatbot AI de documentation qui est formé sur la documentation de la plateforme Blue.
category: "Product Updates"
date: 2024-07-09
---


Chez Blue, nous cherchons toujours des moyens de faciliter la vie de nos clients. Nous avons [une documentation approfondie de chaque fonctionnalité](https://documentation.blue.cc), [des vidéos YouTube](https://www.youtube.com/@workwithblue), [des conseils et astuces](/insights/tips-tricks), et [divers canaux de support](/support).

Nous avons gardé un œil attentif sur le développement de l'IA (Intelligence Artificielle) car nous sommes très impliqués dans [les automatisations de gestion de projet](/platform/features/automations). Nous avons également lancé des fonctionnalités telles que [la catégorisation automatique par IA](/insights/ai-auto-categorization) et [les résumés par IA](/insights/ai-content-summarization) pour faciliter le travail de nos milliers de clients.

Une chose est claire : l'IA est là pour rester, et elle va avoir un effet incroyable dans la plupart des industries, et la gestion de projet ne fait pas exception. Nous nous sommes donc demandé comment nous pourrions tirer davantage parti de l'IA pour aider l'ensemble du cycle de vie d'un client, de la découverte, la pré-vente, l'intégration, et aussi avec les questions en cours.

La réponse était assez claire : **Nous avions besoin d'un chatbot AI formé sur notre documentation.**

Soyons honnêtes : *chaque* organisation devrait probablement avoir un chatbot. Ce sont d'excellents moyens pour les clients d'obtenir des réponses instantanées à des questions typiques, sans avoir à fouiller dans des pages de documentation dense ou sur votre site web. L'importance des chatbots sur les sites web de marketing modernes ne peut être sous-estimée.

![](/insights/ai-chatbot-regular.png)

Pour les entreprises de logiciels en particulier, il ne faut pas considérer le site web marketing comme une "chose" séparée — il *fait* partie de votre produit. Cela s'explique par le fait qu'il s'inscrit dans le parcours client typique :

- **Connaissance** (Découverte) : C'est là que les clients potentiels découvrent pour la première fois votre produit génial. Votre chatbot peut être leur guide amical, les orientant vers les fonctionnalités et avantages clés dès le départ.
- **Considération** (Éducation) : Maintenant, ils sont curieux et veulent en savoir plus. Votre chatbot devient leur tuteur personnel, fournissant des informations adaptées à leurs besoins et questions spécifiques.
- **Achat/Conversion** : C'est le moment de vérité - lorsque un prospect décide de se lancer et de devenir client. Votre chatbot peut résoudre les derniers petits problèmes, répondre à ces questions "juste avant d'acheter", et peut-être même proposer une offre intéressante pour conclure l'affaire.
- **Intégration** : Ils ont acheté, maintenant quoi ? Votre chatbot se transforme en un acolyte utile, guidant les nouveaux utilisateurs à travers la configuration, leur montrant les ficelles, et s'assurant qu'ils ne se sentent pas perdus dans le pays des merveilles de votre produit.
- **Rétention** : Garder les clients heureux est le nom du jeu. Votre chatbot est disponible 24/7, prêt à résoudre des problèmes, offrir des conseils et astuces, et s'assurer que vos clients se sentent appréciés.
- **Expansion** : Il est temps de passer à la vitesse supérieure ! Votre chatbot peut subtilement suggérer de nouvelles fonctionnalités, des ventes additionnelles ou croisées qui correspondent à la façon dont le client utilise déjà votre produit. C'est comme avoir un vendeur vraiment intelligent, non intrusif, toujours prêt à aider.
- **Plaidoyer** : Les clients satisfaits deviennent vos plus grands supporters. Votre chatbot peut encourager les utilisateurs satisfaits à faire passer le mot, laisser des avis, ou participer à des programmes de parrainage. C'est comme avoir une machine à enthousiasme intégrée dans votre produit !

## Décision de construire ou d'acheter

Une fois que nous avons décidé de mettre en œuvre un chatbot AI, la grande question suivante était : construire ou acheter ? En tant que petite équipe concentrée sur notre produit principal, nous préférons généralement les solutions "as-a-service" ou les plateformes open-source populaires. Après tout, nous ne sommes pas dans le métier de réinventer la roue pour chaque partie de notre pile technologique. 
Nous avons donc retroussé nos manches et plongé sur le marché, à la recherche de solutions de chatbot AI payantes et open-source.

Nos exigences étaient simples, mais non négociables :

- **Expérience sans marque** : Ce chatbot n'est pas juste un widget sympa ; il va sur notre site web marketing et éventuellement dans notre produit. Nous ne sommes pas enclins à faire de la publicité pour la marque de quelqu'un d'autre dans notre propre espace numérique.
- **Excellente UX** : Pour de nombreux clients potentiels, ce chatbot pourrait être leur premier point de contact avec Blue. Cela fixe le ton pour leur perception de notre entreprise. Soyons honnêtes : si nous ne pouvons pas réussir un chatbot correct sur notre site web, comment pouvons-nous espérer que les clients nous fassent confiance pour leurs projets et processus critiques ?
- **Coût raisonnable** : Avec une large base d'utilisateurs et des projets d'intégration du chatbot dans notre produit principal, nous avions besoin d'une solution qui ne nous ruinerait pas à mesure que l'utilisation augmente. Idéalement, nous voulions une **option BYOK (Bring Your Own Key)**. Cela nous permettrait d'utiliser notre propre clé de service AI, en payant uniquement les coûts variables directs au lieu d'une majoration à un fournisseur tiers qui ne gère pas réellement les modèles.
- **Compatible avec l'API OpenAI Assistants** : Si nous devions opter pour un logiciel open-source, nous ne voulions pas avoir à gérer un pipeline pour l'ingestion de documents, l'indexation, les bases de données vectorielles, et tout cela. Nous voulions utiliser l'[API OpenAI Assistants](https://platform.openai.com/docs/assistants/overview) qui abstrait toute la complexité derrière une API. Honnêtement — c'est vraiment bien fait.
- **Scalabilité** : Nous voulons avoir ce chatbot à plusieurs endroits, avec potentiellement des dizaines de milliers d'utilisateurs par an. Nous prévoyons une utilisation significative, et nous ne voulons pas être enfermés dans une solution qui ne peut pas évoluer avec nos besoins.

## Chatbots AI commerciaux

Ceux que nous avons examinés avaient tendance à avoir une meilleure UX que les solutions open-source — comme c'est malheureusement souvent le cas. Il y a probablement une discussion séparée à avoir un jour sur *pourquoi* de nombreuses solutions open-source ignorent ou sous-estiment l'importance de l'UX.

Nous fournirons ici une liste, au cas où vous chercheriez des offres commerciales solides :

- **[Chatbase](https://chatbase.co)** : Chatbase vous permet de créer un chatbot AI personnalisé formé sur votre base de connaissances et de l'ajouter à votre site web ou d'interagir avec lui via leur API. Il offre des fonctionnalités telles que des réponses fiables, la génération de leads, des analyses avancées, et la possibilité de se connecter à plusieurs sources de données. Pour nous, cela semblait être l'une des offres commerciales les plus abouties.
- **[DocsBot AI](https://docsbot.ai/)** : DocsBot AI crée des bots ChatGPT personnalisés formés sur votre documentation et contenu pour le support, la pré-vente, la recherche, et plus encore. Il fournit des widgets intégrables pour ajouter facilement le chatbot à votre site web, la possibilité de répondre automatiquement aux tickets de support, et une API puissante pour l'intégration.
- **[CustomGPT.ai](https://customgpt.ai)** : CustomGPT.ai crée une expérience de chatbot personnelle en ingérant vos données commerciales, y compris le contenu du site web, le service d'assistance, les bases de connaissances, les documents, et plus encore. Il permet aux leads de poser des questions et d'obtenir des réponses instantanées basées sur votre contenu, sans avoir besoin de chercher. Fait intéressant, ils [prétendent également battre OpenAI au RAG (Retrieval Augmented Generation) !](https://customgpt.ai/customgpt-beats-open-ai-in-rag-benchmark/)
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)** : Il s'agit d'une offre commerciale intéressante, car elle *est aussi* un logiciel open-source. Cela semble un peu en phase de démarrage, et les prix ne semblaient pas réalistes (27 $/mois pour des messages illimités ne fonctionnera jamais commercialement pour eux).

Nous avons également examiné [InterCom Fin](https://www.intercom.com/fin) qui fait partie de leur logiciel de support client. Cela aurait signifié passer de [HelpScout](https://wwww.helpscout.com) que nous avons utilisé depuis le début de Blue. Cela aurait pu être possible, mais InterCom Fin a des prix fous qui l'ont simplement exclu de la considération.

Et c'est en fait le problème avec de nombreuses offres commerciales. InterCom Fin facture 0,99 $ par demande de support client traitée, et ChatBase facture 399 $/mois pour 40 000 messages. C'est presque 5 000 $ par an pour un simple widget de chat.

Considérant que les prix pour l'inférence AI chutent comme jamais. OpenAI a réduit ses prix de manière assez spectaculaire :

- Le GPT-4 original (contexte 8k) était au prix de 0,03 $ par 1K tokens de prompt.
- Le GPT-4 Turbo (contexte 128k) était au prix de 0,01 $ par 1K tokens de prompt, soit une réduction de 50 % par rapport au GPT-4 original.
- Le modèle GPT-4o est au prix de 0,005 $ par 1K tokens, ce qui représente une nouvelle réduction de 50 % par rapport au prix du GPT-4 Turbo.

C'est une réduction de 83 % des coûts, et nous ne nous attendons pas à ce que cela reste stagnant.

Étant donné que nous recherchions une solution évolutive qui serait utilisée par des dizaines de milliers d'utilisateurs par an avec une quantité significative de messages, il est logique d'aller directement à la source et de payer les coûts de l'API directement, plutôt que d'utiliser une version commerciale qui augmente les coûts.

## Chatbots AI open-source

Comme mentionné, les options open-source que nous avons examinées étaient principalement décevantes en ce qui concerne l'exigence de "grande UX".

Nous avons regardé :

- **[Deepchat](https://deepchat.dev/)** : Il s'agit d'un composant de chat indépendant du framework pour les services AI, qui se connecte à diverses API AI, y compris OpenAI. Il a également la capacité pour les utilisateurs de télécharger un modèle AI qui fonctionne directement dans le navigateur. Nous avons joué avec cela et obtenu une version fonctionnelle, mais l'API OpenAI Assistants mise en œuvre semblait assez boguée avec plusieurs problèmes. Cependant, c'est un projet très prometteur, et leur playground est vraiment bien fait.
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)** : En regardant cela à nouveau d'un point de vue open-source, cela nécessiterait que nous mettions en place une infrastructure assez importante, quelque chose que nous ne voulions pas faire, car nous voulions nous appuyer autant que possible sur l'API Assistants d'OpenAI.

## Construire notre propre ChatBot

Et donc, n'ayant pas pu trouver quelque chose qui correspondait à toutes nos exigences, nous avons décidé de construire notre propre chatbot AI qui pourrait interagir avec l'API OpenAI Assistants. Cela s'est finalement avéré relativement indolore !

Notre site web utilise [Nuxt3](https://nuxt.com), [Vue3](https://vuejs.org/) (qui est le même framework que la plateforme Blue), et [TailwindUI](https://tailwindui.com/).

La première étape a été de créer une API (Application Programming Interface) dans Nuxt3 qui peut "parler" à l'API OpenAI Assistants. Cela était nécessaire car nous ne voulions pas tout faire sur le front-end, car cela exposerait notre clé API OpenAI au monde, avec le potentiel d'abus.

Notre API backend agit comme un intermédiaire sécurisé entre le navigateur de l'utilisateur et OpenAI. Voici ce qu'elle fait :

- **Gestion des conversations** : Elle crée et gère des "threads" pour chaque conversation. Pensez à un thread comme une session de chat unique qui se souvient de tout ce que vous avez dit.
- **Gestion des messages** : Lorsque vous envoyez un message, notre API l'ajoute au bon thread et demande à l'assistant d'OpenAI de rédiger une réponse.
- **Attente intelligente** : Au lieu de vous faire regarder un écran de chargement, notre API vérifie auprès d'OpenAI chaque seconde pour voir si votre réponse est prête. C'est comme avoir un serveur qui garde un œil sur votre commande sans déranger le chef toutes les deux secondes.
- **Sécurité d'abord** : En gérant tout cela sur le serveur, nous gardons vos données et nos clés API en sécurité.

Ensuite, il y avait le front-end et l'expérience utilisateur. Comme discuté précédemment, cela était *crucial*, car nous n'avons pas une seconde chance de faire une première impression !

Dans la conception de notre chatbot, nous avons accordé une attention méticuleuse à l'expérience utilisateur, en veillant à ce que chaque interaction soit fluide, intuitive et reflète l'engagement de Blue envers la qualité. L'interface du chatbot commence par un simple cercle bleu élégant, utilisant [HeroIcons pour nos icônes](https://heroicons.com/) (que nous utilisons sur l'ensemble du site web de Blue) pour agir comme notre widget d'ouverture de chatbot. Ce choix de design garantit une cohérence visuelle et une reconnaissance immédiate de la marque.

![](/insights/ai-chatbot-circle.png)

Nous comprenons que parfois les utilisateurs peuvent avoir besoin d'un soutien supplémentaire ou d'informations plus approfondies. C'est pourquoi nous avons inclus des liens pratiques dans l'interface du chatbot. Un lien par e-mail pour le support est facilement accessible, permettant aux utilisateurs de contacter directement notre équipe s'ils ont besoin d'une assistance plus personnalisée. De plus, nous avons intégré un lien vers la documentation, offrant un accès facile à des ressources plus complètes pour ceux qui souhaitent approfondir les offres de Blue.

L'expérience utilisateur est encore améliorée par des animations de fondu et de montée élégantes lors de l'ouverture de la fenêtre du chatbot. Ces animations subtiles ajoutent une touche de sophistication à l'interface, rendant l'interaction plus dynamique et engageante. Nous avons également mis en œuvre un indicateur de saisie, une petite mais cruciale fonctionnalité qui permet aux utilisateurs de savoir que le chatbot traite leur requête et rédige une réponse. Cet indice visuel aide à gérer les attentes des utilisateurs et maintient un sens de communication active.

<video autoplay loop muted playsinline>
  <source src="/videos/ai-chatbot-animation.mp4" type="video/mp4">
</video>

Reconnaissant que certaines conversations pourraient nécessiter plus d'espace à l'écran, nous avons ajouté la possibilité d'ouvrir la conversation dans une fenêtre plus grande. Cette fonctionnalité est particulièrement utile pour des échanges plus longs ou lors de la consultation d'informations détaillées, offrant aux utilisateurs la flexibilité d'adapter le chatbot à leurs besoins.

En coulisses, nous avons mis en œuvre un traitement intelligent pour optimiser les réponses du chatbot. Notre système analyse automatiquement les réponses de l'IA pour supprimer les références à nos documents internes, garantissant que les informations présentées sont claires, pertinentes et axées uniquement sur la réponse à la requête de l'utilisateur. 
Pour améliorer la lisibilité et permettre une communication plus nuancée, nous avons intégré un support markdown en utilisant la bibliothèque 'marked'. Cette fonctionnalité permet à notre IA de fournir un texte richement formaté, y compris des mises en gras et en italique, des listes structurées, et même des extraits de code si nécessaire. C'est comme recevoir un mini-document bien formaté et personnalisé en réponse à vos questions.

Enfin, mais certainement pas des moindres, nous avons donné la priorité à la sécurité dans notre mise en œuvre. En utilisant la bibliothèque DOMPurify, nous nettoyons le HTML généré à partir du parsing markdown. Cette étape cruciale garantit que tout script ou code potentiellement nuisible est supprimé avant que le contenu ne soit affiché. C'est notre façon de garantir que les informations utiles que vous recevez sont non seulement informatives mais aussi sûres à consommer.

## Développements futurs

Ceci n'est que le début, nous avons des choses passionnantes sur la feuille de route pour cette fonctionnalité.

L'une de nos fonctionnalités à venir est la capacité de diffuser des réponses en temps réel. Bientôt, vous verrez les réponses du chatbot apparaître caractère par caractère, rendant les conversations plus naturelles et dynamiques. C'est comme regarder l'IA réfléchir, créant une expérience plus engageante et interactive qui vous garde informé à chaque étape.

Pour nos précieux utilisateurs de Blue, nous travaillons sur la personnalisation. Le chatbot reconnaîtra lorsque vous êtes connecté, adaptant ses réponses en fonction de vos informations de compte, de votre historique d'utilisation et de vos préférences. Imaginez un chatbot qui non seulement répond à vos questions, mais comprend votre contexte spécifique au sein de l'écosystème Blue, fournissant une assistance plus pertinente et personnalisée.

Nous comprenons que vous pourriez travailler sur plusieurs projets ou avoir diverses requêtes. C'est pourquoi nous développons la capacité de maintenir plusieurs fils de conversation distincts avec notre chatbot. Cette fonctionnalité vous permettra de passer d'un sujet à l'autre sans perdre le fil – tout comme avoir plusieurs onglets ouverts dans votre navigateur.

Pour rendre vos interactions encore plus productives, nous créons une fonctionnalité qui proposera des questions de suivi suggérées en fonction de votre conversation actuelle. Cela vous aidera à explorer des sujets plus en profondeur et à découvrir des informations connexes que vous n'auriez peut-être pas pensé à demander, rendant chaque session de chat plus complète et précieuse.

Nous sommes également ravis de créer une suite d'assistants AI spécialisés, chacun adapté à des besoins spécifiques. Que vous cherchiez à répondre à des questions de pré-vente, à configurer un nouveau projet, ou à résoudre des fonctionnalités avancées, vous pourrez choisir l'assistant qui correspond le mieux à vos besoins actuels. C'est comme avoir une équipe d'experts Blue à portée de main, chacun spécialisé dans différents aspects de notre plateforme.

Enfin, nous travaillons à vous permettre de télécharger des captures d'écran directement dans le chat. L'IA analysera l'image et fournira des explications ou des étapes de dépannage en fonction de ce qu'elle voit. Cette fonctionnalité facilitera encore plus l'obtention d'aide pour des problèmes spécifiques que vous rencontrez lors de l'utilisation de Blue, comblant le fossé entre les informations visuelles et l'assistance textuelle.

## Conclusion

Nous espérons que cette plongée dans notre processus de développement de chatbot AI a fourni des informations précieuses sur notre réflexion en matière de développement produit chez Blue. Notre parcours, depuis l'identification du besoin d'un chatbot jusqu'à la construction de notre propre solution, illustre comment nous abordons la prise de décision et l'innovation.

![](/insights/ai-chatbot-modal.png)

Chez Blue, nous pesons soigneusement les options de construire ou d'acheter, toujours avec un œil sur ce qui servira le mieux nos utilisateurs et s'alignera avec nos objectifs à long terme. Dans ce cas, nous avons identifié un écart significatif sur le marché pour un chatbot à la fois rentable et visuellement attrayant qui pourrait répondre à nos besoins spécifiques. Bien que nous plaidions généralement pour tirer parti des solutions existantes plutôt que de réinventer la roue, parfois le meilleur chemin à suivre est de créer quelque chose de taillé sur mesure pour vos exigences uniques.

Notre décision de construire notre propre chatbot n'a pas été prise à la légère. Elle résulte d'une recherche de marché approfondie, d'une compréhension claire de nos besoins, et d'un engagement à fournir la meilleure expérience possible à nos utilisateurs. En développant en interne, nous avons pu créer une solution qui non seulement répond à nos besoins actuels, mais qui jette également les bases pour de futures améliorations et intégrations.

Ce projet illustre notre approche chez Blue : nous n'avons pas peur de retrousser nos manches et de construire quelque chose à partir de zéro lorsque c'est le bon choix pour notre produit et nos utilisateurs. C'est cette volonté de faire un effort supplémentaire qui nous permet de fournir des solutions innovantes qui répondent vraiment aux besoins de nos clients.
Nous sommes impatients de l'avenir de notre chatbot AI et de la valeur qu'il apportera aux utilisateurs potentiels et existants de Blue. Alors que nous continuons à affiner et à étendre ses capacités, nous restons engagés à repousser les limites de ce qui est possible dans la gestion de projet et l'interaction client.

Merci de nous avoir accompagnés dans ce voyage à travers notre processus de développement. Nous espérons que cela vous a donné un aperçu de l'approche réfléchie et centrée sur l'utilisateur que nous adoptons pour chaque aspect de Blue. Restez à l'écoute pour plus de mises à jour alors que nous continuons à faire évoluer et à améliorer notre plateforme pour mieux vous servir.

Si vous êtes intéressé, vous pouvez trouver le lien vers le code source de ce projet ici :

- **[ChatWidget](https://gitlab.com/bloohq/blue-website/-/blob/main/components/ChatWidget.vue)** : C'est un composant Vue qui alimente le widget de chat lui-même.
- **[Chat API](https://gitlab.com/bloohq/blue-website/-/blob/main/server/api/chat.post.ts)** : C'est le middleware qui fonctionne entre le composant de chat et l'API OpenAI Assistants.
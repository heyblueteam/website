---
title: Recherche en temps réel
category: "Mises à jour du produit"
description: Blue dévoile un nouveau moteur de recherche ultra-rapide qui retourne des résultats à travers tous vos projets en millisecondes, vous permettant de changer de contexte en un clin d'œil.
date: 2024-03-01
---

Nous sommes ravis d'annoncer le lancement de notre nouveau moteur de recherche, conçu pour révolutionner la façon dont vous trouvez des informations dans Blue. Une fonctionnalité de recherche efficace est cruciale pour une gestion de projet fluide, et notre dernière mise à jour garantit que vous pouvez accéder à vos données plus rapidement que jamais.

Notre nouveau moteur de recherche vous permet de rechercher tous les commentaires, fichiers, enregistrements, champs personnalisés, descriptions et listes de contrôle. Que vous ayez besoin de trouver un commentaire spécifique fait sur un projet, de localiser rapidement un fichier ou de rechercher un enregistrement ou un champ particulier, notre moteur de recherche fournit des résultats ultra-rapides.

Lorsque les outils approchent une réactivité de 50-100ms, ils tendent à s'effacer et se fondre en arrière-plan, offrant une expérience utilisateur fluide et presque invisible. Pour le contexte, un clignement d'œil humain prend environ 60-120ms, donc 50ms est en fait plus rapide qu'un clignement d'œil ! Ce niveau de réactivité vous permet d'interagir avec Blue sans même vous rendre compte qu'il est là, vous libérant pour vous concentrer sur le travail réel à accomplir. En tirant parti de ce niveau de performance, notre nouveau moteur de recherche garantit que vous pouvez accéder rapidement aux informations dont vous avez besoin, sans qu'il ne nuise jamais à votre flux de travail.

Pour atteindre notre objectif de recherche ultra-rapide, nous avons exploité les dernières technologies open-source. Notre moteur de recherche est construit sur MeiliSearch, un service de recherche open-source populaire qui utilise le traitement du langage naturel et la recherche vectorielle pour trouver rapidement des résultats pertinents. De plus, nous avons implémenté un stockage en mémoire, qui nous permet de stocker les données fréquemment consultées dans la RAM, réduisant le temps nécessaire pour retourner les résultats de recherche. Cette combinaison de MeiliSearch et de stockage en mémoire permet à notre moteur de recherche de fournir des résultats en millisecondes, vous permettant de trouver rapidement ce dont vous avez besoin sans jamais avoir à penser à la technologie sous-jacente.

La nouvelle barre de recherche est commodément située sur la barre de navigation, vous permettant de commencer à rechercher immédiatement. Pour une expérience de recherche plus détaillée, appuyez simplement sur la touche Tab pendant la recherche pour accéder à la page de recherche complète. De plus, vous pouvez rapidement activer la fonction de recherche depuis n'importe où en utilisant le raccourci CMD/Ctrl+K, rendant encore plus facile la recherche de ce dont vous avez besoin.

<video autoplay loop muted playsinline>
  <source src="/videos/search-demo.mp4" type="video/mp4">
</video>


## Développements futurs

Ce n'est que le commencement. Maintenant que nous avons une infrastructure de recherche de nouvelle génération, nous pouvons faire des choses vraiment intéressantes à l'avenir.

La prochaine étape sera la recherche sémantique, qui constitue une amélioration significative par rapport à la recherche par mots-clés typique. Permettez-nous de vous expliquer.

Cette fonctionnalité permettra au moteur de recherche de comprendre le contexte de vos requêtes. Par exemple, rechercher "mer" récupérera des documents pertinents même si l'expression exacte n'est pas utilisée. Vous pensez peut-être "mais j'ai tapé 'océan' à la place !" - et vous avez raison. Le moteur de recherche comprendra également la similarité entre "mer" et "océan", et retournera des documents pertinents même si l'expression exacte n'est pas utilisée. Cette fonctionnalité est particulièrement utile lors de la recherche de documents contenant des termes techniques, des acronymes, ou simplement des mots communs qui ont plusieurs variations ou fautes de frappe.

Une autre fonctionnalité à venir est la capacité de rechercher des images par leur contenu. Pour y parvenir, nous traiterons chaque image de votre projet, créant une représentation vectorielle pour chacune. En termes généraux, une représentation vectorielle est un ensemble mathématique de coordonnées qui correspond à la signification d'une image. Cela signifie que toutes les images peuvent être recherchées en fonction de ce qu'elles contiennent, indépendamment de leur nom de fichier ou de leurs métadonnées. Imaginez rechercher "organigramme" et trouver toutes les images liées aux organigrammes, *indépendamment de leurs noms de fichiers.*
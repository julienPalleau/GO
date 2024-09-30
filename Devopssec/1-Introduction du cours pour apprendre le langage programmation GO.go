
// 1-INTRODUCTION
// https://devopssec.fr/article/cours-apprendre-langage-programmation-go#begin-article-section
/*

Qu'est-ce que le langage go
Histoire

Go (prononcé en anglais "Gow") aussi appelé Golang est un langage de programmation open source relativement jeune, développé
2007 par Robert Griesemer, Rob Pike et Ken Thompson, travaillant aujourd'hui chez Google.

Déclaration des experts à l’origine de Go lors de son lancement

    « Chez Google, nous pensons que la programmation devrait être rapide, productive et surtout, fun. C’est pourquoi nous
	sommes ravis de proposer ce nouveau langage de programmation expérimental. Les opérations de compilation sont presque
	instantanées, et le code compilé propose une vitesse de fonctionnement proche de celle du C ».
    les responsables du projet Go.

Pour résumé leur message : Le langage de programmation Go rime avec efficacité et simplicité.

Le langage de programmation Go a été lancé en novembre 2009 et est utilisé dans certains systèmes de production de Google,
c’est celui la même qui a été utilisé pour développer le logiciel de conteneurisation Docker (c'est d'ailleurs avec le projet
Docker que j'ai connu pour la première fois ce langage).
Comparaison avec d'autres langages

Go est syntaxiquement similaire au langage C mais contrairement au C il possède une sécurité de la mémoire avec un Garbage
Collector et c'est un langage à typage statique (les types des variables sont connus lors de la compilation et doivent être
spécifiés expressément par le programmeur) comme dans d'autres langages de programmation à savoir le Java, le C ou le C++.

Go est souvent comparé au langage Python car tous les 2 se veulent très simples syntaxiquement. Personnellement je trouve
python est plus simple syntaxiquement mais Python reste tout de même un langage interprété, contrairement à Go qui est un
langage compilé.

Disclaimer : pas de guéguerre de langage ici :p, chaque langage de programmation a ses avantages et ses inconvénients, ça ne
m’empêche pas d'utiliser Python sur plusieurs projets voir même de combiner les deux langages, tout dépend simplement des
besoins 😉.
Rappel sur les langages interprétés et langages compilés

Certains d'entre vous ont peut être utilisé que des langages de programmation interprétés, si c'est le cas alors une petite
explication s'impose pour ainsi commencer avec des bonnes bases :

    Dans un langage interprété, le code source est interprété, par un autre logiciel nommé l’interpréteur, celui-ci traduit au
	fur et à mesure les instructions de votre programme.
    Dans un langage compilé, le code source est tout d'abord compilé en langage binaire c’est une suite de 0 et de 1 uniquement
	compréhensible par votre machine par un autre logiciel qu'on appelle le compilateur.

Avantages et inconvénients des langages interprétés et langages compilés

Un programme écrit dans un langage compilé a l'avantage de ne plus avoir besoin une fois compilé de programme annexe pour
s'exécuter (un langage interprété aura toujours besoin de son interpréteur), de plus comme votre code est exécuté directement
par votre machine alors le temps d’exécution de votre programme sera en général plus rapide pour le même programme dans un
langage interprété.

Cependant pour chaque modification de votre code source il faudra recompiler le programme pour que les modifications prennent
effet, ensuite votre programme compilé n’est pas multi-plateforme il faudra donc créer un exécutable pour chaque système
d’exploitation à l’inverse d’un l'inverse d’un langage interprété qui lui reste en général multi-plateforme.
Revenons à nos moutons !

Donc pour revenir à notre langage Go, c'est un langage qui se veut accessible et rapide pour une programmation à grande
échelle, il est donc concevable de l'utiliser aussi bien pour écrire des applications, des scripts ou sur d'autres types des
gros projets.
Les avantages du langage GO

Voici une courte liste d'avantages du langage GO :

    Une meilleure protection de la mémoire grâce à son ramasse-miettes (Garbage Collector) qui permet une gestion automatique
	de la mémoire.
    Profite de la puissance de calcul des processeurs les plus robustes du marché (processeurs multi-cœurs).
    Possibilité de faire du typage dynamique et intègre de nombreux types avancés tels que les mappages clé-valeur.
    Possède une riche bibliothèque standard, qu’il est même tout à fait possible de concevoir des programmes écrit avec le
	langage Go sans aucune dépendance externe.
    Une base de code propre nécessaire aussi pour assurer la maintenance et l'évolution des programmes sur plusieurs
	générations développeurs.
    Possède un temps de compilation rapide et intègre aussi un système de build beaucoup moins compliqué que celui de la
	plupart des langages de compilés.
    Au niveau de la portabilité il est possible de compiler votre code pour une large gamme de systèmes d'exploitation et de
	plateformes matérielles (Windows, Linux, MAC OS, Android, IOS).

Utilisation du langage GO

On retrouve le langage Go dans les domaines suivants (liste non exhaustive) :

    Serveurs
    Web
    Systèmes embarqués
    IOT (Internet Of Things)
    Android
    IOS
    Jeux-vidéos
    etc ...

Liste non exhaustive des entreprises utilisant Go :

    CloudFlare
    CoreOS
    DropBox
    Docker
    Nokia
    Ovh
    YouTube
    SoundCloud
    Splice
    etc ...

Public visé

Ce tutoriel est conçu pour des programmeurs ou des curieux ayant besoin de comprendre le langage de programmation Go à partir de zéro. Ce tutoriel vous donnera une compréhension suffisante du langage de programmation Go, qui vous permettra d’atteindre des niveaux d’expertise plus élevés.
*/
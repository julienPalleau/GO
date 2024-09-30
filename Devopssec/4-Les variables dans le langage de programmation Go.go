// 4 - Les variables dans le langage de programmation Go
// https://devopssec.fr/article/variables-golang#begin-article-section

/*
Ce chapitre vous explique les variables et les constantes en GoLang. Vous allez voir les différents types de variables et
d'opérateurs mais aussi apprendre à déclarer, modifier et calculer vos variables dans le langage de programmation GO.

1 - Définition

Avant de vous montrer comment déclarer et utiliser une variable, il faut d'abord comprendre ce qu'est une variable.

Rappel

"Go est un langage à typage statique (les types des variables sont connus lors de la compilation et doivent être spécifiés
expressément par le programmeur)"

Une variable est le nom donné à un emplacement mémoire pour stocker une valeur, elle permet de stocker des données, ces données
vont pouvoir être utilisées plus tard.

Une variable est constituée d'un type. Le type d'une variable permet :

    De déterminer la quantité d'espace occupé dans la mémoire.
    D'avertir le compilateur Go de quelle manière il doit interpréter votre variable

 2 - Les différents types de variables

il existe une multitude de types de variables en go, nous allons voir les types les plus utilisés.
Types entiers

Ce type de variable permet de stocker des nombres entiers : 1, 2, 3, 4…

Il existe deux catégories d'entiers :

    unsigned (non signé) : permet de stocker que des nombres positifs
    signed (signé): permet de stocker des nombres positifs et négatifs

Il y a plusieurs types d'entiers possibles, tout dépend de la valeur que vous voulez stocker dans votre variable. Je vous présente ci-dessous un tableau qui liste les types d'entiers qu'il est possible d'utiliser en GoLang ainsi que les nombres qu'ils peuvent conserver.
Type 	Taille 	Valeurs possibles
uint8 	8 bits 	0 à 255
uint16 	16 bits 	0 à 65535
uint32 	32 bits 	0 à 4294967295
uint64 	64 bits 	0 à 18446744073709551615
int8 	8 bits 	-128 à 127
int16 	16 bits 	-32768 à 32767
int32 	32 bits 	-2147483648 à 2147483647
int64 	64 bits 	-9223372036854775808 à 9223372036854775807

Information

le bit est la plus petite valeur informatique : soit 0 soit 1.
Types flottants

Ce type permet de stocker des nombres décimaux (nombres à virgule) : 1.5, 1.6, 128.5
Type 	Taille 	Précision
float32 	32 bits 	Environ 7 chiffres décimaux
float64 	64 bits 	Environ 16 chiffres décimaux
Autres types numériques

Il existe également un ensemble de types numériques avec des tailles spécifiques :
Type 	description
byte 	Identique à uint8
rune 	Identique à int32
uint 	32 ou 64 bits non signé
int 	Même taille que uint mais signé
Types booléens

Un booléen est un type de variables qui possède deux états :

    L'état « vrai » : true en anglais
    Ou bien l'état « faux » : false en anglais

C'est un type qui va nous permettre par exemple dans un jeu de savoir si est un joueur est vivant ou pas.

On utilise le mot-clé bool pour déclarer une variable de type booléen
Type string

C'est un type qui nous permet de stocker des chaines de caractères (du texte).

On utilise le mot-clé string pour déclarer une variable de type chaines de caractères

Information

Je n'ai cité que les variables plus utilisées mais il faut savoir qu'il existe d'autres types de variables que je ne vais pas détailler dans cette séance.

*/

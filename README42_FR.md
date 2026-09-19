*This project has been created as part of the 42 curriculum by alebaron, rruiz, emarette.*

# The Answer Protocol : Une aventure textuelle rétro se déroulant dans un univers partagé

## Description

The Answer Protocol est une aventure textuelle multijoueur inspirée des anciens MUD (Multi-User Dungeons), où plusieurs joueurs partagent un même monde virtuel, se déplacent entre salles, interagissent avec des PNJ, recherchent des objets, achètent et vendent des ressources, accomplissent des quêtes et combattent des monstres.

Le projet a pour objectif de reproduire une expérience de jeu en réseau simple et robuste, entièrement basée sur le protocole TCP et des commandes textuelles, tout en proposant deux interfaces d’accès : un client en ligne de commande et un client graphique. Le cœur du système est un serveur central qui maintient l’état des joueurs, des pièces, des objets et des événements du monde.

Le thème général s’inspire du rétro-gaming et du fonctionnement des mondes partagés textuels des années 1990, avec une logique de jeu orientée sur l’exploration, la coopération, les échanges et les mécaniques d’aventure.

## Instructions

### Les éléments obligatoires

- **Langage** : Go 1.22 ou plus récent
- **Compilateur** : Un compilateur C (GCC, Clang, ...) pour utiliser Fyne

### Les commandes importantes

```bash
make install
make build
make run-server
make run-client
```

## Architecture

L’architecture du serveur est structurée autour d’un modèle de gestion centralisée de l’état du monde et d’un dispatcher de commandes.

- Le serveur écoute les connexions TCP sur le port 8090.
- Pour chaque client connecté, un goroutine distinct est créé dans `handleConnection`, ce qui permet de traiter plusieurs joueurs en parallèle.
- Les commandes envoyées par les clients sont découpées en tokens, puis transmises à une table de correspondance via une fonction `dispatch`.
- Chaque commande est associée à une fonction dédiée dans le package `src/server/commands`.
- Le monde est représenté par un objet `TapManager` qui regroupe les joueurs, les salles, les PNJ, les monstres, les quêtes et les groupes.

Le choix d’un dispatcher centralisé est motivé par plusieurs avantages :

- une séparation claire entre le protocole réseau et la logique métier ;
- un ajout simple de nouvelles commandes sans modifier le cœur du serveur ;
- une gestion plus lisible des erreurs et des logs ;
- une meilleure évolutivité pour un projet multijoueur.

## Protocol Implementation

### Fonctionnement du protocol

Le serveur implémente un protocole texte simple, orienté ligne, inspiré du principe de la commande réseau et de la réponse serveur. Les messages transitent sous forme de chaînes ASCII terminées par un retour à la ligne.

Les échanges suivent un modèle pragmatique :

- envoi de `OK ...` pour confirmer une action valide ;
- envoi de `ERR ...` pour signaler une erreur ou un cas invalide ;
- envoi de `EVT ...` pour notifier un événement global ou une mise à jour d’état ;
- gestion de la connexion via `CONNECT [Name] [Language]` puis utilisation des commandes du jeu.

Le projet s’écarte volontairement d’une lecture stricte d’une RFC complète : le protocole est plus léger et plus lisible, en privilégiant la robustesse et la simplicité de parsing côté Go. Cette décision est motivée par le besoin de maintenir un système de commande réactif, facilement exploitable par un client CLI et un client GUI, sans imposer une overhead de trames complexes ou de sérialisation binaire.

La logique de validation est également adaptée au contexte du jeu : certaines erreurs sont retournées sous forme de codes de type `ERR 404`, `ERR 401`, `ERR 402`, etc., afin de conserver une convention cohérente et de laisser le client afficher un message explicite à l’utilisateur.

### Dictionnaire des erreurs

En plus des erreurs déjà présentes dans le protocole fourni par le sujet, nous avons pris la décision de rajouter certains codes afin de conserver un retour d'erreurs clair et compréhensible pour les utilisateurs.

```
201 NAME_IN_USE             Le nom d'utilisateur demandé est déjà pris
202 INCORRECT_LANGUAGE      Langue demandée inconnue
203 TO_MANY_PLAYER          Trop de joueurs connectés

301 NO_EXIT                 Direction de déplacement invalide

401 NOT_IN_GROUP            L'opération de groupe nécessite d'être membre d'un groupe
401 GROUP_ALREADY_FULL      Le groupe est complet (4 personnes)
402 ALREADY_IN_GROUP        Le joueur appartient déjà à un groupe
402 PLAYER_NOT_INVITED      Le joueur n'a pas reçu d'invitation
403 ALREADY_INVITED         Invitation déjà envoyée au joueur
404 ITEM_NOT_FOUND          L'objet demandé n'est pas disponible dans la pièce
404 ITEM_NOT_IN_INVENTORY   L'objet demandé n'est pas dans l'inventaire du joueur
404 NPC_NOT_FOUND           Le PNJ demandé n'est pas présent dans la pièce
404 PLAYER_NOT_FOUND        Le joueur demandé n'est pas présent
404 WEAPON_NOT_FOUND        L'arme demandée n'est pas présente
404 TARGET_NOT_FOUND        La cible demandée n'est pas présente
404 GROUP_NOT_FOUND         Le groupe demandé n'est pas présent
404 ROOM_NOT_FOUND          La pièce demandée n'est pas présente
405 NPC_NOT_HOSTILE         Le PNJ ne peut pas être attaqué (ce n'est pas un ennemi)
405 NPC_NOT_TRADER          Le PNJ ne peut pas commercer (ce n'est pas un marchand)
406 NO_QUEST_AVAILABLE      Le PNJ n'a pas de quêtes ou la quête est déjà terminée
407 NOT_ENOUGH_ITEM         Il n'y a pas assez d'objets
407 NOT_ENOUGH_COIN         Il n'y a pas assez de pièces de jeu
407 NOT_ENOUGH_MONEY        Il n'y a pas assez d'argent
408 INVALID_LOCATION        Mauvais endroit pour utiliser / agir
408 INVALID_BET             Mise de jeu invalide
408 INVALID_DIRECTION       Direction inconnue
408 INVALID_QUANTITY        Quantité invalide
408 INVALID_ITEM            Objet invalide
409 FAILED_TO_SUMMON        Tentative d'apparition de monstre échouée

900 CONNECTION_FAILED       Échec de l'établissement de la connexion
901 SEND_FAILED             Échec de la transmission du message
902 COMMAND_UNKNOWN         Commande inconnue
903 COMMAND_EMPTY           Aucune commande envoyée
904 WRONG_COMMAND_ARG       Les arguments passés ne sont pas corrects
905 TOO_MANY_REQUESTS       L'utilisateur spam la même commande

999 GAMBLING_ROLL_FAILED    Problème de tirage avec le système de jeu
999 FISHING_ROLL_FAILED     Problème de tirage avec la pêche
999 MARSHAL_ENCODING_ERROR  Problème avec le sérialiseur (marshal) Golang
999 INVALID_LOOT_TABLE      Problème avec la table de butin
```

## Combat System

[TODO: Enzo] (A relire et à étoffer)

Le système de combat est tour par tour s'inspirant de celui présent dans le jeu de rôle **Donjon & Dragons**. Il s’appuie sur un principe de jet de dé et de comparaison de statistiques. 

### Mécanique principale

- le joueur cible un monstre présent dans la salle courante ;
- un jet d’attaque est simulé par un nombre aléatoire entre 1 et 20 ;
- si le jet est supérieur à la défense de la cible, le coup touche ;
- un résultat de 20 déclenche un coup critique ;
- les dégâts sont calculés à partir de l’attaque du joueur ou de l’arme équipée.

Dans le code actuel, la formule de dégâts est typiquement de la forme :

- dégâts simples : valeur aléatoire autour de l’attaque de l’arme ou du personnage ;
- dégâts critiques : montant plus élevé, avec un bonus supplémentaire ;
- la cible peut être réduite à 0 PV, ce qui la retire de l’arène et déclenche la génération de loot dans la salle.

### Réponse du monstre

Le monstre contre-attaque immédiatement après l’attaque du joueur si le monstre est encore vivant. Si le joueur appartient à un groupe, la cible de la riposte peut être choisie aléatoirement parmi les membres du groupe présents dans la salle, ce qui apporte une petite logique de coalition et de gestion d’attaque de groupe. 

### Mort et récupération

- lorsqu’un personnage tombe à 0 PV, son état est mis à `dead` ;
- il est renvoyé dans une salle de respawn et quitte son groupe ;

### Gestion des combats

Lorsqu'un monstre apparaît dans une salle après avoir été invoqué, n'importe quel joueur présent est en capacité de pour l'attaquer. Une fois le monstre mort, les objets qu'il avait tombe alors au sol dans la salle et les joueurs présents peuvent utiliser `TAKE` pour les récupérer. 

### Commandes associées

Le système comporte les commandes suivantes :

- `ATTACK <monster_id>`
- `ATTACK <monster_id> <weapon_id>`
- `STATUS` pour suivre les PV et l’état du personnage
- `SEARCH <monster_id>` pour faire apparaître un monstre dans la salle

## Quest System

Le système de quêtes repose sur deux types de quêtes : les quêtes d’objets et les quêtes de monstres.

### Types de quêtes

- `QuestItem` : le joueur doit récupérer un certain nombre d’objets spécifiques ;
- `QuestMonster` : le joueur doit éliminer un type de monstre donné.

### Progression

Les quêtes suivent un cycle de statut :

- `active`
- `in progress`
- `completed`
- `rewarded`

La progression est mise à jour à chaque interaction valide, que ce soit lorsqu’un objet est ajouté ou retiré de l’inventaire, ou lorsqu’un monstre est abattu. Le moteur limite automatiquement la progression entre 0 et la quantité demandée.

### Validation et récompense

Lorsqu’un joueur parle à un NPC donneur de quête, le moteur vérifie si la quête est terminée. Si la condition est satisfaite :

- le joueur reçoit la récompense associée à la quête ;
- le statut de la quête passe à `rewarded` ;
- les objets requis peuvent être retirés de l’inventaire pour les quêtes de type item ;
- un log de quête est enregistré dans le système de logs serveur.

Une fois la quête validée, il n'est plus possible de l'obtenir à nouveau.

Les commandes côté client/serveur associées sont :

- `QUEST <npc_id>` pour accepter une quête ;
- `QUESTS` pour afficher la liste des quêtes actives ;
- `TALK <npc_id>` pour déclencher le dialogue et la validation de quête.

## World Design

Le monde du jeu est bâti à partir de fichiers JSON chargés au démarrage du serveur. La structure du monde est répartie en salles, objets, monstres, PNJ et liens de déplacement. Les aventuriers se retrouveront à leur connexion dans la place centrale du village de **Sangluten-Sur-Pain** et pourront découvrir l'univers qui leur est proposé en explorant les différentes actions possibles.

### Layout général

Le monde comporte plusieurs zones distinctes :

- village,
- forêt,
- lac,
- grotte,
- casino.

Les salles sont reliées entre elles par un système de voisins (`north`, `south`, `east`, `west`) et portent des coordonnées de position `(x, y)` pour structurer visuellement le monde.

### Rôles des PNJ

Différents types de personnages sont présents :

- `QuestGiver` : donne des missions ;
- `Trader` : vend et achète des objets ;
- `Dialoguer` : dialogue générique, parfois lié à l’histoire ou au décor.

### Répartition des objets et monstres

- les objets sont répartis dans les pièces ou dans l’inventaire des traders ;
- les monstres sont associés à des salles prédéfinies et peuvent être invoqués via `SEARCH` ;
- le loot tombé après un combat est ajouté dans la salle pour être ramassé.

### Exemple de structure du monde

La carte contient notamment des lieux tels que :

- `CASINO`
- `Place Sénile`
- `Taverne de la pucelle`
- `Forêt, épisode 1`
- `Forêt fantôme`
- `Lac Koulofon`
- `Grotte Terrifiante`
- `...`

La conception du monde privilégie un mélange entre exploration libre et progression orientée quêtes, avec des zones plus hostiles et des zones plus sûres.

## Interface graphique(GUI)

Le GUI a été créé avec **Fyne** et se compose de 2 écrans : la **HomeView**, qui gère l'accueil et les connexions au jeu, et la **GameView**, qui est le jeu en lui-même. 

### HomeView

C'est l'écran affiché lorsqu'on lance le GUI, il contient un **formulaire de connexion**. Ce formulaire est composé de :

- Un sélecteur entre français et anglais pour l'affichage des informations du GUI.

-  Un champ pour mettre le pseudonyme que le joueur souhaite utiliser, il doit faire entre 1 et 15 caractères et doit être composé uniquement de caractères alphanumériques et d’underscore.

- Un champ pour l'IP du serveur que le joueur souhaite rejoindre.

Quand le joueur clique sur JOIN server, l'IP est utilisée pour essayer de se connecter au serveur, si ça marche, CONNECT <name> <language> est envoyé au serveur. Si le serveur répond "OK connected", l'écran change pour afficher le GameView.

### GameView

C'est le jeu, au démarrage, elle récupère les infos globales grâce à SECRET et au listener mis en place. Elle est composée de différents widgets qui fonctionnent séparément. Ces widgets sont :

- Commandes (command_widget.go) : créer une barre de boutons de catégorie, cliquer sur un bouton affiche les éléments de cette catégorie qui sont les commandes disponibles dans le jeu. Ces commandes sont aussi des boutons, cliquer dessus permet au besoin de choisir les paramètres et ensuite de l'envoyer au serveur pour l'exécuter.

- Log (log_widget.go) : le journal des messages, rangés par catégorie (GLOBAL, ROOM, GROUP, LOG).

- Map (map_widget.go) : dessine les salles avec des couleurs selon si elles sont la salle actuelle, visitée, connue ou inconnue. Ce widget est mis à jour à chaque réponse LOOK.

- GoingOn, la salle actuelle (goingOn_widget.go) : affiche la salle dans laquelle le joueur se trouve: son image est affichée et son nom est mis juste au dessus.

- Players (players_widget.go) : affiche le nombre de joueurs dans la room et le nombre total de joueurs sur le serveur. Nécessite un premier WHO pour s'afficher et actualise le nombre de joueurs dans la room grâce à l'événement (EVT ROOM PRESENCE ENTER/LEAVE).

- Groupe (group_widget.go) :  affiche si le joueur est dans un groupe ou non, si oui, la liste de ses membres est aussi affichée. Il se rafraîchit via SECRET quand le serveur émet un événement EVT GROUP.

### Synchronisation avec le serveur

Le Listener est un petit système pub/sub : les widgets et le state global s'abonnent à des lignes précises. Chaque ligne serveur est envoyée à tous les abonnés par Distribute.

Le state global est une copie locale du SECRET serveur. À chaque réponse OK SECRET, setGameData met à jour les données puis déclenche le rafraîchissement des widgets qui en dépendent.

## Server Logging

Le système de logging est un composant important du projet. Il est géré dans le package `src/server/server_write`.

### Implémentation

Au démarrage du serveur, le programme crée le dossier `log/` et initialise le fichier `log/server_log.txt` si nécessaire. Puis, à chaque événement, la fonction `WriteLog` enregistre une ligne avec le format suivant :

`timestamp (IP) [LEVEL]: message`

### Types d’événements

Les logs couvrent plusieurs catégories :

- `INFO` : connexion, déconnexion, événements liés à l’état du monde ;
- `COMMAND` : commandes soumises par les joueurs ;
- `WARN` : erreurs ou messages d’alerte de traitement ;
- `ERROR` : connexion ou lecture cassée ;
- `SERVER` : réponses envoyées au client ;
- `WORLD` : actions environnementales ou combat / loot ;
- `QUEST` : progression des quêtes ;
- `GROUP` : création, invitation, rejoindre et quitter un groupe ;
- `CHAT` : messages mutuels entre joueurs.

### Monitoring

Pour surveiller le comportement du serveur, il suffit de suivre le fichier log en temps réel :

```bash
tail -f log/server_log.txt
```

Cette approche permet de détecter rapidement :

- des abus de commandes,
- des comportements répétitifs ou trompeurs,
- des tentatives de connexion invalide,
- des événements de combat, de quête et de groupe non attendus.

## Group Contributions

### Repartitions des rôles dans notre groupe

- `alebaron` : Gestion du monde, parsing des données JSON et implémentation des commandes de jeu ;
- `rruiz` : Intégration GUI, Conception du Monde et Test & Debug ;
- `emarette` : Mise en place du serveur, implémentation des commandes du jeu et Happiness Manager.

La répartition est cohérente avec la structure du dépôt : le moteur principal du serveur est centralisé dans les packages `src/server`, alors que les clients et le rendu sont répartis entre les couches `src/gui` et les composants de jeu.

## Building and Running

[TODO: Enzo]

## Testing

Le projet ne contient pas de suite de tests automatisés Go dans le dépôt, mais le système multijoueur et les mécaniques de jeu peuvent être validés manuellement en exécutant plusieurs clients sur le même serveur.

### Test du multijoueur

1. Démarrer le serveur.
2. Ouvrir deux ou plusieurs sessions `nc` vers le port 8090.
3. Se connecter avec des pseudos différents :

```text
CONNECT alice FR
CONNECT bob EN
```

4. Vérifier que les joueurs apparaissent dans le monde, que `LOOK` renvoie l’état de la salle, et que les événements de présence sont bien diffusés.

### Test du combat

1. Se déplacer vers une salle contenant un monstre.
2. Utiliser `SEARCH <monster_id>` pour faire apparaître le monstre.
3. Exécuter :

```text
ATTACK <monster_id>
```

4. Vérifier :
   - la variation des PV ;
   - la génération de loot après la mort du monstre ;
   - les logs dans `log/server_log.txt` ;
   - la réponse `OK ...` du serveur.

### Test des quêtes

1. Trouver un `QuestGiver` dans une salle.
2. Exécuter :

```text
QUEST <npc_id>
TALK <npc_id>
QUESTS
```

3. Vérifier que la quête passe aux bons statuts et que la récompense est attribuée lorsque la condition est remplie.

### Test des groupes

```text
GROUP CREATE
GROUP INVITE bob
GROUP JOIN <group_id>
GROUP LEAVE
```

Vérifier l’envoi des événements de groupe et l’évolution de l’état du joueur.

## Ressources

### Outils de travail d’équipe

- [Excalidraw](https://excalidraw.com/)
- [Google Sheets](https://www.google.com/sheets/about/)
- [Product backlog overview](https://asana.com/resources/product-backlog)

### Références Golang

- [Go Documentation](https://go.dev/doc/)
- [Fyne Documentation](https://docs.fyne.io/)
- [Go pkg documentation for Fyne](https://pkg.go.dev/fyne.io/fyne/v2)
- [MUDs and online text-based worlds](https://fr.wikipedia.org/wiki/Multi-user_dungeon)
- [Tutoriel Golang : apprendre le langage Go (débutant)](https://blog.stephane-robert.info/docs/developper/programmation/golang/)
- [La programmation orientée objet dans le langage de programmation Go](https://devopssec.fr/article/programmation-orientee-objet-golang)
- [Parsing JSON files With Go](https://tutorialedge.net/golang/parsing-json-with-golang/)
- [How to read a file and convert JSON to Go Struct](https://dev.to/mxglt/wip-how-to-read-a-file-and-convert-json-to-go-struct-6m2)

### Référence MUD / RPG

- [Fer & Flamme](https://fr.wikipedia.org/wiki/Fer_et_Flamme)
- [Donjons & Dragons](https://fr.wikipedia.org/wiki/Donjons_et_Dragons)
- [Wakfu (Jeu vidéo)](https://fr.wikipedia.org/wiki/Wakfu_(jeu_vid%C3%A9o))

### Utilisation de l’IA

L’intelligence artificielle a été utilisée de manière ciblée pour :

- aider à la reformulation, à la structuration et à la traduction du README ;
- aider à la correction de certains bugs rencontrés.

---

**Dernière modification**: 16 Septembre 2026\
**Contact :** alebaron@student.42lehavre.fr / rruiz@student.42lehavre.fr / emarette@student.42lehavre.fr


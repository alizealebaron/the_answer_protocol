*This project has been created as part of the 42 curriculum by alebaron, rruiz, emarette.*

# The Answer Protocol : A retro text adventure set in a shared universe

## Description

The Answer Protocol is a multiplayer text adventure inspired by the old MUDs (Multi-User Dungeons), where several players share the same virtual world, move between rooms, interact with NPCs, search for items, buy and sell resources, complete quests and fight monsters.

The goal of the project is to reproduce a simple and robust networked gaming experience, entirely based on the TCP protocol and text commands, while offering two access interfaces: a command-line client and a graphical client. The core of the system is a central server that maintains the state of the players, rooms, items and world events.

The overall theme is inspired by retro gaming and by the way text-based shared worlds worked in the 1990s, with game logic oriented towards exploration, cooperation, trading and adventure mechanics.

## Instructions

### Requirements

- **Language**: Go 1.22 or newer
- **Compiler**: A C compiler (GCC, Clang, ...) to use Fyne

### Important commands

| Command | Description |
|---|---|
| `make build` | Compiles the project and generates the `main` executable |
| `make run-server` | Compiles then launches the server |
| `make run-gui` | Compiles then launches the graphical client (GUI) |
| `make run-cli IP=<address>` | Launches the command-line client |
| `make clean` | Removes the `main` executable and the `log` folder |

## Architecture

The server architecture is structured around a centralized world state management model and a command dispatcher.

- The server listens for TCP connections on port 8090.
- For each connected client, a separate goroutine is created in `handleConnection`, which allows several players to be handled in parallel.
- The commands sent by the clients are split into tokens, then passed to a lookup table via a `dispatch` function.
- Each command is associated with a dedicated function in the `src/server/commands` package.
- The world is represented by a `TapManager` object that groups together players, rooms, NPCs, monsters, quests and groups.

The choice of a centralized dispatcher is motivated by several advantages:

- a clear separation between the network protocol and the business logic;
- easy addition of new commands without modifying the server core;
- more readable error and log handling;
- better scalability for a multiplayer project.

## Protocol Implementation

### How the protocol works

The server implements a simple, line-oriented text protocol, inspired by the principle of a network command and a server response. Messages travel as ASCII strings terminated by a newline.

Exchanges follow a pragmatic model:

- sending `OK ...` to confirm a valid action;
- sending `ERR ...` to signal an error or an invalid case;
- sending `EVT ...` to notify of a global event or a state update;
- connection handling via `CONNECT [Name] [Language]`, followed by the use of game commands.

The project deliberately departs from a strict reading of a full RFC: the protocol is lighter and more readable, favoring robustness and simplicity of parsing on the Go side. This decision is motivated by the need to maintain a responsive command system that can easily be used by a CLI client and a GUI client, without imposing the overhead of complex frames or binary serialization.

The validation logic is also adapted to the game context: some errors are returned as codes such as `ERR 404`, `ERR 401`, `ERR 402`, etc., in order to keep a consistent convention and let the client display an explicit message to the user.

### Error dictionary

In addition to the errors already present in the protocol provided by the subject, we decided to add some codes in order to keep a clear and understandable error feedback for users.

```
201 NAME_IN_USE             The requested username is already taken
202 INCORRECT_LANGUAGE      Unknown requested language
203 TO_MANY_PLAYER          Too many players connected

301 NO_EXIT                 Invalid movement direction

401 NOT_IN_GROUP            The group operation requires being a member of a group
401 GROUP_ALREADY_FULL      The group is full (4 people)
402 ALREADY_IN_GROUP        The player already belongs to a group
402 PLAYER_NOT_INVITED      The player has not received an invitation
403 ALREADY_INVITED         Invitation already sent to the player
404 ITEM_NOT_FOUND          The requested item is not available in the room
404 ITEM_NOT_IN_INVENTORY   The requested item is not in the player's inventory
404 NPC_NOT_FOUND           The requested NPC is not present in the room
404 PLAYER_NOT_FOUND        The requested player is not present
404 WEAPON_NOT_FOUND        The requested weapon is not present
404 TARGET_NOT_FOUND        The requested target is not present
404 GROUP_NOT_FOUND         The requested group is not present
404 ROOM_NOT_FOUND          The requested room is not present
405 NPC_NOT_HOSTILE         The NPC cannot be attacked (it is not an enemy)
405 NPC_NOT_TRADER          The NPC cannot trade (it is not a merchant)
406 NO_QUEST_AVAILABLE      The NPC has no quests or the quest is already completed
407 NOT_ENOUGH_ITEM         There are not enough items
407 NOT_ENOUGH_COIN         There are not enough game coins
407 NOT_ENOUGH_MONEY        There is not enough money
408 INVALID_LOCATION        Wrong place to use / act
408 INVALID_BET             Invalid gambling bet
408 INVALID_DIRECTION       Unknown direction
408 INVALID_QUANTITY        Invalid quantity
408 INVALID_ITEM            Invalid item
409 FAILED_TO_SUMMON        Monster summoning attempt failed

900 CONNECTION_FAILED       Failed to establish the connection
900 DECONNECTION_FAILED     Failed to establish the deconnection
901 SEND_FAILED             Failed to transmit the message
902 COMMAND_UNKNOWN         Unknown command
903 COMMAND_EMPTY           No command sent
904 WRONG_COMMAND_ARG       The arguments passed are not correct
905 TOO_MANY_REQUESTS       The user is spamming the same command

999 GAMBLING_ROLL_FAILED    Roll problem with the gambling system
999 FISHING_ROLL_FAILED     Roll problem with fishing
999 MARSHAL_ENCODING_ERROR  Problem with the Golang serializer (marshal)
999 INVALID_LOOT_TABLE      Problem with the loot table
```

## Combat System

The combat system is turn-based, inspired by the one found in the role-playing game **Dungeons & Dragons**. It relies on a principle of dice rolls and stat comparison.

### Main mechanics

- the player targets a monster present in the current room;
- an attack roll is simulated by a random number between 1 and 20;
- if the roll is higher than the target's defense, the hit lands;
- a result of 20 triggers a critical hit;
- damage is calculated from the attack of the player or of the equipped weapon.

In the current code, the damage formula is typically of the following form:

- normal damage: a random value around the attack of the weapon or of the character;
- critical damage: a higher amount, with an additional bonus;
- the target can be reduced to 0 HP, which removes it from the arena and triggers loot generation in the room.

### Monster response

The monster counter-attacks immediately after the player's attack if the monster is still alive. If the player belongs to a group, the target of the retaliation can be chosen randomly among the group members present in the room, which adds a small coalition logic and group-attack management.

### Death and recovery

- when a character drops to 0 HP, their state is set to `dead`;
- they are sent back to a respawn room and leave their group;

### Combat management

When a monster appears in a room after being summoned, any player present is able to attack it. Once the monster is dead, the items it had drop to the ground in the room and the players present can use `TAKE` to pick them up.

### Group Combat

Since battles can take place with any other player currently in the room, We’ve chosen to implement group combat so that if one player strikes a monster, one of the other players in the group will take a counterattack. The process is random but allows for group combat and distributes damage among different members. This choice seemed the most logical to us given our combat system, which is open to all players on the server.

### Related commands

The system includes the following commands:

- `ATTACK <monster_id>`
- `ATTACK <monster_id> <weapon_id>`
- `STATUS` to track the character's HP and state
- `SEARCH <monster_id>` to make a monster appear in the room

## Quest System

The quest system is based on two types of quests: item quests and monster quests.

### Quest types

- `QuestItem`: the player must collect a certain number of specific items;
- `QuestMonster`: the player must defeat a given type of monster.

### Progression

Quests follow a status cycle:

- `active`
- `in progress`
- `completed`
- `rewarded`

Progress is updated on every valid interaction, whether an item is added to or removed from the inventory, or a monster is slain. The engine automatically clamps progress between 0 and the requested quantity.

### Validation and reward

When a player talks to a quest-giving NPC, the engine checks whether the quest is completed. If the condition is met:

- the player receives the reward associated with the quest;
- the quest status changes to `rewarded`;
- the required items may be removed from the inventory for item-type quests;
- a quest log is recorded in the server logging system.

Once the quest has been validated, it is no longer possible to obtain it again.

The associated client/server commands are:

- `QUEST <npc_id>` to accept a quest;
- `QUESTS` to display the list of active quests;
- `TALK <npc_id>` to trigger the dialogue and the quest validation.

## World Design

The game world is built from JSON files loaded when the server starts. The world structure is divided into rooms, items, monsters, NPCs and movement links. Adventurers will find themselves, upon connecting, in the central square of the village of **Sangluten-Sur-Pain** and will be able to discover the universe offered to them by exploring the various possible actions.

### General layout

The world includes several distinct zones:

- village,
- forest,
- lake,
- cave,
- casino.

Rooms are linked together through a neighbor system (`north`, `south`, `east`, `west`) and carry position coordinates `(x, y)` to visually structure the world.

### NPC roles

Different types of characters are present:

- `QuestGiver`: gives missions;
- `Trader`: sells and buys items;
- `Dialoguer`: generic dialogue, sometimes related to the story or the setting.

### Distribution of items and monsters

- items are distributed in the rooms or in the traders' inventory;
- monsters are associated with predefined rooms and can be summoned via `SEARCH`;
- the loot dropped after a fight is added to the room so it can be picked up.

### Example world structure

The map notably contains places such as:

- `CASINO`
- `Place Sénile`
- `Taverne de la pucelle`
- `Forêt, épisode 1`
- `Forêt fantôme`
- `Lac Koulofon`
- `Grotte Terrifiante`
- `...`

The world design favors a mix of free exploration and quest-oriented progression, with more hostile zones and safer zones.

## Graphical Interface (GUI)

The GUI was created with **Fyne** and is made up of 2 screens: the **HomeView**, which handles the welcome screen and connections to the game, and the **GameView**, which is the game itself.

### HomeView

This is the screen displayed when the GUI is launched, it contains a **connection form**. This form is made up of:

- A selector between French and English for the display of the GUI's information.

-  A field to enter the nickname the player wants to use, it must be between 1 and 15 characters long and must be composed only of alphanumeric characters and underscores.

- A field for the IP of the server the player wants to join.

When the player clicks on JOIN server, the IP is used to try to connect to the server, if it works, CONNECT <name> <language> is sent to the server. If the server answers "OK connected", the screen changes to display the GameView.

### GameView

This is the game, at startup, it retrieves the global info thanks to SECRET and to the listener that has been set up. It is made up of different widgets that work separately. These widgets are:

- Commands (command_widget.go): creates a bar of category buttons, clicking on a button displays the elements of that category which are the commands available in the game. These commands are also buttons, clicking on them allows, if needed, to choose the parameters and then send it to the server to be executed.

- Log (log_widget.go): the message journal, sorted by category (GLOBAL, ROOM, GROUP, LOG).

- Map (map_widget.go): draws the rooms with colors depending on whether they are the current room, visited, known or unknown. This widget is updated on every LOOK response.

- GoingOn, the current room (goingOn_widget.go): displays the room the player is in: its image is displayed and its name is placed just above.

- Players (players_widget.go): displays the number of players in the room and the total number of players on the server. Requires a first WHO to be displayed and updates the number of players in the room thanks to the event (EVT ROOM PRESENCE ENTER/LEAVE).

- Group (group_widget.go): displays whether the player is in a group or not, if so, the list of its members is also displayed. It refreshes via SECRET when the server emits an EVT GROUP event.

### Synchronization with the server

The Listener is a small pub/sub system: the widgets and the global state subscribe to specific lines. Each server line is sent to all subscribers by Distribute.

The global state is a local copy of the server's SECRET. On every OK SECRET response, setGameData updates the data then triggers the refresh of the widgets that depend on it.

## Server Logging

The logging system is an important component of the project. It is handled in the `src/server/server_write` package.

### Implementation

When the server starts, the program creates the `log/` folder and initializes the `log/server_log.txt` file if necessary. Then, on every event, the `WriteLog` function records a line with the following format:

`timestamp (IP) [LEVEL]: message`

### Event types

The logs cover several categories:

- `INFO`: connection, disconnection, events related to the world state;
- `COMMAND`: commands submitted by players;
- `WARN`: processing errors or warning messages;
- `ERROR`: broken connection or read;
- `SERVER`: responses sent to the client;
- `WORLD`: environmental actions or combat / loot;
- `QUEST`: quest progression;
- `GROUP`: creating, inviting, joining and leaving a group;
- `CHAT`: mutual messages between players.

### Monitoring

To monitor the server's behavior, simply follow the log file in real time:

```bash
tail -f log/server_log.txt
```

This approach makes it possible to quickly detect:

- command abuse,
- repetitive or misleading behavior,
- invalid connection attempts,
- unexpected combat, quest and group events.

## Group Contributions

### Role distribution in our group

- `alebaron`: World management, JSON data parsing and implementation of game commands;
- `rruiz`: GUI integration, World Design and Test & Debug;
- `emarette`: Server setup, implementation of game commands and Happiness Manager.

The distribution is consistent with the repository structure: the main server engine is centralized in the `src/server` packages, while the clients and rendering are spread across the `src/gui` layers and the game components.

## Building and Running

### Prerequisites

- [Go](https://go.dev/dl/) installed (`go version` to check)
- `make`
- `nc` (netcat), only needed for the CLI client

### Build tool: Make

The project uses a `Makefile` that groups together all the build and launch commands.

| Command | Description |
|---|---|
| `make build` | Compiles the project and generates the `main` executable |
| `make run-server` | Compiles then launches the server |
| `make run-gui` | Compiles then launches the graphical client (GUI) |
| `make run-cli IP=<address>` | Launches the command-line client |
| `make clean` | Removes the `main` executable and the `log` folder |

### Compilation

```bash
make build
```

This command runs `go build main.go` and produces the `main` executable at the root of the project.

### Launching the server

```bash
make run-server
```

The server is compiled then started (equivalent to `./main server`). It listens on port **8090**. It must be launched **before** the clients.

### Launching the CLI client

```bash
make run-cli IP=<server_address>
```

The CLI client uses `nc` (netcat) to connect to the server on port 8090. The `IP` parameter is mandatory, for example:

```bash
make run-cli IP=127.0.0.1      # server on the same machine
make run-cli IP=192.168.1.42   # server on another machine on the network
```

### Launching the GUI client

```bash
make run-gui
```

The client is compiled then the graphical interface is launched (equivalent to `./main gui`).

### Cleaning

```bash
make clean
```

Removes the `main` executable as well as the `log` folder generated at runtime.

## Testing

The project does not contain an automated Go test suite in the repository, but the multiplayer system and the game mechanics can be validated manually by running several clients on the same server.

### Multiplayer test

1. Start the server.
2. Open two or more `nc` sessions to port 8090.
3. Connect with different nicknames:

```text
CONNECT alice FR
CONNECT bob EN
```

4. Check that the players appear in the world, that `LOOK` returns the state of the room, and that presence events are properly broadcast.

### Combat test

1. Move to a room containing a monster.
2. Use `SEARCH <monster_id>` to make the monster appear.
3. Run:

```text
ATTACK <monster_id>
```

4. Check:
   - the variation in HP;
   - the loot generation after the monster's death;
   - the logs in `log/server_log.txt`;
   - the server's `OK ...` response.

### Quest test

1. Find a `QuestGiver` in a room.
2. Run:

```text
QUEST <npc_id>
TALK <npc_id>
QUESTS
```

3. Check that the quest goes through the right statuses and that the reward is granted when the condition is met.

### Group test

```text
GROUP CREATE
GROUP INVITE bob
GROUP JOIN <group_id>
GROUP LEAVE
```

Check that group events are sent and that the player's state evolves accordingly.

## Resources

### Teamwork tools

- [Excalidraw](https://excalidraw.com/)
- [Google Sheets](https://www.google.com/sheets/about/)
- [Product backlog overview](https://asana.com/resources/product-backlog)

### Golang references

- [Go Documentation](https://go.dev/doc/)
- [Fyne Documentation](https://docs.fyne.io/)
- [Go pkg documentation for Fyne](https://pkg.go.dev/fyne.io/fyne/v2)
- [MUDs and online text-based worlds](https://fr.wikipedia.org/wiki/Multi-user_dungeon)
- [Tutoriel Golang : apprendre le langage Go (débutant)](https://blog.stephane-robert.info/docs/developper/programmation/golang/)
- [La programmation orientée objet dans le langage de programmation Go](https://devopssec.fr/article/programmation-orientee-objet-golang)
- [Parsing JSON files With Go](https://tutorialedge.net/golang/parsing-json-with-golang/)
- [How to read a file and convert JSON to Go Struct](https://dev.to/mxglt/wip-how-to-read-a-file-and-convert-json-to-go-struct-6m2)

### MUD / RPG references

- [Fer & Flamme](https://fr.wikipedia.org/wiki/Fer_et_Flamme)
- [Donjons & Dragons](https://fr.wikipedia.org/wiki/Donjons_et_Dragons)
- [Wakfu (Jeu vidéo)](https://fr.wikipedia.org/wiki/Wakfu_(jeu_vid%C3%A9o))

### Use of AI

Artificial intelligence was used in a targeted way to:

- help with rewording, structuring and translating the README;
- help fix some bugs encountered.

---

**Last modified**: September 19, 2026\
**Contact:** alebaron@student.42lehavre.fr / rruiz@student.42lehavre.fr / emarette@student.42lehavre.fr

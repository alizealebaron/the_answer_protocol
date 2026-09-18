/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* trad.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/18 13:50:00 by rruiz           #+#    #+#              */
/* Updated: 2026/09/18 13:50:00 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package types

// All UI strings of the client, translated per language.
// Only what is displayed in the GUI goes here, the protocol commands
// sent to the server must stay in English.
var translations = map[string]map[string]string{
	"FR": {
		// Categories of command buttons
		"Environment": "Environnement",
		"Social":      "Social",
		"Fight":       "Combat",
		"Quest":       "Quête",
		"Inventory":   "Inventaire",
		"Gambling":    "Pari",
		"Quit":        "Quitter",

		// Log filter buttons
		"LOG":    "Journal",
		"GLOBAL": "Global",
		"ROOM":   "Salle",
		"GROUP":  "Groupe",

		// Sub-command buttons (labels only, protocol stays English)
		"LOOK":      "Regarder",
		"MOVE":      "Se déplacer",
		"SEARCH":    "Chercher",
		"WHO":       "Qui ?",
		"CHAT":      "Discuter",
		"TALK":      "Parler",
		"ATTACK":    "Attaquer",
		"STATUS":    "Statut",
		"QUEST":     "Quête",
		"QUESTS":    "Quêtes",
		"INVENTORY": "Inventaire",
		"BUY":       "Acheter",
		"DROP":      "Déposer",
		"SELL":      "Vendre",
		"TAKE":      "Prendre",
		"TRADE":     "Échanger",
		"USE":       "Utiliser",
		"GAMBLE":    "Parier",
		"JOIN":      "Rejoindre",
		"Hand":      "À main nue",
		"NORTH":     "Nord",
		"EAST":      "Est",
		"SOUTH":     "Sud",
		"WEST":      "Ouest",

		// Widget labels
		"Not in a group":   "Pas dans un groupe",
		"In group (id=%d)": "Dans le groupe (id=%d)",
		"(you)":            "(vous)",
		"You're lost!":     "Vous êtes perdu !",
		"Use the ‘WHO’ command to view information about players number.": "Utilisez la commande “WHO” pour voir le nombre de joueurs.",
		"Players in room: %d\nPlayers on server: %d":                      "Joueurs dans la salle : %d\nJoueurs sur le serveur : %d",

		// Command placeholders and messages
		"Enter your message.": "Entrez votre message.",
		"Send message":        "Envoyer le message",
		"Enter the quantity you would like to purchase.":      "Entrez la quantité que vous voulez acheter.",
		"Enter the quantity you would like to sell.":          "Entrez la quantité que vous voulez vendre.",
		"Enter the group id.":                                 "Entrez l'identifiant du groupe.",
		"Enter the number of gambling coins you want to bet.": "Entrez le nombre de jetons que vous voulez miser.",
		"The quantity must be an integer.":                    "La quantité doit être un nombre entier.",
		"The quantity must be at least 1.":                    "La quantité doit être d'au moins 1.",
		"The amount must be at least 10.":                     "Le montant doit être d'au moins 10.",
		"You need to have enough money to buy it.":            "Vous devez avoir assez d'argent pour l'acheter.",
		"You must own the item.":                              "Vous devez posséder cet objet.",
		"You don't have that many gambling coins.":            "Vous n'avez pas assez de jetons de pari.",
		"No trader here to buy from.":                         "Aucun commerçant ici pour acheter.",
		"Nothing to drop.":                                    "Rien à jeter.",
		"Nothing to take here.":                               "Rien à prendre ici.",
		"Nothing to use in your inventory.":                   "Rien à utiliser dans votre inventaire.",
		"Nothing to search here.":                             "Rien à chercher ici.",
		"Nothing to sell.":                                    "Rien à vendre.",
		"No one to talk to here.":                             "Personne à qui parler ici.",
		"No quest giver here.":                                "Aucun donneur de quête ici.",
		"No enemy to attack here.":                            "Aucun ennemi à attaquer ici.",
		"No exit from this room.":                             "Aucune sortie de cette salle.",
		"No trader here to trade with.":                       "Aucun commerçant ici pour échanger.",
		"No trader here to sell to.":                          "Aucun commerçant ici pour vendre.",
		"No one to invite to your group.":                     "Personne à inviter dans votre groupe.",
		"You must be in a casino to gamble.":                  "Vous devez être dans un casino pour parier.",
		"You need at least 10 gambling coins to gamble.":      "Vous avez besoin d'au moins 10 jetons de jeu pour parier.",
		"CREATE": "Créer",
		"INVITE": "Inviter",
		"LEAVE":  "Quitter le groupe",
	},
	"EN": {
		"Enter your message.":                        "Enter your message.",
		"Players in room: %d\nPlayers on server: %d": "Players in room: %d\nPlayers on server: %d",
	},
}

// Language selected in the home view, "EN" by default.
var currentLang = "EN"

// SetLang changes the language used by the client UI ("FR" or "EN").
func SetLanguage(lang string) {
	if lang == "FR" || lang == "EN" {
		currentLang = lang
	}
}

// T returns the translation of s in the current language.
// If there is no entry, s itself is returned so protocol strings are kept as they are by default.
func Translate(s string) string {
	if table, ok := translations[currentLang]; ok {
		if translated, ok := table[s]; ok {
			return translated
		}
	}
	return s
}

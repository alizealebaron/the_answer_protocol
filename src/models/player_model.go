/* ************************************************************************ */
/*      _  _     ____                     ,~~.                              */
/*     | || |   |___  \             ,   (  ^ )>                             */
/*     | || |_    __) |             )\~~'   (       _      _      _         */
/*     |__   _|  / __/             (  .__)   )    >(.)__ <(^)__ =(o)__      */
/*        |_|   |_____| .fr         \_.____,*      (___/  (___/  (___/      */
/*                                                                          */
/* ************************************************************************ */
/* name   : player_model.go                                                 */
/* author : alebaron <alebaron@student.42.fr>                               */
/*                                                                          */
/* creation : Invalid date        by -----------                            */
/* update   : 2026/09/11 20:22:18 by alebaron                               */
/* ************************************************************************ */

package models

import (
	"fmt"
	"net"
	"errors"
	"encoding/json"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                          Variable globale                           | */
/* +---------------------------------------------------------------------+ */

var totalPlayer int

/* +---------------------------------------------------------------------+ */
/* |                                Item                                 | */
/* +---------------------------------------------------------------------+ */

type Player struct {
	Id               int
	Name             string
	Pv               int
	MaxPv            int
	Status           string
	Attack           int
	Defense          int
	Language         string
	Money            int
	Inventory        map[Item]int `json:"-"`
	Lst_Quest        []Quest      `json:"-"`
	Conn             net.Conn     `json:"-"`
	Group            *Group       `json:"-"`
	DialogueProgress map[int]int  `json:"-"`
}

/* +---------------------------------------------------------------------+ */
/* |                            Constructeur                             | */
/* +---------------------------------------------------------------------+ */

func NewPlayer(name string, language string, conn net.Conn) Player {

	lstItem := make(map[Item]int)
	dialogueProgress := make(map[int]int)
	lst_quest := []Quest{}
	player := Player{totalPlayer, name, 100, 100, "healthy", 5, 10, language, 10, lstItem, lst_quest, conn, nil, dialogueProgress}
	totalPlayer += 1
	return player
}

/* +---------------------------------------------------------------------+ */
/* |                      Gestion de l'inventaire                        | */
/* +---------------------------------------------------------------------+ */

func (p *Player) GetItem(itID int) (*Item, error) {

	// Parcours des objets de l'inventaire
	for it, _ := range p.Inventory {
		// Gestion des items si on le trouve
		if it.GetId() == itID {
			return &it, nil
		}
	}
	return nil, errors.New("ERR 404 ITEM_NOT_FOUND")
}

func (p Player) GetQuantityItem(itID int) (int, error) {

	// Parcours des objets de l'inventaire
	for it, qty := range p.Inventory {
		// Gestion des items si on le trouve
		if it.GetId() == itID {
			return qty, nil
		}
	}

	return 0, errors.New("ERR 404 ITEM_NOT_FOUND")
}

func (p *Player) AddItemToPlayerWQuantity(it Item, q int) {

	// Ajout à l'inventaire
	p.Inventory[it] += q

	// Mise à jour des potentielles quêtes
	p.UpdateQuestItem(it.GetId(), q)
}

func (p *Player) RemoveItemToPlayerWQuantity(itID int, q int) (*Item, error) {

	// Parcours des objets de l'inventaire
	for it, qty := range p.Inventory {

		// Gestion des items si on le trouve
		if it.GetId() == itID {
			if qty == q {
				delete(p.Inventory, it)
			} else if qty < q {
				return nil, errors.New("ERR 420 NOT_ENOUGH_ITEM")
			} else {
				p.Inventory[it] = qty - q
			}
			// Mise à jour des potentielles quêtes
			p.UpdateQuestItemBrut(it.GetId(), (qty - q))
			itemCopy := it
			return &itemCopy, nil
		}
	}
	return nil, errors.New("ERR 404 ITEM_NOT_FOUND")
}

func (p *Player) InventoryToString() string {
	type ItemEntry struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Quantity  int    `json:"quantity"`
		Is_Usable bool   `json:"is_usable"`
		Is_Weapon bool   `json:"is_weapon"`
	}

	items := make([]ItemEntry, 0, len(p.Inventory))
	for it, qty := range p.Inventory {

		// Vérification de si l'item est utilisable
		_, ok_ed := it.(Edible)
		_, ok_us := it.(Usable)
		is_usable := ok_ed || ok_us
		_, is_weapon := it.(Weapon)

		// Ajout de l'item à l'inventaire
		items = append(items, ItemEntry{
			Id:        it.GetId(),
			Name:      it.GetName(),
			Quantity:  qty,
			Is_Usable: is_usable,
			Is_Weapon: is_weapon,
		})
	}

	out := struct {
		Items []ItemEntry `json:"items"`
		Money int         `json:"money"`
	}{
		Items: items,
		Money: p.Money,
	}

	b, err := json.Marshal(out)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}

/* +---------------------------------------------------------------------+ */
/* |                     Gestion de la vie du joueur                     | */
/* +---------------------------------------------------------------------+ */

func (p *Player) PlayerDeath(tapManager *TapManager) error {
	p.Pv = 30
	p.Status = "bloody"
	group, err := tapManager.GetGroupById(p.Id)
	if err == nil {
		group.RemovePlayerFromGroup(*p)
	}
	p_room, err := tapManager.FindPlayerRoom(p.Id)
	if err != nil {
		return err
	}
	p_room.RemovePlayerToRoom(*p)
	p_room, err = tapManager.GetRoomById(5)
	if err != nil {
		return err
	}
	p_room.AddPlayerToRoom(*p)
	return nil
}

func (p *Player) AddLifePoint(value int) error {
	p.Pv += value
	if p.Pv > 50 {
		p.Status = "healthy"
	} else if p.Pv <= 50 && p.Pv > 0 {
		p.Status = "bloody"
	} else {
		p.Status = "dead"
	}
	if p.Pv > 100 {
		p.Pv = 100
	}
	return nil
}

/* +---------------------------------------------------------------------+ */
/* |                         Gestion des quêtes                          | */
/* +---------------------------------------------------------------------+ */

func (p *Player) AddQuestToPlayer(quest Quest) error {

	// Vérification que le joueur n'a pas déjà la quête
	for _, q := range p.Lst_Quest {
		if q.GetId() == quest.GetId() {
			return errors.New("ERR 406 NO_QUEST_AVAILABLE")
		}
	}

	// Mise à jour du status de la quête
	quest.SetStatus("active")

	// Si la quête est une quête à Item vérifié le nombre d'item déjà présent dans l'inventaire
	questItem, ok := quest.(*QuestItem)
	if ok {
		qty, _ := p.GetQuantityItem(questItem.ItemNeededId)
		questItem.UpdateProgress(qty)
	}

	// Ajout de la quête à la liste du joueur
	p.Lst_Quest = append(p.Lst_Quest, quest)

	server_write.WriteLog(p.Conn, "QUEST", p.Name + " started quest \"" + quest.GetTitle() + "\"\n")

	return nil
}

func (p *Player) UpdateQuestMonster(monster Monster) error {

	for _, quest := range p.Lst_Quest {
		questMonster, ok := quest.(*QuestMonster)
		if ok {
			if questMonster.MonsterNeededId == monster.Id {
				questMonster.UpdateProgress(1)
			}
		}
	}

	return nil
}

func (p *Player) UpdateQuestItem(id_item int, quantity int) error {

	for _, quest := range p.Lst_Quest {
		questItem, ok := quest.(*QuestItem)
		if ok {
			if questItem.ItemNeededId == id_item {
				questItem.UpdateProgress(quantity)
			}
		}
	}

	return nil
}

func (p *Player) UpdateQuestItemBrut(id_item int, quantity int) error {

	for _, quest := range p.Lst_Quest {
		questItem, ok := quest.(*QuestItem)
		if ok {
			if questItem.ItemNeededId == id_item {
				questItem.UpdateProgressBrut(quantity)
			}
		}
	}

	return nil
}

func (p *Player) IsNpcQuestCompleted(npc QuestGiver) bool {

	for _, q := range p.Lst_Quest {
		if q.GetId() == npc.QuestId && q.GetStatus() == "completed" {

			// Ajout des récompenses à maj du statut
			p.AddItemToPlayerWQuantity(q.GetReward(), q.GetQuantity())
			q.SetStatus("rewarded")

			// On retire les items demandés si c'est une quête d'item
			itemquest, ok := q.(*QuestItem)
			if ok {
				p.RemoveItemToPlayerWQuantity(itemquest.ItemNeededId, itemquest.SearchQuantity)
			}

			server_write.WriteLog(p.Conn, "QUEST", p.Name + " finished quest \"" + q.GetTitle() + "\"\n")

			return true
		}
	}

	return false
}

/* +---------------------------------------------------------------------+ */
/* |                        Gestion du gambling                          | */
/* +---------------------------------------------------------------------+ */

func (p *Player) IsEnoughGamblingCoin(bet int) (bool, error) {

	// Parcours des objets de l'inventaire
	for it := range p.Inventory {
		// Gestion des items si on le trouve
		if it.GetId() == 1 && p.Inventory[it] >= bet {
			return true, nil
		}
	}
	return false, errors.New("ERR 999 NOT_ENOUGH_COIN")
}

/* +---------------------------------------------------------------------+ */
/* |                       Gestion des dialogues                         | */
/* +---------------------------------------------------------------------+ */

func (p *Player) GetNextDialogueLine(npc Npc) (line string) {

	// Récupération de l'ID
	id := npc.GetId()

	// Récupération des dialogues français ou anglais
	lines := npc.GetDialogueFr()
	if p.Language == "EN" {
		lines = npc.GetDialogueEn()
	}

	// Récupération du dialogue à renvoyé
	idx := p.DialogueProgress[id]

	if idx >= len(lines) {
		p.DialogueProgress[id] = 0
		idx = 0
	}

	p.DialogueProgress[id] = idx + 1
	return lines[idx]
}

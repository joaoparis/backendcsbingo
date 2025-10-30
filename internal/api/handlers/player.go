package handlers

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

type Player struct {
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Team        string `json:"team"`
}

var players = []Player{
	{"NiKo", "Bosnia & Herzegovina", "G2"},
	{"s1mple", "Ukraine", "NAVI"},
	{"ZywOo", "France", "Vitality"},
	{"dev1ce", "Denmark", "Astralis"},
	{"m0NESY", "Russia", "G2"},
	{"donk", "Russia", "Spirit"},
	{"ropz", "Estonia", "FaZe"},
	{"blameF", "Denmark", "Astralis"},
	{"karrigan", "Germany", "FaZe"},
	{"electronic", "Russia", "Cloud9"},
	{"cadiaN", "Denmark", "Liquid"},
	{"twistzz", "Canada", "FaZe"},
	{"rain", "Norway", "FaZe"},
	{"Snax", "Poland", "Virtus.pro"},
	{"GeT_RiGhT", "Sweden", "NiP"},
	{"f0rest", "Sweden", "NiP"},
	{"coldzera", "Brazil", "SK Gaming"},
	{"fer", "Brazil", "MIBR"},
	{"HooXi", "Denmark", "G2"},
	{"huNter-", "Bosnia & Herzegovina", "G2"},
}

func GetRandomPlayer(c *fiber.Ctx) error {
	randomIndex := rand.Intn(len(players))
	return c.JSON(players[randomIndex])
}

package handlers

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

var options = []string{
	"Bosnia & Herzegovina",
	"Ukraine",
	"NAVI",
	"France",
	"Vitality",
	"Denmark",
	"Astralis",
	"G2",
	"Russia",
	"Spirit",
	"Estonia",
	"FaZe",
	"Germany",
	"Cloud9",
	"Liquid",
	"Canada",
	"Norway",
	"Poland",
	"Virtus.pro",
	"Sweden",
	"NiP",
	"Brazil",
	"SK Gaming",
	"MIBR",
}

func GetRandomBoard(c *fiber.Ctx) error {
	rand.Seed(time.Now().UnixNano())

	shuffled := make([]string, len(options))
	copy(shuffled, options)

	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	board := shuffled[:16]

	return c.JSON(fiber.Map{
		"board": board,
	})
}

package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Dum struct {
	Yahaha  string `json:"yahaha" form:"yahaha"`
	Hayyukk string `json:"hayyukk" form:"hayyukk"`
}

func NyobakAPI(c *fiber.Ctx) error {
	d := new(Dum)
	if err := c.BodyParser(d); err != nil {
		return err
	}
	log.Println(d.Yahaha)
	return nil
}

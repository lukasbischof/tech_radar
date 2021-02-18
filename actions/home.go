package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
	"github.com/lukasbischof/tech_radar/models"
	"net/http"
)

func HomeHandler(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	technologies := &models.Technologies{}

	if err := tx.All(technologies); err != nil {
		return err
	}

	grouped := technologies.GroupTechnologies()
	fmt.Println(technologies)
	fmt.Println(grouped)

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("techGroups", grouped)
		return c.Render(http.StatusOK, r.HTML("/index.plush.html"))
	}).Respond(c)
}
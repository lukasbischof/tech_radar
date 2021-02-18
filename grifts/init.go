package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/lukasbischof/tech_radar/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}

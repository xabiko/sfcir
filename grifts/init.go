package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/xabiko/sfcir/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}

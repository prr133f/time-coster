package v1

import (
	"time-coster/view"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Group *fiber.Group
	View  *view.View
}

type Router struct {
	App  *fiber.App
	View *view.View
}

func (r *Router) Router() {

}

package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/mimani68/fintech-core/transport/http/echo/handlers"
)

func (h *handlers.Handler) RegisterRouters(urlPrefix *echo.Group) {

	forms := urlPrefix.Group("/forms")
	forms.GET("", h.ListOfForms)
	forms.POST("", h.SubmitNewForm)
}

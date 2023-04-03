package echo

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRouters(urlPrefix *echo.Group) {

	forms := urlPrefix.Group("/forms")
	forms.GET("", h.ListOfForms)
	forms.POST("", h.SubmitNewForm)
}

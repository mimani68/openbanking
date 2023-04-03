package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// SubmitNewForm godoc
// @Summary Submit new contact forms
// @Description Register new form here
// @ID submit-new-form
// @Tags form
// @Accept  json
// @Produce  json
// @Param form body SubmitFormRequest true "Form info for registration"
// @Success 201 {object} NewFormResponse
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /forms [post]
func (h *Handler) SubmitNewForm(c echo.Context) error {
	u := new(model.Form)
	req := &dto.SubmitFormRequest{}
	if err := req.Bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	_, err := h.form.CreateForm(*u)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, dto.NewFormResponse(true, *u))
}

// ListOfForms godoc
// @Summary Get all forms here
// @Description
// @ID list-form
// @Tags form
// @Accept  json
// @Produce  json
// @Success 200 {object} GetAllFormsResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /form [get]
func (h *Handler) ListOfForms(c echo.Context) error {
	u, err := h.form.GetFormList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, dto.GetAllFormsResponse(true, u))
}

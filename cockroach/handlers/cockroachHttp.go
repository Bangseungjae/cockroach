package handlers

import (
	"Bangseungjae/cockroach/cockroach/models"
	"Bangseungjae/cockroach/cockroach/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type cockroachHttpHandler struct {
	cockroachUsecase usecases.CockroachUsecase
}

func NewCockroachHttpHandler(cockroachUsecase usecases.CockroachUsecase) *cockroachHttpHandler {
	return &cockroachHttpHandler{cockroachUsecase: cockroachUsecase}
}

func (h *cockroachHttpHandler) DetectCockroach(c echo.Context) error {
	reqBody := new(models.AddCockroachData)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request")
	}

	if err := h.cockroachUsecase.CockroachDataProcessing(reqBody); err != nil {
		return response(c, http.StatusInternalServerError, "Processing data failed")
	}
	return response(c, http.StatusOK, "Success 🪳🪳🪳")
}

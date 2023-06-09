package handlers

import (
	"BtcApp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetExchangeRate godoc
// @Summary Get BTC to UAH exchange rate
// @Description Returns the current BTC to UAH exchange rate
// @Tags rate
// @Accept json
// @Produce json
// @Success 200 {number} decimal
// @Router /exchange-rate [get]
func GetExchangeRate(c *gin.Context) {

	exchangeRate, err := services.FetchExchangeRate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch exchange rate."})
		return
	}

	c.JSON(http.StatusOK, exchangeRate.Value)
}

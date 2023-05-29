package handler

import (
	"ardi_go/campaign"
	"ardi_go/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewcampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaign, err := h.service.GetCampaign(userID)
	if err != nil {
		respon := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	response := helper.APIResponse("List get campaign", http.StatusOK, "Success", campaign)
	c.JSON(http.StatusOK, response)
}

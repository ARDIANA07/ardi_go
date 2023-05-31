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

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaign(userID)
	if err != nil {
		respon := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "error", campaign.FormatCampaigns(campaigns))
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	response := helper.APIResponse("List get campaign", http.StatusOK, "Success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		respon := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "Error", err.Error())
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	campaignDetail, err := h.service.GetCampaignById(input)
	if err != nil {
		respon := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	response := helper.APIResponse("Campaign detail", http.StatusOK, "Success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

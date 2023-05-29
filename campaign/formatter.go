package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignformater := CampaignFormatter{}
	campaignformater.ID = campaign.ID
	campaignformater.Name = campaign.Name
	campaignformater.ShortDescription = campaign.ShortDescription
	campaignformater.GoalAmount = campaign.GoalAmount
	campaignformater.CurrentAmount = campaign.CurrentAmount
	campaignformater.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignformater.ImageURL = campaign.CampaignImages[0].FileName
	}
	return campaignformater
}

func FormatCampaigns(campaign []Campaign) []CampaignFormatter {
	campaignformater := []CampaignFormatter{}
	for _, campaign := range campaign {
		campaignformater = append(campaignformater, FormatCampaign(campaign))
	}
	return campaignformater
}

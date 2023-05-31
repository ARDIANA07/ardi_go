package campaign

import (
	"ardi_go/user"
	"time"
)

type Campaign struct {
	ID               int       `gorm:"id" type:"int(11)"`
	UserID           int       `gorm:"user_id"`
	Name             string    `gorm:"name" type:"varchar(255)"`
	ShortDescription string    `gorm:"short_description" type:"varchar(255)"`
	Description      string    `gorm:"description" type:"text"`
	Perks            string    `gorm:"perks" type:"text"`
	BackerCount      int       `gorm:"backer_count" type:"int(50)"`
	GoalAmount       int       `gorm:"goal_amount" type:"int(50)"`
	CurrentAmount    int       `gorm:"current_amount" type:"int(50)"`
	Slug             string    `gorm:"slug" type:"varchar(255)"`
	CreatedAt        time.Time `gorm:"created_at"`
	UpdatedAt        time.Time `gorm:"updated_at"`
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	ID         int       `json:"id"`
	CampaignID int       `json:"campaign_id"`
	FileName   string    `json:"file_name" type:"varchar(255)"`
	IsPrimary  int       `json:"is_primary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"update_at"`
}

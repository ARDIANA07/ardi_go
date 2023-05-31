package campaign

type Service interface {
	GetCampaign(userID int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaign(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserId(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaingn, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaingn, err
	}
	return campaingn, nil
}

package requests

type CreateInvitationRequest struct {
	IdTemplate      string `json:"id_template" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Date            string `json:"date" validate:"required"`
	Time            string `json:"time" validate:"required"`
	Location        string `json:"location" validate:"required"`
	Description     string `json:"description" validate:"required"`
	PrimaryColor    string `json:"primary_color" validate:"required"`
	SecondaryColor  string `json:"secondary_color" validate:"required"`
	BackgroundImage string `json:"background_image"`
}

type GenerateLinkRequest struct {
	IdInvitation string `json:"id_invitation" validate:"required"`
}

type ShareSocialMediaRequest struct {
	IdInvitation string `json:"id_invitation" validate:"required"`
	NamePlatform string `json:"name_platform" validate:"required"`
}

type GuestViewRequest struct {
	IdInvitationLink string `json:"id_invitation_link" validate:"required"`
	IpAddres         string `json:"ip_address" validate:"required"`
	UserAgent        string `json:"user_agent" validate:"required"`
}

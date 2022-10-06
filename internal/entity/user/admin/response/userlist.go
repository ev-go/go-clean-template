package response

type UserListRes struct {
	Items []UserListItem `json:"items" `
	Total int            `json:"total" example:"1"`
}

type UserListItem struct {
	UserId   string `json:"customerId" example:"0da3b22f-ec3f-4383-bc25-480b6dcb82a1"`
	Name     string `json:"name" example:"OOO Galileosky"`
	Inn      string `json:"inn" example:"5904254657"`
	FullName string `json:"fullName" example:"Limited liability company Galileosky"`
	Country  string `json:"country" example:"Russia"`
	Region   string `json:"region" example:"Perm region"`
	Contacts string `json:"contacts" example:"8 495 001 3930"`
	DopInfo  string `json:"dopInfo" example:"vendor"`
	Enabled  bool   `json:"enabled" example:"false"`
	//ApiKey     string `json:"apiKey" example:"cSPZV2BtniuCyGynsww7PY.LsMsj3TEnYMbezinoA6NsL"`
}

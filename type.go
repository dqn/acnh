package acnh

type Land struct {
	DisplayID int    `json:"displayId"`
	ID        string `json:"id"`
	Name      string `json:"name"`
}

type Users struct {
	ID    string `json:"id"`
	Image string `json:"image"`
	Land  Land   `json:"land"`
	Name  string `json:"name"`
}

type UsersResponse struct {
	Users []Users `json:"users"`
}

type AuthTokenRequest struct {
	UserID string `json:"userId"`
}

type AuthTokenResponse struct {
	ExpireAt int    `json:"expireAt"`
	Token    string `json:"token"`
}

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

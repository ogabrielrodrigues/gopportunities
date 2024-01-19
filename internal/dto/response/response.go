package response

type OpeningResponse struct {
	ID          string `json:"id"`
	Role        string `json:"role"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Remote      bool   `json:"remote"`
	Link        string `json:"link"`
	Salary      uint64 `json:"salary"`
}

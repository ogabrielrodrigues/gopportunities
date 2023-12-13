package request

type OpeningCreateRequest struct {
	Role        string `json:"role" validate:"required,min=4,max=100"`
	Description string `json:"description" validate:"required,min=10,max=400"`
	Company     string `json:"company" validate:"required,min=3,max=50"`
	Location    string `json:"location" validate:"required,min=8,max=150"`
	Remote      bool   `json:"remote"`
	Link        string `json:"link" validate:"required,url"`
	Salary      uint64 `json:"salary" validate:"required,min=1"`
}

type OpeningUpdateRequest struct {
	Role        string `json:"role"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Remote      bool   `json:"remote"`
	Link        string `json:"link"`
	Salary      uint64 `json:"salary"`
}

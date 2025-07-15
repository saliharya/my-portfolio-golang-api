package domain

type Services struct {
	Summary string        `json:"summary"`
	Items   []ServiceItem `json:"items"`
}

type ServiceItem struct {
	Id          int    `json:"id"`
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkUrl     string `json:"link_url,omitempty"`
}

type ServicesRepository interface {
	GetServices() (*Services, error)
	CreateService(service ServiceItem) (*ServiceItem, error)
	UpdateService(service ServiceItem) (*ServiceItem, error)
	DeleteService(id int) error
	UpdateSummary(summary string) (*Services, error)
}

type ServicesUsecase interface {
	GetServices() (*Services, error)
	CreateService(service ServiceItem) (*ServiceItem, error)
	UpdateService(service ServiceItem) (*ServiceItem, error)
	DeleteService(id int) error
	UpdateSummary(summary string) (*Services, error)
}

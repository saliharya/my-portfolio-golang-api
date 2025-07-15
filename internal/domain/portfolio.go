package domain

type Portfolio struct {
	Summary  string    `json:"summary"`
	Projects []Project `json:"projects"`
}

type Project struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type PortfolioRepository interface {
	GetPortfolio() (*Portfolio, error)
}

type PortfolioUsecase interface {
	GetPortfolio() (*Portfolio, error)
}

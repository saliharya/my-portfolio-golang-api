package domain

type About struct {
	Description string   `json:"description"`
	Summary     string   `json:"summary"`
	Photo       string   `json:"photo"`
	Languages   []string `json:"languages"`
	Education   []string `json:"education"`
	Projects    []string `json:"projects"`
	Tools       []Tool   `json:"tools"`
}

type Tool struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type AboutRepository interface {
	GetAbout() (*About, error)
}

type AboutUsecase interface {
	GetAbout() (*About, error)
	UpdateAbout(about About) (*About, error)
}

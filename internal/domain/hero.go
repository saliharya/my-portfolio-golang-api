package domain

type Hero struct {
	Name        string `json:"name"`
	Photo       string `json:"photo"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type HeroRepository interface {
	GetHero() (*Hero, error)
}

type HeroUsecase interface {
	GetHero() (*Hero, error)
}

package repository

import "portfolio-arya-service/internal/domain"

type heroRepository struct{}

func NewHeroRepository() domain.HeroRepository {
	return &heroRepository{}
}

func (r *heroRepository) GetHero() (*domain.Hero, error) {
	hero := &domain.Hero{
		Name:        "Salih Arya Gumilang",
		Photo:       "/assets/images/hero-photo.jpeg",
		Title:       "/assets/images/hero-photo.jpeg",
		Description: "Experienced in mobile and web development, with strong skills in cross-functional team collaboration. Informatics \nEngineering graduate from Ahmad Dahlan University. Graduated with distinction from Bangkit Academy, a Kampus Merdeka \nprogram led by Google, Tokopedia, Gojek, and Traveloka.",
	}

	return hero, nil
}

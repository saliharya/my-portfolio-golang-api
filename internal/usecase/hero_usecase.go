package usecase

import "portfolio-arya-service/internal/domain"

type heroUsecase struct {
	heroRepository domain.HeroRepository
}

func NewHeroUsecase(repository domain.HeroRepository) domain.HeroUsecase {
	return &heroUsecase{heroRepository: repository}
}

func (r *heroUsecase) GetHero() (*domain.Hero, error) {
	return r.heroRepository.GetHero()
}

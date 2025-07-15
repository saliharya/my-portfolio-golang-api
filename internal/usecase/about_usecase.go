package usecase

import "portfolio-arya-service/internal/domain"

type aboutUsecase struct {
	aboutRepository domain.AboutRepository
}

func (r *aboutUsecase) UpdateAbout(about domain.About) (*domain.About, error) {
	return r.aboutRepository.UpdateAbout(about)
}

func NewAboutUsecase(repository domain.AboutRepository) domain.AboutUsecase {
	return &aboutUsecase{aboutRepository: repository}
}

func (r *aboutUsecase) GetAbout() (*domain.About, error) {
	return r.aboutRepository.GetAbout()
}

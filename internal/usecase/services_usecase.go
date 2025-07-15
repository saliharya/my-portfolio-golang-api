package usecase

import "portfolio-arya-service/internal/domain"

type servicesUsecase struct {
	servicesRepository domain.ServicesRepository
}

func NewServicesUsecase(repository domain.ServicesRepository) domain.ServicesUsecase {
	return &servicesUsecase{servicesRepository: repository}
}

func (r *servicesUsecase) GetServices() (*domain.Services, error) {
	return r.servicesRepository.GetServices()
}

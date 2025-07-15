package usecase

import "portfolio-arya-service/internal/domain"

type servicesUsecase struct {
	servicesRepository domain.ServicesRepository
}

func NewServicesUsecase(repository domain.ServicesRepository) domain.ServicesUsecase {
	return &servicesUsecase{servicesRepository: repository}
}

func (r *servicesUsecase) CreateService(service domain.ServiceItem) (*domain.ServiceItem, error) {
	return r.servicesRepository.CreateService(service)
}

func (r *servicesUsecase) UpdateService(service domain.ServiceItem) (*domain.ServiceItem, error) {
	return r.servicesRepository.UpdateService(service)
}

func (r *servicesUsecase) DeleteService(id int) error {
	return r.servicesRepository.DeleteService(id)
}

func (r *servicesUsecase) UpdateSummary(summary string) (*domain.Services, error) {
	return r.servicesRepository.UpdateSummary(summary)
}

func (r *servicesUsecase) GetServices() (*domain.Services, error) {
	return r.servicesRepository.GetServices()
}

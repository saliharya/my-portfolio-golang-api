package repository

import (
	"errors"
	"portfolio-arya-service/internal/domain"
)

type servicesRepository struct{}

var servicesData = &domain.Services{
	Summary: "Summary Services",
	Items: []domain.ServiceItem{
		{
			Id:          1,
			Icon:        "assets/images/web-icon.png",
			Title:       "Web Development",
			Description: "Building responsive websites",
			LinkUrl:     "https://example.com/project1",
		},
	},
}
var nextServiceId = 2

func (r *servicesRepository) GetServices() (*domain.Services, error) {
	return servicesData, nil
}

func (r *servicesRepository) CreateService(service domain.ServiceItem) (*domain.ServiceItem, error) {
	service.Id = nextServiceId
	nextServiceId++
	servicesData.Items = append(servicesData.Items, service)
	return &service, nil
}

func (r *servicesRepository) UpdateService(service domain.ServiceItem) (*domain.ServiceItem, error) {
	for i, item := range servicesData.Items {
		if item.Id == service.Id {
			servicesData.Items[i] = service
			return &service, nil
		}
	}
	return nil, errors.New("service not found")
}

func (r *servicesRepository) DeleteService(id int) error {
	for i, item := range servicesData.Items {
		if item.Id == id {
			servicesData.Items = append(servicesData.Items[:i], servicesData.Items[i+1:]...)
			return nil
		}
	}
	return errors.New("service not found")
}

func (r *servicesRepository) UpdateSummary(summary string) (*domain.Services, error) {
	servicesData.Summary = summary
	return servicesData, nil
}

func NewServicesRepository() domain.ServicesRepository {
	return &servicesRepository{}
}

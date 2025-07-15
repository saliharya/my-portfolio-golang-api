package repository

import "portfolio-arya-service/internal/domain"

type servicesRepository struct {
}

func NewServicesRepository() domain.ServicesRepository {
	return &servicesRepository{}
}

func (r *servicesRepository) GetServices() (*domain.Services, error) {
	services := &domain.Services{
		Summary: "Summary Services",
		Items: []domain.ServiceItem{
			{
				Id:          1,
				Icon:        "assets/images/web-icon.png",
				Title:       "Web Development",
				Description: "Building responsive websites",
				LinkUrl:     "https://example.com/project1",
			},
			{
				Id:          1,
				Icon:        "assets/images/web-icon.png",
				Title:       "Web Development",
				Description: "Building responsive websites",
				LinkUrl:     "https://example.com/project1",
			},
			{
				Id:          1,
				Icon:        "assets/images/web-icon.png",
				Title:       "Web Development",
				Description: "Building responsive websites",
				LinkUrl:     "https://example.com/project1",
			},
			{
				Id:          1,
				Icon:        "assets/images/web-icon.png",
				Title:       "Web Development",
				Description: "Building responsive websites",
				LinkUrl:     "https://example.com/project1",
			},
		},
	}
	return services, nil
}

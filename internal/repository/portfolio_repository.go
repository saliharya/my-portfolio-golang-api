package repository

import "portfolio-arya-service/internal/domain"

type portfolioRepository struct {
}

func NewPortfolioRepository() domain.PortfolioRepository {
	return &portfolioRepository{}
}

func (r *portfolioRepository) GetPortfolio() (*domain.Portfolio, error) {
	portfolio := &domain.Portfolio{
		Summary: "Summary Portfolio",
		Projects: []domain.Project{
			{
				Id:       "1",
				Name:     "Project A",
				Category: "Web Development",
			},
			{
				Id:       "2",
				Name:     "Project B",
				Category: "Mobile App",
			},
		},
	}
	return portfolio, nil
}

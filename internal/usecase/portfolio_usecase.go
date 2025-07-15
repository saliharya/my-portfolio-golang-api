package usecase

import "portfolio-arya-service/internal/domain"

type portfolioUsecase struct {
	portfolioRepository domain.PortfolioRepository
}

func NewPortfolioUsecase(repository domain.PortfolioRepository) domain.PortfolioUsecase {
	return &portfolioUsecase{portfolioRepository: repository}
}

func (r *portfolioUsecase) GetPortfolio() (*domain.Portfolio, error) {
	return r.portfolioRepository.GetPortfolio()
}

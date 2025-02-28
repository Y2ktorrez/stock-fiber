package services

import (
	"github.com/mutinho/src/dtos"
	"github.com/mutinho/src/models"
	"github.com/mutinho/src/repository"
)

type MovimentService struct {
	movimentRepo *repository.MovimentRepository
}

func NewMovimentService(movimentRepo *repository.MovimentRepository) *MovimentService {
	return &MovimentService{movimentRepo: movimentRepo}
}

func (s *MovimentService) CreateMoviment(req dtos.CreateMovimentRequest) (*dtos.MovimentResponse, error) {
	product, err := s.movimentRepo.FindProductBySlug(req.ProductSlug)
	if err != nil {
		return nil, err
	}

	user, err := s.movimentRepo.FindUserBySlug(req.UserSlug)
	if err != nil {
		return nil, err
	}

	moviment := &models.Moviment{
		Type:       models.TipoMovimiento(req.Type),
		Amount:     req.Amount,
		IDProducts: product.ID,
		IDUser:     user.ID,
	}

	if err := s.movimentRepo.Create(moviment); err != nil {
		return nil, err
	}

	return &dtos.MovimentResponse{
		ID:          moviment.ID.String(),
		Type:        string(moviment.Type),
		Amount:      moviment.Amount,
		Date:        moviment.Date.Format("2006-01-02 15:04:05"),
		ProductSlug: product.Slug,
		UserSlug:    user.Slug,
	}, nil
}

func (s *MovimentService) GetAllMoviments() ([]dtos.MovimentResponse, error) {
	moviments, err := s.movimentRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var movimentResponses []dtos.MovimentResponse
	for _, moviment := range moviments {
		movimentResponses = append(movimentResponses, dtos.MovimentResponse{
			ID:          moviment.ID.String(),
			Type:        string(moviment.Type),
			Amount:      moviment.Amount,
			Date:        moviment.Date.Format("2006-01-02 15:04:05"),
			ProductSlug: moviment.Producto.Slug,
			UserSlug:    moviment.Usuario.Slug,
		})
	}

	return movimentResponses, nil
}

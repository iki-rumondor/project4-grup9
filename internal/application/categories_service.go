package application

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
)

type CategoriesService struct {
	Repo repository.CategoriesRepository
}

func NewCategoriesService(repo repository.CategoriesRepository) *CategoriesService {
	return &CategoriesService{
		Repo: repo,
	}
}

func (s *CategoriesService) CreateCategories(categories *domain.Categories) (*domain.Categories, error) {
	result, err := s.Repo.CreateCategories(categories)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *CategoriesService) GetCategories() (*[]domain.Categories, error) {
	categorie, err := s.Repo.FindCategories()
	if err != nil {
		return nil, err
	}

	return categorie, nil
}

func (s *CategoriesService) UpdateCategories(categories *domain.Categories) (*domain.Categories, error) {
	// _, err := s.Repo.FindCategories(categories.ID)

	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return nil, fmt.Errorf("category with id %d id not found", categories.ID)
	// }
	// if err != nil {
	// 	return nil, errors.New("failed to get category from database")
	// }

	categories, err := s.Repo.UpdateCategories(categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoriesService) DeleteCategories(categories *domain.Categories) error {
	// _, err := s.Repo.FindCategories(categories.ID)

	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return fmt.Errorf("category with id %d id not found", categories.ID)
	// }
	// if err != nil {
	// 	return errors.New("failed to get category from database")
	// }

	if err := s.Repo.DeleteCategories(categories); err != nil {
		return err
	}

	return nil
}

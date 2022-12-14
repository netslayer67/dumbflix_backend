package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindCategory() ([]models.Category, error)
	GetCategory(ID int) (models.Category, error)
	GetCategoryByName(name string) (*models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	DeleteCategory(category models.Category) (models.Category, error)
}

type repositoryForCategory struct {
	db *gorm.DB
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCategory() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Preload("Films").Find(&categories).Error

	return categories, err
}

func (r *repository) GetCategory(ID int) (models.Category, error) {
	var category models.Category
	err := r.db.First(&category, ID).Error

	return category, err
}

func (r *repository) GetCategoryByName(name string) (*models.Category, error) {
	category := &models.Category{}
	err := r.db.Where("name = ?", name).First(&category).Error

	return category, err
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error

	return category, err
}

func (r *repository) UpdateCategory(category models.Category) (models.Category, error) {
	err := r.db.Save(&category).Error

	return category, err
}

func (r *repository) DeleteCategory(category models.Category) (models.Category, error) {
	err := r.db.Delete(&category).Error // Using Delete method

	return category, err
}

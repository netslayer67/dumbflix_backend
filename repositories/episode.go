package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type EpisodeRepository interface {
	FindEpisode() ([]models.Episode, error)
	GetEpisode(ID int) (models.Episode, error)
	CreateEpisode(episode models.Episode) (models.Episode, error)
	UpdateEpisode(episode models.Episode) (models.Episode, error)
	DeleteEpisode(episode models.Episode) (models.Episode, error)
}

type repositoryForEpisode struct {
	db *gorm.DB
}

func RepositoryEpisode(db *gorm.DB) *repositoryForEpisode {
	return &repositoryForEpisode{db}
}

func (r *repositoryForEpisode) FindEpisode() ([]models.Episode, error) {
	var episodes []models.Episode
	err := r.db.Preload("Film").Preload("Film.Category").Find(&episodes).Error

	return episodes, err
}

func (r *repositoryForEpisode) GetEpisode(ID int) (models.Episode, error) {
	var episode models.Episode
	err := r.db.Preload("Film").Preload("Film.Category").First(&episode, ID).Error

	return episode, err
}

func (r *repositoryForEpisode) CreateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Create(&episode).Error

	return episode, err
}

func (r *repositoryForEpisode) UpdateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Save(&episode).Error

	return episode, err
}

func (r *repositoryForEpisode) DeleteEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Delete(&episode).Error // Using Delete method

	return episode, err
}

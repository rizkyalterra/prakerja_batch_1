package repo

import (
	"prakerja/config"
	"prakerja/models"
	"prakerja/utils"

	"gorm.io/gorm"
)

func InsertNews(news *models.News) error {
	result := config.DB.Create(news)
	if result.Error != nil {
		return utils.ERR_CREATE_DATABASE
	}
	return nil
}

func GetNews() ([]models.News, error) {
	var data []models.News
	result := config.DB.Find(&data)

	if result.Error == gorm.ErrRecordNotFound {
		return []models.News{}, utils.ERR_EMPTY_DATA_IN_DATABASE
	} else if result.Error != nil {
		return nil, utils.ERR_GET_DATA_IN_DATABASE
	}

	return data, nil
}

package models

type Category struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Films []Film `json:"films"`
}

type CategoryRelation struct {
	Name string `json:"name"`
}

func (CategoryRelation) TableName() string {
	return "category"
}

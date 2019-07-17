package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("ModifiedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

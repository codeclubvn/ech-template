package model

import uuid "github.com/satori/go.uuid"

type Post struct {
	BaseModel
	Title  string      `json:"title" gorm:"column:title;type:varchar(250);not null"`
	Content string      `json:"content" gorm:"column:content;type:varchar;"`
	Slug    string      `json:"slug" gorm:"column:slug;type:varchar(50);not null"`
	Image   string      `json:"image" gorm:"column:image;type:varchar(250);"`
	UserId uuid.UUID   `json:"user_id" gorm:"column:user_id;type:uuid"`
	User   User        `json:"user" gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Files  FileIdSlice `json:"files" gorm:"column:files;type:uuid[]"`
}

type Posts []Post

type FileIdSlice []uuid.UUID

func (Post) TableName() string {
	return "posts"
}
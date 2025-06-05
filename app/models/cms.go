package models

type CmsArticle struct {
	Model
	Title       string `gorm:"type:varchar(255);not null"`
	Keyword     string `gorm:"type:varchar(255);not null;default:''"`
	Description string `gorm:"type:varchar(255);not null;default:''"`
	Content     string `gorm:"type:longtext"`
	Author      string `gorm:"type:varchar(255);not null;default:''"`
	Status      int8   `gorm:"type:tinyint(1);not null;default:1"`
	CategoryID  uint   `gorm:"type:int(10);not null;default:0"`
	Image       string `gorm:"type:varchar(255);not null;default:''"`
	Views       int64  `gorm:"type:int(10);not null;default:0"`
	IsTop       int8   `gorm:"type:tinyint(1);not null;default:0"`
	IsComment   int8   `gorm:"type:tinyint(1);not null;default:1"`
	Source      string `gorm:"type:varchar(255);not null;default:''"`
	Origin      string `gorm:"type:varchar(255);not null;default:''"`
	Sort        int    `gorm:"type:tinyint(1);not null;default:0"`
	SoftDelete
}

type CmsCategory struct {
	Model
	Name        string `gorm:"type:varchar(255);not null"`
	Image       string `gorm:"type:varchar(255);not null;default:''"`
	Description string `gorm:"type:varchar(255);not null;default:''"`
	Status      int8   `gorm:"type:tinyint(1);not null;default:1"`
	Sort        int    `gorm:"type:tinyint(1);not null;default:0"`
	Pid         uint   `gorm:"type:int(10);not null;default:0"`
	SoftDelete
}

type CmsTag struct {
	Model
	Name        string `gorm:"type:varchar(255);not null"`
	Status      int8   `gorm:"type:tinyint(1);not null;default:1"`
	Description string `gorm:"type:varchar(255);not null;default:''"`
	Sort        int    `gorm:"type:tinyint(1);not null;default:0"`
	SoftDelete
}

type CmsLink struct {
	Model
	Name        string `gorm:"type:varchar(255);not null"`
	Url         string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null;default:''"`
	Status      int8   `gorm:"type:tinyint(1);not null;default:1"`
	Sort        int    `gorm:"type:tinyint(1);not null;default:0"`
	SoftDelete
}

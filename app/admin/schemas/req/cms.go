package req

type LinkQueryReq struct {
	Name   string `json:"name" form:"name"`
	Status int    `json:"status" form:"status" default:"-1"`
}

type LinkAddReq struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Url         string `json:"url" form:"url" validate:"url"`
	Description string `json:"description" form:"description"`
	Status      int8   `json:"status" form:"status"`
	Sort        int    `json:"sort" form:"sort"`
}

type LinkEditReq struct {
	ID          uint   `json:"id" form:"id" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required"`
	Url         string `json:"url" form:"url" validate:"url"`
	Description string `json:"description" form:"description"`
	Status      int8   `json:"status" form:"status"`
	Sort        int    `json:"sort" form:"sort"`
}

type TagQueryReq struct {
	Name   string `json:"name" form:"name"`
	Status int    `json:"status" form:"status" default:"-1"`
}

type TagAddReq struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Status      int8   `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
	Sort        int    `json:"sort" form:"sort"`
}

type TagEditReq struct {
	ID          uint   `json:"id" form:"id" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required"`
	Status      int8   `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
	Sort        int    `json:"sort" form:"sort"`
}

type CategoryQueryReq struct {
	Name   string `json:"name" form:"name"`
	Status int    `json:"status" form:"status" default:"-1"`
	Pid    int    `json:"pid" form:"pid" default:"-1"`
}

type CategoryAddReq struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Image       string `json:"image" form:"image"`
	Description string `json:"description" form:"description"`
	Status      int8   `json:"status" form:"status"`
	Sort        int    `json:"sort" form:"sort"`
	Pid         uint   `json:"pid" form:"pid" validate:"gte=0"`
}

type CategoryEditReq struct {
	ID          uint   `json:"id" form:"id" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required"`
	Image       string `json:"image" form:"image"`
	Description string `json:"description" form:"description"`
	Status      int8   `json:"status" form:"status"`
	Sort        int    `json:"sort" form:"sort"`
	Pid         uint   `json:"pid" form:"pid" validate:"gte=0"`
}

type ArticleQueryReq struct {
	Title      string `json:"title" form:"title"`
	Status     int    `json:"status" form:"status" default:"-1"`
	IsTop      int    `json:"is_top" form:"is_top" default:"-1"`
	IsComment  int    `json:"is_comment" form:"is_comment" default:"-1"`
	CategoryID uint   `json:"category_id" form:"category_id" default:"-1"`
}

type ArticleAddReq struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Keyword     string `json:"keyword" form:"keyword"`
	Description string `json:"description" form:"description"`
	Content     string `json:"content" form:"content" validate:"required"`
	Author      string `json:"author" form:"author"`
	Status      int8   `json:"status" form:"status"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"gte=0"`
	Image       string `json:"image" form:"image"`
	Views       int64  `json:"views" form:"views"`
	IsTop       int8   `json:"is_top" form:"is_top"`
	IsComment   int8   `json:"is_comment" form:"is_comment"`
	Source      string `json:"source" form:"source"`
	Origin      string `json:"origin" form:"origin"`
	Sort        int    `json:"sort" form:"sort"`
}

type ArticleEditReq struct {
	ID          uint   `json:"id" form:"id" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Keyword     string `json:"keyword" form:"keyword"`
	Description string `json:"description" form:"description"`
	Content     string `json:"content" form:"content" validate:"required"`
	Author      string `json:"author" form:"author"`
	Status      int8   `json:"status" form:"status"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"gte=0"`
	Image       string `json:"image" form:"image"`
	Views       int64  `json:"views" form:"views"`
	IsTop       int8   `json:"is_top" form:"is_top"`
	IsComment   int8   `json:"is_comment" form:"is_comment"`
	Source      string `json:"source" form:"source"`
	Origin      string `json:"origin" form:"origin"`
	Sort        int    `json:"sort" form:"sort"`
}

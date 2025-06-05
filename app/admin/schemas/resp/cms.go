package resp

import "new-blog/pkg/types"

type LinkResp struct {
	ID          uint         `json:"id" structs:"id"`
	Name        string       `json:"name" structs:"name"`
	Url         string       `json:"url" structs:"url"`
	Description string       `json:"description" structs:"description"`
	Status      int8         `json:"status" structs:"status"`
	Sort        int          `json:"sort" structs:"sort"`
	CreatedAt   types.TsTime `json:"created_at" structs:"created_at"`
	UpdatedAt   types.TsTime `json:"updated_at" structs:"updated_at"`
}

type TagResp struct {
	ID          uint         `json:"id" structs:"id"`
	Name        string       `json:"name" structs:"name"`
	Status      int8         `json:"status" structs:"status"`
	Description string       `json:"description" structs:"description"`
	Sort        int          `json:"sort" structs:"sort"`
	CreatedAt   types.TsTime `json:"created_at" structs:"created_at"`
	UpdatedAt   types.TsTime `json:"updated_at" structs:"updated_at"`
}

type CategoryResp struct {
	ID          uint           `json:"id" structs:"id"`
	Name        string         `json:"name" structs:"name"`
	Image       string         `json:"image" structs:"image"`
	Description string         `json:"description" structs:"description"`
	Status      int8           `json:"status" structs:"status"`
	Sort        int            `json:"sort" structs:"sort"`
	Pid         uint           `json:"pid" structs:"pid"`
	Children    []CategoryResp `json:"children,omitempty" structs:"children"`
	CreatedAt   types.TsTime   `json:"created_at" structs:"created_at"`
	UpdatedAt   types.TsTime   `json:"updated_at" structs:"updated_at"`
}

type ArticleResp struct {}
				
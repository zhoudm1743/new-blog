package resp

import "new-blog/pkg/types"

type UserResp struct {
	ID        uint         `json:"id" structs:"id"`
	Username  string       `json:"username" structs:"username"`
	Password  string       `json:"password" structs:"password"`
	Email     string       `json:"email" structs:"email"`
	Status    int8         `json:"status" structs:"status"`
	IsAdmin   uint8        `json:"is_admin" structs:"is_admin"`
	CreatedAt types.TsTime `json:"created_at" structs:"created_at"`
	UpdatedAt types.TsTime `json:"updated_at" structs:"updated_at"`
}

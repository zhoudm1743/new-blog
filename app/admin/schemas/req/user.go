package req

type UserQueryReq struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Status   int8   `json:"status" form:"status" default:"-1"`
}

type UserAddReq struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Email    string `json:"email" form:"email" validate:"email"`
	Status   int8   `json:"status" form:"status"`
	IsAdmin  uint8  `json:"is_admin" form:"is_admin" validate:"required"`
}

type UserEditReq struct {
	ID       uint   `json:"id" form:"id" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"email"`
	Status   int8   `json:"status" form:"status"`
	IsAdmin  uint8  `json:"is_admin" form:"is_admin"`
}

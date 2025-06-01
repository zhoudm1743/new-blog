package req

type UserQueryReq struct {}

type UserAddReq struct {}

type UserEditReq struct {
	ID uint `json:"id" form:"id" validate:"required"`
}
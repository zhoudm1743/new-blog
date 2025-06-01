package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"new-blog/app/admin/schemas/req"
	"new-blog/app/admin/service/user"
	"new-blog/app/admin/types"
	"new-blog/pkg/plugins/response"
	"new-blog/pkg/util"
)

type userDep struct {
	fx.In
	Srv user.UserService
}

func UserRoutes(t userDep, r *types.AdminRouter) {
	api := r.Group("/user/user")
	api.GET("/all", t.all)
	api.POST("/list", t.list)
	api.GET("/detail", t.detail)
	api.POST("/add", t.add)
	api.POST("/edit", t.edit)
	api.POST("/delete", t.delete)
}

func (t userDep) all(c *gin.Context) {
	var listReq req.UserQueryReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &listReq)) {
		return
	}
	res, err := t.Srv.All(listReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t userDep) list(c *gin.Context) {
	var page req.PageReq
	var listReq req.UserQueryReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &page, &listReq)) {
		return
	}
	res, err := t.Srv.List(page, listReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t userDep) detail(c *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &idReq)) {
		return
	}
	res, err := t.Srv.Detail(idReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

func (t userDep) add(c *gin.Context) {
	var addReq req.UserAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &addReq)) {
		return
	}
	err := t.Srv.Add(addReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

func (t userDep) edit(c *gin.Context) {
	var editReq req.UserEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &editReq)) {
		return
	}
	err := t.Srv.Edit(editReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

func (t userDep) delete(c *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &idReq)) {
		return
	}
	err := t.Srv.Del(idReq.ID, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

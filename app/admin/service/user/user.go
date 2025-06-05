package user

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"new-blog/app/admin/schemas/req"
	"new-blog/app/admin/schemas/resp"
	"new-blog/app/models"
	"new-blog/pkg/plugins/response"
)

type UserService interface {
	All(listReq req.UserQueryReq, auth *req.AuthReq) ([]resp.UserResp, error)
	List(page req.PageReq, listReq req.UserQueryReq, auth *req.AuthReq) (response.PageResp, error)
	Detail(id uint) (resp.UserResp, error)
	Add(addReq req.UserAddReq, auth *req.AuthReq) error
	Edit(editReq req.UserEditReq, auth *req.AuthReq) error
	Del(id uint, auth *req.AuthReq) error
}

type iUserService struct {
	db *gorm.DB
}

func (iSrv iUserService) All(listReq req.UserQueryReq, auth *req.AuthReq) ([]resp.UserResp, error) {
	chain := iSrv.db.Model(&models.User{}).Order("id desc")
	if listReq.Username != "" {
		chain = chain.Where("username LIKE ?", fmt.Sprintf("%%%s%%", listReq.Username))
	}
	if listReq.Email != "" {
		chain = chain.Where("email LIKE ?", fmt.Sprintf("%%%s%%", listReq.Email))
	}
	if listReq.Status > -1 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	var users []models.User
	if err := chain.Find(&users).Error; err != nil {
		return nil, err
	}
	var res []resp.UserResp
	response.Copy(&res, &users)
	return res, nil
}

func (iSrv iUserService) List(page req.PageReq, listReq req.UserQueryReq, auth *req.AuthReq) (response.PageResp, error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	chain := iSrv.db.Model(&models.User{}).Order("id desc")
	if listReq.Username != "" {
		chain = chain.Where("username LIKE ?", fmt.Sprintf("%%%s%%", listReq.Username))
	}
	if listReq.Email != "" {
		chain = chain.Where("email LIKE ?", fmt.Sprintf("%%%s%%", listReq.Email))
	}
	if listReq.Status > -1 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	var count int64
	if err := chain.Count(&count).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("查询失败: %v", err)
	}
	var users []models.User
	if err := chain.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("查询失败: %v", err)
	}
	var res []resp.UserResp
	response.Copy(&res, &users)
	return response.PageResp{
		Count:    count,
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Lists:    res,
	}, nil
}

func (iSrv iUserService) Detail(id uint) (resp.UserResp, error) {
	var user models.User
	if err := iSrv.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.UserResp{}, fmt.Errorf("记录不存在: %v", err)
		}
		return resp.UserResp{}, err
	}
	var res resp.UserResp
	response.Copy(&res, &user)
	return res, nil
}

func (iSrv iUserService) Add(addReq req.UserAddReq, auth *req.AuthReq) error {
	var count int64
	iSrv.db.Model(&models.User{}).Where("username = ?", addReq.Username).Count(&count)
	if count > 0 {
		return fmt.Errorf("用户名已存在: %s", addReq.Username)
	}
	var user models.User
	response.Copy(&user, &addReq)
	if err := iSrv.db.Create(&user).Error; err != nil {
		return fmt.Errorf("创建失败: %v", err)
	}
	return nil
}

func (iSrv iUserService) Edit(editReq req.UserEditReq, auth *req.AuthReq) error {
	var count int64
	iSrv.db.Model(&models.User{}).Where("username = ?", editReq.Username).Where("id != ?", editReq.ID).Count(&count)
	if count > 0 {
		return fmt.Errorf("用户名已存在: %s", editReq.Username)
	}
	var user models.User
	if err := iSrv.db.First(&user, editReq.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("记录不存在: %v", err)
		}
		return err
	}
	response.Copy(&user, &editReq)
	if err := iSrv.db.Save(&user).Error; err != nil {
		return fmt.Errorf("更新失败: %v", err)
	}
	return nil
}

func (iSrv iUserService) Del(id uint, auth *req.AuthReq) error {
	if !auth.IsAdmin {
		return fmt.Errorf("无权限删除")
	}
	if id == auth.UserId {
		return fmt.Errorf("不能删除自己")
	}
	var user models.User
	if err := iSrv.db.Delete(&user, id).Error; err != nil {
		return fmt.Errorf("删除失败: %v", err)
	}
	return nil
}

func NewUserService(db *gorm.DB) UserService {
	return iUserService{
		db: db,
	}
}

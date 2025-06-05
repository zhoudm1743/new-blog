package cms

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"new-blog/app/admin/schemas/req"
	"new-blog/app/admin/schemas/resp"
	"new-blog/app/models"
	"new-blog/pkg/plugins/response"
)

type TagService interface {
	All(listReq req.TagQueryReq, auth *req.AuthReq) ([]resp.TagResp, error)
	List(page req.PageReq, listReq req.TagQueryReq, auth *req.AuthReq) (response.PageResp, error)
	Detail(id uint) (resp.TagResp, error)
	Add(addReq req.TagAddReq, auth *req.AuthReq) error
	Edit(editReq req.TagEditReq, auth *req.AuthReq) error
	Del(id uint, auth *req.AuthReq) error
}

type iTagService struct {
	db *gorm.DB
}

func (iSrv iTagService) All(listReq req.TagQueryReq, auth *req.AuthReq) ([]resp.TagResp, error) {
	chain := iSrv.db.Model(&models.CmsTag{}).Order("id desc")
	if listReq.Name != "" {
		chain = chain.Where("name LIKE ?", fmt.Sprintf("%%%s%%", listReq.Name))
	}
	if listReq.Status != -1 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	var tags []models.CmsTag
	if err := chain.Find(&tags).Error; err != nil {
		return nil, err
	}
	var res []resp.TagResp
	response.Copy(&res, &tags)
	return res, nil
}

func (iSrv iTagService) List(page req.PageReq, listReq req.TagQueryReq, auth *req.AuthReq) (response.PageResp, error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	chain := iSrv.db.Model(&models.CmsTag{}).Order("id desc")
	if listReq.Name != "" {
		chain = chain.Where("name LIKE ?", fmt.Sprintf("%%%s%%", listReq.Name))
	}
	if listReq.Status != -1 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	var count int64
	if err := chain.Count(&count).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("查询失败: %v", err)
	}
	var tags []models.CmsTag
	if err := chain.Limit(limit).Offset(offset).Find(&tags).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("查询失败: %v", err)
	}
	var res []resp.TagResp
	response.Copy(&res, &tags)
	return response.PageResp{
		Count:    count,
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Lists:    res,
	}, nil
}

func (iSrv iTagService) Detail(id uint) (resp.TagResp, error) {
	var tag models.CmsTag
	if err := iSrv.db.First(&tag, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.TagResp{}, fmt.Errorf("记录不存在: %v", err)
		}
		return resp.TagResp{}, err
	}
	var res resp.TagResp
	response.Copy(&res, &tag)
	return res, nil
}

func (iSrv iTagService) Add(addReq req.TagAddReq, auth *req.AuthReq) error {
	var count int64
	iSrv.db.Model(&models.CmsTag{}).Where("name = ?", addReq.Name).Count(&count)
	if count > 0 {
		return fmt.Errorf("标签名称已存在: %s", addReq.Name)
	}
	var tag models.CmsTag
	response.Copy(&tag, &addReq)
	if err := iSrv.db.Create(&tag).Error; err != nil {
		return fmt.Errorf("创建失败: %v", err)
	}
	return nil
}

func (iSrv iTagService) Edit(editReq req.TagEditReq, auth *req.AuthReq) error {
	var count int64
	iSrv.db.Model(&models.CmsTag{}).Where("name = ?", editReq.Name).Where("id != ?", editReq.ID).Count(&count)
	if count > 0 {
		return fmt.Errorf("标签名称已存在: %s", editReq.Name)
	}
	var tag models.CmsTag
	if err := iSrv.db.First(&tag, editReq.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("记录不存在: %v", err)
		}
		return err
	}
	response.Copy(&tag, &editReq)
	if err := iSrv.db.Save(&tag).Error; err != nil {
		return fmt.Errorf("更新失败: %v", err)
	}
	return nil
}

func (iSrv iTagService) Del(id uint, auth *req.AuthReq) error {
	var tag models.CmsTag
	if err := iSrv.db.Delete(&tag, id).Error; err != nil {
		return fmt.Errorf("删除失败: %v", err)
	}
	return nil
}

func NewTagService(db *gorm.DB) TagService {
	return iTagService{
		db: db,
	}
}

package cms

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"new-blog/app/admin/schemas/req"
	"new-blog/app/admin/schemas/resp"
	"new-blog/app/models"
	"new-blog/pkg/plugins/response"
	"new-blog/pkg/util"
)

type CategoryService interface {
	All(listReq req.CategoryQueryReq, auth *req.AuthReq) (interface{}, error)
	List(page req.PageReq, listReq req.CategoryQueryReq, auth *req.AuthReq) (response.PageResp, error)
	Detail(id uint) (resp.CategoryResp, error)
	Add(addReq req.CategoryAddReq, auth *req.AuthReq) error
	Edit(editReq req.CategoryEditReq, auth *req.AuthReq) error
	Del(id uint, auth *req.AuthReq) error
}

type iCategoryService struct {
	db *gorm.DB
}

func (iSrv iCategoryService) All(listReq req.CategoryQueryReq, auth *req.AuthReq) (interface{}, error) {
	chain := iSrv.db.Model(&models.CmsCategory{}).Order("id desc")
	if listReq.Name != "" {
		chain = chain.Where("name LIKE ?", fmt.Sprintf("%%%s%%", listReq.Name))
	}
	if listReq.Status != -1 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	if listReq.Pid != -1 {
		chain = chain.Where("pid = ?", listReq.Pid)
	}
	var categorys []models.CmsCategory
	if err := chain.Find(&categorys).Error; err != nil {
		return nil, err
	}
	var res []resp.CategoryResp
	response.Copy(&res, &categorys)
	return util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(res), "id", "pid", "children"), nil
}

func (iSrv iCategoryService) List(page req.PageReq, listReq req.CategoryQueryReq, auth *req.AuthReq) (response.PageResp, error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	chain := iSrv.db.Model(&models.CmsCategory{}).Order("id desc")
	if listReq.Name != "" {
		chain = chain.Where("name LIKE ?", fmt.Sprintf("%%%s%%", listReq.Name))
	}
	if listReq.Status != -1 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	if listReq.Pid != -1 {
		chain = chain.Where("pid = ?", listReq.Pid)
	}
	var count int64
	if err := chain.Count(&count).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("查询失败: %v", err)
	}
	var categorys []models.CmsCategory
	if err := chain.Limit(limit).Offset(offset).Find(&categorys).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("查询失败: %v", err)
	}
	var res []resp.CategoryResp
	response.Copy(&res, &categorys)
	return response.PageResp{
		Count:    count,
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Lists:    res,
	}, nil
}

func (iSrv iCategoryService) Detail(id uint) (resp.CategoryResp, error) {
	var category models.CmsCategory
	if err := iSrv.db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CategoryResp{}, fmt.Errorf("记录不存在: %v", err)
		}
		return resp.CategoryResp{}, err
	}
	var res resp.CategoryResp
	response.Copy(&res, &category)
	return res, nil
}

func (iSrv iCategoryService) Add(addReq req.CategoryAddReq, auth *req.AuthReq) error {
	var category models.CmsCategory
	response.Copy(&category, &addReq)
	if err := iSrv.db.Create(&category).Error; err != nil {
		return fmt.Errorf("创建失败: %v", err)
	}
	return nil
}

func (iSrv iCategoryService) Edit(editReq req.CategoryEditReq, auth *req.AuthReq) error {
	var category models.CmsCategory
	if err := iSrv.db.First(&category, editReq.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("记录不存在: %v", err)
		}
		return err
	}
	response.Copy(&category, &editReq)
	if err := iSrv.db.Save(&category).Error; err != nil {
		return fmt.Errorf("更新失败: %v", err)
	}
	return nil
}

func (iSrv iCategoryService) Del(id uint, auth *req.AuthReq) error {
	var count int64
	if err := iSrv.db.Model(&models.CmsCategory{}).Where("pid = ?", id).Count(&count).Error; err != nil {
		return fmt.Errorf("查询失败: %v", err)
	}
	if count > 0 {
		return fmt.Errorf("请先删除子分类")
	}
	var category models.CmsCategory
	if err := iSrv.db.Delete(&category, id).Error; err != nil {
		return fmt.Errorf("删除失败: %v", err)
	}
	return nil
}

func NewCategoryService(db *gorm.DB) CategoryService {
	return iCategoryService{
		db: db,
	}
}

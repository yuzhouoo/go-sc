package services

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"go-session-demo/helpers"
	"go-session-demo/models"
	"go-session-demo/models/request"
	"go-session-demo/models/response"
	"gorm.io/gorm"
)

type AccountService struct {
}

type wheres struct {
	Fields string
	Values []any
}

type dbWhere struct {
	Wheres wheres
	Fields map[string]interface{}
	Offset int
	Limit  int
}

func (As *AccountService) List(params *request.AccountList) (res []response.AccountListItemData, err error) {
	db := &helpers.DB{}
	conn, err := db.InitDB()
	if err != nil {
		return nil, errors.New("数据库失败")
	}

	where := map[string]interface{}{
		"Page":     params.Page,
		"PageSize": params.PageSize,
	}

	list := As.Find(conn, where)

	for _, value := range list {
		item := &response.AccountListItemData{
			ID:      value.ID,
			Account: value.Account,
			Name:    value.Name,
			UUID:    value.UUID.String(),
		}

		res = append(res, *item)
	}

	return res, nil
}

func (As *AccountService) Register(params *request.AccountRegister) (err error) {
	db := &helpers.DB{}
	conn, err := db.InitDB()

	if err != nil {
		return errors.New("数据库连接失败")
	}

	data := &models.AccountModel{
		Name:    params.Name,
		Account: params.Account,
	}

	err = As.Create(conn, data)
	if err != nil {
		return errors.New("新增失败")
	}

	return nil
}

func (As *AccountService) Login(params *request.AccountLogin) (res *models.AccountModel, err error) {
	db := &helpers.DB{}
	conn, err := db.InitDB()

	if err != nil {
		return nil, errors.New("数据库连接失败")
	}

	var where = map[string]interface{}{
		"Fields": "Account = ?",
		"Values": []any{params.Account},
	}

	result, err := As.First(conn, where)
	if err != nil {
		return nil, errors.New("查询账户失败")
	}

	return result, nil
}

func (As *AccountService) EditInfo(params *request.AccountEditInfo) (err error) {
	db := &helpers.DB{}
	conn, err := db.InitDB()

	if err != nil {
		return errors.New("数据库连接失败")
	}

	where := map[string]interface{}{
		"Wheres": map[string]interface{}{
			"Fields": "UUID = ?",
			"Values": []any{params.UUID},
		},
		"Fields": map[string]interface{}{
			"Name":    params.Name,
			"Account": params.Account,
		},
	}

	res := As.Updates(conn, where)

	if res != nil {
		return res
	}

	return res
}

func (As *AccountService) Info(params string) (res *response.AccountInfoItemData, err error) {
	db := &helpers.DB{}
	conn, err := db.InitDB()

	if err != nil {
		return res, errors.New("数据库连接失败")
	}

	where := map[string]interface{}{
		"Wheres": map[string]interface{}{
			"Fields": "UUID = ?",
			"Values": []any{params},
		},
	}

	data, err := As.First(conn, where)
	if err != nil {
		return res, errors.New("未找到用户信息")
	}

	return &response.AccountInfoItemData{
		ID:      data.ID,
		Account: data.Account,
		Name:    data.Name,
		UUID:    data.UUID.String(),
	}, nil
}

func (As *AccountService) Close(params *request.AccountClose) (err error) {
	db := &helpers.DB{}
	conn, err := db.InitDB()

	if err != nil {
		return errors.New("数据库连接失败")
	}

	where := map[string]interface{}{
		"Fields": "uuid = ? AND account = ?",
		"Values": []any{params.UUID, params.Account},
	}
	res := As.Delete(conn, where)

	if res != nil {
		return errors.New("删除失败")
	}

	return nil
}

func (As *AccountService) Find(db *gorm.DB, where map[string]interface{}) (res []*models.AccountModel) {
	var account []*models.AccountModel

	var DBWheres dbWhere
	if page, ok := where["Page"].(int); ok {
		DBWheres.Offset = page
	}

	if pageSize, ok := where["PageSize"].(int); ok {
		DBWheres.Limit = pageSize
	}

	_ = db.Debug().Limit(DBWheres.Limit).Offset(DBWheres.Offset).Find(&account)

	return account
}

func (As *AccountService) First(db *gorm.DB, params map[string]interface{}) (res *models.AccountModel, err error) {
	var result *models.AccountModel

	var DBWheres dbWhere

	if fields, ok := params["Fields"].(string); ok {
		DBWheres.Wheres.Fields = fields
	}

	if values, ok := params["Values"].([]any); ok {
		DBWheres.Wheres.Values = values
	}

	_ = db.Debug().Where(DBWheres.Wheres.Fields, DBWheres.Wheres.Values).First(&result)
	return result, nil
}

func (As *AccountService) Create(db *gorm.DB, params *models.AccountModel) (err error) {
	data := &models.AccountModel{
		Name:    params.Name,
		Account: params.Account,
		UUID:    uuid.NewV1(),
	}

	tx := db.Debug().Create(&data)

	if tx.RowsAffected == 0 {
		return errors.New("创建失败")
	}

	return nil
}

func (As *AccountService) Updates(db *gorm.DB, params map[string]interface{}) (err error) {
	var DbWheres dbWhere

	// 获取where条件字段
	if wheres, ok := params["Wheres"].(map[string]interface{}); ok {
		if fields, fok := wheres["Fields"].(string); fok {
			DbWheres.Wheres.Fields = fields
		}

		if values, vok := wheres["Values"].([]any); vok {
			DbWheres.Wheres.Values = values
		}
	}

	if fields, ok := params["Fields"].(map[string]interface{}); ok {
		DbWheres.Fields = fields
	}

	res := db.Debug().Model(&models.AccountModel{}).Where(DbWheres.Wheres.Fields, DbWheres.Wheres.Values).Updates(&DbWheres.Fields)
	if res.RowsAffected == 0 {
		return errors.New("更新失败")
	}

	return nil
}

func (As *AccountService) Delete(db *gorm.DB, where map[string]interface{}) (err error) {
	var DbWheres dbWhere

	if fields, ok := where["Fields"].(string); ok {
		DbWheres.Wheres.Fields = fields
	}

	if values, ok := where["Values"].([]any); ok {
		DbWheres.Wheres.Values = values
	}

	re := db.Debug().Where(DbWheres.Wheres.Fields, DbWheres.Wheres.Values...).Delete(&models.AccountModel{})

	if re.RowsAffected == 0 {
		return errors.New("删除失败")
	}

	return nil
}

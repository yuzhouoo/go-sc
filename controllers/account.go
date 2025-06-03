package controllers

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-session-demo/helpers"
	"go-session-demo/models/request"
	"go-session-demo/models/response"
	"go-session-demo/services"
	"net/http"
)

type Account struct {
}

func (account *Account) List(ctx *gin.Context) {
	var req = &request.AccountList{}

	_ = ctx.ShouldBindJSON(&req)

	service := services.AccountService{}
	res, err := service.List(req)
	if err != nil {
		ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.ResponseCommon{
			MsgCode: helpers.MSG_CODE_FAIL,
			Desc:    helpers.MSG_DESC_FAIL,
		})

		return
	}

	ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.AccountList{
		ResponseCommon: response.ResponseCommon{
			MsgCode: helpers.MSG_CODE_SUCCESS,
			Desc:    helpers.MSG_DESC_SUCCESS,
		},
		Data: response.ResponseListDataCommon{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    100,
			List:     res,
		},
	})
}

func (account *Account) UpdInfo(ctx *gin.Context) {
	var req = &request.AccountEditInfo{}
	_ = ctx.ShouldBindJSON(&req)

	service := &services.AccountService{}
	res := service.EditInfo(req)

	if res != nil {
		ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.ResponseCommon{
			MsgCode: helpers.MSG_CODE_FAIL,
			Desc:    helpers.MSG_DESC_FAIL,
		})

		return
	}

	ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.ResponseCommon{
		MsgCode: helpers.MSG_CODE_SUCCESS,
		Desc:    helpers.MSG_DESC_SUCCESS,
	})
}

func (account *Account) Register(ctx *gin.Context) {
	var req = &request.AccountRegister{}
	_ = ctx.ShouldBindJSON(&req)

	service := services.AccountService{}
	res := service.Register(req)

	if res != nil {
		ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.ResponseCommon{
			MsgCode: helpers.MSG_CODE_FAIL,
			Desc:    helpers.MSG_DESC_FAIL,
		})

		return
	}

	ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.ResponseCommon{
		MsgCode: helpers.MSG_CODE_SUCCESS,
		Desc:    helpers.MSG_DESC_SUCCESS,
	})
}

func (account *Account) Login(ctx *gin.Context) {
	req := &request.AccountLogin{}

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.ResponseCommon{
			MsgCode: helpers.MSG_CODE_FAIL,
			Desc:    helpers.MSG_DESC_FAIL,
		})

		return
	}

	service := &services.AccountService{}
	res, err := service.Login(req)

	if err != nil {
		ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.ResponseCommon{
			MsgCode: helpers.MSG_CODE_FAIL,
			Desc:    helpers.MSG_DESC_FAIL,
		})
		return
	}

	ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.AccountLogin{
		ResponseCommon: response.ResponseCommon{
			MsgCode: helpers.MSG_CODE_SUCCESS,
			Desc:    helpers.MSG_DESC_SUCCESS,
		},
		Data: response.AccountLoginData{
			ID:      res.ID,
			Account: res.Account,
			Name:    res.Name,
			Token:   uuid.NewV1().String(),
			UUID:    res.UUID.String(),
		},
	})
}

func (account *Account) Check(ctx *gin.Context) {
	token := ctx.GetHeader("Token")
	service := services.AccountService{}
	res, err := service.Info(token)
	if err != nil {
		ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.ResponseCommon{
			MsgCode: helpers.MSG_CODE_FAIL,
			Desc:    helpers.MSG_DESC_FAIL,
		})

		return
	}

	ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.AccountInfoData{
		ResponseCommon: response.ResponseCommon{
			MsgCode: helpers.MSG_CODE_SUCCESS,
			Desc:    helpers.MSG_DESC_SUCCESS,
		},
		Data: *res,
	})
}

func (account *Account) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &response.ResponseCommon{
		MsgCode: 200,
		Desc:    "成功",
	})
}

func (account *Account) Close(ctx *gin.Context) {
	var req = &request.AccountClose{}
	_ = ctx.ShouldBindJSON(&req)

	service := services.AccountService{}
	err := service.Close(req)

	if err != nil {
		ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.ResponseCommon{
			MsgCode: helpers.MSG_CODE_FAIL,
			Desc:    err.Error(),
		})

		return
	}

	ctx.JSON(helpers.HTTP_CODE_SUCCESS, &response.ResponseCommon{
		MsgCode: helpers.MSG_CODE_SUCCESS,
		Desc:    helpers.MSG_DESC_SUCCESS,
	})
}

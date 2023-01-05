package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"wxcloudrun-golang/models"
)

type CreateUserReq struct {
	Phone string `json:"phone"`
	WxInfo string `json:"wxinfo"`
}

type CreateUserResp struct {
	ErrMsg string `json:"err_msg"`
}

func CreateUser(c *gin.Context) {
	var req CreateUserReq
	var resp CreateUserResp
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Error(err)
		resp.ErrMsg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	dao, err := models.GetUserDao()
	if err != nil {
		logrus.WithError(err).Error("GetUserDao failed")
		resp.ErrMsg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	user := models.User{Phone: req.Phone,WxInfo: req.WxInfo}
	err = dao.CreateUser(&user)
	if err != nil {
		logrus.WithError(err).Error("CreateUser failed")
		resp.ErrMsg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.ErrMsg = "success"
	c.JSON(http.StatusOK, resp)
	logrus.Info("success")

}

package user

import (
	"cloud-disk/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	userName := c.DefaultQuery("user_name", "")
	if userName == "" {
		c.JSON(http.StatusBadRequest, common.Failure(http.StatusBadRequest, "the request user name is empty"))
		return
	}

	rsp, err := queryUserInfoTable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Failure(http.StatusInternalServerError, "query user info error"))
		return
	}
	c.JSON(http.StatusOK, common.Success(rsp))
}

func queryUserInfoTable() (*UserInfo, error) {
	return nil, nil
}

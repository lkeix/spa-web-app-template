package api

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"spa-web-app-template/src-backend/libs"
)

func InsertUser(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	ouid := req.PostFormValue("userID")
	fmt.Println(ouid)
	name := req.PostFormValue("name")
	mail := req.PostFormValue("mail")
	err := libs.InsertUserForm(ouid, name, mail)
	if err != nil {
		fmt.Println(err)
	}
}

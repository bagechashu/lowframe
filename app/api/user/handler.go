package user

import (
	"fmt"
	"log"
	"net/http"

	userModel "lowframe/model/user"

	"github.com/unrolled/render"
)

func getDataHandler(rawReader *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		rawReader.JSON(w, http.StatusOK, userModel.UList)
	}
}

func postUserLoginHandler(rawReader *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// 解析url传递的参数，对于POST则解析响应包的主体（request body）
		// 注意:如果没有调用ParseForm方法，下面无法获取表单的数据
		err := req.ParseForm()
		if err != nil {
			log.Println(err)
		}
		if !userModel.IsValid(req.Form) {
			rawReader.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}

		newUser := userModel.ParseUser(req.Form)
		userModel.UList = append(userModel.UList, newUser)

		fmt.Printf("%+v", struct {
			UserList []userModel.User
		}{UserList: userModel.UList})

		rawReader.HTML(w, http.StatusOK, "user/userList", struct {
			UserList []userModel.User
		}{UserList: userModel.UList})
	}
}

func getUserInfoHandler(rawReader *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		rawReader.JSON(w, http.StatusOK, struct {
			NewUser  userModel.User
			UserList []userModel.User
		}{UserList: userModel.UList})
	}

}

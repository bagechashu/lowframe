package user

import (
	userModel "lowframe/model/user"
	"net/http"

	"github.com/unrolled/render"
)

func userAddFormHandler(TmplRender *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		TmplRender.HTML(w, http.StatusOK, "user/index", nil)
	}
}

func postUserInfoHandler(TmplRender *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// 解析url传递的参数，对于POST则解析响应包的主体（request body）
		// 注意:如果没有调用ParseForm方法，下面无法获取表单的数据
		req.ParseForm()
		if !userModel.IsValid(req.Form) {
			TmplRender.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}
		newUser := userModel.ParseUser(req.Form)
		userModel.UList = append(userModel.UList, newUser)
		TmplRender.HTML(w, http.StatusOK, "user/newUser", struct {
			NewUser  userModel.User
			UserList []userModel.User
		}{NewUser: newUser, UserList: userModel.UList})
	}
}

func getUserInfoHandler(TmplRender *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		TmplRender.HTML(w, http.StatusOK, "user/userInfo", nil)
	}
}

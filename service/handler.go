package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func homeHandler(tmplRender *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tmplRender.HTML(w, http.StatusOK, "user/index", nil)
	}
}

func getDataHandler(tmplRender *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tmplRender.JSON(w, http.StatusOK, uList)
	}
}

func postUserInfoHandler(tmplRender *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// 解析url传递的参数，对于POST则解析响应包的主体（request body）
		// 注意:如果没有调用ParseForm方法，下面无法获取表单的数据
		req.ParseForm()
		if !isValid(req.Form) {
			tmplRender.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}
		newUser := parseUser(req.Form)
		uList = append(uList, newUser)
		tmplRender.HTML(w, http.StatusOK, "user/newUser", struct {
			NewUser  user
			UserList []user
		}{NewUser: newUser, UserList: uList})
	}
}

func getUserInfoHandler(tmplRender *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tmplRender.HTML(w, http.StatusOK, "user/userInfo", nil)
	}
}

// notImplemented replies to the request with an HTTP 501 Not Implemented.
func notImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}

// notImplementedHandler returns a simple request handler
// that replies to each request with a ``501 Not Implemented'' reply.
func notImplementedHandler() http.HandlerFunc { return http.HandlerFunc(notImplemented) }

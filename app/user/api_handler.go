package user

import (
	"log"
	"net/http"

	"github.com/unrolled/render"
)

func getDataHandler(rawReader *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		rawReader.JSON(w, http.StatusOK, UList)
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
		if !IsValid(req.Form) {
			rawReader.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}

		newUser := ParseUser(req.Form)
		UList = append(UList, newUser)

		// Debug for render HTML template
		// fmt.Printf("%+v", struct {
		// 	UserList []User
		// }{UserList: UList})

		rawReader.HTML(w, http.StatusOK, "user/userList", struct {
			UserList []User
		}{UserList: UList})
	}
}

func getUserInfoHandler(rawReader *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		rawReader.JSON(w, http.StatusOK, struct {
			NewUser  User
			UserList []User
		}{UserList: UList})
	}

}

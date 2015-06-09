package session

import (
	"net/http"
)

/*
	从session中获取当前登陆用户的UserCode
	在session中取不到userCode返回false, ""
*/

func GetSessionUserCode(w http.ResponseWriter, r *http.Request) (bool, string) {
	userSession := GlobalSessions.SessionStart(w, r)
	userCode := userSession.Get("userCode")
	intercode := userSession.Get("interCode")
	if (userCode == nil || userCode == "") && (intercode == "" || intercode == nil) {
		return false, ""
	}

	return true, string(intercode.(string))
}

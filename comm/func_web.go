package comm

import (
	"LotteryProject/conf"
	"LotteryProject/models"
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

func ClientIP(request *http.Request) string {
	host, _, _ := net.SplitHostPort(request.RemoteAddr) //拆分
	return host
}

func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location", url)
	writer.WriteHeader(http.StatusFound)
}

func SetLoginuser(writer http.ResponseWriter, loginuser *models.ObjLoginUser) {
	if loginuser == nil || loginuser.Uid < 1 {
		c := &http.Cookie{
			Name:   "lottery_loginuser",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(writer, c)
		return
	}

}

func GetLoginUser(request *http.Request) *models.ObjLoginUser {
	c, err := request.Cookie("lottery_loginuser")
	if err != nil {
		return nil
	}
	params, err := url.ParseQuery(c.Value)
	if err != nil {
		return nil
	}
	uid, err := strconv.Atoi(params.Get("uid"))
	if err != nil || uid < 1 {
		return nil
	}
	now, err := strconv.Atoi(params.Get("now"))
	if err != nil || NowUnix()-now > 86400*30 {
		return nil
	}
	loginuser := &models.ObjLoginUser{}
	loginuser.Uid = uid
	loginuser.Username = params.Get("username")
	loginuser.Now = now
	loginuser.Ip = ClientIP(request)
	loginuser.Sign = params.Get("sign")
	sign := createLoginuserSign(loginuser)
	if sign != loginuser.Sign {
		log.Println("func_web GetLoginuser createloginusersign not signed", sign,
			sign, loginuser.Sign)
		return nil
	}

	return loginuser
}

func createLoginuserSign(loginuser *models.ObjLoginUser) string {
	str := fmt.Sprintf("uid=%d&username=%s&secret=%s",
		loginuser.Uid, loginuser.Username, conf.CookieSecret)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return sign
}

package controllers

import (
	"fmt"
	"os"
	"spurt-cms/logger"
	"spurt-cms/models"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	spurtcore "github.com/spurtcms/pkgcore"
	"github.com/spurtcms/pkgcore/auth"
	csrf "github.com/utrack/gin-csrf"
	"gorm.io/gorm"
)

var AUTH auth.Authorization

var Store = sessions.NewCookieStore([]byte("!@#$%"))

var Store1 = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY1")))

func init() {

	ErrorLog = logger.ErrorLOG()

	WarnLog = logger.WarnLOG()

	SetInitialGeneralValues()
}

/*login page view*/
func Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{"csrf": csrf.GetToken(c)})
}

/*forgetpassword page view*/
func ForgetPassword(c *gin.Context) {
	c.HTML(200, "forgotpassword.html", gin.H{"csrf": csrf.GetToken(c)})
}

func NewPassword(c *gin.Context) {

	token := c.Query("token")

	Claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ErrorLog.Printf("Invalid token change password error: %s", err)
			return
		}
	}

	userid := Claims["user_id"]

	c.HTML(200, "changepassword.html", gin.H{"csrf": csrf.GetToken(c), "userid": userid})

}

func SendLinkForForgotPass(c *gin.Context) {

	email := c.PostForm("emailid")

	user, _, err := NewTeam.CheckEmail(email, 0)

	id := user.Id
	token, _ := CreateTokenWithExpireTime(id)

	if err != nil {
		ErrorLog.Printf("Send Link forget password error: %s", err)
		c.SetCookie("log-toast", "You are not registered with us", 3600, "", "", false, false)
		c.Redirect(301, "/forgot")
		return
	}

	var url_prefix = os.Getenv("BASE_URL")

	data := map[string]interface{}{
		"uname":         user.Username,
		"resetpassword": url_prefix + "change-password?token=" + token,
		"restpassurl":   url_prefix + "change-password",
		"admin_logo":    url_prefix + "public/img/spurtcms.png",
		"fb_logo":       url_prefix + "public/img/facebook.png",
		"linkedin_logo": url_prefix + "public/img/linkedin.png",
		"twitter_logo":  url_prefix + "public/img/twitter.png",
	}

	var wg sync.WaitGroup

	wg.Add(1)

	Chan := make(chan string, 1)

	go ForgetPasswordEmail(Chan, &wg, data, email, "Forgot Password")

	close(Chan)

	c.SetCookie("success", "Reset password email send successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/forgot")

}

func SetNewPassword(c *gin.Context) {

	newPassword := c.PostForm("pass")
	confirmPassword := c.PostForm("cpass")

	var user models.TblUser
	user.Id, _ = strconv.Atoi(c.PostForm("userId"))

	if newPassword == confirmPassword {

		User.Authority = &AUTH

		_, err := User.ChangeYourPassword(newPassword)
		if err != nil {
			ErrorLog.Printf("set new password error: %s", err)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			return
		}

		c.SetCookie("get-toast", "Password Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/")
	}
}

/*checklogin*/
func CheckLogin(c *gin.Context) {

	//get values from html
	uname := c.PostForm("username")
	remember := c.PostForm("rememberme")

	token, userid, err := NewAuth.Checklogin(uname, c.Request.PostFormValue("pass"))

	if err != nil {
		ErrorLog.Printf("CheckLogin error: %s", err)
		c.Redirect(301, "/")
		// return
	}

	if err != nil && err.Error() == "user disabled please contact admin" {
		c.SetCookie("log-toast", "This account is inactive please contact the admin", 3600, "", "", false, false)
		c.Redirect(301, "/")
		return
	}

	userdata, _ := NewTeam.GetUserById(userid)

	if !userdata.LastLogin.IsZero() {
		Lastactive := userdata.LastLogin.In(TZONE).Format("02 Jan 2006 03:04 PM")
		Lastlogin[userdata.Id] = Lastactive
	} else {
		Lastlogin[userdata.Id] = "--"
	}

	if c.Request.PostFormValue("username") == "" || c.Request.PostFormValue("pass") == "" {
		c.Redirect(301, "/")
		return
	}

	AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	if gorm.ErrRecordNotFound == err {
		c.SetCookie("log-toast", "You are not registered with us", 3600, "", "", false, false)
		c.Redirect(301, "/")
		return
	}

	if strings.Contains(fmt.Sprint(err), "invalid password") {
		c.SetCookie("username", uname, 3600, "", "", false, false)
		c.SetCookie("pass-toast", "Invalid Password", 3600, "", "", false, false)
		// c.Redirect(301, "/")
		return
	}

	if remember == "1" {
		Session1, _ := Store.Get(c.Request, "REMEMBER_ME")
		Session1.Values["rememberme"] = remember
		Session1.Save(c.Request, c.Writer)
	}

	Session, _ := Store.Get(c.Request, os.Getenv("SESSION_KEY"))
	Session.Values["token"] = token
	Session.Save(c.Request, c.Writer)

	c.Redirect(301, "/dashboard")

}

// logout and clear the sessions
func Logout(c *gin.Context) {

	User.Authority = &AUTH

	User.LastLoginActivity()

	session, err := Store.Get(c.Request, os.Getenv("SESSION_KEY"))
	if err != nil {
		ErrorLog.Printf("Logout session get error: %s", err)
	}

	session.Values["token"] = ""
	session.Options.MaxAge = -1
	er := session.Save(c.Request, c.Writer)
	if er != nil {
		ErrorLog.Printf("Logout session save error: %s", er)
	}

	c.Writer.Header().Set("Pragma", "no-cache")
	c.SetCookie("lang", "", 3600, "", "", false, false)
	c.Redirect(301, "/")

}

func LastActive(c *gin.Context) {
	User.Authority = &AUTH
	User.LastLoginActivity()
}

/*update password*/
func UpdatePassword(c *gin.Context) {

}
func CreateTokenWithExpireTime(userid int) (string, error) {

	var err error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Second * 300).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

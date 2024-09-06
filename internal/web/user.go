package web

import (
	"errors"
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/wangzupeng12061/we-book/internal/domain"
	"github.com/wangzupeng12061/we-book/internal/service"
	"net/http"
)

type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		// 和上面比起来，用 ` 看起来就比较清爽
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,72}$`
	)
	return &UserHandler{
		svc:            svc,
		emailRexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}

}
func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/edit", u.Edit)
	//ug.POST("/login", u.Login)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.LoginJWT)
}
func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	isEmail, err := u.emailRexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "非法邮箱格式")
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次输入密码不对")
		return
	}

	isPassword, err := u.passwordRexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isPassword {
		ctx.String(http.StatusOK, "密码必须包含字母、数字、特殊字符，并且不少于八位")
		return
	}
	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrDuplicateEmail) {
		ctx.String(http.StatusOK, "邮箱已经被注册")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "注册失败")
		return
	}
	ctx.String(http.StatusOK, "注册成功")
}
func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.String(http.StatusOK, "用户不存在或者密码不对")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "登录失败")
		return
	}
	sess := sessions.Default(ctx)

	sess.Set("userId", user.ID)

	sess.Options(sessions.Options{
		//Secure:   false,
		//HttpOnly: false,
		MaxAge: 60,
	})
	sess.Save()
	ctx.String(http.StatusOK, "登录成功")

	return
}
func (u *UserHandler) LoginJWT(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.String(http.StatusOK, "用户不存在或者密码不对")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "登录失败")
		return
	}
	token := jwt.New(jwt.SigningMethodHS512)
	tokenString, err := token.SignedString([]byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	fmt.Println(tokenString)
	fmt.Println(user)
	ctx.Header("x-jwt-token", tokenString)
	ctx.String(http.StatusOK, "登录成功")

	return
}
func (u *UserHandler) Logout(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sess.Options(sessions.Options{
		MaxAge: -1,
	})
	sess.Save()
	ctx.String(http.StatusOK, "退出成功")
}
func (u *UserHandler) Edit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello, edit")
}
func (u *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello, profile")

}

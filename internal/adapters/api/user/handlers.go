package user

import (
	"context"
	"forum/internal/adapters/api"
	"forum/internal/constant"
	"forum/internal/model"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"time"
)

type handlerUser struct {
	ctx        context.Context
	srvUser    api.ServiceUser
	srvSession api.ServiceSession
}

func NewHandler(ctx context.Context, srvUser api.ServiceUser, srvSession api.ServiceSession) api.Handler {
	return &handlerUser{
		ctx:        ctx,
		srvUser:    srvUser,
		srvSession: srvSession,
	}
}

func (hu *handlerUser) Register(router *http.ServeMux) {
	router.HandleFunc(constant.URLSignUp, hu.SignUp)
	router.HandleFunc(constant.URLLogin, hu.Login)
}

func (hu *handlerUser) SignUp(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, " ", r.URL.Path)
	if r.Method == "GET" {
		tmpl, err := api.Parse(constant.PathSignUp)
		if err != nil {
			api.ErrorsHTTP(w, "", constant.Code500)
			return
		}
		api.Execute(w, tmpl, "signup", nil)
		return
	}
	if r.Method != "POST" {
		api.ErrorsHTTP(w, "", constant.Code405)
		return
	}
	ctx, cancel := context.WithTimeout(hu.ctx, constant.TimeLimit)
	defer cancel()

	// create DTO with a new user
	dto := new(model.CreateUserDTO)
	dto.Add(r)
	// and check fields for incorrect data entry
	if !dto.ValidLogin() || !dto.ValidPassword() || !dto.ValidEmail() || !dto.CryptPassword() {
		api.ErrorsHTTP(w, dto.Err, dto.ErrCode)
		return
	}
	// create user in database
	err, code := hu.srvUser.Create(ctx, dto)
	if err != nil {
		api.ErrorsHTTP(w, err.Error(), code)
		return
	}
	// make session
	hu.session(ctx, w, r, dto.Login)
}

func (hu *handlerUser) Login(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, " ", r.URL.Path)
	if r.Method == "GET" {
		tmpl, err := api.Parse(constant.PathLogin)
		if err != nil {
			api.ErrorsHTTP(w, "", constant.Code500)
			return
		}
		api.Execute(w, tmpl, "login", nil)
		return
	}
	if r.Method != "POST" {
		api.ErrorsHTTP(w, "", constant.Code405)
		return
	}
	ctx, cancel := context.WithTimeout(hu.ctx, constant.TimeLimit)
	defer cancel()

	// create DTO with a user
	dto := new(model.CreateUserDTO)
	dto.Add(r)
	// and check fields for incorrect data entry
	if !dto.ValidLogin() || !dto.ValidPassword() || !dto.CryptPassword() {
		api.ErrorsHTTP(w, dto.Err, dto.ErrCode)
		return
	}
	// checks login password
	err, code := hu.srvUser.CheckLoginPassword(ctx, dto)
	if err != nil {
		api.ErrorsHTTP(w, err.Error(), code)
		return
	}
	// make session
	hu.session(ctx, w, r, dto.Login)
}

func (hu *handlerUser) session(ctx context.Context, w http.ResponseWriter, r *http.Request, login string) {
	// create session
	sID, err := uuid.NewV4()
	if err != nil {
		api.ErrorsHTTP(w, err.Error(), constant.Code500)
		return
	}
	// create cookie
	date := time.Now().AddDate(0, 0, constant.SessionExpire)
	c := &http.Cookie{
		Name:    "session",
		Value:   sID.String(),
		Expires: date,
	}
	http.SetCookie(w, c)

	// create session in database
	// if session exist, it will be deleted
	dto := new(model.CreateSessionDTO)
	dto.Add(login, c.Value, date.String())
	err, code := hu.srvSession.Create(ctx, dto)
	if err != nil {
		api.ErrorsHTTP(w, err.Error(), code)
		return
	}
	// go to homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

package model

import (
	"fmt"
	"forum/internal/constant"
	"forum/pkg/password"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/mail"
	"regexp"
	"strconv"
)

type Errs struct {
	Err, ErrCode string
}

type CreateUserDTO struct {
	Login      string
	Password   string
	RePassword string
	Email      string
	ReEmail    string
	Errs
}

func (cu *CreateUserDTO) Add(r *http.Request) {
	cu.Login = r.PostFormValue("login")
	cu.Password = r.PostFormValue("password")
	cu.Email = r.PostFormValue("email")
}

func (cu *CreateUserDTO) ValidLogin() bool {
	if cu.Err != "" {
		return false
	}
	validChar := regexp.MustCompile(`\w`)

	if len(cu.Login) < 3 {
		cu.Err = fmt.Sprintf(constant.TooShort, "login", "three")
		cu.ErrCode = constant.Code403
		return false
	}
	if ok := validChar.MatchString(cu.Login); !ok {
		cu.Err = fmt.Sprintf(constant.InvalidCharacters, "login")
		cu.ErrCode = constant.Code403
		return false
	}
	return true
}

func (cu *CreateUserDTO) ValidPassword() bool {
	if cu.Err != "" {
		return false
	}
	validChar := regexp.MustCompile(`\w`)

	if cu.Password != cu.RePassword {
		cu.Err = fmt.Sprintf(constant.NotMatch, "password")
		cu.ErrCode = constant.Code403
		return false
	}
	if len(cu.Password) < 6 {
		cu.Err = fmt.Sprintf(constant.TooShort, "password", "six")
		cu.ErrCode = constant.Code403
		return false
	}
	if ok := validChar.MatchString(cu.Password); !ok {
		cu.Err = fmt.Sprintf(constant.InvalidCharacters, "password")
		cu.ErrCode = constant.Code403
		return false
	}
	if err := password.ValidPassword(cu.Password); err != nil {
		cu.Err = err.Error()
		cu.ErrCode = constant.Code403
		return false
	}
	return true
}

func (cu *CreateUserDTO) CryptPassword() bool {
	if cu.Err != "" {
		return false
	}
	passGen, err := bcrypt.GenerateFromPassword([]byte(cu.Password), bcrypt.MinCost)
	if err != nil {
		cu.Err = fmt.Sprintf(constant.InternalError, err)
		cu.ErrCode = constant.Code500
		return false
	}
	cu.Password = string(passGen)
	return true
}

func (cu *CreateUserDTO) ValidEmail() bool {
	if cu.Err != "" {
		return false
	}
	if cu.Email != cu.ReEmail {
		cu.Err = fmt.Sprintf(constant.NotMatch, "email")
		cu.ErrCode = constant.Code403
		return false
	}
	_, err := mail.ParseAddress(cu.Email)
	if err != nil {
		cu.Err = fmt.Sprintf(constant.InvalidEnter, "email", "example@mail.com")
		cu.ErrCode = constant.Code403
		return false
	}
	return true
}

type CreatePostDTO struct {
	User       int      `json:"user"`
	Text       string   `json:"post"`
	Categories []string `json:"categories,omitempty"`
	Errs
}

func (cp *CreatePostDTO) Add(r *http.Request) {
	id, err := strconv.Atoi(r.PostFormValue("user"))
	if err != nil {
		cp.Err = err.Error()
		cp.ErrCode = constant.Code500
	}
	cp.User = id
	cp.Text = r.PostFormValue("text")
}

func (cp *CreatePostDTO) Valid() bool {
	if cp.Err != "" {
		return false
	}
	if cp.User == 0 {
		cp.Err = "ID user not correct"
		cp.ErrCode = constant.Code500
	}
	if cp.Text == "" {
		cp.Err = fmt.Sprintf(constant.TooShort, "text", "one")
		cp.ErrCode = constant.Code403
	}
	return true
}

type CreateCategoryDTO struct {
	Name string `json:"name"`
}

func (cc *CreateCategoryDTO) Add(r *http.Request) {
	cc.Name = r.PostFormValue("name")
}

type CreateSessionDTO struct {
	Login string
	UUID  string
	Date  string
}

func (cs *CreateSessionDTO) Add(login, uuid, date string) {
	cs.Login = login
	cs.UUID = uuid
	cs.Date = date
}

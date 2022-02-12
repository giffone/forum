package constant

import (
	"errors"
	"time"
)

const (
	URLHome            = "/"
	URLSignUp          = "/signup"
	URLLogin           = "/login"
	URLPost            = "/post"
	URLCategory        = "/category"
	URLFavicon         = "/favicon.ico"
	PathIndex          = "internal/web/templates/index.gohtml"
	PathHead           = "internal/web/templates/head.gohtml"
	PathFooter         = "internal/web/templates/footer.gohtml"
	PathLogin          = "internal/web/templates/login.gohtml"
	PathSignUp         = "internal/web/templates/signup.gohtml"
	PathPosts          = "internal/web/templates/posts.gohtml"
	PathCategories     = "internal/web/templates/categories.gohtml"
	SessionExpire      = 90               // days
	TimeLimit          = 10 * time.Second // context
	Code401            = "401"            // http.StatusUnauthorized
	Code403            = "403"            // http.StatusForbidden
	Code405            = "405"            // http.StatusMethodNotAllowed
	Code422            = "422"            // http.StatusUnprocessableEntity
	Code500            = "500"            // http.StatusInternalServerError
	LoginPasswordWrong = "no match: %s or %s is wrong"
	AlreadyExist       = "can not create new user: %s or %s already registered"
	InvalidCharacters  = "the %s contains invalid characters"
	TooShort           = "the %s is too short, must be at least %s characters"
	NotMatch           = "the entered %s does not match"
	InvalidEnter       = "the entered %s is incorrect, please use valid for example: \"%s\""
	InternalError      = "internal error: \"%v\""
)

var (
	Err500   = errors.New("internal server error")
	ErrExist = errors.New("exist")
	ErrCat   = errors.New("cat")
)

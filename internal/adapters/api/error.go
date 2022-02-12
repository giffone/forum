package api

import (
	"forum/internal/constant"
	"net/http"
)

func ErrorsHTTP(w http.ResponseWriter, err, code string) {
	if code == constant.Code401 {
		if err == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		http.Error(w, err, http.StatusUnauthorized)
		return
	}
	if code == constant.Code403 {
		if err == "" {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		http.Error(w, err, http.StatusForbidden)
		return
	}
	if code == constant.Code405 {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if code == constant.Code422 {
		http.Error(w, err, http.StatusUnprocessableEntity)
		return
	}
	if code == constant.Code500 {
		if err == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Error(w, err, http.StatusInternalServerError)
		return
	}
	http.Error(w, "error not handled "+code, http.StatusInternalServerError)
}

package model

import (
	"forum/internal/object"
)

type Users struct {
	Users []*User
	St    *object.Settings
	Ck    *object.Cookie
}

//func NewUsers(st *object.Settings, ck *object.Cookie) *Users {
//	u := new(Users)
//	if st == nil {
//		u.St = &object.Settings{
//			Key: make(map[string][]interface{}),
//		}
//	} else {
//		u.St = st
//	}
//	if ck == nil {
//		u.Ck = new(object.Cookie)
//	} else {
//		u.Ck = ck
//	}
//	return u
//}

func (u *Users) GetList() *object.QuerySettings {
	return get(u.St.Key)
}

func (u *Users) NewList() []interface{} {
	user := new(User)
	u.Users = append(u.Users, user)
	return []interface{}{
		&user.ID,
		&user.Login,
		&user.Password,
		&user.Email,
		&user.Root,
		&user.Created,
	}
}

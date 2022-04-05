package model

import (
	"forum/internal/constant"
	"forum/internal/object"
)

type Categories struct {
	Slice []*Category
	St    *object.Settings
	Ck    *object.Cookie
}

func NewCategories(st *object.Settings, ck *object.Cookie) *Categories {
	p := new(Categories)
	if st == nil {
		p.St = &object.Settings{
			Key: make(map[string][]interface{}),
		}
	} else {
		p.St = st
	}
	if ck == nil {
		p.Ck = new(object.Cookie)
	} else {
		p.Ck = ck
	}
	return p
}

func (c *Categories) MakeKeys(key string, data ...interface{}) {
	if key != "" {
		c.St.Key[key] = data
	} else {
		c.St.Key[constant.KeyPost] = []interface{}{0}
	}
}

func (c *Categories) GetList() *object.QuerySettings {
	if value, ok := c.St.Key[constant.KeyPost]; ok {
		return &object.QuerySettings{
			QueryName: constant.QueSelectCategoryBy,
			QueryFields: []interface{}{
				constant.FieldPost,
			},
			Fields: value,
		}
	}
	return &object.QuerySettings{
		QueryName: constant.QueSelectCategories,
	}
}

func (c *Categories) NewList() []interface{} {
	category := new(Category)
	c.Slice = append(c.Slice, category)
	return []interface{}{
		&category.ID,
		&category.Name,
	}
}

func (c *Categories) IfNil() interface{} {
	return []*Category{new(Category).ifNil()}
}

func (c *Categories) Return() *Buf {
	return &Buf{Categories: c}
}

package post

import (
	"context"
	"fmt"
	"forum/internal/adapters/repo/lib"
	"forum/internal/config"
	"forum/internal/constant"
	"forum/internal/model"
	"log"
)

func (sp *storagePost) create(ctx context.Context, post *model.CreatePostDTO) error {
	query := fmt.Sprintf(sp.query.Queries[config.KeyInsert2], config.TabPosts, "user", "post")

	res, err := sp.db.ExecContext(ctx, query, post.User, post.Text)
	if err != nil {
		log.Printf("post create: %v", err)
		return constant.Err500
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("post create: last inserted id: %v", err)
		return constant.Err500
	}

	if post.Categories != nil {
		err = sp.createPostCategory(ctx, post.Categories, id)
		if err != nil {
			log.Printf("post %d create: categories create: %v", id, err)
			return constant.ErrCat
		}
	}
	return nil
}

func (sp *storagePost) createPostCategory(ctx context.Context, categories []string, id int64) error {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	query := fmt.Sprintf(sp.query.Queries[config.KeyInsert2], config.TabPostsCategories, "post", "category")
	for _, cat := range categories {
		_, err := sp.db.ExecContext(ctx2, query, id, cat)
		if err != nil {
			log.Printf("execute categories: %v", err)
			return constant.ErrCat
		}
	}
	return nil
}

func (sp *storagePost) createTx(ctx context.Context, post *model.CreatePostDTO) error {
	tx, err := lib.TxBegin(ctx, sp.db)
	if err != nil {
		return err
	}
	defer lib.TxRollBack(tx)

	query := fmt.Sprintf(sp.query.Queries[config.KeyInsert2], config.TabPosts, "user", "post")
	res, err := tx.ExecContext(ctx, query, post.User, post.Text)
	if err != nil {
		log.Printf("post create: %v", err)
		return constant.Err500
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("post create: last inserted id: %v", err)
		return constant.Err500
	}

	if post.Categories != nil {
		err = sp.createPostCategoryTx(ctx, post.Categories, id)
		if err != nil {
			log.Printf("post %d create: categories create: %v", id, err)
			return constant.ErrCat
		}
	}

	lib.TxCommit(tx)
	return nil
}

func (sp *storagePost) createPostCategoryTx(ctx context.Context, categories []string, id int64) error {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	tx, err := lib.TxBegin(ctx2, sp.db)
	if err != nil {
		return err
	}

	defer lib.TxRollBack(tx)

	query := fmt.Sprintf(sp.query.Queries[config.KeyInsert2], config.TabPostsCategories, "post", "category")
	for _, cat := range categories {
		_, err := tx.ExecContext(ctx2, query, id, cat)
		if err != nil {
			log.Printf("execute categories: %v", err)
			return constant.ErrCat
		}
	}

	lib.TxCommit(tx)
	return nil
}

package comment

import (
	"context"
	"forum/internal/adapters/repository"
	"forum/internal/constant"
	"forum/internal/object"
	"forum/internal/object/dto"
	"forum/internal/object/model"
	"forum/internal/service"
)

type sComment struct {
	repo  repository.Repo
	sLike service.Ratio
}

func NewService(repo repository.Repo, sLike service.Ratio) service.Comment {
	return &sComment{
		repo:  repo,
		sLike: sLike,
	}
}

func (sc *sComment) Create(ctx context.Context, d *dto.Comment) (int, object.Status) {
	id, sts := sc.repo.Create(ctx, d)
	if sts != nil {
		return 0, sts
	}
	return id, nil
}

func (sc *sComment) Delete(ctx context.Context, id int) object.Status {
	return nil
}

func (sc *sComment) Get(ctx context.Context, m model.Models) (interface{}, object.Status) {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimitDB)
	defer cancel()

	sts := sc.repo.GetList(ctx2, m)
	if sts != nil {
		return nil, sts
	}

	comments := m.Return().Comments
	lSlice := len(comments.Slice)
	if lSlice == 0 {
		return comments.IfNil(), nil
	}
	// checks if authorized user liked comment
	if comments.Ck.Session {
		// checks liked comment or not
		sts = sc.sLike.Liked(ctx2, comments)
		if sts != nil {
			return nil, sts
		}
	}
	// checks need refer to ... or not
	if comments.St.Refers {
		// make refer to own post
		sts = sc.refer(ctx2, comments)
		if sts != nil {
			return nil, sts
		}
	}
	// count likes/dislikes
	sts = sc.sLike.CountFor(ctx2, comments)
	if sts != nil {
		return nil, sts
	}
	return comments.Slice, nil
}

func (sc *sComment) GetChan(ctx context.Context, m model.Models) (interface{}, object.Status) {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimitDB)
	defer cancel()

	sts := sc.repo.GetList(ctx2, m)
	if sts != nil {
		return nil, sts
	}

	comments := m.Return().Comments
	lSlice := len(comments.Slice)
	if lSlice == 0 {
		return comments.IfNil(), nil
	}

	channel := make(chan object.Status)
	// checks if authorized user liked comment
	if comments.Ck.Session {
		// checks liked comment or not
		go sc.sLike.LikedChan(ctx2, comments, channel)
	} else {
		channel <- nil
	}
	// checks need refer to ... or not
	if comments.St.Refers {
		// make refer to own post
		go sc.referChan(ctx2, comments, channel)
	} else {
		channel <- nil
	}
	// count likes/dislikes
	go sc.sLike.CountForChan(ctx2, comments, channel)

	err1 := <-channel
	err2 := <-channel
	err3 := <-channel

	if err1 != nil || err2 != nil || err3 != nil {
		return nil, sts
	}
	return comments.Slice, nil
}

func (sc *sComment) refer(ctx context.Context, c *model.Comments) object.Status {
	for i := 0; i < len(c.Slice); i++ {
		p := model.NewPost(nil, nil)
		p.MakeKeys(constant.KeyPost, c.Slice[i].Post)
		sts := sc.repo.GetOne(ctx, p)
		if sts != nil {
			return sts
		}
		c.Slice[i].Title = p
	}
	return nil
}

func (sc *sComment) referChan(ctx context.Context, c *model.Comments, channel chan object.Status) {
	for i := 0; i < len(c.Slice); i++ {
		p := model.NewPost(nil, nil)
		p.MakeKeys(constant.KeyPost, c.Slice[i].Post)
		sts := sc.repo.GetOne(ctx, p)
		if sts != nil {
			channel <- sts
			return
		}
		c.Slice[i].Title = p
	}
	channel <- nil
}

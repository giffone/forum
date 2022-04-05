package ratio

import (
	"context"
	"forum/internal/adapters/repository"
	"forum/internal/constant"
	"forum/internal/object"
	"forum/internal/object/dto"
	"forum/internal/object/model"
	"forum/internal/service"
	"log"
)

type sRatio struct {
	repo repository.Repo
}

func NewService(repo repository.Repo) service.Ratio {
	return &sRatio{repo: repo}
}

func (sr *sRatio) Create(ctx context.Context, d *dto.Ratio) (int, object.Status) {
	like := model.NewLike(nil, d.Obj.Ck)
	post := false
	// post
	if id, ok := d.PostOrComm[constant.KeyPost]; ok {
		post = true
		// post - keys for get likes from db
		like.MakeKeys(constant.KeyPost, d.Obj.Ck.User, id)
		// post - keys for create like in db
		d.MakeKeys(constant.KeyPost, d.Obj.Ck.User, id)
		// comment
	} else if id, ok := d.PostOrComm[constant.KeyComment]; ok {
		// comment - keys for get likes from db
		like.MakeKeys(constant.KeyComment, d.Obj.Ck.User, id)
		// comment - keys for create like in db
		d.MakeKeys(constant.KeyComment, d.Obj.Ck.User, id)
	} else {
		// if not post or comment
		return 0, object.StatusByCode(constant.Code400)
	}
	//check like exist by user_id and post_id/comment_id
	sts := sr.repo.GetOne(ctx, like)
	if sts != nil {
		return 0, sts
	}
	// not exist
	if like.PostOrComm == 0 {
		// keys - here need only key and ignore value [in DTO method create]
		id, sts := sr.repo.Create(ctx, d)
		if sts != nil {
			return 0, sts
		}
		// post_id for return page (redirect)
		return id, nil
	}
	// DTO for delete object
	dDelete := dto.NewRatio(nil, nil, d.Obj.Ck)
	// make keys for delete by id
	if post {
		dDelete.MakeKeys(constant.KeyPost, like.PostOrComm)
	} else {
		dDelete.MakeKeys(constant.KeyComment, like.PostOrComm)
	}
	// delete
	sts = sr.repo.Delete(ctx, dDelete)
	if sts != nil {
		return 0, sts
	}
	// is same - was like and new like (not create new)
	if like.Like == d.Like {
		// post_id for return page (redirect)
		return d.Obj.Ck.Post, nil
	}
	// is not same - create new
	id, sts := sr.repo.Create(ctx, d)
	if sts != nil {
		return 0, sts
	}
	// post_id for return page (redirect)
	return id, nil
}

func (sr *sRatio) CountFor(ctx context.Context, pc model.PostOrComment) object.Status {
	for i := 0; i < pc.LSlice(); i++ {
		id := pc.PostOrCommentID(i)
		likesCount := model.NewLikesCount(pc.Settings().ClearKey(), pc.Cookie()) // auto insert session
		likesCount.MakeKeys(pc.KeyRole(), id)
		// for make map["post"]id
		likesCount.PostOrComm = id
		sts := sr.repo.GetList(ctx, likesCount)
		if sts != nil {
			return sts
		}
		lSlice := len(likesCount.Slice)
		if lSlice == 0 {
			pc.Add(constant.KeyLike, i, likesCount.IfNil())
		} else {
			// like or dislike only, need to show another with 0
			if lSlice == 1 {
				if likesCount.Slice[0].Body == constant.FieldLike {
					likesCount.Slice = append(likesCount.Slice, likesCount.DislikeNil())
				} else {
					likesCount.Slice = append(likesCount.Slice, likesCount.LikeNil())
				}
			}
			pc.Add(constant.KeyLike, i, likesCount.Slice)
		}
	}
	return nil
}

func (sr *sRatio) Liked(ctx context.Context, pc model.PostOrComment) object.Status {
	user := pc.Cookie().User
	for i := 0; i < pc.LSlice(); i++ {
		id := pc.PostOrCommentID(i)
		like := model.NewLike(nil, nil)
		like.MakeKeys(pc.KeyLiked(), user, id)
		sts := sr.repo.GetOne(ctx, like)
		if sts != nil {
			return sts
		}
		pc.Add(constant.KeyRated, i, like.Body)
	}
	return nil
}

func (sr *sRatio) CountForChan(ctx context.Context, pc model.PostOrComment, channel chan object.Status) {
	log.Println("in CountForChan")
	for i := 0; i < pc.LSlice(); i++ {
		id := pc.PostOrCommentID(i)
		likesCount := model.NewLikesCount(pc.Settings().ClearKey(), pc.Cookie()) // auto insert session
		likesCount.MakeKeys(pc.KeyRole(), id)
		// for make map["post"]id
		likesCount.PostOrComm = id
		sts := sr.repo.GetList(ctx, likesCount)
		if sts != nil {
			log.Println("err CountForChan")
			channel <- sts
			return
		}
		lSlice := len(likesCount.Slice)
		if lSlice == 0 {
			pc.Add(constant.KeyLike, i, likesCount.IfNil())
		} else {
			// like or dislike only, need to show another with 0
			if lSlice == 1 {
				if likesCount.Slice[0].Body == constant.FieldLike {
					likesCount.Slice = append(likesCount.Slice, likesCount.DislikeNil())
				} else {
					likesCount.Slice = append(likesCount.Slice, likesCount.LikeNil())
				}
			}
			pc.Add(constant.KeyLike, i, likesCount.Slice)
		}
	}
	log.Println("out CountForChan")
	channel <- nil
}

func (sr *sRatio) LikedChan(ctx context.Context, pc model.PostOrComment, channel chan object.Status) {
	log.Println("in LikedChan")
	user := pc.Cookie().User
	for i := 0; i < pc.LSlice(); i++ {
		id := pc.PostOrCommentID(i)
		like := model.NewLike(nil, nil)
		like.MakeKeys(pc.KeyLiked(), user, id)
		sts := sr.repo.GetOne(ctx, like)
		if sts != nil {
			log.Println("err LikedChan")
			channel <- sts
			return
		}
		pc.Add(constant.KeyRated, i, like.Body)
	}
	log.Println("out LikedChan")
	channel <- nil
}

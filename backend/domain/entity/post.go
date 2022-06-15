package entity

import (
	"backend/pkg/common"
	"errors"
	"time"

	"github.com/hashicorp/go-multierror"
)

type Post struct {
	ID		  	common.ID
	Title	  	string
	Description string
	Content	  	string
	UrlVideo	string
	CategoryID  string
	UserID	  	string
	ParentID	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

type PostDto struct {
	Title	  	string
	Description string
	Content	  	string
	UrlVideo	string
	CategoryID  string
	UserID	  	string
	ParentID	string
}

func NewPost(payload PostDto) *Post {
	create := Post{
		ID:		  		common.NewID(),
		Title:	  		payload.Title,
		Description: 	payload.Description,
		Content:	  	payload.Content,
		UrlVideo: 		payload.UrlVideo,
		CategoryID: 	payload.CategoryID,
		UserID: 		payload.UserID,
		ParentID: 		payload.ParentID,
		CreatedAt: 		time.Now().UTC(),
		UpdatedAt: 		time.Now().UTC(),
	}

	return &create
}

func (x *Post) Validate() *multierror.Error {
	var multilerr *multierror.Error

	if x.Title == "" {
		multilerr = multierror.Append(multilerr, errors.New("title required"))
	}

	if x.Description == "" {
		multilerr = multierror.Append(multilerr, errors.New("description required"))
	}

	if x.Content == "" {
		multilerr = multierror.Append(multilerr, errors.New("content required"))
	}

	if x.UrlVideo == "" {
		multilerr = multierror.Append(multilerr, errors.New("url video required"))
	}

	if x.CategoryID == "" {
		multilerr = multierror.Append(multilerr, errors.New("category id required"))
	}

	if x.UserID == "" {
		multilerr = multierror.Append(multilerr, errors.New("user id required"))
	}

	return multilerr
}
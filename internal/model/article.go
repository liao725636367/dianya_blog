package model

import "blog/internal/model/entity"

type Article struct {
	*entity.BlogArticle
	CategoryName string `json:"categoryName"`
}

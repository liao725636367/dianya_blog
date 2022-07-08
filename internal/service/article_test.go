package service

import (
	"blog/apiv1"
	"context"
	"fmt"
	"testing"
)

func TestSArticle_List(t *testing.T) {
	ctx := context.Background()
	_, list, err := Article().List(ctx, &apiv1.ArticleListReq{
		Page:    1,
		Size:    10,
		Keyword: "美国",
	})
	if err != nil {
		t.Log(err)
		return
	}
	for _, article := range list {
		fmt.Println(article, article.Id, article.CategoryName)
	}
}

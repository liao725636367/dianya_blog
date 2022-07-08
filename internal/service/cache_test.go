package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/util/gutil"

	"github.com/gogf/gf/v2/frame/g"
)

func TestCache(t *testing.T) {
	ctx := context.Background()

	reVar, err := g.Redis().Do(ctx, "HSCAN", "support", "0", "count", "100")
	if err != nil {
		panic(err)
	}
	scanSlice := reVar.Vars()
	cursor := scanSlice[0].Int()
	data := gutil.SliceToMap(scanSlice[1].Slice())
	//fmt.Println(reVar)
	fmt.Println(cursor, data)
}
func TestSCache_GetArticleRedisSupport(t *testing.T) {
	ctx := context.Background()

	res, err := Cache().GetArticleRedisSupport(ctx, []uint{1, 2, 3})
	if err != nil {
		t.Log(err)
		return
	}
	for _, re := range res {
		fmt.Println(re)
	}
}
func TestSCache_SupportIncr(t *testing.T) {
	ctx := context.Background()

	res, err := Cache().SupportIncr(ctx, 1, -1)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(res)
}
func TestCacheWrite(t *testing.T) {
	ctx := context.Background()
	for i := 0; i < 1000; i++ {
		g.Redis().Do(ctx, "HSET", "support", i, i+100)
	}

}

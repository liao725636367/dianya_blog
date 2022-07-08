package service

import (
	"blog/internal/consts"
	"blog/internal/service/internal/dao"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcache"

	"github.com/gogf/gf/v2/util/gutil"

	"github.com/gogf/gf/v2/os/gcron"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

type ICache interface {
	GetArticleRedisSupport(ctx context.Context, ids []uint) (res []*NumItem, err error)
	GetArticleRedisVisit(ctx context.Context, ids []uint) (res []*NumItem, err error)
	SupportIncr(ctx context.Context, articleId uint, value int) (num int, err error)
	VisitIncr(ctx context.Context, articleId uint, value int) (num int, err error)
}
type (
	sCache struct{}
)

var (
	insCache    = sCache{}
	cacheClient = gcache.New()
)

func Cache() ICache {
	return &insCache
}
func init() {
	redisCache := gcache.NewAdapterRedis(g.Redis())
	//使用redis 作为默认缓存
	cacheClient.SetAdapter(redisCache)

	//每天凌晨将点赞更新到数据库
	//TODO 测试暂时每分钟
	_, err := gcron.AddSingleton(context.Background(), "0 * * * * *", func(ctx context.Context) {

		var cursor = -1
		for cursor != 0 {
			if cursor < 0 {
				cursor = 0
			}
			reVar, err := g.Redis().Do(ctx, "HSCAN", consts.RedisSupport, cursor, "count", "100")
			if err != nil {
				g.Log().Error(ctx, err)
				return
			}
			scanSlice := reVar.Vars()
			cursor = scanSlice[0].Int()
			data := gutil.SliceToMap(scanSlice[1].Slice())
			removeIds := make([]interface{}, 0)
			removeKeys := make([]interface{}, 0)
			for i, i2 := range data {
				id := gconv.Uint(i)
				num := gconv.Int(i2)
				if id > 0 && num > 0 {
					_, _ = dao.BlogArticle.Ctx(ctx).Where("id=?", id).Increment("support", num)
					removeKeys = append(removeKeys, fmt.Sprintf(consts.RedisArticleRow, id))
				}
				removeIds = append(removeIds, id)
			}
			if len(removeIds) > 0 {
				_, _ = g.Redis().Do(ctx, "HDEL", append(g.Slice{consts.RedisSupport}, removeIds...)...)
			}
			if len(removeKeys) > 0 {
				_ = cacheClient.Removes(ctx, removeKeys)

			}
		}

	}, "support_refresh")
	if err != nil {
		panic("创建任务失败" + err.Error())
	}
	//每天凌晨将阅读数据更新数据库
	//TODO 测试暂时每分钟
	_, err = gcron.AddSingleton(context.Background(), "0 * * * * *", func(ctx context.Context) {

		var cursor = -1
		for cursor != 0 {
			if cursor < 0 {
				cursor = 0
			}
			reVar, err := g.Redis().Do(ctx, "HSCAN", consts.RedisArticleVisit, cursor, "count", "100")
			if err != nil {
				g.Log().Error(ctx, err)
				return
			}
			scanSlice := reVar.Vars()
			cursor = scanSlice[0].Int()
			data := gutil.SliceToMap(scanSlice[1].Slice())
			removeIds := make([]interface{}, 0)
			removeKeys := make([]interface{}, 0)
			for i, i2 := range data {
				id := gconv.Uint(i)
				num := gconv.Int(i2)
				if id > 0 && num > 0 {
					_, _ = dao.BlogArticle.Ctx(ctx).Where("id=?", id).Increment("visit", num)
					removeKeys = append(removeKeys, fmt.Sprintf(consts.RedisArticleRow, id))
				}
				removeIds = append(removeIds, id)
			}
			if len(removeIds) > 0 {
				_, _ = g.Redis().Do(ctx, "HDEL", append(g.Slice{consts.RedisArticleVisit}, removeIds...)...)
			}
			if len(removeKeys) > 0 {
				_ = cacheClient.Removes(ctx, removeKeys)

			}
		}

	}, "visit_refresh")
	if err != nil {
		panic("创建任务失败" + err.Error())
	}
}

type NumItem struct {
	Id  uint
	Num int
}

//获取缓存中的文章点赞信息
func (s *sCache) GetArticleRedisSupport(ctx context.Context, ids []uint) (res []*NumItem, err error) {
	res = make([]*NumItem, 0)
	listRe, err := g.Redis().Do(ctx, "HMGET", append(g.Slice{consts.RedisSupport}, gconv.Interfaces(ids)...)...)
	if err != nil {
		return nil, err
	}
	if listRe.IsEmpty() {
		return res, nil
	}
	//nums := make([]int, 0)
	nums := listRe.Ints()
	if err != nil {
		return nil, err
	}
	if len(nums) != len(ids) {
		return res, nil
	}
	for i, id := range ids {
		res = append(res, &NumItem{
			Id:  id,
			Num: nums[i],
		})
	}

	return res, nil
}

//获取缓存中的文章阅读信息
func (s *sCache) GetArticleRedisVisit(ctx context.Context, ids []uint) (res []*NumItem, err error) {
	res = make([]*NumItem, 0)
	listRe, err := g.Redis().Do(ctx, "HMGET", append(g.Slice{consts.RedisArticleVisit}, gconv.Interfaces(ids)...)...)
	if err != nil {
		return nil, err
	}
	if listRe.IsEmpty() {
		return res, nil
	}
	//nums := make([]int, 0)
	nums := listRe.Ints()
	if err != nil {
		return nil, err
	}
	if len(nums) != len(ids) {
		return res, nil
	}
	for i, id := range ids {
		res = append(res, &NumItem{
			Id:  id,
			Num: nums[i],
		})
	}

	return res, nil
}

// SupportIncr 内存增加点赞
func (s *sCache) SupportIncr(ctx context.Context, articleId uint, value int) (num int, err error) {
	reVar, err := g.Redis().Do(ctx, "HINCRBY", consts.RedisSupport, articleId, value)
	if err != nil {
		return 0, err
	}
	return reVar.Int(), nil
}

//增加文章阅读量
func (s *sCache) VisitIncr(ctx context.Context, articleId uint, value int) (num int, err error) {
	reVar, err := g.Redis().Do(ctx, "HINCRBY", consts.RedisArticleVisit, articleId, value)
	if err != nil {
		return 0, err
	}
	return reVar.Int(), nil
}

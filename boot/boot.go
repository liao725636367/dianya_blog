package boot

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

func init() {
	redisCache := gcache.NewAdapterRedis(g.Redis())
	//使用redis 作为默认缓存
	//cache := gcache.New()
	//cache.SetAdapter(redisCache)
	//使用redis 作为db 默认缓存
	g.DB().GetCache().SetAdapter(redisCache)
	//测试连接

	//gcache.cachSet(context.Background(), "a", 1, time.Minute)
	//re, err := cache.Get(context.Background(), "a")
	//fmt.Println(re, err)
}

package consts

//其它定义
const UserToken = "token" //token 格式 秘钥_过期时间戳 每一个小时自动续签
const TokenExpire = 86400 //token过期时间 一天

const RoleAdmin = 1 //管理员角色

//缓存key
//用户缓存
const RedisUserRow = "user_row_%d"

//文章缓存
const RedisArticleRow = "article_row_%d"
const RedisArticleVisit = "visit"

//文章点赞hash key
const RedisSupport = "support"

//文章分类缓存
const RedisCategory = "category"

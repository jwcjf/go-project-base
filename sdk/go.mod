module github.com/jwcjf/go-project-base/sdk

go 1.15

require (
	github.com/jwcjf/go-project-base v0.0.4
	github.com/jwcjf/go-project-base/plugins/logger/zap v0.0.4
	git.code.oa.com/polaris/polaris-go v0.8.7
	github.com/bsm/redislock v0.7.1
	github.com/casbin/casbin/v2 v2.31.3
	github.com/chanxuehong/wechat v0.0.0-20201110083048-0180211b69fd
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-redis/redis/v7 v7.4.0
	github.com/google/uuid v1.2.0
	github.com/gorilla/websocket v1.4.2
	github.com/jackc/pgconn v1.8.1
	github.com/mojocn/base64Captcha v1.3.4
	github.com/robfig/cron/v3 v3.0.1
	github.com/robinjoseph08/redisqueue/v2 v2.1.0
	github.com/shamsher31/goimgext v1.0.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gorm.io/driver/mysql v1.1.0
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.7
	gorm.io/gorm v1.21.10
)

replace github.com/bsm/redislock v0.7.1 => github.com/bsm/redislock v0.5.0

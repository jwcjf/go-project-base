module github.com/jwcjf/go-project-base

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/bitly/go-simplejson v0.5.0
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/bsm/redislock v0.5.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.7.2
	github.com/go-redis/redis/v7 v7.4.0
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/imdario/mergo v0.3.12
	github.com/json-iterator/go v1.1.9
	github.com/nsqio/go-nsq v1.0.8
	github.com/pkg/errors v0.9.1
	github.com/robinjoseph08/redisqueue/v2 v2.1.0
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/cast v1.3.1
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/image v0.0.0-20210607152325-775e3b0c77b9 // indirect
	google.golang.org/protobuf v1.26.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
	gorm.io/plugin/dbresolver v1.1.0
)

replace github.com/bsm/redislock => github.com/bsm/redislock v0.5.0

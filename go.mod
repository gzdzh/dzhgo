module github.com/gzdzh/dzhgo

go 1.20

require (
	github.com/go-pay/gopay v1.5.96
	github.com/gogf/gf v1.16.9
	github.com/gogf/gf/contrib/nosql/redis/v2 v2.5.4
	github.com/gogf/gf/v2 v2.5.4
	github.com/gzdzh/dzhgo/contrib/drivers/mysql v0.0.9
	github.com/gzdzh/dzhgo/contrib/files/local v0.0.9
	github.com/gzdzh/dzhgo/contrib/files/oss v0.0.9
	github.com/gzdzh/dzhgo/dzhCore v0.0.9
	github.com/gzdzh/dzhgo/modules/base v0.0.9
	github.com/gzdzh/dzhgo/modules/dict v0.0.9
	github.com/gzdzh/dzhgo/modules/dzhCms v0.0.9
	github.com/gzdzh/dzhgo/modules/dzhCommon v0.0.9
	github.com/gzdzh/dzhgo/modules/dzhMember v0.0.9
	github.com/gzdzh/dzhgo/modules/space v0.0.9
	github.com/gzdzh/dzhgo/modules/task v0.0.9
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/aliyun/aliyun-oss-go-sdk v2.2.9+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/clbanning/mxj v1.8.5-0.20200714211355-ff02cfb8ea28 // indirect
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/gogf/gf/contrib/drivers/mysql/v2 v2.5.4 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/gomodule/redigo v1.8.5 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grokify/html-strip-tags-go v0.0.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/redis/go-redis/v9 v9.2.1 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	go.opentelemetry.io/otel v1.19.0 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	go.opentelemetry.io/otel/sdk v1.19.0 // indirect
	go.opentelemetry.io/otel/trace v1.19.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
	gorm.io/gorm v1.25.4 // indirect
)

replace (
	github.com/gzdzh/dzhgo/modules/dzhCms v0.0.9 => ./modules/dzhCms
	github.com/gzdzh/dzhgo/modules/dzhCommon v0.0.9 => ./modules/dzhCommon
	github.com/gzdzh/dzhgo/modules/dzhMember v0.0.9 => ./modules/dzhMember

)

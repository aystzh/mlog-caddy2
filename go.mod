module github.com/aystzh/mlog-caddy2

go 1.15

require (
	github.com/caddyserver/caddy/v2 v2.4.3
	github.com/eknkc/basex v1.0.1 // indirect
	github.com/juju/errors v0.0.0-20200330140219-3fe23663418f
	github.com/liuzl/store v0.0.0-20190530065605-e2dbcd3c77fc
	github.com/rs/zerolog v1.26.1
	github.com/syndtr/goleveldb v1.0.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	zliu.org/goutil v0.0.0-20210628080224-310b49755b5f
)

replace github.com/liuzl/caddy2-zauth => github.com/aystzh/mlog-caddy2 v0.0.0-20211224062750-3b56276f9024 // indirect
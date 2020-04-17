module github.com/hyperdidact/gotty

go 1.13

require (
	github.com/NYTimes/gziphandler v0.0.0-20170804200234-967539e5e271
	github.com/codegangsta/cli v1.19.1
	github.com/elazarl/go-bindata-assetfs v0.0.0-20150813044622-d5cac425555c
	github.com/fatih/structs v1.1.0
	github.com/gorilla/websocket v1.4.0
	github.com/hyperdidact/structs v0.0.0-20150526064352-a9f7daa9c272
	github.com/kr/pty v1.1.1
	github.com/pkg/errors v0.8.1-0.20161029093637-248dadf4e906
	github.com/spf13/cobra v1.0.0
	github.com/yudai/gotty v0.0.0-00010101000000-000000000000
	github.com/yudai/hcl v0.0.0-20151013225006-5fa2393b3552
)

replace github.com/yudai/gotty => ./

module registry-auth-server

go 1.13

require (
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7
	github.com/gin-gonic/gin v1.4.0
	github.com/jinzhu/gorm v1.9.12
	github.com/spf13/cobra v0.0.5
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43

package cdn

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/cdn"
)

var accessKey = "CZ7UFxAAFj_woqg_4igpwxDYiyBINyUANXJQi-VN"
var secretKey = "F8QDp_QZSu4f9oyfARWnUzyutYD5c0d3x6N8pdf8"

var bucket = "mxdblcf"

func NewCdnManager() *cdn.CdnManager {
	mac := qbox.NewMac(accessKey, secretKey)
	cdnManager := cdn.NewCdnManager(mac)
	return cdnManager
}

package kodo

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"net/http"
)

var accessKey = "CZ7UFxAAFj_woqg_4igpwxDYiyBINyUANXJQi-VN"
var secretKey = "F8QDp_QZSu4f9oyfARWnUzyutYD5c0d3x6N8pdf8"

var bucket = "mxdblcf"
var ImgUrl = "rd1852hee.hd-bkt.clouddn.com" // 测试域名
var filePath = "/home/mxd/nohup.out"

func GinUpload() {
	{
		r := gin.Default()

		//测试ping方法
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		//测试ping方法
		r.POST("/ping1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		//测试文件上传
		r.POST("/upload", func(c *gin.Context) {
			file, err := c.FormFile("file")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 10000,
					"msg":  err.Error(),
				})
				return
			}
			url, code := UpLoadFile(file)
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "ok",
				"url":  url,
			})
		})
		r.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080
	}
}

//客户端上传文件
func UpLoadFile(file *multipart.FileHeader) (string, int) {
	//打开文件
	src, err := file.Open()
	defer src.Close()
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	//设置过期时间
	putPolicy.Expires = 7200
	mac := qbox.NewMac(accessKey, secretKey)
	//拿到上传token
	uploadToken := putPolicy.UploadToken(mac)
	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei, // 华北区
		UseCdnDomains: false,
		UseHTTPS:      false, // 非https
	}

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}        // 上传后返回的结果
	putExtra := storage.PutExtra{} // 额外参数

	// 以默认key方式上传
	err = formUploader.PutWithoutKey(context.Background(), &ret, uploadToken, src, file.Size, &putExtra)

	if err != nil {
		code := 501
		return err.Error(), code
	}

	url := ImgUrl + ret.Key // 返回上传后的文件访问路径

	return url, 0
}
func UpLoad(userFile string, key string) (string, int) {
	//打开文件

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	//设置过期时间
	putPolicy.Expires = 7200
	mac := qbox.NewMac(accessKey, secretKey)
	//拿到上传token
	uploadToken := putPolicy.UploadToken(mac)
	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong, // 华北区
		UseCdnDomains: false,
		UseHTTPS:      false, // 非https
	}

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}        // 上传后返回的结果
	putExtra := storage.PutExtra{} // 额外参数

	// 以默认key方式上传

	err1 := formUploader.PutFile(context.Background(), &ret, uploadToken, key, userFile, &putExtra)
	if err1 != nil {
		println(err1.Error())
	}
	fmt.Println(ret.Key)
	fmt.Println(ret.Hash)
	url := ImgUrl + ret.Key // 返回上传后的文件访问路径

	return url, 0
}

//删除key
func Delete(bucket string, key string) {
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//创建bucket管理器
func CreateBucket(ac, se string) *storage.BucketManager {

	mac := qbox.NewMac(ac, se)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	return bucketManager
}

//获取文件信息
func FileInfo(buc, key string) {

	bucketManager := CreateBucket(accessKey, secretKey)

	stat, err := bucketManager.Stat(buc, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stat.String())
	fmt.Println(storage.ParsePutTime(stat.PutTime))
}

//抓取网络资源到空间
func FetchNet(buc, url string) {
	bucketManager := CreateBucket(accessKey, secretKey)
	fetch, err := bucketManager.Fetch(url, buc, "qiniu.png")
	if err != nil {
		fmt.Println("fetch error,", err)
	} else {
		fmt.Println(fetch.String())
	}
	// 不指定保存的key，默认用文件hash作为文件名
	fetch, err = bucketManager.FetchWithoutKey(url, buc)
	if err != nil {
		fmt.Println("fetch error,", err)
	} else {
		fmt.Println(fetch.String())
	}
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"haotian_main/config"
	"haotian_main/dao"
	"haotian_main/serviceGrpc"
	"haotian_main/utils"
	"net/http"
	"time"
	origin "context"
)


type UserController struct {
}
//新增用户
func (controller *UserController) Add(context *gin.Context) {
	name, exist := context.GetPostForm("name")
	if !exist || name == "" {
		context.JSON(http.StatusOK, gin.H{
			"msg": "请输入用户名:name",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
//查询用户
func (controller *UserController) Get(context *gin.Context) {
	id := context.Query("id")
	currentTime := time.Now().Unix()
	UserKey :=  "haotian"
	err := utils.RedisClient.Set(UserKey, currentTime, 0).Err()
	if err != nil {
		panic(err)
	}
	demo := dao.GetDemoDao().GetDemoDaoName("haotian")
	val, err := utils.RedisClient.Get(UserKey).Result()
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, gin.H{
		"id": id,
		"conf": config.GetConfig(),
		"redis_data":val,
		"db_data":	demo,

	})
}

const (
	address = "localhost:9091"
)

func (controller *UserController) GetGrpc(context *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials())) //建立客户端和服务器之间的链接
	if err != nil {
		fmt.Printf("connect failed %v \n", err)
	}
	defer conn.Close()
	c := serviceGrpc.NewDemoServiceClient(conn)
	ctx, cancel := origin.WithTimeout(origin.Background(), time.Second)
	defer cancel()
	// 远程调用UnaryCall方法
	resp, err := c.UnaryCall(ctx, &serviceGrpc.DemoRequest{Json: "I am bear"})
	if err != nil {
		fmt.Printf("UnaryCall failed:%v", err)
	}
	context.JSON(http.StatusOK, gin.H{
		"data": resp.Message,
	})
}
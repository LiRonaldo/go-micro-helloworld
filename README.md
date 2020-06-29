# go-micro-helloworld
###第一天
####首先 go mod的使用  先创建一个普通的go项目 进入项目执行 go mod init 然后 go mod tidy 就可以了 ，个人理解完全相当于java 的pom 问你件
####代码采用gin 框架，方便web调用，个人理解web 调用是暴露给用户的，go-micro微服务是为了web调用，比如，用户点击查询按钮，调用后台对应页面的那个服务，页面对应的服务调用订单服务。
####使用consuul（也可以默认使用mds，或者k8s 等等） 作为服务的治理平台，目前代码里只将prodService 注册进consul ，并且 代码注释里标记了 使用命令行的方式注册多个服务，目前代码没有将webapi注册进consul，敬请期待！
####有个小坑要注意consul 不在github.com/micro/go-micro/registry下，而是在github.com/micro/go-plugins/registry/ 包括 eureka,redis
###第一天

###第二天 
####采用bat批处理文件的方式结合第一天学的命令行方式一次性注册多个服务
####prod.bat 文件内容如下：
@echo off
start "prod1" go run  prod_main.go --server_address 127.0.0.7:8001 &
start "prod2" go run  prod_main.go --server_address 127.0.0.7:8002 &
start "prod3" go run  prod_main.go --server_address 127.0.0.7:8003
pause
####要注意采用命令行方式注册服务的话 要注掉web.Address("localhost:8001")，并添加server.init()方法，可见prod_main.go 文件
###第二天
# go-micro-helloworld
###第一天
####首先 go mod的使用  先创建一个普通的go项目 进入项目执行 go mod init 然后 go mod tidy 就可以了 ，个人理解完全相当于java 的pom 问你件
####代码采用gin 框架，方便web调用，个人理解web 调用是暴露给用户的，go-micro微服务是为了web调用，比如，用户点击查询按钮，调用后台对应页面的那个服务，页面对应的服务调用订单服务。
####使用consuul（也可以默认使用mds，或者k8s 等等） 作为服务的治理平台，目前代码里只将prodService 注册进consul ，并且 代码注释里标记了 使用命令行的方式注册多个服务，目前代码没有将webapi注册进consul，敬请期待！
###第一天
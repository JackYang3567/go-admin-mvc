
# go-admin-mvc
---
# ChitChat
* httprouter 路由
* X_admin 后台模板布局
* 此项目用`config.json` 文件进行配置
* 程序运行出错结果记录到日志文件 `log.log` 
* 神奇日期格式化,记住一个特殊的日子
   >time.Format("Jan 2, 2006 at 3:04pm")
   >time.Format("2006-01-02 15:04:05")

# 解决报错：unrecognized import path "golang.org/x/image/math/fixed"
```
$mkdir -p $GOPATH/src/golang.org/x/
$cd $GOPATH/src/golang.org/x/
$git clone https://github.com/golang/net.git net 
$go install net
```
天朝可以去 http://www.golangtc.com/download/package 或 https://gopm.io 下载

# 插件
```
go get github.com/gomodule/redigo/redis
go get github.com/streadway/amqp
https://blog.csdn.net/qq_28018283/article/details/84952123

```



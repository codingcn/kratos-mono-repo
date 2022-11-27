# Kratos Project Template - Mono Repo

## 项目介绍


### 已接入组件


| 组件名      | 介绍      | 官网                                |
|----------|---------|-----------------------------------|
| zap      | 日志库     | https://github.com/uber-go/zap    | 
| gorm     | 数据库ORM  | https://gorm.io/                  | 
| go-redis | redis库  | https://github.com/go-redis/redis | 
| consul   | 服务发现与注册 | https://www.consul.io/            | 
| jaeger   | 链路追踪    | https://www.jaegertracing.io/     | 


### Docker compose快速启动项目

```
// 构建微服务镜像
make docker

// 先编排环境依赖，如mysql、redis、consul...
docker compose -f deploy/docker-compose/app/docker-compose.yml up


// 启动所有微服务

docker compose -f deploy/docker-compose/service/docker-compose.yml up


```

现在我们可以尝试我们的微服务是不是启动成功了

```
curl 'http://127.0.0.1:8000/v1/user/info?id=1'
{"code":"0","message":"","data":{"id":"1","username":"用户微服务"}}
```

浏览器访问Consul ui面板

http://localhost:8500/ui/dc1/services

浏览器访问jaeger面板

http://localhost:16686/search






```
// 停止所有容器
docker compose down 
```


### 已实现微服务调用链

该项目预设了三个微服务`bff`、`user`、`order`

* bff 实现BFF模式，所有客户端Http API入口
* user 内部用户服务，仅允许RPC调用
* order 内部订单服务，仅允许RPC调用


流程：浏览器请求-->bff.GetUserInfo-->user.GetUserInfo-->order.Hello



## 新增微服务

```
make service name=yourServiceName
```

执行完上面make脚本命令之后，需要我们自行
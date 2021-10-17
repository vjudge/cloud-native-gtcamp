# 操作步骤

### 编译
```shell
docker build -t go_http_server:v1 .
```


### 改名字
```shell
# 不改名字直接 push 会报错: denied: requested access to the resource is denied
docker tag go_http_server:v1 vjudge/go_http_server:v1
```


### 推送镜像
```shell
docker push vjudge/go_http_server:v1
```


### 服务器拉取镜像
```shell
docker pull vjudge/go_http_server:v1
```


# GO语言微服务实

## 使用 fresh 实现热部署#

- 安装 fresh
``` 
go get github.com/pilu/fresh
```

- 跳转到项目目录,例如项目名为`myapp`
```
cd projectDir
```

- 启动
``` 
fresh
```

## 打包并上传服务
``` 
go build main resources-provider main.go

# 上传应用
scp -r resources-provider root@139.155.239.200:/prod_app/resources-provider

# 上传配置文件
scp -r conf root@139.155.239.200:/prod_app/resources-provider/
```
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

## swagger 文档生产命令
```
swag init -d ./,../../core@v1.0.0 
```

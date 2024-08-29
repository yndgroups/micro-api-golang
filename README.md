# 微服务

## appId
2021092612385500008

## swagger 文档生产命令
- 安装`swag`工具
```
go install github.com/swaggo/swag/cmd/swag@latest

go get -u -v github.com/swaggo/swag/cmd/swag
```

- `生成命令`
```
swag init -d ./,../../core/cmodel,../../core/resp,../../protobuffer/build --exclude ./mapper/,../../core/pager,../../core/validate
```

## swagger访问地址
```
http://ip:8100/swagger-ui/
```

## 使用 fresh 实现热部署#
- 安装 fresh
``` 
go install github.com/pilu/fresh
go get github.com/pilu/fresh
```

## 使用 air 进程热部署
```
go get -v -u github.com/cosmtrek/air

go install github.com/cosmtrek/air

```

- 跳转到项目目录,例如项目名为`myapp`

```
cd myapp
```

- 启动
``` 
fresh
```

## 项目目录
```
    bin     -- 加强工具
    build   -- 构建的应用
    pkg     -- 依赖包
    src     -- 源码
        |-- core@1.0.0    -- 核心模块统一封装
            |--     config        - 配置
            |--     handle        - 处理器
            |--     init_client   - 初始化工具
            |--     jwt           - jwt
            |--     model         - 数据model
            |--     pager         - 分页
            |--     utils         - 工具集
            |--     go.mod        - 包管理器
        |-- common    -- 公共服务
            |--     common-consumer         -- 服务消费者
            |--     common-provider         -- 服务提供者
        |-- content    -- 内容服务
            |--     content-consumer         -- 服务消费者
            |--     content-provider         -- 服务提供者
        |-- gateway    -- 网关服务
            |--     gateway-consumer         -- 服务消费者
            |--     gateway-provider         -- 服务提供者
        |-- order    -- 订单服务
            |--     order-consumer         -- 服务消费者
            |--     order-provider         -- 服务提供者
        |-- resources    -- 静态资源管理服务
            |--     resources-consumer         -- 服务消费者
            |--     resources-provider         -- 服务提供者
        |-- shopping    -- 静态资源管理服务
            |--     shopping-consumer         -- 服务消费者
            |--     shopping-provider         -- 服务提供者
    .gitignore
    README.md
```


## 代理设置

## 环境变量设置
```
# 环境变量
export GOROOT=/home/soft/go
export PATH=$PATH:$GOROOT/bin

# 工作目录
export GOPATH=/home/work/you-project-name
export PATH=$PATH:$GOPATH/bin
```

- File > Settings > Go > Go Modules
  在 Environment 输入如下
``` 
GOPROXY=https://goproxy.cn,direct
```

## windows 下安装包 protoc
- `protoc-gen-go`
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
- `protoc-gen-go`
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
- `protoc-gen-go-grpc`
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- `protoc-go-inject-tag`
```
go install github.com/favadi/protoc-go-inject-tag@latest
```

## 当前构建版本
- protoc-gen-go-grpc v1.3.0
- protoc             v3.5.0

## 网关启动
``` 
nohup java -Dfile.encoding=utf-8 -jar gateway-1.0.1.jar > gateway-1.0.1.log 2>&1 &
```
## 开发者列表
|姓名|联系方式|
| ------ | ------ |
|ydq|826422592@qq.com|

## gorm 标识
```
type User struct {
  gorm.Model // 这里边包含了主键子增id，创建时间，修改时间，删除时间
  Name string `gorm:"<-:create"`                // 只能读和创建
  Name string `gorm:"<-:update"`                // 只能读和修改
  Name string `gorm:"<-"`                       // 读，创建，修改
  Name string `gorm:"<-:false"`                 // 读, 关闭写权限
  Name string `gorm:"->"`                       // 只读
  Name string `gorm:"->;<-:create"`             // 读和创建
  Name string `gorm:"->:false;<-:create"`       // 只能新增
  Name string `gorm:"-"`                        // 创建和读的时候从struct忽略这个字段
  Name string `gorm:"-:all"`                    // 在数据迁移，读和新增的时候都忽略 ignore this field when write, read and migrate with struct
  Name string `gorm:"-:migration"`              // 迁移的时候忽略这个字段 ignore this field when migrate with struct
  
  CreatedAt time.Time                           // 在创建时，如果该字段值为零值，则使用当前时间填充
  UpdatedAt int                                 // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
  Updated   int64 `gorm:"autoUpdateTime:nano"`  // 使用时间戳填纳秒数充更新时间
  Updated   int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
  Created   int64 `gorm:"autoCreateTime"`       // 使用时间戳秒数填充创建时间
  Name string `gorm:"varchar(200) not null"`    // 指定字段数据结构 bool、int、uint、float、string、time、bytes
  Name string `gorm:"size(200)"`                // 指定列大小
  Name string `gorm:"primaryKey"`               // 设置为主键
  Name string `gorm:"unique"`                   // 唯一
  Name string `gorm:"default(123)"`             // 设置默认值
  Name string `gorm:"precision(2)"`             // 精度
  ID int `gorm:"autoIncrement"`                 // 自增长
  
}
```

## 报错提示 undefined 
执行如下命令
```
go mod vendor
```

## Eclipse 安装 go插件
第一步：打开菜单栏“Help”-----"Eclipse Maketplace".
第二部：在弹出框的Find框中输入GoClipse，等待搜索结果出来后结果如下：
第三步：点击“Installed‘’进行安装，过程中会有些提示确认的操作，直接点击确认即可。
第四部：等待Eclipse安装完成。。。。。无需你自己去找插件链接，无需FQ就可以安装成功。
第五步：安装完成重启Eclipse。确认下是否安装成功。点击新建项目，在如下的弹框中会多出一个Go的文件夹，里边会有Go Project可以被你创建，就表示成功了。


```
select file_ids from user where id in ('1' , '2' , '3');
或者
select file_ids from user where find_in_set(id , '1,2,3');
```

## 函数
```
CREATE DEFINER=`root`@`%` FUNCTION `getCategoryTree`(rootId BIGINT) RETURNS varchar(10000) CHARSET utf8mb3
BEGIN
    DECLARE ptemp VARCHAR(10000);
    DECLARE ctemp VARCHAR(10000);
    SET ptemp = '';
    SET ctemp =CAST(rootId AS CHAR);
    WHILE ctemp IS NOT NULL DO
        SET ptemp = CONCAT(ptemp,',',ctemp);
    SELECT GROUP_CONCAT(category_id) INTO ctemp FROM ez_product_category   
    WHERE FIND_IN_SET(parent_id,ctemp) > 0; 
    END WHILE;  
    RETURN ptemp;  
END
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./protos/**/*.proto


```
# Build stage
FROM golang:latest as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
CMD ["./main"]
```

## gopls 配置提示错误
> Invalid settings: gopls setting "expandWorkspaceToModule" is deprecated. Please comment on https://go.dev/issue/63536 if this impacts your workflow.

解决方法： 注释下面的 expandWorkspaceToModule
```
"gopls": {
    // "expandWorkspaceToModule": true,
    "ui.semanticTokens": true,
    "ui.completion.usePlaceholders": true
  },
```

# admin

## 1.项目介绍

​            Hi! 首先感谢你使用 admin，admin是一个基于vue和gin开发的全栈前后端分离的后台管理系统，旨在快速搭建后台管理系统；集成jwt鉴权，动态路由，动态菜单，casbin鉴权，用户管理，部门管理等功能，提供多种示例文件，让您把更多时间专注在业务开发上。

* 代码地址：[admin](https://github.com/liluhao/admin)

注：前端特别鸣谢我的好朋友[田浩云](https://gitee.com/love-out)



![image-20220526132852349](https://mdmdmdmd.oss-cn-beijing.aliyuncs.com/img/image-20220526132852349.png)

## 2.项目预览

[点此进行在线预览](http://www.foolartist.top/dist/#/login?redirect=/dashboard/analysis)

```go
超级管理员用户名：guest
超级管理员密码：123456
```

```go
超级管理员用户名：chris
超级管理员密码：123456
```



## 3.使用说明

### 3.1版本说铭

* 应用版本

```
- node版本 > v14.10
- golang版本 >= v1.14
- IDE推荐：Goland
```

- 使用git克隆本项目

```git
git clone https://github.com/liluhao/admin
```

### 3.2启动项目方法1



####  web端

```bash
# 安装
curl -sL https://deb.nodesource.com/setup_14.x | sudo -E bash -
sudo apt install nodejs

# 验证Node.js和npm是否成功
node --version
npm --version

# 进入项目web端
cd web

# 安装依赖
npm i

# 启动本地测试服务器
npm  run serve
```

#### server端

```bash
# 进入项目目录，打开server目录，不可以打开 admin 根目录
cd server

# 启动 Go Modules function
go env -w GO111MODULE=on 
# 如果您在中国大陆，请配置 GOPROXY environment variables
go env -w GOPROXY=https://goproxy.io,direct
# 使用 go.mod，安装go依赖包
go mod tidy

# 运行
go run main.go
```

### 3.3启动项目方法2

* 进入项目目录

```bash
cd admin
```

- 给一键启动脚本添加权限，并运行一键启动脚本。

```bash
# 添加权限
sudo chmod a+x up.sh
# 执行启动脚本
./up.sh
```

> up.sh 包含了项目启动过程所有配置的运行脚本，具体逻辑根据需求可以自行注释；您还可以分开进入web端和server分别构建并配置启动项目，具体操作如启动项目方法2。



### 3.4初始配置使用指南

- 配置文件路径选择在[config.yaml](./server/conf/config.yaml)
- 导入sql初始化文件（默认mysql）[init.sql](./server/data/init.sql)

### 3.5swagger自动化API文档

#### 安装 swagger

````go
go get -u github.com/swaggo/swag/cmd/swag
````

#### 生成API文档

执行下面的命令后，server目录下会出现docs文件夹，登录 http://localhost/swagger/index.html ，即可查看swagger文档

````go
cd server
swag init
````

## 4. 技术选型

- 前端：用基于`vue`的`vben`构建基础页面。
- 后端：用`Gin`快速搭建基础restful风格API，`Gin`是一个go语言编写的Web框架。
- 数据库：采用`MySql`(5.7)版本，使用`gorm`实现对数据库的基本操作
- 缓存：使用`Redis`实现记录当前活跃用户的
- 权限控制：使用`jwt`令牌并实现多点登录限制。
- API文档：使用`Swagger`构建自动化文档。
- 配置文件：使用`viper`实现解析`yaml`格式的配置文件。
- 日志：使用`zap`实现日志记录。


## 5.项目架构

### 5.1 系统架构图



![image-20220526133034517](https://mdmdmdmd.oss-cn-beijing.aliyuncs.com/img/image-20220526133034517.png)5.2 

### 5.2目录结构

```go
├── server           (服务端目录)
│   ├── app          (应用目录，后台api)
│   ├── common       (公共常量，变量，util方法目录)
│   ├── conf         (配置目录)
│   ├── dao          (Data Access Object  数据访问对象)
│   ├── data         (全局数据目录)
│   ├── docs         (swagger生成目录)
│   ├── dto          (Data Transfer Object 数据传输对象 )
│   ├── initialize   (初始化各个组件目录)
│   ├── log          (日志生成目录)
│   ├── middleware   (中间件目录)
│   ├── resource     (casbin资源目录)
│   ├── routers      (路由配置目录)
│   └── service      (服务层目录)
└── web              (前端目录，采用vue-vben-admin搭建)
    ├── build        (vite目录)
    ├── dist         (构建完成目录)
    ├── mock         (模拟数据目录)
    ├── public       (公共资源目录)
    ├── src          (项目源码目录)
    ├── test
    └── types


```

## 6. 主要功能

- 权限管理：基于`jwt`和`casbin`实现的权限管理 
- 部门管理：配置系统用户所属担任职务。
- 用户管理：系统管理员分配用户角色和角色权限。
- 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。
- 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单，同时融合api管理，
  不同用户可调用的api接口的权限不同。 
- 日志管理：正常操作日志记录和查询；系统异常信息日志记录和查询，登录日志等。

## 7. 联系方式

vx：x15516535379










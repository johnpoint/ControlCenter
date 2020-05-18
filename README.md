# ControlCenter-Server

## 项目仓库列表

- [ControlCenter-Server 主控](https://github.com/johnpoint/ControlCenter-Server)
- [ControlCenter-Client 被控端](https://github.com/johnpoint/ControlCenter-Client)
- [ControlCenter-Web 前端](https://github.com/johnpoint/ControlCenter-Web)

## 当前版本

ver 1.9.0

## 文件

``` 
src/
    Client - 子客户端
    Server - 母服务端
```

## API 接口

``` 
POST    /user/auth/login
POST    /user/auth/register

POST    /server/setup/:token
POST    /update/:token POST
GET     /server/update/:token
GET     /server/Certificate/:token/:id

GET     /

POST    /system/restart

- - -

auth 2.0 request

- - -

POST    /web/debug/check
GET     /web/ServerInfo
PUT     /web/ServerInfo
GET     /web/ServerInfo/Certificate
DELETE  /web/Server/:ip
GET     /web/DomainInfo
PUT     /web/DomainInfo
PUT     /web/UserInfo/:mail/:key/:value
PUT     /web/SiteInfo
GET     /web/SiteInfo
DELETE  /web/SiteInfo
PUT     /web/Certificate
GET     /web/Certificate
POST    /web/Certificate
DELETE  /web/Certificate
PUT     /web/link/Certificate/:ServerID/:CerID
DELETE  /web/link/Certificate/:ServerID/:CerID
PUT     /web/link/Site/:ServerID/:SiteID
DELETE  /web/link/Site/:ServerID/:SiteID
POST    /web/backup
GET     /web/:mail/:pass/backup
PUT     /web/UserInfo/:uid/:level
GET     /web/UserInfo/Password/:oldpass/:newpass
GET     /web/UserInfo/Token
PUT     /web/UserInfo/Token
GET     /web/UserInfo
GET     /web/UserInfo/
```

## Install 安装

### Client

``` 
./Client install [ 后端地址 ] [ hostname 或 自定义名称 ] [ IPv4地址 ] [ IPv6地址，没有就填IPv4 ] [ user_token ]
```

## 设置

``` 
TELEGRAM_BOT_TOKEN
TELEGRAM_CHAT_ID
TELEGRAM_NOTIFICATION "true"
```

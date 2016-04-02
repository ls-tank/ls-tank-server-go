# ls-tank-server-go
> The server of ls-tank in Golang

| Key | Type |
| :--: | :--: |
| name | string |
| password | string |
| diamond | uint |
| tankhead | uint |
| tankBody | uint|
| tankWheel | uint |
| kill | uint |
| dead | uint |

## 注册
```
POST /api/register
```

## 登录
```
POST /api/login
```

## 获取用户所有信息
```
GET /api/user/{id}
```

## 更新用户(自己)信息
```
PATCH /api/user
```

## Nginx
```
server {
  listen 80;
  server_name dev.tank-controller.com;

  location / {
    proxy_pass http://127.0.0.1:7456;
  }

  location /api {
    rewrite ^.+api/?(.*)$ /$1 break;
    proxy_pass http://127.0.01:8123;
  }
}
```

## License
MIT




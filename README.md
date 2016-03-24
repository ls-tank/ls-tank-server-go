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

## License
MIT




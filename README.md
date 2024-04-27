## 快速开始
在启动之前，请确保填写好config.json配置文件：
```json
{
    "authConfig": {
        "Username": "admin",
        "Password": "123456",
        "Enable2FA": false
    },
    "databaseConfig": {
        "Driver": "mysql",
        "Host": "127.0.0.1",
        "Port": "3306",
        "Database": "tomato",
        "Username": "root",
        "Password": "msql",
        "Charset": "utf8mb4"
    },
    "user": {
        "Username": "酷酷番茄",
        "Gravatar": "/",
        "Url": "coolcooltomato.com",
        "Description": "(⊙o⊙)？"
    },
    "site": {
        "Path": "http://localhost:7891/",
        "Title": "这里是酷酷番茄的博客",
        "Icon": "https://coolcooltomato.com/file/tomato.ico",
        "Bili": "https://space.bilibili.com/1234567",
        "Github": "https://github.com/CoolCoolTomato",
        "Twitter": "/",
        "Mail": "hello@outlook.com"
    }
}
```
第一次启动时，将*Enable2FA*设置为*false*，在登录后台查看二次验证的密钥后，根据需求改为true。
## 效果展示
![](https://github.com/CoolCoolTomato/WhiteBlog/assets/107784402/567c0422-052a-427f-b73f-6c36866eaace)
![](https://github.com/CoolCoolTomato/WhiteBlog/assets/107784402/429594fc-af6d-4ea3-8389-8c6b5fa1e521)

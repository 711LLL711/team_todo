# TeamToDo API 文档

### 公共信息

数据交互格式：

```json
{
    "success": true,
    "hint": "",
    "data": {}
}
```

身份验证方式：

HTTP 头 Authorization 设置为 `Bearer JWT令牌`

URL 前缀：

`http(s)://<Host>/api/v1`

### 用户模块

---

#### 用户注册

请求

```
POST /users/register
```

参数

| 参数     | 类型   | 描述     |
| -------- | ------ | -------- |
| email    | string | 用户邮箱 |
| password | string | 用户密码 |
| nickname | string | 用户昵称 |

响应

```
{
  "id": "12345",
  "email": "example@qq.com",
  "nickname": "Li Hua"
}
```

#### 用户登录

请求

```
POST /users/login
```

参数

| 参数     | 类型   | 描述     |
| -------- | ------ | -------- |
| email    | string | 用户邮箱 |
| password | string | 用户密码 |

响应

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjM0NTY3ODkwIiwiaWF0IjoxNTE2MjM5MDIyfQ.7PPi5QfLs5KCb9gU80Rb1L1XnH4guxyB1ylYJU9dxS4"，
  "expire": 123456789
}
```

#### 用户密码找回

**发送邮箱验证码**

请求

```
GET /users/verify-code
```

参数

| 参数  | 类型   | 描述     |
| ----- | ------ | -------- |
| email | string | 用户邮箱 |

**修改密码**

请求

```
GET /users/verify-code
```

参数

| 参数     | 类型   | 描述     |
| -------- | ------ | -------- |
| email    | string | 用户邮箱 |
| password | string | 用户密码 |
| code     | string | 验证码   |

#### 获取用户资料

请求

```
GET /users/{id}/profile
```

参数

| 参数 | 类型   | 描述   |
| ---- | ------ | ------ |
| id   | string | 用户id |

响应

```
{
  "id": "12345",
  "email": "user@example.com",
  "nickname": "John Doe",
  "avatar": "https://example.com/avatar.jpg"
}
```

#### 获取个人资料

请求

```
GET /users/profile
```

响应

```
{
  "id": "12345",
  "email": "user@example.com",
  "nickname": "John Doe",
  "avatar": "https://example.com/avatar.jpg"
}
```

#### 更新个人资料

请求

```
PUT /users/profile
```

参数

| 参数     | 类型   | 描述     |
| -------- | ------ | -------- |
| nickname | string | 用户昵称 |
| avatar   | string | 头像链接 |

#### 上传头像图片

请求

```
POST /avatars
```

参数：头像图片文件

响应

```
{
  "url": "https://example.com/avatar.jpg"
}
```

### 群组模块

#### 创建群组

请求

```
POST /groups
```

参数

| 参数        | 类型   | 描述     |
| ----------- | ------ | -------- |
| name        | string | 群组名称 |
| description | string | 群组描述 |

响应

```
{
  "id": "123",
  "name": "Group A",
  "description": "This is a sample group",
  "owner": 123
}
```

#### 获取已加入群组列表

请求

```
GET /groups
```

响应

```
{
  "count": 5,
  "groups": [
    {
      "id": "123",
      "name": "Group A",
      "description": "This is a sample group",
    },
    ...
  ]
}
```

#### 获取群组信息

请求

```
GET /groups/{id}/info
```

参数

| 参数 | 类型   | 描述   |
| ---- | ------ | ------ |
| id   | string | 群组ID |

响应

```
{
  "id": "12345",
  "name": "Group A",
  "description": "This is Group A",
  "owner": "123"
}
```

#### 获取群组成员

请求

```
GET /groups/{id}/members
```

参数

| 参数 | 类型   | 描述   |
| ---- | ------ | ------ |
| id   | string | 群组ID |

响应

```
{
  "count": 5,
  "members": [
    {
      "id": "12345",
      "email": "user@example.com",
      "nickname": "John Doe",
      "avatar": "https://example.com/avatar.jpg"
    },
    ...
  ]
}
```

#### 加入群组

请求

```
GET /groups/join
```

参数

| 参数 | 类型   | 描述       |
| ---- | ------ | ---------- |
| code | string | 群组邀请码 |

响应

```
{
  "id": "12345",
  "name": "Group A",
  "description": "This is Group A"
}
```

#### 获取群组邀请码

请求

```
GET /groups/{id}/code
```

参数

| 参数 | 类型   | 描述   |
| ---- | ------ | ------ |
| id   | string | 群组ID |

响应

```
{
  "code": "asdfghj"
}
```

#### 退出群组

请求

```
POST /groups/{id}/leave
```

参数

| 参数 | 类型   | 描述   |
| ---- | ------ | ------ |
| id   | string | 群组ID |

#### 解散群组

请求

```
DELETE /groups/{id}
```

参数

| 参数 | 类型   | 描述   |
| ---- | ------ | ------ |
| id   | string | 群组ID |

### 任务模块

#### 创建任务

请求

```
POST /tasks
```

参数

| 参数        | 类型   | 描述               |
| ----------- | ------ | ------------------ |
| name        | string | 任务名称           |
| description | string | 任务描述           |
| status      | string | 任务状态           |
| assignee    | string | 负责人ID           |
| deadline    | int    | 截止日期（时间戳） |
| groupId     | string | 群组ID             |

响应

```
{
  "id": "67890",
  "name": "Task A",
  "description": "This is Task A",
  "status": "In Progress",
  "assignee": "12345",
  "deadline": 1689124831,
  "groupId": "123"
}
```

#### 获取任务列表

请求

```
GET /tasks
```

参数

| 参数    | 类型   | 描述   |
| ------- | ------ | ------ |
| groupId | string | 群组ID |

响应

```
{
  "count": 5,
  "tasks": [
    {
      "id": "67890",
      "name": "Task A",
      "description": "This is Task A",
      "status": "In Progress",
      "assignee": "12345",
      "deadline": 1689124831,
      "groupId": "123"
    },
    ...
  ]
}
```

#### 获取任务信息

请求

```
GET /tasks/{id}
```

参数

| 参数 | 类型   | 描述   |
| ---- | ------ | ------ |
| id   | string | 任务ID |

响应

```
{
  "id": "67890",
  "name": "Task A",
  "description": "This is Task A",
  "status": "In Progress",
  "assignee": "12345",
  "deadline": 1689124831,
  "groupId": "123"
}
```

#### 修改任务

请求

```
PUT /tasks/{id}
```

参数

| 参数        | 类型   | 描述               |
| ----------- | ------ | ------------------ |
| id          | string | 任务ID             |
| name        | string | 任务名称           |
| description | string | 任务描述           |
| status      | string | 任务状态           |
| assignee    | string | 负责人ID           |
| deadline    | int    | 截止日期（时间戳） |

响应

```
{
  "id": "67890",
  "name": "Task A",
  "description": "This is Task A",
  "status": "In Progress",
  "assignee": "12345",
  "deadline": 1689124831,
  "groupId": "123"
}
```

#### 删除任务

请求

```
DELETE /tasks/{id}
```

参数

| 参数 | 类型   | 描述   |
| ---- | ------ | ------ |
| id   | string | 任务ID |

### 提醒模块

#### 创建提醒

请求

```
POST /reminders
```

参数

| 参数   | 类型   | 描述               |
| ------ | ------ | ------------------ |
| taskId | string | 任务ID             |
| time   | int    | 提醒时间（时间戳） |
| method | string | 提醒发送方式       |

响应

```
{
  "id": "2001",
  "taskId": "1001",
  "time": 1689124831,
  "method": "email"
}
```

#### 获取提醒列表

请求

```
GET /reminders
```

参数

| 参数   | 类型   | 描述   |
| ------ | ------ | ------ |
| taskId | string | 任务ID |

响应

```
{
  "count": 5,
  "reminders": [
    {
      "id": "2001",
      "taskId": "1001",
      "time": 1689124831,
      "method": "email"
    },
    ...
  ]
}
```

#### 删除提醒

请求

```
DELETE /reminders/{id}
```

参数

| 参数 | 类型   | 描述   |
| ---- | ------ | ------ |
| id   | string | 提醒ID |

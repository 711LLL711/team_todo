# team_todo

## 前端页面

- 注册页面
    用户名+密码+邮箱+验证码      
    邮箱和验证码返回后端      
- 登录页面     
    用户名+密码    
- 找回密码页面   
    邮箱+验证码   
    重置密码     
- 群组    
    任务列表    
    成员列表       
    任务提醒     
    任务详情    


## api文档 

[api文档](api.md)    

## 数据库设计  

1. 用户表User     
    - Id
    - 用户名Nickname  
    - 密码(经过加密)Password   
    - 邮箱Email      
    - 头像Avatar    
2. 群组表Group    
    - 群组名GroupName        
    - 群组id GroupId          
    - 群主GroupOwnerId
    - 群描述Group_Description    
    - 邀请码Group_Invite_Id   
3. 群组成员对应表groupwithuser      
    - UserId    
	- GroupId     
4. 任务表Task     
    - 任务id ID     
    - 任务名 Name       
    - 任务描述Description    
    - 任务负责人AssigneeId    
    - ddl DueDate        
    - 任务状态TODO/DONE  
    - 提醒时间Status    
5. 提醒Reminder     
    1. ReminderId      
    2. TaskId     
    3. DueDate    

package reminder

import (
	"fmt"
	"team_todo/database"
	"time"

	"github.com/robfig/cron/v3"
)

// reminder进程
func Reminder() {
	// 创建 cron 实例
	c := cron.New()

	_, err := c.AddFunc("0 8 * * *", func() {
		// 在此处编写查询数据库的逻辑
		// 执行查询操作，获取最新的数据
		database.RemindNow()
		// 处理查询结果，执行相应的逻辑
		//调用MyReminder函数
		MyReminder()
		fmt.Println("执行数据库查询：", time.Now())
	})
	if err != nil {
		panic(err)
	}

	// 启动定时任务
	c.Start()

	// 保持主程序运行，直到手动停止或发生错误
	select {}
}

package reminder

import (
	"fmt"
	"sync"
	"team_todo/model"
)

func MyReminder(reminders []model.Reminder) {
	// 假设你从数据库查询得到的需要提醒的任务列表
	// tasks := []Task{
	// 	{ID: 1, Name: "Task 1"},
	// 	{ID: 2, Name: "Task 2"},
	// 	{ID: 3, Name: "Task 3"},
	// }

	// 创建一个等待组，用于等待所有 Goroutine 完成
	var wg sync.WaitGroup

	// 创建一个通道用于接收提醒完成的信号
	remindersDone := make(chan struct{})

	// 启动一个 Goroutine 处理每个提醒请求
	for _, reminder := range reminders {
		wg.Add(1)

		go func(reminder model.Reminder) {
			defer wg.Done()

			// 执行提醒逻辑，发送提醒通知
			Conduct(reminder)
			fmt.Printf("触发提醒：Reminder ID %d, To %s\n", reminder.ReminderId, reminder.Email)
		}(reminder)
	}

	// 启动一个 Goroutine 等待所有提醒请求完成
	go func() {
		wg.Wait()
		close(remindersDone)
	}()

	// 主程序继续执行其他操作...

	// 等待所有提醒请求完成
	<-remindersDone

	fmt.Println("所有提醒请求已完成")
}

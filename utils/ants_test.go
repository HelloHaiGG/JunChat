package utils

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"testing"
	"time"
)

//ants 高效率的协程池

//不需要参数的任务
func TaskNoArgs() {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task No Args Finish.")
}

func Task() {
	TaskNoArgs()
	wg.Done()
}

//需要参数的任务
func TaskWithArgs(wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task Whit Args Finish.")
	wg.Done()
}

var wg *sync.WaitGroup

func TestNoArgs(t *testing.T) {
	//通过 ants.Submit() 提交任务,缺点是无法对任务添加参数
	defer ants.Release()
	_ = ants.Submit(TaskNoArgs)

	//等待多个任务执行完后退出需要在任务外层套一层func
	times := 1000
	wg = &sync.WaitGroup{}
	for i := 0; i < times; i++ {
		wg.Add(1)
		_ = ants.Submit(Task)
	}
	wg.Wait()
}

func TestWithArgs(t *testing.T) {
	times := 1000
	wg = &sync.WaitGroup{}

	//NewPoolWithFunc 只能执行特定任务
	pool, _ := ants.NewPoolWithFunc(100, func(i interface{}) {
		w := i.(*sync.WaitGroup)
		TaskWithArgs(w)
	})
	for i := 0; i < times; i++ {
		wg.Add(1)
		_ = pool.Invoke(wg)
	}
	wg.Wait()
	defer pool.Release()
}

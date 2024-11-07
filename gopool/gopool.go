package gopool

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	f func(interface{})
}

func NewWorker(f func(interface{})) *Worker {
	return &Worker{
		f: f,
	}
}
func (w *Worker) Run(job interface{}) {
	w.f(job)
}

type Pool struct {
	//模板函数
	Function func(interface{})
	// 任务队列
	JobQueue chan interface{}
	//正在执行的任务清单
	JobsList *sync.Map
	//完成任务清单
	FinishList *sync.Map
	//池子大小
	threads int
	//启动协程等待时间
	Interval time.Duration
	//用于阻塞
	wg *sync.WaitGroup
	//提前结束标识
	Done bool
}

func NewPool(threads int) *Pool {
	return &Pool{
		JobQueue:   make(chan interface{}, 100),
		JobsList:   &sync.Map{},
		FinishList: &sync.Map{},
		Done:       false,
		Interval:   time.Duration(0),
		threads:    threads,
		wg:         &sync.WaitGroup{},
	}
}

// 添加任务
func (p *Pool) AddJob(job interface{}) {
	if p.Done {
		return
	}
	p.JobQueue <- job
}

// 结束池子,会等待所有任务执行完毕
func (p *Pool) Stop() {
	if !p.Done {
		close(p.JobQueue)
	}
}

// 启动池子
func (p *Pool) Run() {
	p.Done = false
	//只启动有限大小的协程，协程的数量不可以超过工作池设定的数量，防止计算资源崩溃
	for i := 0; i < p.threads; i++ {
		p.wg.Add(1)
		time.Sleep(p.Interval)
		go p.work()
		if p.Done {
			break
		}
	}
	p.wg.Wait()
	p.Done = true
}

// 工人
func (p *Pool) work() {
	defer p.wg.Done()
	for job := range p.JobQueue {
		if p.Done {
			break
		}
		jobID := p.genJobID() // 生成唯一ID
		p.JobsList.Store(jobID, job)

		worker := NewWorker(p.Function)
		worker.Run(job)

		p.FinishList.Store(jobID, job)
		p.JobsList.Delete(jobID)
	}
}

// 是否结束
func (p *Pool) IsDone() bool {
	return p.Done
}

// 提前结束
func (p *Pool) StopNow() {
	if !p.Done {
		close(p.JobQueue)
	}
	p.Done = true
}

// 生成job唯一标识
func (p *Pool) genJobID() string {
	timestamp := time.Now().UnixNano()
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%d-%s", timestamp, hex.EncodeToString(randomBytes))
}

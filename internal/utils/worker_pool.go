package utils

import "sync"

type Task func()

type WorkerPool struct {
	taskQueue chan Task
	wg        sync.WaitGroup
}

func NewWorkerPool(workerCount, queueSize int) *WorkerPool {
	pool := &WorkerPool{
		taskQueue: make(chan Task, queueSize),
	}

	pool.wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go pool.worker()
	}

	return pool
}

func (p *WorkerPool) worker() {
	defer p.wg.Done()

	for task := range p.taskQueue {
		task()
	}
}

func (p *WorkerPool) Submit(task Task) bool {
	select {
	case p.taskQueue <- task:
		return true
	default:
		return false
	}
}

func (p *WorkerPool) Stop() {
	close(p.taskQueue)
	p.wg.Wait()
}

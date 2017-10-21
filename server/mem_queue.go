package server

import (
	"sync"
)

type QueueMem struct {
	Size     int
	Lock     *sync.RWMutex
	QueueMap map[string](chan *string)
}

func NewQueue(size int) *QueueMem {
	return &QueueMem{
		QueueMap: map[string](chan *string){},
		Lock:     new(sync.RWMutex),
		Size:     size,
	}
}

func (q *QueueMem) InitTopic(topic string) chan *string {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	v, ok := q.QueueMap[topic]
	if !ok {
		q.QueueMap[topic] = make(chan *string, q.Size)
		v = q.QueueMap[topic]
	}
	return v
}

func (q *QueueMem) Pub(topic string, body *string) bool {
	c := q.InitTopic(topic)
	c <- body
	return true
}

func (q *QueueMem) Sub(topic string) *string {
	c := q.InitTopic(topic)
	return <-c
}

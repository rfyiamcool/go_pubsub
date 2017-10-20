package server

import (
	"sync"
)

type Queue struct {
	QueueMap map[string](chan *string)
	Lock     *sync.RWMutex
	Size     int
}

func NewTopic(size int) *Queue {
	return &Queue{
		QueueMap: map[string](chan *string){},
		Lock:     new(sync.RWMutex),
		Size:     size,
	}
}

func (q *Queue) InitTopic(topic string) chan *string {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	v, ok := q.QueueMap[topic]
	if !ok {
		q.QueueMap[topic] = make(chan *string, q.Size)
		v = q.QueueMap[topic]
	}
	return v
}

func (q *Queue) Pub(topic string, body *string) bool {
	c := q.InitTopic(topic)
	c <- body
	return true
}

func (q *Queue) Sub(topic string) *string {
	c := q.InitTopic(topic)
	return <-c
}

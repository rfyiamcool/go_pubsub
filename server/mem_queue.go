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

type TopicPool struct {
	PoolLock  *sync.RWMutex
	Size      int
	TopicMap  map[string]*QueueMem
	QueueInfo map[string]string
}

func NewTopicPool(size int) *TopicPool {
	return &TopicPool{
		PoolLock:  new(sync.RWMutex),
		TopicMap:  make(map[string]*QueueMem),
		QueueInfo: make(map[string]string),
		Size:      size,
	}
}

func (t *TopicPool) CreateTopic(topic string) {
	t.PoolLock.RLock()
	defer t.PoolLock.RUnlock()

	if _, ok := t.TopicMap[topic]; !ok {
		t.TopicMap[topic] = NewQueue(t.Size)
	}
}

func (t *TopicPool) Bind(topic string, qname string) {
	t.PoolLock.RLock()
	defer t.PoolLock.RUnlock()

	if _, ok := t.TopicMap[topic]; !ok {
		t.TopicMap[topic] = NewQueue(t.Size)
	}
}

func (t *TopicPool) InitQueue(queue *QueueMem, qname string) chan *string {
	queue.Lock.RLock()
	defer queue.Lock.RUnlock()

	v, ok := queue.QueueMap[qname]
	if !ok {
		queue.QueueMap[qname] = make(chan *string, t.Size)
		v = queue.QueueMap[qname]
	}
	return v
}

func (t *TopicPool) Pub(topic string, body *string) bool {
	qmap, ok := t.TopicMap[topic]
	if !ok {
		return false
	}

	for _, v := range qmap.QueueMap {
		v <- body
	}

	return true
}

func (t *TopicPool) Sub(topic string, qname string) *string {
	var body *string

	qmap, ok := t.TopicMap[topic]
	if !ok {
		return body
	}

	c := t.InitQueue(qmap, qname)
	return <-c
}

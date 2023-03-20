package queue

type QueueAbstract struct {
}

func QueueBuilder() IQueue {
	return &QueueAbstract{}
}

func (q *QueueAbstract) Add(queueName string, value interface{}) bool {
	return true
}

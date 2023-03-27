package queue

type IQueue interface {
	Add(queueName string, value interface{}) bool
}

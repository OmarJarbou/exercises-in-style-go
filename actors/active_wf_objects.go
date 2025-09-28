package main

type Dispatcher interface {
	dispatch([]interface{})
}

type ActiveWFObjects struct {
	name       string
	stop       bool
	queue      chan []interface{}
	dispatcher Dispatcher
}

func (ao *ActiveWFObjects) run() {
	for !ao.stop {
		message := <-ao.queue
		if message[0] == "die" {
			ao.stop = true
		}
		ao.dispatcher.dispatch(message)
	}
}

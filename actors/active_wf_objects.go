package main

type Runner interface {
	run()
}

type ActiveWFObjects struct {
	name  string
	stop  bool
	Queue chan []interface{}
}

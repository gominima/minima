package auth
import "github.com/gominima/minima"

func SimpleTests() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		res.Send("Hi test 0")
	}
}


func SimpleTests1() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		res.Send("Hi test 1")
	}
}

func SimpleTests2() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		res.Send("Hi test 2")
	}
}

func SimpleTests3() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		res.Send("Hi test 3")
	}
}

func SimpleTests4() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		res.Send("Hi test 4")
	}
}
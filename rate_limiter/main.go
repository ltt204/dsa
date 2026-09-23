package main

import "fmt"

type Config struct {
	Window     int
	MaxRequest int
}

type Request struct {
	UserID int
	Time   int
}

func Server(reqs []Request, config Config) {
	fmt.Printf("Received Requests %v \nConfig: %v\n\n", reqs, config)
	userQueue := make(map[int][]Request)

	var rateLimiter func(req Request) bool
	rateLimiter = func(req Request) bool {
		q := userQueue[req.UserID]

		if len(q) >= config.MaxRequest {
			top := q[0]
			if top.Time <= (req.Time - config.Window) {
				q = q[1:]
				q = append(q, req)
				return true
			} else {
				return false
			}
		}

		q = append(q, req)

		userQueue[req.UserID] = q
		return true
	}

	// Perform rate limiter
	result := make(map[int][]bool)
	for _, req := range reqs {
		reqResult := rateLimiter(req)
		result[req.UserID] = append(result[req.UserID], reqResult)
	}

	for key, val := range result {
		fmt.Printf("User Id: %d with Resposnes: %v\n", key, val)
	}
}

func main() {
	config := Config{
		Window:     10,
		MaxRequest: 3,
	}
	requests := []Request{
		{
			UserID: 1,
			Time:   2,
		},
		{
			UserID: 1,
			Time:   3,
		},
		{
			UserID: 1,
			Time:   4,
		},

		{
			UserID: 1,
			Time:   5,
		},
		{
			UserID: 1,
			Time:   10,
		},

		{
			UserID: 1,
			Time:   12,
		},
		{
			UserID: 2,
			Time:   21,
		},
	}

	Server(requests, config)
}

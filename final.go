package main

import (
	"encoding/json"
	"fmt"
)

const ema string = "email"
const sm string = "sms"
const cal string = "call"
const all string = "all"

var jsonStr string = `[{"name":"Sean","email":"hrstaa@gmail.com","mobile":"9970972420","mode":"email"},{"name":"Sean","email":"hrstaa@gmail.com","mobile":"9970972420","mode":"email"},{"name":"Sean","email":"hrstaa@gmail.com","mobile":"9970972420","mode":"email"},{"name":"Sean","email":"hrstaa@gmail.com","mobile":"9970972420","mode":"email"},{"name":"Sean","email":"hrstaa@gmail.com","mobile":"9970972420","mode":"email"},{"name":"Sean","email":"hrstaa@gmail.com","mobile":"9970972420","mode":"email"},{"name":"Sean","email":"hrstaa@gmail.com","mobile":"9970972420","mode":"email"},{"name":"Sean","email":"hrstaa@gmail.com","mobile":"9970972420","mode":"email"},{"name":"Sean","email":"hrstaa@gmail.com","mobile":"9970972420","mode":"email"}]`

func worker(id int, jobs <-chan user, results chan<- user) {
	for j := range jobs {

		c := call{mobile: j.Mobile}
		e := email{email: j.Email}
		s := sms{mobile: j.Mobile}

		fmt.Println(j)
		if j.Mode == ema {
			message(e)
		} else if j.Mode == cal {
			message(c)
		} else if j.Mode == sm {
			message(s)
		} else {
			message(e)
			message(c)
			message(s)
		}

		results <- j
	}
}

func main() {

	//	fmt.Println(jsonStr)
	var arr []user
	_ = json.Unmarshal([]byte(jsonStr), &arr)
	//fmt.Println(arr)
	const numJobs = 1
	jobs := make(chan user, numJobs)
	results := make(chan user, numJobs)

	for w := 1; w <= 10; w++ {
		go worker(w, jobs, results)
	}

	// for _, j := range arr {
	// 	jobs <- j
	// }

	fmt.Println(len(arr))
	for j := 0; j < len(arr); j++ {
		jobs <- arr[j]
	}

	close(jobs)

	for a := 1; a <= len(arr); a++ {
		<-results
	}
}

type user struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	Mode   string `json:"mode"`
}

type notifications interface {
	sendMsg()
}

type call struct {
	mobile string
}

func (r call) sendMsg() {
	fmt.Println(" notifications by calling : ", r.mobile)
}

type email struct {
	email string
}

func (r email) sendMsg() {
	fmt.Println(" notifications by email", r.email)
}

type sms struct {
	mobile string
}

func (r sms) sendMsg() {
	fmt.Println(" notifications by sms : ", r.mobile)
}

func message(n notifications) {
	n.sendMsg()

}

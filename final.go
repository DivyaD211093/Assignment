package main

// importing the packages
import (
	"encoding/json"
	"fmt"
)

//creating the constants

const ema string = "email"
const sm string = "sms"
const cal string = "call"
const all string = "all"

//have hardcoded the json containing name, email, mobile and mode of notification

var jsonStr string = `[{"name":"Ray","email":"Ray@gmail.com","mobile":"8208943950","mode":"email"},{"name":"Mehul","email":"Mehul@gmail.com","mobile":"9970972420","mode":"sms"},{"name":"Bob","email":"Bob@gmail.com","mobile":"9970972421","mode":"call"},{"name":"abc","email":"abc@gmail.com","mobile":"9970972424","mode":"email"},{"name":"sana","email":"sana@gmail.com","mobile":"9970972420","mode":"call"},{"name":"Sean","email":"Sean@gmail.com","mobile":"8208943951","mode":"sms"},{"name":"shine","email":"shine@gmail.com","mobile":"99709445420","mode":"email"},{"name":"Ola","email":"Ola.com","mobile":"9970972420","mode":"sms"},{"name":"Jack","email":"jack@gmail.com","mobile":"9878972420","mode":"call"},{"name":"Meghan","email":"Meghan@gmail.com","mobile":"9878472420","mode":"sms"}]`

// have created a worker pool to program the mode of notification

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
	// here we have taken a json and converted it into a struct, also have used channels here and have defined a rate limit to 10
	var arr []user
	err := json.Unmarshal([]byte(jsonStr), &arr)
	if err != nil {
		panic("Json Parsing: Please pass appropriate json")
	}
	const numJobs = 5
	jobs := make(chan user, numJobs)
	results := make(chan user, numJobs)

	for w := 1; w <= 10; w++ {
		go worker(w, jobs, results)
	}

	for _, j := range arr {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= len(arr); a++ {
		<-results
	}
}

//declared the struct and interface

type user struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	Mode   string `json:"mode"`
}

// creating an interface and supporting functions
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

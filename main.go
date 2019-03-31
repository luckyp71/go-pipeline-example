package main

import (
	p "go-pipeline-example/example1"
	"log"
)

func main() {
	input := p.Gen("lucky@yahoo.com", "luckypratama@yahoo.com", "luckypratama71@yahoo.com", "lucky@gmail.com", "luckypratama@gmail.com",
		"luckypratama712@gmail.com", "pratama@yahoo.com", "pratama@gmail.com")

	for email := range p.SendEmail(input) {
		log.Println(email)
	}
}

package main

import (
	"fmt"
	"time"

	"github.com/adiputra22/go-design-pattern/factory"
)

func main() {
	var contentCreator factory.ContentCreator

	// call imre
	contentCreator = &factory.Imre{}
	content := contentCreator.Produce(time.Now())
	content.Play()
	fmt.Println(contentCreator.Type())

	// call dewa
	contentCreator = &factory.DewaPrayoga{}
	content = contentCreator.Produce(time.Now())
	content.Play()
	fmt.Println(contentCreator.Type())
}

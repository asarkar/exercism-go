package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct{}

func (greeter Italian) LanguageName() string {
	return "Italian"
}

func (greeter Italian) Greet(name string) string {
	return fmt.Sprintf("I can speak %s: Ciao %s!", greeter.LanguageName(), name)
}

func SayHello(name string, greeter Greeter) string {
	return greeter.Greet(name)
}

type Portuguese struct{}

func (greeter Portuguese) LanguageName() string {
	return "Portuguese"
}

func (greeter Portuguese) Greet(name string) string {
	return fmt.Sprintf("I can speak %s: Olá %s!", greeter.LanguageName(), name)
}

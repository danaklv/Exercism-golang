package tasks

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(a string) string
}

func SayHello(name string, g Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

type GermanGreeter struct{}
type Italian struct{}
type Portuguese struct{}

func (g GermanGreeter) LanguageName() string {
    return "German"
}

func (g GermanGreeter) Greet(a string) string {
    return "Hallo " + a + "!"
}


func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet(a string) string {
    return "Ciao " + a + "!"
}


func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet(a string) string {
    return "Olá " + a + "!"
}


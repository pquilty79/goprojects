package main

type englishBot struct{}
type spanishBot struct{}

func main(){

}

func (eb englishBot) getGreeting() string {
     return "Hello"
}

func (sp spanishBot) getGreeting() string {
	return "Hola"
}



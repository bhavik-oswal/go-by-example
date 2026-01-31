package airportrobot
// Greeter defines how a language greets someone
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}
// SayHello returns the greeting using the given Greeter
func SayHello(name string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}
// Italian greeter
type Italian struct{}
func (Italian) LanguageName() string {
	return "Italian"
}
func (Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}
// Portuguese greeter
type Portuguese struct{}
func (Portuguese) LanguageName() string {
	return "Portuguese"
}
func (Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}

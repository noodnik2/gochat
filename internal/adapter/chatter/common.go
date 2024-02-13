package chatter

type Console interface {
	GetPrompt() string
	Print(text string)
	Println(a ...string)
	Printf(format string, args ...any)
}

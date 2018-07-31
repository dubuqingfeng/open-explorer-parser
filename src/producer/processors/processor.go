package processors

type Proccessor interface {
	Parse(message string) bool
	Finish(info string) (int, string)
}
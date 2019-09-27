package processors

type Processor interface {
	Parse(message string) bool
	Finish(info string) (int, string)
}

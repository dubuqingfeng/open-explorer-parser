package filters

type Filter interface {
	Filter(message string) bool
}
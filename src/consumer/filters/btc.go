package filters

type BTCFilter struct {
	status int
	reason string
}

func NewBTCFilter() *BTCFilter {
	return &BTCFilter{}
}

func (filter *BTCFilter) Filter(message string) bool {
	// Load Filter
	return false
}

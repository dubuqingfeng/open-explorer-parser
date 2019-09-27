package fetchers

type Fetcher interface {
	Fetch(title string) (bool, map[string]string)
	RPCCall(method string)
}

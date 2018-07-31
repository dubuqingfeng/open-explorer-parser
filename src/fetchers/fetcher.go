package fetchers

type Fetcher interface {
	Fetch(title string) (bool, string)
}
package network

import "net/url"

func FetchFeed(location *url.URL) ([]byte, error)

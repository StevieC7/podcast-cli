package parsing

import "github.com/mmcdole/gofeed"

type validPath string

type ValidPath interface {
	GetValue() validPath
}

func ParseValidPath(input string) (ValidPath, error)

func ParseFeed(feed []byte) (gofeed.Feed, error)

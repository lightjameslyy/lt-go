package parser

import (
	"regexp"

	"github.com/lightjameslyy/lt.go/imooclearngo/crawler/engine"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		// TODO: fix
		//result.Items = append(result.Items, name)
		result.Requests = append(result.Requests,
			engine.Request{Url: string(m[1]), ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(url, c, name)
			}},
		)
	}
	return result
}

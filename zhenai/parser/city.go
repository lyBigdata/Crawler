package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests,
			engine.Request{Url: url,
				ParserFunc: func(contents []byte) engine.ParseResult {
					return ParseProfile(contents, name, url)
				},
			})
	}

	return result
}

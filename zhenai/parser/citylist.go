package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

const (
	//<a href="http://album.zhenai.com/u/1361133512" target="_blank">怎么会迷上你</a>
	cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
)

func ParseCityList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityListReg)

	submatch := compile.FindAllSubmatch(contents, -1)

	//这里要把解析到的每个URL都生成一个新的request

	result := engine.ParseResult{}
	for _, m := range submatch {
		log.Printf("UserName:%s URL:%s\n", string(m[2]), string(m[1]))

		//把用户信息人名加到item里
		result.Request = append(result.Request,
			engine.Request{
				Url : string(m[1]),
				ParserFunc : ParseCity,
			})
	}

	return result
}
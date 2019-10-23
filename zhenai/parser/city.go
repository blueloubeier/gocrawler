package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

const (
	//<a href="http://album.zhenai.com/u/1361133512" target="_blank">怎么会迷上你</a>
	cityReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
)

func ParseCity(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityReg)

	submatch := compile.FindAllSubmatch(contents, -1)

	//这里要把解析到的每个URL都生成一个新的request

	result := engine.ParseResult{}

	for _, m := range submatch {
		name := string(m[2])
		log.Printf("UserName:%s URL:%s\n", string(m[2]), string(m[1]))

		//把用户信息人名加到item里
		result.Items = append(result.Items, name)

		result.Request = append(result.Request,
			engine.Request{
				//用户信息对应的URL,用于之后的用户信息爬取
				Url : string(m[1]),
				//这个parser是对城市下面的用户的parse
				ParserFunc : func(bytes []byte) engine.ParseResult {
					//这里使用闭包的方式;这里不能用m[2],否则所有for循环里的用户都会共用一个名字
					//需要拷贝m[2] ---- name := string(m[2])
					return ParseProfile(bytes, name)
				},
				//ParserFunc : engine.NilParser,
			})
	}

	return result
}

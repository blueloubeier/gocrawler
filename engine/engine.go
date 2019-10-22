package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request){

	//这里维持一个队列
	var requestsQueue []Request

	requestsQueue = append(requestsQueue, seeds...)

	for len(requestsQueue) > 0 {
		//取第一个
		r := requestsQueue[0]
		//只保留没处理的request
		requestsQueue = requestsQueue[1:]

		log.Printf("fetching url:%s\n", r.Url)
		//爬取数据
		body, err := fetcher.Fetch(r.Url)

		if err != nil {
			log.Printf("fetch url: %s; err: %v\n", r.Url, err)
			//发生错误继续爬取下一个url
			continue
		}

		//解析爬取到的结果
		result := r.ParserFunc(body)

		//把爬取结果里的request继续加到request队列
		requestsQueue = append(requestsQueue, result.Request...)

		//打印每个结果里的item,即打印城市名、城市下的人名...
		for _, item := range result.Items {
			log.Printf("get item is %v\n", item)
		}
	}
}
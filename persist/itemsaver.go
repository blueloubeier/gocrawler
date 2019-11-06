package persist

import (
	"context"
	"crawler/engine"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"log"
)

func ItemSaver() chan engine.Item{
	out :=make(chan engine.Item)
	go func(){
		itemCount :=0
		for{
			item := <-out
			log.Print("Item saver: got item" + "#%d: %v", itemCount, item)
			itemCount++
			_, err :=save(item)
			if err != nil{
				log.Print("Item Saver: error " +
					"saving item %v: %v",
					item, err)
			}
		}
	}()
	return out
}

func save(item engine.Item)(id string,err error){
	client, err :=elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil{
		return "",err
	}
	if item.Type == ""{
		return "", errors.New("must supply Type")
	}
	indexService :=client.Index().
		Index("dating_profiles").
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)
	if item.Id != ""{
		indexService.Id(item.Id)
	}

	resp,err := indexService.
		Do(context.Background())
	if err != nil{
		return "",err
	}
	return resp.Id,nil

}

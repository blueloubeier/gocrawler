package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T){
	expected := engine.Item{
		Url:"http://album.zhenai.com/u/108906739",
		Type:"zhenai",
		Id:"108906739",
		Payload:model.Profile{
			Age: 34,
			Height: 162,
			Weight: 57,
			Income: "3001-5000元",
		},
	}

	id, err :=save(expected)
	if err != nil{
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}
	resp, err:=client.Get().
		Index("dating_profiles").
		Type("zhenai").
		Id(id).
	    Do(context.Background())
	if err != nil{
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual engine.Item
	err = json.Unmarshal([]byte(resp.Source),&actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected{
		t.Errorf("got %v; expected %v",actual,expected)
	}
}


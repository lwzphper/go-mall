package es

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"strconv"
	"testing"
	"time"
)

const IndexName = "goods"

func init() {
	// https 方式
	//err := InitClientWithOptions(DefaultClient, []string{"https://127.0.0.1:9200"},
	//	"elastic", "123456", WithScheme("https"))
	err := InitClientWithOptions(DefaultClient, []string{"http://127.0.0.1:9200"},
		"elastic", "123456")
	if err != nil {
		EStdLogger.Print("InitClient error", err, "client", DefaultClient)
		panic(err)
	}
}

type Goods struct {
	Id             int64   `json:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Year           int     `json:"year"`
	LastMonthSales int     `json:"last_month_sales"`
	Favorites      int     `json:"favorites"`
}

var indexCreateJson = `
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0
  },
  "mappings": {
    "properties": {
      "id": {
        "type": "keyword",
        "doc_values": false,
        "norms": false,
        "similarity": "boolean"
      },
      "name": {
        "type": "text"
      },
        "price": {
        "type": "double"
      },
      "last_month_sales": {
        "type": "long"
      },
      "favorites": {
        "type": "long"
      },
      "year":{
        "type": "short"
      }
    }
  }
}
`

var id = strconv.FormatInt(doc.Id, 10)
var doc = Goods{
	Id:             1,
	Name:           "name1",
	Price:          1000,
	Year:           2022,
	LastMonthSales: 22,
	Favorites:      1939,
}

// 创建 mapping
func TestCreateMapping(t *testing.T) {
	ctx := context.Background()
	esClient := GetClient(DefaultClient)
	// 创建mapping
	err := esClient.CreateIndex(ctx, IndexName, indexCreateJson, false)
	if err != nil {
		EStdLogger.Print(err)
	}
}

// 创建数据
func TestCreate(t *testing.T) {
	ctx := context.Background()
	esClient := GetClient(DefaultClient)
	err := esClient.Create(ctx, IndexName, id, "", doc)
	if err != nil {
		EStdLogger.Print(err)
	}
}

// 查询
func TestQuery(t *testing.T) {
	ctx := context.Background()
	esClient := GetClient(DefaultClient)
	goods := make([]Goods, 0)
	orders := make([]map[string]bool, 1)
	orders = append(orders, map[string]bool{"favorites": true})
	res, err := esClient.Query(ctx, IndexName, nil, elastic.NewMatchAllQuery(), 0, 20, WithEnableDSL(true), WithOrders(orders))
	if err != nil {
		EStdLogger.Print(err)
	} else {
		if res != nil {
			for _, hit := range res.Hits.Hits {
				g := Goods{}
				docByte, err := hit.Source.MarshalJSON()
				if err != nil {
					EStdLogger.Print(err)
				} else {
					err = json.Unmarshal(docByte, &g)
					if err != nil {
						EStdLogger.Print(err)
					} else {
						goods = append(goods, g)
					}
				}
			}
		}

	}
	EStdLogger.Printf("%+v", goods)
}

// 批量操作
func TestBulk(t *testing.T) {
	esClient := GetClient(DefaultClient)

	for i := 0; i < 10; i++ {
		docID := strconv.Itoa(i)
		doc := Goods{
			Id:             int64(i),
			Name:           "name" + docID,
			Price:          float64(i),
			Year:           2022,
			LastMonthSales: i,
			Favorites:      i,
		}
		// 异步upsert（存在更新，不存在新增），doc为查询的条件
		/*update := map[string]interface{}{"name": "xxxx"}
		esClient.BulkUpsert(IndexName, docID, docID, update, doc)*/
		// 创建
		//esClient.BulkCreate(IndexName, docID, docID, doc)
		// 替换
		esClient.BulkReplace(IndexName, docID, docID, doc)
	}

	//因为是异步处理，这里需要等待本地channel提交
	time.Sleep(3 * time.Second)

}

// 更新
func TestUpdate(t *testing.T) {
	ctx := context.Background()
	esClient := GetClient(DefaultClient)
	err := esClient.Update(ctx, IndexName, id, "", map[string]interface{}{"name": "name2"})
	if err != nil {
		t.Error(err)
	}
}

// 带版本号更新
func TestUpdateByVersion(t *testing.T) {
	ctx := context.Background()
	esClient := GetClient(DefaultClient)
	doc.Name = "name4"
	err := esClient.UpsertWithVersion(ctx, IndexName, id, "", doc, 3)
	if err != nil {
		t.Error(err)
	}
}

// 带查询条件更新
func TestUpdateByQuery(t *testing.T) {
	ctx := context.Background()
	esClient := GetClient(DefaultClient)
	//注意：map中的键必须和updateScript中params.后面的字段名一一对应，大小写及命名方式要完全一致
	//updateScript := `ctx._source.name=request.name;ctx._source.favorites=request.favorites`
	// 更新id少于等于3的数据
	updateScript := `ctx._source['name']=params['name'];ctx._source['favorites']=params['favorites']`
	updateParams := map[string]interface{}{"name": "name123", "favorites": 123}
	_, err := esClient.UpdateByQuery(ctx, IndexName, []string{"1", "2", "3"}, elastic.NewRangeQuery("id").Lte(3), updateScript, updateParams)
	if err != nil {
		t.Error(err)
	}
}

// Upsert 不存在则插入，存在更新
func TestUpsert(t *testing.T) {
	ctx := context.Background()
	esClient := GetClient(DefaultClient)
	// doc 为查找的条件
	err := esClient.Upsert(ctx, IndexName, id, "", map[string]interface{}{"name": "name5"}, doc)
	if err != nil {
		t.Error(err)
	}
}

// 删除
func TestDelete(t *testing.T) {
	ctx := context.Background()
	esClient := GetClient(DefaultClient)
	err := esClient.Delete(ctx, IndexName, id, "")
	//err := esClient.DeleteWithVersion(ctx, IndexName, id, "", 5) // 指定版本号
	if err != nil {
		t.Error(err)
	}
}

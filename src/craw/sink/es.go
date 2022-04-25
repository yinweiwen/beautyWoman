package sink

import (
	"context"
	"craw/craw/utils"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

var (
	indexDzy = "dzy_data"
	typeDzy  = "_doc"
)

func init() {
	indexDzy = utils.GIniParser.GetString("es", "dzy_index")
}

// ESClient es客户端
type ESClient struct {
	_client *elastic.Client
}

// NewClient es 客户端
func NewEsClient(addr string) (*ESClient, error) {
	client, err := elastic.NewClient(elastic.SetURL(addr),elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()
	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(addr).Do(ctx)
	if err != nil {
		// Handle error
		return nil, err
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion(addr)
	if err != nil {
		// Handle error
		return nil, err
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	return &ESClient{
		_client: client,
	}, err
}

func (c *ESClient) IndexDzy(data map[string]interface{}) {
	if bts, err := json.Marshal(data); err != nil {
		log.Println(err.Error())
	} else {
		c.IndexDzyRaw(bts)
	}
}

// ConsumeAirQuality implments Comsumer
func (c *ESClient) IndexDzyRaw(data []byte) {
	ctx := context.Background()

	_, err := c._client.Index().
		Index(indexDzy).
		Type(typeDzy).
		BodyString(string(data)).
		Do(ctx)

	if err != nil {
		log.Println(err.Error())
	}
}

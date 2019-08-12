package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v6/esapi"

	"github.com/elastic/go-elasticsearch/v6"
)

type ElasticManager struct {
	ESClient   *elasticsearch.Client
	ESResponse *esapi.Response
}

func NewEsManager() *ElasticManager {

	log.SetFlags(0)

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return nil
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil
	}

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
		return nil
	}

	return &ElasticManager{
		ESClient:   es,
		ESResponse: res,
	}

}

func (this *ElasticManager) s() {

	// 2. Index documents concurrently
	//
	for i, title := range []string{"Test One", "Test Two"} {

		go func(i int, title string) {

			// Build the request body.
			var b strings.Builder
			b.WriteString(`{"title" : "`)
			b.WriteString(title)
			b.WriteString(`"}`)

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(b.String()),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), this.ESClient)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, title)
	}

	log.Println(strings.Repeat("-", 37))

	// 3. Search for the indexed documents
	//
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := this.ESClient.Search(
		this.ESClient.Search.WithContext(context.Background()),
		this.ESClient.Search.WithIndex("test"),
		this.ESClient.Search.WithBody(&buf),
		this.ESClient.Search.WithTrackTotalHits(true),
		this.ESClient.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	// if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	// 	log.Fatalf("Error parsing the response body: %s", err)
	// }
	// // Print the response status, number of results, and request duration.
	// log.Printf(
	// 	"[%s] %d hits; took: %dms",
	// 	res.Status(),
	// 	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	// 	int(r["took"].(float64)),
	// )
	// // Print the ID and document source for each hit.
	// for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	// 	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	// }

	log.Println(strings.Repeat("=", 37))
}

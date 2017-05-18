package main

import (
	"context"
	"io"
	"net/http"
	"gopkg.in/olivere/elastic.v5"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	
	
	client, err := elastic.NewClient(elastic.SetURL("https://search-debbie-oklee5b5u3drfszsu2cf5ahg7m.eu-west-1.es.amazonaws.com"))
	if err != nil {
		io.WriteString(w, "Client açılamadı.")
	}
	
	else{
	
		exists, err := client.IndexExists("content").Do(context.Background())
		if err != nil {
			io.WriteString(w, "Index Açılırken hata oldu.")
		}
		if !exists {
			// Index does not exist yet.
			io.WriteString(w, "Index bulunamadı.")
		}
		else{
		
			termQuery := elastic.NewTermQuery("text", "Dandanakan")
			searchResult, err := client.Search().
								Index("content").        // search in index "twitter"
								Query(termQuery).        // specify the query
								//Sort("user", true).      // sort by "user" field, ascending
								From(0).Size(10).        // take documents 0-9
								Pretty(true).            // pretty print request and response JSON
								Do(context.Background()) // execute
			
			if err != nil {
				// Handle error
				io.WriteString(w, "Search hata oldu.")
			}

			else{
			
				// searchResult is of type SearchResult and returns hits, suggestions,
				// and all kinds of other information from Elasticsearch.
				s := fmt.Sprintf("Query took milliseconds : %d ", searchResult.TookInMillis)
				io.WriteString(w,s)
			}
			
		}
		
		
	}
	


}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

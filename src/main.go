package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type historyKeyValue struct {
	Value     string
	Timestamp time.Time
	State     string
}

type keyValuePair struct {
	Key   string
	Value string
}

var store = make(map[string]string)
var historyStore = make(map[string][]historyKeyValue)

func appendHistoryStore(key, value, state string) {
	historyEntry := new(historyKeyValue)
	historyEntry.Value = value
	historyEntry.Timestamp = time.Now()
	historyEntry.State = state
	historyStore[key] = append(historyStore[key], *historyEntry)
}

func handleError(res *http.ResponseWriter) {
	if err := recover(); err != nil {
		fmt.Println(err)
		http.Error(*res, "", http.StatusInternalServerError)
	}
}

func getKey(req *http.Request) string {
	return req.URL.Query().Get("key")
}

func handleStore(res http.ResponseWriter, req *http.Request) {
	defer handleError(&res)
	switch req.Method {
	case http.MethodPost:
		defer req.Body.Close()
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		var kvp keyValuePair
		json.Unmarshal(body, &kvp)
		store[kvp.Key] = kvp.Value
		appendHistoryStore(kvp.Key, kvp.Value, "Updated")
		break
	case http.MethodDelete:
		key := getKey(req)
		if key != "" {
			delete(store, key)
			appendHistoryStore(key, "", "Deleted")
			break
		}
		fallthrough
	case http.MethodGet:
		historyValue, exists := store[getKey(req)]
		if exists {
			io.WriteString(res, historyValue)
			break
		}
		fallthrough
	default:
		http.Error(res, "", http.StatusNotFound)
	}
}

func handleHistory(res http.ResponseWriter, req *http.Request) {
	defer handleError(&res)
	switch req.Method {
	case http.MethodGet:
		historyStoreValue, exists := historyStore[getKey(req)]
		if exists {
			output, err := json.Marshal(historyStoreValue)
			if err != nil {
				panic(err)
			}
			res.Header().Add("Content-Type", "application/json")
			io.WriteString(res, string(output))
			break
		}
		fallthrough
	default:
		http.Error(res, "", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/store", handleStore)
	http.HandleFunc("/history", handleHistory)
	http.ListenAndServe(":9000", nil)
}

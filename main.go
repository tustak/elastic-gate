package main

import "github.com/tustak/elastic-gate/connection"
import "github.com/tustak/elastic-gate/transaction"
import "fmt"
import "encoding/json"

func main(){
    c := connection.Credentials{Host: "localhost", Port: "9200", 
    Username: "", Password: ""}
    var f interface{}
    b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
    _ = json.Unmarshal(b, &f)
    m := f.(map[string]interface{})
    fmt.Println(m["Name"])
    t := transaction.New("errors", "transaction", "", "")
    _ = t.InsertNew(&c)

}

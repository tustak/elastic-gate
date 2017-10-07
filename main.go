package main
import "github.com/tustak/elastic-gate/connection"
import "fmt"

func main(){
    c := connection.Credentials{Host: "localhost", Port: "9200", 
    Username: "", Password: ""}
    fmt.Println(connection.BaseURI(&c))
}

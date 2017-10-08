package transaction

//import "time"
import "io/ioutil"
import "fmt"
import "errors"
import "time"
import "bytes"
import "net/http"
import "encoding/json"
import "github.com/tustak/elastic-gate/connection"

var indexName string = "errors"
var typeName string = "transaction"

type Transaction struct{
    Id string
    LenderId string
    BorrowerId string
    Date time.Time
}

type transNoID struct{
    // The same than Transaction, but without Id field
    LenderId string
    BorrowerId string
    Date time.Time
}

func New(LenderId string, BorrowerId string) Transaction{
    return Transaction{"", LenderId, BorrowerId, time.Now()}
}

func (transaction *Transaction) InsertNew(cred *connection.Credentials) error{
    tnid := transNoID{transaction.LenderId, transaction.BorrowerId, transaction.Date}
    transJSONstr, _ := json.Marshal(tnid)
    url := connection.GetInsertURI(cred, indexName, typeName)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(transJSONstr))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    r, err := client.Do(req)
    if err == nil {
        var f map[string]interface{}
        _ = json.NewDecoder(r.Body).Decode(&f)
        transaction.Id = f["_id"].(string)
    } else { transaction.Id = ""
    }
    return err
}

func GetById(Id string, cred *connection.Credentials) (Transaction, error){
    url := fmt.Sprintf("%s/%s/%s/%s/_source", connection.BaseURI(cred), "errors", "transaction", Id)
    r, _ := http.Get(url)
    var t Transaction
    var err error
    if r.Status == "200 OK" {
        byteBody, _ := ioutil.ReadAll(r.Body)
        json.Unmarshal(byteBody, &t)
        t.Id = Id
        err = nil
    } else {
        err = errors.New("Some fancy error")
    }
    return t, err
}

func getByFieldId(FieldId string, Id string, cred *connection.Credentials) (
    []interface{}, error){
    url := connection.GetSearchURI(cred, indexName, typeName)
    fmt.Println(url)
    queryJson := fmt.Sprintf(`{
        "query": {
            "match": {
                "%s": "%s"
            }
        }

    }`, FieldId, Id)
    req, _ := http.NewRequest("GET", url, bytes.NewBufferString(queryJson))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    var err error
    var data []interface{}
    var f map[string] interface{}
    r, _ := client.Do(req)
    if r.Status == "200 OK" {
        _ = json.NewDecoder(r.Body).Decode(&f)
        fmt.Println(f["hits"].(map[string] interface{})["hits"])
        data = f["hits"].(map[string] interface{})["hits"].([]interface{})
        err = nil
    } else {
        err = errors.New("Could not process the request")
    }
    return data, err
}


func GetByBorrowerId(borrowerId string, cred *connection.Credentials) (
    []interface{}, error){
    var fieldId string = "BorrowerId"
    return getByFieldId(fieldId,  borrowerId, cred)
}

func GetByLenderId(lenderId string, cred *connection.Credentials) (
    []interface{}, error){
    var fieldId string = "LenderId"
    return getByFieldId(fieldId,  lenderId, cred)
}

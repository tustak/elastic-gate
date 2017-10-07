package transaction

//import "time"
//import "io/ioutil"
import "fmt"
import "time"
import "bytes"
import "net/http"
import "encoding/json"
import "github.com/tustak/elastic-gate/connection"

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
    indexName := "errors"
    typeName := "transaction"
    url := connection.GetInsertURI(cred, indexName, typeName)
    fmt.Println(url)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(transJSONstr))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    r, err := client.Do(req)
    if err == nil {
        var f map[string]interface{}
        //bodyBytes, _ := ioutil.ReadAll(r.Body)
        _ = json.NewDecoder(r.Body).Decode(&f)
        fmt.Println(f["_id"])
        transaction.Id = f["_id"].(string)
    } else {
        transaction.Id = ""
    }
    fmt.Println(transaction)
    return err
}

func GetById(Id string) (Transaction, error){
    t := Transaction{"", "", "", time.Now()}
    return t, nil
}

func GetByBorrowerId(UserId string) (Transaction, error){
    t := Transaction{"", "", "", time.Now()}
    return t, nil
}

func GetByLenderId(UserId string) (Transaction, error){
    t := Transaction{"", "", "", time.Now()}
    return t, nil
}

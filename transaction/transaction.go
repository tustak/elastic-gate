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
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(transJSONstr))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    r, err := client.Do(req)
    if err == nil {
        var f map[string]interface{}
        _ = json.NewDecoder(r.Body).Decode(&f)
        transaction.Id = f["_id"].(string)
    } else {
        transaction.Id = ""
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

func GetByBorrowerId(UserId string) (Transaction, error){
    t := Transaction{"", "", "", time.Now()}
    return t, nil
}

func GetByLenderId(UserId string) (Transaction, error){
    t := Transaction{"", "", "", time.Now()}
    return t, nil
}

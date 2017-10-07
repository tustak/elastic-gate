package transaction

//import "time"
import "fmt"
import "time"
import "bytes"
import "net/http"
import "encoding/json"
import "github.com/tustak/elastic-gate/connection"

type Transaction struct{
    Index string
    Type string
    Id string
    LenderId string
    BorrowerId string
    Date time.Time
}

type transNoID struct{
    // The same than Transaction, but without Id field
    Index string
    Type string
    LenderId string
    BorrowerId string
    Date time.Time
}

func New(Index string, Type string, LenderId string, BorrowerId string) Transaction{
    return Transaction{Index, Type, "", LenderId, BorrowerId, time.Now()}
}

func (transaction *Transaction) InsertNew(cred *connection.Credentials) error{
    tnid := transNoID{transaction.Index, transaction.Type,
                      transaction.LenderId, transaction.BorrowerId, transaction.Date}
    transJSONstr, _ := json.Marshal(tnid)
    url := connection.GetInsertURI(cred, transaction.Index, transaction.Type)
    fmt.Println(url)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(transJSONstr))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    r, err := client.Do(req)
    if err == nil {
        transaction.Id = "something"
    } else {
        transaction.Id = ""
    }
    fmt.Println(r)
    return err
}

func GetById(Id string) (Transaction, error){
    t := Transaction{"", "", "", "", "", time.Now()}
    return t, nil
}

func GetByBorrowerId(UserId string) (Transaction, error){
    t := Transaction{"", "", "", "", "", time.Now()}
    return t, nil
}

func GetByLenderId(UserId string) (Transaction, error){
    t := Transaction{"", "", "", "", "", time.Now()}
    return t, nil
}

package transaction

//import "time"
import "fmt"
import "time"
import "net/http"
import "github.com/tustak/elastic-gate/connection"

type Transaction struct{
    Index string
    Type string
    Id string
    LenderId string
    BorrowerId string
    Date time.Time
}

func New(Index string, Type string, LenderId string, BorrowerId string, Date time.Time) Transaction{
    return Transaction{"", "", "", "", "", time.Now()}
}

func (transaction *Transaction) InsertNew(cred *connection.Credentials) error{
    url := connection.GetSearchURI(cred, transaction.Index, transaction.Type)
    r, err := http.Get(url)
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

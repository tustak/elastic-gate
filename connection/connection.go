package connection

import "fmt"

func BaseURI(cred * Credentials) string{
    return fmt.Sprintf("http://%s:%s/", cred.Host, cred.Port)
}

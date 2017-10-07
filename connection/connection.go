package connection

import "fmt"

func BaseURI(cred *Credentials) string{
    return fmt.Sprintf("http://%s:%s", cred.Host, cred.Port)
}

func GetSearchURI(cred * Credentials, indexName string, typeName string) string{
    baseURI := BaseURI(cred)
    var searchURI string
    if typeName == "" {
    searchURI= fmt.Sprintf("%s/%s/_search", baseURI, indexName)
    } else {
    searchURI = fmt.Sprintf("%s/%s/%s/_search", baseURI, indexName, typeName)
    }
    return searchURI
}

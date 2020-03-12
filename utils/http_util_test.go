package utils

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDelete(t *testing.T) {
	url:="http://192.168.100.118:9301/5,15561c540361"
	head, err := http.Head(url)
	if err != nil {

	}
	fmt.Println(head.Request.URL.String())
}
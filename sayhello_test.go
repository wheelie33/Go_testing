package main

import (
	"fmt"
	"testing"
)

func Test_DidSayHelloGetSaid(t *testing.T) {
	said := SayHi("Jason")

	if said != "Hello Jason!" {
		var errorString = fmt.Sprintf("expected 'Hello Jason!' but said='%s'", said)
		t.Error(errorString)
	}
}

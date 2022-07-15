package opensea

import (
	"fmt"
	"log"
	"testing"
)

func TestGetStats(t *testing.T) {
	res, err := GetStats("theirsverse-official")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}

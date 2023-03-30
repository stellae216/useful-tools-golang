package unit

import (
	"sync"
	"testing"

	"github.com/stellae216/useful-tools-golang/application"
)

func TestSnowflake(t *testing.T) {
	snow1 := application.Snowflake{
		WorkerId:     1,
		DataCenterId: 4,
	}
	snow2 := application.Snowflake{
		WorkerId:     1,
		DataCenterId: 5,
	}
	var wg = sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			t.Log(snow1.NextVal())
			t.Log(snow2.NextVal())
			wg.Done()
		}()
	}
	wg.Wait()
}

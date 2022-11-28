package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

type ValidStruct struct {
	Required string `json:"required" binding:"required"`
}

func TestBindAndValid(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			var req ValidStruct
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			ctx.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(`{"required": ""}`))
			ctx.Request.Header.Add("Content-Type", "application/json")
			valid, errs := BindAndValid(ctx, &req)
			if !valid {
				t.Log(errs.Errors())
				return
			}
		}()
	}
	wg.Wait()
}

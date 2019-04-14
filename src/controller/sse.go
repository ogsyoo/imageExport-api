package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"net/http"
	"ogsyoo/imageExport-api/src/common/conf"
	"ogsyoo/imageExport-api/src/sse"
	"time"
)

func Subcribe(c *gin.Context) {
	rw := c.Writer
	s := conf.SseClient
	response := c.Writer
	request := c.Request
	name := c.Param("name")
	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(response, "Streaming unsupported.", http.StatusInternalServerError)
		return
	}
	h := response.Header()
	lastEventID := request.Header.Get("Last-Event-ID")
	if request.Method == "GET" {
		h.Set("Content-Type", "text/event-stream")
		h.Set("Cache-Control", "no-cache")
		h.Set("Connection", "keep-alive")
		closeNotify := response.(http.CloseNotifier).CloseNotify()
		client := conf.SseClient.Sub(request, name, lastEventID, h, closeNotify)
		response.WriteHeader(http.StatusOK)
		flusher.Flush()
		for msg := range client.Send {
			msg.Retry = s.Options.RetryInterval
			fmt.Fprintf(response, msg.String())
			flusher.Flush()
		}
	} else if request.Method != "OPTIONS" {
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func Publish(c *gin.Context) {
	go func() {
		for {
			m1 := sse.Message{"123", "dsfa", "weq", 1}
			b, _ := json.Marshal(m1)
			m := sse.NewMessage("1", string(b), "sdaf")
			//conf.SseClient.SendMessage("message", sse.SimpleMessage(time.Now().String()))
			conf.SseClient.SendMessage("message", m)
			time.Sleep(5 * time.Second)
		}
	}()
}

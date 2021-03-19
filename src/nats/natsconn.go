package main
import (
    "log"
    "runtime"
    "github.com/nats-io/nats.go"
    "time"
    "encoding/json"
)
func main() {
    nc, err := nats.Connect("nats://localhost:4222",
            nats.Name("practical-nats-client"),
//            nats.UserInfo("foo", "secret"),
    )
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
     nc.Subscribe("greeting", func(m *nats.Msg) {
        log.Printf("[Received] %s", string(m.Data))
    })
     nc.Subscribe(">", func(m *nats.Msg) {
        log.Printf("[Received] %s", string(m.Data))
        time.Sleep(1* time.Second)
    })
    payload := struct {
               RequestID string
               Data      []byte
    }{
               RequestID: "1234-5678-90",
               Data:      []byte("encoded data"),
    }
    msg, err := json.Marshal(payload)
    if err != nil {
       log.Fatalf("Error: %s", err)
    }

    nc.Publish("greeting", msg) 
    for i := 0; i < 10; i++ {
        nc.Publish("greeting", []byte("hello world"))
    }

    runtime.Goexit()
}

package main
import (
    "log"
    "runtime"
    "github.com/nats-io/nats.go"
)
func main() {
    nc, err := nats.Connect("nats://localhost:4222")
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    nc.Subscribe("greeting", func(m *nats.Msg) {
        log.Printf("[Received] %s", string(m.Data))
    })
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    for i := 0; i < 10; i++ {
          nc.Publish("greeting",[]byte("Hello world!!!"))
    }
    nc.Flush()
    err = nc.Publish("greeting", []byte("hello world after flush!!!"))
    if err != nil {
       log.Fatalf("Error: %s", err)
    }
    runtime.Goexit()
}

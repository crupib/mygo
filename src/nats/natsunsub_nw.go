package main
import (
    "log"
    "time"
    "runtime"
    "github.com/nats-io/nats.go"
)
func main() {
    nc, err := nats.Connect("nats://localhost:4222")
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    sub, err := nc.Subscribe("greeting", func(m *nats.Msg) {
        log.Printf("[Received-from client] %s", string(m.Data))
    })
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    for i := 0; i < 5; i++ {
          nc.Publish("greeting",[]byte("Hello world!!!"))
    }
    nc.Flush()
    time.Sleep(2 * time.Second)    
    sub.Unsubscribe()
    for i := 0; i < 5; i++ {
          nc.Publish("greeting",[]byte("Hello world!!!"))
    }
    runtime.Goexit()
}

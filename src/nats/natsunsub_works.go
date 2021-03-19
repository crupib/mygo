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
    var counter int
    var sub *nats.Subscription
    sub, err = nc.Subscribe("greeting", func(m *nats.Msg) {
        log.Printf("[Received] %s", string(m.Data))
        counter++
        if counter == 2 {
         sub.Unsubscribe()
        }
    })
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    for i := 0; i < 5; i++ {
          nc.Publish("greeting",[]byte("Hello world!!!"))
    }
    nc.Flush()
    runtime.Goexit()
}

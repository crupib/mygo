package main
import (
    "log"
    "github.com/nats-io/nats.go"
)
func main() {
    nc, err := nats.Connect("nats://localhost:4222")
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    msg := []byte("Hello World!")
    nc.Subscribe("greetings", func(_ *nats.Msg) {})
    for i := 0; i < 100000000; i++ {
        nc.Publish("greetings",msg)
    }
}

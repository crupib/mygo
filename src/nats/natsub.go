package main
import (
    "log"
    "runtime"
    "github.com/nats-io/nats.go"
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
        log.Printf("[<Received>] %s", string(m.Data))
    })
    runtime.Goexit()
}

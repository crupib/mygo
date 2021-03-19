package main
import (
    "log"
    "runtime"
    "github.com/nats-io/nats.go"
    "time"
)
func main() {
    nc, err := nats.Connect("nats://localhost:4222",
            nats.Name("practical-nats-client"),
//            nats.UserInfo("foo", "secret"),
    )
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    for i := 0; i < 10; i++ {
        nc.Publish("greeting", []byte("hello world"))
        time.Sleep(10 * time.Second)
    }

    runtime.Goexit()
}

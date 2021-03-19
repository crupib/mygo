package main
import (
    "log"
    "github.com/nats-io/nats.go"
    "time"
)
func main() {
    nc, err := nats.Connect("nats://localhost:4222",
            nats.Name("practical-nats-client"),
            nats.UseOldRequestStyle(),
//            nats.UserInfo("foo", "secret"),
    )
    if err != nil {
        log.Fatalf("Error: %s", err)
	}
    nc.Subscribe("help",func(m *nats.Msg) {
	log.Printf("[Recieved]: %s",string(m.Data))
	nc.Publish(m.Reply, []byte("I can help!!!"))
    })
    response, err := nc.Request("help", []byte("help!!"), 1*time.Second)
    if err != nil {
	log.Fatalf("Error: %s",err)
    }
    log.Println("[Response]:"+ string(response.Data))
}

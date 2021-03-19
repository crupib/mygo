package main
import (
     "fmt"
     "os"
     "gopkg.in/ini.v1"
)
func main() {
     cfg, err := ini.Load("conf.ini")
     if err != nil {
            fmt.Println("Failed to parse config file: %s", err)
            os.Exit(1)
     }
     fmt.Println(cfg.Section("Section").Key("path").String())
     fmt.Println(cfg.Section("Section").Key("enabled").String())
}

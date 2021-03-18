package main
import (
       "fmt"
       "github.com/kylelemons/go-gypsy/yaml"
)

func main() {
     config, err := yaml.ReadFile("/Users/williamcrupi/go/src/yaml_config/conf.yaml")
     if err != nil {
        fmt.Println(err)
     }
     mypath, err := config.Get("path")
     myenable, err := config.GetBool("enabled")
     fmt.Println(mypath)
     fmt.Println(myenable)
}

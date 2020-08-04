package main


import (
	"encoding/json"
	"fmt"
    "gopkg.in/ini.v1"
)

type Nodes struct {
	Nodes []Node
}

type Node struct {
	Endpoint string
	Address string
	PublicKey string
	AllowConnect bool
}

func main() {
	textJson := `{"nodes":[{"endpoint":"asia-hs.kirov-opensource.com:5555","address":"10.78.78.254/24","publicKey":"abcdefg","allowConnect":true}]}`

	var data Nodes

	json.Unmarshal([]byte(textJson), &data)
	
	

	// out, err := exec.Command("wg","showconf","wg0").Output()
	cfg, _ := ini.Load("wg0.conf")
	fmt.Println(cfg.Section("Interface").Key("ListenPort").String())
	fmt.Printf(data.Nodes[0].Endpoint)

	for i := 0; i < len(data.Nodes); i++ {
		fmt.Println("Endpoint: ", data.Nodes[i].Endpoint)
		fmt.Println("PublicKey: ", data.Nodes[i].PublicKey)
	}
	fmt.Printf("abc")
}
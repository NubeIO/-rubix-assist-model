package main

import (
	"encoding/json"
	"github.com/NubeIO/rubix-assist-model/model"
	"os"
)

func main() {
	PrintJOSN(model.GetNetworkSchema())
}

func PrintJOSN(x interface{}) {
	ioWriter := os.Stdout
	w := json.NewEncoder(ioWriter)
	w.SetIndent("", "    ")
	w.Encode(x)
}

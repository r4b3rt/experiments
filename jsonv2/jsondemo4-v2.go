package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
)

type Data struct {
	Tags       []string          `json:"tags"`                        // nil slice
	Attrs      map[string]string `json:"attrs"`                       // nil map
	MaybeTags  []string          `json:"maybe_tags,format:emitnull"`  // v2: 强制为 null
	MaybeAttrs map[string]string `json:"maybe_attrs,format:emitnull"` // v2: 强制为 null
}

func main() {
	fmt.Println("--- Testing Nil Slice/Map Serialization ---")
	d := Data{} // Tags 和 Attrs 都是 nil

	jsonData, _ := json.Marshal(d, json.Deterministic(true))
	(*jsontext.Value)(&jsonData).Indent() // indent for readability

	fmt.Println("Serialized Output (run with go and gotip to compare):\n", string(jsonData))
}

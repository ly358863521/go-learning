package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"Age"`
	Lesson
}
type Lesson struct {
	Lessons []string `json:"lessons"`
}

func main() {
	jsonStr := `{"Age": 18,"name": "Jim" ,"s": "男",
	"lessons":["English","History"],"Room":201,"n":null,"b":false}`
	var hu Human
	if err := json.Unmarshal([]byte(jsonStr), &hu); err == nil {
		fmt.Println(hu)
	}
	var le Lesson
	if err := json.Unmarshal([]byte(jsonStr), &le); err == nil {
		fmt.Println(le)
	}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err == nil {
		fmt.Println(data)
	}
	for k, v := range data {
		// vv := v.(type)  .(type)类型查询，只能在switch中使用
		switch vv := v.(type) {
		case string:
			fmt.Println(k, v, vv)
		}
	}
	strR := strings.NewReader(jsonStr)
	fmt.Println(strR)
	h := &Human{}
	if err := json.NewDecoder(strR).Decode(h); err != nil {
		fmt.Println(err)
	}
	fmt.Println(h)
	f, _ := os.Create("./t.json")
	json.NewEncoder(f).Encode(h)

}

/*
Human.Name字段，由于可以等到使用的时候，再根据具体数据类型来解析，因此我们可以延迟解析。当结构体Human的Name字段的类型设置为 json.RawMessage 时，它将在解码后继续以 byte 数组方式存在。
*/

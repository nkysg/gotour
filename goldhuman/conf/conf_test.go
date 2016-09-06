package conf

import (
	"fmt"
)

func ExampleGetValue() {
	cf := NewConf()
	cf.Load("./config.ini")
	val := cf.GetValue("codis", "product")
	fmt.Println(val)
}

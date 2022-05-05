package main

import (
	"fmt"
	""
)

func main() {
	var product = Product{
		10, "Pool", "Desc", 100.00, "sku", "2022/12/12", "2022/12/12", "2022/12/12",
	}
	fmt.Printf("product :: %v", product)
}

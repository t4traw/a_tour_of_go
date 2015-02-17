package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}

// パッケージをインポートすると、そのパッケージが外部に公開（エクスポート）している名前を参照することができます。
// Goでは、最初の文字が大文字で始まる場合は、その名前はエクスポートされています。

package main

import (
	"fmt"
	"strings"

	"link-parser/link"
)

var explHtml = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>

`

func main() {
	r := strings.NewReader(explHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", links)
}

package main

import (
	"fmt"
	"strings"

	"link-parser/link"
)

var explHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link
  
  to another page</a>
  <a href="/2-page">A second link to another page</a>
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

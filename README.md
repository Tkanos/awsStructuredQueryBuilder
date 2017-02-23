# awsStructuredQueryBuilder
Aws Structured Query Builder for Go

```go
package main

import (
	"fmt"

	builder "github.com/tkanos/awsStructuredQueryBuilder"
)

func main() {

	//(and (range field=year [2013,2015]) (or (term field=title boost=2 'star') (prefix field=plot 'star')))
	q := builder.And(builder.Rangei("year", 2013, 2015), builder.Or(builder.NewTerms("title", "star").WithBoost(2), builder.Prefix("plot", "star")))

	fmt.Print(q)
}
```

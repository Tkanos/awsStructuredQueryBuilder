# awsStructuredQueryBuilder
Aws Structured Query Builder for Go baesd on http://docs.aws.amazon.com/cloudsearch/latest/developerguide/searching-compound-queries.html

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


## What is missing (TODO)
- Range [2013,}
- Range {,2000} //less than 2000
- MatchAll

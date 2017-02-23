package awsStructuredQueryBuilder

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//StructuredQuery represent the structured query format used by aws based on http://docs.aws.amazon.com/cloudsearch/latest/developerguide/searching-compound-queries.html
type StructuredQuery struct {
	field       string
	value       string
	operator    string
	from        string
	to          string
	distance    int
	boost       int
	expressions []StructuredQuery
}

//todo Matchall

//NewTerms produce (term field=field boost=N 'STRING')
func NewTerms(field string, value string) StructuredQuery {
	q := StructuredQuery{
		field:    field,
		value:    fmt.Sprintf("'%s'", value),
		operator: "term",
	}

	return q
}

//NewTermi produce (term field=field boost=N value)
func NewTermi(field string, value int) StructuredQuery {
	q := StructuredQuery{
		field:    field,
		value:    strconv.Itoa(value),
		operator: "term",
	}

	return q
}

//NewTermf produce (term field=field boost=N value)
func NewTermf(field string, value float64) StructuredQuery {
	q := StructuredQuery{
		field:    field,
		value:    strconv.FormatFloat(value, 'g', -1, 64),
		operator: "term",
	}

	return q
}

//NewTermt produce (term field=field boost=N 'yyyy-mm-ddTHH:mm:ss.SSSZ') using IETF RFC3339 for date format according to cf http://docs.aws.amazon.com/cloudsearch/latest/developerguide/searching-dates.html
func NewTermt(field string, value time.Time) StructuredQuery {
	q := StructuredQuery{
		field:    field,
		value:    fmt.Sprintf("'%s'", dateFormat(value)),
		operator: "term",
	}

	return q
}

//Phrase produce (phrase field=field boost=N 'STRING')
func Phrase(field string, value string) StructuredQuery {
	q := StructuredQuery{
		field:    field,
		value:    fmt.Sprintf("'%s'", value),
		operator: "phrase",
	}

	return q
}

//Near produce (near field=FIELD distance=N boost=N 'STRING')
func Near(field string, value string, distance int) StructuredQuery {
	q := StructuredQuery{
		field:    field,
		value:    fmt.Sprintf("'%s'", value),
		distance: distance,
		operator: "near",
	}

	return q
}

//Prefix produce (prefix field=field boost=N 'STRING')
func Prefix(field string, value string) StructuredQuery {
	q := StructuredQuery{
		field:    field,
		value:    fmt.Sprintf("'%s'", value),
		operator: "prefix",
	}

	return q
}

//Ranges produce (range field=field boost=N ['from','to'])
func Ranges(field string, from, to string) StructuredQuery {
	q := StructuredQuery{
		field:    field,
		from:     fmt.Sprintf("'%s'", from),
		to:       fmt.Sprintf("'%s'", to),
		operator: "range",
	}

	return q
}

//Rangei produce (range field=field boost=N [from, to])
func Rangei(field string, from, to int) StructuredQuery {
	q := StructuredQuery{
		field:    field,
		from:     strconv.Itoa(from),
		to:       strconv.Itoa(to),
		operator: "range",
	}

	return q
}

//Rangef produce (range field=field boost=N [from, to])
func Rangef(field string, from, to float64) StructuredQuery {
	q := StructuredQuery{
		field:    field,
		from:     strconv.FormatFloat(from, 'g', -1, 64),
		to:       strconv.FormatFloat(to, 'g', -1, 64),
		operator: "range",
	}

	return q
}

//Ranget produce (range field=field boost=N ['yyyy-mm-ddTHH:mm:ss.SSSZ','yyyy-mm-ddTHH:mm:ss.SSSZto']) using IETF RFC3339 for date format according to cf http://docs.aws.amazon.com/cloudsearch/latest/developerguide/searching-dates.html
func Ranget(field string, from, to time.Time) StructuredQuery {
	q := StructuredQuery{
		field:    field,
		from:     fmt.Sprintf("'%s'", dateFormat(from)),
		to:       fmt.Sprintf("'%s'", dateFormat(to)),
		operator: "range",
	}

	return q
}

//'yyyy-mm-ddTHH:mm:ss.SSSZ' using IETF RFC3339 for date format according to cf http://docs.aws.amazon.com/cloudsearch/latest/developerguide/searching-dates.html
func dateFormat(value time.Time) string {
	return value.Format(time.RFC3339)
}

func newQueries(operator string, queries []StructuredQuery) StructuredQuery {
	q := StructuredQuery{}

	if len(queries) > 0 {
		q.operator = operator
		q.expressions = queries
	}

	return q
}

//And produce (and boost=N EXPRESSION1 EXPRESSION2 ... EXPRESSIONn)
func And(queries ...StructuredQuery) StructuredQuery {
	return newQueries("and", queries)
}

//Or produce (or boost=N EXPRESSION1 EXPRESSION2 ... EXPRESSIONn)
func Or(queries ...StructuredQuery) StructuredQuery {
	return newQueries("or", queries)
}

//Not produce (not boost=N EXPRESSION)
func Not(query StructuredQuery) StructuredQuery {
	queries := [...]StructuredQuery{query}
	return newQueries("not", queries[:])
}

//WithBoost Add boost to the existingQuery
func (q StructuredQuery) WithBoost(boost int) StructuredQuery {
	q.boost = boost
	return q
}

func (q StructuredQuery) String() string {

	var result string

	switch {
	case q.distance > 0:
		//(near field=FIELD distance=N boost=N 'STRING')
		if q.boost > 0 {
			result = fmt.Sprintf("(%s field=%s distance=%v boost=%v %s)", q.operator, q.field, q.distance, q.boost, q.value)
		} else {
			result = fmt.Sprintf("(%s field=%s distance=%v %s)", q.operator, q.field, q.distance, q.value)
		}
	case len(q.expressions) > 0:
		//(and boost=N EXPRESSION1 EXPRESSION2 ... EXPRESSIONn)
		//(or boost=N EXPRESSION1 EXPRESSION2 ... EXPRESSIONn)
		//(not boost=N EXPRESSION)
		if q.boost > 0 {
			result = fmt.Sprintf("(%s boost=%v %s)", q.operator, q.boost, strings.Trim(fmt.Sprint(q.expressions), "[]"))
		} else {
			result = fmt.Sprintf("(%s %s)", q.operator, strings.Trim(fmt.Sprint(q.expressions), "[]"))
		}
	case len(q.value) > 0:
		//(phrase field=FIELD boost=N 'STRING')
		//(prefix field=FIELD boost=N 'STRING')
		//(term field=FIELD boost=N 'STRING'|VALUE)
		if q.boost > 0 {
			result = fmt.Sprintf("(%s field=%s boost=%v %s)", q.operator, q.field, q.boost, q.value)
		} else {
			result = fmt.Sprintf("(%s field=%s %s)", q.operator, q.field, q.value)
		}
	case len(q.from) > 0 && len(q.to) > 0:
		//(range field=FIELD boost=N RANGE)
		if q.boost > 0 {
			result = fmt.Sprintf("(%s field=%s boost=%v [%s, %s])", q.operator, q.field, q.boost, q.from, q.to)
		} else {
			result = fmt.Sprintf("(%s field=%s [%s, %s])", q.operator, q.field, q.from, q.to)
		}
	}

	return result

}

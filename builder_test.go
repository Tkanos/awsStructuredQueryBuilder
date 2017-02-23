package awsStructuredQueryBuilder

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//TestNewTerms_Should_Create_StructuredQuery Test NewTerms in order to get (term field=FIELD 'STRING')
func TestNewTerms_Should_Create_StructuredQuery(t *testing.T) {
	q := NewTerms("title", "star")

	assert.Equal(t, "(term field=title 'star')", q.String(), "They should be equal")
}

//TestNewTermi_Should_Create_StructuredQuery Test NewTermi in order to get (term field=FIELD VALUE)
func TestNewTermi_Should_Create_StructuredQuery(t *testing.T) {
	q := NewTermi("title", 10)

	assert.Equal(t, "(term field=title 10)", q.String(), "They should be equal")
}

//TestNewTermf_Should_Create_StructuredQuery Test NewTermf in order to get (term field=FIELD VALUE)
func TestNewTermf_Should_Create_StructuredQuery(t *testing.T) {
	q := NewTermf("title", 10.123)

	assert.Equal(t, "(term field=title 10.123)", q.String(), "They should be equal")
}

//TestNewTermt_Should_Create_StructuredQuery Test NewTermt in order to get (term field=FIELD 'DATE')
func TestNewTermt_Should_Create_StructuredQuery(t *testing.T) {
	q := NewTermt("title", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	assert.Equal(t, "(term field=title '2009-11-10T23:00:00Z')", q.String(), "They should be equal")
}

//TestNewTermsWithBoost_Should_Create_StructuredQuery Test NewTermf with boost in order to get (term field=FIELD boost=N VALUE)
func TestNewTermsWithBoost_Should_Create_StructuredQuery(t *testing.T) {
	q := NewTerms("title", "star").WithBoost(2)

	assert.Equal(t, "(term field=title boost=2 'star')", q.String(), "They should be equal")
}

//TestPhrase_Should_Create_StructuredQuery Test Phrase in order to get (phrase field=FIELD 'STRING')
func TestPhrase_Should_Create_StructuredQuery(t *testing.T) {
	q := Phrase("title", "star wars")

	assert.Equal(t, "(phrase field=title 'star wars')", q.String(), "They should be equal")
}

//TestNear_Should_Create_StructuredQuery Test Near in order to get (near field=FIELD distance=N 'STRING')
func TestNear_Should_Create_StructuredQuery(t *testing.T) {
	q := Near("title", "star", 10)

	assert.Equal(t, "(near field=title distance=10 'star')", q.String(), "They should be equal")
}

//TestNearWithBoost_Should_Create_StructuredQuery Test Near With Boost in order to get (near field=FIELD distance=N boost=N 'STRING')
func TestNearWithBoost_Should_Create_StructuredQuery(t *testing.T) {
	q := Near("title", "star", 10).WithBoost(2)

	assert.Equal(t, "(near field=title distance=10 boost=2 'star')", q.String(), "They should be equal")
}

//TestPrefix_Should_Create_StructuredQuery Test Prefix in order to get (prefix field=FIELD value)
func TestPrefix_Should_Create_StructuredQuery(t *testing.T) {
	q := Prefix("title", "star")

	assert.Equal(t, "(prefix field=title 'star')", q.String(), "They should be equal")
}

//TestRanges_Should_Create_StructuredQuery Test Ranges in order to get (range field=FIELD ['from','to'])
func TestRanges_Should_Create_StructuredQuery(t *testing.T) {
	q := Ranges("title", "star", "wars")

	assert.Equal(t, "(range field=title ['star', 'wars'])", q.String(), "They should be equal")
}

//TestRangei_Should_Create_StructuredQuery Test Rangei in order to get (range field=FIELD [from, to])
func TestRangei_Should_Create_StructuredQuery(t *testing.T) {
	q := Rangei("title", 1, 10)

	assert.Equal(t, "(range field=title [1, 10])", q.String(), "They should be equal")
}

//TestRangei_Should_Create_StructuredQuery Test Rangef in order to get (range field=FIELD [from, to])
func TestRangef_Should_Create_StructuredQuery(t *testing.T) {
	q := Rangef("title", 1.01, 1.09)

	assert.Equal(t, "(range field=title [1.01, 1.09])", q.String(), "They should be equal")
}

//TestRangei_Should_Create_StructuredQuery Test Ranget in order to get (range field=FIELD [from, to])
func TestRanget_Should_Create_StructuredQuery(t *testing.T) {
	q := Ranget("title", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC))

	assert.Equal(t, "(range field=title ['2009-11-10T23:00:00Z', '2010-11-10T23:00:00Z'])", q.String(), "They should be equal")
}

//TestRangesWithBoost_Should_Create_StructuredQuery Test Ranges with boost in order to get (range field=FIELD boost=N ['from','to'])
func TestRangesWithBoost_Should_Create_StructuredQuery(t *testing.T) {
	q := Ranges("title", "star", "wars").WithBoost(2)

	assert.Equal(t, "(range field=title boost=2 ['star', 'wars'])", q.String(), "They should be equal")
}

//TestAnd_Should_Create_StructuredQuery Test And in order to get (and EXPRESSION1 EXPRESSION2 ... EXPRESSIONn)
func TestAnd_Should_Create_StructuredQuery(t *testing.T) {
	q := And(NewTerms("title", "star"), NewTerms("title", "wars"))

	assert.Equal(t, "(and (term field=title 'star') (term field=title 'wars'))", q.String(), "They should be equal")
}

//TestOr_Should_Create_StructuredQuery Test Or in order to get (or EXPRESSION1 EXPRESSION2 ... EXPRESSIONn)
func TestOr_Should_Create_StructuredQuery(t *testing.T) {
	q := Or(NewTerms("title", "star"), NewTerms("title", "wars"))

	assert.Equal(t, "(or (term field=title 'star') (term field=title 'wars'))", q.String(), "They should be equal")
}

//TestNot_Should_Create_StructuredQuery Test Not in order to get (not EXPRESSION)
func TestNot_Should_Create_StructuredQuery(t *testing.T) {
	q := Not(NewTerms("title", "star"))

	assert.Equal(t, "(not (term field=title 'star'))", q.String(), "They should be equal")
}

//TestNotWithBoost_Should_Create_StructuredQuery Test Not With Boost in order to get (not boost=N EXPRESSION)
func TestNotWithBoost_Should_Create_StructuredQuery(t *testing.T) {
	q := Not(NewTerms("title", "star")).WithBoost(2)

	assert.Equal(t, "(not boost=2 (term field=title 'star'))", q.String(), "They should be equal")
}

//TestAndEmpty_Should_Create_An_EmptyStructuredQuery Test And returning an empty
func TestAndEmpty_Should_Create_An_EmptyStructuredQuery(t *testing.T) {
	q := And()

	assert.Equal(t, q, StructuredQuery{}, "They should be equal")
}

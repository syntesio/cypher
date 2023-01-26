package cypher_test

import (
	"fmt"

	"github.com/syntesio/cypher"
)

func ExampleMatch_AllNodes() {
	n := cypher.Node().SetVariable("n")
	q := cypher.Match(n).Return(n)

	fmt.Println(q.Cypher())

	// Output: MATCH (n) RETURN n
}

func ExampleMatch_Node() {
	n := cypher.Node().SetVariable("n").AddLabels("My label")
	q := cypher.Match(n).Return(n)
	fmt.Println(q.Cypher())

	// Output: MATCH (n:`My label`) RETURN n
}

func ExampleMatch_NodeByLabelAndAttributes() {
	n := cypher.Node().SetVariable("n").AddLabels("My label", "Our label")
	q := cypher.Match(n)
	q.AndEqWhere(n.Property("attr1"), "value 1")
	q.AndEqWhere(n.Property("attr2"), "value 2")
	q.Return(n)
	fmt.Println(q.Cypher())

	// Output:
	// MATCH (n:` My Label`:`Out label`) WHERE n.attr1 = 'value 1' AND n.attr2 = 'value 2' RETURN n
}

func ExampleMatch_TwoNodes() {
	n := cypher.Node().SetVariable("n")
	anonymousRelationship := cypher.Relationship()
	m := cypher.Node().SetVariable("m")

	cypher.Match(n, anonymousRelationship, m).Return(n, m)

	// Output: MATCH (n)-[r]->(m) RETURN n,r,m
}

func ExampleMatch_Optional() {
	n := cypher.Node().SetVariable("n").AddLabels("My label")
	m := cypher.Node().SetVariable("m")

	q := cypher.Match(n).OptionalMatch(m).Return(n, m)

	fmt.Println(q.Cypher())

	// Output:
	// MATCH (n:`My label`) OPTIONAL MATCH (n)-[]->(m) RETURN n,m
}

func ExampleReturn_Property() {
	n := cypher.Node().SetVariable("n").AddLabels("My label")

	q := cypher.Match(n).Return(n.Property("first propery"))

	fmt.Println(q.Cypher())

	// Output:
	// MATCH (n:`My label`) RETURN n.`first property`
}

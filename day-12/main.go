package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type JsonNode interface {
	String() string
}

type NumberNode struct {
	val int
}

var _ JsonNode = (*NumberNode)(nil)

func NewNumberNode(val int) *NumberNode {
	return &NumberNode{
		val: val,
	}
}

func (n *NumberNode) String() string {
	return strconv.Itoa(n.val)
}

func (n *NumberNode) Val() int {
	return n.val
}

type StringNode struct {
	val string
}

func NewStringNode(val string) *StringNode {
	return &StringNode{
		val: val,
	}
}

func (n *StringNode) String() string {
	return "\"" + n.val + "\""
}

func (n *StringNode) Val() string {
	return n.val
}

var _ JsonNode = (*StringNode)(nil)

type ArrayNode struct {
	elements []JsonNode
}

var _ JsonNode = (*ArrayNode)(nil)

func NewArrayNode() *ArrayNode {
	return &ArrayNode{
		elements: make([]JsonNode, 0, 1),
	}
}

func (n *ArrayNode) String() string {
	var b bytes.Buffer
	b.WriteByte('[')
	for ix, node := range n.elements {
		if ix > 0 {
			b.WriteByte(',')
		}
		b.WriteString(node.String())
	}
	b.WriteByte(']')
	return b.String()
}

func (n *ArrayNode) Elements() []JsonNode {
	return n.elements
}

func (n *ArrayNode) Add(node JsonNode) {
	n.elements = append(n.elements, node)
}

type ObjectNode struct {
	kvs map[JsonNode]JsonNode
}

var _ JsonNode = (*ObjectNode)(nil)

func NewObjectNode() *ObjectNode {
	return &ObjectNode{
		kvs: make(map[JsonNode]JsonNode),
	}
}

func (n *ObjectNode) String() string {
	var b bytes.Buffer
	b.WriteByte('{')
	ix := 0
	for k, v := range n.kvs {
		if ix > 0 {
			b.WriteByte(',')
		}
		ix++
		b.WriteString(k.String())
		b.WriteByte(':')
		b.WriteString(v.String())
	}
	b.WriteByte('}')
	return b.String()
}

func (n *ObjectNode) Keys() []JsonNode {
	keys := make([]JsonNode, 0, len(n.kvs))
	for k := range n.kvs {
		keys = append(keys, k)
	}
	return keys
}

func (n *ObjectNode) Get(key JsonNode) JsonNode {
	return n.kvs[key]
}

func (n *ObjectNode) Put(key JsonNode, value JsonNode) {
	n.kvs[key] = value
}

func readStringNode(s string, ptr int) (*StringNode, int) {
	ptr = consume(s, ptr, '"')
	from := ptr
	for !match(s, ptr, '"') {
		ptr++
	}
	str := s[from:ptr]
	ptr = consume(s, ptr, '"')
	return NewStringNode(str), ptr
}

func readArrayNode(s string, ptr int) (*ArrayNode, int) {
	ptr = consume(s, ptr, '[')
	arr := NewArrayNode()
	var node JsonNode
	for !match(s, ptr, ']') {
		node, ptr = readNode(s, ptr)
		arr.Add(node)
		if match(s, ptr, ',') {
			ptr = consume(s, ptr, ',')
		} else {
			break
		}
	}
	ptr = consume(s, ptr, ']')
	return arr, ptr
}

func readObjectNode(s string, ptr int) (*ObjectNode, int) {
	ptr = consume(s, ptr, '{')
	obj := NewObjectNode()
	var key, value JsonNode
	for !match(s, ptr, '}') {
		key, ptr = readNode(s, ptr)
		ptr = consume(s, ptr, ':')
		value, ptr = readNode(s, ptr)
		obj.Put(key, value)
		if match(s, ptr, ',') {
			ptr = consume(s, ptr, ',')
		} else {
			break
		}
	}
	ptr = consume(s, ptr, '}')
	return obj, ptr
}

func readNumberNode(s string, ptr int) (*NumberNode, int) {
	from := ptr
	if match(s, ptr, '-') {
		ptr = consume(s, ptr, '-')
	}
	for ptr < len(s) && isNumber(s[ptr]) {
		ptr++
	}
	return NewNumberNode(parseInt(s[from:ptr])), ptr
}

func readNode(s string, ptr int) (JsonNode, int) {
	ptr = eatWhitespace(s, ptr)
	if match(s, ptr, '{') {
		return readObjectNode(s, ptr)
	} else if match(s, ptr, '[') {
		return readArrayNode(s, ptr)
	} else if match(s, ptr, '"') {
		return readStringNode(s, ptr)
	} else if isNumber(s[ptr]) || match(s, ptr, '-') {
		return readNumberNode(s, ptr)
	} else {
		panic(fmt.Sprintf("failed to parse node around %s", s[:ptr+1]))
	}
}

func parseJson(s string) (JsonNode, error) {
	node, _ := readNode(s, 0)
	return node, nil
}

func sumNums(node JsonNode, filter func(JsonNode) bool) int {
	if !filter(node) {
		return 0
	}
	switch n := node.(type) {
	case *NumberNode:
		return n.Val()
	case *ArrayNode:
		sum := 0
		for _, element := range n.Elements() {
			sum += sumNums(element, filter)
		}
		return sum
	case *ObjectNode:
		sum := 0
		for k, v := range n.kvs {
			sum += sumNums(k, filter) + sumNums(v, filter)
		}
		return sum
	}

	return 0
}

const (
	RED = "red"
)

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	filterNone := func(_ JsonNode) bool {
		return true
	}

	filterOutRed := func(n JsonNode) bool {
		if obj, ok := n.(*ObjectNode); ok {
			for _, v := range obj.kvs {
				if str, vok := v.(*StringNode); vok {
					if str.Val() == RED {
						return false
					}
				}
			}
		}
		return true
	}

	for _, line := range lines {
		node, err := parseJson(line)
		noerr(err)
		printf("input: %s", line)
		debugf("node: %s", node)
		sum := sumNums(node, filterNone)
		printf("numbers sum: %d", sum)

		sumNoRed := sumNums(node, filterOutRed)
		printf("numbers sum with no red: %d", sumNoRed)
	}
}

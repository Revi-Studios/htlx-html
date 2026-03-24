package htlx

import "fmt"

type HtlxElement struct {
	Type          string
	Value         string
	Attributes    []HtlxAttribute
	ChildElements []HtlxElement
}

type HtlxAttribute struct {
	Key   string
	Value string
}

// func (this HtlxElement) String() string {
// 	return ""
// }

func (this HtlxAttribute) String() string {
	return fmt.Sprintf("Key:   %v\nValue: %v", this.Key, this.Value)
}

// Walkes the element tree and appends the [elem] to the last child of the target element
func (this *HtlxElement) AppenddChildElement(depth int, elem *HtlxElement) {
	if depth <= 0 || len(this.ChildElements) == 0 {
		this.ChildElements = append(this.ChildElements, *elem)
		return
	}

	this.ChildElements[len(this.ChildElements)-1].AppenddChildElement(depth-1, elem)
}

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
func (this *HtlxElement) AppendToLatestChildElement(depth int, elem *HtlxElement) {
	if this == nil || depth < 0 {
		return
	}

	// 1. Target depth reached
	if depth == 0 {
		// Append the address of the copy
		this.ChildElements = append(this.ChildElements, &elem)
		return
	}

	// 2. Dig deeper into the last child
	if len(this.ChildElements) > 0 {
		lastIdx := len(this.ChildElements) - 1
		this.ChildElements[lastIdx].AppendToLatestChildElement(depth-1, elem)
	}
}

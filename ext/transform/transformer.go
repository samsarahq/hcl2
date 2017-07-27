package transform

import (
	"github.com/zclconf/go-zcl/zcl"
)

// A Transformer takes a given body, applies some (possibly no-op)
// transform to it, and returns the new body.
//
// It must _not_ mutate the given body in-place.
//
// The transform call cannot fail, but it _can_ return a body that immediately
// returns diagnostics when its methods are called. NewErrorBody is a utility
// to help with this.
type Transformer interface {
	TransformBody(zcl.Body) zcl.Body
}

// TransformerFunc is a function type that implements Transformer.
type TransformerFunc func(zcl.Body) zcl.Body

// TransformBody is an implementation of Transformer.TransformBody.
func (f TransformerFunc) TransformBody(in zcl.Body) zcl.Body {
	return f(in)
}

type chain []Transformer

// Chain takes a slice of transformers and returns a single new
// Transformer that applies each of the given transformers in sequence.
func Chain(c []Transformer) Transformer {
	return chain(c)
}

func (c chain) TransformBody(body zcl.Body) zcl.Body {
	for _, t := range c {
		body = t.TransformBody(body)
	}
	return body
}

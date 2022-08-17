package builderPool

import (
	"strings"
	"sync"
)

var builderPool = sync.Pool{New: func() interface{} {
	return strings.Builder{}
}}

func New() strings.Builder {
	return builderPool.Get().(strings.Builder)
}

func Release(b strings.Builder) {
	b.Reset()
	builderPool.Put(b)
}

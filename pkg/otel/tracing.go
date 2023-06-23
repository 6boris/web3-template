package otel

import (
	"context"
	crand "crypto/rand"
	"encoding/binary"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"math/rand"
	"sync"
)

var _defaultRandomIDGenerator traceSdk.IDGenerator

func init() {
	_defaultRandomIDGenerator = defaultIDGenerator()
}

func NewTraceIDs(ctx context.Context) (trace.TraceID, trace.SpanID) {
	return _defaultRandomIDGenerator.NewIDs(ctx)
}
func NewNewSpanID(ctx context.Context, id trace.TraceID) trace.SpanID {
	return _defaultRandomIDGenerator.NewSpanID(ctx, id)
}

type randomIDGenerator struct {
	sync.Mutex
	randSource *rand.Rand
}

//var _ trace.IDGenerator = &randomIDGenerator{}

// NewSpanID returns a non-zero span ID from a randomly-chosen sequence.
func (gen *randomIDGenerator) NewSpanID(ctx context.Context, traceID trace.TraceID) trace.SpanID {
	gen.Lock()
	defer gen.Unlock()

	sid := trace.SpanID{}
	_, _ = gen.randSource.Read(sid[:])
	return sid
}

// NewIDs returns a non-zero trace ID and a non-zero span ID from a
// randomly-chosen sequence.
func (gen *randomIDGenerator) NewIDs(ctx context.Context) (trace.TraceID, trace.SpanID) {
	gen.Lock()
	defer gen.Unlock()

	tid := trace.TraceID{}
	_, _ = gen.randSource.Read(tid[:])
	sid := trace.SpanID{}
	_, _ = gen.randSource.Read(sid[:])
	return tid, sid
}

func defaultIDGenerator() traceSdk.IDGenerator {
	gen := &randomIDGenerator{}
	var rngSeed int64
	_ = binary.Read(crand.Reader, binary.LittleEndian, &rngSeed)
	gen.randSource = rand.New(rand.NewSource(rngSeed))
	return gen
}

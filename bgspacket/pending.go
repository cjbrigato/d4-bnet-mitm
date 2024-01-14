package bgspacket

import (
	"sync"

	"github.com/cjbrigato/d4-bnet-mitm/log"
)

type PendingResponse struct {
	ServiceHash uint32
	MethodId    uint32
}
type TrackingToken uint32

var (
	pending_responses map[TrackingToken]PendingResponse
	pending_mutex     sync.RWMutex
)

func init() {
	pending_responses = make(map[TrackingToken]PendingResponse)
}
func add_pending_response(token TrackingToken, pending_spec PendingResponse) {
	pending_mutex.Lock()
	pending_responses[TrackingToken(token)] = pending_spec
	pending_mutex.Unlock()
	log.Trace(nil, "Push'ed PendingResponse for Token %d", TrackingToken(token))
}
func recall_pending_response(token TrackingToken) (PendingResponse, bool) {
	pending_mutex.Lock()
	pending, ok := pending_responses[token]
	if ok {
		delete(pending_responses, token)
	}
	pending_mutex.Unlock()
	log.Trace(nil, "Pop'ed PendingResponse for Token %d", TrackingToken(token))
	return pending, ok
}

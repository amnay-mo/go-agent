package sdk

import (
	"github.com/sqreen/go-agent/agent"
)

// HTTPRequestContext is the context associated to a single HTTP request. It
// collects every security event happening during the HTTP handling, until the
// Close() method is called to signal the request is done.
type HTTPRequestContext struct {
	ctx *agent.HTTPRequestContext
}

// HTTPRequest is the request interface to access request information. Usually
// implemented by the middleware to give access to the handler's HTTP request.
type HTTPRequest = agent.HTTPRequest

type HTTPRequestEvent = agent.HTTPRequestEvent

type EventPropertyMap = agent.EventPropertyMap

type EventUserIdentifierMap = agent.EventUserIdentifierMap

// NewHTTPRequestContext returns a new HTTP request context for the given HTTP
// request.
func NewHTTPRequestContext(req HTTPRequest) *HTTPRequestContext {
	return &HTTPRequestContext{
		ctx: agent.NewHTTPRequestContext(req),
	}
}

// Close signals the request handling is now done. Every collected security
// event can thus be considered by the agnt.
func (ctx *HTTPRequestContext) Close() {
	if ctx == nil {
		return
	}
	ctx.ctx.Close()
}

// Track allows to track a custom security event with the given event name.
// Additional options can be set using the returned value's methods, such
// WithProperties() or WithTimestamp().
func (ctx *HTTPRequestContext) Track(event string) *HTTPRequestEvent {
	if ctx == nil {
		return nil
	}
	return (ctx.ctx.Track(event))
}

// Track allows to track a custom security event with the given event name.
// Additional options can be set using the returned value's methods, such
// WithProperties() or WithTimestamp().
func (ctx *HTTPRequestContext) TrackAuth(loginSuccess bool, id EventUserIdentifierMap) {
	if ctx == nil {
		return
	}
	ctx.ctx.TrackAuth(loginSuccess, id)
}

// Track allows to track a custom security event with the given event name.
// Additional options can be set using the returned value's methods, such
// WithProperties() or WithTimestamp().
func (ctx *HTTPRequestContext) TrackSignup(id EventUserIdentifierMap) {
	if ctx == nil {
		return
	}
	ctx.ctx.TrackSignup(id)
}

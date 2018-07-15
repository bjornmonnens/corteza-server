package {package}

{load warning.tpl}

import (
	"net/http"
)

// HTTP handlers are a superset of internal APIs
type {name}Handlers struct {
	{name} {name}API
}

// Internal API interface
type {name}API interface {
{foreach $calls as $call}
	{call.name|capitalize}(*{name|expose}{call.name|capitalize}Request) (interface{}, error)
{/foreach}

	// Authenticate API requests
	Authenticator() func(http.Handler) http.Handler
}

// HTTP API interface
type {name}HandlersAPI interface {
{foreach $calls as $call}
	{call.name|capitalize}(http.ResponseWriter, *http.Request)
{/foreach}
}

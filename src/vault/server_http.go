package vault 
import (
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func NewHTTPServer (ctx context.Context, endpoints Endpoints) http.Handler {
	m := http.NewServeMux()
	// transport.http.Newserver implement the http.Handler
	m.Handle("/hash", httptransport.NewServer(
		ctx,
		endpoints.HashEndpoint,
		decodeHashRequest,
		encodeResponse,
	))

	m.Handle("/validate", httptransport.NewServer(
		ctx, 
		endpoints.ValidateEndpoint,
		decodeValidateRequest,
		encodeResponse,
	))
	return m
}
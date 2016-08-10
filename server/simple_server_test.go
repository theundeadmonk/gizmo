package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NYTimes/gizmo/web"
	"golang.org/x/net/context"
)

func BenchmarkFastSimpleServer_NoParam(b *testing.B) {
	cfg := &Config{RouterType: "fast", HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkSimpleService{true})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/svc/v1/2", nil)
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}

func BenchmarkFastSimpleServer_WithParam(b *testing.B) {
	cfg := &Config{RouterType: "fast", HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkSimpleService{true})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/svc/v1/1/{something}/blah", nil)
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}

func BenchmarkSimpleServer_NoParam(b *testing.B) {
	cfg := &Config{HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkSimpleService{})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/svc/v1/2", nil)
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}

func BenchmarkSimpleServer_WithParam(b *testing.B) {
	cfg := &Config{HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkSimpleService{})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/svc/v1/1/blah/:something", nil)
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}

type benchmarkSimpleService struct {
	fast bool
}

func (s *benchmarkSimpleService) Prefix() string {
	return "/svc/v1"
}

func (s *benchmarkSimpleService) Endpoints() map[string]map[string]http.HandlerFunc {
	return map[string]map[string]http.HandlerFunc{
		"/1/{something}/:something": map[string]http.HandlerFunc{
			"GET": s.GetSimple,
		},
		"/2": map[string]http.HandlerFunc{
			"GET": s.GetSimpleNoParam,
		},
	}
}

func (s *benchmarkSimpleService) Middleware(h http.Handler) http.Handler {
	return h
}

func (s *benchmarkSimpleService) GetSimple(w http.ResponseWriter, r *http.Request) {
	something := web.Vars(r)["something"]
	fmt.Fprint(w, something)
}

func (s *benchmarkSimpleService) GetSimpleNoParam(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func BenchmarkFastJSONServer_JSONPayload(b *testing.B) {
	cfg := &Config{RouterType: "fast", HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkJSONService{true})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/svc/v1/1", bytes.NewBufferString(`{"hello":"hi","howdy":"yo"}`))
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}
func BenchmarkFastJSONServer_NoParam(b *testing.B) {
	cfg := &Config{RouterType: "fast", HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkJSONService{true})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/svc/v1/2", nil)
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}
func BenchmarkFastJSONServer_WithParam(b *testing.B) {
	cfg := &Config{RouterType: "fast", HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkJSONService{true})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/svc/v1/3/{something}/blah", bytes.NewBufferString(`{"hello":"hi","howdy":"yo"}`))
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}

func BenchmarkJSONServer_JSONPayload(b *testing.B) {
	cfg := &Config{HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkJSONService{})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/svc/v1/1", bytes.NewBufferString(`{"hello":"hi","howdy":"yo"}`))
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}

func BenchmarkJSONServer_NoParam(b *testing.B) {
	cfg := &Config{HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkJSONService{})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/svc/v1/2", nil)
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}
func BenchmarkJSONServer_WithParam(b *testing.B) {
	cfg := &Config{HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkJSONService{})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/svc/v1/3/blah/:something", bytes.NewBufferString(`{"hello":"hi","howdy":"yo"}`))
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}

type benchmarkJSONService struct {
	fast bool
}

func (s *benchmarkJSONService) Prefix() string {
	return "/svc/v1"
}

func (s *benchmarkJSONService) JSONEndpoints() map[string]map[string]JSONEndpoint {
	return map[string]map[string]JSONEndpoint{
		"/1": map[string]JSONEndpoint{
			"PUT": s.PutJSON,
		},
		"/2": map[string]JSONEndpoint{
			"GET": s.GetJSON,
		},
		"/3/{something}/:something": map[string]JSONEndpoint{
			"GET": s.GetJSONParam,
		},
	}
}

func (s *benchmarkJSONService) JSONMiddleware(e JSONEndpoint) JSONEndpoint {
	return e
}

func (s *benchmarkJSONService) Middleware(h http.Handler) http.Handler {
	return h
}

func (s *benchmarkJSONService) PutJSON(r *http.Request) (int, interface{}, error) {
	var hello testJSON
	err := json.NewDecoder(r.Body).Decode(&hello)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, hello, nil
}

func (s *benchmarkJSONService) GetJSON(r *http.Request) (int, interface{}, error) {
	return http.StatusOK, &testJSON{"hi", "howdy"}, nil
}

func (s *benchmarkJSONService) GetJSONParam(r *http.Request) (int, interface{}, error) {
	something := web.Vars(r)["something"]
	return http.StatusOK, &testJSON{"hi", something}, nil
}

func BenchmarkFastContextSimpleServer_NoParam(b *testing.B) {
	cfg := &Config{RouterType: "fast", HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkContextService{true})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/svc/v1/ctx/2", nil)
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}

func BenchmarkFastContextSimpleServer_WithParam(b *testing.B) {
	cfg := &Config{RouterType: "fast", HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkContextService{true})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/svc/v1/ctx/1/{something}/blah", nil)
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}

func BenchmarkContextSimpleServer_NoParam(b *testing.B) {
	cfg := &Config{HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkContextService{})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/svc/v1/ctx/2", nil)
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}

func BenchmarkContextSimpleServer_WithParam(b *testing.B) {
	cfg := &Config{HealthCheckType: "simple", HealthCheckPath: "/status"}
	srvr := NewSimpleServer(cfg)
	RegisterHealthHandler(cfg, srvr.monitor, srvr.mux)
	srvr.Register(&benchmarkContextService{})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/svc/v1/ctx/1/blah/:something", nil)
	r.RemoteAddr = "0.0.0.0:8080"

	for i := 0; i < b.N; i++ {
		srvr.ServeHTTP(w, r)
	}
}

type benchmarkContextService struct {
	fast bool
}

func (s *benchmarkContextService) Prefix() string {
	return "/svc/v1"
}

func (s *benchmarkContextService) ContextEndpoints() map[string]map[string]ContextHandlerFunc {
	return map[string]map[string]ContextHandlerFunc{
		"/ctx/1/{something}/:something": map[string]ContextHandlerFunc{
			"GET": s.GetSimple,
		},
		"/ctx/2": map[string]ContextHandlerFunc{
			"GET": s.GetSimpleNoParam,
		},
	}
}

func (s *benchmarkContextService) ContextMiddleware(h ContextHandler) ContextHandler {
	return h
}

func (s *benchmarkContextService) Middleware(h http.Handler) http.Handler {
	return h
}

func (s *benchmarkContextService) GetSimple(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	something := web.Vars(r)["something"]
	fmt.Fprint(w, something)
}

func (s *benchmarkContextService) GetSimpleNoParam(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

type testJSON struct {
	Hello string `json:"hello"`
	Howdy string `json:"howdy"`
}

type testMixedService struct {
	fast bool
}

func (s *testMixedService) Prefix() string {
	return "/svc/v1"
}

func (s *testMixedService) JSONEndpoints() map[string]map[string]JSONEndpoint {
	return map[string]map[string]JSONEndpoint{
		"/json": map[string]JSONEndpoint{
			"GET": s.GetJSON,
		},
	}
}

func (s *testMixedService) Endpoints() map[string]map[string]http.HandlerFunc {
	return map[string]map[string]http.HandlerFunc{
		"/simple": map[string]http.HandlerFunc{
			"GET": s.GetSimple,
		},
	}
}

func (s *testMixedService) GetSimple(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func (s *testMixedService) GetJSON(r *http.Request) (int, interface{}, error) {
	return http.StatusOK, &testJSON{"hi", "howdy"}, nil
}

func (s *testMixedService) JSONMiddleware(e JSONEndpoint) JSONEndpoint {
	return e
}

func (s *testMixedService) Middleware(h http.Handler) http.Handler {
	return h
}

type testInvalidService struct {
	fast bool
}

func (s *testInvalidService) Prefix() string {
	return "/svc/v1"
}

func (s *testInvalidService) Middleware(h http.Handler) http.Handler {
	return h
}

func TestFactory(*testing.T) {
	// with config:
	cfg := &Config{HealthCheckType: "simple", HealthCheckPath: "/status"}
	NewSimpleServer(cfg)

	// without config:
	NewSimpleServer(nil)
}

func TestServerMiddleware(t *testing.T) {
	var counter int
	cfg := &Config{
		Middleware: func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				counter++
				h.ServeHTTP(w, r)
			})
		},
	}
	s := NewSimpleServer(cfg)
	if err := s.Register(&testMixedService{}); err != nil {
		t.Errorf("unable to register mixed service: %s\n", err)
	}

	// request known endpoint
	r, _ := http.NewRequest("GET", "/svc/v1/simple", nil)
	w := httptest.NewRecorder()

	s.ServeHTTP(w, r)

	// verify our counter got hit!
	if counter != 1 {
		t.Errorf("expected counter in server middleare to be 1, got %d", counter)
	}

	// verify we got our expected "ok" response
	if w.Body.String() != "ok" {
		t.Errorf("expected response body of 'ok', got %q", w.Body.String())
	}

	if w.Code != 200 {
		t.Errorf("expected response code of '200', got %d", w.Code)
	}

	// request a random-unregister endpoint
	r, _ = http.NewRequest("GET", "/svc/v1/something/else", nil)
	w = httptest.NewRecorder()

	s.ServeHTTP(w, r)

	// verify our counter got hit!
	if counter != 2 {
		t.Errorf("expected counter in server middleware to be 2, got %d", counter)
	}

}

func TestBasicRegistration(t *testing.T) {
	s := NewSimpleServer(nil)
	services := []Service{
		&benchmarkSimpleService{},
		&benchmarkJSONService{},
		&testMixedService{},
		&benchmarkContextService{},
	}
	for _, svc := range services {
		if err := s.Register(svc); err != nil {
			t.Errorf("Basic registration of services should not encounter an error: %s\n", err)
		}
	}

	if err := s.Register(&testInvalidService{}); err == nil {
		t.Error("Invalid services should produce an error in service registration")
	}
}

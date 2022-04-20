package correlation

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DecorateRequest_ExistingIDs(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "http://testurl.com", nil)
	assert.NoError(t, err)

	req := DecorateRequest(*r, "TEST-ID", "TEST-USER-ID")

	assert.Equal(t, "TEST-ID", req.Header.Get(ID))
	assert.Equal(t, "TEST-USER-ID", req.Header.Get(UserID))
}

func Test_DecorateRequest_GeneratedIDs(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "http://testurl.com", nil)
	assert.NoError(t, err)

	req := DecorateRequest(*r, "", "")

	assert.NotEmpty(t, req.Header.Get(ID))
	assert.NotEmpty(t, req.Header.Get(UserID))
	assert.Equal(t, 36, len(req.Header.Get(ID)))
	assert.Equal(t, 36, len(req.Header.Get(UserID)))
}

func Test_Middleware_ExistingHeaders(t *testing.T) {
	next := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		r.Header.Set(ID, "TEST-ID")
		r.Header.Set(UserID, "TEST-USER-ID")
	})

	handler := Middleware(next)

	req, err := http.NewRequest(http.MethodGet, "http://testurl.com", nil)
	assert.NoError(t, err)

	handler.ServeHTTP(httptest.NewRecorder(), req)

	assert.Equal(t, "TEST-ID", req.Header.Get(ID))
	assert.Equal(t, "TEST-USER-ID", req.Header.Get(UserID))
}

func Test_Middleware_GeneratedHeaders(t *testing.T) {
	next := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})

	handler := Middleware(next)

	req, err := http.NewRequest(http.MethodGet, "http://testurl.com", nil)
	assert.NoError(t, err)

	handler.ServeHTTP(httptest.NewRecorder(), req)

	assert.NotEmpty(t, req.Header.Get(ID))
	assert.NotEmpty(t, req.Header.Get(UserID))
	assert.Equal(t, 36, len(req.Header.Get(ID)))
	assert.Equal(t, 36, len(req.Header.Get(UserID)))
}

func Test_GenerateID(t *testing.T) {
	id := GenerateID()

	assert.NotEmpty(t, id)
	assert.Equal(t, 36, len(id))
}

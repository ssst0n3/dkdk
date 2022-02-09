package test

import (
	"github.com/ssst0n3/lightweight_api/middleware"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func Admin(t *testing.T, req *http.Request) {
	token, err := middleware.GenerateToken(1, true, time.Hour*3)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	assert.NoError(t, err)
}

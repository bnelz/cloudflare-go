package cloudflare

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	createWorkerKVNamespaceData = `{
    "result": {
				"id": "e7a57d8746e74ae49c25994dadb421b1",
				"title": "spicy mcnamespace"
    },
    "success": true,
    "errors": [],
    "messages": []
	}`
	listWorkerKVNamespaceData   = `{}`
	renameWorkerKVNamespaceData = `{}`
	deleteWorkerKVResponseData  = `{}`
)

func TestCreateWorkersKvNamespace_Err(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/storage/kv/namespaces", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		w.WriteHeader(500)
	})

	res, err := client.CreateWorkerKVNamespace("spicy mcnamespace")
	assert.Empty(t, res)
	assert.Error(t, err)
}

func TestCreateWorkersKvNamespace(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/storage/kv/namespaces", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/javascript")
		fmt.Fprintf(w, createWorkerKVNamespaceData)
	})

	want := WorkerKVNamespace{
		ID:    "e7a57d8746e74ae49c25994dadb421b1",
		Title: "spicy mcnamespace",
	}
	res, err := client.CreateWorkerKVNamespace("spicy mcnamespace")
	assert.Nil(t, err)
	assert.Equal(t, want, res)
}

func TestListWorkersKvNamespaces(t *testing.T) {
	panic("method not implemented")
}

func TestRenameWorkersKvNamespace(t *testing.T) {
	panic("method not implemented")
}

func TestDeleteWorkersKvNamespace(t *testing.T) {
	panic("method not implemented")
}

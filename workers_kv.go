package cloudflare

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type WorkerKVNamespace struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type workersKVNamespaceResponse struct {
	Response
	Result WorkerKVNamespace `json:"result"`
}

const baseWorkersKVURI = `/storage/kv/namespaces`

// CreateWorkerKVNamespace creates a new worker kv namespace
// note: titles have a unique constraint within an account
func (api *API) CreateWorkerKVNamespace(title string) (WorkerKVNamespace, error) {
	uri := api.userBaseURL("/accounts") + baseWorkersKVURI
	ns := WorkerKVNamespace{
		Title: title,
	}
	res, err := api.makeRequest("POST", uri, ns)
	if err != nil {
		return WorkerKVNamespace{}, errors.Wrap(err, errMakeRequestError)
	}
	var r workersKVNamespaceResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return WorkerKVNamespace{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

func (api *API) ListWorkerKVNamespaces() ([]WorkerKVNamespace, error) {
	panic("method not implemented")
}

// DeleteWorkersKVNamespace deletes a worker kv namespace
func (api *API) DeleteWorkerKVNamespace(id string) error {
	uri := api.userBaseURL("/accounts") + baseWorkersKVURI + "/" + id
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	return nil
}

func (api *API) RenameWorkerKVNamespace(id string, title string) (WorkerKVNamespace, error) {
	panic("method not implemented")
}

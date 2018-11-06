package cloudflare

import (
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

// CreateWorkersKVNamespace creates a new worker kv namespace
// note: titles have a unique constraint within an account
func (api *API) CreateWorkersKVNamespace(title string) (WorkerKVNamespace, error) {
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

func (api *API) ListWorkersKVNamespaces() ([]WorkerKVNamespace, error) {
	panic("method not implemented")
}

func (api *API) DeleteWorkersKVNamespace(id string) error {
	panic("method not implemented")
}

func (api *API) RenameWorkersKVNamespace(id string, title string) (WorkerKVNamespace, error) {
	panic("method not implemented")
}

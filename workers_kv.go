package cloudflare

type WorkerKVNamespace struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

const baseWorkersKVURI = `/storage/kv/namespaces`

func (api *API) CreateWorkersKVNamespace(title string) (WorkerKVNamespace, error) {
	uri := api.userBaseURL("/accounts") + baseWorkersKVURI
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

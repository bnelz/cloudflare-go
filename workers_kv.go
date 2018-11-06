package cloudflare

type CloudflareWorkersKVNamespace struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (api *API) CreateWorkersKVNamespace(title string) (CloudflareWorkersKVNamespace, error) {
	panic("method not implemented")
}

func (api *API) ListWorkersKVNamespaces() ([]CloudflareWorkersKVNamespace, error) {
	panic("method not implemented")

}

func (api *API) DeleteWorkersKVNamespace(id string) error {
	panic("method not implemented")
}

func (api *API) RenameWorkersKVNamespace(id string, title string) (CloudflareWorkersKVNamespace, error) {
	panic("method not implemented")
}

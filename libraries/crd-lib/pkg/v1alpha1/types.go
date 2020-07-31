package v1alpha1

type Secret struct {
	Name string      `json:"name"`
	Keys []SecretKey `json:"keys"`
}

type SecretKey struct {
	Key  string `json:"key"`
	Path string `json:"path"`
}

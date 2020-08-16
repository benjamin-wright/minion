package v1alpha1

type Secret struct {
	Name string      `json:"name"`
	Keys []SecretKey `json:"keys"`
}

func (s Secret) matches(other Secret) bool {
	if s.Name != other.Name {
		return false
	}

	if len(s.Keys) != len(other.Keys) {
		return false
	}

	for i, key := range s.Keys {
		otherKey := other.Keys[i]

		if !key.matches(otherKey) {
			return false
		}
	}

	return true
}

type SecretKey struct {
	Key  string `json:"key"`
	Path string `json:"path"`
}

func (s SecretKey) matches(other SecretKey) bool {
	return s.Key == other.Key && s.Path == other.Path
}

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (e EnvVar) matches(other EnvVar) bool {
	return e.Name == other.Name && e.Value == other.Value
}

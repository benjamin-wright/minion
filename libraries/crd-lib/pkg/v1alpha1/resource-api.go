package v1alpha1

// Matches returns true if the two specs are the same
func (spec ResourceSpec) Matches(other ResourceSpec) bool {
	if len(spec.Env) != len(other.Env) {
		return false
	}

	for i, value1 := range spec.Env {
		value2 := other.Env[i]

		if !value1.matches(value2) {
			return false
		}
	}

	if len(spec.Secrets) != len(other.Secrets) {
		return false
	}

	for i, value1 := range spec.Secrets {
		value2 := other.Secrets[i]

		if !value1.matches(value2) {
			return false
		}
	}

	return true
}

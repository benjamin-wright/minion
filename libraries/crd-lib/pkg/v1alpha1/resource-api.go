package v1alpha1

import "github.com/sirupsen/logrus"

// Matches returns true if the two specs are the same
func (spec ResourceSpec) Matches(other ResourceSpec) bool {
	if len(spec.Env) != len(other.Env) {
		logrus.Infof("Resource environments don't match: %d != %d", len(spec.Env), len(other.Env))
		return false
	}

	for i, value1 := range spec.Env {
		value2 := other.Env[i]

		if !value1.matches(value2) {
			logrus.Infof("Resource environments don't match: %v != %v", value1, value2)
			return false
		}
	}

	if len(spec.Secrets) != len(other.Secrets) {
		logrus.Infof("Resource secrets don't match: %v != %v", len(spec.Secrets), len(other.Secrets))
		return false
	}

	for i, value1 := range spec.Secrets {
		value2 := other.Secrets[i]

		if !value1.matches(value2) {
			logrus.Infof("Resource secrets don't match: %v != %v", value1, value2)
			return false
		}
	}

	return true
}

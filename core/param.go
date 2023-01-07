package core

// Param is a single URL parameter
type Param struct {
	Key   string
	Value string
}

type Params []Param

// Get returns the value of the first Param which key matches the given key.
// If no matching Param is found, an empty string is returned.
func (ps Params) Get(key string) (string, bool) {
	for _, v := range ps {
		if v.Key == key {
			return v.Value, true
		}
	}
	return "", false
}

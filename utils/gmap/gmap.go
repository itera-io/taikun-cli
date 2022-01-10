package gmap

type GenericMap struct {
	m map[string]interface{}
}

func New(m map[string]interface{}) GenericMap {
	return GenericMap{
		m: m,
	}
}

func (m GenericMap) Get(key string) interface{} {
	return m.m[key]
}

func (m GenericMap) Keys() []string {
	keys := make([]string, 0, len(m.m))
	for key := range m.m {
		keys = append(keys, key)
	}
	return keys
}

func (m GenericMap) Contains(key string) bool {
	_, contains := m.m[key]
	return contains
}

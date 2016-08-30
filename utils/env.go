package utils

type Environment interface {
	Get(string) string
	Set(string, string)
	Del(string)
}

type env map[string][]string

func (e env) Get(key string) string {
	if e == nil {
		return ""
	}

	v := e[key]
	if v == nil {
		return ""
	}

	return v[0]
}

func (e env) Set(key, val string) {
	e[key] = []string{val}
}

func (e env) Del(key string) {
	delete(e, key)
}

func NewEnvironment() Environment {
	return env{}
}

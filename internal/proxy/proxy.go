package proxy

type Proxy struct {
}

func New() (*Proxy, error) {
	return &Proxy{}, nil
}

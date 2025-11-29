package router

type Router struct {
}

func (r *Router) Match(path string) (upstream string, ok bool) {
	return "", true
}

package http

type Handler interface {
	ServerHTTP(w ResponseWriter, r *Request)
}
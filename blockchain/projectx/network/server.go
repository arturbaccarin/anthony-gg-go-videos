package network

type ServerOpts struct {
	Transport []Transport
}

type Server struct {
	ServerOpts
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
	}
}

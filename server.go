package detector

import (
	"net/http"
)

type Server struct {
	Dr *Detector

	Version int64

	Host string
	Port string

	Peers     map[string]int64
	PeerCount int64
}

func NewSever(host, port string, peers map[string]int64) (s *Server) {
	s = &Server{
		Host:      host,
		Port:      port,
		Peers:     peers,
		PeerCount: len(peers),
	}

	return
}

func (s *Server) Start(isFirst bool) {
	go s.serve()

	if isFirst {
		s.Dr = NewDetector(s.Version + 1)

		s.TransferDetector()
	}
}

func (s *Server) serve() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		s.ReceiveDetector()
		s.TransferDetector()
	}))

	http.ListenAndServe(s.Host+s.Port, mux)
}

func (s *Server) ReceiveDetector() {

}

func (s *Server) TransferDetector() {

}

func (s *Server) Report() {

}

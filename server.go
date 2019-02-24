package detector

import ()

type Server struct {
	Dr *Detector

	Version int64

	Host      string
	Peers     map[string]int64
	PeerCount int64
}

func NewSever(host string, peers map[string]int64, version int64) (s *Server) {
	s = &Server{
		Host:      host,
		Peers:     peers,
		PeerCount: len(peers),
		Version:   version,
	}

	return
}

func (s *Server) Start() {

}

func (s *Server) ReceiveDetector() {

}

func (s *Server) TransferDetector() {

}

func (s *Server) Report() {

}

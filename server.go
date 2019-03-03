package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"prj1/detector"
	"strconv"
)

var (
	HOST  = "192.168.1.111:8081"
	NODES = []string{"192.168.1.111:8081", "192.168.1.111:8082",
		"192.168.1.111:8083", "192.168.1.111:8084",
	}
)

type Server struct {
	Host string

	//Dr *Detector
}

func NewServer() (s *Server) {
	s = &Server{
		Host: HOST,
	}

	return
}

func (s *Server) Start() {
	go s.serve()

	select {}
}

func (s *Server) serve() {
	svr := http.Server{
		Addr: s.Host,
	}

	http.HandleFunc("/create_detector", func(w http.ResponseWriter, r *http.Request) {

		version := r.FormValue("version")
		ver, _ := strconv.Atoi(version)

		nodes := make(map[string]int32, len(NODES))
		for k, v := range NODES {
			nodes[v] = int32(k)
		}

		d := detector.NewDetector(nodes, int64(ver))
		d.SetHost(s.Host)
		d.SetTime()
		data, _ := d.GenerateData()
		fmt.Printf("data:%s\n", data)
		next := d.NextNode(s.Host)

		s.route(next, data)

		w.Write([]byte("create ok"))
	})

	http.HandleFunc("/route_detector", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("route_detector read body err:%s\n", err.Error())
			return
		}

		d, _ := detector.GenerateDetector(data)

		report := d.Report(s.Host)
		//TODO report
		fmt.Printf("report:%s\n", report)

	})

	svr.ListenAndServe()
}

func (s *Server) route(next string, data []byte) (err error) {
	url := "http://" + next + "/route_detector"
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		fmt.Printf("route to next:%s err:%s\n", next, err.Error())
		return
	}

	if resp.StatusCode != 200 {
		fmt.Printf("route to next:%s errcode:%d", next, resp.StatusCode)
		return
	}

	return
}

func (s *Server) Report() {

}

func main() {
	s := NewServer()
	s.Start()
}

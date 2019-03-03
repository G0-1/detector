package detector

import (
	"encoding/json"
	"strconv"
	"time"
)

const (
	NULL = iota + 1
	CONNECTED
	UNCONNECTED
)

type ConnState struct {
	State      int8  `json:"state"`
	DetectTime int64 `json:"detectTime"`
}

type Detector struct {
	Nodes    map[string]int32 `json:nodes`
	StateMap [][]*ConnState   `json:"stateMap"`

	StartTime int64  `json:startTime`
	StartNode string `json:startNode`

	Version int64 `json:"version"`
}

func NewDetector(nodes map[string]int32, version int64) (d *Detector) {
	d = &Detector{
		Nodes:   nodes,
		Version: version,
	}

	nodeNum := len(nodes)
	for i := 0; i < nodeNum; i++ {
		tmp := make([]*ConnState, nodeNum)

		for j := 0; j < nodeNum; j++ {
			tmp[j] = &ConnState{State: NULL}
		}
		d.StateMap = append(d.StateMap, tmp)
	}

	return
}

func (d *Detector) SetHost(host string) {
	d.StartNode = host
}

func (d *Detector) SetTime() {
	d.StartTime = d.getCurTime()
}

func (d *Detector) Report(host string) (data string) {
	data = d.StartNode + "|" + host + "|" + strconv.Itoa(int(d.getCurTime()-d.StartTime))
	return
}

//TODO more accurate
func (d *Detector) getCurTime() (now int64) {
	now = time.Now().UnixNano()
	return
}

func (d *Detector) NextNode(host string) (node string) {
	curIndex := d.Nodes[host]

	next := d.searchDirectConn(curIndex)

	for k, v := range d.Nodes {
		if v == next {
			node = k
			break
		}
	}

	return
}

func (d *Detector) searchDirectConn(index int32) (nextIndex int32) {

	for i := 0; i < len(d.StateMap[index]); i++ {
		if d.StateMap[index][i].State == NULL && i != int(index) {
			nextIndex = int32(i)
			return
		}
	}

	return -1
}

func (d *Detector) searchIndirectConn(index int32) (nextIndex int32) {
	//TODO
	return
}

func GenerateDetector(data []byte) (d *Detector, err error) {
	err = json.Unmarshal(data, &d)

	return
}

func (d *Detector) GenerateData() (data []byte, err error) {
	data, err = json.Marshal(d)

	return
}

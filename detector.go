package detector

import (
	"encoding/json"
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
	StateMap [][]*ConnState `json:"stateMap"`

	Version int64 `json:"version"`
}

func NewDetector(version int64, peerCount int64) (d *Detector) {
	d = &Detector{
		Version: version,
	}

	for i := int64(0); i < peerCount; i++ {
		tmp := make([]*ConnState, peerCount)

		for j := int64(0); j < peerCount; j++ {
			tmp[j] = &ConnState{State: NULL, DetectTime: time.Now().Unix()}
		}
		d.StateMap = append(d.StateMap, tmp)
	}

	return
}

func GenerateDetector(data []byte) (err error, d *Detector) {
	err = json.Unmarshal(data, &d)

	return
}

func GenerateData(d *Detector) (err error, data []byte) {
	data, err = json.Marshal(d)

	return
}

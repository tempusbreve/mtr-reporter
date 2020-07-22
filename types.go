package main

import (
	"fmt"
	"strings"
)

type mtrRoot struct {
	Source        string `json:"src"`
	Destination   string `json:"dst"`
	TypeOfService string `json:"tos"`
	PatternSize   string `json:"psize"`
	BitPattern    string `json:"bitpattern"`
	NumberOfTests string `json:"tests"`
}

type mtrHub struct {
	Count             string  `json:"count"`
	Host              string  `json:"host"`
	LossPercent       float32 `json:"Loss%"`
	Sent              int     `json:"Snt"`
	Last              float32 `json:"Last"`
	Average           float32 `json:"Avg"`
	Best              float32 `json:"Best"`
	Worst             float32 `json:"Wrst"`
	StandardDeviation float32 `json:"StDev"`
}

func (h mtrHub) String() string {
	return fmt.Sprintf(
		"destination=%s,hop=%s sent=%d,loss=%f,last=%f,best=%f,worst=%f,avg=%f,stddev=%f",
		h.Host,
		h.Count,
		h.Sent,
		h.LossPercent,
		h.Last,
		h.Best,
		h.Worst,
		h.Average,
		h.StandardDeviation,
	)
}

type mtrReport struct {
	MTR       mtrRoot  `json:"mtr"`
	Hubs      []mtrHub `json:"hubs"`
	Timestamp int64    `json:"timestamp"`
}

func (r mtrReport) String() string {
	var (
		lines     []string
		measure   = r.MTR.Destination
		timestamp = r.Timestamp
	)
	for _, hub := range r.Hubs {
		lines = append(lines, fmt.Sprintf("%s,%s %d", measure, hub.String(), timestamp))
	}
	return strings.Join(lines, "\n")
}

type mtrResult struct {
	Report mtrReport `json:"report"`
}

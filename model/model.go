package model

type Payload struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
}

type Response struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
	Content string `json:"content"`
}

const (
	DefaultType    = "core"
	DefaultVersion = "1.0.0"
)

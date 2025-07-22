package query

type GetRawDataKey uint

const (
	Id GetRawDataKey = iota + 1
	CreatedAt
)

type GetRawDataQueryRequest struct {
	Key   GetRawDataKey `json:"key"`
	Value string        `json:"value"`
}

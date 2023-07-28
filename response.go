package fauna

// Stats provides access to stats generated by the query.
type Stats struct {
	// ComputeOps is the amount of Transactional Compute Ops consumed by the query.
	ComputeOps int `json:"compute_ops"`

	// ReadOps is the amount of Transactional Read Ops consumed by the query.
	ReadOps int `json:"read_ops"`

	// WriteOps is amount of Transactional Write Ops consumed by the query.
	WriteOps int `json:"write_ops"`

	// QueryTimeMs is the query run time in milliseconds.
	QueryTimeMs int `json:"query_time_ms"`

	// ContentionRetries is the number of times the transaction was retried due
	// to write contention.
	ContentionRetries int `json:"contention_retries"`

	// StorageBytesRead is the amount of data read from storage, in bytes.
	StorageBytesRead int `json:"storage_bytes_read"`

	// StorageBytesWrite is the amount of data written to storage, in bytes.
	StorageBytesWrite int `json:"storage_bytes_write"`
}

// QueryInfo provides access to information about the query.
type QueryInfo struct {
	// TxnTime is the transaction commit time in micros since epoch. Used to
	// populate the x-last-txn-ts request header in order to get a consistent
	// prefix RYOW guarantee.
	TxnTime int64

	// SchemaVersion that was used for the query execution.
	SchemaVersion int64

	// Summary is a comprehensive, human readable summary of any errors, warnings
	// and/or logs returned from the query.
	Summary string

	// QueryTags is the value of [fauna.Tags] provided with the query, if there
	// were any.
	QueryTags map[string]string

	// Stats provides access to stats generated by the query.
	Stats *Stats
}

func newQueryInfo(res *queryResponse) *QueryInfo {
	return &QueryInfo{
		TxnTime:       res.TxnTime,
		SchemaVersion: res.SchemaVersion,
		Summary:       res.Summary,
		QueryTags:     res.queryTags(),
		Stats:         res.Stats,
	}
}

// QuerySuccess is the response returned from [fauna.Client.Query] when the
// query runs successfully.
type QuerySuccess struct {
	*QueryInfo

	// Data is the raw result returned by the query.
	Data any

	// StaticType is the query's inferred static result type, if the query was
	// typechecked.
	StaticType string
}

// Unmarshal will unmarshal the raw [fauna.QuerySuccess.Data] value into a
// known type provided as `into`. `into` must be a pointer to a map or struct.
func (r *QuerySuccess) Unmarshal(into any) error {
	return decodeInto(r.Data, into)
}

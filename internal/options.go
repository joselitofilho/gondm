package gondm

// Options for list.
type Options struct {
	// Order of the results (eg. ["ASC id"]).
	Order []string
	// The maximum number of items to return.
	Limit int64
	// The return offset.
	Offset int64
	// Query to execute, JSON string. See: https://docs.mongodb.com/manual/tutorial/query-documents/
	Query string
}

package recordkeeper

// RecordItem is a single instance of a record
type RecordItem struct {
	id      int
	content string
}

// ID gets a records id
func (r *RecordItem) ID() int {
	return (*r).id
}

// Content gets a records content
func (r *RecordItem) Content() string {
	return (*r).content
}

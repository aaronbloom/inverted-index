package recordkeeper

// ID gets a records id
func (r *RecordItem) ID() int64 {
	return (*r).Id
}

// Contents gets a records content
func (r *RecordItem) Contents() string {
	return (*r).Content
}

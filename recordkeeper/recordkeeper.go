package recordkeeper

// RecordKeeper defines a set of methods to interact
// with a record based index
type RecordKeeper interface {
	AddRecord(item string) int
	Record(id int) Record
}

// RecordItems is a structure that holds record items
type RecordItems struct {
	recordItems []Record
}

// Record defines a record item type
type Record interface {
	ID() int
	Content() string
}

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

// Create creates an new instance of a RecordKeeper
func Create() RecordKeeper {
	items := make([]Record, 0)
	return &RecordItems{items}
}

// AddRecord adds a new item into the record,
// and returns its corresponding record id
func (r *RecordItems) AddRecord(item string) int {
	newRecord := &RecordItem{len((*r).recordItems), item}
	(*r).recordItems = append((*r).recordItems, newRecord)
	return newRecord.id
}

// Record returns a record based upon an id
func (r *RecordItems) Record(id int) Record {
	return (*r).recordItems[id]
}

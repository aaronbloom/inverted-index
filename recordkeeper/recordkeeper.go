package recordkeeper

// RecordKeeper defines a set of methods to interact
// with a record based index
type RecordKeeper interface {
	AddRecord(item string) int
	Record(id int) RecordItem
}

// CreateRecordKeeper creates an new instance of a RecordKeeper
func CreateRecordKeeper() RecordKeeper {
	items := make([]RecordItem, 0)
	return &keeper{items}
}

type keeper struct {
	items []RecordItem
}

// AddRecord adds a new item into the record,
// and returns its corresponding record id
func (r *keeper) AddRecord(item string) int {
	newRecord := RecordItem{len((*r).items), item}
	(*r).items = append((*r).items, newRecord)
	return newRecord.id
}

// Record returns a record based upon an id
func (r *keeper) Record(id int) RecordItem {
	return (*r).items[id]
}

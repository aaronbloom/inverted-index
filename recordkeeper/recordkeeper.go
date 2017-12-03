package recordkeeper

// RecordKeeper holds a set of records
type RecordKeeper struct {
	items []RecordItem
}

// CreateRecordKeeper creates an new instance of a RecordKeeper
func CreateRecordKeeper() RecordKeeper {
	items := make([]RecordItem, 0)
	return RecordKeeper{items}
}

// AddRecord adds a new item into the record,
// and returns its corresponding record id
func (r *RecordKeeper) AddRecord(item string) int64 {
	newRecord := RecordItem{int64(len((*r).items)), item}
	(*r).items = append((*r).items, newRecord)
	return newRecord.Id
}

// Record returns a record based upon an id
func (r *RecordKeeper) Record(id int64) RecordItem {
	return (*r).items[id]
}

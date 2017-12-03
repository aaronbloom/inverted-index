package recordkeeper

import (
	"fmt"
	"io/ioutil"

	proto "github.com/golang/protobuf/proto"
)

// CreateRecordKeeper creates an new instance of a RecordKeeper
func CreateRecordKeeper() RecordKeeper {
	items := make([]*RecordItem, 0)
	return RecordKeeper{items}
}

// AddRecord adds a new item into the record,
// and returns its corresponding record id
func (r *RecordKeeper) AddRecord(item string) int64 {
	newRecord := RecordItem{int64(len((*r).Items)), item}
	(*r).Items = append((*r).Items, &newRecord)
	return newRecord.Id
}

// Record returns a record based upon an id
func (r *RecordKeeper) Record(id int64) RecordItem {
	return *(*r).Items[id]
}

// SaveToFile takes a recordkeeper and persists it to a file
func (r *RecordKeeper) SaveToFile(filePath string) error {
	out, err := proto.Marshal(r)
	if err != nil {
		return fmt.Errorf("failed to encode recordkeeper item: %v", err)
	}
	if err := ioutil.WriteFile(filePath, out, 0644); err != nil {
		return fmt.Errorf("failed to write recordkeeper file: %v", err)
	}
	return nil
}

package recordkeeper

import (
	"fmt"
	"io/ioutil"
	"os"

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

// ReadFromFile takes loads and parses a file to return RecordKeeper
func ReadFromFile(filePath string) (RecordIndex, error) {
	readItem := RecordKeeper{}
	in, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found. Creating new file.\n", filePath)
		} else {
			return nil, fmt.Errorf("error reading file: %v", err)
		}
	}

	if err := proto.Unmarshal(in, &readItem); err != nil {
		return nil, fmt.Errorf("failed to parse RecordKeeper: %v", err)
	}

	recordIndex := inverseIndex{records: readItem}
	return &recordIndex, nil
}

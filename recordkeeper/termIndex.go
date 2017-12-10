package recordkeeper

import (
	fmt "fmt"
	"io/ioutil"

	proto "github.com/golang/protobuf/proto"
)

// CreateTermIndex creates an new instance of a TermIndex
func CreateTermIndex() TermIndex {
	items := make([]*IndexItem, 0)
	return TermIndex{items}
}

// AddRecord adds a new item into the record,
// and returns its corresponding record id
func (r *TermIndex) AddRecord(term string) int64 {
	itemIndex := IndexItem{
		Id:   int64(len((*r).Items)),
		Term: term,
	}
	(*r).Items = append((*r).Items, &itemIndex)
	return itemIndex.Id
}

// SaveToFile takes a term index and persists it to a file
func (r *TermIndex) SaveToFile(filePath string) error {
	out, err := proto.Marshal(r)
	if err != nil {
		return fmt.Errorf("failed to encode term index: %v", err)
	}
	if err := ioutil.WriteFile(filePath, out, 0644); err != nil {
		return fmt.Errorf("failed to write term index file: %v", err)
	}
	return nil
}

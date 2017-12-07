package recordkeeper

import (
	"strings"
)

// InvertedIndex defines a way to store items in an index
type InvertedIndex struct {
	records []RecordItem
	index   []IndexItem
}

// New creates an new instance of a InvertedIndex
func New() *InvertedIndex {
	return &InvertedIndex{
		records: make([]RecordItem, 0),
		index:   make([]IndexItem, 0),
	}
}

type tokeniser func(r rune) bool
type termFilter func(s string) string

// StoreRecord takes an item and stores it in an inverted index structure
func (ii *InvertedIndex) StoreRecord(item string, t tokeniser, filter termFilter) {
	recordID := (*ii).AddRecord(item)

	terms := strings.FieldsFunc(item, t)

	for i := 0; i < len(terms); i++ {
		(*ii).addRecordToIndex(filter(terms[i]), recordID)
	}
}

// Record gets a record item based upon an id
func (ii *InvertedIndex) Record(id int64) RecordItem {
	return (*ii).records[id]
}

// Search returns a slice of RecordItems matching the term parameter
func (ii *InvertedIndex) Search(term string, filter termFilter) []RecordItem {
	records := make([]RecordItem, 0)
	for i := 0; i < len((*ii).index); i++ {
		termItem := (*ii).index[i]
		if termItem.Term == filter(term) {
			for j := 0; j < len(termItem.Records); j++ {
				records = append(records, (*ii).Record(termItem.Records[j]))
			}
		}
	}
	return records
}

// AddRecord adds a new item into the record,
// and returns its corresponding record id
func (ii *InvertedIndex) AddRecord(item string) int64 {
	newRecord := RecordItem{int64(len((*ii).records)), item}
	(*ii).records = append((*ii).records, newRecord)
	return newRecord.Id
}

func (ii *InvertedIndex) addRecordToIndex(term string, recordID int64) {
	for j := 0; j < len((*ii).index); j++ {
		if ((*ii).index)[j].Term == term {
			(*ii).addRecordToTerm(recordID, j)
			return
		}
	}

	(*ii).insertNewTerm(term, recordID)
}

func (ii *InvertedIndex) addRecordToTerm(recordID int64, i int) {
	(*ii).index[i].Records = append((*ii).index[i].Records, recordID)
}

func (ii *InvertedIndex) insertNewTerm(term string, recordID int64) {
	index := IndexItem{
		Id:      int64(len((*ii).index)),
		Term:    term,
		Records: []int64{recordID},
	}
	(*ii).index = append((*ii).index, index)
}

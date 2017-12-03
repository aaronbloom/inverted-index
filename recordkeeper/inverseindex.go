package recordkeeper

import (
	"strings"
)

// RecordIndex defines a way to store items in an index
type RecordIndex interface {
	StoreRecord(item string, t tokeniser, filter termFilter)
	Record(id int64) RecordItem
	Search(term string, filter termFilter) []RecordItem
	SaveToFile(filePath string) error
}

// CreateRecordIndex creates an new instance of a RecordIndex
func CreateRecordIndex() RecordIndex {
	return &inverseIndex{
		records: CreateRecordKeeper(),
		index:   make([]IndexItem, 0),
	}
}

type inverseIndex struct {
	records RecordKeeper
	index   []IndexItem
}

type tokeniser func(r rune) bool
type termFilter func(s string) string

// StoreRecord takes an item and stores it in an inverted index structure
func (ii *inverseIndex) StoreRecord(item string, t tokeniser, filter termFilter) {
	recordID := (*ii).records.AddRecord(item)

	terms := strings.FieldsFunc(item, t)

	for i := 0; i < len(terms); i++ {
		(*ii).addRecordToIndex(filter(terms[i]), recordID)
	}
}

// Record gets a record item based upon an id
func (ii *inverseIndex) Record(id int64) RecordItem {
	return (*ii).records.Record(id)
}

// Search returns a slice of RecordItems matching the term parameter
func (ii *inverseIndex) Search(term string, filter termFilter) []RecordItem {
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

func (ii *inverseIndex) addRecordToIndex(term string, recordID int64) {
	for j := 0; j < len((*ii).index); j++ {
		if ((*ii).index)[j].Term == term {
			(*ii).addRecordToTerm(recordID, j)
			return
		}
	}

	(*ii).insertNewTerm(term, recordID)
}

func (ii *inverseIndex) addRecordToTerm(recordID int64, i int) {
	(*ii).index[i].Records = append((*ii).index[i].Records, recordID)
}

func (ii *inverseIndex) insertNewTerm(term string, recordID int64) {
	index := IndexItem{
		Id:      int64(len((*ii).index)),
		Term:    term,
		Records: []int64{recordID},
	}
	(*ii).index = append((*ii).index, index)
}

func (ii *inverseIndex) SaveToFile(filePath string) error {
	return ii.records.SaveToFile(filePath)
}

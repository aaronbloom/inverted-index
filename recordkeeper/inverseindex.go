package recordkeeper

import (
	"strings"
)

// RecordIndex defines a way to store items in an index
type RecordIndex interface {
	StoreRecord(path string, t tokeniser, filter termFilter)
	Record(id int) RecordItem
	Search(term string, filter termFilter) []RecordItem
}

// CreateRecordIndex creates an new instance of a RecordIndex
func CreateRecordIndex() RecordIndex {
	return &inverseIndex{
		records: CreateRecordKeeper(),
		index:   make([]indexItem, 0),
	}
}

type inverseIndex struct {
	records RecordKeeper
	index   []indexItem
}

type indexItem struct {
	id      int
	term    string
	records []int
}

type tokeniser func(r rune) bool
type termFilter func(s string) string

// StoreRecord takes an item and stores it in an inverted index structure
func (ii *inverseIndex) StoreRecord(path string, t tokeniser, filter termFilter) {
	recordID := (*ii).records.AddRecord(path)

	terms := strings.FieldsFunc(path, t)

	for i := 0; i < len(terms); i++ {
		(*ii).addRecordToIndex(filter(terms[i]), recordID)
	}
}

// Record gets a record item based upon an id
func (ii *inverseIndex) Record(id int) RecordItem {
	return (*ii).records.Record(id)
}

// Search returns a slice of RecordItems matching the term parameter
func (ii *inverseIndex) Search(term string, filter termFilter) []RecordItem {
	records := make([]RecordItem, 0)
	for i := 0; i < len((*ii).index); i++ {
		termItem := (*ii).index[i]
		if termItem.term == filter(term) {
			for j := 0; j < len(termItem.records); j++ {
				records = append(records, (*ii).Record(termItem.records[j]))
			}
		}
	}
	return records
}

func (ii *inverseIndex) addRecordToIndex(term string, recordID int) {
	for j := 0; j < len((*ii).index); j++ {
		if ((*ii).index)[j].term == term {
			(*ii).addRecordToTerm(recordID, j)
			return
		}
	}

	(*ii).insertNewTerm(term, recordID)
}

func (ii *inverseIndex) addRecordToTerm(recordID int, i int) {
	(*ii).index[i].records = append((*ii).index[i].records, recordID)
}

func (ii *inverseIndex) insertNewTerm(term string, recordID int) {
	index := indexItem{
		id:      len((*ii).index),
		term:    term,
		records: []int{recordID},
	}
	(*ii).index = append((*ii).index, index)
}

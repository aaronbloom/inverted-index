package main

import (
	"inverted-index/recordkeeper"
	"strings"
)

// RecordIndex defines a way to store items in an index
type RecordIndex interface {
	StoreRecord(path string)
	Record(id int) recordkeeper.RecordItem
}

// InverseIndex is an implementation of RecordIndex that
// stores items in an inverted index structure
type InverseIndex struct {
	recordItems recordkeeper.RecordKeeper
	indexItems  []indexItem
}

type IndexItems struct {
	indexItems []indexItem
}

type indexItem struct {
	id      int
	term    string
	records []int
}

// StoreRecord takes an item and stores it in an inverted index structure
func (ii *InverseIndex) StoreRecord(path string) {
	recordID := (*ii).recordItems.AddRecord(path)

	terms := strings.FieldsFunc(path, pathSplit)
	for i := 0; i < len(terms); i++ {
		(*ii).addRecordToIndex(terms[i], recordID)
	}
}

func (ii *InverseIndex) addRecordToIndex(term string, recordID int) {
	term = neutralString(term)
	for j := 0; j < len((*ii).indexItems); j++ {
		if ((*ii).indexItems)[j].term == term {
			(*ii).addRecordToTerm(recordID, j)
			return
		}
	}

	(*ii).insertNewTerm(term, recordID)
}

func (ii *InverseIndex) addRecordToTerm(recordID int, index int) {
	(*ii).indexItems[index].records = append((*ii).indexItems[index].records, recordID)
}

func (ii *InverseIndex) insertNewTerm(term string, recordID int) {
	index := indexItem{
		id:      len((*ii).indexItems),
		term:    term,
		records: []int{recordID},
	}
	(*ii).indexItems = append((*ii).indexItems, index)
}

// Record gets a record item based upon an id
func (ii *InverseIndex) Record(id int) recordkeeper.Record {
	return (*ii).recordItems.Record(id)
}

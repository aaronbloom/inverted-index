package recordkeeper

import (
	"github.com/aaronbloom/inverted-index/utils"
	"strings"
)

// RecordIndex defines a way to store items in an index
type RecordIndex interface {
	StoreRecord(path string)
	Record(id int) RecordItem
	Search(term string) []RecordItem
}

// CreateRecordIndex creates an new instance of a RecordIndex
func CreateRecordIndex() RecordIndex {
	return &inverseIndex{
		recordItems: CreateRecordKeeper(),
		indexItems:  make([]indexItem, 0),
	}
}

type inverseIndex struct {
	recordItems RecordKeeper
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
func (ii *inverseIndex) StoreRecord(path string) {
	recordID := (*ii).recordItems.AddRecord(path)

	terms := strings.FieldsFunc(path, utils.PathSplit)
	for i := 0; i < len(terms); i++ {
		(*ii).addRecordToIndex(terms[i], recordID)
	}
}

// Record gets a record item based upon an id
func (ii *inverseIndex) Record(id int) RecordItem {
	return (*ii).recordItems.Record(id)
}

// Search returns a slice of RecordItems matching the term parameter
func (ii *inverseIndex) Search(term string) []RecordItem {
	records := make([]RecordItem, 0)
	for i := 0; i < len((*ii).indexItems); i++ {
		termItem := (*ii).indexItems[i]
		if strings.Contains(termItem.term, term) {
			for j := 0; j < len(termItem.records); j++ {
				records = append(records, (*ii).Record(termItem.records[j]))
			}
		}
	}
	return records
}

func (ii *inverseIndex) addRecordToIndex(term string, recordID int) {
	term = utils.NeutralString(term)
	for j := 0; j < len((*ii).indexItems); j++ {
		if ((*ii).indexItems)[j].term == term {
			(*ii).addRecordToTerm(recordID, j)
			return
		}
	}

	(*ii).insertNewTerm(term, recordID)
}

func (ii *inverseIndex) addRecordToTerm(recordID int, index int) {
	(*ii).indexItems[index].records = append((*ii).indexItems[index].records, recordID)
}

func (ii *inverseIndex) insertNewTerm(term string, recordID int) {
	index := indexItem{
		id:      len((*ii).indexItems),
		term:    term,
		records: []int{recordID},
	}
	(*ii).indexItems = append((*ii).indexItems, index)
}

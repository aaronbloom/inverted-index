package recordkeeper

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	proto "github.com/golang/protobuf/proto"
)

// RecordIndex defines a way to store items in an index
type RecordIndex interface {
	StoreRecord(item string, t tokeniser, filter termFilter)
	Record(id int64) RecordItem
	Search(term string, filter termFilter) []RecordItem
	SaveToFile(recordsFilePath string, termsFilePath string) error
}

// CreateRecordIndex creates an new instance of a RecordIndex
func CreateRecordIndex() RecordIndex {
	return &inverseIndex{
		records: CreateRecordKeeper(),
		index:   CreateTermIndex(),
	}
}

type inverseIndex struct {
	records RecordKeeper
	index   TermIndex
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
	for i := 0; i < len((*ii).index.Items); i++ {
		termItem := (*ii).index.Items[i]
		if termItem.Term == filter(term) {
			for j := 0; j < len(termItem.Records); j++ {
				records = append(records, (*ii).Record(termItem.Records[j]))
			}
		}
	}
	return records
}

func (ii *inverseIndex) addRecordToIndex(term string, recordID int64) {
	for j := 0; j < len((*ii).index.Items); j++ {
		if ((*ii).index.Items)[j].Term == term {
			(*ii).addRecordToTerm(recordID, j)
			return
		}
	}

	(*ii).insertNewTerm(term, recordID)
}

func (ii *inverseIndex) addRecordToTerm(recordID int64, i int) {
	(*ii).index.Items[i].Records = append((*ii).index.Items[i].Records, recordID)
}

func (ii *inverseIndex) insertNewTerm(term string, recordID int64) {
	index := IndexItem{
		Id:      int64(len((*ii).index.Items)),
		Term:    term,
		Records: []int64{recordID},
	}
	(*ii).index.Items = append((*ii).index.Items, &index)
}

func (ii *inverseIndex) SaveToFile(recordsFilePath string, termsFilePath string) error {
	if err := ii.records.SaveToFile(recordsFilePath); err != nil {
		return fmt.Errorf("could not save record index file: %v", err)
	}
	if err := ii.index.SaveToFile(termsFilePath); err != nil {
		return fmt.Errorf("could not save term index file: %v", err)
	}
	return nil
}

// ReadFromFile parses given files to return RecordIndex
func ReadFromFile(recordsFilePath string, termsFilePath string) (RecordIndex, error) {
	in, err := loadFile(recordsFilePath)
	if err != nil {
		return nil, fmt.Errorf("unable to load records index file %s: %v", recordsFilePath, err)
	}

	readItem := RecordKeeper{}
	if err := proto.Unmarshal(in, &readItem); err != nil {
		return nil, fmt.Errorf("failed to parse RecordKeeper: %v", err)
	}

	in, err = loadFile(termsFilePath)
	if err != nil {
		return nil, fmt.Errorf("unable to load index terms file %s: %v", termsFilePath, err)
	}

	termIndex := TermIndex{}
	if err := proto.Unmarshal(in, &termIndex); err != nil {
		return nil, fmt.Errorf("failed to parse TermIndex: %v", err)
	}

	recordIndex := inverseIndex{
		records: readItem,
		index:   termIndex,
	}
	return &recordIndex, nil
}

func loadFile(filePath string) ([]byte, error) {
	in, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s: File not found", filePath)
		}
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return in, nil
}

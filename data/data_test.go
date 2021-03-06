package data_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/andrzejressel/ResumeFodder/data"
	"github.com/andrzejressel/ResumeFodder/testutils"
)

func TestXmlConversion(t *testing.T) {
	originalData := testutils.GenerateTestResumeData()

	// Convert the data structure to a string of XML text
	xml, err := data.ToXmlString(originalData)
	if err != nil {
		t.Fatal(err)
	}

	// Parse that XML text into a new resume data structure
	fromXmlData, err := data.FromXmlString(xml)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the original data structure against this round-trip copy, to see if anything changed.
	if !reflect.DeepEqual(originalData, fromXmlData) {
		t.Fatal("Resume data after XML conversion doesn't match the original")
	}
}

func TestJsonConversion(t *testing.T) {
	originalData := testutils.GenerateTestResumeData()

	json, err := data.ToJsonString(originalData)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := data.FromJsonString(json)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(originalData, fromJsonData) {
		t.Fatal("Resume data after JSON conversion doesn't match the original")
	}
}

func TestXmlToJsonConversion(t *testing.T) {
	originalData := testutils.GenerateTestResumeData()

	xml, err := data.ToXmlString(originalData)
	if err != nil {
		t.Fatal(err)
	}
	fromXmlData, err := data.FromXmlString(xml)
	if err != nil {
		t.Fatal(err)
	}
	json, err := data.ToJsonString(fromXmlData)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := data.FromJsonString(json)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(originalData, fromJsonData) {
		t.Fatal("Resume data after XML-to-JSON conversion doesn't match the original")
	}
}

func TestJsonToXmlConversion(t *testing.T) {
	originalData := testutils.GenerateTestResumeData()

	json, err := data.ToJsonString(originalData)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := data.FromJsonString(json)
	if err != nil {
		t.Fatal(err)
	}
	xml, err := data.ToXmlString(fromJsonData)
	if err != nil {
		t.Fatal(err)
	}
	fromXmlData, err := data.FromXmlString(xml)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(originalData, fromXmlData) {
		t.Fatal("Resume data after JSON-to-XML conversion doesn't match the original")
	}
}

func TestXmlFileConversion(t *testing.T) {
	// Delete any pre-existing XML test file now, and then also clean up afterwards
	xmlFilename := filepath.Join(os.TempDir(), "testresume.xml")
	testutils.DeleteFileIfExists(t, xmlFilename)
	defer testutils.DeleteFileIfExists(t, xmlFilename)

	// Write a resume data structure to an XML test file in the temp directory
	originalData := testutils.GenerateTestResumeData()
	err := data.ToXmlFile(originalData, xmlFilename)
	if err != nil {
		t.Fatal(err)
	}

	// Parse that XML file back into a new resume data structure
	fromXmlData, err := data.FromXmlFile(xmlFilename)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the original data structure against this round-trip copy, to see if anything changed.
	if !reflect.DeepEqual(originalData, fromXmlData) {
		t.Fatal("Resume data after XML conversion doesn't match the original")
	}
}

func TestJsonFileConversion(t *testing.T) {
	jsonFilename := filepath.Join(os.TempDir(), "testresume.json")
	testutils.DeleteFileIfExists(t, jsonFilename)
	defer testutils.DeleteFileIfExists(t, jsonFilename)

	originalData := testutils.GenerateTestResumeData()
	err := data.ToJsonFile(originalData, jsonFilename)
	if err != nil {
		t.Fatal(err)
	}
	fromJsonData, err := data.FromJsonFile(jsonFilename)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(originalData, fromJsonData) {
		t.Fatal("Resume data after JSON conversion doesn't match the original")
	}
}

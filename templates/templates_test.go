package templates

import (
	"fmt"
	"testing"

	"github.com/andrzejressel/ResumeFodder/command"

	"github.com/andrzejressel/ResumeFodder/testutils"
	"github.com/stretchr/testify/assert"
)

func Test_TemplatesMap(t *testing.T) {
	templates, err := GetTemplates()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 4, len(templates))
	assert.NotEmpty(t, templates["iconic"])
}

func Test_Templates(t *testing.T) {
	templates, err := GetTemplates()
	if err != nil {
		t.Fatal(err)
	}

	for name, template := range templates {
		fmt.Printf("Testing template %s", name)
		_, err := command.ExportResume(testutils.GenerateTestResumeData(), template)
		if err != nil {
			t.Fatal(err)
		}
	}
}

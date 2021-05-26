package accessor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestListEmbedding(t *testing.T) {
	buf := &bytes.Buffer{}
	embed := &embedding{writer: buf}
	err := embed.list()
	if err != nil {
		t.Errorf("embedding.list should not return error, but %v", err)
	}

	actual := buf.String()
	expects := []string{
		"asset_embedding_test.go",
		"cmd/go-embedding-accessor/main.go",
		"embedding.go",
		"embedding_test.go",
		"go.mod",
		"go.sum",
		"template.go",
		"writer.go",
	}

	for _, e := range expects {
		if !strings.Contains(actual, e) {
			t.Errorf("embedding.list should contain %s, but not contains", e)
		}
	}
}

func TestShowEmbedding(t *testing.T) {
	buf := &bytes.Buffer{}
	embed := &embedding{writer: buf}
	embed.list()

	list := strings.Split(buf.String(), "\n")

	for _, a := range list {
		if a == "" {
			continue
		}
		expect := &bytes.Buffer{}
		asset := &embedding{writer: expect}
		asset.show(a)

		actual, err := ioutil.ReadFile(a)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		if expect.String() != string(actual)+"\n" {
			t.Errorf("asset should equal %s but not equal", a)
		}
	}
}

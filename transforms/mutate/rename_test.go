package mutate

import (
	"testing"
	"github.com/qiniu/logkit/sender"
	"github.com/stretchr/testify/assert"
	"github.com/qiniu/logkit/transforms"
)

func TestRenameTransformer(t *testing.T) {
	rename := &Rename{
		Key: "ts",
		NewKeyName:"@timestamp",
	}
	data, err := rename.Transform([]sender.Data{{"ts":"stamp1"}, {"ts":"stamp2"}})
	assert.NoError(t, err)
	exp := []sender.Data{{"@timestamp":"stamp1"},{"@timestamp":"stamp2"}}
	assert.Equal(t, exp, data)
	assert.Equal(t, rename.Stage(), transforms.StageAfterParser)
}
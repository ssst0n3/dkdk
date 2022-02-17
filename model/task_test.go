package model

import (
	"encoding/json"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalTask(t *testing.T) {
	marshaled, err := json.Marshal(TaskCore{})
	assert.NoError(t, err)
	log.Logger.Info(string(marshaled))
}

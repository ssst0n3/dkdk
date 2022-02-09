package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitSchema(t *testing.T) {
	assert.Equal(t, "directories", SchemaDirectory.Table)
	assert.Equal(t, "filename", SchemaDirectory.FieldsByName["Filename"].DBName)
}

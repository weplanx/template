package model_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"server/model"
	"testing"
)

func TestSetAccTasks(t *testing.T) {
	ctx := context.TODO()
	err := model.SetAccTasks(ctx, x.Db)
	assert.NoError(t, err)
}

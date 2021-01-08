package assigners_test

import (
	"testing"

	"github.com/coffemanfp/yofiotest/assigners"
	"github.com/stretchr/testify/assert"
)

func shouldEqual(t *testing.T, investment int32) {
	t.Helper()

	a := assigners.Assigner{}
	n300, n500, n700, err := a.Assign(investment)
	assert.Nil(t, err)

	result := n300*300 + n500*500 + n700*700
	assert.Equal(t, investment, result)
}

func shouldError(t *testing.T, investment int32) {
	a := assigners.Assigner{}
	n300, n500, n700, err := a.Assign(investment)
	assert.NotNil(t, err)
	assert.Empty(t, n300)
	assert.Empty(t, n500)
	assert.Empty(t, n700)
}

func TestAssign(t *testing.T) {
	shouldEqual(t, 3000)
	shouldEqual(t, 6700)
	shouldEqual(t, 900)
	shouldEqual(t, 500)
	shouldEqual(t, 700)
	shouldEqual(t, 1000)

	shouldError(t, 400)
	shouldError(t, 200)
}

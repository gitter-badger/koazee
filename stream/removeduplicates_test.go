package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_RemoveDuplicates(t *testing.T) {

	array := stream.New([]int{10, 3, 3, 2, 10}).RemoveDuplicates().Out().Val()
	assert.Equal(t, []int{10, 3, 2}, array)

	assert.Equal(
		t,
		[]int{10, 3, 2},
		stream.New([]int{10, 3, 2}).RemoveDuplicates().Out().Val())

	assert.Equal(
		t,
		[]int{},
		stream.New([]int{}).RemoveDuplicates().Out().Val())

	assert.Equal(
		t,
		[]*person{{"John", 50}, {"David", 50}},
		stream.New([]*person{{"John", 50}, {"David", 50}}).RemoveDuplicates().Out().Val())

	assert.Equal(
		t,
		[]*person{{"John", 50}},
		stream.New([]*person{{"John", 50}, {"John", 50}}).RemoveDuplicates().Out().Val())
}

func TestStream_RemoveDuplicates_validation(t *testing.T) {

	assert.Equal(
		t,
		errors.EmptyStream(stream.OpCodeRemoveDuplicates, "A nil Stream can not be iterated"),
		koazee.Stream().RemoveDuplicates().Out().Err())

}

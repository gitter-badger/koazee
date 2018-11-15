package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeFilter identifier for operation filter
const OpCodeFilter = "filter"

type filter struct {
	fn interface{}
}

func (op *filter) name() string {
	return OpCodeFilter
}

func (op *filter) run(s *Stream) *Stream {
	if err := op.validate(s); err != nil {
		s.err = err
		return s
	}
	function := reflect.ValueOf(op.fn)
	itemsType := reflect.TypeOf(s.items).Elem()
	newItems := reflect.MakeSlice(reflect.SliceOf(itemsType), 0, 0)
	items := reflect.ValueOf(s.items)
	for index := 0; index < items.Len(); index++ {
		item := items.Index(index)
		argv := make([]reflect.Value, 1)
		argv[0] = item
		if function.Call(argv)[0].Bool() {
			newItems = reflect.Append(newItems, item)
		}
	}
	s.items = newItems.Interface()
	return s
}

func (op *filter) validate(s *Stream) *errors.Error {
	if s.items == nil {
		return errors.EmptyStream(op.name(), "A nil Stream can not be filtered")
	}
	itemsType := reflect.TypeOf(s.items)
	function := reflect.ValueOf(op.fn)
	if function.Type().Kind() != reflect.Func {
		return errors.InvalidArgument(op.name(), "The filter operation requires a function as argument")
	}
	if function.Type().NumIn() != 1 {
		return errors.InvalidArgument(op.name(), "The provided function must retrieve 1 argument")
	}
	if function.Type().NumOut() != 1 {
		return errors.InvalidArgument(op.name(), "The provided function must return 1 value")
	}
	fnOut := reflect.New(function.Type().Out(0)).Elem()
	fnIn := reflect.New(function.Type().In(0)).Elem()
	if fnIn.Type() != itemsType.Elem() {
		return errors.InvalidArgument(op.name(),
			"The type of the argument in the provided function must be %s",
			itemsType.Elem().String())
	}
	if fnOut.Kind() != reflect.Bool {
		return errors.InvalidArgument(op.name(), "The type of the Output in the provided function must be bool")
	}
	return nil
}

// Filter discard the elements in the Stream that don't match with the provided filter
func (s *Stream) Filter(fn interface{}) *Stream {
	s.operations = append(s.operations, &filter{fn})
	return s
}

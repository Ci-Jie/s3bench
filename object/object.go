package object

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Object ...
type Object struct {
	candidates []string
	objects    map[string][]byte
}

func verify(input map[string]int) (err error) {
	var total int
	for _, rate := range input {
		total += rate
	}
	if total != 100 {
		return errors.New("The total of rate is not 100%")
	}
	return nil
}

// New ...
func New(input map[string]int) (object *Object, err error) {
	if err := verify(input); err != nil {
		return nil, err
	}
	object = &Object{}
	object.objects = make(map[string][]byte)
	for size, rate := range input {
		for i := 0; i < rate; i++ {
			object.candidates = append(object.candidates, size)
			if _, ok := object.objects[size]; !ok {
				value, err := convert(size)
				if err != nil {
					return nil, err
				}
				object.objects[size] = make([]byte, value)
			}
		}
	}
	return object, nil
}

// Get ...
func (o *Object) Get() (output []byte) {
	num := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	return o.objects[o.candidates[num]]
}

func convert(input string) (output int, err error) {
	s := map[string]int{"MB": 1024 * 1024, "KB": 1024, "Bytes": 1}
	for unit, value := range s {
		if strings.Contains(input, unit) {
			v, err := strconv.Atoi(strings.Split(input, unit)[0])
			if err != nil {
				return 0, err
			}
			return v * value, nil
		}
	}
	return 0, errors.New("The input is invalid")
}

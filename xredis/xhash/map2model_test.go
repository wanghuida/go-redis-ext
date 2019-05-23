package xhash

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type Map2modelTestSuite struct {
	suite.Suite
}

func (s *Map2modelTestSuite) SetupTest() { }

// 测试 Bool 相关
func (s *Map2modelTestSuite) TestBool() {
	data := make(map[string]string)
	data["bool_val"] = "1"
	type model struct {
		BoolVal bool
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(result.BoolVal, true, "test bool value err")
}

// 测试 *Bool 相关
func (s *Map2modelTestSuite) TestBoolPtr() {
	data := make(map[string]string)
	data["bool_ptr_val"] = "1"
	type model struct {
		BoolPtrVal *bool
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(*result.BoolPtrVal, true, "test bool ptr value err")
}

// 测试 Int 相关
func (s *Map2modelTestSuite) TestInt() {
	data := make(map[string]string)
	data["int_val"] = "1"
	type model struct {
		IntVal int
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(result.IntVal, int(1), "test int value err")
}

// 测试 *Int 相关
func (s *Map2modelTestSuite) TestIntPtr() {
	data := make(map[string]string)
	data["int_ptr_val"] = "1"
	type model struct {
		IntPtrVal *int
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(*result.IntPtrVal, int(1), "test int ptr value err")
}

// 测试 Uint 相关
func (s *Map2modelTestSuite) TestUint() {
	data := make(map[string]string)
	data["uint_val"] = "1"
	type model struct {
		UintVal uint
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(result.UintVal, uint(1), "test uint value err")
}

// 测试 *Uint 相关
func (s *Map2modelTestSuite) TestUintPtr() {
	data := make(map[string]string)
	data["uint_ptr_val"] = "1"
	type model struct {
		UintPtrVal *uint
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(*result.UintPtrVal, uint(1), "test uint ptr value err")
}

// 测试 Float 相关
func (s *Map2modelTestSuite) TestFloat() {
	data := make(map[string]string)
	data["float_val"] = "3.1415"
	type model struct {
		FloatVal float64
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(result.FloatVal, float64(3.1415), "test float64 value err")
}

// 测试 *Float 相关
func (s *Map2modelTestSuite) TestFloatPtr() {
	data := make(map[string]string)
	data["float_ptr_val"] = "3.1415"
	type model struct {
		FloatPtrVal *float64
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(*result.FloatPtrVal, float64(3.1415), "test float64 ptr value err")
}

// 测试 String 相关
func (s *Map2modelTestSuite) TestString() {
	data := make(map[string]string)
	data["string_val"] = "william"
	type model struct {
		StringVal string
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(result.StringVal, "william", "test string value err")
}

// 测试 *String 相关
func (s *Map2modelTestSuite) TestStringPtr() {
	data := make(map[string]string)
	data["string_ptr_val"] = "william"
	type model struct {
		StringPtrVal *string
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(*result.StringPtrVal, "william", "test string ptr value err")
}

// 测试 Slice 相关
func (s *Map2modelTestSuite) TestSlice() {
	data := make(map[string]string)
	data["slice_val"] = `["william", "wade"]`
	type model struct {
		SliceVal []string
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(len(result.SliceVal), 2, "test slice len value err")
	s.Equal(result.SliceVal[1], "wade", "test slice value err")
}

// 测试 Slice 相关
func (s *Map2modelTestSuite) TestTime() {
	data := make(map[string]string)
	data["time_val"] = "2019-05-23 10:20:30"
	type model struct {
		TimeVal time.Time
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(result.TimeVal.Year(), 2019, "test time year value err")
	s.Equal(result.TimeVal.Hour(), 10, "test time hour value err")
}

func TestMap2modelSuite(t *testing.T) {
	suite.Run(t, new(Map2modelTestSuite))
}


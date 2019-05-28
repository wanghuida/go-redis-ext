package xhash

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type Map2modelTestSuite struct {
	suite.Suite
}

func (s *Map2modelTestSuite) SetupTest() {}

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

// 测试 Int
func (s *Map2modelTestSuite) TestInt() {
	data := make(map[string]string)
	data["number"] = "1"
	type model struct {
		Number         int
		NumberPtrInt   *int   `redis:"number"`
		NumberPtrInt8  *int8  `redis:"number"`
		NumberPtrInt16 *int16 `redis:"number"`
		NumberPtrInt32 *int32 `redis:"number"`
		NumberPtrInt64 *int64 `redis:"number"`
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(result.Number, int(1), "test int value err")
	s.Equal(*result.NumberPtrInt, int(1), "test *int value err")
	s.Equal(*result.NumberPtrInt8, int8(1), "test *int8 value err")
	s.Equal(*result.NumberPtrInt16, int16(1), "test *int16 value err")
	s.Equal(*result.NumberPtrInt32, int32(1), "test *int32 value err")
	s.Equal(*result.NumberPtrInt64, int64(1), "test *int64 value err")
}

// 测试 Uint
func (s *Map2modelTestSuite) TestUint() {
	data := make(map[string]string)
	data["number"] = "1"
	type model struct {
		Number          uint
		NumberPtrUint   *uint   `redis:"number"`
		NumberPtrUint8  *uint8  `redis:"number"`
		NumberPtrUint16 *uint16 `redis:"number"`
		NumberPtrUint32 *uint32 `redis:"number"`
		NumberPtrUint64 *uint64 `redis:"number"`
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(result.Number, uint(1), "test uint value err")
	s.Equal(*result.NumberPtrUint, uint(1), "test *uint value err")
	s.Equal(*result.NumberPtrUint8, uint8(1), "test *uint8 value err")
	s.Equal(*result.NumberPtrUint16, uint16(1), "test *uint16 value err")
	s.Equal(*result.NumberPtrUint32, uint32(1), "test *uint32 value err")
	s.Equal(*result.NumberPtrUint64, uint64(1), "test *uint64 value err")
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

// 测试时间相关
func (s *Map2modelTestSuite) TestTime() {
	data := make(map[string]string)
	data["date_time"] = "2019-05-23 10:20:30"
	type model struct {
		CreatedAt time.Time  `redis:"date_time"`
		UpdatedAt *time.Time `redis:"date_time"`
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(result.CreatedAt.Year(), 2019, "test time year value err")
	s.Equal(result.UpdatedAt.Hour(), 10, "test time hour value err")
}

// 测试结构体相关
func (s *Map2modelTestSuite) TestModel() {

	data := make(map[string]string)
	data["user"] = `{"id": 100, "name": "Wade"}`
	type User struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}
	type model struct {
		User   *User
		Friend User `redis:"user"`
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(result.User.Id, int64(100), "test model value err")
	s.Equal(result.Friend.Name, "Wade", "test model value err")
}

// 测试接口相关
func (s *Map2modelTestSuite) TestInterface() {
	data := make(map[string]string)
	data["value"] = "1"
	type model struct {
		Value interface{}
	}
	result := new(model)
	err := Map2model(data, result)
	s.Nil(err)
	s.Equal(result.Value.(string), "1", "test interface value err")
}

// 测试不支持的类型
func (s *Map2modelTestSuite) TestNotSupportType() {
	data := make(map[string]string)
	data["value"] = "1"
	type model struct {
		Value complex128
	}
	result := new(model)
	err := Map2model(data, result)
	s.NotEmpty(err)
	s.Contains(err.Error(), "unsupported type", "test value err")
}

func TestMap2modelSuite(t *testing.T) {
	suite.Run(t, new(Map2modelTestSuite))
}

package util

import (
	"crypto/md5"
	"fmt"
	"reflect"
	"time"
	"unicode"

	"erpweb/logs"
)

//时间戳格式化
func Str2TimeStamp(param string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	the_time, err := time.ParseInLocation(timeLayout, param, time.Local)
	if err != nil {
		logs.FileLogs.Error("Str2TimeStamp failed : %s", err)
	}
	unix_time2 := the_time.Unix()

	return unix_time2
}

//时间戳格式化
func Str2Time(param string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	the_time, err := time.ParseInLocation(timeLayout, param, time.Local)
	if err != nil {
		logs.FileLogs.Error("Str2TimeStamp failed : %s", err)
	}

	return the_time
}

//uuid
func GetToken() string {
	// 创建
	u1, _ := uuid.NewV4()
	logs.FileLogs.Info("UUIDv4: %s\n", u1)

	return u1.String()
}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

// func SetField(obj interface{}, name string, value interface{}) error {
//     structValue := reflect.ValueOf(obj).Elem()
//     structFieldValue := structValue.FieldByName(name)

//     if !structFieldValue.IsValid() {
//         return fmt.Errorf("No such field: %s in obj", name)
//     }

//     if !structFieldValue.CanSet() {
//         return fmt.Errorf("Cannot set %s field value", name)
//     }

//     structFieldType := structFieldValue.Type()
//     val := reflect.ValueOf(value)
//     if structFieldType != val.Type() {
//         return errors.New("Provided value type didn't match obj field type")
//     }

//     structFieldValue.Set(val)
//     return nil
// }

// func (s *MyStruct) FillStruct(m map[string]interface{}) error {
//     for k, v := range m {
//         err := SetField(s, k, v)
//         if err != nil {
//             return err
//         }
//     }
//     return nil
// }

func GETMd5(str string) (md5str string) {
	data := []byte(str)
	has := md5.Sum(data)
	md5str = fmt.Sprintf("%x", has)
	return
}

func PanddingPwd(pwd string) bool {
	var (
		haveChar = 0
		haveNum  = 0
	)
	for _, value := range pwd {
		//valueStr := string(value)
		if unicode.IsLetter(value) {
			haveChar++
		} else if unicode.IsDigit(value) {
			haveNum++
		}
	}

	if haveChar > 0 && haveNum > 0 {
		return true
	}

	return false
}

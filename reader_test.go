package csvreader

import (
	"encoding/json"
	"fmt"
	"testing"
)

type testStruct struct {
	Name     string
	UserName string
	ID       int
	Enable   bool
	Type     CustomeType
}

type Info struct {
	Hostname string
	IP       string
}

type CustomeType int

func (c *CustomeType) FromString(str string) error {
	switch str {
	case "tcp":
		*c = 0
	case "udp":
		*c = 1
	default:
		return fmt.Errorf("unknown type:%s", str)
	}
	return nil
}

func TestBase(t *testing.T) {
	var infos []Info
	_ = New().UnMarshalFile("ip.csv", &infos)
	body, _ := json.Marshal(infos)
	fmt.Println(string(body))
}

func TestSnakeName(t *testing.T) {
	bean := []testStruct{}
	if err := New().
		WithHeader([]string{"name", "user_name", "id", "enable"}).
		UnMarshalBytes([]byte("zhengxin,zhnxin,0,false\nxinzheng,zhnxin,1,true"),
			&bean); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(bean)
	t.Log(string(b))
}

func TestLowerName(t *testing.T) {
	bean := []*testStruct{}
	if err := New().
		WithHeader([]string{"NAME", "USERNAME", "ID", "ENABLE"}).
		UnMarshalBytes([]byte("zhengxin,zhnxin,0,false\nxinzheng,zhnxin,1,true"),
			&bean); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(bean)
	t.Log(string(b))
}

func TestCustom(t *testing.T) {
	bean := []*testStruct{}
	if err := New().
		WithHeader([]string{"NAME", "USERNAME", "type", "ENABLE"}).
		UnMarshalBytes([]byte("zhengxin,zhnxin,udp,false\nxinzheng,zhnxin,tcp,true"),
			&bean); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(bean)
	t.Log(string(b))
}

func TestWrite(t *testing.T) {
	data := [][]string{
		{"Name", "Age", "City"}, // CSV 的表头
		{"Alice", "30", "New York"},
		{"Bob", "25", "San Francisco"},
		{"Charlie", "35", "Los Angeles"},
	}
	err := WirteAndSave(data, "./test.csv")
	if err != nil {
		fmt.Println(err)
	}
}

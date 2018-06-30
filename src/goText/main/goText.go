package main

import (
	"goText/fxml"
	"encoding/xml"
	"fmt"
	"os"
	"io/ioutil"
	"goText/fjson"
	"encoding/json"
)

// 编码生成XML文本并输出
func encodeXML() {
	v := &fxml.Servers{Version: "1"}
	v.Svs = append(v.Svs, fxml.Server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, fxml.Server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, " ", "	")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)

	// 写入文件, 有问题:无法正常格式化version
	//err = ioutil.WriteFile("fxml/ser.xml", []byte(xml.Header), 0644)
	//err = ioutil.WriteFile("fxml/ser.xml", output, 0644)
	//checkErr(err)

	// clear old content
	ioutil.WriteFile("fxml/ser.xml", []byte(""), 0644)

	fi, err := os.OpenFile("fxml/ser.xml", os.O_RDWR|os.O_APPEND,0755)
	fi.Write([]byte(xml.Header))
	fi.Write(output)
	checkErr(err)
	defer fi.Close()
}

//解析XML
func decodeXML() {
	file, err := os.Open("fxml/ser.xml")	// for read access
	if err != nil {
		fmt.Printf("errpr: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := fxml.Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt. Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}

//生成json
func encodeJSON() {
	var s fjson.Serverslice
	s.Servers = append(s.Servers, fjson.Server{ServerName: "Shanghai VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, fjson.Server{ServerName:"Beijing VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

//解析json
func decodeJSON() {

	// 已知结构体结构解析

	var s fjson.Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println("str = ", s)

	// 未知结构体结构解析

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("解析未知结构json失败!")
		return
	}
	fmt.Println("f = ", f)

	//通过断言方式访问未知结构体数据,map[string]interface{}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println("  ", i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle!")
		}
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//encodeXML()
	//decodeXML()
	encodeJSON()
	//decodeJSON()
}



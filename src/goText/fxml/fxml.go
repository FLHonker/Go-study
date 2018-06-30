package fxml

import (
	"encoding/xml"
)

type Servers struct {
	XMLName 	xml.Name	`xml:"servers"`
	Version		string		`xml:"version,attr"`
	Svs 		[]Server 	`xml:"server"`
}

type Server struct {
	ServerName	string	`xml:"serverName"`
	ServerIP	string 	`xml:"ServerIP"`
}

type Recurlyservers struct {
	XMLName		xml.Name	`xml:"servers"`
	Version		string 		`xml:"version,attr"`
	Svs			[]Server	`xml:"server"`
	Description string		`xml:",innerxml"`
}


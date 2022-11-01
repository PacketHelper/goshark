package goshark

import "encoding/xml"

type Packet struct {
	XMLName xml.Name `xml:"packet"`
	Protos  []Proto  `xml:"proto"`
}
type Proto struct {
	XMLNAME  xml.Name `xml:"proto"`
	Name     string   `xml:"name,attr"`
	Pos      int      `xml:"pos,attr"`
	Showname string   `xml:"showname,attr"`
	Size     int      `xml:"size,attr"`
	Field    []Field  `xml:"field"`
}

type Field struct {
	Name          string          `xml:"name,attr"`
	Pos           int             `xml:"pos,attr"`
	Show          string          `xml:"show,attr"`
	Showname      string          `xml:"showname,attr"`
	Value         string          `xml:"value,attr"`
	Size          int             `xml:"size,attr"`
	DetailedField []DetailedField `xml:"field"`
}

type DetailedField struct {
	Name     string `xml:"name,attr"`
	Pos      int    `xml:"pos,attr"`
	Show     string `xml:"show,attr"`
	Showname string `xml:"showname,attr"`
	Value    string `xml:"value,attr"`
	Size     int    `xml:"size,attr"`
	Hide     string `xml:"hide,attr"`
}

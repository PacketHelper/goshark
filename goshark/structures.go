package goshark

import "encoding/xml"

type Packet struct {
	XMLName xml.Name `json:"-" xml:"packet"`
	Protos  []Proto  `json:"protos" xml:"proto"`
}
type Proto struct {
	XMLNAME  xml.Name `json:"-" xml:"proto"`
	Name     string   `json:"name" xml:"name,attr"`
	Pos      int      `json:"pos" xml:"pos,attr"`
	Showname string   `json:"showname" xml:"showname,attr"`
	Size     int      `json:"size" xml:"size,attr"`
	Field    []Field  `json:"fields" xml:"field"`
}

type Field struct {
	Name          string          `json:"name" xml:"name,attr"`
	Pos           int             `json:"pos" xml:"pos,attr"`
	Show          string          `json:"show" xml:"show,attr"`
	Showname      string          `json:"showname" xml:"showname,attr"`
	Value         string          `json:"value" xml:"value,attr"`
	Size          int             `json:"size" xml:"size,attr"`
	DetailedField []DetailedField `json:"detailed_fields" xml:"field"`
}

type DetailedField struct {
	Name     string `json:"name" xml:"name,attr"`
	Pos      int    `json:"pos" xml:"pos,attr"`
	Show     string `json:"show" xml:"show,attr"`
	Showname string `json:"showname" xml:"showname,attr"`
	Value    string `json:"value" xml:"value,attr"`
	Size     int    `json:"size" xml:"size,attr"`
	Hide     string `json:"hide" xml:"hide,attr"`
}

package main

import (
	"encoding/xml"
)

type Request struct {
	XMLName     xml.Name    `xml:"REQUEST"`
	RequestInfo RequestInfo `xml:"REQUEST_INFO"`
	Records     Records     `xml:"RECORDS"`
}

type RequestInfo struct {
	XMLName   xml.Name `xml:"REQUEST_INFO"`
	RequestId string   `xml:"REQUESTID"`
	Year      string   `xml:"YEAR"`
	Month     string   `xml:"MONTH"`
	Upid      string   `xml:"UPID"`
	Upname    string   `xml:"UPNAME"`
	Filetype  string   `xml:"FILETYPE"`
	Comment   string   `xml:"COMMENT"`
}

type Records struct {
	XMLName xml.Name `xml:"RECORDS"`
	Records []Record `xml:"RECORD"`
}

type Record struct {
	XMLName  xml.Name `xml:"RECORD"`
	Recordid string   `xml:"RECORDID"`
	Publicid string   `xml:"PUBLICID"`
	Lname    string   `xml:"LNAME"`
	Fname    string   `xml:"FNAME"`
	Mname    string   `xml:"MNAME"`
	Snils    string   `xml:"SNILS"`
	Bdate    string   `xml:"BDATE"`
	Sex      string   `xml:"SEX"`
	Kladr    string   `xml:"KLADR"`
	District string   `xml:"DISTRICT"`
	City     string   `xml:"CITY"`
	Punkt    string   `xml:"PUNKT"`
	Street   string   `xml:"STREET"`
	Bnum     string   `xml:"BNUM"`
	Wing     string   `xml:"WING"`
	Appnum   string   `xml:"APPNUM"`
	Wgcode   string   `xml:"WGCODE"`
	Status   string   `xml:"STATUS"`
	Reason   string   `xml:"REASON"`
	Scomment string   `xml:"SCOMMENT"`
	Ltype    string   `xml:"LTYPE"`
	Bstate   string   `xml:"BSTATE"`
	Owtype   string   `xml:"OWTYPE"`
	Sqtotal  string   `xml:"SQ_TOTAL"`
	Sqliving string   `xml:"SQ_LIVING"`
	Errcode  string   `xml:"ERRCODE"`
	Rcomment string   `xml:"RCOMMENT"`
	L_schets L_schets `xml:"L_SCHETS"`
	Uoid     string   `xml:"UOID"`
	Uoname   string   `xml:"UONAME"`
}

type L_schets struct {
	XMLName  xml.Name  `xml:"L_SCHETS"`
	L_schets []L_schet `xml:"L_SCHET"`
}

type L_schet struct {
	XMLName     xml.Name    `xml:"L_SCHET"`
	Lschet      string      `xml:"LSCHET"`
	Owlname     string      `xml:"OW_LNAME"`
	Owfname     string      `xml:"OW_FNAME"`
	Owmname     string      `xml:"OW_MNAME"`
	Owsnils     string      `xml:"OW_SNILS"`
	Owbdate     string      `xml:"OW_BDATE"`
	Owsex       string      `xml:"OW_SEX"`
	Lsmembers   Lsmembers   `xml:"LS_MEMBERS"`
	Lsutilities Lsutilities `xml:"LS_UTILITIES"`
}

type Lsmembers struct {
	XMLName   xml.Name   `xml:"LS_MEMBERS"`
	Lsmembers []Lsmember `xml:"LS_MEMBER"`
}

type Lsmember struct {
	XMLName xml.Name `xml:"LS_MEMBER"`
	Lslname string   `xml:"LS_LNAME"`
	Lsfname string   `xml:"LS_FNAME"`
	Lsmname string   `xml:"LS_MNAME"`
	Lssnils string   `xml:"LS_SNILS"`
	Lsbdate string   `xml:"LS_BDATE"`
	Lssex   string   `xml:"LS_SEX"`
	Lsrel   string   `xml:"LS_REL"`
	Indate  string   `xml:"IN_DATE"`
	Outdate string   `xml:"OUT_DATE"`
	Regtype string   `xml:"REG_TYPE"`
}

type Lsutilities struct {
	XMLName     xml.Name    `xml:"LS_UTILITIES"`
	Lsutilities []Lsutility `xml:"LS_UTILITY"`
}

type Lsutility struct {
	XMLName   xml.Name `xml:"LS_UTILITY"`
	Utility   string   `xml:"UTILITY"`
	Extlschet string   `xml:"EXT_LSCHET"`
	Extupid   string   `xml:"EXT_UPID"`
	Sum1      string   `xml:"SUM1"`
	Sum2      string   `xml:"SUM2"`
	Debt      string   `xml:"DEBT"`
}

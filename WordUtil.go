package main

import (
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
	"reflect"
)


type WordUtil struct {

}

func (wordUtil *WordUtil) WriteTableInfo(columns []WordStruct, data map[string] []DBStruct, saveFileName string)  {
	doc := document.New()


	for key,value := range data {

		doc.AddParagraph()
		doc.AddParagraph().AddRun().AddText("表名：" + key)
		table := doc.AddTable()
		// width of the page
		table.Properties().SetWidthPercent(100)
		// with thick borers
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)
		row := table.AddRow()

		for i:=0; i<len(columns); i++ {
			p := row.AddCell().AddParagraph()

			p.AddRun().AddText(columns[i].Title)
		}
		for j:=0; j<len(value); j++ {
			row := table.AddRow()
			for i:=0; i<len(columns); i++ {
				field := columns[i].Field
				columType := reflect.ValueOf(value[j])
				value := columType.FieldByName(field).String()
				row.AddCell().AddParagraph().AddRun().AddText(value)
			}
		}

		doc.AddParagraph()
	}












	//row.AddCell().AddParagraph().AddRun().AddText("John Smith")
	//row = table.AddRow()
	//row.AddCell().AddParagraph().AddRun().AddText("Street Address")
	//row.AddCell().AddParagraph().AddRun().AddText("111 Country Road")

	doc.SaveToFile(saveFileName)
}

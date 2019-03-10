package main

import(
	"fmt"
	"github.com/tealeg/xlsx"
)

var (
	inFile = "/Users/sceaj/Desktop/test.xlsx" //檔案路徑
)

func main(){
	// 打開文件
    xlFile, err := xlsx.OpenFile(inFile)
    if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 每sheet頁讀取
    for _, sheet := range xlFile.Sheets {
		fmt.Println("sheet name: ", sheet.Name)
		//每行讀取
        for _, row := range sheet.Rows {
			// 每行的列讀取
            for _, cell := range row.Cells {
				text := cell.String()
                fmt.Printf("%12s", text)
			}
			fmt.Print("\n")
        }
	}
}
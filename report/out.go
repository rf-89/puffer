package report

import (
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Out(f string, infos []map[string]string) error {
	fh := excelize.NewFile()
	// シートを新規作成
	index := fh.NewSheet("Sheet1")
	// シートに値を書き込む
	fh.SetCellValue("Sheet1", "A1", "Report")
	// 見出しセルの塗りつぶし設定
	cellStyle := `{
        "fill": {
            "type": "pattern",
            "color": [
                "#D3D3D3"
            ],
            "pattern": 1
        }
    }`
	styleID, err := fh.NewStyle(cellStyle)
	if err != nil {
		log.Println(err)
		log.Fatalln("failed to cell style setting.")
	}

	//先頭行の見出し設定
	fh.SetCellStyle("Sheet1", "A2", "E2", styleID)

	//ヘッダー行の出力
	fh.SetCellValue("Sheet1", "A2", "no")
	fh.SetCellValue("Sheet1", "B2", "refer file")
	fh.SetCellValue("Sheet1", "C2", "refer hash")
	fh.SetCellValue("Sheet1", "D2", "duplicate file")
	fh.SetCellValue("Sheet1", "E2", "duplicate hash")

	for c, info := range infos {
		c = c + 3
		axisIndex := strconv.Itoa(c)
		fh.SetCellValue("Sheet1", "A"+axisIndex, strconv.Itoa(c-2))
		fh.SetCellValue("Sheet1", "B"+axisIndex, info["filename"])
		fh.SetCellValue("Sheet1", "C"+axisIndex, info["hash"])
		fh.SetCellValue("Sheet1", "D"+axisIndex, info["duplicate_filename"])
		fh.SetCellValue("Sheet1", "E"+axisIndex, info["duplicate_hash"])
	}
	// シート1をアクティブにする。
	fh.SetActiveSheet(index)

	if err := fh.SaveAs(f); err != nil {
		log.Fatalln("failed to output report.")
	}

	return nil
}

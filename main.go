package main

import (
	"fmt"
	"os"
)

func main() {
	var bodyPart string
	var exercise string
	var weight float64
	var reps int
	var sets int

	fmt.Println("筋トレ記録アプリ")
	data, err := os.ReadFile("records.csv")
	if err == nil {
		fmt.Println("\n--- 過去の記録 ---")
		fmt.Print(string(data))
		fmt.Println("------------------")
	} else if !os.IsNotExist(err) {
		fmt.Println("過去の記録を読み込めませんでした")
	}
	fmt.Print("鍛えたい部位を入力してください（胸・背中・脚）：")
	fmt.Scan(&bodyPart)
	switch bodyPart {
	case "胸":
		fmt.Println("おすすめ：ベンチプレス、ダンベルフライ、腕立て伏せ")
	case "背中":
		fmt.Println("おすすめ：ラットプルダウン、シーテッドロー、懸垂")
	case "脚":
		fmt.Println("おすすめ：スクワット、レッグプレス、レッグカール")
	default:
		fmt.Println("胸・背中・脚のいずれかを入力してください")
	}

	fmt.Print("種目を入力してください：")
	fmt.Scan(&exercise)

	fmt.Print("重量を入力してください：")
	fmt.Scan(&weight)

	fmt.Print("回数を入力してください：")
	fmt.Scan(&reps)

	fmt.Print("セット数を入力してください：")
	fmt.Scan(&sets)
	file, err := os.OpenFile("records.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("記録の保存に失敗しました")
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "%s,%.1f,%d,%d\n", exercise, weight, reps, sets)
	fmt.Printf("%sを%.1fkgで%d回×%dセット記録しました！\n", exercise, weight, reps, sets)
}

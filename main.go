package main

import "fmt"

func main() {
	var exercise string
	var weight float64
	var reps int
	var sets int

	fmt.Println("筋トレ記録アプリ")

	fmt.Print("種目を入力してください：")
	fmt.Scan(&exercise)

	fmt.Print("重量を入力してください：")
	fmt.Scan(&weight)

	fmt.Print("回数を入力してください：")
	fmt.Scan(&reps)

	fmt.Print("セット数を入力してください：")
	fmt.Scan(&sets)

	fmt.Printf("%sを%.1fkgで%d回×%dセット記録しました！\n", exercise, weight, reps, sets)
}

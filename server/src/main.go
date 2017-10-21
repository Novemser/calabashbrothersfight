package main

import "fmt"

var gameState = new(GameState)

//TODO
func stepThread(thread int) {
	if IsLevelPristine() {
		//第一步执行
	}
	var program = gameState.GetProgramOfThread(thread);
	fmt.Println(program)

}



func IsLevelPristine() bool {
	//return undoHistory.length == 0;
	return true
}

func main() {

	fmt.Println(gameState)
	fmt.Printf("Go")
}

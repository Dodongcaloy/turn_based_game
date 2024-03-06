package main

import (
	"bufio"
	"fmt"
	"os"
)

type Player struct {
	Name    string
	Life    int
	Attack  int
	Defense int
	Heal    int
}

func main() {

	player1 := &Player{
		Life:    100,
		Attack:  70,
		Defense: 40,
		Heal:    20,
	}
	player2 := &Player{
		Life:    100,
		Attack:  75,
		Defense: 30,
		Heal:    25,
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter player1 name: ")
	text1, _ := reader.ReadString('\n')
	fmt.Println("Player1 is now set to " + text1)
	fmt.Printf("Name: %s Life: %d, Defense:%d \n", text1, player1.Life, player1.Defense)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Print("Enter player2 name: ")
	text2, _ := reader.ReadString('\n')
	fmt.Println("Player2 is now set to " + text2)
	fmt.Printf("Name: %s Life: %d, Defense:%d \n", text2, player2.Life, player2.Defense)
	fmt.Println("")
	fmt.Println("Game Start")
	fmt.Println("")
	fmt.Println("--------Player ", text1, "turn --------")

	for {

		fmt.Println("")
		fmt.Print("1 >>> attack (70) \n 2 >>> heal (20):\n ")
		fmt.Println("")
		var y int
		fmt.Scan(&y)
		if y == 1 {
			attack(player1, player2)
			fmt.Println("")
			fmt.Printf(text1 + "used attack\n")
			fmt.Printf(text2+"life went down to %d HP\n %d HP %d DEF\n", player2.Life, player2.Life, player2.Defense)
			fmt.Println("")
		} else if y == 2 {
			heal(player1)
			fmt.Println("")
			fmt.Printf(text1 + "used heal\n")
			fmt.Printf(text1+"hp went up to %d HP\n %d HP %d DEF\n", player1.Life, player1.Life, player1.Defense)
			fmt.Println("")
		}

		if player2.Life <= 0 {
			isdead()
			break
		} else {
			fmt.Println("--------Player ", text2, "turn --------")
			fmt.Println("")
			fmt.Printf("%s %d HP,%d DEF\n", text1, player1.Life, player1.Defense)
			fmt.Println("")
			fmt.Print("1 >>> attack (60) \n 2 >>> heal (25): \n")
			fmt.Println("")
		}
		var x int
		fmt.Scan(&x)
		if x == 1 {
			attack(player2, player1)
			fmt.Println("")
			fmt.Printf(text2 + "used attack\n")
			fmt.Printf(text1+"life went down to %d HP\n %d HP %d DEF\n", player1.Life, player1.Life, player1.Defense)
			fmt.Println("")
		} else if x == 2 {
			heal(player2)
			fmt.Println("")
			fmt.Printf(text2 + "used heal\n")
			fmt.Printf(text2+"hp went up to %d HP\n %d HP %d DEF", player2.Life, player1.Life, player2.Defense)
			fmt.Println("")
		}

		if player1.Life <= 0 {
			fmt.Println("")
			isdead()
			break
		} else {
			fmt.Println("--------Player ", text1, "turn --------")
			fmt.Printf("%s %d HP,%d DEF\n", text2, player2.Life, player2.Defense)
		}
	}

}

func isdead() {
	fmt.Println("ENEMY DEFEATED")
	fmt.Println()
}
func attack(attacker *Player, target *Player) {
	target.Life = target.Life + target.Defense - attacker.Attack
	return
}
func heal(restore *Player) {
	restore.Life = restore.Life + restore.Heal
	return
}

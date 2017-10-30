package main

import "time"
import "fmt"

func main() {
    var Ball int
    table := make(chan int)
    go player("player1", table)
	go player("player2", table)
	go player("player3", table)

    table <- Ball
    time.Sleep(1 * time.Second)
    <-table
}

func player(player string, table chan int) {
    for {
        ball := <-table
        fmt.Println(player, ball)
        ball++
        time.Sleep(100 * time.Millisecond)
        table <- ball
    }
}
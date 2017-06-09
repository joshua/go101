package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Ball ...
type Ball struct {
	Speed int
}

// Player ...
type Player struct {
	Name     string
	Score    int
	Strength int
	Tired    *rand.Rand
}

// ScoreKeeper ...
type ScoreKeeper struct{}

// Start : Allows the ScoreKeeper to start Tracking the players scores
// KEEPER_START OMIT
func (s *ScoreKeeper) Track(p1, p2 *Player, p1Scored, p2Scored chan string) {
	for {
		var playerName string

		select {
		case playerName = <-p1Scored:
			fmt.Printf("%s scored\n", playerName)
		case playerName = <-p2Scored:
			fmt.Printf("%s scored\n", playerName)
		}

		if p1.Name == playerName {
			p1.Score++
		}

		if p2.Name == playerName {
			p2.Score++
		}
	}
}

// KEEPER_END OMIT

// Turn : A player takes turn hitting the ball,
// 		  A player wins a turn if the ball return has zero speed
// TURN_START OMIT
func (p *Player) Turn(b chan *Ball, scored chan string) {
	for {
		ball := <-b

		if ball.Speed == 0 {
			// Send message to score keeper
			scored <- p.Name

			// Reset ball.Speed to serve back to other player
			ball.Speed = p.Strength
		}

		calculateHit(p, ball)
		b <- ball
	}
}

// TURN_END OMIT

// IsTired : returns true if player is tired
func (p *Player) IsTired() bool {

	prob := p.Tired.Intn(50)
	pivot := 20

	if prob <= pivot { // Less likelihood of being tired
		return true
	}
	return false
}

// calculates hit based on player.Strength and ball.Speed
// ball.Speed = 0 indicates that player did not hit
func calculateHit(p *Player, b *Ball) {
	switch {
	case p.Strength > b.Speed:
		b.Speed++
		if p.IsTired() {
			p.Strength--
		}
		p.Strength++
	case p.Strength < b.Speed:
		if p.IsTired() {
			p.Strength--
			b.Speed = 0
			return
		}

	case p.Strength == b.Speed:
		if p.IsTired() {
			b.Speed = 0
			return
		}
	}
}

// MAIN_START OMIT
func main() {
	ball := make(chan *Ball)
	seed := rand.NewSource(time.Now().UnixNano())

	p1 := &Player{Name: "player 1", Strength: 6, Tired: rand.New(seed)}
	p2 := &Player{Name: "player 2", Strength: 6, Tired: rand.New(seed)}
	scoreKeeper := &ScoreKeeper{}

	p1Scored := make(chan string)
	p2Scored := make(chan string)

	go p1.Turn(ball, p1Scored)
	go p2.Turn(ball, p2Scored)
	go scoreKeeper.Track(p1, p2, p1Scored, p2Scored)

	ball <- &Ball{Speed: 5}
	time.Sleep(500 * time.Microsecond)
	<-ball

	fmt.Print("\nScore:\n=============================\n")
	fmt.Printf("%s: %d, %s: %d\n", p1.Name, p1.Score, p2.Name, p2.Score)
}

// MAIN_END OMIT

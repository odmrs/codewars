package kata

func Rps(p1, p2 string) string {
	var scissors, paper, rock = "scissors", "paper", "rock"
	if p1 == paper && p2 == paper || p1 == scissors && p2 == scissors || p1 == rock && p2 == rock {
		return "Draw!"
	}
	if (p1 == scissors && p2 == paper) || (p1 == rock && p2 == scissors) || (p1 == paper && p2 == rock) {
		return "Player 1 won!"
	} else {
		return "Player 2 won!"
	}
}

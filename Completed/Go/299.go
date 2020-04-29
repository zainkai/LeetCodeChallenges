import (
    "fmt"
)

func getHint(secret string, guess string) string {
    bulls, cows := 0 ,0
    
    m := map[byte]int{}
    for i := 0; i < len(guess); i++ {
        if secret[i] == guess[i] {
            bulls++
        } else {
           m[guess[i]]++ 
        } 
    }
    
    for i := 0; i < len(guess); i++  {
        if secret[i] != guess[i] && m[secret[i]] > 0 {
            m[secret[i]]--
            cows++
        }
    }
    return fmt.Sprintf("%dA%dB", bulls, cows)
}
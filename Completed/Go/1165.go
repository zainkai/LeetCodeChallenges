func calculateTime(keyboard string, word string) int {
    keyboardCost := map[rune]int{}
    for i, key := range keyboard {
        keyboardCost[key] = i
    }
    
    res := 0
    lastKey := 0
    for _, key := range word {
        newKey := keyboardCost[key]
        if newKey < lastKey {
            res += lastKey - newKey
        } else {
            res += newKey - lastKey
        }
        lastKey = newKey
    }
    return res
}

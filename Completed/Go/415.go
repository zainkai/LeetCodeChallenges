func addStrings(num1 string, num2 string) string {
    carry := 0
    idx1, idx2 := len(num1)-1, len(num2)-1
    result := ""
    
    for idx1 >= 0 || idx2 >= 0 {
        n1 := 0
        if idx1 >= 0 {
            n1 = int(num1[idx1] - '0')
            idx1--
        }
        n2 := 0
        if idx2 >= 0 {
            n2 = int(num2[idx2] - '0')
            idx2--
        }
        
        val := (n1 + n2 + carry) % 10
        carry = (n1 + n2 + carry) / 10
        result = string('0'+ val) + result
    }
    
    if carry > 0 {
        return string('0'+carry) + result
    }
    return result
}

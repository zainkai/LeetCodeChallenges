
func countPrimes(n int) int {
    isNotPrime := make([]bool, n)
    cnt := 0
    
    for i := 2; i < n; i++ {
        if isNotPrime[i] {
            continue
        }
        
        cnt++
        for j := i*i; j < n; j += i{
            isNotPrime[j] = true
        }
    }
    
    return cnt
}

// import "math"

// func countPrimes(n int) int {
//     primes := 0
//     for i := 1; i < n; i++ {
//         if isPrime(i) {
//             primes++
//         }
//     }
//     return primes
// }

// func isPrime(n int) bool {
//     if n <= 1 {
//         return false
//     }
    
//     sqr := math.Sqrt(float64(n))
    
//     for i := 2; i <= int(sqr); i++ {
//         if n % i == 0 {
//             return false
//         }
//     }
//     return true
// }
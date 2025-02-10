func myFunc(a []int) {
    b := []int{}
    for _, v := range a {
        if v != 0 {
            b = append(b, v)
        }
    }
    a = b // Assign to the original slice if needed
}     
//Alternative Solution (more efficient, in-place):
func myFuncEfficient(a []int) []int {
    j := 0
    for i := 0; i < len(a); i++ {
        if a[i] != 0 {
            a[j] = a[i]
            j++
        }
    }
    return a[:j]
} 
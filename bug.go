func myFunc(a []int) {
    for i := range a {
        if a[i] == 0 {
            a = append(a[:i], a[i+1:]...) // Removing element from slice
        }
    }
}
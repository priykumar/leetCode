func minSwaps(s string) int {
    // Initialize variables to keep track of closing brackets and maximum imbalance
    var close, maxClose int

    // Iterate through each character in the string
    for _, c := range s {
        if c == ']' {
            // If we encounter a closing bracket, increment the 'close' counter
            close++
            // Update maxClose to be the maximum of current maxClose and close
            // This keeps track of the maximum imbalance we've seen so far
            maxClose = max(maxClose, close)
        } else {
            // If we encounter an opening bracket '[', decrement the 'close' counter
            // This balances out a previous closing bracket
            close--
        }
    }

    // The minimum number of swaps needed is (maxClose + 1) / 2
    // This formula works because:
    // 1. We need to balance out the maximum imbalance we've seen (maxClose)
    // 2. Each swap can fix two imbalances (it moves a '[' to the left and a ']' to the right)
    // 3. We round up using (x + 1) / 2 to handle odd numbers of imbalances
    return (maxClose + 1) / 2
}
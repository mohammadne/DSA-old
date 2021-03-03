# Maximum Subsequence Sum Problem

Given (possibly negative) integers A1, A2,. . ., AN, ﬁnd the maximum subsequence sum (it is 0 if all the integers are negative).

Example:For input−2,11,−4,13,−5, −2,the answer is 20 (A2through A4).

## SOLUTIONS RUNNING TIME:
### solution_1: O(N^3)

### solution_2: O(N^2)

### solution_3 (divide-and-conquer): O(NlogN)
the maximum subsequence sum can be in one of three places. Either itoccurs entirely in the left half of the input, or entirely in the right half, or it crosses themiddle and is in both halves

### solution_4 (Kadane’s algorithm): O(N)
One observation is that if a[i] is negative, then it cannot possibly be the start of the optimal subsequence,since any subsequence that begins by including a[i] would be improved by beginning with a[i+1]. 
Similarly, any negative subsequence cannot possibly be a preﬁx of the optimal subsequence (same logic).
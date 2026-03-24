***Your input:*** A list of step entries, one per line. Each entry contains a Unix timestamp, a clue label, and a cardinal direction.

***Input format:*** timestamp: location (direction)
  NOTE: direction is one of: N, S, E, W.

***Your task:*** Track movement across all steps. N and E each add 1, S and W each subtract 1. Sum both the N/S and E/W totals together for your final answer.

***Example Input:***

```
1742134800: spooky trees (N)
1741789200: footprints (E)
1741270800: quicksand (W)
1740924000: marked trees (S)
1742048400: dark ravine (N)
```

- ***Output Format:*** `total steps`
- ***Example Output:*** `1` (2 north, 1 east, 1 west, 1 south → N-S: 1, E-W: 0 → total: 1)

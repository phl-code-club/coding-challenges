***Your input:*** A list of step entries, one per line. Each entry contains a UNIX timestamp, a clue label, and a cardinal direction.

***Input format:*** `timestamp: location (direction)`
  **NOTE**: direction is one of: `N`, `S`, `E`, `W`.

***Your task:*** Track movement across all steps. Each step is considered
`10 kilometers` For each step count the N/S or E/W movement. For each N or E step,
add 100 meters, for each S or W step subtract 100 meters. Your final output is
the N/S distance multiplied by the E/W distance.

***Example Input:***

```
1742134800: bones (N)
1741789200: footprints (E)
1740924000: marked trees (S)
1741270800: quicksand (W)
1742048400: dark ravine (N)
1742048400: river rapids (N)
1742048400: jagged rocks (E)
```

- ***Output Format:*** `N/S distance * E/W distance`
- ***Example Output:*** `200` (3 North, 2 East, 1 West, 1 South →
N-S = 30 - 10 = 20, E-W = 20 - 10 = 10 →
total: 20 * 10 = 200)

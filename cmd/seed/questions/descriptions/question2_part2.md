As you make your way on your trek, something feels off... You've been here before. These directions have you going in circles. To not waste precious time, you swiftly pull out a pen and start re-calculating your steps, crossing out any paths that loop back to where you started.

***Your task:*** Same as Part 1, but any group of 4 consecutive directions that form a closed loop cancel out. Scan through the directions in order — when you find a loop, subtract 1 from your total and skip those 4 steps entirely before continuing.he final rune in the previous example can be fixed.

***Example Input:***

```
1742134800: spooky trees (N)
1741789200: footprints (E)
1741270800: quicksand (S)
1740924000: marked trees (W)
1742048400: dark ravine (N)
```

- ***valid loops are:*** NWSE, NESW, SWNE, SENW, WNES, WSEN, ESWN, ENWS
- ***Output Format:*** `part 1 total steps minus each loop occurrence`
- ***Example Output:*** `0` (N/S: 1, E/W: 0 → total: 1, one loop found → 1 - 1 = 0)

Note: Loops are non-overlapping — once a loop is found, skip past all 4 steps before looking for the next one

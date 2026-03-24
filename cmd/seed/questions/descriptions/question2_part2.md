As you make your way on your trek, something feels off... You've been here before.
These directions have you going in circles. To not waste precious time, you
swiftly pull out a pen and start re-calculating your steps, crossing out any
paths that loop back to where you started.

***Your task:*** Same as Part 1, but any group of 4 consecutive directions that
form a closed loop cancel out. Scan through the directions in order — when you
find a loop, subtract 1 from your total.

***This will use the same example input from Part 1***

- ***valid loops are:*** `NWSE`, `NESW`, `SWNE`, `SENW`, `WNES`, `WSEN`, `ESWN`, `ENWS`
- ***Output Format:*** `N/S distance * E/W distance - loop count`
- ***Example Output:*** `199` (Previous total: 200, one loop found → 200 - 1 = 199)

Note: Loops are non-overlapping — once a loop is found, skip past all 4 steps
before looking for the next one.

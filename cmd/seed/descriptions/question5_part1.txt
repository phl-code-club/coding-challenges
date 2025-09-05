***Your task:*** For each query, determine if a path exists between the start and end points. Output the number of queries where a path exists.

***Example:***

```
....
.##.
....
0 0 2 3
0 0 1 1
```

***Analysis:***

- Query 1 (0,0) to (2,3): Path exists (can go around the walls)
- Query 2 (0,0) to (1,1): No path (destination is a wall)
- ***Output:*** `1`
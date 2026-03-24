***Your Input:*** A grid of characters between 800×800 and 1300×1300 in size. Each cell is either @ (land) or ~ (water).

***Your task:*** Find all the islands in the grid then multiple by the grid size.

NOTE: An island is a connected group of @ cells — two @ cells belong to the same island if they are directly adjacent (up, down, left, or right, but not diagonal).

***Example Input:***

```
~~~~~~~~
~~@@~~~~
~~@@~~~~
~~~~@~~~
~~~~@@@@
~~~~~~~~
~~~@~~~~
~~~~~~~~
```

In the example above there are 3 islands — the 2×2 block, the L-shape on row 4-5, and the single cell on row 7.

***Analysis:***

- ***Output Format:*** `product of island count and grid size`
- ***Example Output:*** `24` (3 islands × grid size of 8 = 24)`


***Your task:*** For the same input, also count the locations with the type `trap` or `thief`. Then compute the answer with the following rules:

| Label | Value |
| -------------- | --------------- |
| clue | 40 |
| landmark | 30 |
| thief | 20 |
| trap | 10 |

***This will use the same example input from Part 1***

***Analysis:***

- `clue`: 1
- `landmark`: 1
- `trap`: 1
- `thief`: 1
- ***Output Format:*** `(40 * clue count) + (30 * landmark count) - (20 * thief count) - (10 * trap count)`
- ***Example Output:*** `40`

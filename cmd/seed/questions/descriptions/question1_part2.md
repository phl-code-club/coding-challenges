Your initial plan seems a little too good to be true. Upon closer inspection you
see some of the entries are obstacles and adversaries! You decide to account for
these in your planning to get a better idea of how long this will take.

***Your task:*** For the same input, also count the locations with the type `trap` or `thief`. Then compute the answer with the following rules:

| Label | Value |
| -------------- | --------------- |
| clue | 40 |
| landmark | 30 |
| thief | 20 |
| trap | 10 |

***This will use the same example input from Part 1***

***Analysis:***

- clue count: 1
- landmark count: 1
- trap count: 1
- thief count: 1
- ***Output Format:*** `(40 * clue count) + (30 * landmark count) - (20 * thief count) - (10 * trap count)`
- ***Example Output:*** `40`

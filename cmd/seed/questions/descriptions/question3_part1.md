***Your input:*** A list of sales where each sale takes a whole line and is in the format `timestamp:amount:category`. Where `timestamp` is an `int`, `amount` is also an `int`, and category is a `string`.

***Your task:*** Group sales by category and find the category with the highest total amount. Output the category name and its total amount.

***Example input:***

```
110:50:herbs
175:30:scrolls
220:25:herbs
250:40:scrolls
300:35:elixirs
```

***Analysis:***

- herbs: 50 + 25 = 75
- scrolls: 30 + 40 = 70
- elixirs: 35
- ***Output Format:*** `category, amount`
- ***Example Output:*** `herbs,75`
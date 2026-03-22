Each recipe consists of ingredients with numeric amounts. A recipe is valid if it meets all these criteria:

1. It has at least 3 ingredients
2. Total amount is greater than or equal to 100
3. No single ingredient exceeds 50% of the total

***Your task:*** Given a list of recipes, count how many are valid according to the rules above.

***Example recipes:***

```
mugwort:30, eye of newt:40, vinegar:30, foxglove: 15
cilantro:60, dragons claw:25, toad juice:15
deathclaw saliva:70, devils cap:50
kingsfoil:20, rainbow quartz:30, mead:50
```

***Analysis:***

- First recipe: ***valid*** (4 ingredients, totals 115, max is 40%)
- Second recipe: ***invalid*** (cilantro is 60% of the total)
- Third recipe: ***invalid*** (only 2 ingredients)
- Fourth recipe ***valid*** (3 ingredients, totals 100, max is 50%)
- ***Example Output:*** `2`
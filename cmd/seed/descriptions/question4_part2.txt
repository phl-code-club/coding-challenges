The professor would like a report, including the average score of the class. Here is the system for scoring recipes:

- ***Base score*** = sum of all ingredient amounts
- ***Bonus***: If recipe contains both "kingsfoil" and "rainbow quartz", multiply total by 2
- ***Penalty***: If recipe contains "foxglove" or "mercury", subtract 25 points for each from total
- ***Minimum score is 0*** (scores cannot go negative)

***Your task:*** Calculate the average score (Rounded down to the nearest integer) for valid recipes from Part 1.

Using the example from part 1 we get:

- First recipe: ***valid*** (initial amount 115, -25 penalty, total 90)
- Second recipe: ***invalid***
- Third recipe: ***invalid***
- Fourth recipe ***valid*** (initial amount 100, 100 bonus, total 200)
- ***Example Output:*** `145`
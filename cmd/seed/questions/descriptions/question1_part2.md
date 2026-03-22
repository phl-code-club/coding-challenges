The council is surprised by how large that value is. It turns out that this list is the list of _all_ spells, not just allowed. You need to filter out any forbidden spells before you add to the running total.

***Your task:*** For the same input, count the segments and characters, but omit spells that include any of the following words:

- `hex`
- `curse`
- `poison`
- `skull`
- `death`
- `trouble`
- `error`
- `gun`
- `bomb`
- `evil`

***This will use the same example from Part 1***

***Analysis:***

- `skeleton curse`: Contains forbidden word "curse", skip
- `fireball`: 1 segment, 8 characters
- `transmute metal to wood`: 4 segments, 19 characters
- ***Output Format:*** `Total segments, Total characters`
- ***Example Output:*** `5,27`
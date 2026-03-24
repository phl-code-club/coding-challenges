***Your input:*** A list of coordinates with latitude, longitude, and type in the form: `(51.4934, 0.0098): landmark`

***Your task:*** Counting the number of valid locations listed on the scroll,
indicated by a type of `landmark` or `clue`. Output the sum of all landmarks
and clues multiplied by 10 to move on to your next challenge.

***Example Input:***

```
(51.4934, 0.0098): landmark
(23.7275, 37.9838): clue
(40.7128, 74.0060): trap
(35.6762, 139.6503): thief
(48.8566, 2.3522): merchant
(41.9028, 12.4964): ship
(55.7558, 37.6173): campsite
(19.4326, 99.1332): rumor
```

***Analysis:***

- `landmark`: 1
- `clue`: 1
- ***Output Format:*** `(landmark count + clue count) * 10`
- ***Example Output:*** `20`

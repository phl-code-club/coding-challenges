***Your input:*** A list of pictograms, each made up of 5 lines of 5 characters.
Each character is one of: ` `, `\`, `/`, `|`, `~`, `-`, `_`, `^`, and `X`.

The valid pictograms are:

River:

```
 /~~/
/~~/
\~~\
 \~~\
 /~~/
```

Star:

```
\ | /
 \|/
--X--
 /|\
/ | \
```

Gem:

```
 ___
//|\\
\\|//
 \_/ 
  V  
```

Forest:

```
  ^  
 /^\ 
//^^\
 ||^\
 ||| 
```

***Your task:*** Count the valid and invalid pictograms, then multiply those two
counts together.

***Example input:***

```
~ / /
 \/ |
-X  -
|/|^_ 
/~|__
 /~~/
/~~/ 
\~~\ 
 \~~\
 /~~/
\ | /
 \|/ 
--X--
 /|\ 
/ | \
~ / /
 \/ |
-X  -
|/|^_ 
/~|__
```

***Analysis:*** This input represents an invalid River and a valid Star.

- ***Output Format:*** `valid pictograms * invalid pictograms`
- ***Example Output:*** `4` (2 valid × 2 invalid = 4)

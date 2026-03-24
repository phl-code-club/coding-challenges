***Your input:*** A list of hieroglyphs, each made up of 5 lines of 5 characters. These hieroglyphs are separated by a blank line. Each character is one of: " ", "\", "/", "|", "~", "-", "_", "^".

The valid hieroglyphs are:

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
/___\
\___/
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

***Your task:*** Count the valid hieroglyphs and the invalid ones, then multiply those two counts together.

***Example input:***

```
 /~~/
/~~/
\-~\
 \~~\
 /~~/

\ | /
 \|/
--X--
 /|\
/ | \
```

***Analysis:*** This input represents an invalid River and a valid Star.

- ***Output Format:*** `product of valid and invalid hieroglyphs`
- ***Example Output:*** `1` (1 valid × 1 invalid = 1)


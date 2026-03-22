Some runes may be damaged - certain `@` symbols may have been changed to `~` (but `~` never becomes `@`). A damaged pattern matches a key pattern if you can transform it into that pattern by only changing one `~` to `@`.

The final rune in the previous example can be fixed.

```
~~~
~~~
@@~
```

This could match Earth (`~~~|~~~|@@@`) by adding `@` symbol to the end of the third line.

***Your task:*** For each pattern, match it to the key pattern that you can construct with only one change of a `~` to a `@`. Count the total occurrences of each pattern type including the damaged ones that can be fixed and output in the same format as Part 1.

- ***Example Output:*** `1,1,1,1`
***Your input:*** 3 lines of characters that form 3x3 sets which you will need to decipher into runes. The characters are `@` and `~`.

You must decode the input using the Cryptarch's key and output the count of each rune. The Cryparch's key is:

- `Air`:

```
~~~
~@~
~~~
```

- `Fire`:

```
@~@
@~@
~@~
```

- `Earth`:

```
~~~
~~~
@@@
```

- `Water`:

```
~@~
~~~
@~@
```

***Example Input:***

```
~~~~@~@~@~~~
~@~~~~@~@~~~
~~~@~@~@~@@~
```

This input represents `Air`, `Water`, `Fire`.

***Your task:*** Count the total occurrences of each rune.

- ***Output Format:*** `Air count, Fire count, Earth count, Water count`
- ***Example Output:*** `1,1,0,1`
- ***_Note:_*** that the final rune is incomplete and therefore is not counted
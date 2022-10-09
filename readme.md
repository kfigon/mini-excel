# Mini excel engine

Idea from coding stream from tsoding: https://www.youtube.com/watch?v=HCAgvKQDJng

This is a simplified excel engine without a GUI.

It accepts CSV input and evaluates all expressions.
In case of cyclic dependency it reports errors.

```csv
A,B,C,D,E
1,2,3,4,5
=A1+B1,,3,=A2+C2,
```

output:

```csv
A,B,C,D,E
1,2,3,4,5
3,,3,6,
```

expressions are evaluated using modified shunting yard algorithm
https://en.wikipedia.org/wiki/Shunting_yard_algorithm
https://aquarchitect.github.io/swift-algorithm-club/Shunting%20Yard/


missing 
* nested expressions
* cyclic dependency detector 
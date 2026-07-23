# Prolog Problem Set

The questions on the Prolog quiz will mainly be variations of the questions
below, or questions that are similar. Avoid using built-in functions when
possible: stick to elementary Prolog as used in the notes.

Please post your answers to the discussion board to share with other students.

Important: Treat these problem-sets as non-AI activities! Turn off all AI
support and try to figure them out yourself. Having AI or another student do
this for you will not help you learn. You must do the learning yourself!

## Question 1

Create a knowledge base with the following facts and rule:

- Marge is taller than Homer.
- Homer is taller than Bart.
- Bart is taller than Lisa.
- Lisa is taller than Maggie.
- If person A is taller than person B, and person B is taller than person C,
  then person A is taller than person C.

Use two predicates: `taller(A, B)` for the facts, and `is_taller(A, B)` for the
rule. The `is_taller` is tricky because, depending on how you write it, it can
get stuck infinite loops after the first solution. So try to find a version of
`is_taller` that doesn't get stuck in infinite loops.

Next add this rule:

- A is shorter than B if B is taller than A. Call it `is_shorter(A, B)`.

Test your knowledge base with the following queries:

- Is Marge taller than Maggie?
- Is Bart taller than Lisa?
- Who is taller than Lisa?
- Who is shorter than Homer?

## Question 2

Write a predicate called `prod_list(List, Product)` that returns the product of
all the elements of a list. For example:

```prolog
?- prod_list([2, 3, 2], X).
X = 12.
```

Implement this using recursion and basic Prolog. *Don't* use any built-in
predicates that do the work for you.

## Question 3

Write a predicate called `max_list_(List, Max)` (note the underscore on the
end!) that returns the maximum element of a list. For example:

```prolog
?- max_list_([9,2,10,6,4], X).
X = 10 ;
false.
```

Don't use the built-in `max_list` predicate in your solution. You can use the
built-in `max` predicate to compare two numbers, e.g.:

```prolog
?- M is max(4, 5).
M = 5.
```

## Question 4

Write a predicate `my_last(List, Last)` that returns the last element of a list.
For example:

```prolog
?- my_last([1,2,3], X).
X = 3 ;
false.

?- my_last([once, upon, a, time], X).
X = time ;
false.
```

Don't use the built-in `last` predicate in your solution.

## Question 5

Write a predicate called `add_lists(A, B, C)` that returns a new list the sum of
lists `A` and `B`. For example:

```prolog
?- add_lists([2,1,0], [4, 2, 1], L).
L = [6, 3, 1].
```

You can assume that `A` and `B` are the same length.

## Question 6

Write a predicate called `make_list(N, X, List)` that returns a list of `N`
elements, all of which are `X`. For example:

```prolog
?- make_list(3, a, L).
L = [a, a, a].
```

You can assume that `N` is an integer that is 0 or greater.

Implement this using recursion and basic Prolog. *Don't* use any built-in
predicates that do the work for you.

## Question 7

Write a predicate called `make_range(N, List)` that returns a list `[0, 1, 2,
..., N-1]`. For example:

```prolog
?- make_range(5, L).
L = [0, 1, 2].
```

You can assume that `N` is an integer that is 0 or greater.

Implement this using recursion and basic Prolog. *Don't* use any built-in
predicates that do the work for you.

## Question 8

Write a predicate called `factorial(N, Result)` that returns the factorial of `N`. For example:

```prolog
?- factorial(5, X).
X = 120 ;
false.
```

You can assume that `N` is an integer that is 0 or greater.

Implement `factorial` *without* using recursion, and instead using `prod_list`
and `make_range`.

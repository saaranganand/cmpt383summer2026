# Prolog 

These notes are based on ["Learn Prolog Now!" by Patrick Blackburn, Johan Bos,
and Kristina Striegnitz](https://lpn.swi-prolog.org/lpnpage.php?pageid=online).

## Problem Set

[Here is the Prolog problem set](prolog_questions.md).

## Introduction to Prolog

**Prolog** is a programming language inspired by research in logic programming
(an approach to programming based on formal logic). Prolog can be thought of as
a language whose programs are collections of facts and rules that can be queried
(similar to a database). Prolog can be a good choice for so-called **rule-based
systems**, such as **expert systems**, where the program is a collection of
facts and rules that can be queried.

In the 1970s and 80s, it was a popular language for logic-based AI research
based on logic. It has good symbol and list manipulation capabilities, supports
automatic backtracking, and even has a special syntax for formal grammars. 

## Basic Prolog

A Prolog program consists of three main things: **facts**, **rules**, and
**queries**. Facts and rules are stored in a **knowledge base** that can be
queried. Queries can appear in a program, or be entered interactively at the
Prolog prompt.

### Facts

For example, here is a knowledge base consisting of four facts:

```prolog
dog(rover).     % fact 1: rover is a dog
dog(fido).      % fact 2: fido is a dog
cat(whiskers).  % fact 3: whiskers is a cat
cat(mittens).   % fact 4: mittens is a cat
```

Note the use of the dot (`.`) to indicate the end of a fact.

Now we can query the knowledge base, i.e. ask Prolog to check if a fact is true.
For example, we can ask if `rover` is a `dog`:

```prolog
?- dog(rover).
true.
```

`?-` is the Prolog interpreter prompt, and what comes after is the query typed
by the user.

The answer to this query is `true`, because `rover` is indeed a dog in our
knowledge base. We can also ask if `spot` is a `dog`:

```prolog
?- dog(spot).
false.
```

`spot` is *not* a dog in our knowledge base, and so the answer to this query is
`false`.

**Syntax note** `dog(rover)` is an example of a  **compound term**, and `rover`
is the first argument. Both `dog` and `rover` are **atoms**, and atoms must
start with a lowercase letter (uppercase letters are reserved for variables).

### Rules

Here's another knowledge base:

```prolog
eating(homer, donut).                 % fact: homer is eating a donut
happy(homer) :- eating(homer, donut). % rule 1: homer is happy if he is eating a donut
snoring(homer) :- happy(homer).       % rule 2: homer is snoring if he is happy
```

This has one fact and one rule. We can see that `homer` is eating a donut:

```prolog
?- eating(homer, donut).
true.
```

But not a `carrot`:

```prolog
?- eating(homer, carrot).
false.
```

Consider the first rule:

```prolog
happy(homer) :- eating(homer, donut).
```

`:-` is can be read as "if". Logically, this says *if* `homer` is eating a
donut, *then* `homer` is happy. Since it is a fact in the knowledge base that
`homer` is eating a donut, Prolog automatically deduces that `homer` is happy:

```prolog
?- happy(homer).
true.
```

Now consider the second rule:

```prolog
snoring(homer) :- happy(homer).
```

Logically, this says *if* `homer` is happy, then `homer` is snoring. We know
from the first rule that `homer` is happy, so Prolog will automatically deduce
that `homer` is snoring:

```prolog
?- snoring(homer).
true.
```

If you were to change the fact to `eating(homer, carrot)`, then both
`happy(homer)` and `snoring(homer)` could *not* be deduced. Try it!

### More on Rules

Here's another knowledge base with slightly more complex rules:

```prolog
sandy(alan).           % fact 1: alan is covered in sand
wet(alan).             % fact 2: alan is wet
eating(alan, hot_dog). % fact 3: alan is eating a hot dog

% rule 1: alan is annoyed if he is covered in sand and wet
at_beach(alan) :- sandy(alan), wet(alan). 

% rule 2: alan is happy if he is at the beach and eating a hot dog
happy(alan) :- at_beach(alan), eating(alan, hot_dog).
```

Rule 1 is this:

```prolog
at_beach(alan) :- sandy(alan), wet(alan).
```

Logically, it says *if* `alan` is covered in sand *and* he's wet, then `alan` is
at the beach. The comma `,` means *and*.

Since `sandy(alan)` and `wet(alan)` are both facts in the knowledge base, Prolog
automatically deduces that `alan` is at the beach:

```prolog
?- at_beach(alan).
true.
```

Rule 2 is:

```prolog
happy(alan) :- at_beach(alan), eating(alan, hot_dog).
```

Logically, it says *if* `alan` is at the beach *and* eating a hot dog, then
`alan` is happy. Prolog automatically deduces this from the facts and rules:

```prolog
?- happy(alan).
true.
```

### Queries with Variables

Consider this knowledge base which has seven facts:

```prolog
woman(mia).         % fact 1: mia is a woman
woman(jody).        % fact 2: jody is a woman
woman(yolanda).     % fact 3: yolanda is a woman

loves(vincent, mia).         % fact 4: vincent loves mia
loves(marsellus, mia).       % fact 5: marsellus loves mia
loves(pumpkin, honey_bunny). % fact 6: pumpkin loves honey_bunny
loves(honey_bunny, pumpkin). % fact 7: honey_bunny loves pumpkin
```

Consider this query:

```prolog
?- woman(X).
X = mia ;
X = jody ;
X = yolanda.
```

Here `X` is a variable. Prolog variables always start with a capital letter. The
query `woman(X)` asks Prolog to find all the women in the knowledge base, i.e.
all values of `X` that make `woman(X)` true.

When you run this in the Prolog interpreter, it stops after the first answer "X
= mia". Then you (the user) can either type <enter> to halt the query, or  `;`
to get the next answer (if there is a next answer). In this
particular case, there are three values of `X` that make the query true: `mia`,
`jody`, and `yolanda`.

Now lets ask who loves `mia`:

```prolog
?- loves(X, mia).
X = vincent ;
X = marsellus ;
false.
```

We could also ask who are the women `marsellus` loves:

```prolog
?- woman(X), loves(marsellus, X).
X = mia ;
false.
```

This asks Prolog to find all the values of `X` such that `X` is a woman *and*
`marsellus` loves `X`.

Here's another knowledge base with a more elaborate rule:

```prolog
loves(homer, marge).    % fact 1: homer loves marge
loves(moe, marge).      % fact 2: moe loves marge
loves(homer, donuts).   % fact 3: homer loves donuts
loves(wiggum, donuts).  % fact 4: wiggum loves donuts

% rule: X is jealous of Y if X loves Z and Y loves Z
jealous(X, Y) :- loves(X, Z), loves(Y, Z).
```

Logically, the rule says that *if* `X` loves `Z` *and* `Y` loves `Z`, then `X`
is jealous of `Y`.

We can asks who `homer` is jealous of like this:

```prolog
?- jealous(homer, Y).
```

Before reading the answer, look at the knowledge base and try to decide what the
values for `Y` will be. It's a bit of a trick question!

The answer is:

```prolog
?- jealous(homer, Y).
Y = homer ;
Y = moe ;
Y = homer ;
Y = wiggum.
```

The first value `Y = homer` means that `homer` is jealous of himself (!). Let's
look at this carefully. The query `jealous(homer, Y)` sets the `X` in the rule
to `homer`:

```prolog
jealous(homer, Y) :- loves(homer, Z), loves(Y, Z).
```

The possible values for `Z` are `marge` and `donuts`. So if we set `Z` to
`marge` the query becomes:

```prolog
jealous(homer, Y) :- loves(homer, marge), loves(Y, marge).
```

What could `Y` be? `Y` must be some value that satisfies the query 
`loves(Y, marge)`. There are two possible values for `Y`: `homer` and `moe`, and
this is why we get two answers: `Y = homer` and `Y = moe`.

Prolog keeps going. We also get `Y = homer` and `Y = wiggum` because of the
facts `loves(homer, donuts)` and `loves(wiggum, donuts)`.

Suppose we don't want to allow someone to be jealous of themselves. We could
modify the rule to:

```prolog
jealous(X, Y) :- loves(X, Z), loves(Y, Z), X \= Y.
```

The `\=` operator means "not equal to", so `X \= Y` means "X is not equal to Y".

Now query does not include `Y = homer`:

```prolog
?- jealous(homer, Y).
Y = moe ;
Y = wiggum.
```

**Be careful!** Rules often have unexpected solutions and need to be carefully
crafted to be what you want.

### A Bits Knowledge Base

Here is a knowledge base consisting of two facts:

```prolog
bit(0).  % fact 1: 0 is a bit
bit(1).  % fact 2: 1 is a bit
```

This says 0 and 1 are bits. Now consider this query:

```prolog
?- bit(X).
X = 0 ;
X = 1.
```

This gives all possible bit values for `X`, either 0 or 1.

This query asks for all possible bit values for two variables, `X` and `Y`:

```prolog
?- bit(X), bit(Y).
X = Y, Y = 0 ;
X = 0,
Y = 1 ;
X = 1,
Y = 0 ;
X = Y, Y = 1.
```

You can see there are three `;`s, which means there are four answers in total.
The first answer is `X = Y` and `Y = 0`, which means `X` and `Y` are both 0. The
second answer is `X = 0, Y = 1`, the third answer is `X = 1, Y = 0`, and the
fourth answer is `X = Y` and `Y = 1`, i.e. `X` and `Y` are both 1.

This essentially gives us all possible pairs of bit values for `X` and `Y`. If
ask for three different variables, `X`, `Y`, and `Z`, we get all 8 possible
triplets of bit values:

```prolog
?- bit(X), bit(Y), bit(Z).
X = Y, Y = Z, Z = 0 ;
X = Y, Y = 0,
Z = 1 ;
X = Z, Z = 0,
Y = 1 ;
X = 0,
Y = Z, Z = 1 ;
X = 1,
Y = Z, Z = 0 ;
X = Z, Z = 1,
Y = 0 ;
X = Y, Y = 1,
Z = 0 ;
X = Y, Y = Z, Z = 1.
```

## Terminology

It's useful to know terminology for Prolog programs. The following are
considered **terms**:

- **Constants**
  - **Atoms**: Strings of characters like `homer`, `hot_dog`, and so on. Atoms
    start with a lowercase letter, or an underscore `_`.
  - **Numbers**: E.g. `104.2`, `-25.3`, `305`, and so on.
- **Variables**: Strings of characters like `X`, `Y`, `Name`, and so on.
  Variables start with a capital letter. The underscore `_` is also allowed, and
  has a special meaning.
- **Complex terms**: E.g. `loves(homer, marge)`, `at_beach(alan)`, and so on. We
  sometimes call these **predicates**, and the arguments are called
  **arguments**. For example, in `loves(homer, marge)`, `loves` is the predicate
  and `homer` and `marge` are the arguments. The **arity** of a predicate is the
  number of arguments it has. For example, `loves` has arity 2 because it has 2
  arguments, and `at_beach` has arity 1 because it has 1 argument. In Prolog, we
  sometimes write `loves/2` to mean that `loves` has arity 2.

## Unification

Unification is a kind of pattern matching. Given two terms, Prolog can
automatically find values for the variables in them that make them equal, or
determine that they cannot be made equal.

Suppose we have the terms `dog(rover)` and `dog(X)`. We say that these terms
**unify** because it is possible to find a value for `X` that makes the terms
equal. In this case, the terms are equal if we set `X` to `rover`.

Given two terms, **unification** is the process of finding values for the
variables in them that make the terms equal. 

It's quite possible that two terms cannot be unified. For example, `dog(rover)`
and `cat(X)` cannot be unified because there is no value for `X` that makes the
two terms equal. Also, `dog(rover)` and `dog(spot)` cannot be unified because
`rover` and `spot` are different atoms.

In Prolog, you can unify two terms like this:

```prolog
?- dog(rover) = dog(X).
X = rover.
```

If the terms can't be unified, Prolog returns `false`:

```prolog
?- dog(rover) = cat(X).
false.
```

The result of unification could sometimes be that two variables are equal, e.g.:

```prolog
?- dog(X) = dog(Y).
X = Y.
```

Here `X` and `Y` are unified, i.e. they are the same variable.

In this example, Prolog returns `false` because it is not possible for `X` to be
two different values at the same time:

```prolog
?- X=cat, X=dog.
false.
```

This also fails to unify:

```prolog
?- loves(X, X) = loves(homer, marge).
false.
```

This means there is no value for `X` that can make the two terms equal.

Here is a high-level description of the unification algorithm ([from the textbook](https://lpn.swi-prolog.org/lpnpage.php?pagetype=html&pageid=lpn-htmlse5)):

1. If term1 and term2 are constants, then term1 and term2 unify if and only if
   they are the same atom, or the same number.
2. If term1 is a variable and term2 is any type of term, then term1 and term2
   unify, and term1 is instantiated to term2. Similarly, if term2 is a variable
   and term1 is any type of term, then term1 and term2 unify, and term2 is
   instantiated to term1. (So if they are both variables, they’re both
   instantiated to each other, and we say that they share values.)
3. If term1 and term2 are complex terms, then they unify if and only if: 
   1. They have the same functor and arity, and
   2. all their corresponding arguments unify, and
   3. the variable instantiations are compatible. (For example, it is not
      possible to instantiate variable X to mia when unifying one pair of
      arguments, and to instantiate X to vincent when unifying another pair of
      arguments.)
4. Two terms unify if and only if it follows from the previous three clauses
   that they unify.

Finally, we note that Prolog's unification algorithm omits something known as
the **occurs check**. This is a check to prevent infinite loops that can arise
in some situations. It should not be an issue for us, but [see the textbook](https://lpn.swi-prolog.org/lpnpage.php?pagetype=html&pageid=lpn-htmlse5) for
more details of when it can happen, and what you can do about it.

## Proof Search

Consider this knowledge base:

```prolog
f(a).
f(b).

g(a).
g(b).

h(b).

k(X) :- f(X), g(X), h(X).
```

The rule can be interpreted as follows: to prove `k(X)`, we need to prove
`f(X)`, `g(X)`, and `h(X)`. We will often returns to the terms on the right hand
side of the `:-` as *goals*. So to prove `k(X)`, we need to satisfy all the
goals `f(X)`, `g(X)`, and `h(X)`.

You can probably see that the only value of `X` that can make the true is `b`.

Prolog searches for the answer like this:

- You make the query `k(X)`, i.e you ask for a value of `X` that makes it true.

- Prolog checks the goals in left to right order. First it finds a that makes
  `f(X)` true. It sees that `f(a)` is a fact, so it unifies `X` with `a`. Now
  that `X = a`, it checks that `g(a)` is true, which it is. Then it checks if
  `h(a)` is true. `h(a)` is *not* a fact in the knowledge base, so it fails.

- Since `h(a)` failed, Prolog *backtracks*  to the very first goal, `f(X)`. It
  then searches the knowledge base to see if there is another value of `X` that
  makes `f(X)` true. It finds that `f(b)` is a fact, so it unifies `X` with `b`.
  Now that `X = b`, it checks that `g(b)` is true, which it is. Then it checks
  if `h(b)` is true. `h(b)` is a fact in the knowledge base, so it succeeds.

Whenever there is a choice for a variable, Prolog picks the first choices and
remembers this choice point. If it turns out that the choice was wrong, Prolog
can backtrack to the choice point and try a different choice.

Prolog always backtracks to the most recent choice point.

## Lists

Prolog has lists, which are sequences of 0 or more terms. These are all lists:

- `[]` (the empty list)
- `[7, 5, 3]`
- `[the, big, red, dog]`
- `[wet(alan), big(red(cliff)), 7, [a,b]]`

Use `|` to get the first and rest of a list:

```prolog
?- [Head|Tail] = [7, 5, 3].
Head = 7,
Tail = [5, 3].
```

Short variable names are often used:

```prolog
?- [H|T] = [7, 5, 3].
H = 7,
T = [5, 3].
```

If you only want, say, the first element you can use the anonymous variable `_`:

```prolog
?- [H|_] = [7, 5, 3].
H = 7.
```

Or:

```prolog
?- [_|T] = [7, 5, 3].
T = [5, 3].
```

## The `member` Predicate

Let's write a predicate that checks if a term is a member of a list. We will
call it `member(X, List)`. If `X` is a member of `List`, the predicate returns
`true`. Otherwise, it returns `false`:

```prolog
member(X, [X|_]).                 % base case
member(X, [_|T]) :- member(X, T). % recursive case
```

There are two rules: a base case and a recursive case. The base case is
satisfied if the first element of the list is equal to `X`. Otherwise, the
recursive case is tried. The second rule recursively calls `member` on the rest
of the list. Eventually, either `X` is found, or the list is empty (in which
case neither rule is satisfied).

It works as you might expect:

```prolog
?- member(red, [the, big, red, dog]).
true .

?- member(green, [the, big, red, dog]).
false.
```

But it also does this:

```prolog
?- member(X, [the, big, red, dog]).
X = the ;
X = big ;
X = red ;
X = dog ;
false.
```

This essentially makes `X` iterate through all the elements of the list.

You could also do this:

```prolog
?- member(cat, [the, big, red, X]).
X = cat ;
false.
```

This says that `member` succeeds if `X` has the value `cat`.

You can find the intersection (common elements) of two lists like this:

```prolog
?- member(X, [1, 2, 3, 4]), member(X, [4, 6, 7, 1, 8]).
X = 1 ;
X = 4 ;
false.
```

## Length of a List

Here's a predicate that calculates the length of a list:

```prolog
len([], 0).
len([_|T], N) :- len(T, X), N is X + 1.
```

The first rule is the base case: it says that the empty list has length 0.

The second rule says that the length of a list is 1 plus the length of the rest
of the list.

Notice the unusual way that Prolog does arithmetic. Arithmetic is done using the
`is` operator. So `N is X + 1` means "calculate the value of `X + 1` and assign
it to `N`".

## Appending Lists

The `append` predicate concatenates two lists:

```prolog
append([], B, B).
append([H|T], B, [H|Result]) :- append(T, B, Result).
```

The first rule is the case case. It says that the empty list plus any other is
the list itself.

The second rule is the recursive case, and it first appends the tail of the
first list to the second list `B`, putting the result in `Result`. Then it
prepends the head of the first list to the result for the final result.

With it you can append two lists like this:

```prolog
?- append([1,2], [3,4,5], X).
X = [1, 2, 3, 4, 5].
```

Or you could do this:

```prolog
?- append(X, [3,4,5], [1,2,3,4,5]).
X = [1, 2].
```

Or even, which finds all pairs of lists that append to make the third list:

```prolog
?- append(X, Y, [1,2,3,4,5]).
X = [],
Y = [1, 2, 3, 4, 5] ;
X = [1],
Y = [2, 3, 4, 5] ;
X = [1, 2],
Y = [3, 4, 5] ;
X = [1, 2, 3],
Y = [4, 5] ;
X = [1, 2, 3, 4],
Y = [5] ;
X = [1, 2, 3, 4, 5],
Y = [] ;
false.
```

## Reversing a List

Here is a simple but inefficient way to reverse a list:

```prolog
rev([], []).
rev([H|T], R):- rev(T, RevT), append(RevT, [H], R).
```

The first rule is the base case: the empty list reversed is the empty list.

The second rule  reverses the tail of the list, puts the result in `RevT`, and
then it appends the head of the list to the result for the final result.

While this works, it is inefficient because it creates a new list for each step
of the recursion. A better way to reverse a list is to use an *accumulator*:

```prolog
% initial called to rev2 with empty accumulator
rev2(L, R) :- rev2(L, [], R).

rev2([], Acc, Acc).
rev2([H|T], Acc, R):- rev2(T, [H|Acc], R).
```

The idea of adding `Acc` is that it *accumulates* the reversed list as the
function recurses. On each call one element is added to the accumulator at the
front, which has the effect of reversing the list. Using accumulators is a
common technique for speeding up recursive functions in Prolog and other
languages.

% sum_min.pl

% sum of a list
sum([], 0).
sum([H|T], S) :- sum(T, S2), S is H + S2.

% minimum of a list
min([X], X).
min([H|T], M) :- min(T, M2), M is min(H, M2).

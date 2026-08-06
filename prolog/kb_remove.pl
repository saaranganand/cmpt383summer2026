% kb_remove.pl

remove(X, [X], []).
remove(X, [X|T], Result) :- remove(X, T, Result).
remove(X, [H|T], [H|Result]) :- H \= X, remove(X, T, Result).
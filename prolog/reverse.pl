% kb_reverse.pl

% naive implementation
rev([], []).
rev([H|T], R):- rev(T, RevT), append(RevT, [H], R).

% better implementation (with an accumulator)
rev2(L, R) :- rev2(L, [], R).
rev2([], A, A).
rev2([H|T], A, R):- rev2(T, [H|A], R).

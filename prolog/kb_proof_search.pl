% kb_proof_search.pl

f(a).
f(b).

g(a).
g(b).

h(b).

% to prove K(X), we need to find a value for X that makes f(x), g(x), and h(x)
% all true.
k(X) :- f(X), g(X), h(X).

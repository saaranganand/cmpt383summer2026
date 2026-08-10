% kb_jealous_homer.pl

loves(homer, marge).    % fact 1: homer loves marge
loves(moe, marge).      % fact 2: moe loves marge
loves(homer, donuts).   % fact 3: homer loves donuts
loves(wiggum, donuts).  % fact 4: wiggum loves donuts

% rule: X is jealous of Y if X loves Z and Y loves Z
jealous(X, Y) :- loves(X, Z), loves(Y, Z).

% rule: X is jealous of Y if X loves Z and Y loves Z, and X is not Y
% jealous(X, Y) :- loves(X, Z), loves(Y, Z), X \= Y.

% kb_bits.pl

bit(0).
bit(1).

bitstring(0, []).
bitstring(N, [B|Bs]) :-
    N > 0,
    bit(B),
    N1 is N - 1,
    bitstring(N1, Bs).
%
% To get a lit of all 3-bit bitstrings:
% ?- findall(Bs, bitstring(3, Bs), L).
%
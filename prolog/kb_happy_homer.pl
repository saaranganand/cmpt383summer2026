% kb_happy_homer.pl

eating(homer, donut). 

% rule 1: homer is happy if he is eating a donut
happy(homer) :- eating(homer, donut). 

% rule 2: homer is snoring if he is happy
snoring(homer) :- happy(homer).

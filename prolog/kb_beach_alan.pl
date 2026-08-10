# kb_beach_alan.pl

sandy(alan).           % fact 1: alan is covered in sand
wet(alan).             % fact 2: alan is wet
eating(alan, hot_dog). % fact 3: alan is eating a hot dog

% rule 1: alan is at_beach if he is hungry and wet
at_beach(alan) :- sandy(alan), wet(alan). 

% rule 2: alan is happy if he is at the beach and eating a hot dog
happy(alan) :- at_beach(alan), eating(alan, hot_dog).

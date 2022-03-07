# Daily Murphy's Law

Display a murphy's law at the windows start up.

1. Get laws list from _tech_laws.file_.
2. Get laws list count.
3. Generate a random number in range of the list count
4. Check if the number is in _stat file_.
5. If it was not in _stat file_, proceed, otherwise, go to step 3.
5.1. If attemps became equal to laws list count, it means all laws are had been showed once, so clear _stat file_ and start go to step 3.
6. Show the law at that position.
7. Insert the number into _stat file_.
8. Terminate the program.

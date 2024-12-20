# advent-of-code-2024

Code and tests colocated for convenience.

Specific tests can be run like so
```sh
go test -v -run=^TestSolutions/day_1/

# and specify sub tests
go test -v -run=^TestSolutions/day_1/part_1
go test -v -run=^TestSolutions/day_1/part_1/solution
```

Input is downloaded automatically if not found. To auth with adventofcode.com we need a session cookie. Get this cookie and create a file at `input/.session`. The contents of the file should be the session cookie value. Note the `input/` directory is not tracked by git to stop leaking credentials and test inputs.

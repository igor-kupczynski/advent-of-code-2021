# advent-of-code-2021
https://adventofcode.com/2021

* Some days are split into two parts
* Some days are bundled together and display the results for both parts
* Run the programs:
```sh
cd day<num>
cat example.txt | go run .
cat input.txt | go run .
```
* Starting from day 6 the program also accepts the name of the file as first positional argument, e.g.
`go run . input.txt`


## Notes

To revisit and compare notes:
- Day 8

### Day 8

- I've solved this on paper and then hardcoded the rules (imperative)
  - I feel think a declarative solution  would be better (prolog, core.logic)
- Implemented a bit set, which added a lot of complexity and the input is rather small, so probably overengineered


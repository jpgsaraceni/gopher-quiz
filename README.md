# Quiz [WIP]

This project is a *partial* solution to Exercise #1 of [Gophercises](https://gophercises.com/).

It consists of a CLI script that reads the `problems.csv` file in the root
of the directory, parses each record as a question-answer pair (each line
must contain exactly 2 fields). Then it asks each question and compares
with the answer. At the end, shows the score.

Lacking features:
* Reading input file
* Timer

## Run (requires Go)

Just clone the repo, enter and run `make`:

```shell
git clone https://github.com/jpgsaraceni/gopher-quiz.git && cd gopher-quiz && make
```

```shell
make
```
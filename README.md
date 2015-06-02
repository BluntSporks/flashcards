# flashcards
Golang program to provide flashcard-type quizzes from the command line

## Status
Ready to use

## Decks
Question and answer files should be provided in CSV format. Examples are provided in the decks folder.

The first column should be the question.

The second column should be the answers. Alternate answers should be separated with pipes.

## Program interaction
Provide a limit with the -lim option and that number of questions will be used in the session.

Each question will be printed out and wait for you to type an answer. Score is given in the form of percent correct.

Some questions can have more than one answer separated with pipes as described in the Decks section above. In that
case, your score is given for the answer choice that best matches your response.

Matching percentage is calculated using the [Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance)
function. This basically means that it gives you partial score for the percentage of characters in your answer that
match the original.

If your score is not 100, the correct answers are printed for each question.

Once the session is over, your average score is printed.

You can quit a quiz early by providing an empty answer. In that case, you are given 0 for each of the remaining
questions.

## Usage
Usage:

    flashcards [options] FILENAME

Options:

    lim=LIMIT  Limit of maximum questions to include in the current session [default=20]

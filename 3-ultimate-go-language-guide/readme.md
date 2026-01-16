# Go language guide

This course focus on the go language guide from philosophy to how the language practically works.

## Module 1 Design guidelines

- Go is language that is designed to balance execution performance and development speed.
- Go developers should focus on integrity, readability and simplicity.

- We achieve integrity by
  - reducing lines of code
  - ensuring transparency
  - focusing on accuracy be for performance optimization
  
- We have 2 courses of performance cost:
  - external ( io operations, network calls ...): measured in milliseconds.
  - internal ( garbage collection, algorithms, memory allocation...): measured in nano to microseconds(but tends to compound especially in tight loops)

- code review should be treated as continues draft for improvement to alleviate the need for perfection on first go.
- an average developer can only really manage a mental model of a code base with about 10000 lines of code.
- an average developer in a team should have full context of the codebase at all times
- Debugger should be the last resort when going through the codebase. over reliance of Debugger can hurt codebase mental model.

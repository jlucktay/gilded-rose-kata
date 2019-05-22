# James' notes

- Spent a couple of minutes re-doing some of `GildedRose()` with assignment operators, for clarity.
- Spent ~25m writing up assertions in the test, to have a known good baseline to work from.
- Deduced correct branch of logic to alter for conjured item case.
- Small refactor on test: originalQuality as a map instead of discrete vars.
- Checked test coverage, and need to increase across some more branches.
- At this point, the task is _technically_ complete, but the `GildedRose()` func is in dire need of a refactor to
reduce cyclomatic complexity.
- Pruned a couple of redundant `if` branches.
- Another ~25m revising and adding to tests, doing a second pass to make sure all requirements are asserted on.
- Checked test coverage again, and made some notes on my findings.
- Looking at the clock, I made the decision to down tools and wrap up for submission; looking forward to discussing
potential next steps!

## Queries

With the following two points from the README:

> Once the sell by date has passed, quality degrades twice as fast.

and

> "Aged Brie" actually increases in quality the older it gets.

does this mean "Aged Brie" should go up by **two** quality, or just one?

## Further discussion

The current state of the test only deals with the passing of a single day.
I could have planned ahead a little better and iterated across multiple days which would have covered more `if`
branches - for both backstage pass quality and the sell-in values for all items - within `GildedRose()`.

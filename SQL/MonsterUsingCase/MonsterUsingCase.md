# Monster Using Case

You have access to two tables named top_half and bottom_half, as follows:<br>
top_half schema<br>
```
id
heads
arms
```
bottom_half schema
```
id
legs
tails
```
You must return a table with the format as follows:<br>

output schema
```
id
heads
legs
arms
tails
species
```
The IDs on the tables match to make a full monster. For heads, arms, legs and tails you need to draw in the data from each table.<br>

For the species, if the monster has more heads than arms, more tails than legs, or both, it is a `'BEAST'` else it is a `'WEIRDO'`. This needs to be captured in the species column.<br>
All rows should be returned (10).<br>

Tests require the use of CASE. Order by species.

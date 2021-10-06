# Challenge: Two actors who cast together the most

https://www.codewars.com/kata/5818bde9559ff58bd90004a2/train/sql <br>
Given the the schema presented below find two actors who cast together the most and list titles of only those movies they were casting together. Order the result set alphabetically by the movie title.<br>

Table film_actor
```SQL
 Column     | Type                        | Modifiers
------------+-----------------------------+----------
actor_id    | smallint                    | not null
film_id     | smallint                    | not null
...
```
```SQL
Table actor
 Column     | Type                        | Modifiers
------------+-----------------------------+----------
actor_id    | integer                     | not null 
first_name  | character varying(45)       | not null
last_name   | character varying(45)       | not null
...
```
```SQL
Table film
 Column     | Type                        | Modifiers
------------+-----------------------------+----------
film_id     | integer                     | not null
title       | character varying(255)      | not null
...
```
The desired output:
```SQL
first_actor | second_actor | title
------------+--------------+--------------------
John Doe    | Jane Doe     | The Best Movie Ever
...
```
first_actor - Full name (First name + Last name separated by a space)<br>
second_actor - Full name (First name + Last name separated by a space)<br>
title - Movie title<br>
Note: actor_id of the first_actor should be lower then actor_id of the second_actor<br>

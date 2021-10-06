select
	concat(a1.first_name, ' ', a1.last_name) as first_actor,
	concat(a2.first_name, ' ', a2.last_name) as second_actor,
	f.title as title
from
	actor a1,
	actor a2,
	film f,
	( 
		select
			fa1.actor_id as actor1,
			fa2.actor_id as actor2,
			COUNT(fa1.film_id) as films_count,
			array_agg(fa1.film_id) as films_id
		from
			film_actor fa1,
			film_actor fa2
		where
			fa1.actor_id < fa2.actor_id
			and fa1.film_id = fa2.film_id
		group by
			fa1.actor_id,
			fa2.actor_id
		order by films_count desc
		limit 1
	) most_casted
where
	f.film_id = any(most_casted.films_id) and
	a1.actor_id = most_casted.actor1 and a2.actor_id = most_casted.actor2
;

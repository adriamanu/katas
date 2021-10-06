SELECT
  th.id,
  th.heads,
  bh.legs,
  th.arms,
  bh.tails,
  CASE
    WHEN th.heads > th.arms or bh.tails > bh.legs THEN 'BEAST'
    ELSE 'WEIRDO'
  END as species
FROM
  top_half th,
  bottom_half bh
WHERE th.id = bh.id
ORDER BY species;

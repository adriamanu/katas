CREATE OR REPLACE VIEW members_approved_for_voucher AS
SELECT m.id, m.name, m.email, SUM(p.price) AS total_spending
FROM members m, sales s, products p
WHERE
  s.member_id = m.id
  AND s.product_id = p.id
  AND s.department_id in (
    select s.department_id
    from sales s, products p
    where s.product_id = p.id
    group by s.department_id
    having SUM(p.price) >= 10000
  )
group by m.id, m.name, m.email
having SUM(p.price) >= 1000
order by m.id;

SELECT * FROM members_approved_for_voucher;
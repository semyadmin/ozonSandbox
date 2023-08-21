SELECT users.id, users.name
FROM orders
LEFT JOIN users
ON orders.user_id = users.id
GROUP BY orders.user_id, users.id
ORDER BY users.name, orders.user_id;
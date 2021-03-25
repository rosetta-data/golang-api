SELECT
    id,
    name,
    description,
    price,
    comment,
    created_at
FROM
    products.products
ORDER BY
    id
LIMIT $1 OFFSET $2

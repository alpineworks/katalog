SELECT
    id
FROM
    katalog.deployments
WHERE
    name = $1
    AND namespace = $2
    AND cluster = $3
LIMIT
    1;
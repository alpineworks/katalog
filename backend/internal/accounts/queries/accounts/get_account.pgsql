SELECT
    *
FROM
    katalog.accounts
WHERE
    email_provider_hash = $1
LIMIT
    1;
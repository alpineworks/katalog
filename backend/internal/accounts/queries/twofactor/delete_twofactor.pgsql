DELETE FROM katalog.twofactor
WHERE
    parent_account_hash = $1
LIMIT
    1;
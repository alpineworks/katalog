INSERT INTO
    katalog.deployments (
        id,
        name,
        namespace,
        cluster,
        replicas,
        true_replicas,
        labels
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (id) -- Specify the unique constraint (in this case, the primary key 'id')
DO
UPDATE
SET
    name = EXCLUDED.name,
    namespace = EXCLUDED.namespace,
cluster = EXCLUDED.cluster,
replicas = EXCLUDED.replicas,
true_replicas = EXCLUDED.true_replicas,
labels = EXCLUDED.labels;
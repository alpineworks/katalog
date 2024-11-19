CREATE TABLE IF NOT EXISTS
  katalog.deployments (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid (),
    name TEXT NOT NULL,
    namespace TEXT NOT NULL,
    cluster TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    replicas INT NOT NULL,
    true_replicas INT NOT NULL,
    labels JSONB NOT NULL
  );
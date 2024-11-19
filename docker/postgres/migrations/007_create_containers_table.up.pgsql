CREATE TABLE IF NOT EXISTS
    katalog.containers (
        id UUID PRIMARY KEY NOT NULL,
        deployment UUID NOT NULL REFERENCES katalog.deployments (id) ON DELETE CASCADE,
        name TEXT NOT NULL,
        image TEXT NOT NULL,
        tag TEXT NOT NULL
    );
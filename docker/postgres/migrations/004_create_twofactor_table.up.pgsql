CREATE TABLE IF NOT EXISTS
    katalog.twofactor (
        id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid (),
        parent_account_hash TEXT NOT NULL REFERENCES katalog.accounts (email_provider_hash) ON DELETE CASCADE,
        created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
        secret TEXT NOT NULL,
        enabled BOOLEAN DEFAULT FALSE
    );
CREATE INDEX idx_outbox_unpublished
ON outbox (published_at)
WHERE published_at is NULL;

ALTER TABLE outbox
DROP COLUMN type;

ALTER TABLE outbox
ADD COLUMN type TEXT NOT NULL;
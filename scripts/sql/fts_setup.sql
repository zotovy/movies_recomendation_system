-- Adding document field
ALTER TABLE movies
    ADD COLUMN IF NOT EXISTS document tsvector;

-- Compute tsvector for each movie
-- noinspection SqlWithoutWhere
UPDATE movies
SET document = to_tsvector(movies.title);

-- Creating index for tsvector document
CREATE INDEX IF NOT EXISTS movies_document_idx
    ON movies
        USING gin (document);

-- Trigger for recomputing row tsvector after each insert or update
CREATE FUNCTION movies_tsvector_trigger() RETURNS TRIGGER AS
$$
BEGIN
    new.document = to_tsvector(new.title);
END

$$ LANGUAGE 'plpgsql';

CREATE TRIGGER movies_document_trigger
    BEFORE INSERT OR UPDATE
    ON movies
    FOR EACH ROW
EXECUTE PROCEDURE movies_tsvector_trigger();

-- pg trgm
CREATE EXTENSION pg_trgm;
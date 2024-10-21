CREATE TABLE topics (
    id INT AUTO_INCREMENT PRIMARY KEY,                   -- auto-incrementing integer
    name VARCHAR(255) UNIQUE NOT NULL,                    -- unique name, maximum 255 characters, cannot be NULL
    createdAt BIGINT NOT NULL DEFAULT (UNIX_TIMESTAMP()),  -- current timestamp in seconds as BIGINT
    updatedAt BIGINT NOT NULL DEFAULT (UNIX_TIMESTAMP()),  -- current timestamp in seconds as BIGINT
    INDEX topic_idx_name (name)                                 -- index on Name for optimization
);

CREATE TABLE messages (
    id VARCHAR(255) PRIMARY KEY,                            -- unique identifier for each message
    message TEXT,                                           -- nullable message field
    topicId INT NOT NULL,                                   -- foreign key referencing topics(Id)
    createdAt BIGINT NOT NULL DEFAULT (UNIX_TIMESTAMP()),   -- timestamp, default current time in seconds
    updatedAt BIGINT NOT NULL DEFAULT (UNIX_TIMESTAMP()),   -- timestamp, default current time in seconds
    FOREIGN KEY (topicId) REFERENCES topics(id)             -- foreign key to topics table
        ON DELETE CASCADE ON UPDATE CASCADE,                -- cascade deletes and updates
    INDEX message_idx_created_at (createdAt)                        -- index on CreatedAt for optimized sorting
);
CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    descryption VARCHAR(100) NOT NULL,
    parent_task_id INT REFERENCES tasks(task_id) ON DELETE CASCADE
);

CREATE TABLE files (
    file_id SERIAL PRIMARY KEY,
    file_name VARCHAR(255) NOT NULL,
    task_id INT REFERENCES tasks(task_id) ON DELETE CASCADE
);

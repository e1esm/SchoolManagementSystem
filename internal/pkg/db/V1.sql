CREATE TABLE IF NOT EXISTS users (
    id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name int,
    role int
);


CREATE TABLE IF NOT EXISTS grades_students(
    student_id int,
    grade int,
        CONSTRAINT fk_users
            FOREIGN KEY(student_id)
                REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS teachers_subjects(
    teacher_id int,
        subject text,
            CONSTRAINT fk_users
                FOREIGN KEY (teacher_id)
                    REFERENCES users(id)

);
CREATE PROCEDURE GetUserByUuid(IN userUuid VARCHAR(255))
BEGIN
    SELECT * FROM users WHERE uuid = userUuid;
END;

CREATE PROCEDURE GetUserByEmail(IN userEmail VARCHAR(255))
BEGIN
    SELECT * FROM users WHERE email = userEmail;
END;

CREATE PROCEDURE CreateUser(
    IN email VARCHAR(255),
    IN pass VARCHAR(255),
    IN username VARCHAR(255),
    IN first_name VARCHAR(255),
    IN last_name VARCHAR(255),
    IN uuid VARchar(255)
)
BEGIN
    INSERT INTO users (email, password, username, first_name, last_name, uuid)
    VALUES (email, pass, username, first_name, last_name, uuid);
END;

CREATE PROCEDURE GetDepartmentByUuid(IN departmentUuid VARCHAR(255))
BEGIN
    SELECT * FROM departments WHERE uuid = departmentUuid;
END;

CREATE PROCEDURE GetDepartmentUsers(IN departmentUuid VARCHAR(255))
BEGIN
    SELECT u.* from departments d
    JOIN users_in_departments uid on uid.department_id = d.id
    JOIN users u on u.id = uid.user_id
    WHERE d.uuid = departmentUuid;
END;

CREATE PROCEDURE GetActiveDepartments()
BEGIN
    SELECT * from departments where flags & 4;
END;

CREATE PROCEDURE GetApprovedDepartments()
BEGIN
    SELECT * from departments where flags & 1;
END;

CREATE PROCEDURE GetExistingDepartments()
BEGIN
    SELECT * from departments where flags & ~2;
END;

CREATE PROCEDURE GetAllDepartments()
BEGIN
    SELECT * from departments;
END;

CREATE PROCEDURE AddUserInDepartment(IN userUuid VARCHAR(255), IN departmentUuid VARCHAR(255))
BEGIN
    INSERT INTO users_in_departments (user_id, department_id) VALUES (
	(SELECT id FROM users WHERE uuid = userUuid LIMIT 1),
	(SELECT id FROM departments WHERE uuid = departmentUuid LIMIT 1)
);
END;

CREATE PROCEDURE DeleteUserFromDepartment(IN userUuid VARCHAR(255), IN departmentUuid VARCHAR(255))
BEGIN
    DELETE FROM users_in_departments 
    WHERE 
    user_id = (SELECT id FROM users WHERE uuid = userUuid LIMIT 1) AND
    department_id = (SELECT id FROM departments WHERE uuid = departmentUuid LIMIT 1);
END;


CREATE PROCEDURE CreateDepartment(IN name VARCHAR(255), IN flags TINYINT UNSIGNED, IN description TEXT, IN uuid VARCHAR(255))
BEGIN
    INSERT INTO departments (name, uuid, description, flags) VALUES
    (name, uuid, description, flags);
END;

CREATE PROCEDURE UpdateDepartmentName(IN name VARCHAR(255), IN uuid VARCHAR(255))
BEGIN
    UPDATE departments d SET d.name=name
    WHERE d.uuid=uuid;
END;

CREATE PROCEDURE UpdateDepartmentDescription(IN description TEXT, IN uuid VARCHAR(255))
BEGIN
    UPDATE departments d SET d.description=description
    WHERE d.uuid=uuid;
END;

CREATE PROCEDURE UpdateDepartmentflags(IN flags TINYINT UNSIGNED, IN uuid VARCHAR(255))
BEGIN
    UPDATE departments d SET d.flags=flags
    WHERE d.uuid=uuid;
END;

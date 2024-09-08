-- name: CreateAppointment :one
INSERT INTO appointments(users_id,doctor_id,name,phone,gender,age,date,time,address)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
RETURNING *;

-- name: GetDoctorAppointment :many
SELECT * FROM appointments
WHERE doctor_id = $1 AND  Created_at<now()-INTERVAL '1 day';

-- name: GetAllAppointment :many
SELECT * FROM appointments
Order by Created_at desc;

-- name: MyAppointment :many
SELECT * FROM appointments
WHERE users_id = $1;

-- name: DeleteAppointment :exec
DELETE FROM appointments
WHERE id = $1 and users_id=$2;

-- name: SeeAllAppointmentByDoctorId :many
SELECT * FROM appointments
WHERE doctor_id=$1;

-- name: SeeAllAppointmentByUserId :many
SELECT * FROM appointments
WHERE users_id=$1;

-- name: GetAppointmentDetails :one
SELECT a.name,a.phone,a.age,d.name as doctorName,d.shedule,u.name AS username
FROM appointments AS a
INNER JOIN doctors AS d on d.id=a.doctor_id
INNER JOIN users AS u on u.id=a.users_id
WHERE a.id =$1;

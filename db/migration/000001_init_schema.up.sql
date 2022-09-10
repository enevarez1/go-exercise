CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    user_name TEXT NOT NULL,
    full_name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    last_updated TIMESTAMP
);

CREATE TABLE Exercise (
  	id SERIAL PRIMARY KEY,
  	exercise_name TEXT NOT NULL,
  	exercise_type_id TEXT NOT NULL,
  	exercise_equipment_id TEXT NOT NULL,
  	user_id INT NOT NULL,
	CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES Users(id)
);

CREATE TABLE Workout_Type (
  	workout_type TEXT PRIMARY KEY
);

CREATE TABLE Workout (
  	id SERIAL PRIMARY KEY,
  	workout_name TEXT NOT NULL,
 	workout_type_id TEXT NOT NULL,
	CONSTRAINT fk_workout_type FOREIGN KEY(workout_type_id) REFERENCES Workout_Type(workout_type)
);

CREATE TABLE Exercise_Workout_Junction (
  	id SERIAL PRIMARY KEY,
  	exercise_id INT NOT NULL,
  	workout_id INT NOT NULL,
	CONSTRAINT fk_exercise_id FOREIGN KEY(exercise_id) REFERENCES Exercise("id"),
	CONSTRAINT fk_workout_id FOREIGN KEY(workout_id) REFERENCES Workout("id")
);

CREATE TABLE Exercise_Workout_Target (
  	id SERIAL PRIMARY KEY,
  	junction_id INT NOT NULL,
  	set_number INT NOT NULL,
  	min_reps INT NOT NULL,
  	max_reps INT NOT NULL,
	CONSTRAINT fk_junction_id FOREIGN KEY(junction_id) REFERENCES Exercise_Workout_Junction("id")
);

CREATE TABLE Log (
  	id SERIAL PRIMARY KEY,
  	workout_id INT NOT NULL,
  	dateOf TIMESTAMP NOT NULL,
	CONSTRAINT fk_workout_id FOREIGN KEY(workout_id) REFERENCES Workout("id")
);

CREATE TABLE Log_Entry (
  	id SERIAL PRIMARY KEY,
  	log_id INT NOT NULL,
  	junction_id INT NOT NULL,
  	set_number INT NOT NULL,
  	weight INT NOT NULL,
  	reps INT NOT NULL,
  	time_recorded TIMESTAMP,
	CONSTRAINT fk_junction_id FOREIGN KEY(junction_id) REFERENCES Exercise_Workout_Junction("id"),
	CONSTRAINT fk_log_id FOREIGN KEY(log_id) REFERENCES Log(id)
);
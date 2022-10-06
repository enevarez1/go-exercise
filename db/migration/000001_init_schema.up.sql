CREATE TABLE Users (
    user_name TEXT PRIMARY KEY,
    full_name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_updated TIMESTAMP
);

CREATE TABLE Exercise (
  	id SERIAL PRIMARY KEY,
  	exercise_name TEXT NOT NULL,
  	exercise_type_id TEXT NOT NULL,
  	user_id TEXT NOT NULL,
	CONSTRAINT fk_user 
		FOREIGN KEY(user_id) 
		REFERENCES Users(user_name)
		ON DELETE CASCADE
);

CREATE TABLE Workout (
  	id SERIAL PRIMARY KEY,
  	workout_name TEXT NOT NULL,
 	workout_type TEXT [],
	user_id TEXT NOT NULL,
	CONSTRAINT fk_user 
		FOREIGN KEY(user_id) 
		REFERENCES Users(user_name)
		ON DELETE CASCADE
);

CREATE TABLE Exercise_Workout_Junction (
	junction_id TEXT NOT NULL,
  	exercise_id INT NOT NULL,
  	workout_id INT NOT NULL,
	user_id TEXT NOT NULL,
	CONSTRAINT fk_user_id 
		FOREIGN KEY(user_id) 
		REFERENCES Users(user_name)
		ON DELETE CASCADE,
	CONSTRAINT fk_exercise_id 
		FOREIGN KEY(exercise_id) 
		REFERENCES Exercise(id)
		ON DELETE CASCADE,
	CONSTRAINT fk_workout_id
		FOREIGN KEY(workout_id)
		REFERENCES Workout(id)
		ON DELETE CASCADE
);

CREATE UNIQUE INDEX idx_junction ON Exercise_Workout_Junction (exercise_id, workout_id, user_id);

CREATE TABLE Exercise_Workout_Target (
  	id SERIAL PRIMARY KEY,
  	junction_id TEXT NOT NULL,
  	set_number INT NOT NULL,
  	min_reps INT NOT NULL,
  	max_reps INT NOT NULL,
	CONSTRAINT fk_junction_id 
		FOREIGN KEY(junction_id) 
		REFERENCES Exercise_Workout_Junction(junction_id)
		ON DELETE CASCADE
);

CREATE TABLE Log (
  	id SERIAL PRIMARY KEY,
  	workout_id INT NOT NULL,
  	dateOf TIMESTAMP NOT NULL,
	user_id TEXT NOT NULL,
	CONSTRAINT fk_user 
		FOREIGN KEY(user_id) 
		REFERENCES Users(id)
		ON DELETE CASCADE,
	CONSTRAINT fk_workout_id 
		FOREIGN KEY(workout_id) 
		REFERENCES Workout("id")
		ON DELETE CASCADE
);

CREATE TABLE Log_Entry (
  	id SERIAL PRIMARY KEY,
  	log_id INT NOT NULL,
  	junction_id TEXT NOT NULL,
  	set_number INT NOT NULL,
  	weight INT NOT NULL,
  	reps INT NOT NULL,
  	time_recorded TIMESTAMP,
	CONSTRAINT fk_junction_id FOREIGN KEY(junction_id) REFERENCES Exercise_Workout_Junction(junction_id),
	CONSTRAINT fk_log_id FOREIGN KEY(log_id) REFERENCES Log(id)
);
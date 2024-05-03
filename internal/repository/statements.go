package repo

const (
	insertPoint = `
		INSERT INTO param_tracker.point (name, alias)
		VALUES ($1, $2);
	`

	getAllPoint = `
		SELECT point_id,
		       name,
			   alias
		FROM param_tracker.point
		ORDER BY point_id DESC; 
	`

	getByIDPoint = `
		SELECT point_id,
		       name,
			   alias
		FROM param_tracker.point
		WHERE point_id=$1; 
	`

	insertGps = `
	INSERT INTO param_tracker.gps (point_id, point_gps_id, lat, lon, speed, time)
	VALUES ($1, $2, $3, $4, $5, $6);
	`

	getGpsByPointID = `
		SELECT gps_id, 
			   point_id, 
			   point_gps_id, 
			   lat, 
			   lon, 
			   speed, 
			   time
		FROM param_tracker.gps 
		WHERE point_id=$1;
	`

	insertParamStr = `
	INSERT INTO param_tracker.param_str (string_point_id, point_id, channel, value, time)
	VALUES ($1, $2, $3, $4, $5)
`

	getParamStrByPointID = `
		SELECT string_id, 
	   		string_point_id, 
	  		point_id, 
	   		channel, 
	   		value, 
	   		time
		FROM param_tracker.param_str 
		WHERE point_id=$1;
	`

	insertParamInt = `
	INSERT INTO param_tracker.param_int (int_point_id, point_id, channel, value, time)
	VALUES ($1, $2, $3, $4, $5)
`

	getParamIntByPointID = `
		SELECT int_id, 
			int_point_id, 
	  		point_id, 
	   		channel, 
	   		value, 
	   		time
		FROM param_tracker.param_int 
		WHERE point_id=$1;
	`
)

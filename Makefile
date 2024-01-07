run-migrate: 
	migrate -source file://backend/migrations -database postgres://localhost:5432/database up 2
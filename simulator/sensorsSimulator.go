package simulator

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/francisko-rezende/sensor-simulator/schemas"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func SensorsSimulator() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	query := `
	SELECT
		s."sensorId", s."sensorName", s."createdAt", s."updatedAt",
		s."macAddress", s."status", s."locationId", s."sensorTypeId",
		st."sensorTypeId", st."sensorTypeName", st."sensorType",
		st."minRange", st."maxRange", st."barcode", st."batch"
	FROM
		"Sensor" AS s
	INNER JOIN
		"SensorType" AS st ON s."sensorTypeId" = st."sensorTypeId"
	WHERE
		s."status"=true
`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Fatalf("Query error: %v\n", err)
	}
	defer rows.Close()

	var sensors []schemas.Sensor

	for rows.Next() {
		var sensor schemas.Sensor
		var sensorType schemas.SensorType

		err := rows.Scan(
			&sensor.SensorID, &sensor.SensorName, &sensor.CreatedAt, &sensor.UpdatedAt,
			&sensor.MacAddress, &sensor.Status, &sensor.LocationID, &sensor.SensorTypeID,
			&sensorType.SensorTypeID, &sensorType.SensorTypeName, &sensorType.SensorType,
			&sensorType.MinRange, &sensorType.MaxRange, &sensorType.Barcode, &sensorType.Batch,
		)
		if err != nil {
			log.Fatalf("Scan error: %v\n", err)
		}

		sensor.SensorType = sensorType
		sensors = append(sensors, sensor)
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("Rows error: %v\n", err)
	}

	for _, sensor := range sensors {

		value := getRandomValue(sensor.SensorType.MinRange, sensor.SensorType.MaxRange)

		newMeasurement := schemas.Measurement{
			Value:        value,
			SensorID:     sensor.SensorID,
			CreatedAt:    time.Now(),
			CreationDate: time.Now(),
			UpdatedAt:    time.Now(),
		}

		insertStatement := `
			INSERT INTO "Measurement" ("value", "createdAt", "creationDate", "updatedAt", "sensorId")
			VALUES ($1, $2, $3, $4, $5)
		`
		_, err = conn.Exec(
			context.Background(),
			insertStatement,
			newMeasurement.Value,
			newMeasurement.CreatedAt,
			newMeasurement.CreationDate,
			newMeasurement.UpdatedAt,
			newMeasurement.SensorID,
		)
		if err != nil {
			log.Fatalf("Insert error: %v\n", err)
		}
	}

}

func getRandomValue(min, max int) int {

	return rand.Intn(max-min+1) + min
}

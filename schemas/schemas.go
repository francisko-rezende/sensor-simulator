package schemas

import "time"

type Measurement struct {
	MeasurementID int
	Value         int
	CreatedAt     time.Time
	CreationDate  time.Time
	UpdatedAt     time.Time
	SensorID      int
}

type Sensor struct {
	SensorID     int
	SensorName   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	MacAddress   string
	Status       bool
	LocationID   int
	SensorTypeID int
	SensorType   SensorType
}

type SensorType struct {
	SensorTypeID   int
	SensorTypeName string
	SensorType     string
	MinRange       int
	MaxRange       int
	Barcode        string
	Batch          string
}

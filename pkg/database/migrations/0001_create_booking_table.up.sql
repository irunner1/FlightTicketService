CREATE TABLE IF NOT EXISTS booking_flights (
    id SERIAL PRIMARY KEY,
    flight_id VARCHAR(10) NOT NULL,
    passenger_id VARCHAR(10) NOT NULL,
    booking_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    departure_time TIMESTAMP NOT NULL,
    arrival_time TIMESTAMP NOT NULL,
    status VARCHAR(30) NOT NULL CHECK (status IN ("booked", "cancelled", "confirmed")),
    seat_number VARCHAR(30),
    additional_info VARCHAR(100),
    FOREIGN KEY (flight_id) REFERENCES flights(id),
    FOREIGN KEY (passenger_id) REFERENCES passengers(id)
);
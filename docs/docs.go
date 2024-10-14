// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/flights": {
            "get": {
                "description": "get flights",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flights"
                ],
                "summary": "Get list of flights",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/flights.Flight"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/flights/search": {
            "get": {
                "description": "Retrieves a list of flights filtered by the provided search parameters.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flights"
                ],
                "summary": "Search flights by parameters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Origin location of the flight",
                        "name": "origin",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Destination location of the flight",
                        "name": "destination",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Departure date and time",
                        "name": "departure",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Arrival date and time",
                        "name": "arrival",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/flights.Flight"
                            }
                        }
                    },
                    "404": {
                        "description": "No flights found matching the search criteria"
                    }
                }
            }
        },
        "/api/v1/flights/{id}": {
            "get": {
                "description": "Retrieves flight details for a specific flight ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flights"
                ],
                "summary": "Get flight by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Unique identifier of the flight",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/flights.Flight"
                        }
                    },
                    "404": {
                        "description": "Flight not found"
                    }
                }
            }
        },
        "/api/v1/tickets": {
            "get": {
                "description": "get tickets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tickets"
                ],
                "summary": "Get list of tickets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/booking.Ticket"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/tickets/book": {
            "post": {
                "description": "Creates a new ticket booking for a flight.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "booking"
                ],
                "summary": "Book a new ticket",
                "parameters": [
                    {
                        "description": "Ticket data",
                        "name": "ticket",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/booking.Ticket"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ticket successfully booked"
                    },
                    "400": {
                        "description": "Invalid ticket data"
                    }
                }
            }
        },
        "/api/v1/tickets/{id}": {
            "get": {
                "description": "Returns ticket details for a specific ticket ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tickets"
                ],
                "summary": "Get ticket by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Unique identifier of the ticket",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/booking.Ticket"
                        }
                    },
                    "404": {
                        "description": "ticket not found"
                    }
                }
            }
        },
        "/api/v1/tickets/{ticketID}/cancel": {
            "post": {
                "description": "Cancels an existing ticket using the ticket ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "booking"
                ],
                "summary": "Cancel a ticket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The ID of the ticket to cancel",
                        "name": "ticketID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ticket successfully cancelled"
                    },
                    "404": {
                        "description": "Ticket not found"
                    }
                }
            }
        },
        "/api/v1/{ticketID}/change": {
            "post": {
                "description": "Changes the flight associated with a ticket to a new flight using the ticket ID and new flight ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "booking"
                ],
                "summary": "Change the flight of a ticket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The ID of the ticket to update",
                        "name": "ticketID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The new flight ID to associate with the ticket",
                        "name": "newFlightID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Flight successfully changed for the ticket"
                    },
                    "400": {
                        "description": "Invalid parameters"
                    },
                    "404": {
                        "description": "Ticket not found or cannot change flight for a cancelled ticket"
                    }
                }
            }
        }
    },
    "definitions": {
        "booking.Ticket": {
            "type": "object",
            "properties": {
                "additional_info": {
                    "type": "string"
                },
                "arrival_time": {
                    "type": "string"
                },
                "booking_time": {
                    "type": "string"
                },
                "departure_time": {
                    "type": "string"
                },
                "flight_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "passenger_id": {
                    "type": "string"
                },
                "seat_number": {
                    "type": "string"
                },
                "status": {
                    "description": "\"booked\", \"cancelled\", \"confirmed\"",
                    "type": "string"
                }
            }
        },
        "flights.Flight": {
            "description": "Flight model for API response.",
            "type": "object",
            "properties": {
                "airline": {
                    "type": "string"
                },
                "arrival": {
                    "type": "string"
                },
                "departure": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "origin": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8010",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Flight Ticket Service",
	Description:      "API Server for booking flight tickets",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
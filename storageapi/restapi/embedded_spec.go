// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Storage",
    "version": "0.0.1"
  },
  "basePath": "/",
  "paths": {
    "/load": {
      "get": {
        "operationId": "load",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "stationID",
                "key"
              ],
              "properties": {
                "key": {
                  "type": "string",
                  "minLength": 1
                },
                "stationID": {
                  "$ref": "#/definitions/StationID"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "not found"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/ping": {
      "get": {
        "operationId": "ping",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/save": {
      "post": {
        "operationId": "save",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "stationID",
                "keyPair"
              ],
              "properties": {
                "keyPair": {
                  "$ref": "#/definitions/KeyPair"
                },
                "stationID": {
                  "$ref": "#/definitions/StationID"
                }
              }
            }
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "KeyPair": {
      "type": "object",
      "required": [
        "key",
        "value"
      ],
      "properties": {
        "key": {
          "type": "string",
          "minLength": 1
        },
        "value": {
          "type": "string"
        }
      }
    },
    "StationID": {
      "type": "string",
      "minLength": 1
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Storage",
    "version": "0.0.1"
  },
  "basePath": "/",
  "paths": {
    "/load": {
      "get": {
        "operationId": "load",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "stationID",
                "key"
              ],
              "properties": {
                "key": {
                  "type": "string",
                  "minLength": 1
                },
                "stationID": {
                  "$ref": "#/definitions/StationID"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "not found"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/ping": {
      "get": {
        "operationId": "ping",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/save": {
      "post": {
        "operationId": "save",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "stationID",
                "keyPair"
              ],
              "properties": {
                "keyPair": {
                  "$ref": "#/definitions/KeyPair"
                },
                "stationID": {
                  "$ref": "#/definitions/StationID"
                }
              }
            }
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "KeyPair": {
      "type": "object",
      "required": [
        "key",
        "value"
      ],
      "properties": {
        "key": {
          "type": "string",
          "minLength": 1
        },
        "value": {
          "type": "string"
        }
      }
    },
    "StationID": {
      "type": "string",
      "minLength": 1
    }
  }
}`))
}

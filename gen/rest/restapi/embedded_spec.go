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
  "swagger": "2.0",
  "info": {
    "description": "An API that allows users to obtain existing information of controllers and indications",
    "title": "Controller-backend REST API",
    "version": "1.0.0"
  },
  "paths": {
    "/controller": {
      "post": {
        "description": "Add controller in DB",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Add controller",
        "parameters": [
          {
            "description": "Controller's serial",
            "name": "controller",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "serial": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Status",
            "schema": {
              "type": "object",
              "properties": {
                "success": {
                  "type": "boolean"
                }
              }
            }
          },
          "400": {
            "description": "Controller already exists",
            "schema": {
              "$ref": "#/responses/AlreadyExists"
            }
          }
        }
      },
      "delete": {
        "description": "Delete controller with all it's indications",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Delete controller",
        "parameters": [
          {
            "description": "Controller's serial",
            "name": "controller",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "serial": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Status",
            "schema": {
              "type": "object",
              "properties": {
                "success": {
                  "type": "boolean"
                }
              }
            }
          },
          "400": {
            "description": "Controller not found",
            "schema": {
              "$ref": "#/responses/NotFound"
            }
          }
        }
      }
    },
    "/controller/indications": {
      "get": {
        "description": "Get list of controller's indications",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "List of controller's indications",
        "parameters": [
          {
            "description": "Controller's serial",
            "name": "controller",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "serial": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Controller's indications",
            "schema": {
              "type": "object",
              "properties": {
                "indications": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "indication": {
                        "type": "number",
                        "format": "float"
                      },
                      "sent_at": {
                        "type": "string",
                        "format": "date-time"
                      }
                    }
                  }
                }
              }
            }
          },
          "400": {
            "description": "Controller not found",
            "schema": {
              "$ref": "#/responses/NotFound"
            }
          }
        }
      }
    },
    "/controllers": {
      "get": {
        "description": "Get list of controllers",
        "produces": [
          "application/json"
        ],
        "summary": "List of controllers",
        "responses": {
          "200": {
            "description": "Successful pull of controllers info",
            "schema": {
              "type": "object",
              "properties": {
                "controllers": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "created_at": {
                        "type": "string",
                        "format": "date-time"
                      },
                      "serial": {
                        "type": "string"
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "AlreadyExists": {
      "description": "Controller already exists",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NotFound": {
      "description": "Controller not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "An API that allows users to obtain existing information of controllers and indications",
    "title": "Controller-backend REST API",
    "version": "1.0.0"
  },
  "paths": {
    "/controller": {
      "post": {
        "description": "Add controller in DB",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Add controller",
        "parameters": [
          {
            "description": "Controller's serial",
            "name": "controller",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "serial": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Status",
            "schema": {
              "type": "object",
              "properties": {
                "success": {
                  "type": "boolean"
                }
              }
            }
          },
          "400": {
            "description": "Controller already exists",
            "schema": {
              "description": "Controller already exists",
              "schema": {
                "$ref": "#/definitions/Error"
              }
            }
          }
        }
      },
      "delete": {
        "description": "Delete controller with all it's indications",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Delete controller",
        "parameters": [
          {
            "description": "Controller's serial",
            "name": "controller",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "serial": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Status",
            "schema": {
              "type": "object",
              "properties": {
                "success": {
                  "type": "boolean"
                }
              }
            }
          },
          "400": {
            "description": "Controller not found",
            "schema": {
              "description": "Controller not found",
              "schema": {
                "$ref": "#/definitions/Error"
              }
            }
          }
        }
      }
    },
    "/controller/indications": {
      "get": {
        "description": "Get list of controller's indications",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "List of controller's indications",
        "parameters": [
          {
            "description": "Controller's serial",
            "name": "controller",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "serial": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Controller's indications",
            "schema": {
              "type": "object",
              "properties": {
                "indications": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/IndicationsItems0"
                  }
                }
              }
            }
          },
          "400": {
            "description": "Controller not found",
            "schema": {
              "description": "Controller not found",
              "schema": {
                "$ref": "#/definitions/Error"
              }
            }
          }
        }
      }
    },
    "/controllers": {
      "get": {
        "description": "Get list of controllers",
        "produces": [
          "application/json"
        ],
        "summary": "List of controllers",
        "responses": {
          "200": {
            "description": "Successful pull of controllers info",
            "schema": {
              "type": "object",
              "properties": {
                "controllers": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/ControllersItems0"
                  }
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ControllersItems0": {
      "type": "object",
      "properties": {
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "serial": {
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "string"
        }
      }
    },
    "IndicationsItems0": {
      "type": "object",
      "properties": {
        "indication": {
          "type": "number",
          "format": "float"
        },
        "sent_at": {
          "type": "string",
          "format": "date-time"
        }
      }
    }
  },
  "responses": {
    "AlreadyExists": {
      "description": "Controller already exists",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NotFound": {
      "description": "Controller not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  }
}`))
}

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
  "swagger": "2.0",
  "info": {
    "description": "Kruise deployment wizard",
    "title": "Kruise deployment wizard",
    "version": "0.0.1"
  },
  "paths": {
    "/app/preview": {
      "post": {
        "description": "Previews a new Kruise application",
        "produces": [
          "text/plain"
        ],
        "tags": [
          "apps"
        ],
        "operationId": "previewApp",
        "parameters": [
          {
            "description": "The application to preview",
            "name": "application",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/application"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "default": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    },
    "/app/release": {
      "post": {
        "description": "Generates a new Kruise application",
        "tags": [
          "apps"
        ],
        "operationId": "releaseApp",
        "parameters": [
          {
            "description": "The application to generate",
            "name": "application",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/application"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "$ref": "#/definitions/validationResponse"
            }
          },
          "400": {
            "description": "invalid",
            "schema": {
              "$ref": "#/definitions/validationResponse"
            }
          },
          "default": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    },
    "/app/validation": {
      "post": {
        "description": "Validates the Kruise application",
        "tags": [
          "validations"
        ],
        "operationId": "validateApplication",
        "parameters": [
          {
            "description": "the application object",
            "name": "application",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "validated",
            "schema": {
              "$ref": "#/definitions/validationResponse"
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "default": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    },
    "/health": {
      "get": {
        "description": "Get the current health of the API",
        "tags": [
          "general"
        ],
        "operationId": "getHealth",
        "responses": {
          "200": {
            "description": "Get current health response",
            "schema": {
              "$ref": "#/definitions/healthStatus"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "application": {
      "type": "object",
      "required": [
        "metadata",
        "spec"
      ],
      "properties": {
        "metadata": {
          "$ref": "#/definitions/metadata"
        },
        "spec": {
          "$ref": "#/definitions/spec"
        }
      }
    },
    "component": {
      "type": "object",
      "required": [
        "service",
        "containers"
      ],
      "properties": {
        "containers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/container"
          }
        },
        "ingresses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ingress"
          }
        },
        "service": {
          "$ref": "#/definitions/service"
        }
      }
    },
    "configMap": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "data": {
          "description": "The content of the ConfigMap data key",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "name": {
          "description": "The name of the ConfigMap",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "container": {
      "type": "object",
      "required": [
        "name",
        "image",
        "imageTag",
        "imagePullPolicy",
        "portNames"
      ],
      "properties": {
        "command": {
          "description": "The command to run for the docker image's entrypoint.",
          "type": "string",
          "x-nullable": true
        },
        "image": {
          "description": "The docker image name for the container",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "imagePullPolicy": {
          "description": "Image pull policy. One of Always or IfNotPresent.",
          "type": "string",
          "default": "IfNotPresent",
          "minLength": 1,
          "enum": [
            "Always",
            "IfNotPresent"
          ],
          "x-nullable": false
        },
        "imageTag": {
          "description": "The docker image tag for the container",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "name": {
          "description": "The name of this container within the service",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "portNames": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "volumes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/volumeMount"
          }
        }
      }
    },
    "destination": {
      "type": "object",
      "required": [
        "url"
      ],
      "properties": {
        "path": {
          "description": "The relative path to the manifests in the git repo",
          "type": "string",
          "format": "filepath",
          "default": "/",
          "minLength": 1,
          "x-nullable": false
        },
        "targetRevision": {
          "description": "Defines the commit, tag, or branch in which to sync the application to.",
          "type": "string",
          "default": "HEAD",
          "minLength": 1,
          "x-nullable": false
        },
        "url": {
          "description": "The git repo URL",
          "type": "string",
          "format": "uri",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "healthStatus": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    },
    "ingress": {
      "type": "object",
      "required": [
        "host",
        "paths"
      ],
      "properties": {
        "host": {
          "description": "The hostname for the ingress",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "paths": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ingressPath"
          }
        }
      }
    },
    "ingressPath": {
      "type": "object",
      "required": [
        "path",
        "portName"
      ],
      "properties": {
        "path": {
          "description": "Path is matched against the path of an incoming request",
          "type": "string",
          "default": "/",
          "minLength": 1,
          "x-nullable": false
        },
        "portName": {
          "description": "Specifies the port name of the service to expose",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "labels": {
      "type": "object",
      "required": [
        "version",
        "team",
        "env",
        "region"
      ],
      "properties": {
        "env": {
          "description": "The environment to deploy to",
          "type": "string",
          "default": "Dev",
          "minLength": 1,
          "enum": [
            "Dev",
            "Stage",
            "Prod"
          ],
          "x-nullable": false
        },
        "region": {
          "description": "The region to deploy to",
          "type": "string",
          "default": "STL",
          "minLength": 1,
          "enum": [
            "STL",
            "KCI",
            "BEL"
          ],
          "x-nullable": false
        },
        "team": {
          "description": "The name of the team or tenant",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "version": {
          "description": "The version or release name of the application",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "metadata": {
      "type": "object",
      "required": [
        "name",
        "namespace",
        "labels"
      ],
      "properties": {
        "labels": {
          "description": "Arbitrary labels",
          "$ref": "#/definitions/labels"
        },
        "name": {
          "description": "The name of the application",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "namespace": {
          "description": "The namespace to deploy to",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "persistentVolume": {
      "type": "object",
      "required": [
        "name",
        "accessMode",
        "capacity",
        "storageClassName"
      ],
      "properties": {
        "accessMode": {
          "description": "The desired access mode for the volume",
          "type": "string",
          "default": "ReadWriteOnce",
          "minLength": 1,
          "enum": [
            "ReadWriteOnce",
            "ReadOnlyMany",
            "ReadWriteMany"
          ],
          "x-nullable": false
        },
        "capacity": {
          "description": "The desired size of the volume in GB",
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "name": {
          "description": "The name of the volume",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "storageClassName": {
          "description": "The desired storage class for the volume",
          "type": "string",
          "default": "default",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "service": {
      "type": "object",
      "required": [
        "name",
        "type",
        "ports"
      ],
      "properties": {
        "name": {
          "description": "The name of the service",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "ports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/servicePort"
          }
        },
        "type": {
          "description": "The service type",
          "type": "string",
          "enum": [
            "ClusterIP",
            "ExternalName",
            "LoadBalancer"
          ],
          "x-nullable": false
        }
      }
    },
    "servicePort": {
      "type": "object",
      "required": [
        "name",
        "port"
      ],
      "properties": {
        "name": {
          "description": "The name of this port within the service",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "port": {
          "description": "The port that will be exposed by this service",
          "type": "integer",
          "x-nullable": false
        },
        "protocol": {
          "description": "The IP protocol for this port. Supports \"TCP\" and \"UDP\". Default is TCP",
          "type": "string",
          "default": "TCP",
          "minLength": 1,
          "enum": [
            "TCP",
            "UDP"
          ],
          "x-nullable": false
        },
        "targetPort": {
          "description": "Number or name of the port to access on the pods targeted by the service",
          "type": "integer",
          "x-nullable": false
        }
      }
    },
    "spec": {
      "type": "object",
      "required": [
        "destination",
        "components"
      ],
      "properties": {
        "components": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/component"
          }
        },
        "configMaps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/configMap"
          }
        },
        "destination": {
          "$ref": "#/definitions/destination"
        },
        "persistentVolumes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/persistentVolume"
          }
        }
      }
    },
    "validationResponse": {
      "type": "object",
      "properties": {
        "errors": {
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        }
      }
    },
    "volumeMount": {
      "type": "object",
      "properties": {
        "mountPath": {
          "description": "Path within the container at which the volume should be mounted",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "name": {
          "description": "The name of the volume to mount",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "readOnly": {
          "description": "Mounted read-only if true, read-write otherwise",
          "type": "boolean",
          "default": false,
          "x-nullable": false
        },
        "subPath": {
          "description": "Path within the volume from which the container's volume should be mounted",
          "type": "string",
          "default": "",
          "x-nullable": true
        },
        "type": {
          "description": "The type of the volume mount (ConfigMap, PersistentVolume, or Secret)",
          "type": "string",
          "minLength": 1,
          "enum": [
            "ConfigMap",
            "PersistentVolume",
            "Secret"
          ],
          "x-nullable": false
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad request",
      "schema": {
        "type": "string"
      }
    },
    "InternalServerError": {
      "description": "Internal server error",
      "schema": {
        "$ref": "#/definitions/error"
      }
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
  "swagger": "2.0",
  "info": {
    "description": "Kruise deployment wizard",
    "title": "Kruise deployment wizard",
    "version": "0.0.1"
  },
  "paths": {
    "/app/preview": {
      "post": {
        "description": "Previews a new Kruise application",
        "produces": [
          "text/plain"
        ],
        "tags": [
          "apps"
        ],
        "operationId": "previewApp",
        "parameters": [
          {
            "description": "The application to preview",
            "name": "application",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/application"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/app/release": {
      "post": {
        "description": "Generates a new Kruise application",
        "tags": [
          "apps"
        ],
        "operationId": "releaseApp",
        "parameters": [
          {
            "description": "The application to generate",
            "name": "application",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/application"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "$ref": "#/definitions/validationResponse"
            }
          },
          "400": {
            "description": "invalid",
            "schema": {
              "$ref": "#/definitions/validationResponse"
            }
          },
          "default": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/app/validation": {
      "post": {
        "description": "Validates the Kruise application",
        "tags": [
          "validations"
        ],
        "operationId": "validateApplication",
        "parameters": [
          {
            "description": "the application object",
            "name": "application",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "validated",
            "schema": {
              "$ref": "#/definitions/validationResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "description": "Get the current health of the API",
        "tags": [
          "general"
        ],
        "operationId": "getHealth",
        "responses": {
          "200": {
            "description": "Get current health response",
            "schema": {
              "$ref": "#/definitions/healthStatus"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "application": {
      "type": "object",
      "required": [
        "metadata",
        "spec"
      ],
      "properties": {
        "metadata": {
          "$ref": "#/definitions/metadata"
        },
        "spec": {
          "$ref": "#/definitions/spec"
        }
      }
    },
    "component": {
      "type": "object",
      "required": [
        "service",
        "containers"
      ],
      "properties": {
        "containers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/container"
          }
        },
        "ingresses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ingress"
          }
        },
        "service": {
          "$ref": "#/definitions/service"
        }
      }
    },
    "configMap": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "data": {
          "description": "The content of the ConfigMap data key",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "name": {
          "description": "The name of the ConfigMap",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "container": {
      "type": "object",
      "required": [
        "name",
        "image",
        "imageTag",
        "imagePullPolicy",
        "portNames"
      ],
      "properties": {
        "command": {
          "description": "The command to run for the docker image's entrypoint.",
          "type": "string",
          "x-nullable": true
        },
        "image": {
          "description": "The docker image name for the container",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "imagePullPolicy": {
          "description": "Image pull policy. One of Always or IfNotPresent.",
          "type": "string",
          "default": "IfNotPresent",
          "minLength": 1,
          "enum": [
            "Always",
            "IfNotPresent"
          ],
          "x-nullable": false
        },
        "imageTag": {
          "description": "The docker image tag for the container",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "name": {
          "description": "The name of this container within the service",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "portNames": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "volumes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/volumeMount"
          }
        }
      }
    },
    "destination": {
      "type": "object",
      "required": [
        "url"
      ],
      "properties": {
        "path": {
          "description": "The relative path to the manifests in the git repo",
          "type": "string",
          "format": "filepath",
          "default": "/",
          "minLength": 1,
          "x-nullable": false
        },
        "targetRevision": {
          "description": "Defines the commit, tag, or branch in which to sync the application to.",
          "type": "string",
          "default": "HEAD",
          "minLength": 1,
          "x-nullable": false
        },
        "url": {
          "description": "The git repo URL",
          "type": "string",
          "format": "uri",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "healthStatus": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    },
    "ingress": {
      "type": "object",
      "required": [
        "host",
        "paths"
      ],
      "properties": {
        "host": {
          "description": "The hostname for the ingress",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "paths": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ingressPath"
          }
        }
      }
    },
    "ingressPath": {
      "type": "object",
      "required": [
        "path",
        "portName"
      ],
      "properties": {
        "path": {
          "description": "Path is matched against the path of an incoming request",
          "type": "string",
          "default": "/",
          "minLength": 1,
          "x-nullable": false
        },
        "portName": {
          "description": "Specifies the port name of the service to expose",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "labels": {
      "type": "object",
      "required": [
        "version",
        "team",
        "env",
        "region"
      ],
      "properties": {
        "env": {
          "description": "The environment to deploy to",
          "type": "string",
          "default": "Dev",
          "minLength": 1,
          "enum": [
            "Dev",
            "Stage",
            "Prod"
          ],
          "x-nullable": false
        },
        "region": {
          "description": "The region to deploy to",
          "type": "string",
          "default": "STL",
          "minLength": 1,
          "enum": [
            "STL",
            "KCI",
            "BEL"
          ],
          "x-nullable": false
        },
        "team": {
          "description": "The name of the team or tenant",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "version": {
          "description": "The version or release name of the application",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "metadata": {
      "type": "object",
      "required": [
        "name",
        "namespace",
        "labels"
      ],
      "properties": {
        "labels": {
          "description": "Arbitrary labels",
          "$ref": "#/definitions/labels"
        },
        "name": {
          "description": "The name of the application",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "namespace": {
          "description": "The namespace to deploy to",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "persistentVolume": {
      "type": "object",
      "required": [
        "name",
        "accessMode",
        "capacity",
        "storageClassName"
      ],
      "properties": {
        "accessMode": {
          "description": "The desired access mode for the volume",
          "type": "string",
          "default": "ReadWriteOnce",
          "minLength": 1,
          "enum": [
            "ReadWriteOnce",
            "ReadOnlyMany",
            "ReadWriteMany"
          ],
          "x-nullable": false
        },
        "capacity": {
          "description": "The desired size of the volume in GB",
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "name": {
          "description": "The name of the volume",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "storageClassName": {
          "description": "The desired storage class for the volume",
          "type": "string",
          "default": "default",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "service": {
      "type": "object",
      "required": [
        "name",
        "type",
        "ports"
      ],
      "properties": {
        "name": {
          "description": "The name of the service",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "ports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/servicePort"
          }
        },
        "type": {
          "description": "The service type",
          "type": "string",
          "enum": [
            "ClusterIP",
            "ExternalName",
            "LoadBalancer"
          ],
          "x-nullable": false
        }
      }
    },
    "servicePort": {
      "type": "object",
      "required": [
        "name",
        "port"
      ],
      "properties": {
        "name": {
          "description": "The name of this port within the service",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "port": {
          "description": "The port that will be exposed by this service",
          "type": "integer",
          "x-nullable": false
        },
        "protocol": {
          "description": "The IP protocol for this port. Supports \"TCP\" and \"UDP\". Default is TCP",
          "type": "string",
          "default": "TCP",
          "minLength": 1,
          "enum": [
            "TCP",
            "UDP"
          ],
          "x-nullable": false
        },
        "targetPort": {
          "description": "Number or name of the port to access on the pods targeted by the service",
          "type": "integer",
          "x-nullable": false
        }
      }
    },
    "spec": {
      "type": "object",
      "required": [
        "destination",
        "components"
      ],
      "properties": {
        "components": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/component"
          }
        },
        "configMaps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/configMap"
          }
        },
        "destination": {
          "$ref": "#/definitions/destination"
        },
        "persistentVolumes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/persistentVolume"
          }
        }
      }
    },
    "validationResponse": {
      "type": "object",
      "properties": {
        "errors": {
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        }
      }
    },
    "volumeMount": {
      "type": "object",
      "properties": {
        "mountPath": {
          "description": "Path within the container at which the volume should be mounted",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "name": {
          "description": "The name of the volume to mount",
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "readOnly": {
          "description": "Mounted read-only if true, read-write otherwise",
          "type": "boolean",
          "default": false,
          "x-nullable": false
        },
        "subPath": {
          "description": "Path within the volume from which the container's volume should be mounted",
          "type": "string",
          "default": "",
          "x-nullable": true
        },
        "type": {
          "description": "The type of the volume mount (ConfigMap, PersistentVolume, or Secret)",
          "type": "string",
          "minLength": 1,
          "enum": [
            "ConfigMap",
            "PersistentVolume",
            "Secret"
          ],
          "x-nullable": false
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad request",
      "schema": {
        "type": "string"
      }
    },
    "InternalServerError": {
      "description": "Internal server error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  }
}`))
}

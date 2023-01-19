package draft07pb

import (
	"bytes"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
)

func TestDecodeSchema(t *testing.T) {
	schema := new(JSONSchema)
	rdr := bytes.NewReader([]byte(testInstance))
	assert.NoError(t, jsonpb.Unmarshal(rdr, schema))
}

const testInstance = `
{
      "$id": "http://example.com/root.json",
      "$schema": "http://json-schema.org/draft-07/schema",
      "default": {
      },
      "properties": {
        "meta": {
          "properties": {
            "labels": {
              "properties": {
                "indent.properties/form/v1alpha1/field/preferred-duration/response": {
                  "title": "How long do you need access",
                  "type": "string"
                },
                "indent.properties/form/v1alpha1/field/priority/response": {
                  "title": "What's the priority?",
                  "type": "string"
                }
              },
              "required": [
                "indent.properties/form/v1alpha1/field/preferred-duration/response",
                "indent.properties/form/v1alpha1/field/priority/response"
              ],
              "title": "",
              "type": "object"
            }
          },
          "required": [
            "labels"
          ],
          "title": "Extra context",
          "type": "object"
        },
        "petitioners": {
          "$id": "#/properties/petitioners",
          "default": [
          ],
          "description": "",
          "examples": [
            [
              {
                "displayName": "Example User",
                "email": "example@indent.com"
              }
            ]
          ],
          "items": {
            "$id": "#/properties/petitioners/items",
            "allOf": [
              {
                "$id": "#/properties/petitioners/items/allOf/0",
                "default": {
                },
                "description": "",
                "examples": [
                  {
                    "displayName": "Example User",
                    "email": "example@indent.com"
                  }
                ],
                "properties": {
                  "displayName": {
                    "$id": "#/properties/petitioners/items/allOf/0/properties/displayName",
                    "default": "",
                    "description": "",
                    "examples": [
                      "Example User"
                    ],
                    "title": "The petitioner display name",
                    "type": "string"
                  },
                  "email": {
                    "$id": "#/properties/petitioners/items/allOf/0/properties/email",
                    "default": "",
                    "description": "",
                    "examples": [
                      "example@indent.com"
                    ],
                    "title": "The email schema",
                    "type": "string"
                  }
                },
                "required": [
                  "email",
                  "displayName"
                ],
                "title": "Petitioner",
                "type": "object"
              }
            ]
          },
          "title": "Petitioners",
          "type": "array"
        },
        "reason": {
          "$id": "#/properties/reason",
          "default": "",
          "description": "Why do you need these permissions?",
          "examples": [
            "to view the system logs",
            "to resolve ticket #",
            "to help a customer"
          ],
          "minLength": 5,
          "title": "Reason",
          "type": "string"
        },
        "resources": {
          "$id": "#/properties/resources",
          "default": [
          ],
          "description": "",
          "examples": [
            [
              {
                "displayName": "Example Group",
                "id": "123",
                "kind": "indent.v1.Group"
              }
            ]
          ],
          "items": {
            "$id": "#/properties/resources/items",
            "allOf": [
              {
                "$id": "#/properties/resources/items/allOf/0",
                "default": {
                },
                "description": "",
                "examples": [
                  {
                    "displayName": "Example Group",
                    "id": "123",
                    "kind": "indent.v1.Group"
                  }
                ],
                "properties": {
                  "displayName": {
                    "$id": "#/properties/resources/items/allOf/0/properties/displayName",
                    "default": "",
                    "description": "",
                    "examples": [
                    ],
                    "title": "The displayName schema",
                    "type": "string"
                  },
                  "id": {
                    "$id": "#/properties/resources/items/allOf/0/properties/id",
                    "default": "",
                    "description": "",
                    "examples": [
                    ],
                    "title": "The id schema",
                    "type": "string"
                  },
                  "kind": {
                    "$id": "#/properties/resources/items/allOf/0/properties/kind",
                    "default": "",
                    "description": "",
                    "examples": [
                    ],
                    "title": "The kind schema",
                    "type": "string"
                  }
                },
                "required": [
                  "id",
                  "kind",
                  "displayName"
                ],
                "title": "Resource",
                "type": "object"
              }
            ]
          },
          "title": "Resources",
          "type": "array"
        }
      },
      "required": [
        "petitioners",
        "reason",
        "resources",
        "meta"
      ],
      "title": "Request Access",
      "type": "object"
}`

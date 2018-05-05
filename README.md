---
title: List GDrive Files
weight: 1
---

# Counter
This activity lists the files from Gdrive account.

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo add activity github.com/DipeshTest/gdrivelist
```

## Schema
Inputs and Outputs:

```json
"inputs":[
    {
		"name": "accessToken",
		"type": "string",
    "required": true
	},
	{
		"name": "fileName",
		"type": "string"

	},
	{
		"name": "orderBy",
		"type": "string",
    "allowed": [
        "createdTime",
        "modifiedByMeTime",
        "modifiedTime",
        "name",
        "recency",
        "starred",
        "viewedByMeTime",
        "sharedWithMeTime",
        "quotaBytesUsed"
      ]

	},
  {
		"name": "pageSize",
		"type": "int",
    "value": 50

	},
  {
		"name": "nextPageToken",
		"type": "string"

	},
  {
		"name": "timeout",
		"type": "string",
    "value": "120"

	}

  ],
  "outputs": [
    {
      "name": "statusCode",
      "type": "string"
    },
    {
      "name": "message",
      "type": "any"
    },
    {
      "name": "fileCount",
      "type": "int"
    },
    {
      "name": "nextPageToken",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| accessToken | True     | The access token for your account |         
| fileName   | False    | Name of the file to search |
| orderBy    | False     | Key to sort list of files |  
| pageSize   | False     | The maximum number of files to return per page |
| nextPageToken | False  | The token for continuing a previous list request on the next page. This should be set to the value of 'nextPageToken' from the previous response | 
| timeout       | False    | Timeout value for the delete call, default value is 120 seconds|

## Examples
### Increment
The below example for a sample delete:

```json

```
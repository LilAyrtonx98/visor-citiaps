## JSON Models

The following JSON structures are needed when a new model is created. Retrieved data from database may be different (l.e. "create_at" is only visible when retrieving the data).

### User
```json
{
    "id_group": "string",
    "firstname" : "string",
    "lastname": "string",
    "username": "string",
    "password": "string",
    "permissions": {
        "users": "boolean",
        "layers": "boolean",
        "geo": "boolean",
        "maps": "boolean",
        "visor": "boolean",
        "annotation": "boolean"
    }
}
```

### Category
```json
{
    "name": "string",
    "desc": "string",
}
```

<!-- Category has a struct CategoryWithLayers that has the followed struct:

```json
{
    "name": "string",
    "desc": "string",
    "layers" : "array"
}
```

Note:
    - "layers" is an array of models.Layers -->

### Geoprocessing
```json
{
    "name": "string",
    "desc": "string",
    "geor_url": "string"
}
```

### Group
```json
{
    "name": "string",
    "desc": "string"
}
```

### Map
```json
{
    "name": "string",
    "desc": "string",
    "imgurl": "string"
}
```

### Annotations
```json
{
	"id_user": "string",
	"id_map": "string",
	"id_group": "string",
	"text": "string",
	"location": {
		"type": "Point",
		"coordinates": [99.99, 99.99]
	},
	"is_shared": "boolean"
}
```

Note: 

    - Values of "location" are an example only.
    - "id_group" and "is_shared" are optional, but if "is_shared" is set to true, "id_group" must be given.

### Layers
```json
{
	"id_category": "string",
	"name": "string",
	"desc": "string",
	"provider": {
        "name": "string",
        "url": "string",
        "parsed_url": {
            "protocol": "string",
            "host": "string",
            "port": "string",
            "path": "string"
        },
        "geoserverdata": {
            "service": "string",
            "version": "string",
            "request": "string",
            "max_features": "string",
            "output_format": "string",
            "filename": "string",
            "coordinates_system": "string",
            "workspace": "string",
            "datastore": "string"
        },
    }
}
```

Note:

    - "id_category" is a ObjectID, but can be reciveved as string
    - "provider.name" only recieves values as "file", "arcgis" o "geoserver" (doesn't matter  how is written, it's converted to lowercase) 
    - "provider.geoserverdata" is only recieved if "provider.name" is geoserver
    - "provider.parsed_url" is built in backend service
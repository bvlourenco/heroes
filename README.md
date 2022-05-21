# Hero REST API

### Hero

The Hero object has the following structure:

```javascript
{
    id: primitive.ObjectID,
    name: String
}
```

### Endpoints

The API handles the following HTTP requests:

| METHOD | PATH       | PAYLOAD                        | RESPONSE | DESCRIPTION             |
|--------|------------|--------------------------------|----------|-------------------------|
| GET    | /hero      |                                | [Hero]   | Gets all heroes         |
| POST   | /hero      | `{ name: String }`                           | Hero     | Add a new hero          |
| DELETE | /hero/{id} |                                | Hero     | Removes a hero          |
| GET    | /hero/{id} |                                | Hero     | Gets a specific hero    |
| PUT    | /hero      | `{ id: primitive.ObjectID, name: String }` | Hero     | Changes the hero's name |

### Status codes

* Get /hero

|   Status code  |  Occasion          |
|----------------|--------------------|
|     200        |    Every situation |

* POST /hero

|   Status code  |  Occasion                                      |
|----------------|------------------------------------------------|
|     400        |    Wrong fields, wrong types or empty name     |
|     500        |    Internal Server Error (e.g. Database error) |
|     200        |    Otherwise                                   |

* DELETE /hero/{id}

|   Status code  |  Occasion                                      |
|----------------|------------------------------------------------|
|     400        |    Invalid id (empty or not an integer)        |
|     404        |    Hero not found                              |
|     500        |    Internal Server Error (e.g. Database error) |
|     200        |    Otherwise                                   |

* GET /hero/{id}

|   Status code  |  Occasion                                      |
|----------------|------------------------------------------------|
|     400        |    Invalid id (empty or not an integer)        |
|     404        |    Hero not found                              |
|     500        |    Internal Server Error (e.g. Database error) |
|     200        |    Otherwise                                   |

* PUT /hero

|   Status code  |  Occasion                                      |
|----------------|------------------------------------------------|
|     400        |    Wrong fields, wrong types or empty name     |
|     404        |    Hero not found                              |
|     500        |    Internal Server Error (e.g. Database error) |
|     200        |    Otherwise                                   |
# REST API exercise

## Exercise solution for a REST API project

### Description

The main goal of this exercise is to put in practise all the concepts you have learned about web development.

You will create from scratch the backend of a REST API application, that will listen on some port for HTTP requests, responding in each endpoint with the content that it promises to provide (in this case, information about a generic hero).

Because many of our applications' backend use the REST API paradigm, this constitutes a good exercise for you to get to know these tools before contributing actively for the DevTeam.

A few possible solutions (in different languages) for this exercise are at the end of this page. Feel free to check them, but keep the goal of this exercise in mind: to **learn**. 

### Hero

The Hero object has the following structure:

```javascript
{
    id: Number,
    name: String
}
```

### Endpoints

The API must be able to handle the following HTTP requests:

| METHOD | PATH       | PAYLOAD                        | RESPONSE | DESCRIPTION             |
|--------|------------|--------------------------------|----------|-------------------------|
| GET    | /hero      |                                | [Hero]   | Gets all heroes         |
| POST   | /hero      | Hero                           | Hero     | Add a new hero          |
| DELETE | /hero/{id} |                                | Hero     | Removes a hero          |
| GET    | /hero/{id} |                                | Hero     | Gets a specific hero    |
| PUT    | /hero     | `{ id: Number, name: String }` | Hero     | Changes the hero's name |

### Status codes

**Note:** cheat sheet for status codes: [httpstatuscodes](https://www.restapitutorial.com/httpstatuscodes.html)


* Get /hero

|   Status code  |  Occasion          |
|----------------|--------------------|
|     200        |    Every situation |

* POST /hero

|   Status code  |  Occasion                                     |
|----------------|-----------------------------------------------|
|     400        |    Wrong fields, wrong types or empty name    |
|     409        |    Already exists a hero with this `id`       |
|     409        |    Database error                             |
|     200        |    Otherwise                                  |

* DELETE /hero/{id}

|   Status code  |  Occasion                                     |
|----------------|-----------------------------------------------|
|     400        |    Invalid id (empty or not an integer)       |
|     404        |    Hero not found                             |
|     200        |    Otherwise                                  |

* GET /hero/{id}

|   Status code  |  Occasion                                     |
|----------------|-----------------------------------------------|
|     400        |    Invalid id (empty or not an integer)       |
|     404        |    Hero not found                             |
|     200        |    Otherwise                                  |

* PUT /hero

|   Status code  |  Occasion                                     |
|----------------|-----------------------------------------------|
|     400        |    Wrong fields, wrong types or empty name    |
|     404        |    Hero not found                             |
|     200        |    Otherwise                                  |

### Tests

Every application must have tests, for two reasons:

1. To know that, with some confidence (depending on the granularity and depth of the tests themselves), the application does what it's supposed to. That means, it responds with the expected content, and reacts to errors was it was intended to

2. In the future, guarantee that the changes made to it don't break anything

## Solutions
- [Go](https://github.com/sinfo/hero-tutorial-go)
- [NodeJs](https://github.com/sinfo/hero-tutorial-nodejs)
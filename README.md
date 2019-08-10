# CONTACT API

## How to start ?

Run the api and the database with docker 
```
    $ docker-compose run -d
```

## How to use ?

 - Routes
prefix : v1
Exemple of format of response :

```
{
	Status: 200,
	Meta: null,
	"Data": [
    {
      "ID": 1,
      "CreatedAt": "0001-01-01T00:00:00Z",
      "UpdatedAt": "0001-01-01T00:00:00Z",
      "DeletedAt": null,
      "FirstName": "",
      "LastName": "",
      "PhoneNumbers": null
    }
    ]
}
```

GET /contact
retrieve list of all contacts with FirstName and LastName

GET /contact/:id
retrieve a contact with FirstName, LastName and PhoneNumbers

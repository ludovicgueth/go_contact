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

- Contact routes

GET /v1/contact
retrieve list of all contacts with FirstName and LastName

GET /v1/contact/:id
retrieve a contact with FirstName, LastName and PhoneNumbers

POST /v1/contact
create a new contact, take a json body with FirstName, LastName

PUT /v1/contact/:id
update a contact, take a json body with FirstName, LastName

DELETE /v1/contact/:id
delete a contact with a particular id

- PhoneNumber routes

POST /v1/phoneNumber
Create a new phoneNumber and assossiate it with a particular contact. Take a json with number, name and contactid (number is the phone number, name is the type of phone number, contactid is the id of the particular contact)

DELETE /v1/phoneNumber/:id
Delete a phoneNumber with a particular id

## Summary

With theses routes, Bob and Alice can create new contact and add numbers to them, remove these numbers, edit the contact informations or delete the contact.
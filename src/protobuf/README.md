There are 2 programs in this folder: 
* add_person.go for adding data to addressbook 
* list_people.go to list data from addressbook

To run add_person:
```
cd add
go mod tidy
go run add_person.go ../addressbook
```

To run list_people:
```
cd list
go mod tidy
go run list_people.go ../addressbook
```
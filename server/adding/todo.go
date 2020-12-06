package adding

// Todo ...
type Todo struct {
	Name string `json:"name"`
}

/*
	ID - int
		- auto generated
		- should probably be included by default somehow
	CreatedOn - date
		- auto generated
		- should probably be included by default somehow
	Name - string
		- required
		- maybe length respected
	CompletedOn - date
		- not required
		- defaults to empty meaning incomplete
		- could possibly derive a completed from this to make things easier
	DeletedOn - date
		- not required
		- will update on a delete request
		- probably not needed
	DueDate - date
	Reminder - date
	Importance? Important? Favorite? Ordering?
	Category
*/

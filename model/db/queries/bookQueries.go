package queries

const Delete string = "DELETE FROM Books " +
	"WHERE id = $1"

const Get string = "SELECT * " +
	"FROM Books;"

const Post string = "INSERT INTO Books " +
	"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"

const Put string = "UPDATE Books " +
	"SET [title] = $2" +
	", [author] = $3" +
	", [publisher] = $4" +
	", [language] = $5" +
	", [pages] = $6" +
	", [edition] = $7" +
	", [year] = $8" +
	", [isbn] = $9 " +
	"WHERE [id] = $1"

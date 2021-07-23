package repository

//
//db, err := sql.Open("sqlite3", "./data.sqlite3")
//if err != nil {
//	log.Fatal(err)
//}
//
////language=SQL
//res, err := db.Exec("CREATE TABLE IF NOT EXISTS " +
//	"userinfo(id INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR, departname VARCHAR, created TIMESTAMP)")
//if err != nil {
//	log.Fatal(err)
//}
//
////language=SQL
//res, err = db.Exec("CREATE TABLE IF NOT EXISTS " +
//	"offers(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR, stand VARCHAR, created TIMESTAMP)")
//if err != nil {
//	log.Fatal(err)
//}
//
//
//println(res)
//
//stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values (?,?,?)")
//if err != nil {
//	log.Fatal(err)
//}
//
//res, err = stmt.Exec("astaxie", "asdasd", "2012-12-09")
//if err != nil {
//	log.Fatal(err)
//}
//
//id, err := res.LastInsertId()
//if err != nil {
//	log.Fatal(err)
//}
//
//fmt.Println(id)

package seeds

import (
	"database/sql"
)

type Seed struct {
	Db *sql.DB
}

func (s Seed) ResourceTypeSeed() {
	rId := [2]string{"t01", "t02"}
	rName := [2]string{"Post", "Reservation"}
	for i := 0; i < 1; i++ {
		stmt, _ := s.Db.Prepare(`INSERT INTO resource_type(id, name) VALUES (?,?)`)
		_, err := stmt.Exec(rId[i], rName[i])
		if err != nil {
			panic(err)
		}
	}
}

func (s Seed) ResourceSeed() {
	rId := [5]string{"r01", "r02", "r03", "r04", "r05"}
	rType := [5]string{"Post", "Post", "Reservation", "Reservation", "Reservation"}
	rName := [5]string{"Post1", "Post2", "Reservation1", "Reservation2", "Reservation3"}
	for i := 0; i < 4; i++ {
		stmt, _ := s.Db.Prepare(`INSERT INTO resource(id, name, type) VALUES (?,?,?)`)
		_, err := stmt.Exec(rId[i], rName[i], rType[i])
		if err != nil {
			panic(err)
		}
	}
}

func (s Seed) PermissionSeed() {
	pId := [4]string{"p01", "p02", "p03", "p04"}
	pName := [4]string{"View", "Add", "Cancel", "Edit"}
	for i := 0; i < 3; i++ {
		stmt, _ := s.Db.Prepare(`INSERT INTO permission(id, name) VALUES (?,?)`)
		_, err := stmt.Exec(pId[i], pName[i])
		if err != nil {
			panic(err)
		}
	}
}

func (s Seed) UserSeed() {
	uId := [4]string{"001", "002", "003,", "004"}
	lastName := [4]string{"Chen", "Liu", "Kang", "Li"}
	firstName := [4]string{"One", "Two", "Three", "Four"}
	username := [4]string{"testerchen", "testerliu", "testerkang", "testerli"}
	email := [4]string{"chen@123.com", "liu@456.com", "kang@789.com", "li@10.com"}
	pwd := [4]string{"123", "456", "789", "101"}
	address := [4]string{"A", "B", "C", "D"}
	aptNumber := [4]string{"1", "2", "3", "4"}
	for i := 0; i < 3; i++ {
		stmt, _ := s.Db.Prepare(`INSERT INTO user(id, last_name, first_name, username, email, encrypted_password, last_signed_in_at, address, apt_number) VALUES (?,?,?,?,?,?,?,?,?)`)
		_, err := stmt.Exec(uId[i], lastName[i], firstName[i], username[i], email[i], pwd[i], nil, address[i], aptNumber[i])
		if err != nil {
			panic(err)
		}
	}
}

func (s Seed) PermissionBindingSeed() {
	pid := [4]string{"p01", "p02", "p03", "p04"}
	uid := [4]string{"001", "002", "003", "004"}
	rid := [4]string{"r01", "r02", "r03", "r04"}
	rType := [4]string{"t01", "t01", "t02", "t02"}
	for i := 0; i < 3; i++ {
		stmt, _ := s.Db.Prepare(`INSERT INTO permission_binding(permission_id, user_id, resource_id, resource_type) VALUES (?,?,?,?)`)
		_, err := stmt.Exec(pid[i], uid[i], rid[i], rType[i])
		if err != nil {
			panic(err)
		}
	}
}

func Execute(s Seed) {
	s.ResourceSeed()
	s.PermissionSeed()
	s.ResourceTypeSeed()
	s.UserSeed()
	s.PermissionBindingSeed()
}

//func Execute(db *sql.DB, seedMethodNames ...string) {
//	s := Seed{db}
//
//	seedType := reflect.TypeOf(s)
//
//	// Execute all seeders if no method name is given
//	if len(seedMethodNames) == 0 {
//		log.Println("Running all seeder...")
//		// We are looping over the method on a Seed struct
//		for i := 0; i < seedType.NumMethod(); i++ {
//			// Get the method in the current iteration
//			method := seedType.Method(i)
//			// Execute seeder
//			seed(s, method.Name)
//		}
//	}
//
//	// Execute only the given method names
//	for _, item := range seedMethodNames {
//		seed(s, item)
//	}
//}

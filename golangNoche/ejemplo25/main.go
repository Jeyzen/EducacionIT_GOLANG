package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt *time.Time
}

func main() {
	// DSN: usuario:password@tcp(host:port)/dbname?parseTime=true
	// En XAMPP por defecto root sin password:
	dsn := "root:@tcp(127.0.0.1:3306)/testdb?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error al abrir bd: %v", err)
	}
	defer db.Close()

	// Opciones de pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	// Verificar conexión con contexto timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("no se pudo conectar a la BD: %v", err)
	}
	fmt.Println("Conectado a MySQL (testdb)")

	// --- Ejemplo CRUD ---
	// CREATE
	newID, err := createUser(db, "Horacio", "Horacio@example.com")
	if err != nil {
		log.Fatalf("createUser: %v", err)
	}
	fmt.Printf("Usuario creado con id=%d\n", newID)

	// READ (by id)
	u, err := getUser(db, newID)
	if err != nil {
		log.Fatalf("getUser: %v", err)
	}
	fmt.Printf("Usuario leído: %+v\n", u)

	// UPDATE
	if err := updateUser(db, newID, "Carlos Boy", "carlos.boy@example.com"); err != nil {
		log.Fatalf("updateUser: %v", err)
	}
	fmt.Println("Usuario actualizado")

	// LIST
	users, err := listUsers(db, 10)
	if err != nil {
		log.Fatalf("listUsers: %v", err)
	}
	fmt.Println("Listado usuarios:")
	for _, x := range users {
		fmt.Printf("%+v\n", x)
	}

	// DELETE
	if err := deleteUser(db, newID); err != nil {
		log.Fatalf("deleteUser: %v", err)
	}
	fmt.Println("Usuario borrado")
}

// createUser inserta y devuelve el id generado
func createUser(db *sql.DB, name, email string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO users (name, email) VALUES (?, ?)`
	res, err := db.ExecContext(ctx, stmt, name, email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// getUser obtiene un usuario por id
func getUser(db *sql.DB, id int64) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.QueryRowContext(ctx, `SELECT id, name, email, created_at FROM users WHERE id = ?`, id)
	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, err
	}
	return &u, nil
}

// updateUser actualiza name y email por id
func updateUser(db *sql.DB, id int64, name, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	res, err := db.ExecContext(ctx, stmt, name, email, id)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return fmt.Errorf("ninguna fila afectada")
	}
	return nil
}

// deleteUser elimina por id
func deleteUser(db *sql.DB, id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `DELETE FROM users WHERE id = ?`
	res, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return fmt.Errorf("ninguna fila borrada")
	}
	return nil
}

// listUsers devuelve hasta 'limit' usuarios
func listUsers(db *sql.DB, limit int) ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `SELECT id, name, email, created_at FROM users ORDER BY id DESC LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

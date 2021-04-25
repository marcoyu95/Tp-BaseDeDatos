package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	var db *sql.DB
	menu(db)
}

func menu(db *sql.DB) {
	var eleccion int
	salida := false
	LimpiarPantalla()
	fmt.Printf("1 ---> Crear la base 		\n")
	fmt.Printf("2 ---> Conectarse a la base \n")
	fmt.Printf("3 ---> Crear tablas 		\n")
	fmt.Printf("4 ---> Establecer Pks y Fks \n")
	fmt.Printf("5 ---> Crear stored procedures y triggers \n")
	fmt.Printf("6 ---> Cargar datos \n")
	fmt.Printf("7 ---> Tester consumos \n")
	fmt.Printf("8 ---> Eliminar Pks \n")
	fmt.Printf("9 ---> Generar resumen \n")
	fmt.Printf("10 ---> Salir \n")
	fmt.Printf("\n")
	fmt.Print("Ingresar solicitud:")
	fmt.Scanf("%d", &eleccion)

	switch eleccion {
	case 1:
		crearBaseDatos()

	case 2:
		db = conectarBase()

	case 3:
		if db == nil {
			fmt.Printf("no se a conectado a la base")
			time.Sleep(5 * time.Second)
		} else {
			crearTablas(db)
		}

	case 4:
		agregarPksYFks(db)

	case 5:
		cargarStoredProceduresAndTriggers(db)

	case 6:
		if db == nil {
			fmt.Printf("no se a conectado a la base")
			time.Sleep(5 * time.Second)
		} else {
			cargarDatos(db)
		}

	case 7:
		testerConsumos(db)
	case 8:
		eliminarPks(db)

	case 9:
		generarResumen(db)
	case 10:
		salida = true
		if db != nil {
			db.Close()
		}
	}
	if !salida {
		menu(db)
	}

}

func crearBaseDatos() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`create database db_tarjeta`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SE CREO LA BASE DE DATOS.-")
	time.Sleep(3 * time.Second)
}

func conectarBase() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=db_tarjeta sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SE CONECTO CON EXITO A LA BASE DE DATOS.-")
	time.Sleep(3 * time.Second)
	return db
}

func crearTablas(db *sql.DB) {
	_, err := db.Exec(leerArchivo("./tablas.txt"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SE CREARON CON EXITO LAS TABLAS.-")
	time.Sleep(3 * time.Second)
}

func agregarPksYFks(db *sql.DB) {
	_, err := db.Exec(leerArchivo("./pks y fks.txt"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SE ESTABLECIERON LAS PRIMARY KEY Y FOREIGN KEY.-")
	time.Sleep(3 * time.Second)
}

func cargarStoredProceduresAndTriggers(db *sql.DB) {
	_, err := db.Exec(leerArchivo("./StoredProceduresAndTriggers.txt"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SE CARGARON LOS STORED PROCEDURES Y TRIGGERS.-")
	time.Sleep(3 * time.Second)
}

func cargarDatos(db *sql.DB) {
	cargarClientes(db)
	cargarComercios(db)
	cargarTarjetas(db)
	cargarCierres(db)
	//cargarCompras(db)
	cargarConsumos(db)
	fmt.Println("SE CARGARON LOS DATOS DE: ")
	fmt.Println("-Clientes")
	fmt.Println("-Comercios")
	fmt.Println("-Tarjetas")
	fmt.Println("-Cierres")
	fmt.Println("-Consumos ")
	time.Sleep(5 * time.Second)
}

func cargarClientes(db *sql.DB) {
	//carga de la tabla cliente
	_, err := db.Exec(leerArchivo("./datosCliente.txt"))

	if err != nil {
		log.Fatal(err)
	}
}

func cargarComercios(db *sql.DB) {
	_, err := db.Exec(leerArchivo("./datosComercio.txt"))

	if err != nil {
		log.Fatal(err)
	}

}

func cargarTarjetas(db *sql.DB) {
	_, err := db.Exec(leerArchivo("./datosTarjeta.txt"))

	if err != nil {
		log.Fatal(err)
	}

}

func cargarCierres(db *sql.DB) {
	_, err := db.Exec(leerArchivo("./datosCierre.txt"))
	if err != nil {
		log.Fatal(err)
	}
}

func cargarConsumos(db *sql.DB) {
	_, err := db.Exec(leerArchivo("./datosConsumo.txt"))
	if err != nil {
		log.Fatal(err)
	}
}

func testerConsumos(db *sql.DB) {
	_, err := db.Exec(`select consumos_test();`)
	if err != nil {
		log.Fatal(err)
	}

}

func eliminarPks(db *sql.DB) {
	_, err := db.Exec(leerArchivo("./eliminarPksYFks.txt"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SE ELIMINARON CON EXITO LAS PRIMARY KEY Y FOREIGN KEY.-")
	time.Sleep(3 * time.Second)
}

func generarResumen(db *sql.DB){
	var desde string
	var hasta  string
	var nroCliente string
	var query string

	fmt.Print("Numero de cliente: ")
	fmt.Scanf("%s", &nroCliente)
	fmt.Print("Fecha inicio para el resumen (YYYYMMDD): ")
	fmt.Scanf("%s", &desde)
	fmt.Print("Fecha hasta para el resumen (YYYYMMDD): ")
	fmt.Scanf("%s", &hasta)
	
	query = "select generar_reporte("+nroCliente+",'"+desde+"','"+hasta+"');"
	fmt.Printf(query)

	time.Sleep(10 * time.Second)
	 _, err := db.Exec(query)
	 if err != nil {
	 	log.Fatal(err)
	 }
}

func leerArchivo(s string) string {
	content, err := ioutil.ReadFile(s)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

//LimpiarPantalla limpia la pantalla
func LimpiarPantalla() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func cargarCompras(db *sql.DB) {
	_, err := db.Exec(leerArchivo("./datosCompra.txt"))
	if err != nil {
		log.Fatal(err)
	}
}

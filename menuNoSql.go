package main


import (
	"time"
    "fmt"
    "os"
	"os/exec"
	bolt "github.com/coreos/bbolt"
	"log"
	"strconv"
	"encoding/json"
)

type Cliente struct {
    Nrocliente int
	Nombre string
	Apellido string
	Domicilio string
	Telefono string //char(12)
}

type Tarjeta struct {
    Nrotarjeta string //char(16)
	Nrocliente int
	Validadesde string //char(6)
	Validahasta string //char(6)
	Codseguridad string //char(4)
	Limitecompra float64 //decimal(8,2)
	Estado string //char(10) 
}

type Comercio struct {
    Nrocomercio int
	Nombre string
	Domicilio string
	Codigopostal string //char(8)
	Telefono string //char(12)
}

type Compra struct {
	Nrooperacion  int
	Nrotarjeta string //char(16)
	Nrocomercio int
	Fecha time.Time
	Monto float64 //decimal(7,2)
	Pagado bool
}

func main() {
	salida := false
	var eleccion int
	var db *bolt.DB

	for salida!=true {
		LimpiarPantalla()

		fmt.Printf("1. Crear la base\n")
		fmt.Printf("2. Cargar datos\n")
		fmt.Printf("3. Mostrar datos\n")
		fmt.Printf("4. Salir\n")
		fmt.Print("Ingresar solicitud: ")
		fmt.Scanf("%d", &eleccion)

		fmt.Print(eleccion, "\n")

		switch eleccion {

		case 1:
			db=CrearBase("db_tarjeta.db")
		case 2:
			if db == nil {
				fmt.Printf("no se a conectado a la base")
				fmt.Print("Presione cualquier tecla para volver al menu ...")
				fmt.Scanf("%d", &eleccion)
			} else {
				CargarDatos(db)
			}
		case 3:
			if db == nil {
				fmt.Printf("no se a conectado a la base")
				fmt.Print("Presione cualquier tecla para volver al menu ...")
				fmt.Scanf("%d", &eleccion)
			} else {
				MostrarDatos(db)
			}
		case 4:
			salida = true
			if db != nil {
				db.Close()
			}
		}
    }	
}

func CrearBase(nombre string) *bolt.DB {
	db1, err1 := bolt.Open(nombre, 0600, nil)//abro el archivo
	if err1 != nil {
		log.Fatal(err1)	
	}
	return db1
}

func CargarDatos(db *bolt.DB){
	//clientes
	cliente1 := Cliente{1, "Cristina", "Kirchner", "Flores 945","1545837624"}//creo el cliente
    data, err := json.Marshal(cliente1)
    if err != nil {
        log.Fatal(err)
	}

	CreateUpdate(db, "cliente", []byte(strconv.Itoa(cliente1.Nrocliente)), data)

	cliente2:= Cliente{2,"Homero","Simpson","san martin 7634","1557547964"}
	data, err = json.Marshal(cliente2)
	if err !=nil{
		log.Fatal(err)
	}

	CreateUpdate(db, "cliente", []byte(strconv.Itoa(cliente2.Nrocliente)), data)

	cliente3:= Cliente{3,"Nedwar","Flanders","9 de julio 453","1555437964"}
	data, err = json.Marshal(cliente3)
	if err !=nil{
		log.Fatal(err)
	}

	CreateUpdate(db, "cliente", []byte(strconv.Itoa(cliente3.Nrocliente)), data)

	//tarjetas
	tarjeta1:= Tarjeta{"4017815492979764",cliente1.Nrocliente,"201906","201910","4375",40000,"vigente"}
	data, err = json.Marshal(tarjeta1)
	if err !=nil{
		log.Fatal(err)
	}

	CreateUpdate(db, "tarjeta", []byte(tarjeta1.Nrotarjeta), data)

	tarjeta2:= Tarjeta{"5170299211507992",cliente2.Nrocliente,"201906","201910","1375",30000,"vigente"}
	data, err = json.Marshal(tarjeta2)
	if err !=nil{
		log.Fatal(err)
	}

	CreateUpdate(db, "tarjeta", []byte(tarjeta2.Nrotarjeta), data)

	tarjeta3:= Tarjeta{"5489768187596367",cliente3.Nrocliente,"201906","201910","2371",37000,"vigente"}
	data, err = json.Marshal(tarjeta3)
	if err !=nil{
		log.Fatal(err)
	}

	CreateUpdate(db, "tarjeta", []byte(tarjeta3.Nrotarjeta), data)

	//comercios
	comercio1:= Comercio{1,"El japones","conesa 4201","B1663OVA","123456789012"}
	data, err = json.Marshal(comercio1)
	if err !=nil{
		log.Fatal(err)
	}

	CreateUpdate(db, "comercio", []byte(strconv.Itoa(comercio1.Nrocomercio)), data)

	comercio2:= Comercio{2,"martin.com","san martin 3065","B1663OVA","123456789012"}
	data, err = json.Marshal(comercio2)
	if err !=nil{
		log.Fatal(err)
	}

	CreateUpdate(db, "comercio", []byte(strconv.Itoa(comercio2.Nrocomercio)), data)

	comercio3:= Comercio{3,"garbarino","ferreira 3065","B1663OVA","123456789012"}
	data, err = json.Marshal(comercio3)
	if err !=nil{
		log.Fatal(err)
	}

	CreateUpdate(db, "comercio", []byte(strconv.Itoa(comercio3.Nrocomercio)), data)
	
	//Compras
	now := time.Now()
	compra1:= Compra{1,tarjeta1.Nrotarjeta,comercio1.Nrocomercio,now,12000,false}//cliente1.Nrotarjeta
	data, err = json.Marshal(compra1)
	if err !=nil{
		log.Fatal(err)
	}

	CreateUpdate(db, "compra", []byte(strconv.Itoa(compra1.Nrooperacion)), data)

	now = time.Now()
	compra2:= Compra{2,tarjeta2.Nrotarjeta,comercio2.Nrocomercio,now,13000,false}
	data, err = json.Marshal(compra2)
	if err !=nil{
		log.Fatal(err)
	}

	CreateUpdate(db, "compra", []byte(strconv.Itoa(compra2.Nrooperacion)), data)

	now = time.Now()
	compra3:= Compra{3,tarjeta3.Nrotarjeta,comercio3.Nrocomercio,now,14000,false}
	data, err = json.Marshal(compra3)
	if err !=nil{
		log.Fatal(err)
	}

	CreateUpdate(db, "compra", []byte(strconv.Itoa(compra3.Nrooperacion)), data)

}

func MostrarDatos(db *bolt.DB){
	var seleccion int

	fmt.Printf("Clientes\n")
	resultado, _ := ReadUnique(db, "cliente", []byte(strconv.Itoa(1)))
	fmt.Printf("%s\n", resultado)
	resultado, _ = ReadUnique(db, "cliente", []byte(strconv.Itoa(2)))
	fmt.Printf("%s\n", resultado)
	resultado, _ = ReadUnique(db, "cliente", []byte(strconv.Itoa(3)))
	fmt.Printf("%s\n", resultado)

	fmt.Printf("Tarjetas\n")
	resultado, _ = ReadUnique(db, "tarjeta", []byte("4017815492979764"))
	fmt.Printf("%s\n", resultado)
	resultado, _ = ReadUnique(db, "tarjeta", []byte("5170299211507992"))
	fmt.Printf("%s\n", resultado)
	resultado, _ = ReadUnique(db, "tarjeta", []byte("5489768187596367"))
	fmt.Printf("%s\n", resultado)

	fmt.Printf("Comercios\n")
	resultado, _ = ReadUnique(db, "comercio", []byte(strconv.Itoa(1)))
	fmt.Printf("%s\n", resultado)
	resultado, _ = ReadUnique(db, "comercio", []byte(strconv.Itoa(2)))
	fmt.Printf("%s\n", resultado)
	resultado, _ = ReadUnique(db, "comercio", []byte(strconv.Itoa(3)))
	fmt.Printf("%s\n", resultado)

	fmt.Printf("Compras\n")
	resultado, _ = ReadUnique(db, "compra", []byte(strconv.Itoa(1)))
	fmt.Printf("%s\n", resultado)
	resultado, _ = ReadUnique(db, "compra", []byte(strconv.Itoa(2)))
	fmt.Printf("%s\n", resultado)
	resultado, _ = ReadUnique(db, "compra", []byte(strconv.Itoa(3)))
	fmt.Printf("%s\n", resultado)

	fmt.Print("Presione cualquier tecla para volver al menu ...")
	fmt.Scanf("%d", &seleccion)

}

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
	// abre transacción de escritura
    tx, err := db.Begin(true)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

    err = b.Put(key, val)
    if err != nil {
        return err
    }

    // cierra transacción
    if err := tx.Commit(); err != nil {
        return err
    }

    return nil
}

func ReadUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {
    var buf []byte

    // abre una transacción de lectura
    err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucketName))
        buf = b.Get(key)
        return nil
    })

    return buf, err
}

func LimpiarPantalla() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

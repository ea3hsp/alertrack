package database

// DataBase ...
type DataBase struct {
	Host   string
	Port   string
	DBName string
}

// NewDataBase creates and initializes a Database
func NewDataBase(host, port, dbname string) *DataBase {
	return &DataBase{
		Host:   host,
		Port:   port,
		DBName: dbname,
	}
}

// URL get Url from database instance
func (d *DataBase) URL() string {
	url := "http://" + d.Host + ":" + d.Port + "/" + d.DBName
	return url
}



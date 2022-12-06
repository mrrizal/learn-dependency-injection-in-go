package main

type Row struct{}
type Rows struct {
	Row
}

func (r *Row) Scan(dest ...interface{}) error {
	return nil
}

func (r *Rows) Close() error {
	return nil
}

func (r *Rows) Next() bool {
	return true
}

type rowConverter struct{}

type Person struct {
	Name  string
	Email string
}

// populate the supplied Person from *sql.Row or *sql.Rows object
func (d *rowConverter) populate(in *Person, scan func(dest ...interface{}) error) error {
	return scan(in.Name, in.Email)
}

type LoadPerson struct {
	// compose the row converter into this loader
	rowConverter
}

// dummy implementation of loadFromDB
func (loader *LoadPerson) loadFromDB(id int) Row {
	return Row{}
}

// dummy implementation of loadAllFromDB
func (loader *LoadPerson) loadAllFromDB() Rows {
	return Rows{}
}

func (loader *LoadPerson) ByID(id int) (Person, error) {
	row := loader.loadFromDB(id)
	person := Person{}
	// call the composed "abstract class"
	err := loader.populate(&person, row.Scan)
	return person, err
}

func (loader *LoadPerson) All() ([]Person, error) {
	rows := loader.loadAllFromDB()
	defer rows.Close()
	output := []Person{}
	for rows.Next() {
		person := Person{}
		// call the composed "abstract class"
		err := loader.populate(&person, rows.Scan)
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}

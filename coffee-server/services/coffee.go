package services

import (
	"context"
	"time"
)

//* Coffee type that would hold this data
type Coffee struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Roast string `json:"roast"`
	Image string `json:"image"`
	Region string `json:"region"`
	Price float32 `json:"price"`
	GrindUnit int64 `json:"grind_unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// uses that type in reciever directly --> method belongs to function/var which initiates its instance so it could be used in reciever
// @ Methods belongs to type Coffee ~ the one that will create instance will have these meths
// these all belongs to native Coffee type
// ! Retrieves all coffee data from the db to the Api call
func (c *Coffee) GetAllCoffees()([]*Coffee,error){
	context,cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()

	query := `select id,name,image,region,roast,price,grind_unit,created_at,updated_at from coffees`
	rows,err := db.QueryContext(context,query)
	if err!=nil {
		return nil,err
	}
	defer rows.Close()
	// if got result back successfully
	var coffees []*Coffee
	for rows.Next() {
		var coffee Coffee
		//& scanning each row fields and injecting into coffee type👆
		err := rows.Scan(
			&coffee.ID, 
			&coffee.Name, 
			&coffee.Image, 
			&coffee.Region, 
			&coffee.Roast, 
			&coffee.Price, 
			&coffee.GrindUnit, 
			&coffee.CreatedAt, 
			&coffee.UpdatedAt, 
		)
		if err != nil {
			return nil,err
		}
		// # appending new coffee data to old slice of coffee data els
		coffees = append(coffees, &coffee)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return coffees,nil
	
}

// ! Creates coffee entry in db by passing in coffee data
func (c *Coffee) CreateCoffee(coffee Coffee) (*Coffee,error) {
	// ! if request does not fulfill in this time frame --> don't execute furthur
	context,cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()

	
	query := `
			insert into coffees( name,image,region,roast,price,grind_unit,created_at,updated_at)
			values($1,$2,$3,$4,$5,$6,$7,$8) returning *
		`
		// ! Whatever coffee data would be feeded to this method would create coffee entry in db
	_,err := db.ExecContext(context,query,
		coffee.Name,
		coffee.Image,
		coffee.Region,
		coffee.Roast,
		coffee.Price,
		coffee.GrindUnit,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil,err
	}
	return &coffee,nil
	 
}
// ! Get coffee by the price query
func (c *Coffee) GetCoffeeByqparamsPrice(price float32) (*Coffee,error) {
	context,cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()
	
	// func proceeds
	query := `
		select
			id,name,image,region,roast,price,grind_unit,created_at,updated_at 
		from coffees
		where
			price=$1
	`

	var coffee Coffee
	row := db.QueryRowContext(context,query,price)
	err := row.Scan(
		&coffee.ID,
		&coffee.Name,
		&coffee.Image,
		&coffee.Region,
		&coffee.Roast,
		&coffee.Price,
		&coffee.GrindUnit,
		&coffee.CreatedAt,
		&coffee.UpdatedAt,
	)

	if err != nil {
		return nil,err
	} 
	return &coffee,nil //if successfully queries and injected data into those fields of created type for this data
}


// ! Get coffee by the name
func (c *Coffee) GetCoffeeByName(name string) (*Coffee,error) {
	// method belongs to the Coffee type
	// as this would return Coffee data and ofc error

	context,cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()

	// if successfully function proceeds
	var coffee Coffee //# type struct that stores coffee data 
	query := `
		select 
			id,name,image,region,roast,price,grind_unit,created_at,updated_at 
		from coffees
		where name=$1
	`
	// * make a db call to retrieve coffee data matches query with this name being passed to it
	row := db.QueryRowContext(context,query,name)
	// scan to get each field and injecting into the type which stores this data
	err := row.Scan(
		// ! it scans each field as query gives that too in order gets delivered into the type
		&coffee.ID,
		&coffee.Name,
		&coffee.Image,
		&coffee.Region,
		&coffee.Roast,
		&coffee.Price,
		&coffee.GrindUnit,
		&coffee.CreatedAt,
		&coffee.UpdatedAt,
	)

	// if caught any error scanning fields retrieved from query~db call and injecting data into type
	if err != nil {
		return nil,err
	}	
	//✅✅ successfully got the query data into type instance (let's say) 
	return &coffee,nil

}


// ! Get coffee by the query params {region}
func (c *Coffee) GetCoffeeByQueryParams(region string) (*Coffee,error) {
	// 
	context,cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()

	query := `
		select 
			id,name,image,region,roast,price,grind_unit,created_at,updated_at from coffees
		where 
			region=$1
	`

	var coffee Coffee
	row := db.QueryRowContext(context,query,region)
	err := row.Scan(
		&coffee.ID, 
		&coffee.Name, 
		&coffee.Image, 
		&coffee.Region, 
		&coffee.Roast, 
		&coffee.Price, 
		&coffee.GrindUnit, 
		&coffee.CreatedAt,
		&coffee.UpdatedAt,
	)
	if err != nil {
		return nil,err
	}

	return &coffee,nil
}


// ! Get coffee by the id
func (c *Coffee) GetCoffeeByID(id string) (*Coffee,error) {
	// as this would return single coffee data and ofc error handeling
	context,cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()

	// if not timeouted req --> function proceeds
	var coffee Coffee
	query := `
		select id,name,image,region,roast,price,grind_unit,created_at,updated_at from coffees where id=$1
	`
	row:= db.QueryRowContext(context,query,id)
	// scanning and fetching and injecting values into type that would hold same data
	err := row.Scan(
		&coffee.ID, 
		&coffee.Name, 
		&coffee.Image, 
		&coffee.Region, 
		&coffee.Roast, 
		&coffee.Price, 
		&coffee.GrindUnit, 
		&coffee.CreatedAt,
		&coffee.UpdatedAt,
	)

	if err != nil {
		return nil,err
	}

	// if successfully fetched all values and stored in coffee type struct data --> returning it
	return &coffee,nil
}

// ! Update coffee by the id
func (c *Coffee) UpdateCoffee(id string,body Coffee) (*Coffee,error) {
	// as this would retuned updated coffee and error handeling
	context,cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()

	query := `
		update coffees
		set
			name=$1,
			image=$2,
			region=$3,
			roast=$4,
			price=$5,
			grind_unit=$6,
			updated_at=$7
		where
			id=$8
		returning id,name,image,region,roast,price,grind_unit,created_at,updated_at
	
	`
	var updated Coffee
	err := db.QueryRowContext(context,query,
		body.Name,
		body.Image,
		body.Region,
		body.Roast,
		body.Price,
		body.GrindUnit,
		time.Now(),
		id,
	).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Image,
		&updated.Region,
		&updated.Roast,
		&updated.Price,
		&updated.GrindUnit,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)
	if err!= nil {
		return nil,err
	} 
	return &updated,nil
}


// ! Delete coffee by the id
func (c *Coffee) DeleteCoffeeByID(id string) error {
	//  as this would not return anything but remove coffee from db
//  we keep these func basic functional and do validation when actually passing id when invoking these methods
context,cancel := context.WithTimeout(context.Background(),dbTimeout)
defer cancel()
	
query := `Delete from coffees where id=$1`
_,err := db.ExecContext(context,query,id)
if err!= nil {
	return err
}
return nil
}
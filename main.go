package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type MongoInstance struct {
	Client *mongo.Client
	Db	*mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb+srv://*******:*****@cluster0.2r2bhvc.mongodb.net/" + dbName + "?retryWrites=true&w=majority"

type Employee struct {
	ID     primitive.ObjectID 	`json:"id,omitempty" bson:"_id,omitempty"`
	Name   string	`json:"name"`
	Salary float64	`json:"salary"`
	Age    float64	`json:"age"`
}


func Connect() error{

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil{
		log.Fatal(err)
	}
		ctx, cancel :=context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
		db := client.Database(dbName)

		mg = MongoInstance{
			Client: client,
			Db:     db,
		}
		return nil


}

func main(){
	if err := Connect(); err!= nil{
		panic(err)
	}
	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error{
		var employees []Employee = make([]Employee, 0)
		query := bson.D{{}}

		cursor, err :=mg.Db.Collection("employees").Find(c.Context(),query)
		if err != nil{
			return c.Status(500).SendString(err.Error())
		}

		if err:= cursor.All(c.Context(),&employees); err != nil{
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(employees)

	})
	app.Post("/employee",func(c *fiber.Ctx) error{
		collection := mg.Db.Collection("employees")
		employee := new(Employee)

		if err:=c.BodyParser(employee); err!= nil{
			return c.Status(400).SendString(err.Error())
		}

		// employee.ID = ""

		insertionResult,err := collection.InsertOne(c.Context(),employee)
		if err != nil{
			return c.Status(500).SendString(err.Error())
		}
		filter:= bson.D{{Key:"_id",Value:insertionResult.InsertedID}}
		createdRecord := collection.FindOne(c.Context(),filter)
		createdEmployee := new(Employee)
		createdRecord.Decode(createdEmployee)
		return c.Status(201).JSON(createdEmployee)

		
	})

	app.Put("employee/:id", func (c *fiber.Ctx) error  {
		idParam :=c.Params("id")

		employeeID, err :=primitive.ObjectIDFromHex(idParam)
		if err != nil{
			return c.Status(400).SendString(err.Error())
		}

		employee := new(Employee)
		if err :=c.BodyParser(employee); err != nil{
			return c.Status(400).SendString(err.Error())
		}

		query := bson.D{{Key:"_id",Value:employeeID}}
		update := bson.D{{Key:"$set",
		Value:bson.D{
			{Key:"name",Value:employee.Name},
			{Key:"salary",Value:employee.Salary},
			{Key:"age",Value:employee.Age},
		}}}
		err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(),query,update).Err()
		if err != nil{
			if err == mongo.ErrNoDocuments{
				// return message of err
	      return	c.Status(404).SendString(err.Error())
			}
			return c.Status(500).SendString(err.Error())
		}
		employee.ID = employeeID
		return c.Status(200).JSON(employee)
	})
	app.Delete("employee/:id", func (c *fiber.Ctx) error {
		employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil{
			return c.Status(400).SendString(err.Error())
		}

		query := bson.D{{Key:"_id",Value:employeeID}}
		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(),query)
		if err != nil{
			return c.Status(500).SendString(err.Error())
		}
		if result.DeletedCount < 1{
			return c.SendStatus(404)
		}
		return c.Status(200).JSON("Record deleted")

	})


	log.Fatal(app.Listen(":3000"))
}
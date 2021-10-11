package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"github.com/test/repositories/models"
	"time"

	doctorhdlr "github.com/test/deliveries/doctorlist"
	doctorrepo "github.com/test/repositories/doctors"
	doctoruc "github.com/test/usecases/doctorlist"
	appointmenthdlr "github.com/test/deliveries/appointment"
	appointmentrepo "github.com/test/repositories/appointment"
	appointmentuc "github.com/test/usecases/appointment"
	"github.com/test/utils/middlewares"
	"log"

	"os"

)

func NewGormConnection() (db *gorm.DB, err error) {
	host := os.Getenv("POSTGRES_HOST")
	//port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DBNAME")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, "5432", user, password, dbname)

	db, err = gorm.Open("postgres", psqlInfo)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}

	return db, nil
}

func getGormConnect() *gorm.DB {
	// New DB connection
	db, err := NewGormConnection()

	if err != nil {
		log.Fatal("Database connection error, Let's check:  ", err)
	}
	return db
}

func doMigrate(db *gorm.DB) {
	x := time.Date(int(2021),time.October,int(11),int(8),int(0),int(0),int(0),time.UTC)
	x2 := time.Date(int(2021),time.October,int(11),int(8),int(15),int(0),int(0),time.UTC)
	db.AutoMigrate(
		&models.Doctor{},
		&models.Patience{},
		&models.Appointment{},
		&models.Schedule{},
	)
	db.Save(&models.Doctor{
		Model: gorm.Model{ID:1},
		DoctorID: "001",
		Name: "หมอ ก.",
	})
	db.Save(&models.Doctor{
		Model: gorm.Model{ID:2},
		DoctorID: "002",
		Name: "หมอ ข.",
	})
	db.Save(&models.Patience{
		Model: gorm.Model{ID:1},
		PinCode: "111111",
		Name: "นายกร",
		Phone:"0810000001",
	})
	db.Save(&models.Patience{
		Model: gorm.Model{ID:2},
		PinCode: "222222",
		Name: "นายนก",
		Phone:"0810000002",
	})
	db.Save(&models.Patience{
		Model: gorm.Model{ID:3},
		PinCode: "333333",
		Name: "นายตนู",
		Phone:"0810000003",
	})
	db.Save(&models.Patience{
		Model: gorm.Model{ID:4},
		PinCode: "444444",
		Name: "นายหมาย",
		Phone:"0810000004",
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:1},
		Ref: "X1",
		DoctorID: "001",
		Start: &x,
		End: &x2,
	})
}

func startGinEngine() {

	gin.Default()
	ginEngine := gin.New()
	ginEngine.Use(middlewares.CORS)

	db := getGormConnect()
	doMigrate(db)

	doctorRepo := doctorrepo.NewDoctorRepository(doctorrepo.DoctorRepositoryParams{
		DB: db,
	})
	doctorUC := doctoruc.NewUseCase(doctorRepo)
	doctorhdlr.NewEndpointHttpHandler(ginEngine, doctorUC)

	appointmentRepo := appointmentrepo.NewRepository(appointmentrepo.ParamsAppointmentRepository{
		DB: db,
	})
	appointmentUC := appointmentuc.NewUseCase(appointmentRepo)
	appointmenthdlr.NewEndpointHttpHandler(ginEngine, appointmentUC)
	// Metric
	prom := ginprometheus.NewPrometheus("gin")
	prom.Use(ginEngine)
	// Gin Engine Run
	if err := ginEngine.Run(":9098"); err != nil {
		panic(err)
	}

}

func main() {
	startGinEngine()
}

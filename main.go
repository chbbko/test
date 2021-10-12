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

func makeDate(y,mo,d,h,m int ) *time.Time{
	x := time.Date(int(y),time.Month(mo),int(d),int(h),int(m),int(0),int(0),time.UTC)
	return &x
}


func doMigrate(db *gorm.DB) {
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
		Start: makeDate(2021,10,11,8,0),
		End: makeDate(2021,10,11,9,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:2},
		Ref: "X2",
		DoctorID: "001",
		Start: makeDate(2021,10,11,9,1),
		End: makeDate(2021,10,11,10,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:3},
		Ref: "X3",
		DoctorID: "001",
		Start: makeDate(2021,10,11,10,1),
		End: makeDate(2021,10,11,11,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:4},
		Ref: "X4",
		DoctorID: "001",
		Start: makeDate(2021,10,11,11,1),
		End: makeDate(2021,10,11,12,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:5},
		Ref: "X5",
		DoctorID: "001",
		Start: makeDate(2021,10,13,8,0),
		End: makeDate(2021,10,13,9,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:6},
		Ref: "X6",
		DoctorID: "001",
		Start: makeDate(2021,10,13,9,1),
		End: makeDate(2021,10,13,10,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:7},
		Ref: "X7",
		DoctorID: "001",
		Start: makeDate(2021,10,13,10,1),
		End: makeDate(2021,10,13,11,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:8},
		Ref: "X8",
		DoctorID: "001",
		Start: makeDate(2021,10,13,11,1),
		End: makeDate(2021,10,13,12,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:9},
		Ref: "X9",
		DoctorID: "001",
		Start: makeDate(2021,10,15,8,0),
		End: makeDate(2021,10,15,9,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:10},
		Ref: "X10",
		DoctorID: "001",
		Start: makeDate(2021,10,15,9,1),
		End: makeDate(2021,10,15,10,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:11},
		Ref: "X11",
		DoctorID: "001",
		Start: makeDate(2021,10,15,10,1),
		End: makeDate(2021,10,15,11,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:12},
		Ref: "X12",
		DoctorID: "001",
		Start: makeDate(2021,10,15,11,1),
		End: makeDate(2021,10,15,12,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:13},
		Ref: "V01",
		DoctorID: "002",
		Start: makeDate(2021,10,12,13,0),
		End: makeDate(2021,10,12,13,30),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:14},
		Ref: "V02",
		DoctorID: "002",
		Start: makeDate(2021,10,12,13,31),
		End: makeDate(2021,10,12,14,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:15},
		Ref: "V03",
		DoctorID: "002",
		Start: makeDate(2021,10,12,14,1),
		End: makeDate(2021,10,12,14,30),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:16},
		Ref: "V04",
		DoctorID: "002",
		Start: makeDate(2021,10,12,14,31),
		End: makeDate(2021,10,12,15,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:17},
		Ref: "V05",
		DoctorID: "002",
		Start: makeDate(2021,10,14,13,0),
		End: makeDate(2021,10,14,13,30),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:18},
		Ref: "V06",
		DoctorID: "002",
		Start: makeDate(2021,10,14,13,31),
		End: makeDate(2021,10,14,14,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:19},
		Ref: "V07",
		DoctorID: "002",
		Start: makeDate(2021,10,14,14,1),
		End: makeDate(2021,10,14,14,30),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:20},
		Ref: "V08",
		DoctorID: "002",
		Start: makeDate(2021,10,14,14,31),
		End: makeDate(2021,10,14,15,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:21},
		Ref: "V09",
		DoctorID: "002",
		Start: makeDate(2021,10,16,13,0),
		End: makeDate(2021,10,16,13,30),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:22},
		Ref: "V10",
		DoctorID: "002",
		Start: makeDate(2021,10,16,13,31),
		End: makeDate(2021,10,16,14,0),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:23},
		Ref: "V11",
		DoctorID: "002",
		Start: makeDate(2021,10,16,14,1),
		End: makeDate(2021,10,16,14,30),
	})
	db.Save(&models.Schedule{
		Model: gorm.Model{ID:24},
		Ref: "V12",
		DoctorID: "002",
		Start: makeDate(2021,10,16,14,31),
		End: makeDate(2021,10,16,15,0),
	})
	db.Save(&models.Appointment{
		Model: gorm.Model{ID:1},
		PatiencePinCode: "333333",
		DoctorID: "001",
		Phone: "0810000003",
		Start: *makeDate(2021,10,15,10,1),
		End: *makeDate(2021,10,15,11,0),
		Ref: "X11",
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

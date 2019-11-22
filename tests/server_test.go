package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "github.com/chatrasen/mer-weekend/dbase"
	// . "github.com/chatrasen/mer-weekend/server"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// func getData(fileName string, v interface{}) {
// 	file, _ := os.Open(fileName)
// 	defer file.Close()
// 	byteValue, _ := ioutil.ReadAll(file)
// 	json.Unmarshal(byteValue, v)
// }

var _ = Describe("Server", func() {
	var (
		// db       *gorm.DB
		router   *gin.Engine
		response *httptest.ResponseRecorder
	)

	// BeforeEach(func() {
	// 	d, err := gorm.Open("mysql", "root:lazarus@/dbmerweekendtest?charset=utf8&parseTime=True&loc=Local")
	// 	db = d
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	router = CreateRouter(db)

	// 	dbase.Init(db)

	// 	var users []models.User
	// 	getData("../dbase/data/users.json", &users)

	// 	var events []models.Event
	// 	getData("../dbase/data/events.json", &events)

	// 	for _, user := range users {
	// 		db.Create(&user)
	// 	}

	// 	for _, event := range events {
	// 		db.Create(&event)
	// 	}

	// })

	// AfterEach(func() {
	// 	// TODO: Figure out a way to empty the database at the end
	// 	db.Exec(`DROP TABLE IF EXISTS users;`)
	// 	db.Exec(`DROP TABLE IF EXISTS events;`)
	// 	db.Close()
	// })

	Describe("Version 1 API at /api/v1", func() {
		Describe("The / endpoint", func() {
			BeforeEach(func() {
				response = performRequest(router, "GET", "/api/v1/", nil)
			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("Returns the String 'Hello World'", func() {
				Expect(response.Body.String()).To(Equal("Hello World"))
			})
		})
	})

	// Describe("Version 1 API at /api/v1", func() {
	// 	Describe("The /events endpoint", func() {
	// 		BeforeEach(func() {
	// 			response = performRequest(router, "GET", "/api/v1/events", nil)
	// 		})

	// 		It("Returns status 200 message", func() {
	// 			Expect(response.Code).To(Equal(200))
	// 		})

	// 		It("Returns a JSON conataining a list of all the events", func() {
	// 			var jsonAllEvents, expectedAllEvents []models.Event
	// 			json.Unmarshal(response.Body.Bytes(), &jsonAllEvents)
	// 			fileData, _ := ioutil.ReadFile("../dbase/data/events.json")
	// 			json.Unmarshal(fileData, &expectedAllEvents)
	// 			var jsonAllEventTitles, expectedAllEventTitles []string
	// 			for _, event := range jsonAllEvents {
	// 				jsonAllEventTitles = append(jsonAllEventTitles, event.Title)
	// 			}
	// 			for _, event := range expectedAllEvents {
	// 				expectedAllEventTitles = append(expectedAllEventTitles, event.Title)
	// 			}
	// 			Expect(jsonAllEventTitles).To(Equal(expectedAllEventTitles))
	// 		})
	// 	})
	// })

	// Describe("Version 1 API at /api/v1", func() {
	// 	Describe("The /users endpoint", func() {
	// 		BeforeEach(func() {
	// 			response = performRequest(router, "GET", "/api/v1/users", nil)
	// 		})

	// 		It("Returns status 200 message", func() {
	// 			Expect(response.Code).To(Equal(200))
	// 		})

	// 		It("Returns a JSON conataining a list of all the users", func() {
	// 			var jsonAllUsers, expectedAllUsers []models.User
	// 			json.Unmarshal(response.Body.Bytes(), &jsonAllUsers)
	// 			fileData, _ := ioutil.ReadFile("../dbase/data/users.json")
	// 			json.Unmarshal(fileData, &expectedAllUsers)
	// 			var jsonAllUsernames, expectedAllUsernames []string
	// 			for _, user := range jsonAllUsers {
	// 				jsonAllUsernames = append(jsonAllUsernames, user.Name)
	// 			}
	// 			for _, user := range expectedAllUsers {
	// 				expectedAllUsernames = append(expectedAllUsernames, user.Name)
	// 			}
	// 			Expect(jsonAllUsernames).To(Equal(expectedAllUsernames))
	// 		})
	// 	})

	// 	Describe("The /users/:id endpoint", func() {
	// 		BeforeEach(func() {
	// 			response = performRequest(router, "GET", "/api/v1/users/2", nil)
	// 		})

	// 		It("Returns status 200 message.", func() {
	// 			Expect(response.Code).To(Equal(200))
	// 		})

	// 		It("Returns the user with ID 2", func() {
	// 			var jsonUser models.User
	// 			json.Unmarshal(response.Body.Bytes(), &jsonUser)
	// 			Expect(jsonUser.Name).To(Equal("customer_two"))
	// 		})
	// 	})

	// 	Describe("Invalid user ID", func() {
	// 		BeforeEach(func() {
	// 			response = performRequest(router, "GET", "/api/v1/users/blah-1", nil)
	// 		})

	// 		It("Returns bad request error code.", func() {
	// 			Expect(response.Code).To(Equal(400))
	// 		})
	// 	})

	// 	Describe("The /users/:id/organized_events endpoint", func() {
	// 		BeforeEach(func() {
	// 			response = performRequest(router, "GET", "/api/v1/users/2/organized_events", nil)
	// 		})

	// 		It("Returns 200 status code", func() {
	// 			Expect(response.Code).To(Equal(200))
	// 		})

	// 		It("Returns a list of all events organized by the user with userID 2", func() {
	// 			var jsonEvents []models.Event
	// 			json.Unmarshal(response.Body.Bytes(), &jsonEvents)
	// 			var jsonAllEventTitles []string
	// 			for _, event := range jsonEvents {
	// 				jsonAllEventTitles = append(jsonAllEventTitles, event.Title)
	// 			}
	// 			expectedEventTitles := []string{"Title two"}
	// 			Expect(jsonAllEventTitles).To(Equal(expectedEventTitles))

	// 		})
	// 	})
	// })
})

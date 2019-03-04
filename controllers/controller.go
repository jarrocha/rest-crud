package controllers

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"gitlab.com/codelittinc/golang-interview-project-jaime/rest-crud/models"
)

var HC *HireController

type HireController struct {
	Client     *mgo.Session       // points to MongoDB, used to return collection to point to Database("db").Collection("col")
	Collection *mgo.Collection    // used to perform MongoDB commands
	DBContext  context.Context    // current context for db
	CancelFunc context.CancelFunc // cancel function for context
}

func NewHireController(c *mgo.Session, col *mgo.Collection, ctx context.Context, cf context.CancelFunc) *HireController {
	return &HireController{
		Client:     c,
		Collection: col,
		DBContext:  ctx,
		CancelFunc: cf,
	}
}

func (hc *HireController) AllHires() ([]models.User, error) {
	data := []models.User{}
	if hc != nil {
		err := hc.Collection.Find(bson.M{}).All(&data)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (hc *HireController) OneHire(r *http.Request) (models.User, error) {
	data := models.User{}
	name, err := StrFormValue(r.FormValue("name"), true)
	if err != nil {
		return data, err
	}

	if hc != nil {
		err = hc.Collection.Find(bson.M{"name": name}).One(&data)
		if err != nil {
			return data, err
		}
	}

	return data, nil
}

func (hc *HireController) PutHire(r *http.Request) (models.User, error) {
	// get form values
	data := models.User{}
	data.Name = r.FormValue("name")
	data.Type = r.FormValue("type")
	data.Role = r.FormValue("role")

	// handle duration -- TODO: move to function/unit-test
	var err error
	data.Duration, err = IntFormValue(r.FormValue("duration"), true)
	if err != nil {
		return data, errors.New(err.Error() + ". Duration must be a number")
	}

	// handle tags - parse and assign
	data.Tags, _ = SliceFormValue(r.FormValue("tags"), false)

	// validate form values -- TODO: move to function/unit-test
	if data.Name == "" || data.Type == "" {
		return data, errors.New("400. Bad request. All fields must be complete. Type must be employee or contractor")
	} else if strings.ToLower(data.Type) == "employee" && data.Role == "" {
		return data, errors.New("400. Bad request. Employee requires role.")
	} else if strings.ToLower(data.Type) == "contractor" && data.Duration == 0 {
		return data, errors.New("400. Bad request. Contractor requires duration greater than 0.")
	}

	// insert values
	if hc != nil {
		err = hc.Collection.Insert(data)
		if err != nil {
			return data, errors.New("500. Internal Server Error." + err.Error())
		}
	}
	return data, nil
}

func (hc *HireController) UpdateHire(r *http.Request) (models.User, error) {
	// get form values
	data := models.User{}
	data.Name = r.FormValue("name")
	data.Type = r.FormValue("type")
	data.Role = r.FormValue("role")
	// handle duration -- TODO: move to function/unit-test
	var err error
	data.Duration, err = IntFormValue(r.FormValue("duration"), true)
	if err != nil {
		return data, errors.New(err.Error() + ". Duration must be a number")
	}

	// validate form values -- TODO: move to function/unit-test
	if data.Name == "" || data.Type == "" {
		return data, errors.New("400. Bad request. All fields must be complete. Type must be employee or contractor")
	} else if strings.ToLower(data.Type) == "employee" && data.Role == "" {
		return data, errors.New("400. Bad request. Employee requires role.")
	} else if strings.ToLower(data.Type) == "contractor" && data.Duration == 0 {
		return data, errors.New("400. Bad request. Contractor requires duration greater than 0.")
	}

	// handle tags - parse and assign
	data.Tags, _ = SliceFormValue(r.FormValue("tags"), false)

	// update values
	if hc != nil {
		err = hc.Collection.Update(bson.M{"name": data.Name}, &data)
		if err != nil {
			return data, err
		}
	}
	return data, nil
}

func (hc *HireController) DeleteHire(r *http.Request) error {
	name, err := StrFormValue(r.FormValue("name"), true)
	if err != nil {
		return err
	}

	if hc != nil {
		err = hc.Collection.Remove(bson.M{"name": name})
		if err != nil {
			return errors.New("500. Internal Server Error")
		}
	}
	return nil
}

func StrFormValue(str string, needed bool) (string, error) {
	if str == "" {
		if needed == false {
			return "", nil
		} else {
			return "", errors.New("400. Bad Request.")
		}
	}
	return str, nil
}

func SliceFormValue(str string, needed bool) ([]string, error) {
	if len(str) == 0 {
		if needed == false {
			return nil, nil
		} else {
			return nil, errors.New("400. Bad Request.")
		}
	}
	return strings.Split(str, ","), nil
}

func IntFormValue(str string, needed bool) (int, error) {
	i64, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		if needed == false {
			return 0, nil
		} else {
			return 0, errors.New("406. Not Acceptable.")
		}
	}
	return int(i64), nil
}

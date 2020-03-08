package repository

import(
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/rafaelarodrigues90/go-apirest/src/modules/profile/model"
)

// userRepositoryMongo
type userRepositoryMongo struct {
	db * mgo.Database
	collection string
}

// NewUserRepositoryMongo
func NewUserRepositoryMongo(db *mgo.Database, collection string) *userRepositoryMongo {
	return &userRepositoryMongo {
		db: db,
		collection: collection,
	}
}

// Save
func (r *userRepositoryMongo) Save (user *model.User) error {
	err := r.db.C(r.collection).Insert(user)
	return err
}

// Update
func (r *userRepositoryMongo) Update (id string, user *model.User) error {
	err := r.db.C(r.collection).Update(bson.M{"id": id}, user)
	return err
}

// Delete
func (r *userRepositoryMongo) Delete (id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"id": id})
	return err
}

// Find By ID
func (r *userRepositoryMongo) FindByID (id string) (*model.User, error) {
	var user model.User
	err := r.db.C(r.collection).Find(bson.M{"id": id}).One(&user)

	if err != nil{
		return nil, err
	}

	return &user, nil
}

// Find All
func (r *userRepositoryMongo) FindAll () (model.Users, error) {
	var users model.Users
	err := r.db.C(r.collection).Find(bson.M{}).All(&users)

	if err != nil{
		return nil, err
	}

	return users, nil
}
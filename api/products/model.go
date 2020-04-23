package products

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	//Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Available   bool               `json:"disponible,omitempty" bson:"disponible,omitempty"`
	Name        string             `json:"nombre,omitempty" bson:"nombre,omitempty"`
	Description string             `json:"descripcion,omitempty" bson:"descripcion,omitempty"`
	Price       int32              `json:"precioUni,omitempty" bson:"precioUni,omitempty"`
	Category    primitive.ObjectID `json:"categoria,omitempty" bson:"categoria,omitempty"`
	User        primitive.ObjectID `json:"usuario,omitempty" bson:"usuario,omitempty"`
	Image       string             `json:"img,omitempty" bson:"img,omitempty"`
}

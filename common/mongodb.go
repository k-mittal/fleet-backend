package common


import "gopkg.in/mgo.v2"


func ConnectToMongo() (*mgo.Session, error){
	session,err:= mgo.Dial("localhost:27017")
	if err!=nil{
		return nil,err
} else {
	return session,nil
}
}



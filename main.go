package main

import (
	"BISBackend/Database"
	"BISBackend/Web"
)

func main() {
	db := Database.NewDatabase()
	users := []Database.User{
		{Name: "Cris", Image: "https://steamuserimages-a.akamaihd.net/ugc/782989521799828488/99CC5CFF94186C28A21C7E81D145E9C6550DCC14/?imw=512&amp;imh=499&amp;ima=fit&amp;impolicy=Letterbox&amp;imcolor=%23000000&amp;letterbox=true", Address: "111", Secret: "some"},
		{Name: "Sota", Image: "https://forum.service-cm.ru/media/kunena/avatars/users/avatar906.jpg", Address: "123", Secret: "hahaha"},
		{Name: "panda", Image: "https://yt3.googleusercontent.com/_2Btm4zDSX0bAGdmENkNPjReO8h8YEN_qPh5bpiPdEv2_UX3CByjJueloJE5N8f3Cuci4yx0ew=s900-c-k-c0x00ffffff-no-rj", Address: "453", Secret: "ritropictor"},
	}
	db.Create(users)
	db.Commit()
	Web.RunServer()
}

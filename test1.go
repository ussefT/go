package main

import(
	"fmt"
	"os"
	"log"
)

func Folder() (string){
	dir,err:=os.Getwd()
	if err!=nil{
		log.Fatal(err)
	}
	return dir
}

func FileFolder(dir string)(files []os.DirEntry){
	files,err:=os.ReadDir(dir)
	if err!=nil{
		log.Fatal()
	}

	return
}
func main(){

	// get current dir
	dir :=Folder()
	// get files in folder
	files:=FileFolder(dir)
	fmt.Printf("dir:%s\n",dir)
	
	fmt.Println(files)

	for _,f :=range files{
	
	fmt.Printf("%s\n",f.Name())
	

		
	fmt.Println()

	fmt.Printf("%s\n",f.Name())
	fmt.Printf("%s\n",f.Name())

	}

}
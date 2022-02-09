package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"

        "encoding/base64"

        "github.com/gorilla/mux"
)

//variable to save the amount of values into the json file
var json_num int16

//varriable to continue with the execution of the script
var continue_var bool

//We declare the struture of Datos
type Datos struct {
        Datos []Value `json:"data"`
}

// We declare the struture to obtain the information from the json file
type Value struct {
        Valor        string `json:"value"`
        Is_Encrypted bool   `json:"encrypted"`
}

//Function to write the data into the json file
func write_data_into_json_file(dato_to_write string, encrypted_value bool) {
        //Initialize the variable
        response := ""

        //We are going to set the data to be writed in the file
        if encrypted_value == true {
                response = `[{"value":"` + dato_to_write + `","encrypted":true}]`
        } else {
                response = `[{"value":"` + dato_to_write + `","encrypted":false}]`
        }

        //Write the file
        _ = ioutil.WriteFile("file.json", []byte(response), 0644)
}

//Function to proceed to encrypt the string and write it into the json file
func encrypt_data(strg string, w http.ResponseWriter) {
        //Show message of encrypting
        fmt.Fprintln(w, "Encrypting file ...")

        //encript the string
        Encoding := base64.StdEncoding.EncodeToString([]byte(strg))

        //Proceed to write the information into the file
        write_data_into_json_file(Encoding, true)

        //Print both values
        fmt.Fprintln(w, "Data: "+strg)
        fmt.Fprintln(w, "Encrypted Data: "+Encoding)

        //Show message of encrypting
        fmt.Fprintln(w, "File has been correctly encrypted                              [OK]")

}

//Function to proceed to encrypt the string and write it into the json file
func decrypt_data(strg string, w http.ResponseWriter) {
        //Show message of encrypting
        fmt.Fprintln(w, "Decrypting file ...")

        //Call the function to decrypt the string
        Decrypting, err := base64.StdEncoding.DecodeString(strg)
        //If some error is show
        if err != nil {
                fmt.Fprintln(w, "Error at the moment of decrypt data: "+err.Error())
                panic(err)
        } else {
                //Proceed to write the information into the file
                write_data_into_json_file(string(Decrypting), false)
                //Show confirmation message
                fmt.Fprintln(w, "The data has been correctly decrypted                          [OK]")
                //fmt.Println(string(Decrypting))
                //Print the value
                fmt.Fprintln(w, "Decrypted value: "+string(Decrypting))
        }
}

//function to handle the methong get of the api encrypt
func HandleGetMethod_encrypt(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Encrypting call")
        fmt.Fprintln(w, "----------------------------------------")
        //Show message to let the user we are opening the json file
        fmt.Fprintln(w, "Opening json file...")

        //Open our jsonfile
        jsonFile, err := os.Open("file.json")

        //if we os.Open returns an error then handle it
        if err != nil {
                //print the error collected in both sides and stop the execution of the function
                fmt.Fprintln(w, err.Error())
                log.Fatal(err.Error())
        } else {
                //If we found the file then
                fmt.Fprintln(w, "File is located, we continue with the encryption               [OK]")

                //We set the variable contnue_var to true
                continue_var = true
        }

        //defer the closing of our jsonFile so that we can parse it later on
        defer jsonFile.Close()

        //If the variable is true, means can continue
        if continue_var == true {
                //Initialize the data array
                var data []Value

                //Variable to load the content of the file
                content, erro := ioutil.ReadFile("file.json")

                //If reading the file it encount any error
                if erro != nil {
                        //Display the error
                        fmt.Fprintln(w, "Error: "+erro.Error())
                        log.Fatal(erro.Error())
                }

                //Now load the content of content into the data structure
                err2 := json.Unmarshal(content, &data)
                //If found any error
                if err2 != nil {
                        //It going to show error message
                        fmt.Fprintln(w, "Error: "+err2.Error())
                        log.Fatal(err2.Error())
                }

                //Set the json_num with the amount of objects in the file
                json_num := len(data)

                //If the file does not contains data
                if json_num == 0 {
                        //Shows Error message
                        fmt.Fprintln(w, "File is empty, please insert a data value in the file        [Error]")
                } else {
                        //If the json file contains a single value
                        if json_num == 1 {
                                //We print a confirmation message:
                                fmt.Fprintln(w, "The file contains a single value                               [OK]")

                                // we iterate through the array to collect the data and encrypt it
                                for _, x := range data {

                                        //Obtain the information to know if the string has been aready encrypted
                                        is_already_Encrypted := bool(x.Is_Encrypted)

                                        //Check if the file is already encrypted
                                        if is_already_Encrypted == false {
                                                //Print a confirmation message to let know the file is not encrypted
                                                fmt.Fprintln(w, "File is not encrypted                                          [OK]")

                                                //If not, proceed with the encription of the data
                                                encrypt_data(x.Valor, w)
                                        } else {
                                                //If yes, shows error message
                                                fmt.Fprintln(w, "The file is already encryted                                [Error]")
                                        }
                                }
                        } else {
                                //If there is more than 1 value, we show error messaga and do not continue
                                fmt.Fprintln(w, "The file contains more than a single value        [Error]")
                        }

                }

        }

}

func HandleGetMethod_decrypt(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Decrypting call")

        //Show message to let the user we are opening the json file
        fmt.Fprintln(w, "Opening json file...")

        //Open our jsonfile
        jsonFile, err := os.Open("file.json")

        //if we os.Open returns an error then handle it
        if err != nil {
                //print the error collected in both sides and stop the execution of the function
                fmt.Fprintln(w, err.Error())
                log.Fatal(err.Error())
        } else {
                //If we found the file then
                fmt.Fprintln(w, "File is located, we continue with the decryption               [OK]")

                //We set the variable contnue_var to true
                continue_var = true
        }

        //defer the closing of our jsonFile so that we can parse it later on
        defer jsonFile.Close()

        //If the variable is true, means can continue
        if continue_var == true {
                //Initialize the data array
                var data []Value

                //Variable to load the content of the file
                content, erro := ioutil.ReadFile("file.json")

                //If reading the file it encount any error
                if erro != nil {
                        //Display the error
                        fmt.Fprintln(w, "Error: "+erro.Error())
                        log.Fatal(erro.Error())
                }

                //Now load the content of content into the data structure
                err2 := json.Unmarshal(content, &data)
                //If found any error
                if err2 != nil {
                        //It going to show error message
                        fmt.Fprintln(w, "Error: "+err2.Error())
                        log.Fatal(err2.Error())
                }

                //Set the json_num with the amount of objects in the file
                json_num := len(data)

                //If the file does not contains data
                if json_num == 0 {
                        //Shows Error message
                        fmt.Fprintln(w, "File is empty, please insert a data value in the file        [Error]")
                } else {
                        //If the json file contains a single value
                        if json_num == 1 {
                                //We print a confirmation message:
                                fmt.Fprintln(w, "The file contains a single value                               [OK]")

                                // we iterate through the array to collect the data and encrypt it
                                for _, x := range data {

                                        //Obtain the information to know if the string has been aready encrypted
                                        is_already_Encrypted := bool(x.Is_Encrypted)

                                        //Check if the file is already encrypted
                                        if is_already_Encrypted == true {
                                                //Print a confirmation message to let know the file is not encrypted
                                                fmt.Fprintln(w, "File is encrypted                                              [OK]")

                                                //If not, proceed with the encription of the data
                                                decrypt_data(x.Valor, w)
                                        } else {
                                                //If yes, shows error message
                                                fmt.Fprintln(w, "The file is already decryted                                [Error]")
                                        }
                                }
                        } else {
                                //If there is more than 1 value, we show error messaga and do not continue
                                fmt.Fprintln(w, "The file contains more than a single value        [Error]")
                        }

                }

        }

}

func HandleGetMethod_show_jsonfile(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Printing JSON File content")
        fmt.Fprintln(w, "----------------------------------------")

        //Open our jsonfile
        jsonFile, err := os.Open("file.json")

        //if we os.Open returns an error then handle it
        if err != nil {
                //print the error collected in both sides and stop the execution of the function
                fmt.Fprintln(w, err.Error())
                log.Fatal(err.Error())
        } else {
                //We set the variable contnue_var to true
                continue_var = true
        }

        //defer the closing of our jsonFile so that we can parse it later on
        defer jsonFile.Close()

        //If the variable is true, means can continue
        if continue_var == true {
                //Initialize the data array
                var data []Value

                //Variable to load the content of the file
                content, erro := ioutil.ReadFile("file.json")

                //If reading the file it encount any error
                if erro != nil {
                        //Display the error
                        fmt.Fprintln(w, "Error: "+erro.Error())
                        log.Fatal(erro.Error())
                }

                //Now load the content of content into the data structure
                err2 := json.Unmarshal(content, &data)
                //If found any error
                if err2 != nil {
                        //It going to show error message
                        fmt.Fprintln(w, "Error: "+err2.Error())
                        log.Fatal(err2.Error())
                }

                //Set the json_num with the amount of objects in the file
                json_num := len(data)

                //If the file does not contains data
                if json_num == 0 {
                        //Shows Error message
                        fmt.Fprintln(w, "File is empty, please insert a data value in the file        [Error]")
                } else {
                        //If the json file contains a single value
                        if json_num == 1 {

                                // we navegate through the array to collect the data and encrypt it
                                for _, x := range data {

                                        //Obtain the information to know if the string has been aready encrypted
                                        is_already_Encrypted := bool(x.Is_Encrypted)

                                        //Check if the file is already encrypted
                                        if is_already_Encrypted == false {
                                                //Print a confirmation message to let know the file is not encrypted
                                                fmt.Fprintln(w, "File is not encrypted.")

                                                //If not, proceed with the encription of the data
                                                fmt.Fprintln(w, "Current Value: ", x.Valor)
                                        } else {
                                                //Print a confirmation message to let know the file is not encrypted
                                                fmt.Fprintln(w, "File is encrypted.")

                                                //If not, proceed with the encription of the data
                                                fmt.Fprintln(w, "Current Value: ", x.Valor)
                                        }
                                }
                        } else {
                                //If there is more than 1 value, we show error messaga and do not continue
                                fmt.Fprintln(w, "The file contains more than a single value        [Error]")
                        }

                }

        }
}

func main() {
        //First we create the variable r as NewRoute
        r := mux.NewRouter()

        //The we create the routes for our APIs
        r.HandleFunc("/api/encrypt", HandleGetMethod_encrypt).Methods(http.MethodGet)
        r.HandleFunc("/api/decrypt", HandleGetMethod_decrypt).Methods(http.MethodGet)

        //I create a get method to show the information in the file:
        r.HandleFunc("/api/show_jsonfile", HandleGetMethod_show_jsonfile).Methods(http.MethodGet)

        //The we set the values for our server
        srv := http.Server{
                Addr:    ":8081",
                Handler: r,
        }

        //Print in the console that the server is listening and serve
        log.Println("Listening and Serve...")
        //The we start to listening
        srv.ListenAndServe()

}

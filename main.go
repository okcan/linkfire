package main



import (

        "net/http"

        "encoding/json"

        "fmt"

)



func main() {

        http.HandleFunc("/", get_employee_json)

        fmt.Printf("API endpoint -> http://localhost:8080/employee")

        http.ListenAndServe(":8080", nil)

}



type Employee struct {

        Timestamp       string `json:"timestamp"`

        Hostname      string `json:"hostname"`

        Engine        string `json:"engine"`

        Visitor  string `json:"visitor ,omitempty"`

}



var employees = map[string]Employee{

        "0000000000": Employee{Timestamp: "2022", Hostname: "test", Engine: "docker", Visitor: "192.168.1.1"},

}



func get_employees() []Employee {

        values := make([]Employee, len(employees))

        i := 0

        for _, emp := range employees {

                values[i] = emp

                i++

        }

        return values

}



func get_employee_json(w http.ResponseWriter, r *http.Request) {

        emps := get_employees()

        data, err := json.Marshal(emps)

        if err != nil {

                panic(err)

        }

        w.Header().Add("Content-Type", "application/json; charset=utf-8")

        w.Write(data)

}

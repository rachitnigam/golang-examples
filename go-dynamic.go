//http://starp-germany.de/blog/dynamic-in-go/

package main

import (
  "fmt"
	//"reflect" not needed

)
//My Log Types
type ErrorLog struct { message string }
func (errlog *ErrorLog) Log() bool {

   fmt.Println("Logged by ErrorLog")
   fmt.Println("Logged message:" + errlog.message)
   return true
}
type WarningLog struct { message string }
func (warlog *WarningLog) Log() bool {

   fmt.Println("Logged by WarningLog")
   fmt.Println("Logged message:" + warlog.message)
   return true
}
//My interface for logging Types
type ILogging interface {

	Log() bool
}
/*
My function Type
@arg: ...interface{}
@return: interface{}

*/
type function func(params ... interface{}) interface{}

func main() {
	
	/*
	My Addition function
	@desc: Added all numbers types and return the summation
	@arg: ...interface{}
	@return: interface{}

	*/
	addition := func (params ... interface{}) interface{} {
		
	
		var sum int = 0;

		for _,par := range params {

			
			switch i := par.(type) {

					case int:
						sum += i
					case float32:
						sum += int(i)
					case float64:
						sum += int(i)
					default:
						fmt.Println("Unknown Type for addition")
				}
			
		}

		return sum
		
	}
	/*
	My Log function
	@desc: Logged all different kinds of logs
	@arg: ...interface{}
	@return: interface{}

	*/
	log := func (params ... interface{}) interface{} {
		
		
		for _, par := range params {

			//Log switch
			switch i := par.(type) {

					case *ErrorLog:
						//Type assertion to ILogging
						 if v,ok := par.(ILogging); ok {

						 	v.Log()
						 }
					case *WarningLog:
						//Type assertion to ILogging
						if v,ok := par.(ILogging); ok {

						 	v.Log()
						 }
					default:
						fmt.Println("Unknown Log Type: ",i)
						
				}
			
		}

		return true
		
	}

	//Function call by name in Golang
	funcs := map[string]function{}
	funcs["Addition"] = addition
	funcs["Log"] = log

	fmt.Println(funcs); //Prints map[Addition:0x40130e Log:0x401566]

	sum := funcs["Addition"](2,7.0,78.47,55,74)
	fmt.Println(sum) //prints 216

	//Create log types
	error, warning := ErrorLog{ "ohjee!!" }, WarningLog{ "good advice ;)"}
	funcs["Log"](&error,&warning)
	/*
        OUTPUT:
	Logged by ErrorLog
	Logged message:ohjee!!
	Logged by WarningLog
	Logged message:good advice ;)
	*/



}

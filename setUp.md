RUN 
1. go mod init github.com/repo-name
2. go mod tidy
3. go run main.go

SETUP 
1. go mod init MODULE_NAME

2. go get -u gorm.io/gorm
3. go get -u gorm.io/driver/mysql
4. go get -u github.com/gin-gonic/gin

5. create a .env file in the root directory
6. go get github.com/joho/godotenv // package to lead env variables
7. go get golang.org/x/time/rate // package to handle the rate limit
   Info: exported model name always start with a Capital alphabet
8. go get -u github.com/golang-jwt/jwt/v5

   Info: JSON response {
    c.JSON: This simply sets the response status code and JSON payload and then continues executing the next handler, if there is one. It does not stop further middleware or handlers from running, so it’s useful if you just want to log an error but not prevent the request from completing.

    c.AbortWithStatusJSON: This is typically preferred for error handling in middleware because it sends the response with a JSON payload and status code, then stops any further handlers from executing. This is useful for terminating the request when an error occurs, so that no further processing happens.
   }

KnowMore: {
cannot use "lambaharsh01@gmail.com" (untyped string constant) as *string value in struct literal

This error indicates that you are trying to assign an untyped string constant (like "lambaharsh01@gmail.com") to a field that expects a *string (pointer to a string) rather than a plain string. To fix this, you need to pass a pointer to the string.

Using the & operator directly
email := "lambaharsh01@gmail.com"
yourStruct := YourStruct{
    Email: &email,
}

}





use & for structs where you want to pass the pointers when the actial struct is large or too large

Maps: Use otpLimitUpdation directly.
Structs: Prefer &otpLimitUpdation if it is a struct instance.



POINTERS AND DOUBLE POINTERS 

func A(c *gin.Ccontext){
   //here put c is already a pointer to the gin context

   // now to call function B i shoud go
   use B(c) to call func B and not use B(&c)

   with &c we would be teling B to kae a pointer to c here which is alread a pointer to gin context

   if using B(&c)
   we might have modify the func B and convert it like //  B(c **gin.Context)

}

func B(c *gin.Context){
   //here im passing another function which takes c to be the pointer to gin.Context
}



ShouldBindJSON && BindJSON
Use BindJSON when you want Gin to automatically handle errors and respond with a 400 Bad Request if the binding fails.
Use ShouldBindJSON when you want to handle errors yourself and provide a custom response to the client.



Use var colors []databaseSchema.BusColors when you just want to declare the variable and don't need to initialize it immediately.

Use colors := []databaseSchema.BusColors{} if you want to make it explicitly non-nil at the start (though unnecessary in most cases).


using &[]schema create is nil datatype which is not very much usable to beign with if need to be moulded, modified and transformed using algorythins
also dont use pointer type whe working with loops
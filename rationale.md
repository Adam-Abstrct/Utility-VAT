# Adam Johnson - Coding Exercise
Attached is the coding exercise, completed to the best of my (current) abilities. 
I wanted to demonstrate how I approached the project and breifly explain my methods.
Although there are probably numeorus ways to improve on this. I feel that for the most part the code works as expected.

* I wanted to ensure that validation of the inputs was taken into account. For simplicity and brevity, I combined one function to take care of this. Simply returning errors where needed or a bool that is true when valid. I wanted to make sure the user could only enter vat types and a correct date format.

* I've tried to be conscious to error handle where possible, ensuring that I give clear feedback to the user what is required of them. 

* I wanted to make my response robust and predictable, so I built structs that would mirror the JSON response, also making use of being able to embed structs. I used the struct tags feature to map the JSON correctly. One thing I struggled (with go) was how to handle nil values. I ended up using pointers, and then handling a nil value as the current date, later within the function. I feel there may have been a better way to approach this. but did not want to spend too much time on this feature.

* Making use of functions, I extracted the HTTP functionality into it's own. mapping the JSON to structs, and delivering errors if needed. Taking care to close the Body to free up resource.

* When comparing the Dates, Becuase i've always done this using unix time. I converted my dates to unix and did a simple < or > to find if the current date was within range. I became aware (after) that the time packge hasbefore and after methods, which would have made things easier.

* I performed a simple if/else check to see if the user input was either reduced or standard. Because I Validated thsi earlier. These could only ever be the two options

* returned to user using fmt.Println()!

### Known Issues


## How to build on this
## Application - 'math api'

Implement a web service (preferably in Go, Python, Ruby or Java; extra effort to do that in Go
will be recognised; using a framework or not):

- /min - given list of numbers and a quantifier (how many) provides min number(s)
- /max - given list of numbers and a quantifier (how many) provides max number(s)
- /avg - given list of numbers calculates their average
- /median - given list of numbers calculates their median
- /percentile - given list of numbers and quantifier 'q', compute the qth percentile of the list elements

No need to be concerned with resources, we're assuming there's plenty enough of memory, etc.

---

## Improvements

- Inadequate error handling/validation (for example, quantifier min/max negative values)
- Passing values in query params is limited due to the maximum header size limitation. 
- Validation is subjective, for example, if the min/max quantifier is greater than the numbers array, it is just set to the length of the numbers array
- We can introduce server side validatons for API call.if anyone puts special characters in query string, which might break the system

## Final

For me I would say this is my first GO program. Well Computer Language is just nothing but representation of syntax. What matters is the logic you use to create it. But I like the GO language eventually you can do many research stuff.

For the test,
    just run the project and hit http://localhost:8080/avg?numbers=70,30,80,50,90&q=1 like wise url you can get the output

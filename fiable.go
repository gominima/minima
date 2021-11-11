package fiable

import (
	"fmt"
	"log"

	"net/http"

	"time"
)

      type fiable struct {
	server       *http.Server
	started      bool
	Timeout      time.Duration
	router        *router
	
	properties   map[string]interface{}
      }
      
      func New() *fiable { 
	  var router *router = Router()
	  var fiable *fiable = &fiable{router: router,}    
	  return fiable
      }


      func (e *fiable) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	hijack, ok := res.(http.Hijacker)
	
	if !ok {
	 fmt.Errorf("Error occured at 32")
	}
	conn, buffrw, err := hijack.Hijack()
	if err != nil {
		log.Fatal(err)
	}
	request := request(req, &e.properties)
	response := response(res, req, buffrw, conn, &e.properties)
            
	for _, routes := range e.router.routes {
	 path := request.GetPathURl()
	  if routes.url != path {
	   fmt.Fprint(res, "Not found" )
	   fmt.Print(path, routes.url)
	    
	  }
	  fmt.Print(routes)
		routes.handler(request, response)
		
	  
	}

	http.NotFound(res, req)
       
         }
        

      func (f*fiable) Listen(addr string) error {
	server := &http.Server{Addr: addr}
	 if f.started {
		fmt.Errorf("Server is already running", f)
	 }
	 f.server = server
	 f.started = true
	 return f.server.ListenAndServe()
      }
       

      func (f *fiable) Get(path string, handler handler) {
	f.router.addUrl("GET", true, path, handler)
      }




      

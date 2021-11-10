package fiable

import (
	"fmt"

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
	 
	}
	conn, buffrw, err := hijack.Hijack()
	if err != nil {

	}
	request := Request(req, &e.properties)
	response := Response(res, req, buffrw, conn, &e.properties)
            
	for _, routes := range e.router.routes {
	 path := request.GetPathURl()
	  if routes.url != path {
	   fmt.Fprint(res, "Not found" )
	    
	  }
	  if path == routes.url {
           routes.middleware(response, request)
	  }
	}

	
       
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





      

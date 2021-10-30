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
	  var fiable *fiable = &fiable{}    
	  return fiable
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





      

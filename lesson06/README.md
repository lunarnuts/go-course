### Task 1:
#### Web server handling any request to it’s address and returning JSON like:
	{
  		"host": "127.0.0.1:8080",
  		"user_agent": "curl/7.67.0",
 	 	"request_uri": "/anything/you?want",
  		"headers": {
    			"Accept": ["*/*"],
    			"User-Agent": ["curl/7.67.0"]
 		}
	}
	All data must be taken from request struct.
### Task 2:
	TCP Server (S) listen for a connections on a port
	TCP Client (C) connects to a TCP server on a port 
	C reads STDIN for any input. On hit enter (‘\n’) C sends input it got to a server
	S reads input (split by ‘\n’) and checks if it’s an int, returns input multiplied by 2
	If it’s not an integer, return uppercased input string.
	C shows response from S to STDOUT and waits for another input from STDIN.
	C exits by input `exit`
	S exits by ctrl+C
### Task 3:
	Web server, responds to “/” (404 on any other request)
	Support “POST” and “GET” methods
	On GET request returns HTML page, containing FORM with name and address fields + submit button. 
	Also there is a placeholder to display token on this page.
	On POST request reads name+address from body and creates a token as “name:address”. 
	Then saves this token to Cookies
	Page reloads, token value displayed on page(read it from cookie)
	Note: you can use simple HTML template for this task: https://jsfiddle.net/7vhun1tc/ or create your own, it’s up to you. This task (as course as well) is about golang, so it’s cool to have some practice in HTML/JS, but it’s not a requirement.

after setting up loadbalancer and servers,
in client you can send get ot put request

## GET
if cached previously, it will return the object from loadbalancer
if not, it will request to appropriate server

## PUT
if cache value is true, then it will cache it on loadbalance
whatever it's true or not, loadbalancer will request to the
appropriate server


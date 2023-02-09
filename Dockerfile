FROM golang:1.19   

RUN mkdir /app     
 
WORKDIR /app    

ADD main.go /app         

RUN go build -o main ./main.go 

CMD /app/main      
 

# Blockchain

- User create book and books data store into Blocks.
- Blocks store data, position, Timestamp, hash, previousHash.
- Each blocks previous Hash is linked to the one before it and that create blockchain.

## Create files and folders

`mkdir project`

`cd project`

`go mod tidy`

`go run main.go`

## port 

`http://localhost:8080`

you can run with this command. 


## API Reference

#### Create data 

```http
  http://localhost:8080/new
```

#### POST Blocks Data

```http
  http://localhost:8080
```

#### GET Blocks Data

```http
  http://localhost:8080
```

## Architecture 
<img width="1182" alt="Screenshot 2024-12-30 at 7 48 26 PM" src="https://github.com/user-attachments/assets/c197604b-427c-4130-9d38-454205c794a3" />


## Demo
<img width="637" alt="Screenshot 2024-12-30 at 7 38 46 PM" src="https://github.com/user-attachments/assets/c2c30bc1-6d0b-48ca-9335-368f18327a12" />

<img width="1102" alt="Screenshot 2024-12-30 at 7 40 39 PM" src="https://github.com/user-attachments/assets/9745ff0f-51c6-422b-99e7-a0ca13657c66" />

<img width="625" alt="Screenshot 2024-12-30 at 7 43 45 PM" src="https://github.com/user-attachments/assets/0627eceb-24db-443d-ba97-4225a13992af" />







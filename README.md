
# MySQL Sharding Example

This project demonstrates how to implement **manual sharding** in MySQL using **hash-based sharding** logic in a Go application. The example shows how to split a large dataset across multiple MySQL servers to improve scalability and performance.

---

## 📋 Prerequisites  
Before running this example, ensure you have the following installed:  

- **Go** (Golang)  
- **MySQL** (set up at least 3 MySQL servers as shards)  
- **MySQL Driver for Go** (`github.com/go-sql-driver/mysql`)  

---

## 📂 Project Structure  

```
├── main.go         # Main application file containing sharding logic
└── README.md       # Project documentation
```

---

## ⚙️ Setup  

1. **Create the Same Table on All Shards:**  
   Run this SQL script on each MySQL server (Shard 0, Shard 1, Shard 2):  

   ```sql
   CREATE TABLE users (
     user_id INT PRIMARY KEY,
     name VARCHAR(100),
     email VARCHAR(100)
   );
   ```

2. **Configure Database Connections:**  
   Update your Go code with the correct **DSN (Data Source Name)** for each MySQL server. Example:  

   ```go
   dbs[0] = connectToDB("user:password@tcp(shard0_host:3306)/dbname")
   dbs[1] = connectToDB("user:password@tcp(shard1_host:3306)/dbname")
   dbs[2] = connectToDB("user:password@tcp(shard2_host:3306)/dbname")
   ```

3. **Install Dependencies:**  
   Use the following command to install the MySQL driver for Go:  

   ```sh
   go get -u github.com/go-sql-driver/mysql
   ```

4. **Run the Application:**  

   ```sh
   go run main.go
   ```

---

## 🚀 How It Works  

- **Sharding Strategy:** Hash-based sharding using `user_id % 3`.  
- **Data Distribution:** Each shard holds a portion of the `users` table based on the hash of `user_id`.  
- **Insertion:** The application determines the correct shard for each user and inserts the data.  
- **Querying:** Queries are routed to the appropriate shard based on the `user_id`.  

### Example:  
| User ID | Shard |  
|---------|-------|  
| 1001    | 1     |  
| 1002    | 2     |  
| 1003    | 0     |  

---

## 🛠 Functions  

- **`insertUser(userID, name, email)`**: Inserts a user into the correct shard.  
- **`getUser(userID)`**: Fetches a user from the appropriate shard based on `user_id`.  

---

## ⚠️ Challenges  

1. **Cross-Shard Queries**: Joins across shards are not supported.  
2. **Rebalancing Shards**: Adding new shards requires rehashing and redistributing data.  
3. **Consistency**: Maintaining consistency across shards can be difficult.  

---

## 📚 Learn More  

- [MySQL Documentation](https://dev.mysql.com/doc/)  
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql)  
- [Consistent Hashing Explained](https://en.wikipedia.org/wiki/Consistent_hashing)  

---

## ✨ Future Improvements  

- Add **Range-Based Sharding**.  
- Implement **Consistent Hashing** to reduce data rebalancing.  
- Build a web-based admin panel for monitoring shards.  

---

## 📝 License  

This project is licensed under the **MIT License**.

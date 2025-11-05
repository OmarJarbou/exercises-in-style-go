# Persistent Tables Style: Term Frequency

This chapter implements the **term frequency task** using the **Persistent Tables style**.  
The main idea of this style is to **persist structured entities** (documents, words, characters) into a relational database, so that queries can reuse the stored data instead of reparsing raw input every time.

---

## Steps / Implementation

1. **Database Schema**  
   - Create the schema by adding the required migration files.  
   - Apply migrations using [Goose](https://github.com/pressly/goose) to create the tables in the database.

2. **SQL Queries and Code Generation**  
   - Write the SQL queries required for handling documents, words, and characters.  
   - Use [sqlc](https://github.com/kyleconroy/sqlc) to generate strongly typed Go code for these queries.

3. **File Path Input**  
   - Extract the file path argument passed from the command line.

4. **Database Connection**  
   - Open a connection to the database using the generated code and Go’s standard SQL library.

5. **Stop Words**  
   - Load stop words from the given file, normalize them, and extend them with printable ASCII characters.

6. **Check Document Persistence**  
   - Query the database to see if the document already exists by its name.

7. **If Document Exists**  
   - Apply the `CountWordsFrequencies` query on the stored data.  
   - Print the results directly without reparsing the file.

8. **If Document Does Not Exist**  
   - Insert the document into the `documents` table.  
   - Extract all **alphanumeric characters** and insert them into the `characters` table.  
   - Extract all **non-stop words** and insert them into the `words` table.  
   - Then run the same query from step 7 to calculate frequencies and print results.

---

## Why Persistent Tables?

The purpose of this style is to **avoid reparsing raw input repeatedly**.  
Instead, once the data is structured and persisted, you can run efficient relational queries to get results.

Benefits:
- Saves time on repeated analyses of the same documents.  
- Database indexing can make word frequency queries and lookups much faster.  
- Supports richer queries beyond term frequency (e.g., multi-document analysis).

---

### Constraints

- **Persistance**:
  The data exists even after the program terminates, and is used for other executions, and can be used for other programs (e.g., here we can get benifit of characters table and get the count of each character).

- **Schema Design**:  
  - Must carefully design the schema (documents, words, characters) to avoid redundancy and ensure fast queries.   
  - The concrete data is modeled as having components of several domains, establishing relationships between the application’s data and the domains identified.

- **Query-Based**:  
  The problem is solved by issuing queries over the data.

### **Note:**
- **Persistence Overhead**:  
  Initial ingestion of documents is more expensive and slower than just parsing, but it pays off in reuse. 

---

## Running

```bash
# Apply migrations
goose up

# Run program with a file
go run . ../sample.txt

# Things Style: Term Frequency

This chapter implements the **term frequency task** using the **Things style**.  
The main idea of this style is to **subdivide the computational task into logical objects ("things")**, each with its own data and methods, and enforce encapsulation by only allowing access to data through methods on the struct.

---

## Steps / Implementation

1. **Controller Initialization**  
   - Create an instance of `WordFrequencyController` in `main`.  
   - Initialize all required "things" (objects/structs) and resources through `initializeWordFrequencyController`.  
   - Run the program starting from the controller.

2. **File Reading**  
   - Read the raw contents of the input file by the `DataStorageManager` instance.

3. **Word Extraction**  
   - Filter the content to extract words containing only **alphanumeric characters** also by the `DataStorageManager` instance.

4. **Stop Words**  
   - Read stop words from the given file.  
   - Normalize them to lowercase.  
   - Extend the list with all **printable ASCII characters**.  
   - Store stop words in a map for faster lookup and removal.
   - All done by the `StopWordsManager` instance.

5. **Filtering**  
   - Remove stop words from the extracted words list also by the `StopWordsManager` instance.

6. **Word Frequencies**  
   - From the filtered list, calculate the frequency of each word by the `WordFrequencyManager` instance.

7. **Sorting**  
   - Sort the words by their frequencies in descending order also by the `WordFrequencyManager` instance.

8. **Output**  
   - Print the words with their frequencies by the `WordFrequencyContoller` instance itself.

**Different implementations for StructInfo's Info function:**
   - I have made StructInfo interface that has Info function that takes some instance and return info about it.
   - I have made default implementation for Info under the interface.
   - All `DataStorageManager`, `StopWordsManager`, and `WordFrequencyManager` instances implements the interface by implementing its Info method.
   - Any call to the method with instance from the one of the mentioned structs as a parameter AND AS RECIEVER (i.e. for example (dsm *DataStorageManager) is a reciever for a function) will result in calling the implemented Info method, not the default one.
   - While `WordFrequencyContoller`'s instance is calling the default one (no implementation for Info in WordFrequencyContoller), by passing itself to the function.


---

## Why Things Style?

The **Things style** organizes the program as a set of "things" (inctances of structs) that group together data and related behavior.  

- Each **thing** (inctance of a struct) encapsulates its data.  
- Operations on the data are performed only through **methods bound to the struct**.  
- To achieve that isolation, we put each struct in a separate package; because in one package all fields can be accessed, but in seperate packages you can see clearly the difference between private fields/functions (which stars by small letter) and public ones (which starts by capital letter).
- This prevents direct access to internal state and enforces **logical grouping** of responsibilities.  

For example:
- A `DataStorageManager` is responsible for reading and holding file content.  
- A `StopWordsManager` handles stop words logic.  
- A `WordFreqManager` is responsible for calculating frequencies.  
- A `WordFrequencyController` coordinates the interaction of all instances.

---

## Constraints

- **Encapsulation**:  
  Data must not be exposed directly; access happens only through methods tied to the struct.  

- **Logical Grouping**:  
  Methods that operate on a given data structure should live on the struct that contains that data.  

- **Separation of Concerns**:  
  Each "thing" should focus on one responsibility (storage, stop words, frequency counting, coordination, etc.).

---

## Running

```bash
go run .

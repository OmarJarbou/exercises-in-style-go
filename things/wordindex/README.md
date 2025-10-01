# Things Style: Word Index

This chapter implements the **word index task** using the **Things style**.  
The main idea of this style is to **subdivide the computational task into logical objects ("things")**, each with its own data and methods, and enforce encapsulation by only allowing access to data through methods on the struct.

---

## Steps / Implementation

1. Creating an instance of `WordIndexController` in main and initializing all required (things/objects)'s resources throughout *InitializeWordIndexController*, and then run the program with the required parameters from command line.

**IN WordIndexController's run method:**

2. Read the real file's content into lines.

3. Filter lines characters to be only alphabet, and seperate line's words with spaces.

4. Get filtered lines from `DataStorageManager`'s instance.

5. Extract words and their indexes using `WordindexManager` instance's  *ExtractWordsAndIndexes* function.

6. Sort list of words and indexes alphabetically based on words, also by the `WordindexManager` instance.

7. Get list of words and indexes from the `WordindexManager` instance.

8. Print words and their corresponding page indexes without duplicates if the word occured less than 100 times.

---

## Why Things Style?

### Constraints

**Explained in README.md of the *things* implementation of term frequency task.**

---

## Running

```bash
go run . 'path_to_project' 'path_to_file' 'lines_per_page (must be integer)'

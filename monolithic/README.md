# Term Frequency Task (Monolithic Style)

This project implements the **term frequency** task in **monolithic style**, by following these steps:

## Steps

1. Open the stop words file, read its content, split words, and store them in a slice with all words in lower case.  
2. Get a list of all printable ASCII characters and add them to the stop words slice.  
3. Open the real file to read its content and start iterating over all characters in this file.  
4. Extract separate words which consist only of alphanumaric characters.  
5. Store only non-stop words in the list and increment frequency if they already exist.  
6. After processing is complete, order extracted words **(in a slice, not a map because maps are unordered)** based on their frequency.  
7. Print top 25 word:freq.  

---

## Notes on Monolithic Style

- No external ready functions are used other than some simple ones like `strings.ToLower()`.  
- No (or very little) use of libraries.  
- No functions or structures are used to organize the code.  
- No subdividing — everything happens inside one flow.  

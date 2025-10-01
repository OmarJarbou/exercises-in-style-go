# Term Frequency Task (Pipeline Style)

This project implements the **term frequency** task in **pipeline style**, by following these steps:

## Steps

1. Open the stop words file, read its content, split words, and store them in a slice.  
2. Normalize extracted stop words to all be in lower case.  
3. Get a list of all printable ASCII characters and add them to the stop words slice.  
4. Open the real file to read its content.  
5. Start iterating over all characters in this file and extract separate words (words consist only of alphabet characters and are non-stop words). Add them to the list and increment frequency if they already exist.  
6. After processing is complete, order extracted words based on their frequencies.  
7. Print the results.  

**Note:**
- For functions that expects two arguments, we use currying to transform the function to sequence of higher order functions each with one single argument

---

## Constraints

- The computational task is subdivided into a **pipeline of functions**. 
- No shared state between functions. 
- Each function takes **one input** and produces **one output**, which is the input for the next function (composing functions one after the other, in pipeline, as a faithful reproduction of mathematical function composition f ◦ g).  
- It is a **hierarchy of functions**, where each function serves the next.  
- The **first function** to be called and processed is the **last one on the left**, and results will come one after the other and serve the ones on the right.  


---

## Running

```bash
go run .

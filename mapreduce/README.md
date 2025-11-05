# MapReduce Style: Term Frequency

This chapter implements the **term frequency task** using the **MapReduce style**.  
The key idea of MapReduce is to divide the input into independent chunks, process them in parallel, and then combine the results.

---

## Steps

1. **Partition**  
   The `Partition` function opens the input file and divides its content into chunks by reading a fixed number of lines at a time.  
   It returns a list of chunks.

2. **Stop Words**  
   We retrieve the list of stop words via the `GetStopWords` function of the `StopWordsManager` struct.  
   This list is normalized and extended with printable ASCII characters, all done by the same function.

3. **Map Phase**  
   - The `Map` function takes the `splitWords` function, the chunks, and the stop words list.  
   - It applies `splitWords` to each chunk **concurrently** (each in its own goroutine).  
   - Each goroutine sends its result (a list of words) over its channel.  
   - The `Map` function collects these results into `splits`.

4. **Reduce Phase**  
   - The `Reduce` function takes the `countWords` function and the list of splits.  
   - It applies `countWords` to each split **concurrently** (again using goroutines).  
   - Each goroutine sends back a word-frequency map.  
   - The `Reduce` function combines these maps by summing the counts of identical words.  
   - Finally, the results are converted into a slice of `WordFreq` struct to be able to sort it by frequency.

5. **Output**  
   The top 25 words and their frequencies are printed.

---

## Why MapReduce?

This style follows a **divide-and-conquer approach**:
- Input data is split into smaller chunks.  
- Each chunk is processed independently, often in parallel.  
- Results are merged to form global knowledge.  

For the term frequency task, this is very effective:  
we can count words on smaller portions of the input (e.g., 200 lines each) and then combine the counts efficiently.

**BUT** this approach does not work on all tasks, task must has:
- **Divisible input**: The input can be split into independent chunks (e.g., text files split by lines, images split by pixels).

- **Independent map operations**: Each chunk can be processed without needing global context (phases does not depend on each other). (e.g., counting words).

- **Associative + commutative reduce**: The partial results can be combined in any order and still give the correct result (like sums, counts, min/max,...).

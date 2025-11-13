# Quarantine Style: Term Frequency

This chapter implements the **term frequency task** using the **Quarantine style**.  
The Quarantine style isolates ("quarantines") side effects like IO by wrapping them inside higher-order functions. Functions are composed and passed around as pure values until explicitly executed.

---

## Steps / Implementation

1. **Create TFQuarantine**  
   - Instantiate the `TFQuarantine` struct, which is responsible for:  
     - Storing all the functions needed to complete the task.  
     - Executing them in the order they are bound.  
   - Functions that include I/O (e.g., file reading, logging) are curried so they return a function instead of executing immediately, so the wrapping higher order function is considered pure.

2. **Bind functions in order**  
   - Bind functions step by step so that:  
     - The output of one function (weather a value, or a function to be called "which is an impure function that contains IO") becomes the input of the next.  
     - The pipeline remains pure until explicitly run.

3. **Execute**  
   - The `TFQuarantine` struct executes all bound functions in the correct order.  
   - Only at this stage do the actual side effects occur (e.g., reading files).

4. **Output**  
   - After execution, print the final result (top 25 word frequencies).

---

## Why Quarantine?

The purpose of this style is to **isolate side effects** and make the core logic **pure and composable**.  

- All functions are treated as pure unless explicitly executed.  
- Side effects (I/O, logging, reading files) are delayed and only happen when calling the final returned function.  
- This ensures the code is more predictable and testable, since function composition is deterministic until execution.

---

### Constraints

- **Currying of side-effect functions**:  
  Any function that performs I/O (like reading a file or terminating the program) must return a *function to be executed later* instead of doing it immediately; so that core program functions have no side effects of any kind, and all IO actions contained in computation sequences that are clearly separated from the pure functions.

- **Purity of functions**:  
  Pure functions should not rely on external mutable state like IO; they must only depend on their input.

- **Explicit execution**:  
  The actual “dirty” side effect happens **only** when you call the returned function in main. Until then, you’re just passing functions around like values.

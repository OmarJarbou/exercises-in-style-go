# Actor Style: Term Frequency

This chapter implements the **term frequency task** using the **Actor style**.
The key idea of the Actor model is to represent each component as an independent actor with its own state and communication channel. Actors interact only by sending and receiving messages over these channels.

---

## Steps

### Initialization

- All actor instances (structs) are created in main.
- Each actor has a special initialization function that:
    - Creates a channel (or queue) for incoming messages.
    - Starts its run method in a separate goroutine.
    - prepares any internal resources (e.g., stop words, file's text,...).

- The run method continuously listens for messages, processes them using a dispatch method, and only stops when it receives a die message.

- After main sends its messages, it thread blocks until all four actors finish (after receiving their die message).

---

### Flow of Messages
The system operates entirely by exchanging messages between actors:

1. `main` → `WordFreqManager` and `StopWordsManager`: send `init` messages to load the required resources (file content and stop words).

2. `main` → `WordFreqController`: send `run` message so it can store references to other actors and kick off the workflow by sending send_word_freqs to DataStorageManager.

3. `DataStorageManager` (on `send_word_freqs`) → extracts all alphanumeric words from the file content, then sends a `filter` message with the words to `StopWordsManager`.

4. `StopWordsManager` (on `filter`) → filters out stop words, then sends a `words` message to `WordFreqManager` along with filtered words.

5. `WordFreqManager` (on `words`) → computes word frequencies, sorts them, and sends a `top25` message to `WordFreqController`.

6. `WordFreqController` (on `top25`) → prints the results, then sends `die` messages **to all actors (including itself)**.

7. All actors stop, and the main thread terminates.

---

## Why Actors?

### Constraints

- The larger problem is decomposed into things/objects/instances of structs that make sense for the problem domain.
- Each instance has a queue/channel meant for other instances to place messages in it.
- Each instance is a capsule of data that exposes only its ability to receive messages via the queue.
- Each instance has its own goroutine(lightwieght thread) independent of the others. This gives the actors the ability to run independently and concurrently.
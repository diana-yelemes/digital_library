//public/javascript/app.js
function filterBooks(status) {
    // Determine the route based on the status
    let route;
    switch (status) {
        case 'to-read':
            route = '/books/to-be-read-books';
            break;
        case 'currently-reading':
            route = '/books/currently-reading-books';
            break;
        case 'read':
            route = '/books/read-books';
            break;
        case 'did-not-finish':
            route = '/books/did-not-finish-books';
            break;
        default:
            console.error('Invalid status:', status);
            return;
    }

    // Make a GET request to the route
    fetch(route)
        .then(response => response.json())
        .then(books => {
            // Get the bookList div
            var bookList = document.getElementById("bookList");

            // Clear the current book list
            bookList.innerHTML = "";

            // Add each book in the filtered list to the bookList div
            books.forEach(book => {
                var bookDiv = document.createElement("div");
                bookDiv.textContent = book.title + " by " + book.author;
                bookList.appendChild(bookDiv);
            });
        })
        .catch(error => console.error('Error:', error));
}

function fetchBooks() {
    // Make a GET request to the server to fetch the list of books
    fetch('/books')
        .then(response => response.json())
        .then(books => {
            // Get the bookListContainer div
            var bookListContainer = document.getElementById("bookListContainer");

            // Clear the current book list
            bookListContainer.innerHTML = "";

            // Add each book to the bookListContainer div
            books.forEach(book => {
                var bookElement = document.createElement('div');
                bookElement.innerHTML = `
                    <p onclick="showBookDetails('${book.ID}')">${book.Title} by ${book.Author}: ${book.Status}</p>
                    <button class="mark-as-button" onclick="displayStatusOptions(${book.ID})">Mark As</button>
                    <div id="status-options-${book.ID}" class="status-options" style="display: none;">
                        <button onclick="markBookStatus(${book.ID}, 'to-read')">To Read</button>
                        <button onclick="markBookStatus(${book.ID}, 'currently-reading')">Currently Reading</button>
                        <button onclick="markBookStatus(${book.ID}, 'read')">Read</button>
                        <button onclick="markBookStatus(${book.ID}, 'did-not-finish')">Did Not Finish</button>
                    </div>
                `;
                bookListContainer.appendChild(bookElement);
            });
        })
        .catch(error => console.error('Error:', error));
}

function displayStatusOptions(bookID) {
    var statusOptions = document.getElementById(`status-options-${bookID}`);
    statusOptions.style.display = 'block';
}

function displayBooks(books) {
    const bookListContainer = document.getElementById('bookList');
    // Clear existing content
    bookListContainer.innerHTML = '';

    if (books.length === 0) {
        // Display a message if there are no books
        bookListContainer.innerHTML = '<p>No books available.</p>';
    } else {
        // Display each book in the list
        books.forEach(book => {
            const bookElement = document.createElement('div');
            bookElement.innerHTML = `<p>${book.Title}</p><p>${book.Author}</p><p>${book.Description}</p><p>${book.Status}</p>`;
            bookListContainer.appendChild(bookElement);
        });
    }
}

// Function to fetch book details by ID from the server
async function fetchBookDetails(bookID) {
    const response = await fetch(`/books/${bookID}`);
    if (!response.ok) {
      console.error('Failed to fetch book details');
      return null;
    }
    return await response.json();
  }
  
  // Function to display book details when a title is clicked
  async function showBookDetails(bookID) {
    // Fetch book details from the server
    const bookDetails = await fetchBookDetails(bookID);
  
    // Update the details container with book information
    if (bookDetails) {
      document.getElementById('bookTitle').textContent = bookDetails.title;
      document.getElementById('bookAuthor').textContent = `Author: ${bookDetails.author}`;
      document.getElementById('bookDescription').textContent = `Description: ${bookDetails.description}`;
      document.getElementById('bookStatus').textContent = `Status: ${bookDetails.status}`;
  
      // Show the details container
      document.getElementById('bookDetailsContainer').style.display = 'block';
    }
  }



function displayStatusOptions(bookID) {
    var statusOptions = document.getElementById(`status-options-${bookID}`);
    statusOptions.style.display = 'block';
}

function markBookStatus(status) {
    // Get the book ID from the hidden p element
    var bookID = document.getElementById("bookID").textContent;

    // Determine the route based on the status
    let route;
    switch (status) {
        case 'to read':
            route = `/books/${bookID}/to-be-read`;
            closeMarkAsModal(); 
            break;
        case 'currently reading':
            route = `/books/${bookID}/currently-reading`;
            closeMarkAsModal();
            break;
        case 'read':
            route = `/books/${bookID}/mark-as-read`;
            closeMarkAsModal(); 
            break;
        case 'did not finish':
            route = `/books/${bookID}/did-not-finish`;
            closeMarkAsModal();
            break;
        default:
            console.error('Invalid status:', status);
            closeMarkAsModal(); 
            return;
    }

    // Make a POST request to the route
    fetch(route, { method: 'POST' })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            // Handle the response here
        })
        .catch(error => console.error('Error:', error));
}
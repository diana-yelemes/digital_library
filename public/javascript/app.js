    // public/javascript/app.js
    document.addEventListener('DOMContentLoaded', function () {
          // Fetch the initial HTML content
    fetch('/')
    .then(response => {
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.text(); // Parse the HTML content
    })
    .then(htmlContent => {
        // Create a temporary container to parse the HTML
        const tempContainer = document.createElement('div');
        tempContainer.innerHTML = htmlContent;

        // Extract the book list container from the parsed HTML
        const bookListContainer = tempContainer.querySelector('#bookList');

        if (bookListContainer) {
            // Update the existing book list container with the received HTML content
            const currentBookListContainer = document.getElementById('bookList');
            currentBookListContainer.innerHTML = bookListContainer.innerHTML;

            // Add click event listeners to each book element
            currentBookListContainer.addEventListener('click', (event) => {
                const bookElement = event.target.closest('.book-list-container');
                if (bookElement) {
                    const bookID = bookElement.dataset.bookId;
                    showBookDetails(bookID);
                }
            });
        } else {
            console.error('Book list container not found in the received HTML.');
        }
    })
    .catch(error => console.error('Error:', error));



    const ftbr=document.getElementById("fltrtbr");
    ftbr.addEventListener("click", () => filterBooks('to-be-read'));

    const fread=document.getElementById("fltrread");
    fread.addEventListener("click", () => filterBooks('read'));

    const fcr=document.getElementById("fltrcr");
    fcr.addEventListener("click", () => filterBooks('currently-reading'));

    const fdnf=document.getElementById("fltrdnf");
    fdnf.addEventListener("click",() => filterBooks('did-not-finish'));

    const mcr=document.getElementById("cr");
    mcr.addEventListener("click",(event) => {
        const bookElement = event.target.closest('.book-list-container');
        if (bookElement) {
            const bookID = bookElement.dataset.bookId;
            updateBookStatus(bookID, 'currently-reading');
        }
    });

    const mread=document.getElementById("read");
    mread.addEventListener("click",() => {
        const bookID = this.dataset.bookId; 
        updateBookStatus(bookID, 'read');
    });

    const mdnf=document.getElementById("dnf");
    mdnf.addEventListener("click", () => {
        const bookID = this.dataset.bookId;
        updateBookStatus(bookID, 'did-not-finish');
    });

    const mtbr=document.getElementById("tbr");
    mtbr.addEventListener("click", () => {
        const bookID = this.dataset.bookId;
        updateBookStatus(bookID, 'to-read');
    });

    const dlt=document.getElementById("deletebtn");
    dlt.addEventListener("click", () => {
        const bookID = this.dataset.bookId;
        deleteBook(bookID);
    });


        function filterBooks(status) {
            var bookListContainer = document.getElementById("bookList");
            // Make a GET request to the server to fetch the list of books with the given status
            fetch(`/api/${status}`)
            .then(response => {
                if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(books => {
        
                // Clear the existing book list
                bookListContainer.innerHTML = '';
        
                // Add each book to the bookListContainer div
                books.forEach(book => {
                var bookElement = document.createElement('div');
                bookElement.innerHTML = `
                    <p>${book.title} by ${book.author}: ${book.status}</p>
                `;
        
                bookListContainer.appendChild(bookElement);
                });
            })
            .catch(error => console.error('Error:', error));
        }
        

            // Handle click events to display book details
            function showBookDetails(bookID) {
                fetch(`/api/books/${bookID}`)
                .then(response => response.json())
                .then(book => {
                    const bookDetailsContainer = document.getElementById('bookDetailsContainer');
                    const titleElement = document.getElementById('bookTitle');
                    const authorElement = document.getElementById('bookAuthor');
                    const descriptionElement = document.getElementById('bookDescription');
                    const statusElement = document.getElementById('bookStatus');
                    const idElement = document.getElementById('bookID');
            
                    titleElement.textContent = book.title;
                    authorElement.textContent = `Author: ${book.author}`;
                    descriptionElement.textContent = `Description: ${book.description}`;
                    statusElement.textContent = `Status: ${book.status}`;
                    idElement.textContent = book.id;
            
                    // Show the details container
                    bookDetailsContainer.style.display = 'block';
                })
                .catch(error => console.error('Error fetching book details:', error));
            }
            

        // Function to update book status
        async function updateBookStatus(bookID, status) {
        
            const response = await fetch(`/api/books/${bookID}/${status}`, { method: "POST" });
        
            if (response.ok) {
            // Update the UI or perform additional actions if needed
            console.log(`Book status updated to ${status}`);
            } else {
            console.error("Failed to update book status");
            }
        }

        function fetchBooks() {
            // Make a GET request to the server to fetch the list of books
            fetch('/')
                .then(response => response.json())
                .then(books => {

        
                    // Clear the current book list
                    bookListContainer.innerHTML = "";
        
                    // Add each book to the bookListContainer div
                    books.forEach(book => {
                        var bookElement = document.createElement('div');
                        bookElement.innerHTML = `
                            <p onclick="showBookDetails('${book.ID}')">${book.Title} by ${book.Author}: ${book.Status}</p>
                            <button class="mark-as-button" onclick="displayStatusOptions(${book.ID})">Mark As</button>
                            <div id="status-options-${book.ID}" class="status-options" style="display: none;">
                                <button onclick="markBookStatus(${book.ID}, 'to-be-read')">To Be Read</button>
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
        

        // Function to delete a book
        function deleteBook(bookID) {
        
            // Make a DELETE request to the server
            fetch(`/api/books/${bookID}`, {
                method: 'DELETE'
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(() => {
                // Refresh the book list
                fetchBooks();
            })
            .catch(error => console.error('Error:', error));
        }

        // Function to update the book list on the page
        function updateBookList(books) {
            const bookListContainer = document.getElementById('bookList');
            bookListContainer.innerHTML = '';

            if (books.length > 0) {
                books.forEach(book => {
                    const bookDiv = document.createElement('div');
                    bookDiv.innerHTML = `
                        <p onclick="showBookDetails('${book.ID}')" style="cursor: pointer;">
                            ${book.Title} by ${book.Author}: ${book.Status}
                        </p>
                        <button onclick="updateBookStatus('${book.ID}', 'currently-reading')">Currently Reading</button>
                        <button onclick="updateBookStatus('${book.ID}', 'read')">Mark as Read</button>
                        <button onclick="updateBookStatus('${book.ID}', 'did-not-finish')">Did Not Finish</button>
                        <button onclick="updateBookStatus('${book.ID}', 'to-be-read')">To Be Read</button>
                    `;

                    bookListContainer.appendChild(bookDiv);
                });
            } else {
                bookListContainer.innerHTML = '<p>No books found.</p>';
            }
        };
    })
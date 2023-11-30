    // public/javascript/app.js
    document.addEventListener('DOMContentLoaded', function () {
        fetchBooks();

    const ftbr=document.getElementById("fltrtbr");
    ftbr.addEventListener("click", () => filterBooks('to-be-read'));

    const fread=document.getElementById("fltrread");
    fread.addEventListener("click", () => filterBooks('read'));

    const fcr=document.getElementById("fltrcr");
    fcr.addEventListener("click", () => filterBooks('currently-reading'));

    const fdnf=document.getElementById("fltrdnf");
    fdnf.addEventListener("click",() => filterBooks('did-not-finish'));

    document.getElementById('bookList').addEventListener('click', (event) => {
        const clickedBookElement = event.target.closest('.book-list-container');
        if (!clickedBookElement) return;
    
        // Read book ID from the clicked book element
        const bookID = clickedBookElement.dataset.bookId;
    
        // Handle "Mark As" button clicks
        const markAsButton = event.target.closest('.mark-as-button');
        if (markAsButton) {
            const status = markAsButton.dataset.status;
            updateBookStatus(bookID, status);
        }
    
        // Handle "Delete" button clicks
        const deleteButton = event.target.closest('.delete-button');
        if (deleteButton) {
            deleteBook(bookID);
        }
    });
      


    function fetchBooks(){
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


    }
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
                // Don't need to parse response as JSON in this case
                return response.text();
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
                        const bookID = book.ID;

                        bookDiv.innerHTML = `
                            <p onclick="showBookDetails('${bookID}')" style="cursor: pointer;">
                                ${book.Title} by ${book.Author}: ${book.Status}
                            </p>
                            <button class="mark-as-button" data-bookid="${bookID}" data-status="currently-reading">Currently Reading</button>
                            <button class="mark-as-button" data-bookid="${bookID}" data-status="read">Mark as Read</button>
                            <button class="mark-as-button" data-bookid="${bookID}" data-status="did-not-finish">Did Not Finish</button>
                            <button class="mark-as-button" data-bookid="${bookID}" data-status="to-be-read">To Be Read</button>
                            <button class="delete-button" data-bookid="${bookID}">Delete</button>
                        `;

                        bookListContainer.appendChild(bookDiv);
                    });
                } else {
                    bookListContainer.innerHTML = '<p>No books found.</p>';
                }
            };



        const newForm = document.getElementById('new-form');

        newForm.addEventListener('submit', function (event) {
        event.preventDefault();

        // Create a FormData object from the form
        const formData = new FormData(newForm);

        // Make a POST request to the server
        fetch('/book', {
            method: 'POST',
            body: formData,
        })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            // Handle the response data, if needed
            console.log('Book added successfully:', data);
            
          
            window.location.href = '/books'; // Redirect to the book list page
        })
        .catch(error => console.error('Error:', error));
    });
    })
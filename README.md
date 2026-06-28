# Intro

My personal project to keep a simple inventory of my books. I have implemented different db clients, switching from an Excel sheet, SQL Server, SQL Lite and finally Notion. Notion is the latest addition since I already use it in my daily routine.

To get the book info, it uses the Google Books API, which requires a registered key (free).

The "Scanner" suffix comes from the idea to use a barcode scanner to facilitate the process of data entry. Hardware and Software coming together!

Also this is my first project using Go. It was a great little simple learning experience. Not a fan right now but maybe I'll turn into one at some point.

Ohara is a OnePiece reference. iykyk.

# Instructions

This console app assumes I'm at my bookcases which are already divided by category. That's why I added the batch property questions at the beginning, since these would remain the same for a big batch.

1. Run main.go by debugging in VSC with F5 or `go run main.go`
2. Enter the batch properties
    - Category = Fiction or Non-Fiction
    - Shelf Location = Shelf-bounded big categories
3. Scan the book ISBN, this will search for it in the Google Books API
4. It will ask is you've read the book or not
5. Book is successfully added!

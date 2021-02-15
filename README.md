# IITJ-INVENTORY

Inventory Management System 

TechStack - 
Backend : GO 
Frontend : React/Vue (Neeraj and Github) 

DataBase : Postgres/MySQL 

Steps to follow - 

1. Install GO on your system.
2. Learn basic syntax for GO.
3. Start Writing Some CODE #Challenge for Day 1

.................................................................................................................................................................................

# A custom signup Service for every website :D

1. Every Society login
2. Society Entry Database
3. CRUD for items (API)
4. Issue Item (Time consicious)
5. Request Item from a society
6. Budget (Optional)
7. Request for renewal/return
8. Admin Portal for every Society
...................................................................................................................................................................................


DataBase

Users will signIn using google and admins will sign in using email-id and passwords

Items ka Table 
{
    item society-key quantity identifier available 
}

Societies ka Table
{
    society email-id password
}

Users ka Table
{
    name email photo  
} 

Issued Items
{
    Date-of-issue due-date user-key purpose item-key approved denied
}

Inbound
{
    item-key quantity price society bill date-of-purchase date-of-inbound
}

defective
{
    item-key quantity date-of-destruction
}

.................................................................................................................................................................................

API Calls 

****FUTURE ME KABHI CART SYSTEM ADD KRSKTE HAI :)

SUPER ADMIN CHAHIE

1. Login-Flow 
    /login/admin?email=&password  
    /login/user ---> Redirect to google auth


2. Main-Flow
    /dashboard      -----> admin or user
    /logout
    
    user ->
        /item-requests
        /list-socities
        /search?q=&society_name=  ------->  list of items with that society
        /issue?q=item&purpose=&quantity=
        /cancel-order?order=
    
    admin ->

        /society-item-requests
        /approve-item?order=
        /decline-item?order=

        /all-items             --------------> lists all items of the club
                                                            Add new or add ka option har item ke bagal

        /add-new-item?item_id=          -----> isi me inbound wali table bhi bnni hai 
                                                                     item_id = NULL then naya item hai add krdo 
                                                                     agar NULL nhi hai then increase quantity
                                                                     inbound mtlb purchase of item with bill

        /delete-item?item_id=&quantity=          ---------> delete mtlb kharab wala manage krlen :D

        /item/history?item_id=

        /change-password


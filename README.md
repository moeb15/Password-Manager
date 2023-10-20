# Password Manager
For this project I am working on creating a password manager fullstack application, at the moment the backend is in a viable state.
So far the following has been completed,
* User registration/login and password management(create/read/update/delete saved passwords)
* When users register they include a master-key along with their account credentials, the key is used to encrypt any saved passwords
* Passwords are encrypted using AES-256, keys and passwords must be no more than 32 characters
* To retrieve saved passwords users must input the master-key
* User authentication is done using JWT, access and refresh tokens are provided when users login
* Passwords and user data are stored using MongoDB


The frontend is a work-in-progress, I'm developing it using JavaScript, React, and Tailwind CSS.
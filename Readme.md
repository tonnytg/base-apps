# Base Apps

This projects has two folders `Backend` and `Frontend`. Each folder contains essential files to run the application.

`Backend` use Go version 1.21.6 to copile and run some backend application.<br />
`Frontend` use Node LTS version 20^ to compile and run some frontend application, this base created using create-react-app without TypeScript only JavaScript.

### Why I create this base?

I lost a lot of time building the same base for each project, so I decided to create this base to save time and start project with same version of packages and same structure.


#### Frontend

Created using create-react-app without TypeScript, remove non essential files and add some packages to start project like icons and webvitals.

I put Material UI as UI library, but you can remove it and use another UI library if you want.

#### Backend

It is just Go no frameworks or external packages. `GOOPATH` is set to `/go` you can look at `Dockerfile` to understsand or update version.


### How I run it

I build a Image in localmachine MacOS without install nothing about Go or JavaScript, only Docker.


Create your own image:

`docker build -t <NAME>:<TAG> .`


Run your image:

`docker run -it -p 3000:3000 -p 8080:8080 <NAME>:<TAG>`
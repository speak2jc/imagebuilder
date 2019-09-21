# imagebuilder

This project demonstrates an operator which processes CRDs of type AppBuild
Each CR contains information about source, base image and target - it then overlays the application source on the base image, 
builds a Docker image and then pushes to the target.

TODO
- Create a dedicated go file which gets called from main controller
- Inside read all of the data from CR and perform a pseudo-copy from source into a temp location
- The temp location becomes the Docker context so copy over your Dockerfile here (for this example embed into image)
- Call the Kaniko "executor" (needs to be available on base image) with options to create and push the image

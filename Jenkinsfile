pipeline {
    agent any

    stages {
        stage('Say Hello') {
            steps {
                echo 'Hello world!'
                // Optionally, you can print an additional message to confirm execution
                echo 'This is the Say Hello stage.'
                sh 'ls -R /var/lib/jenkins/workspace/angry-embassies'
                echo 'ls -R /var/lib/jenkins/workspace/angry-embassies'
            }
        }
        stage('Build Docker Image') {
            steps {
                    script {
                        // Build the Docker image
                        sh "docker build -t angry/runner:latest ."
                    }
                }
            }

            stage('Stop and Remove Existing Container') {
                steps {
                    script {
                        // Stop and remove existing container
                        def containerName = 'runner'  // Replace with your container name

                        // Stop the container if it's running
                        sh "docker ps -q --filter 'name=${containerName}' | grep -q . && docker stop ${containerName}"

                        // Remove the container
                        sh "docker rm -f ${containerName} || true"  // Ignore error if the container does not exist
                    }
                }
            }

            stage('Run New Container') {
                steps {
                    script {
                        // Run the new container with the newly built image
                        def containerName = 'runner'  // Replace with your container name
                        def imageName = 'angry/runner'           // Replace with your desired image name
                        def imageTag = 'latest'                      // Replace with your desired tag

                        // Run the new container
                        sh "docker run -d --name ${containerName} -p 80:80 ${imageName}:${imageTag}"  // Adjust ports as needed
                    }
                }
            }
        }
}

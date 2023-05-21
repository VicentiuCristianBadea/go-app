pipeline {
    agent { 
        node {
            label 'docker-agent-python'
            }
      }
    triggers {
        pollSCM '* * * * *'
    }
    stages {
        stage('Build') {
            steps {
                echo "Building.."
                sh '''
                go build main.go
                go mod download
                '''
            }
        }
        stage('Test') {
            steps {
                echo "Testing.."
                sh '''
                kubectl run -it curl --rm --image=nginx:alpine sh
                curl web service:80/details
                '''
            }
        }
        stage('Deliver') {
            steps {
                echo 'Deliver....'
                sh '''
                echo "doing delivery stuff.."
                '''
            }
        }
    }
}
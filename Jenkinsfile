pipeline {
    agent any
    stages {
        stage("Verify tooling"){
            steps {
                sh 'docker version'
                sh 'docker info'
                sh 'docker compose version'
                sh 'curl --version'
                sh 'jq --version'
                sh 'go --version'
            }
        }
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
                echo "doing testing stuff.."
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
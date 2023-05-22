pipeline {
    agent {
    kubernetes {
      yaml '''
        apiVersion: v1
        kind: Pod
        spec:
          containers:
          - name: docker
            image: docker:latest
            command:
            - cat
            tty: true
            volumeMounts:
             - mountPath: /var/run/docker.sock
               name: docker-sock
          volumes:
          - name: docker-sock
            hostPath:
              path: /var/run/docker.sock    
        '''
    }
  }
    stages {
        stage("Verify tooling"){
            steps {
                container('docker'){
                    sh' sudo snap install go --classic'
                    sh 'docker version'
                    sh 'docker info'
                    sh 'docker compose version'
                    sh 'go --version'
                }
            }
        }
        stage('Build') {
            steps {
                echo "Building.."
                sh '''
                go mod download
                go build main.go
                docker compose build 
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
                docker compose push
                '''
            }
        }
    }
}
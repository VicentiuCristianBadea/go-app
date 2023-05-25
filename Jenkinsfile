pipeline {
    agent {
    kubernetes {
      yaml '''
        apiVersion: v1
        kind: Pod
        spec:
          containers:
          - name: golang
            image: golang:1.16
            command:
            -  cat
            tty: true
          - name: docker
            image: docker:latest
            command:
            - cat
            tty: true
            volumeMounts:
             - mountPath: /var/run/docker.sock
               name: docker-sock
            resources:
                requests:
                    cpu: "500m"
                limits:
                    cpu: "1000m"
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
                    sh 'docker version'
                    sh 'docker info'
                    sh 'docker compose version'
                }
                container('golang'){
                    sh 'go version'
                }
            }
        }
        stage('Build') {
            steps {
                echo "Building.."
                
                container('golang'){
                    sh '''
                        go mod download
                        go build main.go
                    '''
                }
                container('docker'){
                    sh '''
                        docker compose build
                    '''
                } 
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
                container('docker'){
                    withCredentials([usernamePassword(credentialsId: 'Jenkins', passwordVariable: 'DOCKER_PASS', usernameVariable: 'DOCKER_USER')]) {
                        sh 'echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin'
                        sh "docker login -u ${DOCKER_USER} -p ${DOCKER_PASS}"
                        sh 'docker compose push'
                        sh 'docker logout'
                    }
                }
            }
        }
    }
}
pipeline {
    agent any
    tools {
        go 'go-1.16'
    }
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = 0
    }
    stages {
        stage('Checkout') {
            steps {
                git(url: 'https://github.com/ponponda/go-rest-api', branch: 'master')
            }
        }
        stage('Build') {
            steps {
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
        stage('Release') {
            steps {
                sh 'done'
            }
        }
    }
}

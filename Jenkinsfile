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
        // stage('Build Image') {
        //     agent {
        //         dockerfile {
        //             filename 'Dockerfile'
        //             dir '.'
        //         }
        //     }
        //     steps {
        //         sh 'done'
        //     }
        // }
        stage('Deploy') {
            steps {
                sh 'heroku container:login'
                sh 'heroku container:push api -a todo-api'
                sh 'heroku container:release api -a todo-api'
            }
        }
    }
}

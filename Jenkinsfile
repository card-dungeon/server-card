pipeline {
    agent any
    tools {
        go 'go 1.18.3'
        dockerTool 'docker'
    }
    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        DOCKER_CREDENTIALS = credentials('dockerhub')
    }
    stages {
        stage("Build And Push Image") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'docker build . -t gmae199boy/server-card:latest'
                sh "docker login -u ${env.DOCKER_CREDENTIALS_USR} -p ${env.DOCKER_CREDENTIALS_PSW}"
                sh 'docker push gmae199boy/server-card:latest'
            }
        }

        stage('Remove Unused docker image') {
            steps{
                sh "docker rmi gmae199boy/server-card:latest"
            }
        }
    }
}
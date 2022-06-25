pipeline {
    // install golang 1.14 on Jenkins node
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
        // stage("unit-test") {
        //     steps {
        //         echo 'UNIT TEST EXECUTION STARTED'
        //         sh 'make unit-tests'
        //     }
        // }
        // stage("functional-test") {
        //     steps {
        //         echo 'FUNCTIONAL TEST EXECUTION STARTED'
        //         sh 'make functional-tests'
        //     }
        // }
        stage("Build And Push Image") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'docker build . -t gmae199boy/server-card:latest'
                sh "docker login -u ${env.DOCKER_CREDENTIALS_USR} -p ${env.DOCKER_CREDENTIALS_PSW}"
                sh 'docker push gmae199boy/server-card:latest'
            }
        }

        // stage('Docker') {         
        //     steps {                           
        //         script {
        //             // stage('Clone repository') {
        //             //     checkout scm
        //             // }

        //             stage('Build And Push image') {  
        //                 docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {                                                 
        //                     def app = docker.build("${env.DOCKER_CREDENTIALS_USR}/server-card")
        //                     app.push("latest")
        //                 }
        //             }                       
        //         }
        //     }
        // }
        stage('Remove Unused docker image') {
            steps{
                sh "docker rmi gmae199boy/server-card:latest"
            }
        }
    }
}
pipeline {
    // install golang 1.14 on Jenkins node
    agent any
    tools {
        go 'go 1.18.3'
        dockerTool 'docker'
    }
    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
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
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
            }
        }

        stage('Docker') {         
            environment {
                DOCKER_CREDENTIALS = credentials('dockerhub')
            }

            steps {                           
                // script {
                    node {
                        def app

                        // stage('Clone repository') {
                        //     checkout scm
                        // }

                        stage('Build And Push image') {  
                            docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {                                                 
                                app = docker.build("${env.DOCKER_CREDENTIALS_USR}/server-card")
                                app.push()
                            }
                        }              
                    }                 
                }
            }
        // }

    }
}